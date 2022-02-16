package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	helloWorld := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloWorld)
}
