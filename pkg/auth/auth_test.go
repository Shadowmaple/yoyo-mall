package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {
	s := "admin"
	result, err := Encrypt(s)
	if err != nil {
		t.Error("encrypt error:", err)
	}
	// fmt.Println("encrypt result:", result)
	err = Compare(result, s)
	if err != nil {
		t.Error("compare error:", err)
	}
	t.Log("test ok")
}
