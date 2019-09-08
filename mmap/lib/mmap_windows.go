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
const num           = 1e2

type Item struct {
	Object interface{}
}

type Mmap struct {
	fileName string
	size int
	mode OpenMode
	fileHandle syscall.Handle
	mapHandle syscall.Handle
	addr []byte
	item map[string]Item
}

func NewMmap() *Mmap {
	return &Mmap{
		fileHandle: syscall.InvalidHandle,
		mapHandle: syscall.InvalidHandle,
		item: make(map[string]Item),
	}
}

func (m *Mmap) Open(filename string, openMode OpenMode, cnt int) error {
	size := int(unsafe.Sizeof(byte(0))) * num

	m.fileName = filename
	access := uint32(syscall.GENERIC_READ | syscall.GENERIC_WRITE)
	mode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE)

	fh, err := syscall.CreateFile(syscall.StringToUTF16Ptr(filename), access, mode, nil, syscall.CREATE_NEW, 0, 0)
	if err != nil {
		if !os.IsExist(err) {
			log.Println(err)
			m.Close()
			return err
		}
		if err := syscall.DeleteFile(syscall.StringToUTF16Ptr(filename)); err != nil {
			log.Println(err)
			m.Close()
			return err
		}
		if fh, err = syscall.CreateFile(syscall.StringToUTF16Ptr(filename), access, mode, nil, syscall.CREATE_NEW, 0, 0); err != nil {
			log.Println(err)
			m.Close()
			return err
		}
		if _, err := syscall.SetFilePointer(fh, int32(size), nil, syscall.FILE_BEGIN); err != nil {
			log.Println(err)
			m.Close()
			return err
		}
		if err := syscall.SetEndOfFile(fh); err != nil {
			log.Println(err)
			m.Close()
			return err
		}
	}
	m.fileHandle = fh

	var prot uint32
	if openMode == ReadWrite {
		prot = syscall.PAGE_READWRITE
		access = syscall.FILE_MAP_READ | syscall.FILE_MAP_WRITE
	} else {
		prot = syscall.PAGE_READONLY
		access = syscall.FILE_MAP_READ
	}
	mh, err := syscall.CreateFileMapping(fh, nil, prot, 0, uint32(size), nil)
	if  err != nil {
		log.Println(err)
		m.Close()
		return err
	}
	m.mapHandle = mh

	addr, err := syscall.MapViewOfFile(mh, access, 0, 0, uintptr(size))
	if err != nil {
		log.Println(err)
		m.Close()
		return err
	}
	//map_array := (*[num]byte)(unsafe.Pointer(addr))
	map_array := (*[num]byte)(unsafe.Pointer(addr))[:size]

	for i := 0; i < num; i++ {
		map_array[i] = byte(i)
	}
	log.Println(map_array)

	m.addr = map_array[:num]
	m.mode = openMode
	m.size = size
	return nil
}

func (m *Mmap) Close()  {
	if len(m.addr) != 0 {
		_ = syscall.UnmapViewOfFile(uintptr(unsafe.Pointer(&m.addr)))
	}
	if m.mapHandle != syscall.InvalidHandle {
		_ = syscall.CloseHandle(m.mapHandle)
		m.mapHandle = syscall.InvalidHandle
	}
	if m.fileHandle != syscall.InvalidHandle {
		_ = syscall.CloseHandle(m.fileHandle)
		m.fileHandle = syscall.InvalidHandle
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
