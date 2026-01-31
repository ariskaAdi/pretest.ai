package config

type Config struct {
	App    AppConfig
	Genkit Genkit
}

type AppConfig struct {
	Name string
	Port string
}

type Genkit struct {
	GoogleAIAPIKey string
	Port           string
	Environment    string
}

var Cfg Config