package main

import (
	"inetaddr"
	"log"
)

func main() {
	for i := range inetaddr.DnsServerS {
		ip, err := inetaddr.DnsCheck(inetaddr.DnsServerS[i])
		if err != nil {
			log.Println("error: ", inetaddr.DnsServerS[i].Server, " has error ", err)
		}
		log.Println(inetaddr.DnsServerS[i].Server, ip.String())
	}

	for j := range inetaddr.HttpServerS {
		ip, err := inetaddr.HttpCheck(inetaddr.HttpServerS[j])
		if err != nil {
			log.Println("error: ", inetaddr.HttpServerS[j].Server, " has error ", err)
		}
		log.Println(inetaddr.HttpServerS[j].Server, ip.String())
	}
}
