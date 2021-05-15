package dcmd

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/testbase"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/goatnet"

	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"

	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/envs"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
)

func engineSimpleStory(t *testing.T, engine ocservices.Engine) {
	t.Parallel()
	var (
		err       error
		outBuffer = bytes.NewBuffer(make([]byte, 10000))
		errBuffer = bytes.NewBuffer(make([]byte, 10000))
		cwd       filesystem.Filespace
		io        app.IO
		envs      = envs.NewEnvironments()
	)
	if cwd, err = diskfs.NewFilespace("."); err != nil {
		t.Error(err)
		return
	}
	io = gio.NewIO(gio.IOParams{
		In: gio.NewAppInput(strings.NewReader(`
		set -x

		# test echo
		echo "echo-ok"
		
		#test pwd
		pwd=$(pwd)
		echo "cwd:$pwd"

		#test env
		echo "ENVVAR:$ENVVAR"

		# test mapping path
		if [[ -d /cwd/test ]]; then
			echo "/cwd/test:ok"
		else
			echo "/cwd/test:fail"
		fi

		`)),
		Out: gio.NewOutput(outBuffer),
		Err: gio.NewOutput(errBuffer),
		CWD: cwd,
	})
	if err = envs.SetAll(map[string]string{
		"ENVVAR": "ENVVALUE",
	}); err != nil {
		t.Error(err)
		return
	}
	if err = engine.Run(ocservices.Container{
		IO:         io,
		Image:      "docker.io/alpine",
		WorkDir:    "/cwd",
		Entrypoint: "sh",
		Envs:       envs,
		FSVolumes: map[string]ocservices.FSVolume{
			"/cwd/test": {
				Filespace: cwd,
			},
		},
		Privileged: false,
	}); err != nil {
		t.Error(err)
		return
	}
	output := outBuffer.String()
	if !strings.Contains(output, "echo-ok") {
		t.Errorf("Expected echo-ok \nOutput:%s", output)
		return
	}
	if !strings.Contains(output, "cwd:/cwd") {
		t.Errorf("Expected current working directory as /cwd \nOutput:%s", output)
		return
	}
	if !strings.Contains(output, "/cwd/test:ok") {
		t.Errorf("Expected mapped /cwd/test directore \nOutput:%s", output)
		return
	}
}

func enginePortStory(t *testing.T, engine ocservices.Engine) {
	t.Parallel()
	var (
		err       error
		io        app.IO
		freePort  int
		output    string
		scp       = scope.New(scope.Params{})
		outBuffer = bytes.NewBuffer(make([]byte, 10000))
		errBuffer = bytes.NewBuffer(make([]byte, 10000))
		cwd       filesystem.Filespace
	)
	if cwd, err = diskfs.NewFilespace("."); err != nil {
		t.Error(err)
		return
	}
	if freePort, err = goatnet.GetFreePort(); err != nil {
		t.Error(err)
		return
	}
	io = gio.NewIO(gio.IOParams{
		In:  gio.NewAppInput(strings.NewReader(``)),
		Out: gio.NewOutput(outBuffer),
		Err: gio.NewOutput(errBuffer),
		CWD: cwd,
	})
	go func() {
		engine.Run(ocservices.Container{
			IO:         io,
			Image:      "tutum/hello-world",
			Privileged: false,
			Ports: map[int]int{
				80: freePort,
			},
			Scope: scp,
		})
	}()
	if output, err = testbase.ReadURLLoop("http://localhost:"+strconv.Itoa(freePort), 20, time.Second/2); err != nil {
		scp.Kill()
		t.Error(err)
		return
	}
	expected := "Hello world!"
	if !strings.Contains(output, expected) {
		scp.Kill()
		t.Errorf("Expected output contains '%s' \nOutput:%s", expected, output)
		return
	}
	scp.Kill()
	scp.Close()
}

func engineEscapeEnvStory(t *testing.T, engine ocservices.Engine) {
	t.Parallel()
	var (
		err       error
		outBuffer = bytes.NewBuffer(make([]byte, 10000))
		errBuffer = bytes.NewBuffer(make([]byte, 10000))
		cwd       filesystem.Filespace
		io        app.IO
		envs      = envs.NewEnvironments()

		untructValue = `echo "$(echo '3')"`
	)
	if cwd, err = diskfs.NewFilespace("."); err != nil {
		t.Error(err)
		return
	}
	io = gio.NewIO(gio.IOParams{
		In: gio.NewAppInput(strings.NewReader(`
		set -x
		echo "$ENVVAR"
		`)),
		Out: gio.NewOutput(outBuffer),
		Err: gio.NewOutput(errBuffer),
		CWD: cwd,
	})
	if err = envs.SetAll(map[string]string{
		"ENVVAR": untructValue,
	}); err != nil {
		t.Error(err)
		return
	}
	if err = engine.Run(ocservices.Container{
		IO:         io,
		Image:      "docker.io/alpine",
		WorkDir:    "/cwd",
		Entrypoint: "sh",
		Envs:       envs,
	}); err != nil {
		t.Error(err)
		return
	}
	output := outBuffer.String()
	if !strings.Contains(output, untructValue) {
		t.Errorf("Expected %s \nOutput:%s", untructValue, output)
		return
	}
}
