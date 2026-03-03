package main

import (
	"github.com/robfig/cron/v3"
	"fmt"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("hello")
	})

	c.Start()

	select {}
}