package goemitter

import (
	"reflect"
	"runtime"
	"sync"
)

// EventEmitter EventEmitterStruct
type EventEmitter struct {
	mutex      sync.Mutex
	handlerMap map[string][]func(...interface{})
}

// AddListener alias for on
func (emitter *EventEmitter) AddListener(eventName string, listener func(...interface{})) *EventEmitter {
	return emitter.On(eventName, listener)
}

// Emit Synchronously calls each of the listeners registered for the event named eventName, in the order they were registered, passing the supplied arguments to each.
// Returns true if the event had listeners, false otherwise.
func (emitter *EventEmitter) Emit(eventName string, args ...interface{}) bool {
	if len(emitter.handlerMap[eventName]) == 0 {
		return false
	}
	go func(listerners []func(...interface{})) {
		for _, listerner := range listerners {
			listerner(args...)
		}
	}(emitter.handlerMap[eventName])
	return true
}

// EventNames Returns events for registered listeners.
func (emitter *EventEmitter) EventNames() []string {
	return getHandlerMapKeys(emitter.handlerMap)
}

// // GetMaxListeners Returns the current max listener value for the EventEmitter which is either set by emitter.setMaxListeners(n) or defaults to EventEmitter.defaultMaxListeners.
// func (emitter *EventEmitter) GetMaxListeners() int {
// 	return 0
// }

// ListenerCount Returns the number of listeners listening to the event named eventName.
func (emitter *EventEmitter) ListenerCount(eventName string) int {
	return len(emitter.handlerMap[eventName])
}

// Listeners Returns a copy of the array of listeners for the event named eventName.
func (emitter *EventEmitter) Listeners(eventName string) []func(...interface{}) {
	listeners := make([]func(...interface{}), len(emitter.handlerMap[eventName]))
	copy(listeners, emitter.handlerMap[eventName])
	return listeners
}

// On Adds the listener function to the end of the listeners array for the event named eventName. No checks are made to see if the listener has already been added. Multiple calls passing the same combination of eventName and listener will result in the listener being added, and called, multiple times.
func (emitter *EventEmitter) On(eventName string, listener func(...interface{})) *EventEmitter {
	if listener == nil {
		return emitter
	}
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()
	if emitter.handlerMap == nil {
		emitter.handlerMap = make(map[string][]func(...interface{}))
	}
	emitter.handlerMap[eventName] = append(emitter.handlerMap[eventName], listener)
	return emitter
}

// Off Alias for RemoveListener().
func (emitter *EventEmitter) Off(eventName string, listener func(...interface{})) *EventEmitter {
	return emitter.RemoveListener(eventName, listener)
}

// Once Adds a one-time listener function for the event named eventName. The next time eventName is triggered, this listener is removed and then invoked.
func (emitter *EventEmitter) Once(eventName string, listener func(...interface{})) {
}

// PrependListener Adds the listener function to the beginning of the listeners array for the event named eventName. No checks are made to see if the listener has already been added. Multiple calls passing the same combination of eventName and listener will result in the listener being added, and called, multiple times.
func (emitter *EventEmitter) PrependListener(eventName string, listener func(...interface{})) {
}

// PrependOnceListener Adds a one-time listener function for the event named eventName to the beginning of the listeners array. The next time eventName is triggered, this listener is removed, and then invoked.
func (emitter *EventEmitter) PrependOnceListener(eventName string, listener func(...interface{})) {
}

// RemoveAllListeners Removes all listeners, or those of the specified eventName.
// It is bad practice to remove listeners added elsewhere in the code, particularly when the EventEmitter instance was created by some other component or module (e.g. sockets or file streams).
// Returns a reference to the EventEmitter, so that calls can be chained.
func (emitter *EventEmitter) RemoveAllListeners(eventNames []string) *EventEmitter {
	if len(eventNames) == 0 {
		eventNames = getHandlerMapKeys(emitter.handlerMap)
	}
	for _, eventName := range eventNames {
		for i := range emitter.handlerMap[eventName] {
			emitter.handlerMap[eventName][i] = nil
		}
		delete(emitter.handlerMap, eventName)
	}
	return emitter
}

// RemoveListener Removes the specified listener from the listener array for the event named eventName.
// removeListener() will remove, at most, one instance of a listener from the listener array.
// If any single listener has been added multiple times to the listener array for the specified eventName, then removeListener() must be called multiple times to remove each instance.
func (emitter *EventEmitter) RemoveListener(eventName string, listener func(...interface{})) *EventEmitter {
	for idx, registredListener := range emitter.handlerMap[eventName] {
		registeredFn := getListenerFunc(registredListener)
		listenerFn := getListenerFunc(listener)
		if registeredFn == nil || listenerFn == nil {
			continue
		}
		if registeredFn.Entry() == listenerFn.Entry() {
			emitter.handlerMap[eventName] = deleteFromListeners(emitter.handlerMap[eventName], idx)
			break
		}
	}
	// if there's no listeners under this event, then remove this event from EventEmitter
	if len(emitter.handlerMap[eventName]) == 0 && emitter.handlerMap[eventName] != nil {
		emitter.handlerMap[eventName] = nil
		delete(emitter.handlerMap, eventName)
	}
	return emitter
}

// SetMaxListeners Set MaxListeners
func (emitter *EventEmitter) SetMaxListeners(n int) {
}

// RawListeners Returns a copy of the array of listeners for the event named eventName, including any wrappers (such as those created by .once()).
func (emitter *EventEmitter) RawListeners(eventName string) {
}

func deleteFromListeners(a []func(...interface{}), i int) []func(...interface{}) {
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = nil // or the zero value of T
	a = a[:len(a)-1]
	return a
}

func getListenerFunc(listener func(...interface{})) *runtime.Func {
	return runtime.FuncForPC(reflect.ValueOf(listener).Pointer())
}

func getHandlerMapKeys(handlerMap map[string][]func(...interface{})) []string {
	if len(handlerMap) == 0 {
		return nil
	}
	keys := make([]string, len(handlerMap))
	i := 0
	for k := range handlerMap {
		keys[i] = k
		i++
	}
	return keys
}
