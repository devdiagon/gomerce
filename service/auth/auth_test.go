package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == password {
		t.Errorf("expected hash to be different\n original: %s \n got: %s", password, hash)
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte(password)) {
		t.Errorf("expected password to match hash")
	}

	if ComparePasswords(hash, []byte("notmatchpassword")) {
		t.Errorf("expected password to not match hash")
	}
}
