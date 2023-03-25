package main

import (
	. "github.com/EyTest-bang/goSTL/linkedList"
	"log"
)

func cmp(a, b int) int {
	return a - b
}

func main() {
	dll := New[int](cmp)
	for i := 0; i < 100; i++ {
		if i < 50 {
			dll.InsertRight(i)
		} else {
			dll.InsertLeft(i)
		}
	}
	for i := 0; i < 100; i++ {
		if !dll.Exist(i) {
			log.Fatalf("Fail: %d should be in the list\n", i)
		}
	}
	wantRight, wantLeft := 49, 99
	for i := 0; i < 50; i++ {
		if getRight := dll.Right(); getRight != wantRight {
			log.Fatalf("Fail: want %d, get %d\n", wantRight, getRight)
		}
		if getLeft := dll.Left(); getLeft != wantLeft {
			log.Fatalf("Fail: want %d, get %d\n", wantLeft, getLeft)
		}
		wantRight--
		wantLeft--
		dll.DeleteRight()
		dll.DeleteLeft()
	}
	log.Println("PASS")
}
