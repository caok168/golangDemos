package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 根据IP和掩码获取网络
 */
func main() {
	name := "192.168.1.97"

	ip := net.ParseIP(name)

	mask := ip.DefaultMask()

	network := ip.Mask(mask)

	fmt.Fprintf(os.Stdout, "network: %s", network.String()) //192.168.1.0

	// ip:      192.168.1.97
	// mask :   255.255.255.0
	// network :192.168.1.0
}
