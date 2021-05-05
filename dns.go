package inetaddr

import (
	"fmt"
	"net"

	dns "github.com/Focinfi/go-dns-resolver"
)

func DnsCheck(record DnsServer) (net.IP, error) {
	// Set timeout and retry times
	dns.Config.SetTimeout(uint(2))
	dns.Config.RetryTimes = uint(4)

	results, err := dns.Exchange(record.Record, fmt.Sprintf("%s%s", record.Server, ":53"), record.RecordType)
	if err != nil {
		return nil, fmt.Errorf("Response error: %s", err)
	}

	for _, r := range results {
		ip := net.ParseIP(r.Content)
		if ip == nil {
			return nil, fmt.Errorf("Response error: not ip address %s", ip.String())
		}
		return ip, nil
	}
	return nil, fmt.Errorf("Error: unknown error")
}
