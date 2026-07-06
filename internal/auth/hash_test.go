package auth

import "testing"

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	password := "secret123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}
	if hash == "" {
		t.Fatal("expected non-empty hash")
	}

	ok, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected valid password to verify")
	}
}

func TestCheckPasswordHash_InvalidPassword(t *testing.T) {
	password := "secret123"
	wrongPassword := "badpass"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	ok, err := CheckPasswordHash(wrongPassword, hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash returned error for wrong password: %v", err)
	}
	if ok {
		t.Fatal("expected invalid password to fail verification")
	}
}
