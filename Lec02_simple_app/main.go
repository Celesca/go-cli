package main

import (
	"flag"
	"fmt"
)

func main() {

	var name string
	flag.StringVar(&name, "name", "Firstname Lastname", "Yourname")
	flag.Parse()

	fmt.Printf("Hello %s", name)

}
