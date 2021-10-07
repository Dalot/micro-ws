package observer

import (
	"fmt"
	"github.com/pkg/errors"
)

//go:generate moq -out ./mock_observer.go . IObserver:IObserverMock
type IObserver interface {
	Update([]byte)
	GetID() string
}

type Observer struct {
	ID string
}

//go:generate moq -out ./mock_database.go . Database:databaseMock
type Database interface {
	GetByObservable(name string) ([]IObserver, error)
	SetObservable(name string, observers []IObserver) error
	Set(name string, obs IObserver) error
	GetAll() map[string][]IObserver
	SetAll(map[string][]IObserver)
}

type Packet struct {
	Name string // The name of the observable, say 'CRICKET' or 'FOOTBALL'
	Msg  []byte
	DB   Database
}

func newPacket(msg []byte) *Packet {
	return &Packet{
		Msg: msg,
	}
}

func (p *Packet) Register(o IObserver, name string) error {
	if name == "" {
		return errors.New("name must not be empty")
	}

	err := p.DB.Set(name, o)
	if err != nil {
		return errors.Wrap(err, "register observer")
	}

	return nil
}

func (p *Packet) Deregister(o IObserver, name string) error {
	var list []IObserver
	var err error

	if name == "" {
		return errors.New("name must not be empty")
	}

	if list, err = p.DB.GetByObservable(name); err != nil {
		return errors.Wrap(err, "deregister observer")
	}

	for i, obs := range list {
		if o.GetID() == obs.GetID() {
			list = append(list[0:i], list[i+1:]...)
			p.DB.SetObservable(name, list)
			return nil
		}
	}

	return fmt.Errorf("could not deregister, did not find any observer with id %s on the observable %s", o.GetID(), name)
}

func (p *Packet) NotifyAll() error {
	observers, err := p.DB.GetByObservable(p.Name)
	if err != nil {
		return errors.Wrap(err, "notify all observers")
	}
	for _, observer := range observers {
		observer.Update(p.Msg)
	}

	return nil
}