package main

import (
	"fmt"

	simplemultialgo "github.com/moncho/go-simplemultialgo"
)

func main() {
	algo, err := simplemultialgo.NiceHashMultiAlgo(map[string]int{"scrypt": 1, "x11": 7, "quark": 12})
	if err != nil {
		panic("Run for your life!!")
	}

	fmt.Printf("Most profitable algo is %s, currently paying %v\n", algo.Name, algo.Paying)
}
