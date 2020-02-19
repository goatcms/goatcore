package smtpmail_test

import (
	"os"
	"testing"

	"github.com/goatcms/goatcore/goatmail/smtpmail"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

type TestConfig struct {
	SenderConfig smtpmail.Config `json:"sender"`
	FromAddress  string          `json:"fromAddress"`
	ToAddress    string          `json:"toAddress"`
}

func LoadTestConfig(t *testing.T) (config *TestConfig) {
	var err error
	config = &TestConfig{}
	if err = goaterr.ToError(goaterr.AppendError(nil,
		InjectEnv("GOATCORE_TEST_SMTP_FROM_ADDRESS", &config.FromAddress),
		InjectEnv("GOATCORE_TEST_SMTP_TO_ADDRESS", &config.ToAddress),
		InjectEnv("GOATCORE_TEST_SMTP_SERVER", &config.SenderConfig.SMTPAddr),
		InjectEnv("GOATCORE_TEST_SMTP_USERNAME", &config.SenderConfig.AuthUsername),
		InjectEnv("GOATCORE_TEST_SMTP_PASSWORD", &config.SenderConfig.AuthPassword),
	)); err != nil {
		t.Skip(err.Error())
	}
	InjectEnv("GOATCORE_TEST_SMTP_IDENTITY", &config.SenderConfig.AuthIdentity) // <- is not required
	return config
}

func InjectEnv(name string, dest *string) (err error) {
	if *dest = os.Getenv(name); *dest == "" {
		return goaterr.Errorf("%v operating environment is required", name)
	}
	return nil
}
