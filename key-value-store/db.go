package main

import (
	"fmt"
	"log"
)

type command struct {
	key    string
	value  int
	action string
}

type db struct {
	kvs          map[string]int
	tmpCache     map[string]int //not used yet
	commandCache []command
}

func NewDatabase() *db {
	return &db{
		kvs:          make(map[string]int),
		tmpCache:     make(map[string]int),
		commandCache: []command{},
	}
}

func (d *db) Set(key string, value int) {
	d.kvs[key] = value
	d.commandCache = append(d.commandCache, command{
		key:    key,
		value:  value,
		action: "set",
	})
}

func (d *db) Get(key string) (value int, ok bool) {
	if v, ok := d.kvs[key]; ok {
		return v, ok
	}
	return 0, false
}

func (d *db) Unset(key string) {
	delete(d.kvs, key)
	d.commandCache = append(d.commandCache, command{
		key:    key,
		action: "unset",
	})
}

func (d *db) Begin() {
	d.commandCache = append(d.commandCache, command{
		action: "begin",
	})
}

// If I remember the requirements correctly, ROLLBACK should undo any changes made since the last BEGIN command
func (d *db) Rollback() {
	for i := len(d.commandCache) - 1; i >= 0; i-- {
		// Iterate through command cache and delete all commandCache until you reach BEGIN
		if d.commandCache[i].action != "begin" {
			d.deleteLastCommand()
			continue
		}
		// You've reached BEGIN and that's the only command left in the cache
		// There is only trasaction and we restore the original state
		if i == 0 {
			log.Printf("No state changes made since last commit")
			return
		}

		//iterate backwards from BEGIN until you reach another BEGIN and execute all commandCache between those 2 statements
		var j int
	INNER:
		for j = i - 1; j >= 0; j-- {
			if d.commandCache[j].action == "begin" {
				break INNER
			}
		}

		if j == 0 {
			log.Printf("No state changes made since last commit")
			return
		}

		for k := j + 1; k <= i-1; k++ {
			fmt.Println(i, j, k, d.commandCache[k])
			d.executeCommand(d.commandCache[k])
		}
	}
}

func (d *db) executeCommand(c command) {
	if c.action == "set" {
		d.Set(c.key, c.value)
	} else if c.action == "unset" {
		d.Unset(c.key)
	} else {
		log.Fatal("Unknown command")
	}
}

func (d *db) Commit() {
	//NO OP
}

func (d *db) End() {
	// Delete all commands from commandCache
	d.commandCache = d.commandCache[:0]
}

func (d *db) deleteLastCommand() {
	l := len(d.commandCache)
	if l <= 0 {
		return
	}
	d.commandCache = d.commandCache[:l-1]
}
