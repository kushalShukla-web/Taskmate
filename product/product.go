package product

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type User struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Check       bool   `json:"is_check"`
}
type Project struct {
	Id int `json:"id"`
}

func Getfunc() ([]*User, error) {
	rows, err := Db.Query("SELECT id, task, date, description, is_check FROM todo")
	if err != nil {
		log.Println("Query failed:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*User
	for rows.Next() {
		var p User
		if err := rows.Scan(&p.Id, &p.Task, &p.Date, &p.Description, &p.Check); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	log.Println("Data fetched successfully")
	return products, nil
}

func Addfunc(x io.ReadCloser) error {
	data, err := io.ReadAll(x)
	if err != nil {
		return fmt.Errorf("Error %v", err)

	}
	var products []*User
	err = json.Unmarshal(data, &products)
	if err != nil {
		return fmt.Errorf("Error %v", err)
	}

	query := "INSERT INTO todo (id, task , date, description,is_check) VALUES ($1, $2, $3, $4, $5)"
	for _, val := range products {
		_, err := Db.Exec(query, val.Id, val.Task, val.Date, val.Description, val.Check)
		if err != nil {
			fmt.Printf("ERROR inserting data: %v\n", err) // Print the error
			return fmt.Errorf("ERROR %v", err)
		}
	}
	fmt.Printf("HEllo")
	return nil
}
func Deletefunc(x io.ReadCloser) {
	data, err := io.ReadAll(x)
	if err != nil {
		fmt.Errorf("ERROR %v", err)
	}
	var products []*Project
	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Printf("ERROR %v", err)
		return
	}
	query := "DELETE FROM todo WHERE id = $1"
	for _, val := range products {
		_, err := Db.Exec(query, val.Id)
		if err != nil {
			fmt.Printf("Error %v", err)
			return
		}
	}
}

func Replacefunc(x io.ReadCloser) {
	data, err := io.ReadAll(x)
	if err != nil {
		fmt.Errorf("ERROR %v", err)
	}
	var products []*User
	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Printf("ERROR%V", err)
	}
	query := "UPDATE todo SET task=$1 , date=$2, description=$3 , is_check=$4 WHERE id=$5"
	for _, val := range products {
		_, err := Db.Exec(query, val.Task, val.Date, val.Description, val.Check, val.Id)
		if err != nil {
			fmt.Printf("Error %v", err)
			return
		}
	}
}
