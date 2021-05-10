package repository

import (
	"algorath/algorath"
	"testing"
)

var (
	db = New()
)

func TestGetCredentialOK(t *testing.T){

	_, err := db.GetCredential()

	if err != nil {
		t.Errorf("TestGetCredentialOK returned unexpected error: got %v want %v", err, nil)
		return
	}

}

func TestUpdateCredentialOK(t *testing.T){

	err := db.UpdateCredential(algorath.Credentials{
		APIKey:    "testing",
		APISecret: "testing2",
	})

	if err != nil {
		t.Errorf("TestUpdateCredentialOK returned unexpected error: got %v want %v", err, nil)
		return
	}

}

func TestGetCredentialKO(t *testing.T){

	db.DeleteConn()

	_, err := db.GetCredential()

	if err != nil {
		t.Errorf("TestGetCredentialKO returned unexpected error: got %v want %v", err, nil)
		return
	}

}

func TestUpdateCredentialKO(t *testing.T){

	err := db.UpdateCredential(algorath.Credentials{
		APIKey:    "testing",
		APISecret: "testing2",
	})

	if err != nil {
		t.Errorf("TestUpdateCredentialKO returned unexpected error: got %v want %v", err, nil)
		return
	}

}