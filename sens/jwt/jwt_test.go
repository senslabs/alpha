package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	m := map[string]string{
		"OrgId": "139284a8-5a2a-4982-a92f-e58a94dde1e1",
	}
	token, _ := GenerateAccessToken(m, 365*24*time.Hour)

	fmt.Println(token)

	// fmt.Println(exec.Command("echo", "export", "TOKEN="+token, ">>", "~/.bash_profile").Run())
}

func TestVerifyAccessToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU1NjEwMTQsImlhdCI6MTU4NTU2MDExNCwiaXNzIjoic2Vuc2xhYnMuaW8iLCJzdWIiOiJ7XCJBXCI6XCJCXCIsXCJDXCI6XCJEXCJ9In0.tQc1IlfpZTw5JXPAW-JVLy7k0RYz8YGc28m-BV1c7_c"
	fmt.Println(VerifyTokenString(token))
}
