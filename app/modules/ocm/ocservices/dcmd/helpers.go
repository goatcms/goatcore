package dcmd

import (
	"io"
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// InitSequence return init sequence reader
func InitSequence(envs commservices.Environments) (reader io.Reader, err error) {
	var (
		initCode = "\nset -e\nset +x\n"
		eofTag   = "EOF" + varutil.RandString(10, varutil.UpperAlphaBytes)
	)
	if envs == nil {
		return strings.NewReader(initCode), nil
	}
	for key, value := range envs.All() {
		initCode += key + "=$(cat <<" + eofTag + "\n" + value + "\n" + eofTag + "\n)\n"
		initCode += "export " + key + "\n"
	}
	sshCert := envs.SSHCert()
	if sshCert.Public != "" || sshCert.Secret != "" {
		if sshCert.Public == "" {
			return nil, goaterr.Errorf("SSHCert: Public key is required")
		}
		if sshCert.Secret == "" {
			return nil, goaterr.Errorf("SSHCert: Secret key is required")
		}
		initCode += "mkdir -p ~/.ssh\n"
		initCode += "cat <<" + eofTag + " >> ~/.ssh/id_rsa.pub \n" + sshCert.Public + "\n" + eofTag + "\n"
		initCode += "chmod 400 ~/.ssh/id_rsa.pub\n"
		initCode += "cat <<" + eofTag + " >> ~/.ssh/id_rsa \n" + sshCert.Secret + "\n" + eofTag + "\n"
		initCode += "chmod 400 ~/.ssh/id_rsa\n"
	}
	return strings.NewReader(initCode), nil
}

// MapVolumens return docker style formatted volumens
func MapVolumens(volumens map[string]ocservices.FSVolume) (args []string, err error) {
	const (
		rowSieze = 2
	)
	var (
		all        = make([]string, len(volumens)*rowSieze)
		i          = 0
		fs         filesystem.LocalFilespace
		volumePath string
		ok         bool
	)
	for containerPath, vfVolume := range volumens {
		if fs, ok = vfVolume.Filespace.(filesystem.LocalFilespace); !ok {
			return nil, goaterr.Errorf("Open container services support only filesystem.LocalFilespace as volume and take %T", vfVolume.Filespace)
		}
		if volumePath, err = varutil.ReduceAbsPath(vfVolume.Path); err != nil {
			return nil, err
		}
		if containerPath, err = varutil.ReduceAbsPath(containerPath); err != nil {
			return nil, err
		}
		volumePath = fs.LocalPath() + volumePath
		all[i*rowSieze] = "-v"
		all[i*rowSieze+1] = volumePath + ":/" + containerPath
		i++
	}
	return all, nil
}

// MapPorts return docker style formatted ports
func MapPorts(ports map[int]int) (args []string, err error) {
	const (
		rowSieze = 2
	)
	var (
		all = make([]string, len(ports)*rowSieze)
		i   = 0
	)
	for containerPort, hostPort := range ports {
		all[i*rowSieze] = "-p"
		all[i*rowSieze+1] = strconv.Itoa(hostPort) + ":" + strconv.Itoa(containerPort)
		i++
	}
	return all, nil
}
