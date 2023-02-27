package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Postgresql PostgresqlConf `json:"Postgresql"`
}

type PostgresqlConf struct {
	Host     string `json:"Host"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DBName   string `json:"DBName"`
	Port     int    `json:"Port"`
}
