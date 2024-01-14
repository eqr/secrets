package cmd

import (
	"context"
	"fmt"
	"github.com/eqr/secrets/app/config"
	"github.com/eqr/secrets/app/server"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.InitConfig(ConfigPath)
		srv, err := server.New(cfg)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Fatal(srv.Start())
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var ConfigPath string

func Build() {
	RootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "./config.yml", "path to the configuration file")
}
