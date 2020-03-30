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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU1NjEwMTQsImlhdCI6MTU4NTU2MDExNCwiaXNzIjoic2Vuc2xhYnMuaW8iLCJzdWIiOiJ7XCJBXCI6XCJCXCIsXCJDXCI6XCJEXCJ9In0.tQc1IlfpZTw5JXPAW-JVLy7k0RYz8YGc28m-BV1c7_c"
	fmt.Println(VerifyTokenString(token))
}
