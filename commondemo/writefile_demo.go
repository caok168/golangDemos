package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func main() {

	data := Float32ToByte(13.4)

	aaa := []float32{1.222222, 3.24555553, 3.1111111}
	//aaa := []float32{1.222222, 3.2455556,  3.1111112

	f, err := os.Create("./commondemo/bbb.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range aaa {
		fmt.Fprint(f, a)
	}

	write("./commondemo/aaa.txt", data)
}

// write data append to file
func write(fileName string, data []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println("err:", err)

		return err
	}

	return nil
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}
