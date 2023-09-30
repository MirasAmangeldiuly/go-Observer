package main

import "fmt"

type Observer interface {
	Update(weaponName string)
}

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type Weapon struct {
	observers []Observer
	name      string
	updated   bool
}

func NewWeapon(name string) *Weapon {
	return &Weapon{name: name}
}

func (w *Weapon) Register(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *Weapon) Deregister(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *Weapon) NotifyAll() {
	for _, observer := range w.observers {
		observer.Update(w.name)
	}
}

func (w *Weapon) UpdateAvailability() {
	fmt.Printf("Новая версия оружия %s доступна!\n", w.name)
	w.updated = true
	w.NotifyAll()
}
