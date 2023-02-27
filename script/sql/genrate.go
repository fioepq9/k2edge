package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
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

func (d DSN) Postgresql() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		d.host,
		d.username,
		d.password,
		d.db,
		d.port,
	)
}

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./dao", // output directory, default value is ./query
		Mode:              gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// Initialize a *gorm.DB instance
	db, err := gorm.Open(mysql.Open(NewDSN("root", "1234567890", "outlg.xyz", 3306, "k2edge").MySQL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
