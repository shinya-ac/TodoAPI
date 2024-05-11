package main

import (
	"fmt"

	config "github.com/shinya-ac/TodoAPI/configs"
)

func main() {
	fmt.Println(config.Config.Host)
}
