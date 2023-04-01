package main

import (
	"fmt"
	j2e "github.com/bellwood4486/sample-go-json2excel"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./data/userlist.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	start := time.Now()
	defer func(s time.Time) {
		fmt.Printf("elapsed: %s\n", time.Since(s))
	}(start)

	l := &j2e.UserList{}
	err = l.ParseJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("uers: %d\n", len(l.Users))
}
