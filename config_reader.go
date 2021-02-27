package main

type (
	// Route ...
	Route struct {
		Src     string            `json:"src"`
		Headers map[string]string `json:"headers"`
	}
	// Config ...
	Config struct {
		Routes []Route `json:"routes"`
	}
)

func configReader() {}
