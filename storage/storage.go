package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/climanager/bins"
)

func SaveBinToJson(bin bins.Bin) error {
	byteArr, err := json.Marshal(bin)
	if err != nil {
		return err
	}

	file, err := os.Create("bin.json")
	defer file.Close()

	_, err = file.WriteString(string(byteArr))
	if err != nil {
		return err
	}
	fmt.Println("Запись успешна.")
	return nil
}

func ReadJson() ([]bins.Bin, error) {
	jsonFile, err := os.ReadFile("bin.json")
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %w", err)
	}

	var emptyBin []bins.Bin
	err = json.Unmarshal(jsonFile, &emptyBin)
	if err != nil {
		return nil, fmt.Errorf("Ошибка преобразования JSON: %w", err)
	}

	return emptyBin, nil

}
