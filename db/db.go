package db

import (
	"database/sql"
	"echo-template/conf"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error //global variable karena banyak diulang

func Init(){
	config := conf.GetConfig()

	//DSN
	connectionString := config.DB_USERNAME + ":" + config.DB_PASSWORD+ "@tcp("+ config.DB_HOST+ ":" + config.DB_PORT + ")/" + config.DB_NAME

	//INITITATE CONNECTION TO SQL WITH DSN
	db, err = sql.Open("mysql", connectionString)

	if err != nil{
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil{
		panic("DSN Invalid")
	}
}



func CreateCon()*sql.DB{

	return db

}