package main

import (
	"bufio"
	"fmt"
	"github.com/rodolfoag/gow32"
	"os"
)

func main() {
	m := os.Args[1]
	_, err := gow32.CreateMutex(m)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Mutex %s create. Press enter to quit.\n", m)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
