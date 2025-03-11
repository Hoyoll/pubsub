package crate

type Crate[K comparable, P any] struct {
	Packet   map[K]P
	Fallback P
}

func New[K comparable, P any]() *Crate[K, P] {
	crate := Crate[K, P]{
		Packet: make(map[K]P),
	}
	return &crate
}

func (this *Crate[K, P]) Get(name K) P {
	packet, ok := this.Packet[name]
	if !ok {
		return this.Fallback
	}
	return packet
}

func (this *Crate[K, P]) Add(name K, packet P) {
	this.Packet[name] = packet
}

func (this *Crate[K, P]) Rmv(name K) {
	_, ok := this.Packet[name]
	if ok {
		delete(this.Packet, name)
	}
}

func (this *Crate[K, P]) Fall(packet P) {
	this.Fallback = packet
}
