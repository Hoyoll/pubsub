package vemiter

type container[KEY comparable, ARGS any] struct {
	Worker map[KEY]*func(ARGS)
	Error  *func(ARGS)
}

func New[KEY comparable, ARGS any]() *container[KEY, ARGS] {
	error := func(a ARGS) {
		println("Uncaught error here!")
	}
	container := container[KEY, ARGS]{
		Worker: make(map[KEY]*func(ARGS)),
		Error:  &error,
	}
	return &container
}

func (this *container[KEY, ARGS]) Store(name KEY, process *func(ARGS)) *container[KEY, ARGS] {
	this.Worker[name] = process
	return this
}

func (this *container[KEY, ARGS]) Get(name KEY) *func(ARGS) {
	result, ok := this.Worker[name]
	if !ok {
		return this.Error
	}
	return result
}

func (this *container[KEY, ARGS]) Catch(process *func(ARGS)) *container[KEY, ARGS] {
	this.Error = process
	return this
}

func (this *container[KEY, ARGS]) Emit(name KEY, dependency ARGS) {
	result, ok := this.Worker[name]
	if !ok {
		(*this.Error)(dependency)
	} else {
		(*result)(dependency)
	}
}

func (this *container[KEY, ARGS]) All(dependency ARGS) {
	for _, worker := range this.Worker {
		(*worker)(dependency)
	}
}

func (this *container[KEY, ARGS]) Remove(name KEY) {
	_, ok := this.Worker[name]
	if ok {
		delete(this.Worker, name)
	} else {
		println("No worker named", name)
	}
}
