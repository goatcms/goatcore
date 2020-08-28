package dcmd

import (
	"testing"

	"github.com/goatcms/goatcore/varutil"

	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"

	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/filesystem"
)

func TestMapVolumensSuccess(t *testing.T) {
	var (
		args     []string
		err      error
		volumens = map[string]ocservices.FSVolume{}
		fs       filesystem.Filespace
		expected string
	)
	if fs, err = diskfs.NewFilespace("/tmp/t1"); err != nil {
		t.Error(err)
		return
	}
	volumens["/path/1/in/container"] = ocservices.FSVolume{
		Filespace: fs,
		Path:      "pathone",
	}
	if fs, err = diskfs.NewFilespace("/tmp/t2"); err != nil {
		t.Error(err)
		return
	}
	volumens["/path/2/in/container"] = ocservices.FSVolume{
		Filespace: fs,
		Path:      "pathtwo",
	}
	if args, err = MapVolumens(volumens); err != nil {
		t.Error(err)
		return
	}
	if args[0] != "-v" || args[2] != "-v" {
		t.Errorf("Expected args[0] and arg[2] equals to -v and take '%s' and '%s'", args[0], args[2])
	}
	expected = "/tmp/t1pathone:/path/1/in/container"
	if !varutil.IsArrContainStr(args, expected) {
		t.Errorf("Expected %s in arguments: %v", expected, args)
	}
	expected = "/tmp/t2pathtwo:/path/2/in/container"
	if !varutil.IsArrContainStr(args, expected) {
		t.Errorf("Expected %s in arguments: %v", expected, args)
	}
}
