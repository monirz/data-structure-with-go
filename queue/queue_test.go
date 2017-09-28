package main

import "testing"

func TestQueue(t *testing.T) {
	enQueue("Go FTW!")

	if size != 1 {
		t.Errorf("error, got: %d, want: %d.", size, 1)
	}

	//one item allready en-queued
	for i := 0; i < MAX-1; i++ {
		enQueue(i)
	}
	if size != MAX {
		t.Errorf("error, got: %d, want: %d.", size, MAX)
	}

	err := enQueue("another item")

	if err == nil {
		t.Errorf("error, queue is suppose to be full with error")
	}

	deQueue()

	if size != MAX-1 {
		t.Errorf("error, got: %d, want: %d.", size, MAX)
	}

	err = enQueue("another item")

	if err != nil {
		t.Errorf("error, got: %v, want: %v.", nil, err)
	}
}
