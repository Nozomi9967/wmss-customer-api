package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID32() string {
	rawId := uuid.New().String()
	Id := strings.ReplaceAll(rawId, "-", "")
	return Id
}
