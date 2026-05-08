package main

import (
	service "demoGo/service/user"
	utils "demoGo/utils/aes"
	"encoding/json"
	"fmt"
)

func main() {
	str := utils.AESDecrypt("MFvciXC6H8SAjWcp4/c9iUHlIQvOwUtGnFUztX1TzpexfmOIwpBjjasaFV+lM/tV")
	fmt.Println(str)

	list := service.GetUserList()

	jsonBytes, err := json.MarshalIndent(list, "", "  ")

	if err != nil {
		panic(fmt.Sprintf("转换 JSON 失败: %v", err))
	}

	fmt.Println(string(jsonBytes))
}
