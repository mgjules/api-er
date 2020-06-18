package config

import "github.com/spf13/viper"

// CasbinConfig holds casbin-related config
type CasbinConfig struct {
	Model  string `mapstructure:"model"`
	Policy string `mapstructure:"policy"`
}

// SecurityConfig holds security-related config
type SecurityConfig struct {
	PasswordSalt string `mapstructure:"password_salt"`
	StoreSecret  string `mapstructure:"store_secret"`
	SessionKey   string `mapstructure:"session_key"`
	CSRFSecret   string `mapstructure:"csrf_secret"`

	Casbin CasbinConfig `mapstructure:"casbin"`
}

// DBConfig holds db-related config
type DBConfig struct {
	Dialect  string
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

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

	HTTP     HTTPConfig     `mapstructure:"http"`
	DB       DBConfig       `mapstructure:"db"`
	Security SecurityConfig `mapstructure:"security"`
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
