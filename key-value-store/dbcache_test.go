package main

import (
	"testing"
)

func TestGetSetCacheNoCommit(t *testing.T) {
	var expected = map[string]int{
		"a": 0,
		"b": 0,
	}

	t.Run("TestSet", func(t *testing.T) {
		db := NewDatabaseCache()
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

func TestGetSetCacheCommit(t *testing.T) {
	var expected = map[string]int{
		"a": 5,
		"b": 7,
	}

	t.Run("TestSet", func(t *testing.T) {
		db := NewDatabaseCache()
		db.Begin()
		db.Set("a", 5)
		db.Set("b", 7)
		db.Commit()

		for k, v := range expected {
			actualV, _ := db.Get(k)
			if actualV != v {
				t.Errorf("got %d for key %s, expected %d", actualV, k, v)
			}
		}
	})

}
func TestRollbackCache(t *testing.T) {
	var expected = map[string]int{
		"a": 6,
		"b": 8,
	}

	t.Run("TestRollback", func(t *testing.T) {
		db := NewDatabaseCache()

		db.Begin()
		db.Set("a", 6)
		db.Set("b", 8)
		db.Commit()

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
