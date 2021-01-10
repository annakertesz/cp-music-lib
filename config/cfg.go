package config

type Config struct {
	BoxConfig      BoxConfig
	SengridConfig  SengridConfig
	SmtpConfig     SmtpEmailConfig
	SongFolder     int
	CoverFolder    int
	DefaultPicture int
	PsqlInfo       string
	Url            string
}

type SengridConfig struct {
	SengridAPIKey  string
	SenderName     string
	SenderEmail    string
	AdminEmail     string
	DeveloperEmail string
}

type BoxConfig struct {
	ClientID     string
	ClientSecret string
	PrivateKey   string
	Token        string
}

type SmtpEmailConfig struct {
	ServerAddress string // e.g. smtp.gmail.com:587
	UserName      string // gmail address
	Password      string // gmail password
	Host          string // e.g. smtp.gmail.com
	ToEmail       string // email address to send emails to
	FromEmail     string // email address that this email is sent from.
	// can be set to something else, but this might trigger spam filters.
}
