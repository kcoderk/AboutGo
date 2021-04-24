package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	size := 1024 * 1024  * 2
	f, err := os.Create("test.log")
	if err != nil {
		log.Fatal(err)
	}
	//testarr := []string{"test"}
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString("test")
		if i%1000==0 {
			f.WriteString(buffer.String())
		}
	}

}
