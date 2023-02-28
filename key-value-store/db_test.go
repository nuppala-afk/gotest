package main

import (
	"testing"
)

func TestGetSet(t *testing.T) {
	var expected = map[string]int{
		"a": 0,
		"b": 7,
	}

	t.Run("TestSet", func(t *testing.T) {
		db := NewDatabase()
		db.Begin()
		db.Set("a", 5)
		db.Set("b", 7)
		db.Unset("a")
		for k, v := range expected {
			actualV, _ := db.Get(k)
			if actualV != v {
				t.Errorf("got %d for key %s, expected %d", actualV, k, v)
			}
		}
	})

}
func TestRollback(t *testing.T) {
	var expected = map[string]int{
		"a": 0,
		"b": 8,
	}

	t.Run("TestRollback", func(t *testing.T) {
		db := NewDatabase()

		db.Begin()
		db.Set("a", 6)
		db.Set("b", 8)

		db.Begin()
		db.Set("a", 5)
		db.Set("b", 8)
		db.Unset("a")

		db.Begin()
		db.Set("a", 5)
		db.Set("b", 7)
		db.Unset("a")
		db.Rollback()

		for k, v := range expected {
			actualV, _ := db.Get(k)
			if actualV != v {
				t.Errorf("got %d for key %s, expected %d", actualV, k, v)
			}
		}
	})
}

func TestCommitRollback(t *testing.T) {
	var expected = map[string]int{
		"a": 0,
		"b": 8,
	}

	t.Run("TestRollback", func(t *testing.T) {
		db := NewDatabase()

		db.Begin()
		db.Set("a", 6)
		db.Set("b", 8)

		db.Begin()
		db.Set("a", 5)
		db.Set("b", 8)
		db.Unset("a")
		db.Commit() //still not handled correctly, tmpCache version seems to work well

		db.Begin()
		db.Set("a", 5)
		db.Set("b", 7)
		db.Rollback()

		for k, v := range expected {
			actualV, _ := db.Get(k)
			if actualV != v {
				t.Errorf("got %d for key %s, expected %d", actualV, k, v)
			}
		}
	})
}
