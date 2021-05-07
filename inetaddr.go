package inetaddr

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

var wg sync.WaitGroup

const (
	OnlyDNS = iota
	OnlyHTTP
	BothMethod
)

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

func InetAddr(flagCheckType int) string {
	messages := make(chan net.IP, len(HttpServerS)+len(DnsServerS))
	var sortIP []kv
	res := make(map[string]int)

	if flagCheckType == OnlyHTTP || flagCheckType == BothMethod {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range HttpServerS {
				wg.Add(1)
				go checkerHttp(HttpServerS[j], messages)
			}
		}()

	}
	if flagCheckType == OnlyDNS || flagCheckType == BothMethod {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range DnsServerS {
				wg.Add(1)
				go checkerDns(DnsServerS[i], messages)
			}
		}()
	}

	wg.Wait()
	close(messages)
	for range messages {
		ip := <-messages
		res[ip.String()] = res[ip.String()] + 1
	}

	if len(res) == 0 {
		return ""
	}

	for k, v := range res {
		sortIP = append(sortIP, kv{k, v})
	}
	sort.Slice(sortIP, func(i, j int) bool {
		return sortIP[i].Value > sortIP[j].Value
	})

	if len(sortIP) == 0 {
		return ""
	}

	return sortIP[0].Key
}

func IpAddrString(flagCheckType int) (string, error) {
	strIP := InetAddr(flagCheckType)
	ip := net.ParseIP(strIP)
	if ip == nil {
		return "", fmt.Errorf("Response error: There is not ip address")
	}
	return strIP, nil
}

func IpAddrIP(flagCheckType int) (net.IP, error) {
	ip := net.ParseIP(InetAddr(flagCheckType))
	if ip == nil {
		return nil, fmt.Errorf("Response error: There is not ip address ")
	}
	return ip, nil
}
