package postgres

import "fmt"

type Config struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
	Retries  uint8
}

func NewDefaultConfig() *Config {
	return &Config{
		Host:     "db",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		Retries:  3,
	}
}

func (cfg *Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)
}
