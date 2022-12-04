//go:build freebsd || (linux && (mips || mips64 || mips64le || mipsle || ppc64 || ppc64le))

package oss

const (
	iocNone  = 1
	iocWrite = 4
	iocRead  = 2

	iocSizebits = 13
)
