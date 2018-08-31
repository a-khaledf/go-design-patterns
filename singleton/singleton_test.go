package singleton

import (
	"testing"
)

func TestNew(t *testing.T) {
	singleton := New("test")
	if singleton.Name != "test" {
		t.Errorf("New is not working properly, %s not correct", singleton.Name)
	}

	anotherSingleton := New("anotherTest")
	if anotherSingleton.Name != singleton.Name {
		t.Errorf("Singleton is not working properly")
	}
}
