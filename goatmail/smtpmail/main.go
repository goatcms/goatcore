package smtpmail

type Config struct {
	SmtpAddr     string `json:"smtpAddr"`
	AuthUsername string `json:"authUsername"`
	AuthPassword string `json:"authPassword"`
	AuthHost     string `json:"authHost"`
	AuthIdentity string `json:"authIdentity"`
}
