package crate

type crate[K comparable, P any] struct {
	Packet   map[K]P
	Fallback P
}

func New[K comparable, P any]() *crate[K, P] {
	crate := crate[K, P]{
		Packet: make(map[K]P),
	}
	return &crate
}

func (this *crate[K, P]) Get(name K) P {
	packet, ok := this.Packet[name]
	if !ok {
		return this.Fallback
	}
	return packet
}

func (this *crate[K, P]) Add(name K, packet P) {
	this.Packet[name] = packet
}

func (this *crate[K, P]) Rmv(name K) {
	_, ok := this.Packet[name]
	if ok {
		delete(this.Packet, name)
	}
}

func (this *crate[K, P]) Fall(packet P) {
	this.Fallback = packet
}
