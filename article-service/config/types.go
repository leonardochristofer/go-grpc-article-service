package config

type Config struct {
	Application struct {
		ServiceName    string
		ServiceVersion string
		Env            string
		ServerPort     string
	}
	Logger  LoggerConfig
	DBRead  DBConfig
	DBWrite DBConfig
}

type LoggerConfig struct {
	Level           string
	Fulltimestamp   bool
	TimestampFormat string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type GrpcClientService struct {
	Host string
}
