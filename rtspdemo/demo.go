package main

import (
	"fmt"
	"golangDemos/rtspdemo/rtsp"
	"net/http"
)

func main() {

	session := rtsp.NewSession()

	url := "rtsp://admin:lx123456@192.168.8.10:554/Streaming/Channels/701?transportmode=unicast"
	url = "rtsp://admin:lx123456@192.168.17.4:554/Streaming/Channels/101?transportmode=unicast"
	resp, _ := session.Describe(url, "")

	if resp.StatusCode == http.StatusUnauthorized {
		res := resp.Header.Get("Www-Authenticate")
		if res != "" {
			resp, _ = session.Describe(url, res)
		}

		fmt.Println(res)
	}

	var currStatus int = 0
	currCode := ""

	if resp == nil {
		currCode = "Connect Bad"
	} else {
		if resp.StatusCode == http.StatusOK { // status OK
			currStatus = 1 //ONLINE
		} else {
			if len(resp.Status) > 0 {
				currCode = resp.Status
			} else {
				// all info response is empty
				currCode = "Response Empty"
			}
		}
	}

	fmt.Println("currStatus:", currStatus, ",currCode:", currCode)

}
