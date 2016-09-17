package gow32

import (
	"syscall"
	"testing"
)

var (
	mutextName = "TGoKernel32"
)

func Test_CreateSemaphore(t *testing.T) {
	_, err := CreateMutex(mutextName)
	if err != nil {
		t.Fatalf("Error creating mutex: %s", err.Error())
	}

	_, err = CreateMutex(mutextName)
	if err == nil {
		t.Errorf("CreateMutex should return an error")
	}
	if n := int(err.(syscall.Errno)); n != ERROR_ALREADY_EXISTS {
		t.Errorf("CreateMutex error should be %d, got %d", ERROR_ALREADY_EXISTS, n)
	}
}
