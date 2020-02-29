# goemitter

[![Build Status](https://travis-ci.org/ShadyZOZ/goemitter.svg?branch=master)](https://travis-ci.org/ShadyZOZ/goemitter)
[![codecov](https://codecov.io/gh/ShadyZOZ/goemitter/branch/master/graph/badge.svg)](https://codecov.io/gh/ShadyZOZ/goemitter)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShadyZOZ/goemitter)](https://goreportcard.com/report/github.com/ShadyZOZ/goemitter)

A go EventEmitter inspired from Node.js's EventEmitter.

---

Development in progress, plan to implement most is not all of the Node.js's EventEmitter api.

## node api list

- [ ] Event: 'newListener'
- [ ] Event: 'removeListener'
- [ ] ~~EventEmitter.listenerCount(emitter, eventName)~~ `deprecated`
- [ ] EventEmitter.defaultMaxListeners
- [x] emitter.addListener(eventName, listener)
- [x] emitter.emit(eventName[, ...args])
- [x] emitter.eventNames()
- [ ] emitter.getMaxListeners()
- [x] emitter.listenerCount(eventName)
- [x] emitter.listeners(eventName)
- [x] emitter.off(eventName, listener)
- [x] emitter.on(eventName, listener)
- [ ] emitter.once(eventName, listener)
- [ ] emitter.prependListener(eventName, listener)
- [ ] emitter.prependOnceListener(eventName, listener)
- [ ] emitter.removeAllListeners([eventName])
- [x] emitter.removeListener(eventName, listener)
- [ ] emitter.setMaxListeners(n)
- [ ] emitter.rawListeners(eventName)
- [ ] ~~emitter\[Symbol.for('nodejs.rejection')\](err, eventName[, ...args])~~ `experimental`

## TODOs

- [ ] Current implement of event handling will not promise the correct execution of `Emit` called before `removeAllListeners` or `removeListener`, thus a different approach needs to be done
- [ ] Error handling mechanism
