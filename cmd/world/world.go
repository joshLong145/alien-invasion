package world

import (
	"fmt"
	"math/rand"
	"strings"
)

//👾🛸
type Alien struct {
	Id        int
	MoveCount int
}

// Location repersents a city allowing for other locations to be associated on cardinal directions
type Location struct {
	Name      string
	North     *Location
	South     *Location
	East      *Location
	West      *Location
	Occupants []*Alien
	IsIsland  bool
}

// World is a collection of Locations related by its name to the struct instance
type World struct {
	Locations map[string]*Location
}

func NewWorld(path string, aliens int) (*World, error) {
	world := &World{
		Locations: make(map[string]*Location),
	}
	err := parseWorld(path, world)

	if err != nil {
		return nil, err
	}

	err = placeAliens(world, aliens)

	if err != nil {
		return nil, err
	}

	return world, nil
}

func (w *World) Run() error {
	// start the game loop
	// ordering goes
	// 1) move aliens to new locations
	// 2) check for a fight at each
	for w.shouldRun() {
		w.moveAliens()
		w.checkForFights()
	}

	return nil
}

// Display locations and alien counts
func (w *World) String() string {
	builder := strings.Builder{}
	for k, v := range w.Locations {
		if v.North != nil {
			builder.Write([]byte(fmt.Sprintf("\t%s\n", v.North.Name)))
		} else {
			builder.Write([]byte(fmt.Sprintf("\t%s\n", "none")))
		}
		builder.Write([]byte("\t|\n"))

		if v.West != nil && v.East != nil {
			builder.Write([]byte(fmt.Sprintf("%s - %s - %s\n", v.West.Name, k, v.East.Name)))
		} else if v.West == nil && v.East == nil {
			builder.Write([]byte(fmt.Sprintf("%s - %s - %s\n", "none", k, "none")))
		} else if v.West == nil {
			builder.Write([]byte(fmt.Sprintf("%s - %s - %s\n", "none", k, v.East.Name)))
		} else if v.East == nil {
			builder.Write([]byte(fmt.Sprintf("%s - %s - %s\n", v.West.Name, k, "none")))
		}
		builder.Write([]byte("\t|\n"))

		if v.South != nil {
			builder.Write([]byte(fmt.Sprintf("\t%s\n", v.South.Name)))
		} else {
			builder.Write([]byte(fmt.Sprintf("\t%s\n", "none")))
		}

		builder.Write([]byte("\n"))
		builder.Write([]byte(fmt.Sprintf("%s aliens: %d\n", v.Name, len(v.Occupants))))
	}

	return builder.String()
}

func (w *World) shouldRun() bool {
	keepRunning := false // assume the world should stop until proven false
	for _, v := range w.Locations {
		for _, a := range v.Occupants {
			if a.MoveCount <= 10000 && !v.IsIsland {
				keepRunning = true
				break // dont waste time checking the rest
			}
		}
		if keepRunning {
			break // dont keep iterating if we already know we are running still
		}
	}

	return keepRunning

}

func (w *World) checkForFights() {
	for k, v := range w.Locations {
		if len(v.Occupants) > 1 {
			fmt.Printf("a fight has broken out at: %s\n", v.Name)
			// check for neighbors and dereference
			if w.Locations[k].East != nil {
				w.Locations[k].East.West = nil
			}
			if w.Locations[k].West != nil {
				w.Locations[k].West.East = nil
			}
			if w.Locations[k].North != nil {
				w.Locations[k].North.South = nil
			}
			if w.Locations[k].South != nil {
				w.Locations[k].South.North = nil
			}
			fmt.Printf("%s has been destroyed by Alien %d and Alien %d\n", k, w.Locations[k].Occupants[0].Id, w.Locations[k].Occupants[1].Id)
			delete(w.Locations, k)
		}
	}
}

func (w *World) moveAliens() {
	for _, v := range w.Locations {
		for i, a := range v.Occupants {
			//store the number of valid directions
			valid_moves := 0
			// relate a valid direction to a number for the rand move calculation
			dir_map := make(map[int]string)
			if v.North != nil {
				valid_moves++
				dir_map[valid_moves] = "north"
			}
			if v.South != nil {
				valid_moves++
				dir_map[valid_moves] = "south"
			}
			if v.East != nil {
				valid_moves++
				dir_map[valid_moves] = "east"
			}
			if v.West != nil {
				valid_moves++
				dir_map[valid_moves] = "west"
			}

			// if there are no valid moves, its an island
			if valid_moves < 1 {
				// While not the immediate place to check if there is an island,
				// it's more convient to do so here than to search around neighbors of the city
				// blown up in a fight
				v.IsIsland = true
				continue
			}

			dir := rand.Intn(valid_moves)
			a.MoveCount += 1
			switch dir_map[dir] {
			case "north":
				v.North.Occupants = append(v.North.Occupants, a)
			case "south":
				v.South.Occupants = append(v.South.Occupants, a)
			case "east":
				v.East.Occupants = append(v.East.Occupants, a)
			case "west":
				v.West.Occupants = append(v.West.Occupants, a)
			}
			if len(v.Occupants) > 2 {
				v.Occupants = append(v.Occupants[:i], v.Occupants[i+1:]...)
			} else {
				v.Occupants = make([]*Alien, 0)
			}
		}
	}
}
