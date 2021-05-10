package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type DbScript struct {
	File        string
	Script      string
	Description string
}

type DbVersion struct {
	Version int `db:"version"`
}

var schema_version = `
CREATE TABLE IF NOT EXISTS db_version (
	version int NOT NULL
);

INSERT OR IGNORE INTO db_version (version) VALUES (0) ;
`

func SqlLiteConnect(dataSourceName string) (*sqlx.DB, error) {
	bbddPath := "bbdd/"

	if _, exists := os.LookupEnv("BBDD_PATH"); exists {
		bbddPath = os.Getenv("BBDD_PATH")
	}

	bbddPath = bbddPath + dataSourceName
	db, err := sqlx.Connect("sqlite3", bbddPath)

	if err != nil {
		fmt.Errorf(filepath.Abs(bbddPath))
	}

	return db, err
}


func CheckDatabase(db *sqlx.DB, scripts []DbScript) int {

	versionDb, err := GetBdVersion(db)
	if err != nil {
		fmt.Print("Version table not found. Created a new one: ", err)
		db.MustExec(schema_version)
	} else {
		totalVersions := len(scripts)

		for currentVersion := versionDb; currentVersion < totalVersions; {
			currentVersion++
			fmt.Print("Executing script - ", currentVersion)

			actualScript := scripts[currentVersion-1]

			if len(actualScript.Script) > 0 {
				err = executeScript(db, actualScript.Script)
			} else if len(actualScript.File) > 0 {
				err = ExecuteFileScript(db, actualScript.File)
			}

			if err == nil {
				// Update the version in the table
				err = updateVersion(db, currentVersion)
			}

			if err != nil {
				fmt.Errorf(err.Error())
				break
			}
			fmt.Print("Executed succesfully - ", currentVersion)
		}
	}
	return 0
}

func GetBdVersion(db *sqlx.DB) (int, error) {
	var versions []DbVersion

	err := db.Select(&versions, "SELECT * FROM db_version")

	if err != nil {
		fmt.Errorf(err.Error())

	} else {
		return versions[0].Version, nil
	}

	return initDatabase(db)
}

func executeScript(db *sqlx.DB, sql string) error {
	_, err := db.Exec(sql)

	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	return nil
}

func ExecuteFileScript(db *sqlx.DB, filePath string) error {

	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	sql := string(bytes)
	_, err = db.Exec(sql)

	if err != nil {
		fmt.Errorf(err.Error())
		return err

	}

	return nil
}

func initDatabase(db *sqlx.DB) (int, error) {
	_, err := db.Exec(schema_version)

	if err != nil {
		fmt.Errorf(err.Error())
		return -1, errors.New("Error initializing versioned database.")

	} else {
		return 0, nil
	}
}

func updateVersion(db *sqlx.DB, version int) error {
	return executeScript(db, "UPDATE db_version SET version = "+strconv.Itoa(version)+" where 1=1;")
}
