package database

import (
	"backend/observer"
	"fmt"
	"sync"
)

type Observers struct {
	data map[string][]observer.IObserver
	lock sync.RWMutex
}

func (o *Observers) Init() {
	o.data = map[string][]observer.IObserver{
		"CRICKET":  {},
		"BASEBALL": {},
		"FOOTBALL": {},
		"SOCCER":   {},
		"NBA":      {},
	}
}

func (o *Observers) GetByObservable(name string) ([]observer.IObserver, error) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	data, exists := o.data[name]
	if !exists {
		return nil, fmt.Errorf("observable %s not found", name)
	}

	return data, nil
}

func (o *Observers) Set(name string, obs observer.IObserver) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	_, exists := o.data[name]
	if !exists {
		return fmt.Errorf("observable %s not found", name)
	}

	o.data[name] = append(o.data[name], obs)

	return nil
}

func (o *Observers) SetObservable(name string, observers []observer.IObserver) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	_, exists := o.data[name]
	if !exists {
		return fmt.Errorf("observable %s not found", name)
	}

	o.data[name] = observers

	return nil
}

func (o *Observers) GetAll() map[string][]observer.IObserver {
	o.lock.RLock()
	defer o.lock.Unlock()

	return o.data
}

func (o *Observers) SetAll(data map[string][]observer.IObserver) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.data = data
}
