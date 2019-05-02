package encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	raw := []byte("password123")
	key := []byte("asdf1234qwer7894")
	str, err := Encrypt(raw, key)
	if err == nil {
		t.Log("suc", str)
	} else {
		t.Fatal("fail", err)
	}
}

func TestDncrypt(t *testing.T) {
	raw := "pqjPM0GJUjlgryzMaslqBAzIknumcdgey1MN+ylWHqY="
	key := []byte("asdf1234qwer7894")
	str, err := Dncrypt(raw, key)
	if err == nil {
		t.Log("suc", str)
	} else {
		t.Fatal("fail", err)
	}
}
