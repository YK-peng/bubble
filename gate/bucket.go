package gate

import (
	"gate/conf"
	"sync"
)

//Channel桶
type Bucket struct {
	//c     *conf.Bucket
	cLock sync.RWMutex        // 保护chs
	chs   map[string]*Channel // pair<subKey, *channel>
}

func NewBucket(c *conf.Bucket) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel, c.Channel)
	//b.c = c
	return
}

func (b *Bucket) ChannelCount() int {
	return len(b.chs)
}

func (b *Bucket) Put(ch *Channel) (err error) {
	b.cLock.Lock()
	// close old channel
	if dch := b.chs[ch.Key]; dch != nil {
		dch.Close()
	}
	b.chs[ch.Key] = ch
	b.cLock.Unlock()
	return
}

func (b *Bucket) Del(dch *Channel) {
	var (
		ok bool
		ch *Channel
	)
	b.cLock.Lock()
	if ch, ok = b.chs[dch.Key]; ok {
		if ch == dch {
			delete(b.chs, ch.Key)
		}
	}
	b.cLock.Unlock()
}

func (b *Bucket) Channel(key string) (ch *Channel) {
	b.cLock.RLock()
	ch = b.chs[key]
	b.cLock.RUnlock()
	return
}