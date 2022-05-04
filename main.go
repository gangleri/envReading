package main

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v6"
)

type appConfig struct {
	Home         string   `env:"HOME"`
	Hosts        []string `env:"HOSTS" envSeparator:","`
	IsProduction bool     `env:"PRODUCTION"` // env var can be 1 or true
	Endpoint     url.URL  `env:"ENDPOINT"`
}

func main() {
	c := appConfig{}
	if err := env.Parse(&c); err != nil {
		panic(err)
	}
	fmt.Println(c.Home)
	fmt.Println(c.Hosts[0])
	fmt.Println(c.Hosts)
	fmt.Println(c.IsProduction)
	fmt.Printf("%+v\n", c.Endpoint)
}
