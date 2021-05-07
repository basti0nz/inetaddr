package inetaddr

import (
	"net"

	dns "github.com/Focinfi/go-dns-resolver"
)

type DnsServer struct {
	Server     string
	Record     string
	RecordType dns.QueryType
}

type HttpServer struct {
	Server string
}

type IPAddrResolver interface {
	IpAddrString() (string, error)
	IpAddrIP() (net.IP, error)
}
