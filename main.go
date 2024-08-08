package main

import (
	"fmt"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Print(err)
	}

	// n, err := s.Write([]byte("test"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {
		buf := make([]byte, 2)
		_, err = s.Read(buf)
		if err != nil {
			fmt.Print(err)
		}
		// log.Printf("%v", len(buf))
		fmt.Printf("%x \n", buf[:len(buf)-1])
	}
}
