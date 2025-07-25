package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() error {
	fileData, err := os.ReadFile("byte.json")
	if err != nil {
		return err
	}
	fmt.Println("Файл:")
	fmt.Println(string(fileData))
}

func checkJSON(str string) bool {
	return strings.Contains(str, ".json")
}
