package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"

	"github.com/blck-snwmn/subslicebench/cmd/schema/fbs"
)

func read(filename string) *fbs.UserPositions {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fbs.GetSizePrefixedRootAsUserPositions(b, 0)
}

func open(filename string) []byte {
	b := read(filename)
	var p fbs.UserPosition
	b.Poss(&p, 0)
	nametmp := p.Name()
	name := make([]byte, len(nametmp))
	copy(name, nametmp)
	return name
}

func open2(filename string) []byte {
	b := read(filename)
	var p fbs.UserPosition
	b.Poss(&p, 0)
	return p.Name()
}

func gen() map[int][]byte {
	return map[int][]byte{}
}

func main() {
	use2 := flag.Bool("b", false, "")
	flag.Parse()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v KB\n", memStats.Alloc/1024)
	runtime.GC()
	mm := gen()
	for i := 0; i < 1000; i++ {
		var f func(string) []byte
		if *use2 {
			f = open2
		} else {
			f = open
		}
		b := f("./testdata/data")
		runtime.GC()
		mm[i%11] = b
	}
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v KB\n", memStats.Alloc/1024)
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v KB\n", memStats.Alloc/1024)

	fmt.Fprintln(io.Discard, mm)
}
