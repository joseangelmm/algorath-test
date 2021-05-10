package repository

import (
	"algorath/algorath"
	"fmt"
	"github.com/jmoiron/sqlx"
)


type DatabaseI interface {
	GetCredential() (algorath.Credentials, error)
	UpdateCredential(algorath.Credentials) error
}

type Database struct {
	conn   *sqlx.DB
}

func New() *Database {

	db := &Database{}
	bbddName := "database.db"

	var err error
	db.conn, err = SqlLiteConnect(bbddName)

	if err != nil {
		fmt.Errorf("New() - Error connecting database: ", err)
	} else {
		fmt.Print("New() - Database connected")
		status := CheckDatabase(db.conn, scripts)
		fmt.Print("New() - Database executed schema: ", status)
		version, _ := GetBdVersion(db.conn)
		fmt.Print("New() - Database version: ", version)
	}

	return db
}

func (db *Database) GetCredential() (algorath.Credentials, error) {
	var credential []algorath.Credentials
	err := db.conn.Select(&credential, "SELECT * FROM credentials WHERE ID=0")
	if err != nil {
		fmt.Errorf("GetCredential() - Error getting credential from database:", err)
	} else {
		return credential[0], nil
	}
	return algorath.Credentials{}, err
}


func (db *Database) UpdateCredential(credentials algorath.Credentials) (error) {
	var nc int64

	n, err := db.conn.NamedExec("UPDATE credentials SET APIKEY=:APIKEY, APISECRET=:APISECRET WHERE ID = 0", credentials)

	if err != nil {
		fmt.Errorf("UpdateCredential() - Error updating credentials:", err)
	} else {
		nc, _ = n.RowsAffected()
		if nc > 0 {
			fmt.Printf("UpdateCredential() - Updated", nc, "rows. credential")
		} else {
			fmt.Printf("UpdateCredential() - Updated", nc, "rows")
		}
		return nil
	}
	return err
}

func (db *Database) DeleteConn(){
	db.conn = nil
}