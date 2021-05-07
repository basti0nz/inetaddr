package inetaddr

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var myIP net.IP
var err error

func init() {
	/* load test data */
	myIP, err = HttpCheck(HttpServer{"https://checkip.amazonaws.com/"})
	if err != nil {
		panic(fmt.Sprintf("Expected IP but has  %s", err))
	}
}

func TestInetAddr(t *testing.T) {
	assert.Equal(t, myIP.String(), InetAddr(BothMethod))
}

func TestIpAddrString(t *testing.T) {
	ip, err := IpAddrString(OnlyDNS)
	if err != nil {
		t.Errorf("Expected IP but has  %s", err)
	}
	assert.Equal(t, myIP.String(), ip)
}

func TestIpAddrIP(t *testing.T) {
	ip, err := IpAddrIP(OnlyHTTP)
	if err != nil {
		t.Errorf("Expected IP but has  %s", err)
	}
	assert.Equal(t, myIP, ip)

}
