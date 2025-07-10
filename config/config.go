package config

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	DSN string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig() *Config {
	return &Config{
		MySQL: MySQLConfig{
			DSN: "root:1234567890@tcp(127.0.0.1:3306)/TestGo?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Redis: RedisConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}
}
