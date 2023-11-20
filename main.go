package main

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
