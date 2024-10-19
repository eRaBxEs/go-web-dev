package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Ehioje Henry Erabor"

	str := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset=UTF-8>
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	`)

	// create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating a file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
