package util

import (
	"testing"
	"fmt"
)

func TestEncrypt(t *testing.T) {
	text := []byte("My name is Astaxie")
	key := []byte("the-key-has-to-be-32-bytes-long!")

	ciphertext, err := encrypt(text, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s => %x\n", text, ciphertext)

	plaintext, err := decrypt(ciphertext, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x => %s\n", ciphertext, plaintext)
}
