package mmap

import (
	"log"
	"os"
	"syscall"
	"unsafe"
)

type OpenMode int

const (
	ReadWrite OpenMode = iota
	ReadOnly
)

type Item struct {
	Object interface{}
}

const (
	InvalidHandle = -1
	num           = 1e2
)

type Mmap struct {
	fileName string
	size     int
	offset   int64
	mode     OpenMode
	fd       *os.File
	addr     []byte
	item     map[string]Item
}

func NewMmap() *Mmap {
	return &Mmap{
		item: make(map[string]Item),
	}
}

func (m *Mmap) Open(filename string, openMode OpenMode, cnt int) error {
	size := int(unsafe.Sizeof(byte(0))) * num

	map_file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = map_file.Seek(int64(size-1), 0)
	if err != nil {
		log.Println(err)
		m.Close()
		return err
	}
	_, err = map_file.Write([]byte(" "))
	if err != nil {
		log.Println(err)
		m.Close()
		return err
	}

	var prot int
	if openMode == ReadWrite {
		prot = syscall.PROT_READ|syscall.PROT_WRITE
	} else {
		prot = syscall.PROT_READ
	}
	mmap, err := syscall.Mmap(int(map_file.Fd()), 0, size, prot, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		m.Close()
		return err
	}
	map_array := (*[num]byte)(unsafe.Pointer(&mmap[0]))
	m.fd = map_file
	m.addr = map_array[:num]
	return nil
}

func (m *Mmap) Close() {
	if len(m.addr) != 0 {
		_ = syscall.Munmap(m.addr)
	}
	if m.fd != nil {
		_ = m.fd.Close()
	}
}

func (m *Mmap) Write(data []byte) {
	for i := 0; i < len(data); i++ {
		m.addr[i] = data[i]
	}
}

func (m *Mmap) Read(size int) []byte {
	return m.addr[:size]
}
