//package encryption works on encryption and decryption operation for the key provided
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//It will return a cipher stream
func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	h := md5.New()
	h.Write([]byte(key))
	cipherkey := h.Sum(nil)

	//Use the hash key as cipherkey
	block, err := aes.NewCipher(cipherkey)
	if err != nil {
		return nil, err
	}

	//Add in a stream
	return cipher.NewCFBEncrypter(block, iv), nil
}

//Encrypt function will take a key and a text value which we need to encrypt and return encrypted value
func Encrypt(key, textValue string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	//Create a Hasher

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(textValue))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}

	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(textValue))

	//Return encrypted element
	return fmt.Sprintf("%x", ciphertext), nil

}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("Unable to write iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

//Returns a stream from key and iv
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	h := md5.New()
	h.Write([]byte(key))
	cipherkey := h.Sum(nil)

	//Use the hash key as cipherkey
	block, err := aes.NewCipher(cipherkey)
	if err != nil {
		return nil, err
	}

	return cipher.NewCFBDecrypter(block, iv), nil
}

//Decrypt function will take a key and a hashencrypted cipher text and return decrypted value
func Decrypt(key, encryptedValue string) (string, error) {

	ciphertext, err := hex.DecodeString(encryptedValue)
	if err != nil {
		fmt.Println("Error")
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("Cipher value too short.")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	ciphertext_string := fmt.Sprintf("%s", ciphertext)

	return ciphertext_string, nil
}

//DecryptReader reads the iv and returns cipher's stream reader
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("Unable to read the iv")
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}
