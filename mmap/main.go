package main

import (
	"log"

	"github.com/k-nhr/mmap/lib"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	name := "/tmp/hoge.map"
	newmap := mmap.NewMmap()
	if err := newmap.Open(name, mmap.ReadWrite, 0x200); err != nil {
		log.Fatal(err)
	}
	newmap.Write([]byte("hoge"))
	bytes := newmap.Read(len("hoge"))
	log.Println(string(bytes))
	newmap.Write([]byte("hoge fugaugu"))
	bytes = newmap.Read(len("hoge fugaugu"))
	log.Println(string(bytes))
	newmap.Close()
}
