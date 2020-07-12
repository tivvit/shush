package backend

import (
	"errors"
	"sync"
	"time"
)

type data struct {
	Data string     `json:"data"`
	Ttl  *time.Time `json:"ttl,omitempty"`
}

type InMem struct {
	data   map[string]data
	close  chan bool
	closed bool
	wg     sync.WaitGroup
	lock   sync.RWMutex
}

func NewInMem() *InMem {
	wg := sync.WaitGroup{}
	im := &InMem{
		wg:     wg,
		lock:   sync.RWMutex{},
		data:   map[string]data{},
		closed: false,
		close:  make(chan bool),
	}
	wg.Add(1)
	go im.curator()
	return im
}

func (im InMem) Get(key string) (string, error) {
	if im.closed {
		return "", errors.New("closed")
	}
	im.lock.RLock()
	defer im.lock.RUnlock()
	if v, ok := im.data[key]; ok {
		if v.Ttl != nil && time.Now().After(*v.Ttl) {
			return "", errors.New("key not found")
		}
		return v.Data, nil
	} else {
		return "", errors.New("key not found")
	}
}

func (im InMem) GetAll() (map[string]string, error) {
	if im.closed {
		return nil, errors.New("closed")
	}
	im.lock.RLock()
	defer im.lock.RUnlock()
	d := make(map[string]string, len(im.data))
	t := time.Now()
	for k, v := range im.data {
		if v.Ttl != nil && t.After(*v.Ttl) {
			continue
		}
		d[k] = v.Data
	}
	return d, nil
}

func (im InMem) Set(key string, value string, ttl time.Duration) error {
	if im.closed {
		return errors.New("closed")
	}
	im.lock.Lock()
	defer im.lock.Unlock()
	var t *time.Time
	if ttl > 0 {
		tn := time.Now().Add(ttl)
		t = &tn
	}
	im.data[key] = data{
		Data: value,
		Ttl:  t,
	}
	return nil
}

func (im *InMem) Close() error {
	im.close <- true
	im.wg.Wait()
	im.closed = true
	return nil
}

func (im *InMem) curator() {
	defer im.wg.Done()
	for {
		select {
		case <-im.close:
			break
		default:
			d := make(map[string]data, len(im.data))
			t := time.Now()
			for k, v := range im.data {
				if v.Ttl != nil && t.After(*v.Ttl) {
					continue
				}
				d[k] = v
			}
			im.lock.Lock()
			im.data = d
			im.lock.Unlock()
		}
		time.Sleep(500 * time.Millisecond)
	}
}
