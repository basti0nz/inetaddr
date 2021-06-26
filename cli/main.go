package main

import (
	"fmt"
	"github.com/basti0nz/inetaddr"
)

const (
	OnlyDNS = iota
	OnlyHTTP
	BothMethod
)

func main() {
	ip, err := inetaddr.IpAddrString(OnlyDNS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	ipIP, err := inetaddr.IpAddrIP(OnlyHTTP)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipIP.String())

	ip, err = inetaddr.IpAddrString(BothMethod)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
