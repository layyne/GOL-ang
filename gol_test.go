package main

import (
	"reflect"
	"testing"
)

func TestBoard(t *testing.T) {
	t.Run("board.toggle sets and unsets the board cell", func(t *testing.T) {
		var board *Board = newBoard(5, 5)
		board.toggle(1, 2)

		got := board.at(1, 2)
		want := true

		if got != want {
			t.Errorf("Error: got %t want %t", got, want)
		}

		board.toggle(1, 2)

		got = board.at(1, 2)
		want = false

		if got != want {
			t.Errorf("Error: got %t want %t", got, want)
		}
	})

	t.Run("board updates correctly", func(t *testing.T) {
		var got *Board = newBoard(3, 3)

		got.toggle(0, 1)
		got.toggle(1, 1)
		got.toggle(2, 1)

		got.update()
		got.updateCount()

		var want *Board = newBoard(3, 3)
		want.toggle(1, 0)
		want.toggle(1, 1)
		want.toggle(1, 2)
		want.updateCount()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error: got %+v want %+v", got, want)
		}
	})
}
