package main

//go get -u github.com/robfig/cron
import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	crontest()
}

func crontest() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	c.Start()

	select {}
}
