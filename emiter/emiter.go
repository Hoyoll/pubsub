package emiter

type container[KEY comparable, ARGS any, RESULT any] struct {
	Worker map[KEY]*func(ARGS) RESULT
	Error  *func(ARGS) RESULT
}

func New[KEY comparable, ARGS any, RESULT any]() *container[KEY, ARGS, RESULT] {
	container := container[KEY, ARGS, RESULT]{
		Worker: make(map[KEY]*func(ARGS) RESULT),
	}
	return &container
}

func (this *container[KEY, ARGS, RESULT]) Store(name KEY, process *func(ARGS) RESULT) *container[KEY, ARGS, RESULT] {
	this.Worker[name] = process
	return this
}

func (this *container[KEY, ARGS, RESULT]) Get(name KEY) *func(ARGS) RESULT {
	result, ok := this.Worker[name]
	if !ok {
		return this.Error
	}
	return result
}

func (this *container[KEY, ARGS, RESULT]) Catch(process *func(ARGS) RESULT) *container[KEY, ARGS, RESULT] {
	this.Error = process
	return this
}

func (this *container[KEY, ARGS, RESULT]) Emit(name KEY, dependency ARGS) RESULT {
	result, ok := this.Worker[name]
	if !ok {
		return (*this.Error)(dependency)
	}
	return (*result)(dependency)
}

func (this *container[KEY, ARGS, RESULT]) All(dependency ARGS) map[KEY]RESULT {
	m := make(map[KEY]RESULT)
	for key, process := range this.Worker {
		m[key] = (*process)(dependency)
	}
	return m
}

func (this *container[KEY, ARGS, RESULT]) Remove(name KEY) {
	_, ok := this.Worker[name]
	if ok {
		delete(this.Worker, name)
	} else {
		println("No worker named", name)
	}
}