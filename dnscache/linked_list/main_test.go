package main

/*
$go test -run . -bench . -benchmem
goos: darwin
goarch: amd64
BenchmarkFindIfExist50Yes-12       	100000000	        13.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist50No-12        	100000000	        16.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist500Yes-12      	100000000	        12.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist500No-12       	100000000	        12.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist5000Yes-12     	100000000	        12.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist5000No-12      	100000000	        17.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist50000Yes-12    	100000000	        10.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIfExist50000No-12     	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
PASS
*/

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFindIfExist50Yes(b *testing.B) {
	size := 50
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("my-domain40")
	}
}

func BenchmarkFindIfExist50No(b *testing.B) {
	size := 50
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("does-not-exist")
	}
}

func BenchmarkFindIfExist500Yes(b *testing.B) {
	size := 500
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("my-domain400")
	}
}

func BenchmarkFindIfExist500No(b *testing.B) {
	size := 500
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("does-not-exist")
	}
}

func BenchmarkFindIfExist5000Yes(b *testing.B) {
	size := 5000
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("my-domain400")
	}
}

func BenchmarkFindIfExist5000No(b *testing.B) {
	size := 5000
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("does-not-exist")
	}
}

func BenchmarkFindIfExist50000Yes(b *testing.B) {
	size := 50000
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("my-domain400")
	}
}

func BenchmarkFindIfExist50000No(b *testing.B) {
	size := 50000
	myCache := newLRU(size)
	domain := "my-domain"
	for i := 0; i < size; i++ {
		myCache.add(
			newDNSEntry(
				fmt.Sprint(domain, strconv.Itoa(i)),
				fmt.Sprint("85.234.34", strconv.Itoa(i)),
			),
		)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myCache.findIfExist("does-not-exist")
	}
}
