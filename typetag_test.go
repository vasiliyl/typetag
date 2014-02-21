package typetag

import (
	"testing"
)

type mytype struct {
	a int
	b string
}

func TestGetTag(t *testing.T) {
	reg := New()
	reg.Register("t", &mytype{})

	instance := &mytype{}
	tag, ok := reg.TagFor(instance)
	if !ok {
		t.Errorf("expected tag to be found")
	}
	if tag != "t" {
		t.Errorf("expected tag to be `t`, got `%s`", tag)
	}
}

func TestGetInstance(t *testing.T) {
	reg := New()
	reg.Register("t", &mytype{})

	instance, ok := reg.InstanceFor("t")
	if !ok {
		t.Errorf("expected tag to be registered")
	}

	switch i := instance.(type) {
	case *mytype:
		if i == nil {
			t.Errorf("expected instance to be allocated, but got nil")
		}
	default:
		t.Errorf("expected instance of `*mytype`, got `%T`", instance)
	}
}

func TestNoTag(t *testing.T) {
	reg := New()
	tag, ok := reg.TagFor(&mytype{})
	if ok {
		t.Errorf("expected tag not to be found, but got `%s`", tag)
	}
}

func TestNoInstance(t *testing.T) {
	reg := New()
	instance, ok := reg.InstanceFor("t")
	if ok {
		t.Errorf("expected tag not to be found, but got something of type `%T`", instance)
	}
}
