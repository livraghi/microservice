package main

import "fmt"

func main() {
	cfg, err := LoadConfiguration()
	if err != nil {
		panic(err)
	}
	fmt.Printf("configuration loaded: %v\n", cfg)
}
