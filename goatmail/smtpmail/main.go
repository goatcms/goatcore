package smtpmail

type Config struct {
	SmtpAddr     string `json:"smtpAddr"`
	AuthUsername string `json:"authUsername"`
	AuthPassword string `json:"authPassword"`
	AuthIdentity string `json:"authIdentity"`
}
