package main

import (
	"fmt"

	"github.com/byounghoonkim/go-dot"
)

func main() {

	type Config struct {
		Server   string
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

	configFolder, err := d.GetConfigFolder()
	if err != nil {
		panic(err)
	}
	fmt.Println("config folder : ", configFolder)

	configPath, err := d.GetConfigPath(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println("config path   : ", configPath)

	config2 := Config{}
	err = d.Load(&config2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", config2)

}
