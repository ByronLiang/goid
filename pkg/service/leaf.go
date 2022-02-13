package service

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/ByronLiang/goid/pkg/db"

	"github.com/ByronLiang/goid/pkg/model"
)

var Leaf = NewLeaf()

type leaf struct {
	mu   sync.RWMutex
	data map[int64]*leafNode
}

type leafNode struct {
	domainId int64
	Current  int64
	Max      int64
	Min      int64
	percent  int64
	buffer   chan int64
	stop     chan struct{}
}

func NewLeaf() *leaf {
	return &leaf{
		mu:   sync.RWMutex{},
		data: make(map[int64]*leafNode),
	}
}

func (l *leaf) FakeLeafNode(domain int, size int64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i := 1; i <= domain; i++ {
		rand.Seed(time.Now().Unix())
		start := rand.Int63()
		node := &leafNode{
			domainId: int64(i),
			Current:  start,
			buffer:   make(chan int64, size),
			stop:     make(chan struct{}),
		}
		l.data[node.domainId] = node
		node.FakeWatch()
	}
}

func (l *leaf) AddLeafNode(nodes []*model.Leaf, size int64, percent int64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, node := range nodes {
		max := node.MaxId + node.Step - 1
		min := max - percent*node.Step/100
		node := &leafNode{
			domainId: node.DomainId,
			Current:  node.MaxId,
			Max:      max,
			Min:      min,
			percent:  percent,
			buffer:   make(chan int64, size),
			stop:     make(chan struct{}),
		}
		l.data[node.domainId] = node
		node.Watch()
	}
}

func (l *leaf) Get(domainId int64) (int64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if node, ok := l.data[domainId]; ok {
		num, ok := node.get()
		if ok {
			return num, nil
		}
		return 0, errors.New("leaf buffer had closed")
	}
	return 0, errors.New("domainId no exist")
}

func (ln *leafNode) FakeWatch() {
	go func() {
		for {
			select {
			case ln.buffer <- ln.Current:
				ln.Current++
			case <-ln.stop:
				// 回收未使用号段
				close(ln.buffer)
				return
			default:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
}

func (ln *leafNode) Watch() {
	go func() {
		for {
			select {
			case ln.buffer <- ln.Current:
				ln.Current++
				if ln.Current >= ln.Min {
					// 申请下一批序列号
					ln.resize()
				}
			case <-ln.stop:
				// 回收未使用号段
				close(ln.buffer)
				return
			default:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
}

func (ln *leafNode) get() (int64, bool) {
	num, ok := <-ln.buffer
	return num, ok
}

func (ln *leafNode) resize() {
	l, err := db.LeafDao.GetByDomainId(ln.domainId)
	if err != nil {
		return
	}
	rowRes := db.LeafDao.UpdateMaxId(l.MaxId, ln.domainId, l.MaxId+l.Step)
	if rowRes == 0 {
		return
	}
	ln.Max = l.MaxId + l.Step - 1
	ln.Min = ln.Max - ln.percent*l.Step/100
}
