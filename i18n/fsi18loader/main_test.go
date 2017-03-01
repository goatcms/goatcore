package fsi18loader

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/i18n/i18mem"
)

const (
	enJSON = `{
		"en": {
			"welcome": "Welcome"
		}
	}`
	plJSON = `{
		"pl": {
			"welcome": "Witaj"
		}
	}`
	enFormJSON = `{
		"en": {
			"form": {
				"validator": {
					"min_length": "The minimum length of the field is %v characters"
				}
			}
		}
	}`
	plFormJSON = `{
		"pl": {
			"form": {
				"validator": {
					"min_length": "Minimalna długość pola to %v znaków"
				}
			}
		}
	}`

	PLTranslate = "pl.form.validator.min_length"
	ENTranslate = "en.form.validator.min_length"
)

func TestSimpleLoad(t *testing.T) {
	i18 := i18mem.NewI18N()

	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("pl.json", []byte(plJSON), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("en.json", []byte(enJSON), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("forms/enform.json", []byte(enFormJSON), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("forms/plform.json", []byte(plFormJSON), 0777); err != nil {
		t.Error(err)
		return
	}

	if err = Load(fs, "./", i18, nil); err != nil {
		t.Error(err)
		return
	}

	pl, err := i18.Translate(PLTranslate, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if "Minimalna długość pola to 2 znaków" != pl {
		t.Errorf("incorrect translation for key %s (%s)", PLTranslate, pl)
		return
	}
	en, err := i18.Translate(ENTranslate, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if "The minimum length of the field is 2 characters" != en {
		t.Errorf("incorrect translation for key %s (%s)", ENTranslate, en)
		return
	}
}
