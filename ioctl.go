package oss

import (
	"syscall"
	"unsafe"
)

const (
	iocNrbits   = 8
	iocTypebits = 8
	iocNrshift  = 0

	iocTypeshift = iocNrshift + iocNrbits
	iocSizeshift = iocTypeshift + iocTypebits
	iocDirshift  = iocSizeshift + iocSizebits
	iocSizemask  = (1 << iocSizebits) - 1
)

const (
	iocIn    = iocWrite << iocDirshift
	iocOut   = iocRead << iocDirshift
	iocInout = iocIn | iocOut
)

func _io(x, y int) int {
	return iocNone | (x << 8) | y
}

func _ior(x, y int) int {
	var t int32

	return iocOut | ((int(unsafe.Sizeof(t)) & iocSizemask) << 16) | (x << 8) | y
}

func _iowr(x, y int) int {
	var t int32

	return iocInout | ((int(unsafe.Sizeof(t)) & iocSizemask) << 16) | (x << 8) | y
}

func _iori(x, y int) int {
	var t audioBufInfo

	return iocOut | ((int(unsafe.Sizeof(t)) & iocSizemask) << 16) | (x << 8) | y
}

func ioctl(fd uintptr, req int, val int) (int, error) {
	var err error
	v := int32(val)

	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(unsafe.Pointer(&v)))
	if e != 0 {
		err = e
	}

	return int(v), err
}

func ioctlI(fd uintptr, req int) (audioBufInfo, error) {
	var err error
	var info audioBufInfo

	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(unsafe.Pointer(&info)))
	if e != 0 {
		err = e
	}

	return info, err
}
