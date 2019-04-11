package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	test()

	time.Sleep(10 * time.Minute)

}

func test() {
	input := ""
	input = "rtsp://admin:lx123456@192.168.8.10:554/Streaming/Channels/101?transportmode=unicast"
	//input = "rtsp://admin:lx123456@192.168.8.10:554/Streaming/Channels/401?transportmode=unicast"

	cmdline := []string{
		"-rtsp_transport",
		"tcp",
		//"-fflags",
		//"+genpts",
		"-i",
		input,
		"-vcodec",
		"copy",
		"-acodec",
		"copy",
		"-f",
		"flv",
		"rtmp://192.168.8.12:1935/live/input222/112",
	}

	out, err := exec.Command("ffmpeg", cmdline...).CombinedOutput()
	fmt.Println("out:", string(out))
	if err != nil {
		fmt.Println("err,", err)

	}

	cmd := exec.Command("ffmpeg", cmdline...)

	//cmd.Process.Kill()

	cmd.Start()

	//cmd.Run()

	state := cmd.ProcessState.String()
	fmt.Println(state)

	fmt.Println(cmd)

	err = cmd.Wait()
	fmt.Println("err:", err)

	time.Sleep(1 * time.Minute)

	cmd.Process.Kill()
}
