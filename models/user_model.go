package models

import (
	"echo-template/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=100"`
	City string `json:"city" validate:"required,min=2,max=100"`
}

type UserPostResponse struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type UpdateResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

//VALIDATION
var validate = validator.New()

func ValidateUser(user User) error {
	return validate.Struct(user)
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


func PostUser(name string, city string)(Response, error){
	var res Response

	user := User{Name: name, City: city}

	if err := validate.Struct(user); err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Validation failed"
		res.Data = err.Error()
		return res, nil
	}

	con := db.CreateCon()


	//CHECK IF NAME ALREADY EXIST OR NO
	var exist bool
	checkId :="SELECT EXISTS(SELECT 1 FROM user WHERE name = ?)"

	err := con.QueryRow(checkId, name).Scan(&exist)
	if err != nil {
		return res, err
	}

	if exist {
		res.Status = http.StatusConflict
		res.Message = "User already exists"
		return res, nil
	}
	

	sqlStatement := "INSERT user (name, city) VALUES (?,?)"

	stmt, err := con.Prepare(sqlStatement)
	
	if err != nil{
		return res, err
	}

	result, err := stmt.Exec(name, city)
	if err != nil{
		return res, err
	}

	lastInsertedId, err:= result.LastInsertId()
	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Add User Data"
	res.Data =UserPostResponse{
		Id: lastInsertedId,
		Name: name,
		City:city,
	}
	return res, nil

	}


	func UpdateUser(id int, name string, city string)(Response, error){
		var res Response
		con := db.CreateCon()

		sqlStatement :="UPDATE user SET name =? ,  city=? WHERE id=?"

		stmt, err :=con.Prepare(sqlStatement)
		if err!=nil{
			return res,err
		}

		_, err=stmt.Exec(name,city,id)
		if err !=nil{
			return res,err
		}

		res.Status = http.StatusOK
		res.Message = "success Edit data"
		res.Data= UpdateResponse{
			Id: id,
			Name: name,
			City: city,
		}
		return res, nil
	}


	func DeleteUser(id int)(Response,error){
		var res Response
		con:= db.CreateCon()

		query :="DELETE FROM user WHERE id=?"
		

		stmt, err:= con.Prepare(query)
		if err!=nil{
			return res,err
		}

		result, err :=stmt.Exec(id)
		if err!=nil{
			return res,err
		}

		rowsAffected, err := result.RowsAffected()
		if err!=nil{
			return res, err
		}
		res.Status=http.StatusOK
		res.Message="Data deleted"
		res.Data=map[string]int64{
			"id":rowsAffected,
		}
		return res,nil
	}