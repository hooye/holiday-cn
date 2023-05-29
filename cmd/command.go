package main

import (
	"fmt"

	"github.com/hooye/holiday-cn/pkg/test"
	"github.com/urfave/cli/v2"
)

func startcmd() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "启动",
		Flags: []cli.Flag{
			// 参数
			// &cli.StringFlag{
			// 	Name:  "test",
			// 	Usage: "test task",
			// },

			// bool 参数 默认为 false
			&cli.BoolFlag{
				Name:  "t",
				Usage: "test, ",
			},
		},
		Action: func(c *cli.Context) error {
			test := c.Bool("t")
			StartTask(test)
			return nil
		},
	}
}

func StartTask(t bool) {

	if t {
		test.LocalTest()
	} else {
		fmt.Println("制作中")
	}

}
