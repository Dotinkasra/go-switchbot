package main

import (
	"fmt"

	"github.com/Dotinkasra/go-switchbot/module"
)

func main() {
	d := new(module.Devices)
	devices, err := d.GetDevices()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(devices)

	s := new(module.Scenes)
	seces, err := s.GetScenes()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(seces)
}
