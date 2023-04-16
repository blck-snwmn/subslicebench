package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
)

func read(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func open(filename string) []byte {
	var buf = make([]byte, 2)
	b := read(filename)
	copy(buf, b)
	return buf
}
func open2(filename string) []byte {
	b := read(filename)
	return b[0:2]
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
	for i := 0; i < 100; i++ {
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
