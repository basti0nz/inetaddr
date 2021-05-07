package inetaddr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDnsCheck(t *testing.T) {
	myIP, err := HttpCheck(HttpServer{"https://checkip.amazonaws.com/"})
	if err != nil {
		t.Errorf("Expected IP but has  %s", err)
	}

	for i := range DnsServerS {
		ip, err := DnsCheck(DnsServerS[i])
		if err != nil {
			t.Error("error: ", DnsServerS[i].Server, " has error ", err)
		}
		assert.Equal(t, myIP, ip)
	}

}
