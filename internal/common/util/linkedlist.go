package util

import (
	"container/list"
	"sync"
)

type Equalable interface {
	EqualTo(interface{}) bool
}

type LinkedList struct {
	list *list.List
	mux  sync.RWMutex
}

func (c *LinkedList) Len() int {
	return c.list.Len()
}

func (c *LinkedList) PushBack(element interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.list.PushBack(element)
}

func (c *LinkedList) PushBacks(elements []interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, element := range elements {
		c.list.PushBack(element)
	}
}

func (c *LinkedList) Remove(equalableObj Equalable) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for e := c.list.Front(); e != nil; e = e.Next() {
		if equalableObj.EqualTo(e.Value) {
			c.list.Remove(e)
		}
	}
}

func (c *LinkedList) Removes(equalableObjs []Equalable) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, equalableObj := range equalableObjs {
		c.Remove(equalableObj)
	}
}

func (c *LinkedList) CheckElements(runnable func(interface{}) error) error {
	if c.list.Len() == 0 {
		return nil
	}

	for e := c.list.Front(); e != nil; e = e.Next() {
		if e == nil {
			continue
		}

		if err := runnable(e.Value); err != nil {
			return err
		}
	}

	return nil
}

func (c *LinkedList) FindElement(equalableObj Equalable) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	for e := c.list.Front(); e != nil; e = e.Next() {
		if e == nil {
			continue
		}

		if equalableObj.EqualTo(e.Value) {
			return e.Value
		}

	}

	return nil
}

func CreateLinkedList() *LinkedList {
	linkedList := new(LinkedList)
	linkedList.list = list.New()
	return linkedList
}
