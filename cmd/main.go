package main

import (
	"os"
	"strconv"

	"github.com/joshLong145/alien-invasion/cmd/world"
)

func main() {
	args := os.Args

	aliens, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err.Error())
	}

	file_path := args[1]
	w, err := world.NewWorld(file_path, aliens)

	if err != nil {
		panic(err.Error())
	}

	w.Run()

	println(w.String())
}
