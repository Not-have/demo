package main

import (
	utils "demoGo/utils/aes"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	str := utils.AESDecrypt("MFvciXC6H8SAjWcp4/c9iUHlIQvOwUtGnFUztX1TzpexfmOIwpBjjasaFV+lM/tV")
	fmt.Println(str)
}
