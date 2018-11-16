package gow32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
	procReleaseMutex = kernel32.NewProc("ReleaseMutex")
)

func ReleaseMutex(id uintptr) error {
	_, _, err := procReleaseMutex.Call(
		id,
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return nil
	default:
		return err
	}
}

func CreateMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}
