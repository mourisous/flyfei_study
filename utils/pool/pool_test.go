package pool

import (
	"fmt"
	"testing"
)

func TestNewPoolWithSize(t *testing.T) {
	pool := NewPoolWithSize(1)
	defer pool.Release()

	pool.RunGo(func() {
		println("hello world1")
		panic("test")
	})
	pool.RunGo(func() {
		println("hello world2")
		panic("test")
	})

	pool.Wait()

	fmt.Println("done all")
}
