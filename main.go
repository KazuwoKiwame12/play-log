package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	output, _ := os.Open("1_toFile/output.log")

	fmt.Println("ログの中身を書き出します")
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}
