package world

import (
	"fmt"
	"math/rand"
)

func placeAliens(w *World, amount int) error {

	if amount < 1 {
		return fmt.Errorf("cannot place a negative amount")
	}

	cities := make([]string, 0)
	for k, v := range w.Locations {
		// if a location is marked as an island, then we skip placing aliens
		if !v.IsIsland {
			cities = append(cities, k)
		}
	}

	for i := 0; i < amount; i++ {
		loc := rand.Intn(len(cities) - 1)
		w.Locations[cities[loc]].Occupants = append(w.Locations[cities[loc]].Occupants, &Alien{
			Id:        i,
			MoveCount: 0, // start the count at 0 as being born is not moving
		})
	}

	return nil
}
