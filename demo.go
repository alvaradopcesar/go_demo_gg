package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o demo.linux demo.go

type User struct {
	Id        int64
	Nombre    string
	Email     string
	Direccion string
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("error !!")
		fmt.Println(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(User{}, "User").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	return dbmap
}

func GetUsers(c *gin.Context) {
	var users []User
	_, err := dbmap.Select(&users, "Select id,nombre,email , direccion from user")
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": "no user(s) into the table"})
	} else {
		c.JSON(200, users)
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
	}

	r.Run(":8080")

}
