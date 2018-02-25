package smtpmail

// Config provide smtp mail configuration
type Config struct {
	// SMTPAddr is SMTP server address
	SMTPAddr string `json:"smtpAddr"`
	// AuthUsername is username for SMTP server
	AuthUsername string `json:"authUsername"`
	// AuthPassword is password for SMTP server
	AuthPassword string `json:"authPassword"`
	// AuthIdentity is identity for SMTP server
	AuthIdentity string `json:"authIdentity"`
}
