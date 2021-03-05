package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func createConnection() *sql.DB  {
	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1 :3306)/golearnrest")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Student struct {
	Id           string `form:"id" json:"id"`
	FirstName    string `form:"firstname" json:"firstname"`
	LastName     string `form:"lastname" json:"lastname"`
	Identifier   string `form:"identifier" json:"identifier"`
	Email  		 string `form:"email" json:"email"`
	PhoneNumber  string `form:"phonenumber" json:"phonenumber"`
}

type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    []Student
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	var students Student
	var arr_student []Student
	var response Response
	
	db := createConnection()
	defer db.Close()
	
	rows, err := db.Query("SELECT id, first_name, last_name, email, identifier, phone_number FROM students")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(
			&students.Id,
			&students.FirstName,
			&students.LastName,
			&students.Email,
			&students.Identifier,
			&students.PhoneNumber); err != nil {
			log.Fatal(err.Error())
		}else {
			arr_student = append(arr_student, students)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_student

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/students", returnAllProducts).Methods("GET")
	http.Handle("/", router)
	fmt.Println("connected to port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}