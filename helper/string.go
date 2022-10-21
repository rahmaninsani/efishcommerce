package helper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"math/rand"
	"strings"
	"time"
)

func RandomString(stringLength int, alphabet []rune) string {
	rand.Seed(time.Now().UnixNano())

	alphabetSize := len(alphabet)
	var stringBuilder strings.Builder

	for i := 0; i < stringLength; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		stringBuilder.WriteRune(ch)
	}

	return stringBuilder.String()
}

func GenerateOrderCode() string {
	alphabet := []rune("0123456789")
	stringLength := 10
	randomString := RandomString(stringLength, alphabet)

	currentTime := time.Now()
	orderDate := currentTime.Format("20060102")

	orderCode := fmt.Sprintf("ORD-%s-%s", orderDate, randomString)
	return orderCode
}

func GenerateFileName(fileExtension string) string {
	return fmt.Sprintf("%s.%s", uuid.New(), fileExtension)
}

func GenerateSlug(text string) string {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	stringLength := 12
	randomString := RandomString(stringLength, alphabet)

	text = slug.MakeLang(text, "id")
	return fmt.Sprintf("%s-%s", text, randomString)
}
