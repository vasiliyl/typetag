package typetag

import (
	"errors"
	"reflect"
)

var (
	ErrUnsupportedType = errors.New("unsupported instance type")
)

type Registry struct {
	m map[string]reflect.Type
}

func New() *Registry {
	return &Registry{
		m: map[string]reflect.Type{},
	}
}

func (r *Registry) typeOf(instance interface{}) reflect.Type {
	ptr := reflect.TypeOf(instance)
	if ptr.Kind() != reflect.Ptr {
		panic(ErrUnsupportedType)
	}

	str := ptr.Elem()
	if str.Kind() != reflect.Struct {
		panic(ErrUnsupportedType)
	}

	return str
}

// На данном этапе поддерживаем исключительно указатели на структуры
func (r *Registry) Register(tag string, instance interface{}) {
	r.m[tag] = r.typeOf(instance)
}

func (r *Registry) TagFor(instance interface{}) (tag string, ok bool) {
	t := r.typeOf(instance)
	for tag, typ := range r.m {
		if typ == t {
			return tag, true
		}
	}
	return
}

func (r *Registry) InstanceFor(tag string) (instance interface{}, ok bool) {
	typ, ok := r.m[tag]
	if !ok {
		return
	}
	return reflect.New(typ).Interface(), true
}
