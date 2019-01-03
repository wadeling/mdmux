package main

import (
	"errors"
	"log"
)

type Trigger struct {
	EventList [][]*Event
}

const MaxEventNum = 10

func (tr *Trigger) DumpEventList() {
	for i, v := range tr.EventList {
		//log.Println("i", i)
		for h, k := range v {
			src := ""
			uuid := ""
			if k != nil {
				src = k.Src
				uuid = k.UUid

				log.Printf("i:%v,h:%v,uuid:%s,src:%s", i, h, uuid, src)
			}
			//log.Printf("i:%v,h:%v,uuid:%s,src:%s", i, h, uuid, src)
		}
	}
}

func NewTrigger() *Trigger {
	tr := &Trigger{
		EventList: make([][]*Event, MaxEventNum),
	}

	for i := range tr.EventList {
		tmp := make([]*Event, MaxEventNum)
		tr.EventList[i] = tmp
	}

	return tr
}

func (tr *Trigger) addSrcEvent(evList []*Event, ev *Event) error {
	for i, v := range evList {
		if v == nil {
			evList[i] = ev
			return nil
		}
	}
	evList = append(evList, ev)
	return nil
}

func (tr *Trigger) AddEvent(ev *Event) error {
	for i, v := range tr.EventList {
		if v[0] == nil {
			//empty,add new event
			v[0] = ev
			return nil
		}

		if v[0].Src == ev.Src {
			//tr.EventList[i] = append(v, ev)
			//return nil
			return tr.addSrcEvent(tr.EventList[i], ev)
		}

	}
	return errors.New("event queue is full")
}

func (tr *Trigger) PopTopEvents() ([]*Event, error) {
	result := make([]*Event, MaxEventNum)
	for _, v := range tr.EventList {
		if v[0] != nil {
			result = append(result, v[0])
			v = append(v[:1], v[2:]...)
		}
	}
	return result, nil
}

func (tr *Trigger) GetEventsFront() ([]*Event, error) {
	result := make([]*Event, MaxEventNum)
	for _, v := range tr.EventList {
		if v[0] != nil {
			result = append(result, v[0])
		}
	}
	return result, nil
}
