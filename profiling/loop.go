package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"unsafe"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")

func loopThrough(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("i is %v and ", i)
		if i%2 == 0 {
			fmt.Println("even!")
		} else {
			fmt.Println("odd")
		}
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}() // () so the function will be executed

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	hello := "hello"
	fmt.Printf("the size of %v is %v", hello, unsafe.Sizeof(hello)) // 8

	loopThrough(1000)
}
