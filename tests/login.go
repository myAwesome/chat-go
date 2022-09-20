package main

import (
	"fmt"
)

func main() {

	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Purple = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"
	var White = "\033[97m"

	fmt.Println(fmt.Sprintf("\033[97m%s\u001B[0m", "This is White"))
	fmt.Println(White + "This is White" + Reset)

	fmt.Println("\033[31m" + "This is Red" + Reset)
	fmt.Println(Red + "This is Red" + Reset)
	fmt.Println(Green + "This is Green" + Reset)
	fmt.Println(Yellow + "This is Yellow" + Reset)
	fmt.Println(Blue + "This is Blue" + Reset)
	fmt.Println(Purple + "This is Purple" + Reset)
	fmt.Println(Cyan + "This is Cyan" + Reset)
	fmt.Println(Gray + "This is Gray" + Reset)

}
