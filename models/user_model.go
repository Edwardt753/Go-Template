package models

import (
	"echo-template/db"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func FetchAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM user"

	rows, err := con.Query(sqlStatement)
	
	if err!= nil{
		return res, err
	}

	defer rows.Close()


	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.City)
		if err !=nil{
			return res,err
		}
		arrobj=append(arrobj, obj)
	}
	res.Status=http.StatusOK
	res.Message="Success get all data"
	res.Data=arrobj

	return res, nil
}