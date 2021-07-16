package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/k-nhr/binary-parser/builder"
)

func main() {
	// [変数名]:[バイト列のインデックス]:[変数の型]:[型に依存した設定内容]
	f := "Flag:0:bool:7 Temp:1:int:13:/10 Long:2:float:32"
	s := strings.Split(f, " ")

	b := builder.NewStructBuilder()

	for _, v := range s {
		p := strings.Split(v, ":")
		switch p[2] {
		case "bool":
			b.AddField(p[0], reflect.TypeOf(true))
		case "int":
			b.AddField(p[0], reflect.TypeOf(1))
		case "float":
			b.AddField(p[0], reflect.TypeOf(1.1))
		}

	}

	person := b.Build()

	i := person.NewInstance()
	i.SetBool("Flag", true)
	i.SetInt("Temp", 8)
	i.SetFloat("Long", 23.45)

	fmt.Println(i.Value())
	fmt.Println(i.Pointer())
}
