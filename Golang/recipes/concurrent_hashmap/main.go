package main

import (
	"fmt"
	"log"
)

/*
	The safe map is implemented by storing the map in a goroutine
	The only way to access the map is via channels
*/

type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type UpdateFunc func(interface{}, bool) interface{}

type safeMap chan commandData

type commandData struct {
	action  commandAction
	key     string
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[string]interface{}
	updater UpdateFunc
}

type commandAction int

const (
	removeItem commandAction = iota
	end
	find
	insert
	length
	update
)

func (sm safeMap) Insert(key string, value interface{}) {
	sm <- commandData{action: insert, key: key, value: value}
}

func (sm safeMap) Delete(key string) {
	sm <- commandData{action: removeItem, key: key}
}

type findResult struct {
	value interface{}
	found bool
}

func(sm safeMap) Find(key string) (value interface{}, found bool) {
	reply := make(chan interface{})
	sm <- commandData{action: find, key :key, result:reply}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (sm safeMap) Len() int {
    reply := make(chan interface{})
    sm <- commandData{action: length, result: reply}
    return (<-reply).(int)
}

func (sm safeMap) Update(key string, updater UpdateFunc) {
    sm <- commandData{action: update, key: key, updater: updater}
}

func (sm safeMap) Close() map[string]interface{} {
    reply := make(chan map[string]interface{})
    sm <- commandData{action: end, data: reply}
    return <-reply
}


func (sm safeMap) run() {
	store := make(map[string]interface{})
	for command := range sm {
        switch command.action {
        case insert:
            store[command.key] = command.value
        case removeItem:
            delete(store, command.key)
        case find:
            value, found := store[command.key]
            command.result <- findResult{value, found}
        case length:
            command.result <- len(store)
        case update:
            value, found := store[command.key]
            store[command.key] = command.updater(value, found)
        case end:
            close(sm)
            command.data <- store
        }
	}
}
func New() safeMap {
	sm := make(safeMap)
	go sm.run()
	return sm
}
func main() {
	log.Println("here")
	sm := New()
	sm.Insert("greeting1", "hiya hun")
	sm.Insert("greeting2", 2)
	if greeting, ok := sm.Find("greeting1"); ok {
		fmt.Println(greeting)
	}
}
