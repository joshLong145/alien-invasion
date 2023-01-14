package world

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Primary file parser
//
// will currently abort in operation
// if a line being parsed is incorrectly formatted
// an error is return and the world wont be created
// this can be improved to ignore incorrect file inputs
// with a pre processor
func parseWorld(path string, w *World) error {
	f, err := os.Open(path)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	graph := make(map[string][]string)

	// pre process the file into a adjacency matrix
	for scanner.Scan() {
		// trim whitespace character left after parsing
		city := strings.Trim(scanner.Text(), "\u200b")
		parts := strings.Split(city, " ")
		if len(parts) == 1 {
			return fmt.Errorf("input file not properly formatted")
		}
		name := parts[0]

		graph[name] = parts[1:]
		w.Locations[name] = &Location{
			Name: name,
		}
	}

	// process the graph to connect locations
	for k, v := range graph {
		loc, ok := w.Locations[k]
		// if the location isnt found we skip it for connecting, making it an island.
		if !ok {
			fmt.Println("could not find city")
			w.Locations[k] = &Location{
				Name:     k,
				IsIsland: true,
			}
			continue
		}

		for _, direction := range v {
			values := strings.Split(direction, "=")
			// if the input was found to not be formatted properly abort.
			if len(values) != 2 {
				return fmt.Errorf("input file not properly formatted")
			}
			connect, ok := w.Locations[values[1]]

			if !ok {
				connect = &Location{
					Name: values[1],
				}
				w.Locations[values[1]] = connect
			}

			//normalize the location assuming also only numeric characters
			values[0] = strings.ToLower(values[0])
			switch values[0] {
			case "north":
				loc.North = connect
				connect.South = loc
			case "south":
				loc.South = connect
				connect.North = loc
			case "east":
				loc.East = connect
				connect.West = loc
			case "west":
				loc.West = connect
				connect.East = loc
			}
		}
	}

	return nil
}
