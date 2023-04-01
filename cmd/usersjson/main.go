package main

import (
	"fmt"
	j2e "github.com/bellwood4486/sample-go-json2excel"
	"log"
	"os"
)

func main() {
	amount := 1000000
	l := j2e.UserList{
		Users: make([]*j2e.User, 0, amount),
	}
	for i := 1; i <= amount; i++ {
		l.Users = append(l.Users, &j2e.User{
			Name:    fmt.Sprintf("user%07d", i),
			Age:     20,
			Profile: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam",
		})
	}

	err := l.ToJSON(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
