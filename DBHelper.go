package IGFXICommon

import (
"fmt"
"database/sql"
_ "github.com/denisenkom/go-mssqldb"
)

type IDbHelper interface {
	ExecuteQuery(query string, connectionString string) (*sql.Rows, error)
}

type DbHelper struct {}

func getDB(connectionString string) (*sql.DB, error){
	//https://www.educative.io/edpresso/how-to-create-and-utilize-packages-in-go
	db, err := sql.Open("mssql", connectionString)

	if err != nil {
		fmt.Println(err.Error())
		return db, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return db, err
	}
	return db, nil
}

func (d DbHelper) ExecuteQuery(query string, connectionString string) (*sql.Rows, error){
	db, err := getDB(connectionString)

	defer db.Close()

	var rows *sql.Rows
	rows, err = db.Query(query)

	return rows, err
}