package config

import (
	"errors"
	"log"
)

type cache struct {
	Db map[string]interface{}
}

var Session *cache

func init() {
	Session = newCache()
	log.Println("Cache ready!")
}

func newCache() *cache {
	return &cache{}
}

func (c *cache) Set(key string, value interface{}) {
	c.Db = make(map[string]interface{})
	c.Db[key] = value
}

func (c *cache) Get(key string) (interface{}, error) {
	v, ok := c.Db[key]
	if !ok {
		return nil, errors.New("Can't find the session key.")
	}
	return v, nil
}

func (c *cache) Del(key string) {
	delete(c.Db, key)
}
