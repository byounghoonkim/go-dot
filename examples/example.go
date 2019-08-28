package main

import (
	"fmt"

	"github.com/byounghoonkim/go-dot"
)
	

func main() {

	type Config struct {
		Server string
		Username string
	}

	config := Config{
		"server1",
		"user1",
	}

	d := dot.New()
	err := d.Save(&config)

	if err != nil {
		panic(err)
	}

	config2 := Config{}
	err = d.Load(&config2)
	if err != nil {
		panic(err)
	}

	fmt.Println(config2)

}