package eth

import (
	"fmt"
	"testing"
)

const (
	ACCOUNT1_ADDRESS    = "0xE7bc6d2F28B68626106391332fEdFD31A3725bBb"
	ACCOUNT1_PRIVATEKEY = "47C629B5130B6E8BBDE4BB0B72898A5464500C9ADB482F05D4173F599239A426"
	ACCOUNT2_ADDRESS    = "0x0698c06FC0c46f57CA561E561b03E2b42522455f"
	ACCOUNT2_PRIVATEKEY = "C6AA02C1CE2D091E39A963C5E8E88F9A70AE13F3FA569ECDAD5D9A79F906ED81"
	ACCOUNT3_ADDRESS    = "0x2C8ce8efc5d3cf7A4A3833764fCc307bA98a3067"
	ACCOUNT3_PRIVATEKEY = "678CEB514C039BBAE87F28EC6CF5966CF303F123253436314D55381483C6502E"
)

func TestEtherTransaction(t *testing.T) {
	privateKeyHex := ACCOUNT3_PRIVATEKEY
	to := ACCOUNT2_ADDRESS
	var ether float64 = 0.01
	tx, err := EtherTransaction(privateKeyHex, to, ether)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
}

func TestTokenTransaction(t *testing.T) {
	privateKeyHex := ACCOUNT1_PRIVATEKEY
	to := ACCOUNT2_ADDRESS
	var token float64 = 1.5
	tx, err := TokenTransaction(privateKeyHex, to, token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
}
