# Benchmarking and Profiling in GO

## Code profiling, flame graphs, memory usage

I leaned heavily of a few blog articles, mainly this one: https://blog.golang.org/profiling-go-programs

**Requirements and dependencies**
* install go and go-tools if you haven’t already
https://golang.org/doc/install

* to create graphical svgs with either pprof or go-torch, install graphviz:http://www.graphviz.org/
  `brew install graphviz`

* also to use go-torch, it requires brendanGregg’s flamegraph in the GOPATH (I dropped it right in the repo)
  * FlameGraph: https://github.com/brendangregg/FlameGraph
git@github.com:brendangregg/FlameGraph.git
  * go-torch: https://github.com/uber/go-torch
    git@github.com:uber/go-torch.git
go-torch binaryname cpu.prof


**Running**

There are multiple ways to run `pprof` on your code, either via `go-test`, as a command-line input straight in the code, or via the http/server listen.

### Via Command-line
`import 	"runtime/pprof"`
create flags to pass in if want to run pprof
note: if you run different types of profiling they could interact with each other to give false numbers

```var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")

if c.CPUPROFILE != "" {
		f, err := os.Create(c.CPUPROFILE)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("starting StartCPUProfile")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

if c.MEMPROFILE != "" {
		f, err := os.Create(c.MEMPROFILE)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		fmt.Println("starting WriteHeapProfile")
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}```

compile binary
`$ go build main.go`
run binary with command-line flags for profile that you want
`$ ./main --cpuprofile=cpu.prof`
`$ ./main --memprofile=mem.prof`

Now that you've created the profiling file, it's time to look at it in pprof.

### To Use pprof
`$ go tool pprof binaryname profilename` aka `$go tool pprof main cpu.prof`
Once inside pprof, can look at the topN items on the cpu (or memory usage) with
`top10`
sort by the 4th and 5th columns with `-cum` `top20 -cum`
You can create an svg file with calls in boxes and their callers with `web`
To view the source code for a particular call use `list` ie `list myawesomefunction`

### Via testing framework
or can run benchmarking with go test

```
import "testing"

func BenchmarkMyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
    myfunction(5)
  }
} ```


then run
`go test -run FFF -bench Bench -memprofile mem2.prof --memprofilerate 1`
FFF is to regex match no tests to run, so it just executes the benchmarking
memprofilerate will grab all memory instead of just a sampling


### Via http/pprof
* For use with currently running binaries, when perhaps there is a rampup in cpu or memory consumption that one wants to skip.

`import _ "net/http/pprof"`

```go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()```

compile and run as usual
then on browser go to
 http://localhost:6060/debug/pprof/heap or  http://localhost:6060/debug/pprof or  http://localhost:6060/debug/pprof/profile (can also use curl if no broswer access)

### Miscellaneous

* Can also use runtime.Memstats https://golang.org/pkg/runtime/#MemStats to get memory information

plop this inside function of interest
```
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Printf("mem.TotalAlloc: %v\n", mem.TotalAlloc)
	log.Printf("mem.HeapAlloc: %v\n", mem.HeapAlloc)
	log.Printf("mem.HeapSys: %v\n", mem.HeapSys)
```

* To get memory size of a variable, use `unsafe.Sizeof(variablename)` with a prints statement
* To get type of a variable, use %T in a print statement like so `fmt.Printf("%T", hello)  -> returns string`


Full list of pprof commands
```
Entering interactive mode (type "help" for commands)
(pprof) help

 Commands:
   cmd [n] [--cum] [focus_regex]* [-ignore_regex]*
       Produce a text report with the top n entries.
       Include samples matching focus_regex, and exclude ignore_regex.
       Add --cum to sort using cumulative data.
       Available commands:
         callgrind    Outputs a graph in callgrind format
         disasm       Output annotated assembly for functions matching regexp or address
         dot          Outputs a graph in DOT format
         eog          Visualize graph through eog
         evince       Visualize graph through evince
         gif          Outputs a graph image in GIF format
         gv           Visualize graph through gv
         list         Output annotated source for functions matching regexp
         pdf          Outputs a graph in PDF format
         peek         Output callers/callees of functions matching regexp
         png          Outputs a graph image in PNG format
         proto        Outputs the profile in compressed protobuf format
         ps           Outputs a graph in PS format
         raw          Outputs a text representation of the raw profile
         svg          Outputs a graph in SVG format
         tags         Outputs all tags in the profile
         text         Outputs top entries in text form
         top          Outputs top entries in text form
         tree         Outputs a text rendering of call graph
         web          Visualize graph through web browser
         weblist      Output annotated source in HTML for functions matching regexp or address
   peek func_regex
       Display callers and callees of functions matching func_regex.

   dot [n] [focus_regex]* [-ignore_regex]* [>file]
       Produce an annotated callgraph with the top n entries.
       Include samples matching focus_regex, and exclude ignore_regex.
       For other outputs, replace dot with:
       - Graphic formats: dot, svg, pdf, ps, gif, png (use > to name output file)
       - Graph viewer:    gv, web, evince, eog

   callgrind [n] [focus_regex]* [-ignore_regex]* [>file]
       Produce a file in callgrind-compatible format.
       Include samples matching focus_regex, and exclude ignore_regex.

   weblist func_regex [-ignore_regex]*
       Show annotated source with interspersed assembly in a web browser.

   list func_regex [-ignore_regex]*
       Print source for routines matching func_regex, and exclude ignore_regex.

   disasm func_regex [-ignore_regex]*
       Disassemble routines matching func_regex, and exclude ignore_regex.

   tags tag_regex [-ignore_regex]*
       List tags with key:value matching tag_regex and exclude ignore_regex.

   quit/exit/^D
 	     Exit pprof.

   option=value
       The following options can be set individually:
           cum/flat:           Sort entries based on cumulative or flat data
           call_tree:          Build context-sensitive call trees
           nodecount:          Max number of entries to display
           nodefraction:       Min frequency ratio of nodes to display
           edgefraction:       Min frequency ratio of edges to display
           focus/ignore:       Regexp to include/exclude samples by name/file
           tagfocus/tagignore: Regexp or value range to filter samples by tag
                               eg "1mb", "1mb:2mb", ":64kb"

           functions:          Level of aggregation for sample data
           files:
           lines:
           addresses:

           unit:               Measurement unit to use on reports

           Sample value selection by index:
            sample_index:      Index of sample value to display
            mean:              Average sample value over first value

           Sample value selection by name:
            alloc_space        for heap profiles
            alloc_objects
            inuse_space
            inuse_objects

            total_delay        for contention profiles
            mean_delay
            contentions

   :   Clear focus/ignore/hide/tagfocus/tagignore```
