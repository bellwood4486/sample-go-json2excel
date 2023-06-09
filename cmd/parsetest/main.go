package main

import (
	"flag"
	"fmt"
	j2e "github.com/bellwood4486/sample-go-json2excel"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var parseCase = flag.Int("case", 1, "parse case")

func main() {
	flag.Parse()

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

	// 重いJSONからExcelを作る
	fmt.Printf("parse case: %d\n", *parseCase)
	switch *parseCase {
	case 1:
		err = j2e.ToExcelCase1(f)
	case 2:
		err = j2e.ToExcelCase2(f)
	case 3:
		err = j2e.ToExcelCase3(f)
	case 4:
		// JSONのパースだけを行う
		_, err = j2e.ParseJSONBulk(f)
	case 5:
		// JSONのパースだけを行う
		_, err = j2e.ParseJSONStream(f)
	default:
		log.Fatalf("unknown case: %d", *parseCase)
	}
	if err != nil {
		log.Fatal(err)
	}

	// メモリプロファイルを取得する
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		//runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
