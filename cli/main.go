package main

import (
	"log"
	"os"

	"github.com/Layr-Labs/incredible-squaring-avs/cli/actions"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{config.ConfigFileFlag}
	app.Commands = []cli.Command{
		{
			Name:    "register-operator-with-eigenlayer",
			Aliases: []string{"rel"},
			Usage:   "registers operator with eigenlayer (this should be called via eigenlayer cli, not plugin, but keeping here for convenience for now)",
			Action:  actions.RegisterOperatorWithEigenlayer,
		},
		{
			Name:    "deposit-into-strategy",
			Aliases: []string{"d"},
			Usage:   "deposit tokens into a strategy",
			Action:  actions.DepositIntoStrategy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "strategy-addr",
					Usage:    "Address of Strategy contract to deposit into",
					Required: true,
				},
				cli.StringFlag{
					Name:     "amount",
					Usage:    "amount of tokens to deposit into strategy",
					Required: true,
				},
			},
		},
		{
			Name:    "register-operator-with-avs",
			Aliases: []string{"r"},
			Usage:   "registers bls keys with pubkey-compendium, opts into slashing by avs service-manager, and registers operators with avs registry",
			Action:  actions.RegisterOperatorWithAvs,
		},
		{
			Name:    "deregister-operator-with-avs",
			Aliases: []string{"d"},
			Action: func(ctx *cli.Context) error {
				log.Fatal("Command not implemented.")
				return nil
			},
		},
		{
			Name:    "print-operator-status",
			Aliases: []string{"s"},
			Usage:   "prints operator status as viewed from incredible squaring contracts",
			Action:  actions.PrintOperatorStatus,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
