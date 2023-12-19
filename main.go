package main

import (
	"fmt"
	"goweb/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
	}

}
