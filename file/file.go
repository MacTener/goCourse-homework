package file

import (
	"fmt"
	"os"
)

func ReadFile() error {
	fileData, err := os.ReadFile("byte.json")
	if err != nil {
		return err
	}
	fmt.Println(string(fileData))
}
