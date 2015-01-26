package cgzip

import "C"

import (
	"hash"
	"unsafe"
)

type adler32Hash struct {
	adler C.uLong
}

func NewAdler32() hash.Hash32 {
	a := &adler32Hash{}
	a.Reset()
	return a
}

func (a *adler32Hash) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		a.adler = C.adler32(a.adler, (*C.Bytef)(unsafe.Pointer(&p[0])), (C.uInt)(len(p)))
	}
	return len(p), nil
}

func (a *adler32Hash) Sum(b []byte) []byte {
	s := a.Sum32()
	b = append(b, byte(s>>24))
	b = append(b, byte(s>>16))
	b = append(b, byte(s>>8))
	b = append(b, byte(s))
	return b
}

func (a *adler32Hash) Reset() {
	a.adler = C.adler32(0, (*C.Bytef)(unsafe.Pointer(nil)), 0)
}

func (a *adler32Hash) Size() int {
	return 4
}

func (a *adler32Hash) BlockSize() int {
	return 1
}

func (a *adler32Hash) Sum32() uint32 {
	return uint32(a.adler)
}

func Adler32Combine(adler1, adler2 uint32, len2 int) uint32 {
	return uint32(C.adler32_combine(C.uLong(adler1), C.uLong(adler2), C.z_off_t(len2)))
}
