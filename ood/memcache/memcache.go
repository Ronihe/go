package memcache

import (
	"time"
)

const intMax int = 0x7fffffff

type cacheNode struct {
	value      int
	willExpire bool
	expireTime time.Time
}

type client map[int]cacheNode

type MemcacheInterface interface {
	Get(key int) int
	Set(key, value int, ttl ...int) int
	Delete(key int) error
	Decrease(key, delta int) int
	Increase(key, delta int) int
}

func (c *client) Get(key int) int {
	if _, ok := (*c)[key]; !ok {
		return intMax
	}

	cacheNode := (*c)[key]
	if !cacheNode.willExpire || cacheNode.expireTime.After(time.Now()) {
		return cacheNode.value
	}

	return intMax
}

func (c *client) Set(key, val int, ttl ...int) int {

	cacheNode := cacheNode{
		value: val,
	}
	if len(ttl) > 0 {
		cacheNode.willExpire = true
		cacheNode.expireTime = time.Now().Add(time.Second * time.Duration(ttl[0]))
	} else {
		cacheNode.willExpire = false
	}
	(*c)[key] = cacheNode
	return cacheNode.value
}

func (c *client) Delete(key int) error {
	delete(*c, key)
	return nil
}

func (c *client) Increase(key, delta int) int {
	val := c.Get(key)
	if val == intMax {
		return intMax
	}

	node := (*c)[key]
	node.value = val + delta
	(*c)[key] = node
	return node.value
}

func (c *client) Decrease(key, delta int) int {
	val := c.Get(key)
	if val == intMax {
		return intMax
	}

	node := (*c)[key]
	node.value = val - delta
	(*c)[key] = node
	return node.value
}

func NewMemcache() MemcacheInterface{
	return &client{}
}
