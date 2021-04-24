package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
)

func handle(chunk []byte) {

}

func main() {
	fileName := "test.log"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("cannot able to read the file", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	bytespool := sync.Pool{New: func() interface{} {
		bytess := make([]byte, 500*1024)
		return bytess
	}}
	for {
		buf := bytespool.Get().([]byte)
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}
	nextline,err:=r.ReadBytes('\n')
	if err!=io.EOF{
		buf=append(buf,nextline)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		wg.Done()
	}()
	wg.Wait()
}
