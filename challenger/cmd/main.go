package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/incredible-squaring-avs/challenger"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "credible-squaring-challenger"
	app.Usage = "Credible Squaring Challenger"
	app.Description = "Service that challenges wrong response to the task."

	app.Action = challengerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func challengerMain(ctx *cli.Context) error {

	log.Println("Initializing Challenger...")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	chal, err := challenger.NewChallenger(config)
	if err != nil {
		return err
	}

	err = chal.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}
