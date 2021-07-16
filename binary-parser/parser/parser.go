package parser

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

// Bytes2str converts []byte to string
func Bytes2str(bytes ...byte) string {
	strs := []string{}
	for _, b := range bytes {
		strs = append(strs, fmt.Sprintf("%02x", b))
	}
	return strings.Join(strs, " ")
}

// Bytes2uint32 converts []byte to uint32
func Bytes2uint32(bytes ...byte) uint32 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint32(append(padding, bytes...))
	return i
}

// Bytes2int32 converts []byte to int32
func Bytes2int32(bytes ...byte) int32 {
	if 0x7f < bytes[0] {
		mask := uint32(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := Bytes2uint32(bytes...)
		i = (^i + 1) & mask
		return int32(-i)

	} else {
		i := Bytes2uint32(bytes...)
		return int32(i)
	}
}

// Bytes2uint64 converts []byte to uint64
func Bytes2uint64(bytes ...byte) uint64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	return i
}

// Bytes2int64 converts []byte to int64
func Bytes2int64(bytes ...byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := Bytes2uint64(bytes...)
		i = (^i + 1) & mask
		return int64(-i)

	} else {
		i := Bytes2uint64(bytes...)
		return int64(i)
	}
}

// Bytes2float32 converts []byte to float32
func Bytes2float32(bytes ...byte) float32 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint32(append(padding, bytes...))
	f32 := math.Float32frombits(i)
	return f32
}

// Bytes2float64 converts []byte to float64
func Bytes2float64(bytes ...byte) float64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	f64 := math.Float64frombits(i)
	return f64
}
