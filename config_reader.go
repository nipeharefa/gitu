package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type (
	// Route ...
	Route struct {
		Src     string            `json:"src"`
		Headers map[string]string `json:"headers"`
		Rewrite string            `json:"rewrite"`
	}
	// Config ...
	Config struct {
		Routes []Route `json:"routes"`
	}
)

// ReadConfig ...
func ReadConfig(filename string) *Config {

	c := &Config{}

	f, err := os.Open(filename)
	if err != nil {
		// logger.Error("e")
		logger.Fatal("S")
	}
	defer f.Close()

	b := &bytes.Buffer{}
	_, err = io.Copy(b, f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b.Bytes(), c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
