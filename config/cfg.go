package config

type Config struct {
	BoxConfig      BoxConfig
	EmailConfig    EmailConfig
	SongFolder     int
	CoverFolder    int
	DefaultPicture int
	PsqlInfo       string
	Url            string
}

type EmailConfig struct {
	SenderName     string
	SenderEmail    string
	AdminEmail     string
	DeveloperEmail string
	SmtpConfig 	   SmtpEmailConfig
}

type BoxConfig struct {
	ClientID     string
	ClientSecret string
	PrivateKey   string
	Token        string
}

type SmtpEmailConfig struct {
	ServerAddress string //e.g. smtp.gmail.com:587
	UserName      string
	Password      string
	Host          string // e.g. smtp.gmail.com
}
