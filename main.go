package main

import (
	"flag"
	"log"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var local bool

func init() {
	flag.BoolVar(&local, "local", true, "run service on local")
	flag.Parse()
}

func main() {
	if local {
		if err := godotenv.Load(); err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	e := echo.New()

}
