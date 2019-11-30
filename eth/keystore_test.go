package eth

import "testing"

import "fmt"

func TestNewKeystore(t *testing.T) {
	account, err := NewKeystore("12345678")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}
