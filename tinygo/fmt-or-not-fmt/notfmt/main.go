package main

import "os"

func main() {
	os.Stdout.WriteString("Content-Type: text/plain\n\nTesting\n")
}
