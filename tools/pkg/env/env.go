package env

import (
	"github.com/joho/godotenv"
)

func Read(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil {
		return err
	}
	return nil
}
