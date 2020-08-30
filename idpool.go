package tools

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type IdPool struct {
	mutex  sync.Mutex
	exists map[int]bool
	li     list.List
}

func (p *IdPool) Init(size int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.exists = make(map[int]bool)
	for i := 0; i < size; i++ {
		p.li.PushBack(i)
		p.exists[i] = true
	}
}
func (p *IdPool) Get() (id int, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ele := p.li.Front()
	if ele == nil {
		err = errors.New("No id can be use\n")
		return
	}
	id = ele.Value.(int)
	p.exists[id] = false
	p.li.Remove(ele)
	return
}
func (p *IdPool) Put(id int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	exists, ok := p.exists[id]
	if !ok {
		panic(fmt.Sprintf("ID(%d) out range value", id))
	}
	if exists {
		panic(fmt.Sprintf("ID(%d) is exists already", id))
	}
	p.exists[id] = true
	p.li.PushFront(id)
}
