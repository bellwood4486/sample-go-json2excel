package main

import (
	"flag"
	"fmt"
	j2e "github.com/bellwood4486/sample-go-json2excel"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	f, err := os.Open("./data/userlist.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // サンプルコードなのでエラーハンドリングは省略

	// パースの時間を計測する
	start := time.Now()
	defer func(s time.Time) {
		fmt.Printf("elapsed: %s\n", time.Since(s))
	}(start)

	flag.Parse()
	// CPUプロファイルを取得する
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// 重いJSONのパース処理
	l := &j2e.UserList{}
	err = l.ParseJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("uers: %d\n", len(l.Users))

	// メモリプロファイルを取得する
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
