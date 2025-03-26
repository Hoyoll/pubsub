package crate

import "sync"

type Crate[K comparable, P any] struct {
	Packet   *sync.Map
	Fallback P
}

func New[K comparable, P any]() *Crate[K, P] {
	crate := &Crate[K, P]{
		Packet: &sync.Map{},
	}
	return crate
}

func (this *Crate[K, P]) Get(name K) P {
	packet, ok := this.Packet.Load(name)
	if !ok {
		return this.Fallback
	}
	return packet.(P)
}

func (this *Crate[K, P]) Add(name K, packet P) {
	this.Packet.Store(name, packet)
}

func (this *Crate[K, P]) Rmv(name K) {
	this.Packet.Delete(name)
}

func (this *Crate[K, P]) Iter(process func(K, P) bool) {
	this.Packet.Range(func(key, value any) bool {
		return process(key.(K), value.(P))
	})
}

func (this *Crate[K, P]) Fall(packet P) {
	this.Fallback = packet
}
