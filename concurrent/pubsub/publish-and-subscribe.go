package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

type Publisher struct {
	m 			sync.RWMutex
	buffer  	int
	timeout 	time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(buffer int, timeout time.Duration) *Publisher {
	return &Publisher{
		buffer: buffer,
		timeout: timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
	ch := make(subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Remove(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p * Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}