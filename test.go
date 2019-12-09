package main

import (
    "bufio"
    "bytes"
	"encoding/binary"
    "fmt"
	"io"
    "math"
)

const (
    varintTwoBytes = 0xfd
    varintFourBytes = 0xfe
    varintEightBytes = 0xff
)

func PutVarUintSimple(buffer []byte, value uint64) uint {
    if value < varintTwoBytes {
        //fmt.Println("put 1 byte")
        buffer[0] = uint8(value)
        return 1
    } else if value <= math.MaxUint16 {
        //fmt.Println("put 2 bytes")
        buffer[0] = uint8(varintTwoBytes)
        binary.BigEndian.PutUint16(buffer[1:], uint16(value))
        return 3
    } else if value <= math.MaxUint32 {
        //fmt.Println("put 4 bytes")
        buffer[0] = uint8(varintFourBytes)
        binary.BigEndian.PutUint32(buffer[1:], uint32(value))
        return 5
    } else {
        //fmt.Println("put 8 bytes")
        buffer[0] = uint8(varintEightBytes)
        binary.BigEndian.PutUint64(buffer[1:], uint64(value))
        return 9
    }
}

func ReadNBytes(reader io.ByteReader, buffer []byte, number_bytes int) error {
    for i := 0; i < number_bytes; i++ {
        //fmt.Println("  read byte #", i)
        value, err := reader.ReadByte()
        if err != nil {
            return err
        }
        buffer[i] = value
    }
    return nil
}

func ReadVarUintSimple(reader io.ByteReader) (uint64, error) {
    value, err := reader.ReadByte()
    if err != nil {
        return 0, err
    }

    switch value {
    case varintEightBytes:
        //fmt.Println("read 8 bytes")
        buffer := make([]byte, 8)
        err = ReadNBytes(reader, buffer, 8)
        if err != nil {
            return 0, err
        }
        return binary.BigEndian.Uint64(buffer), nil
    case varintFourBytes:
        //fmt.Println("read 4 bytes")
        buffer := make([]byte, 4)
        err = ReadNBytes(reader, buffer, 4)
        if err != nil {
            return 0, err
        }
        return uint64(binary.BigEndian.Uint32(buffer)), nil
    case varintTwoBytes:
        //fmt.Println("read 2 bytes")
        buffer := make([]byte, 2)
        err = ReadNBytes(reader, buffer, 2)
        if err != nil {
            return 0, err
        }
        return uint64(binary.BigEndian.Uint16(buffer)), nil
    default:
        //fmt.Println("return 1 byte")
        return uint64(value), nil
    }
    return 0, nil
}

func test_int(i uint64) {
    //fmt.Println("testing:", i)
	var buf [10]byte
	n := PutVarUintSimple(buf[:], i)
    reader := bufio.NewReader(bytes.NewReader(buf[0:n]))
    value, err := ReadVarUintSimple(reader)
    if err != nil {
        fmt.Println("ERROR", i)
        panic("error with read")
    }
    if value != i {
        fmt.Println("ERROR MISMATCH", i, value)
        panic("error mismatch")
    }
    //fmt.Println("Pass:", i)
}

func main() {
    var i uint64 = 0
    for {
        test_int(i)
        i++
        if i % 100000 == 0 {
            fmt.Println(i, "...")
        }
    }
}

