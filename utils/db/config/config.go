package config

type Config struct {
	Mysql struct {
		Datasource string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}
