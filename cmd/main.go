package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/signshine/give-a-sign/api/handlers/http"
	"github.com/signshine/give-a-sign/app"
	"github.com/signshine/give-a-sign/config"
)

var configPath = flag.String("config", "config.json", "configuration file path")

func main() {
	flag.Parse()
	godotenv.Load(".env")

	if v := os.Getenv("CONFIG_FILE"); len(v) > 0 {
		*configPath = v
	}

	cfg := config.MustReadConfig(*configPath)

	app := app.NewMustApp(cfg)

	log.Fatal(http.Run(app))
}
