package svc

import (
	"context"
	"fmt"
	"k2edge/dao"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/docker/docker/client"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	DockerClient   *client.Client
	DAO            *dao.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	_, err = dockerCli.Ping(context.TODO())
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		c.Postgresql.Host,
		c.Postgresql.User,
		c.Postgresql.Password,
		c.Postgresql.DBName,
		c.Postgresql.Port,
	)))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		DockerClient:   dockerCli,
		DAO:            dao.Use(db),
	}
}
