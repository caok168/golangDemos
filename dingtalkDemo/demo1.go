package main

import (
	"fmt"
	"golangDemos/dingtalkDemo/dingrobot"
)

func main() {
	robot := dingrobot.New("49541b9cc490f42ba199832469a5c3d93ad38d8dcdbadcfc601332bb4b926c11")

	//robot.Text("Hi")
	//robot.Markdown("**HI**", "World")
	//robot.Link("Google", "Google homepage ", "https://www.google.com", "")

	err := robot.AtMobiles("18613875370").Text("test at someone")
	if err != nil {
		fmt.Println(err.Error())
	}
}
