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

package ipfs

import (
	"encoding/json"
	"github.com/Loopring/ringminer/config"
	"github.com/Loopring/ringminer/eventemiter"
	"github.com/Loopring/ringminer/log"
	"github.com/Loopring/ringminer/types"
	"github.com/ipfs/go-ipfs-api"
	"sync"
)

type IPFSListener struct {
	options config.IpfsOptions
	sh      *shell.Shell
	subs    []*shell.PubSubSubscription
	stop    chan struct{}
	lock    sync.RWMutex
}

func NewListener(options config.IpfsOptions) *IPFSListener {
	l := &IPFSListener{}

	l.options = options

	l.sh = shell.NewLocalShell()
	for _, topic := range options.Topics {
		sub, err := l.sh.PubSubSubscribe(topic)
		if err != nil {
			log.Fatalf("ipfs listener create sub scribe error:%s", err.Error())
		}
		l.subs = append(l.subs, sub)
	}

	return l
}

func (l *IPFSListener) Start() {
	l.stop = make(chan struct{})

	for i := 0; i < len(l.subs); i++ {
		go func(i int) {
			for {
				if record, err := l.subs[i].Next(); nil != err {
					log.Errorf("err:%s", err.Error())
				} else {
					data := record.Data()
					ord := &types.Order{}
					err := json.Unmarshal(data, ord)
					if err != nil {
						log.Errorf("failed to accept data %s", err.Error())
					} else {
						log.Debugf("accept data from ipfs %s", string(data))
						eventemitter.Emit(eventemitter.OrderBookPeer, ord)
					}
				}
			}
		}(i)
	}
}

func (listener *IPFSListener) Stop() {
	listener.lock.Lock()
	close(listener.stop)
	listener.lock.Unlock()
}

func (listener *IPFSListener) Restart() {

}

func (listener *IPFSListener) Name() string {
	return "ipfs-listener"
}
