package main

import (
	"github.com/eqr/secrets/app/cmd"
	"github.com/eqr/secrets/app/config"
)

var cfg *config.Config

func init() {
	cmd.Build()
}

func main() {
	cmd.Execute()
}
