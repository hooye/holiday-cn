package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "holiday"
	app.Usage = "Obtains Chinese Holidays"

	// 调用单独的函数来设置 Commands, Flags 和 Before
	app.Commands = []*cli.Command{
		startcmd(),
	}

	// app.Flags = setFlags()
	// app.Before = checkArgs()

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
