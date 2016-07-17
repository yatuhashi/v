package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			v := i*16 + j
			str := "\033[38;5;" + strconv.Itoa(v) + "m" + strconv.Itoa(v) + "\033[0m "
			os.Stdout.Write([]byte(str))
		}
		fmt.Println("\n")
	}
}
