package dcmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"

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
			"/cwd/test": ocservices.FSVolume{
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
