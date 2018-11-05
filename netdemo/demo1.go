package main

import (
	"fmt"
	"net"
)

func main() {

	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIpNet bool
		err     error
		ipv4    string
	)

	if addrs, err = net.InterfaceAddrs(); err != nil {
		fmt.Println(err)
		return
	}
	for _, addr = range addrs {
		// 这个网络地址是IP地址：ipv4,ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				fmt.Println("ipv4:", ipv4)
			}
		}
	}
}
