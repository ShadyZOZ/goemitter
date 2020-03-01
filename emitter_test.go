package goemitter_test

import (
	"sync"
	"testing"

	"github.com/shadyzoz/goemitter"
	. "github.com/smartystreets/goconvey/convey"
)

type MyEmitter struct {
	goemitter.EventEmitter
}

func TestEventEmitter(t *testing.T) {
	Convey("Test EventEmitter.On", t, func() {
		Convey("Should work with a single listener and without arguments", func() {
			var (
				result int
				wg     sync.WaitGroup
			)
			emitter := MyEmitter{}
			wg.Add(1)
			emitter.On("event_a", func(args ...interface{}) {
				result++
				wg.Done()
			})
			emitter.Emit("event_a")
			wg.Wait()
			So(result, ShouldEqual, 1)
		})

		Convey("Should work with a single listener and with arguments", func() {
			var (
				result int
				wg     sync.WaitGroup
			)
			emitter := MyEmitter{}
			wg.Add(1)
			emitter.On("event_a", func(args ...interface{}) {
				result = args[0].(int)
				wg.Done()
			})
			emitter.Emit("event_a", 20)
			wg.Wait()
			So(result, ShouldEqual, 20)
		})

		Convey("Should work with no listener handlers", func() {
			emitter := MyEmitter{}
			emitter.On("event_a", nil)
			emitter.Emit("event_a", 20)
		})
	})

	Convey("Test EventEmitter.AddListener", t, func() {
		Convey("Should work like EventEmitter.On", func() {
			var (
				result int
				wg     sync.WaitGroup
			)
			emitter := MyEmitter{}
			wg.Add(1)
			emitter.AddListener("event_a", func(args ...interface{}) {
				result++
				wg.Done()
			})
			emitter.Emit("event_a")
			wg.Wait()
			So(result, ShouldEqual, 1)
		})
	})

	Convey("Test EventEmitter.EventNames", t, func() {
		Convey("Should work with 0 event", func() {
			emitter := MyEmitter{}
			So(emitter.EventNames(), ShouldBeEmpty)
		})

		Convey("Should work with 1 event", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			So(emitter.EventNames(), ShouldHaveLength, 1)
			So(emitter.EventNames(), ShouldContain, "event_a")
		})

		Convey("Should work with 2 events", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_b", dummyHandler)
			So(emitter.EventNames(), ShouldHaveLength, 2)
			So(emitter.EventNames(), ShouldContain, "event_a")
			So(emitter.EventNames(), ShouldContain, "event_b")
		})
	})

	Convey("Test EventEmitter.ListenerCount", t, func() {
		Convey("Should work with 0 listener", func() {
			emitter := MyEmitter{}
			So(emitter.ListenerCount("unkonw Listner"), ShouldEqual, 0)
		})

		Convey("Should work with 1 listener", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
		})

		Convey("Should work with 2 listeners", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 2)
		})
	})

	Convey("Test EventEmitter.Listeners", t, func() {
		Convey("Should work with 0 listener", func() {
			emitter := MyEmitter{}
			So(emitter.Listeners("unkonw Listner"), ShouldHaveLength, 0)
		})

		Convey("Should work with 1 listener", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			So(emitter.Listeners("event_a"), ShouldHaveLength, 1)
		})

		Convey("Should work with 2 listeners", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_a", dummyHandler)
			So(emitter.Listeners("event_a"), ShouldHaveLength, 2)
		})
	})

	Convey("Test EventEmitter.RemoveAllListeners", t, func() {
		Convey("Should work with nil eventNames passed in", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_b", dummyHandler)
			emitter.On("event_b", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			So(emitter.ListenerCount("event_b"), ShouldEqual, 2)
			emitter.RemoveAllListeners(nil)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
			So(emitter.ListenerCount("event_b"), ShouldEqual, 0)
		})

		Convey("Should work with eventNames", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_b", dummyHandler)
			emitter.On("event_b", dummyHandler)
			emitter.On("event_c", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			So(emitter.ListenerCount("event_b"), ShouldEqual, 2)
			So(emitter.ListenerCount("event_c"), ShouldEqual, 1)
			emitter.RemoveAllListeners([]string{"event_a", "event_b"})
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
			So(emitter.ListenerCount("event_b"), ShouldEqual, 0)
			So(emitter.ListenerCount("event_c"), ShouldEqual, 1)
		})
	})

	Convey("Test EventEmitter.RemoveListener", t, func() {
		Convey("Should work with 1 listener", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
		})

		Convey("Should work with 2 same listeners", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 2)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", nil)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
		})

		Convey("Should work with 2 different listeners", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			dummyHandler2 := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			emitter.On("event_a", dummyHandler2)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 2)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", dummyHandler2)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
		})

		Convey("Should the same event still work after clean up", func() {
			emitter := MyEmitter{}
			dummyHandler := func(...interface{}) {}
			emitter.On("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.RemoveListener("event_a", dummyHandler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
		})
	})

	Convey("Test EventEmitter.Off", t, func() {
		Convey("Should work with 1 listener", func() {
			emitter := MyEmitter{}
			var (
				result int
				wg     sync.WaitGroup
			)
			handler := func(args ...interface{}) {
				result++
				wg.Done()
			}
			emitter.On("event_a", handler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			emitter.Off("event_a", handler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 0)
			emitter.On("event_a", handler)
			So(emitter.ListenerCount("event_a"), ShouldEqual, 1)
			wg.Add(1)
			emitter.Emit("event_a")
			wg.Wait()
			So(result, ShouldEqual, 1)
		})
	})
}
