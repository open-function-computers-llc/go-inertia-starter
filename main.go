package main

import (
	"embed"
	"errors"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/go-inertia-starter/server"
)

//go:embed dist/*
var dist embed.FS

// version can be changed by build flags
var Version = "latest-dev"

func main() {
	// fmt.Println("KCE Teacher App - v: " + Version)
	err := verifyVaildEnv()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	s := server.NewServer(dist, os.Getenv("APP_URL"))

	err = s.Serve()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func verifyVaildEnv() error {
	requiredKeys := []string{
		"APP_PORT",
		"APP_URL",
		"APP_ENV",
	}

	for _, key := range requiredKeys {
		if os.Getenv(key) == "" {
			return errors.New("missing env: " + key)
		}
	}

	requiredKeysThatAreInts := []string{
		"APP_PORT",
	}
	for _, key := range requiredKeysThatAreInts {
		_, err := strconv.Atoi(os.Getenv(key))
		if err != nil {
			return err
		}
	}

	return nil
}
