package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
    viper.AutomaticEnv()

    viper.SetDefault("APP_PORT", "8080")
    
    if err := viper.ReadInConfig(); err != nil {
        log.Println("config file not found, using env only")
    }

    Cfg = Config{
        App: AppConfig{
            Name: viper.GetString("APP_NAME"),
            Port: viper.GetString("APP_PORT"),
        },
        Genkit: Genkit{
            GoogleAIAPIKey: viper.GetString("GOOGLE_AI_API_KEY"),
            Port:           viper.GetString("GENKIT_PORT"),
            Environment:    viper.GetString("ENV"),
        },
    }
}

//         GoogleAIAPIKey: apiKey,
//         Port:           port,
//         Environment:    env,
//     }, nil
// }