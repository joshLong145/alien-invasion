package world_test

import (
	"testing"

	"github.com/joshLong145/alien-invasion/cmd/world"
	"github.com/stretchr/testify/assert"
)

func TestNewWorldCreateValid(t *testing.T) {
	const (
		FILE_PATH_VALID   = "/Users/bean/Documents/code/alien-invasion/cmd/world/test_cases/world_valid.txt"
		FILE_PATH_INVALID = "/Users/bean/Documents/code/alien-invasion/cmd/world/test_cases/world_invalid.txt"
	)

	t.Run("Should create world with 5 locations", func(t *testing.T) {
		world, err := world.NewWorld(FILE_PATH_VALID, 100)

		assert.Nil(t, err)
		assert.Equal(t, len(world.Locations), 5)
	})

	t.Run("Should create world with 5 locations", func(t *testing.T) {
		world, err := world.NewWorld(FILE_PATH_VALID, 100)

		assert.Nil(t, err)
		total := 0

		for _, l := range world.Locations {
			total += len(l.Occupants)
		}

		assert.Equal(t, total, 100)
	})

	t.Run("Should return error creating world negative aliens", func(t *testing.T) {
		_, err := world.NewWorld(FILE_PATH_VALID, -1)

		assert.Error(t, err)
	})

	t.Run("Should return error parsing invalid file formatting", func(t *testing.T) {
		_, err := world.NewWorld(FILE_PATH_INVALID, 100)
		assert.Error(t, err)
	})
}
