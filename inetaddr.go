package inetaddr

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

var wg sync.WaitGroup

type kv struct {
	Key   string
	Value int
}

func checkerDns(srv DnsServer, ch chan net.IP) {
	defer wg.Done()
	ip, err := DnsCheck(srv)
	if err != nil {
		return
	}
	ch <- ip
}

func checkerHttp(srv HttpServer, ch chan net.IP) {
	defer wg.Done()
	ip, err := HttpCheck(srv)
	if err != nil {
		return
	}
	ch <- ip
}

func InetAddr() string {
	messages := make(chan net.IP, len(HttpServerS)+len(DnsServerS))
	var sortIP []kv
	res := make(map[string]int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := range HttpServerS {
			wg.Add(1)
			go checkerHttp(HttpServerS[j], messages)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range DnsServerS {
			wg.Add(1)
			go checkerDns(DnsServerS[i], messages)
		}
	}()

	wg.Wait()
	close(messages)
	for range messages {
		ip := <-messages
		res[ip.String()] = res[ip.String()] + 1
	}
	for k, v := range res {
		sortIP = append(sortIP, kv{k, v})
	}
	sort.Slice(sortIP, func(i, j int) bool {
		return sortIP[i].Value > sortIP[j].Value
	})

	return sortIP[0].Key
}

func IpAddrString() (string, error) {
	strIP := InetAddr()
	ip := net.ParseIP(strIP)
	if ip == nil {
		return "", fmt.Errorf("Response error: not ip address %s", InetAddr())
	}
	return strIP, nil
}

func IpAddrIP() (net.IP, error) {
	ip := net.ParseIP(InetAddr())
	if ip == nil {
		return nil, fmt.Errorf("Response error: not ip address %s", InetAddr())
	}
	return ip, nil
}
