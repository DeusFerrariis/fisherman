package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	watchCommand := cli.Command{
		Name:   "watch",
		Action: WatchAction,
	}

	app := &cli.App{
		Name:     "fisherman",
		Usage:    "watch a file or directory and then execute a command after",
		Commands: []*cli.Command{&watchCommand},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func WatchAction(cCtx *cli.Context) error {
	if cCtx.NArg() < 2 {
		fmt.Println("ERR: path and command are needed.")
		return nil
	}

	watchPath := cCtx.Args().Get(0)
	watchCommand := cCtx.Args().Get(1)

	color.Green("\n🎣 FISHERMAN:")
	fmt.Println("\tWatching - \"" + watchPath + "\" for changes\n")

	wd := WatchDirectory{
		Path:         watchPath,
		LastChecked:  time.Now(),
		SleepSeconds: 1,
	}

	if err := wd.Watch(func() {
		if err := executeCommand(watchCommand); err != nil {
			fmt.Println(err)
		}

		fmt.Println("")
	}); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}

func executeCommand(command string) error {
	yellow := color.New(color.FgYellow)
	fmt.Println("Executing \"" + yellow.Sprintf(command) + "\" now")
	fmt.Println("\nOutput Caught 🎣:\n---")
	split := strings.Split(command, " ")

	cmd := exec.Command(split[0], split[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	fmt.Println("---\n")

	return err
}
