package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	postgresdriver = "postgres"
	host           = "localhost"
	port           = 5432
	user           = "postgres"
	password       = "postgres"
	dbname         = "go"
)

type Pessoa struct {
	id    int
	nome  string
	email string
}

var db *sql.DB
var err error

func main() {

	var conec = "host=localhost port=5432 user=postgres password=postgres dbname=go sslmode=disable"
	// var conec = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Printf("acessando %s ... ", dbname)

	db, err = sql.Open("postgres", conec)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conectou!!")
	}

	defer db.Close()

	//sqlSelect()
	//sqlSelectID()
	//sqlInsert()
	//sqlUpdate()
	//sqlDelete()

	sqlSelect()
}

func sqlSelect() {

	sqlStatement, err := db.Query("SELECT * FROM pessoa ")
	if err != nil {
		panic(err.Error())
	}

	for sqlStatement.Next() {

		var pessoa Pessoa

		err = sqlStatement.Scan(&pessoa.id, &pessoa.nome, &pessoa.email)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%d\t%s\t%s \n", pessoa.id, pessoa.nome, pessoa.email)
	}
}

func sqlSelectID() {

	var pessoa Pessoa

	sqlStatement := "SELECT * FROM pessoa where id = $1"

	err = db.QueryRow(sqlStatement, 1).Scan(&pessoa.id, &pessoa.nome, &pessoa.email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%d\t%s\t%s \n", pessoa.id, pessoa.nome, pessoa.email)
}

func sqlInsert() {

	sqlStatement := fmt.Sprintf("INSERT INTO pessoa (nome, email) VALUES ($1, $2)")

	insert, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}

	result, err := insert.Exec("teste", "testee@teste.com") //atualziar email para cada inserção
	if err != nil {
		panic(err.Error())
	}

	affect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(affect)
}

func sqlUpdate() {

	sqlStatement := fmt.Sprintf("update pessoa set nome=$1 where id=$2")

	update, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}

	result, err := update.Exec("teste update", 2)
	if err != nil {
		panic(err.Error())
	}

	affect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(affect)
}

func sqlDelete() {

	sqlStatement := fmt.Sprintf("delete from pessoa where id=$1")

	delete, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}

	result, err := delete.Exec(5)
	if err != nil {
		panic(err.Error())
	}

	affect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(affect)
}
