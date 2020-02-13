// https://github.com/dgryski/trifles/blob/master/cachetest/random/random.go
package cache

import (
	"github.com/devforfu/blogapp/app/blog"
	"math/rand"
)

type PostsCache struct {
	capacity int
	data 	 map[string]*blog.Post
	keys 	 []string
}

func New(capacity int) *PostsCache {
	return &PostsCache{
		capacity: capacity,
		data: 	  make(map[string]*blog.Post),
		keys:	  make([]string, capacity),
	}
}

func (c *PostsCache) Get(key string) *blog.Post {
	return c.data[key]
}

func (c *PostsCache) Set(key string, value *blog.Post) {
	slot := len(c.data)
	if len(c.data) == c.capacity {
		slot = rand.Intn(c.capacity)
		delete(c.data, c.keys[slot])
	}
	c.keys[slot] = key
	c.data[key] = value
}

var Default = New(10)
