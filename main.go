package main

import (
	"fmt"
	"github.com/zerosimonn/blog-aggregator/internal/config"
	"log"
)


func main() {
	cfg,err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("zerosimonn")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", cfg)
}
