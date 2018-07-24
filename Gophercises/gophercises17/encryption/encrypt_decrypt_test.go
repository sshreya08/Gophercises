package encryption

import (
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	ciphertext, err := Encrypt("test-key", "test-value")
	if ciphertext == "" {
		t.Errorf("Error: %d.", err)
	}
}

func TestEncryptNegative(t *testing.T) {
	ciphertext, err := Encrypt("test-key", "test-value")
	if ciphertext == "" {
		t.Errorf("Error: %d.", err)
	}
}

func TestDecrypt(t *testing.T) {
	ciphertext, err := Encrypt("test-key", "test-encrypted-value")
	if ciphertext == "" {
		t.Errorf("Error: %d.", err)
	}
}

func TestDecryptNegative(t *testing.T) {
	ciphertext, err := Encrypt("test-key1", "test-encrypted-value")
	if err != nil {
		t.Errorf("Error for %s: %d.", ciphertext, err)
	}
}

func TestEncryptWriter(t *testing.T) {
	testFile, err := os.OpenFile("test", os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := EncryptWriter("test-key", testFile)
	if err != nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
}

func TestEncryptWriterNegative(t *testing.T) {
	testFile, err := os.OpenFile("test", os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := EncryptWriter("test-key", testFile)
	if cipherstream == nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
}

func TestDecryptReader(t *testing.T) {
	testFile, err := os.OpenFile("test", os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := DecryptReader("test-key", testFile)
	if err != nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
}

func TestDecryptReaderNegative(t *testing.T) {
	testFile, err := os.OpenFile("test", os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := DecryptReader("test-key", testFile)
	if cipherstream == nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
}
