package svc

import (
	"fmt"
	"k2edge/master/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DSN struct {
	username string
	password string
	host     string
	port     int
	db       string
}

func NewDSN(username, password, host string, port int, db string) DSN {
	return DSN{
		username: username,
		password: password,
		host:     host,
		port:     port,
		db:       db,
	}
}

func (d DSN) MySQL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=30s",
		d.username,
		d.password,
		d.host,
		d.port,
		d.db,
	)
}

type ServiceContext struct {
	Config config.Config
	DatabaseClient *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	databaseClient, err := gorm.Open(mysql.Open(NewDSN("root", "1234567890", "outlg.xyz", 3306, "k2edge").MySQL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DatabaseClient: databaseClient,
	}
}