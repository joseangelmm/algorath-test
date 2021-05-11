package repository

import (
	"algorath/algorath"
	"testing"
)


func TestGetCredentialKO(t *testing.T){

	db := New()
	_, err := db.GetCredential()

	if err != nil {
		t.Errorf("TestGetCredentialOK returned error %s but we wanted nil", err.Error())
		return
	}

}

func TestUpdateCredentialOK(t *testing.T){

	db := New()
	err := db.UpdateCredential(algorath.Credentials{
		APIKey:    "testing",
		APISecret: "testing2",
	})

	if err != nil {
		t.Errorf("TestUpdateCredentialOK returned error %s but we wanted nil", err.Error())
		return
	}

}