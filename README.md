# gow32
Windows APIs for Golang

CreateMutex: creates a mutex object on Windows. Usefull for ensuring that only one instance of the application is open. CreateMutex will return the error ERROR_ALREADY_EXISTS ( **int(err.(syscall.Errno))** ) if the mutex was already created by another program.

<pre><code>
package main

import (
	"bufio"
	"fmt"
	"github.com/rodolfoag/gow32"
	"os"
	"syscall"
)

func main() {
	m := os.Args[1]
	_, err := gow32.CreateMutex(m)
	if err != nil {
		fmt.Printf("Error: %d - %s\n", int(err.(syscall.Errno)), err.Error())
	} else {
		fmt.Printf("Mutex %s create. Press enter to quit.\n", m)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
</code></pre>

