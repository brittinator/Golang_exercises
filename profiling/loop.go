package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"unsafe"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 10000; i++ {
		fmt.Printf("i is %v and ", i)
		if i%2 == 0 {
			fmt.Println("even!")
		} else {
			fmt.Println("odd")
		}
	}
	hello := "hello"
	unsafe.Sizeof(hello)

}
