package config

type Config struct {
	BoxConfig      BoxConfig
	SengridConfig  SengridConfig
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
	ServerAddress string //e.g. smtp.gmail.com:587
	UserName      string
	Password      string
	Host          string // e.g. smtp.gmail.com
}
