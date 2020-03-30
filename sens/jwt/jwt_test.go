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

func TestVerifyAccessToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU1NjA0NjcsImlhdCI6MTU4NTU1OTU2NywiaXNzIjoic2Vuc2xhYnMuaW8iLCJzdWIiOiJ7XCJBXCI6XCJCXCIsXCJDXCI6XCJEXCJ9In0.jAhDZRvcxB3oN6jFlDM8kmOcRuW7NUQN78UnOyl4Vk1PLrDIyNIEUCgy0CDBAOT9pUkErK4WstKzEK6EUuT9LQ"
	fmt.Println(VerifyTokenString(token))
}
