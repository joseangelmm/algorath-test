package repository

import (
	"algorath/algorath"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DatabaseI interface {
	GetCredential() (algorath.Credentials, error)
	UpdateCredential(algorath.Credentials) error
}

var scriptCredential = `CREATE TABLE IF NOT EXISTS credentials (
	ID int PRIMARY KEY,	
	APIKEY text ,
	APISECRET text
);
INSERT INTO credentials (ID,APIKEY,APISECRET) VALUES (0,"","");
`

type Database struct {
	conn   *sqlx.DB
}

func New() *Database {

	db := &Database{}
	var err error

	bbddName := "database.db"

	db.conn, err = sqlx.Connect("sqlite3",bbddName)

	if _, err := os.Stat(bbddName); err == nil {
		fmt.Printf("Database already exists\n");
	} else {
		fmt.Printf("Database does not exist so it will be created\n");
		db.conn.MustExec(scriptCredential)

	}

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func (db *Database) GetCredential() (algorath.Credentials, error) {
	var credential []algorath.Credentials
	err := db.conn.Select(&credential, "SELECT APIKEY, APISECRET FROM credentials WHERE ID=0")
	if err != nil {
		fmt.Printf("GetCredential() - Error getting credential from database: %s", err.Error())
	} else {
		return credential[0], nil
	}
	return algorath.Credentials{}, err
}


func (db *Database) UpdateCredential(credentials algorath.Credentials) (error) {
	var nc int64

	n, err := db.conn.NamedExec("UPDATE credentials SET APIKEY=:APIKEY, APISECRET=:APISECRET WHERE ID = 0", credentials)

	if err != nil {
		fmt.Printf("UpdateCredential() - Error updating credentials: %s", err)
	} else {
		nc, _ = n.RowsAffected()
		if nc > 0 {
			fmt.Printf("UpdateCredential() - Updated %d rows. credential", nc)
		} else {
			fmt.Printf("UpdateCredential() - Updated %d rows", nc)
		}
		return nil
	}
	return err
}
