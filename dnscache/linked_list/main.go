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
	ip   string
	prev *dnsEntry
	next *dnsEntry
}

func newDNSEntry(domain, ip string) *dnsEntry {
	// // TODO: validations
	if len(domain) == 0 || len(ip) == 0 {
		return &dnsEntry{}
	}
	// TODO: regex validation on ip
	return &dnsEntry{
		domain: domain,
		ip:     ip,
	}
}

func newLRU(size int) *lru {
	head := &dnsEntry{domain: "head-placeholder"}
	tail := &dnsEntry{domain: "tail-placeholder"}
	head.next = tail
	tail.prev = head
	dnsMap := make(map[string]*dnsEntry)
	dnsMap[head.domain] = head
	dnsMap[tail.domain] = tail

	return &lru{
		dnsMap: dnsMap,
		size:   size,
		head:   head,
		tail:   tail,
	}
}

type lru struct {
	// set size
	size int
	// head node
	head *dnsEntry
	// end
	tail *dnsEntry

	// lru is composed of a map of doubly-linked nodes
	// doubly linked for ease of deletion & addition
	// map for fast lookup
	// the order is in the prev & next aspects of a node
	dnsMap map[string]*dnsEntry // key is domain, val is node
}

func (c *lru) findIfExist(domainName string) *dnsEntry {
	// check if already in cache
	entry, found := c.dnsMap[domainName]
	if !found {
		// would do a lookup & add to cache
		// log.Println("miss")
		return &dnsEntry{}
	}

	// log.Print("hit")
	return entry
}

func (c *lru) add(toAdd *dnsEntry) {
	// check if exists
	if entry, found := c.dnsMap[toAdd.domain]; found {
		// log.Println("found while trying to add")
		// remove pointers to entry
		prev := entry.prev
		next := entry.next

		prev.next = next
		next.prev = prev

		toAdd = entry
	} else {
		// fmt.Println("not found ", len(c.dnsMap))
		if len(c.dnsMap) == c.size {
			// log.Println("deleting due to cache size")
			// this operation will overflow our cache,
			// need to clear some room
			c.removeTail()
		}
	}

	c.dnsMap[toAdd.domain] = toAdd
	// add to beginning of list
	// point to old head
	toAdd.prev = nil
	toAdd.next = c.head
	c.head.prev = toAdd

	// toAdd becomes head
	c.head = toAdd
}

func (c *lru) remove(domain string) {
	// find in map
	toRemoveNode, ok := c.dnsMap[domain]
	if !ok {
		// not actually in map/list
		// log.Printf("cannot remove something that doesn't exist in cache: %v", domain)
		return
	}
	// set tail to be now the last undeleted node
	if c.tail.domain == domain {
		c.tail = c.tail.prev
	}

	// remove links in list
	prevNode := toRemoveNode.prev
	nextNode := toRemoveNode.next
	prevNode.next = nextNode
	nextNode.prev = prevNode

	// remove in map
	delete(c.dnsMap, domain)
}

func (c *lru) removeTail() {
	oldTail := c.tail
	newTail := c.tail.prev
	newTail.next = nil

	oldTail.prev = nil
	delete(c.dnsMap, oldTail.domain)

	c.tail = newTail

}

func (c *lru) String() string {
	str := fmt.Sprintf("lru max size: %v\nhead: %v\ntail: %v\n", c.size, c.head, c.tail)
	current := c.head
	for {
		str += fmt.Sprintf("\n%v", current)
		if current.next == nil {
			break
		}
		current = current.next
	}
	return str
}

func (d *dnsEntry) String() string {
	return fmt.Sprintf("%v : %v", d.domain, d.ip)
}

func main() {
	myCache := newLRU(5)

	domain := "my-domain"
	for i := 0; i < 6; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}

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
