package inetaddr

import (
	dns "github.com/Focinfi/go-dns-resolver"
)

var HttpServerS = []HttpServer{
	{"https://ifconfig.me/"},
	{"https://icanhazip.com/"},
	{"https://ident.me/"},
	{"https://ipecho.net/plain"},
	{"https://checkip.amazonaws.com/"},
	{"https://wgetip.com/"},
	{"https://ip.tyk.nu/"},
	{"https://bot.whatismyipaddress.com/"},
	{"https://ipof.in/txt"},
	{"https://l2.io/ip"},
	{"https://eth0.me/"},
}

var DnsServerS = []DnsServer{
	{"resolver3.opendns.com", "myip.opendns.com", dns.TypeA},
	{"resolver4.opendns.com", "myip.opendns.com", dns.TypeA},
	{"resolver2.opendns.com", "myip.opendns.com", dns.TypeA},
	{"resolver1.opendns.com", "myip.opendns.com", dns.TypeA},
	{"ns1-1.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	{"za.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	{"zb.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	{"zc.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	{"zd.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	{"ns1.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
	{"ns2.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
	{"ns3.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
}
