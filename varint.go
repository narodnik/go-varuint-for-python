package main

import (
    "bufio"
    "bytes"
    "C"
	"encoding/binary"
)

//export Add
func Add(a, b byte) byte {
    return a + b
}

//export PutVarint
func PutVarint(buffer []byte, i int64) int {
    n := binary.PutVarint(buffer[:], i)
    return n
}

//export ReadVarint
func ReadVarint(buffer []byte) int64 {
    reader := bufio.NewReader(bytes.NewReader(buffer))
    value, _ := binary.ReadVarint(reader)
    return value
}

func main() {
}
