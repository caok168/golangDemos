package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 查看主机服务器（service）占用的端口，这些服务，都是tcp或者udp的
 */
func main() {
	port, err := net.LookupPort("tcp", "telnet") //查看telnet服务器使用的端口

	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}

	fmt.Fprintf(os.Stdout, "telnet port: %d", port)
}
