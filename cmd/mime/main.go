package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	conf "github.com/therealfakemoot/slayer/conf"
)

func main() {
	var config = flag.String("config", "conf.toml", "config file")

	flag.Parse()

	f, err := os.Open(*config)
	if err != nil {
		log.Fatal("could not open config file")
	}
	auth := conf.LoadAuth(f)

	fmt.Printf("%#v\n", auth)
}
