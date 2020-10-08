package config


type Config struct {
	BoxConfig BoxConfig
	SengridConfig SengridConfig
	SongFolder int
	CoverFolder int
	PsqlInfo string
	Url string
}

type SengridConfig struct {
	SengridAPIKey string
	SenderName string
	SenderEmail string
	AdminEmail string
	DeveloperEmail string
}

type BoxConfig struct {
	ClientID string
	ClientSecret string
	PrivateKey string
	Token string
}