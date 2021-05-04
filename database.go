package inetaddr

import (
	dns "github.com/Focinfi/go-dns-resolver"
)

var HttpServerS = []HttpServer{
	HttpServer{"https://ifconfig.me/"},
	HttpServer{"https://icanhazip.com/"},
	HttpServer{"https://ident.me/"},
	HttpServer{"https://ipecho.net/plain"},
	HttpServer{"https://checkip.amazonaws.com/"},
	HttpServer{"https://wgetip.com/"},
	HttpServer{"https://ip.tyk.nu/"},
	HttpServer{"https://corz.org/ip"},
	HttpServer{"https://bot.whatismyipaddress.com/"},
	HttpServer{"https://ipof.in/txt"},
	HttpServer{"https://l2.io/ip"},
	HttpServer{"https://eth0.me/"},
}

var DnsServerS = []DnsServer{
	DnsServer{"resolver3.opendns.com", "myip.opendns.com", dns.TypeA},
	DnsServer{"resolver4.opendns.com", "myip.opendns.com", dns.TypeA},
	DnsServer{"resolver2.opendns.com", "myip.opendns.com", dns.TypeA},
	DnsServer{"resolver1.opendns.com", "myip.opendns.com", dns.TypeA},
	DnsServer{"ns1-1.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	DnsServer{"za.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	DnsServer{"zb.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	DnsServer{"zc.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	DnsServer{"zd.akamaitech.net", "whoami.akamai.net", dns.TypeA},
	DnsServer{"ns1.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
	DnsServer{"ns2.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
	DnsServer{"ns3.google.com", "o-o.myaddr.l.google.com", dns.TypeTXT},
}
