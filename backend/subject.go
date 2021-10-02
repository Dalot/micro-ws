package main

import (
	"errors"
	"fmt"
)

type packet struct {
	name string // The name of the observable, say 'CRICKET' or 'FOOTBALL'
	msg  []byte
}

type ObserversTable map[string][]*observer

var observers = ObserversTable{
	"CRICKET":  []*observer{},
	"BASEBALL": []*observer{},
	"FOOTBALL": []*observer{},
	"SOCCER":   []*observer{},
	"NBA":      []*observer{},
}

func newPacket(msg []byte) *packet {
	return &packet{
		msg: msg,
	}
}

func (p *packet) register(o observer, name string) error {
	if name == "" {
		return errors.New("name must not be empty")
	}

	for observableName, list := range observers {
		if name == observableName {
			list = append(list, &o)
			observers[observableName] = list
			return nil
		}
	}

	return fmt.Errorf("could not register, did not find any observable named %s", name)
}

func (p *packet) deregister(o observer, name string) error {
	list := observers[name]
	for i, obs := range list {
		if o.getID() == (*obs).getID() {
			list = append(list[0:i], list[i+1:]...)
			observers[name] = list
			return nil
		}
	}

	return fmt.Errorf("could not deregister, did not find observer with id %s on the observable %s", o.getID(), name)
}

func (p *packet) notifyAll() {
	for _, observer := range observers[p.name] {
		(*observer).update(p.msg)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
