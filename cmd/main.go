package main

import (
	"fmt"
	"os"

	"github.com/ashutoshraina/myootadapter/mygrpcadapter"
)

func main() {
	println("Hello World")
	addr := ""
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	s, err := mygrpcadapter.NewMyGrpcAdapter(addr)
	if err != nil {
		fmt.Printf("unable to start server: %v", err)
		os.Exit(-1)
	}

	shutdown := make(chan error, 1)
	go func() {
		s.Run(shutdown)
	}()
	_ = <-shutdown
}
