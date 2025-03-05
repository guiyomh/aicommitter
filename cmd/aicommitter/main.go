package main

import "github.com/guiyomh/aicommitter/internal/cli"

func main() {
	cli := cli.New()
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
