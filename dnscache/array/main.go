package main

import (
	"fmt"
	"strconv"
)

// write a dns cache using LRU
// fixed size
// a dns cache is where you can go to find ip addy that's connected to domain
type dnsEntry struct {
	// google.com -> 76.345.234
	// google.com -> 76.345.112
	domain string
	// ip later could be a list of any number of ips
	ip string
}

func newDNSEntry(domain, ip string) dnsEntry {
	// TODO: validations
	if len(domain) == 0 || len(ip) == 0 {
		return dnsEntry{}
	}
	// TODO: regex validation on ip
	return dnsEntry{
		domain: domain,
		ip:     ip,
	}
}

func newLRU(size int) lru {
	return lru{
		list:   make([]dnsEntry, 0, size),
		dnsMap: make(map[string]int, size),
		size:   size,
	}
}

type lru struct {
	// set size
	size int
	// lru is composed of  list and a map
	// list for order
	list []dnsEntry
	// map for fast lookup
	dnsMap map[string]int // key is domain, val is index in list
}

func (c *lru) findIfExist(domainName string) dnsEntry {
	// check if already in cache
	index, found := c.dnsMap[domainName]
	if !found {
		// would do a lookup & add to cache
		return dnsEntry{}
	}
	// add this entry to front of cache
	c.add(c.list[index])
	return c.list[index]

}

func (c *lru) add(toAdd dnsEntry) {
	// check if exists
	if index, found := c.dnsMap[toAdd.domain]; found {
		if index == 0 {
			return
		}
		// remove from location in list
		c.list = append(c.list[0:index], c.list[index:]...)
	}

	// check size
	if len(c.dnsMap) > c.size {
		// remove last
		c.remove(c.list[len(c.list)-1].domain)
	}

	// add to list
	c.list = append([]dnsEntry{toAdd}, c.list[0:]...)
	c.updateIndexes(0, "add")
	// add will always be index 0
	c.dnsMap[toAdd.domain] = 0
}

func (c *lru) updateIndexes(position int, operation string) {
	switch operation {
	case "add":
		for k, v := range c.dnsMap {
			c.dnsMap[k] = v + 1
		}
	case "remove":
		for k, v := range c.dnsMap {
			if v > position {
				c.dnsMap[k] = v - 1
			}
		}
	}
}

func (c *lru) remove(domain string) {
	// find in map
	index, ok := c.dnsMap[domain]
	if !ok {
		// not actually in map/list
		return
	}
	// remove in map
	delete(c.dnsMap, domain)
	// remove in list
	c.list = append(c.list[0:index], c.list[index+1:]...)
	c.updateIndexes(index, "remove")
}

func main() {
	myCache := newLRU(5)
	domain := "my-domain"
	for i := 0; i < 4; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}

	fmt.Println(myCache)

	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "0")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "1")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "2")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "3")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "4")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "5")))
	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "6")))

	fmt.Println(myCache)

	myCache.remove(fmt.Sprint(domain, "3"))
	fmt.Println(myCache)

	fmt.Println(myCache.findIfExist(fmt.Sprint(domain, "3")))

}
