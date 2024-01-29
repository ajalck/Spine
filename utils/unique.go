package utils

import (
	"github.com/google/uuid"
)

func GenerateUniqueString() string {
	defer func() string {
		 err := recover()
		 return err.(string)
	}()
	return uuid.NewString()
}
