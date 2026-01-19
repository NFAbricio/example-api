package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Enviromnents struct {
	StripKey       string `mapstructure:"STRIPE_KEY"`
	UserRoleSecret string `mapstructure:"USER_ROLE_SECRET"`

	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
}

func LoadEnvs() (*Enviromnents, error) {
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error to load file: %w", err)
		}
		viper.AutomaticEnv()

		env := &Enviromnents{}
		if err := viper.Unmarshal(env); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		return env, nil
	}

	return nil, fmt.Errorf("file .env not found") //?
}