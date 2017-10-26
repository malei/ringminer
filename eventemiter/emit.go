/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package eventemitter

import (
	"sync"
)

type Topic int

func (t Topic) Name() string {
	switch t {
	case RingMined:
		return "RingMined"
	case OrderCanceled:
		return "OrderCanceled"
	case OrderFilled:
		return "OrderFilled"
	case Fork:
		return "Fork"
	case RingSubmitFailed:
		return "RingSubmitFailed"
	case Transaction:
		return "Transaction"
	default:
		return "EventEmit"
	}
}

const (
	RingMined Topic = iota
	OrderCanceled
	OrderFilled
	Fork
	RingSubmitFailed
	Transaction
)

var watchers map[string][]*Watcher
var mtx *sync.Mutex

type EventData interface{}

type Watcher struct {
	Concurrent bool
	Handle     func(eventData EventData) error
}

func Un(topic string, watcher *Watcher) {
	watchersTmp := []*Watcher{}
	for _, w := range watchers[topic] {
		if w != watcher {
			watchersTmp = append(watchersTmp, w)
		}
	}
	watchers[topic] = watchersTmp
}

func On(topic string, watcher *Watcher) {
	mtx.Lock()
	defer mtx.Unlock()
	if _, ok := watchers[topic]; !ok {
		watchers[topic] = make([]*Watcher, 0)
	}
	watchers[topic] = append(watchers[topic], watcher)
}

func Emit(topic string, eventData EventData) {
	for _, ob := range watchers[topic] {
		if ob.Concurrent {
			go ob.Handle(eventData)
		} else {
			ob.Handle(eventData)
		}
	}
}

func init() {
	watchers = make(map[string][]*Watcher)
	mtx = &sync.Mutex{}
}