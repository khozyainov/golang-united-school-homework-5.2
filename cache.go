package cache

import "time"

type Cache struct {
	data map[string]record
}

type record struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{data: map[string]record{}}
}

func (c *Cache) Get(key string) (string, bool) {
	r, ok := c.data[key]
	if !r.deadline.IsZero() && r.deadline.Before(time.Now()) {
		delete(c.data, key)
		return "", false
	}
	return r.value, ok
}

func (c *Cache) Put(key, value string) {
	c.data[key] = record{value: value}
}

func (c *Cache) Keys() []string {
	var keys []string
	now := time.Now()
	for k, r := range c.data {
		if !r.deadline.IsZero() && r.deadline.Before(now) {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = record{value: value, deadline: deadline}
}
