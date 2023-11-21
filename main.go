package main

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/bootstrap"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()
	bootstrap.NewApp().Execute()
}
