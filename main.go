package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("tesseract", "tax.jpg", "output.txt")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running Tesseract:", err)
		return
	}

	// Read the extracted text from "output.txt" file here
}
