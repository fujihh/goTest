package awesome_test

import (
	"gohelloworld/awesome"
	"testing"
)

func TestHello(t *testing.T) {
	if awesome.Hello() != "Hello, world." {
		t.Fatal("Wrong output")
	}
}
