package inetaddr

import (
  dns "github.com/Focinfi/go-dns-resolver"
  "log"
)

func main() {
  domains := []string{"google.com", "twitter.com"}
  types := []dns.QueryType{dns.TypeA, dns.TypeNS, dns.TypeMX, dns.TypeTXT}

  // Set timeout and retry times
  dns.Config.SetTimeout(uint(2))
  dns.Config.RetryTimes = uint(4)

  // Simple usage
  if results, err := dns.Exchange("google.com", "119.29.29.29", dns.TypeA); err == nil {
    for _, r := range results {
      log.Println(r.Record, r.Type, r.Ttl, r.Priority, r.Content)
    }
  } else {
    log.Fatal(err)
  }

  // Create and setup resolver with domains and types
  resolver := dns.NewResolver("119.29.29.29")
  resolver.Targets(domains...).Types(types...)
  // Lookup
  res := resolver.Lookup()

  //res.ResMap is a map[string]*ResultItem, key is the domain
  for target := range res.ResMap {
    log.Printf("%v: \n", target)
    for _, r := range res.ResMap[target] {
      log.Println(r.Record, r.Type, r.Ttl, r.Priority, r.Content)
    }
  }
}