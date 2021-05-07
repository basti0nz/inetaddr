package main

import (
	"fmt"
	"inetaddr"
)

func main() {
	ip, err := inetaddr.IpAddrString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	ipIP, err := inetaddr.IpAddrIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipIP.String())
}
