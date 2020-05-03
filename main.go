package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func display(items []string) {
	for i, item := range items {
		fmt.Printf("[%d]\t %s\n", i, item)
	}
}

func readFromFile(filename string) []string {
	return nil
}

func writeToFile(filename string, items []string) {

}

func removeItem(i int, items []string) []string {
	return nil
}

func main() {
	cliFlags := flag.String("add", "done")
	items := readFromFile("/tmp/yata4f.txt")

	args := os.Args[1:]

	if len(args) > 0{
		if strings.Contains(args.[0], "-a"){
			
		}
	}
	display(items)
}
