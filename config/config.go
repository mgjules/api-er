package config

import "github.com/spf13/viper"

// CORSConfig holds cors-related config
type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

// HTTPConfig holds http-related config
type HTTPConfig struct {
	Host string
	Port string

	CORS CORSConfig `mapstructure:"cors"`
}

// Config holds app-wide config
type Config struct {
	Prod bool

	HTTP HTTPConfig `mapstructure:"http"`
}

// Load loads the config
func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config

	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
