package main

import (
	"embed"
	"errors"
	"os"
	"strconv"

	"github.com/open-function-computers-llc/server-run-inertia/server"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed dist/*
var dist embed.FS

func main() {
	port, url, err := processENV()
	if err != nil {
		panic("Trouble setting up server: " + err.Error())
	}

	s, err := server.New(port, url, dist)
	if err != nil {
		panic("Trouble setting up server: " + err.Error())
	}

	err = s.Serve()
	if err != nil {
		panic("Trouble serving from server: " + err.Error())
	}
}

func processENV() (int, string, error) {
	requiredENV := []string{
		"APP_PORT",
		"APP_URL",
	}

	for _, env := range requiredENV {
		check := os.Getenv(env)
		if check == "" {
			return 0, "", errors.New("missing env: " + env)
		}
	}

	webPort := os.Getenv("APP_PORT")
	portInt, err := strconv.Atoi(webPort)
	if err != nil {
		panic("Invalid APP_PORT: " + err.Error())
	}

	return portInt, os.Getenv("APP_URL"), nil
}
