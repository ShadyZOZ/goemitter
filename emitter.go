package goemitter

import "sync"

// EventEmitter EventEmitterStruct
type EventEmitter struct {
	mutex      sync.Mutex
	handlerMap map[string][]func(...interface{})
	channelMap map[string]chan []interface{}
}

// On add event handler
func (em *EventEmitter) On(eventName string, eventHandler func(...interface{})) {
	em.mutex.Lock()
	defer em.mutex.Unlock()
	if em.handlerMap == nil {
		em.handlerMap = make(map[string][]func(...interface{}))
	}
	em.handlerMap[eventName] = append(em.handlerMap[eventName], eventHandler)
	if em.channelMap == nil {
		em.channelMap = make(map[string]chan []interface{})
	}
	if em.channelMap[eventName] == nil {
		em.channelMap[eventName] = make(chan []interface{})
	}
	go func() {
		for {
			select {
			case args, ok := <-em.channelMap[eventName]:
				if !ok {
					continue
				}
				for _, fn := range em.handlerMap[eventName] {
					fn(args...)
				}
			}
		}
	}()
}

// Emit emit event
func (em *EventEmitter) Emit(eventName string, args ...interface{}) {
	go func() {
		em.channelMap[eventName] <- args
	}()
}
