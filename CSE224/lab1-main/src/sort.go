package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Record struct {
	length []byte
	key    []byte
	value  []byte
}

type Cmp []Record

func (a Cmp) Len() int      { return len(a) }
func (a Cmp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Cmp) Less(i, j int) bool {
	return bytes.Compare(a[i].key, a[j].key) < 0
}

func ReadBigEndianUint32(buffer []byte) uint32 {
	if len(buffer) < 4 {
		panic("buffer too short to read uint32")
	}
	return binary.BigEndian.Uint32(buffer[:])
}

// Write a big-endian uint32 to a byte slice of length at least 4
func WriteBigEndianUint32(buffer []byte, num uint32) {
	if len(buffer) < 4 {
		panic("buffer too short to write uint32")
	}
	binary.BigEndian.PutUint32(buffer, num)
}

func fileAccess() {
	var data [4]byte = [4]byte{0x00, 0x00, 0x00, 0x01}
	num := ReadBigEndianUint32(data[:])
	fmt.Println(num) // Output: 1

	// Writing a big-endian uint32 to a byte slice
	var buffer [4]byte
	WriteBigEndianUint32(buffer[:], num)
	fmt.Println(buffer) // Output: [0 0 0 1]

	// Attempting to write a big-endian uint32 to a 2-byte buffer
	var shortBuffer [2]byte
	WriteBigEndianUint32(shortBuffer[:], num) // This will cause a panic
	fmt.Println(shortBuffer)                  // This line will not be reached due to the error
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening the file.")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var records []Record
	for {
		var record Record
		length := make([]byte, 4)
		err = binary.Read(reader, binary.BigEndian, &length)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Failed to read 4 bytes for length :", err)
		}
		n := ReadBigEndianUint32(length)
		record.length = length
		fmt.Println("Read length :", length)

		key := make([]byte, 10)
		_, err := io.ReadFull(reader, key)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Failed to read key :", err)
		}
		record.key = key

		value := make([]byte, n-10)
		_, err = io.ReadFull(reader, value)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Failed to read value :", err)
		}
		record.value = value

		records = append(records, record)
	}

	output, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Failed creating output file :", output)
	}
	defer output.Close()

	sort.Sort(Cmp(records))
	for i := 0; i < len(records); i++ {
		record := records[i]
		buf := make([]byte, 0, ReadBigEndianUint32(record.length)+4)
		buf = append(buf, record.length...)
		buf = append(buf, record.key...)
		buf = append(buf, record.value...)
		output.Write(buf)
	}
}
