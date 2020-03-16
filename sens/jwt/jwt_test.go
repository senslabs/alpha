package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	m := map[string]string{
		"A": "B",
		"C": "D",
	}

	fmt.Println(GenerateTemporaryToken(m))
}
