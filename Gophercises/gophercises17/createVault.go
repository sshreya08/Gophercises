//package gophersises17 works on vault operation on key value file and contains other packages for command line and encryption and decryption
package gophersises17

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/sonam/Gophercises/gophercises17/encryption"
)

//Pass encoding key and value to the vault
func NewVault(encodeKey string) Vault {
	return Vault{
		encodeKey:    encodeKey,
		encodeValues: make(map[string]string),
	}
}

//Pass encoding key and value to the vault
func File(encodeKey, file string) FileVault {
	return FileVault{
		encodeKey:    encodeKey,
		filePath:     file,
		encodeValues: make(map[string]string),
	}
}

//Create a Struct of Vault
type Vault struct {
	encodeKey    string
	encodeValues map[string]string
}

//Create a Struct of Vault
type FileVault struct {
	encodeKey    string
	filePath     string
	mutex        sync.Mutex
	encodeValues map[string]string
}

//Load the file if exists else return error
func (v *FileVault) load() error {
	file, err := os.Open(v.filePath)
	if err != nil {
		v.encodeValues = make(map[string]string)
		return nil
	}
	defer file.Close()
	r, err := encryption.DecryptReader(v.encodeKey, file)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

//Decode the reader and return decoded values
func (v *FileVault) readKeyValues(reader io.Reader) error {
	decoded := json.NewDecoder(reader) //Decode from the decryptReader
	return decoded.Decode(&v.encodeValues)
}

//Encrypt in a JSON file
func (v *FileVault) save() error {

	file, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	w, err := encryption.EncryptWriter(v.encodeKey, file)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

//Encode the writer and return encoded values
func (v *FileVault) writeKeyValues(writer io.Writer) error {
	encoded := json.NewEncoder(writer) //Encode it to json and write to writer
	return encoded.Encode(v.encodeValues)
}

//Get function returns the Vault value from the key provided
func (v *FileVault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.encodeValues[key]
	if !ok {
		return "", errors.New("No value provided for the key")
	}
	return value, nil
}

//Set function sets the Vault value from the key and value provided
func (v *FileVault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.encodeValues[key] = value
	err = v.save()
	return nil
}
