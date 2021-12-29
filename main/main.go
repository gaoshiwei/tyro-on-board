package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var Db *sqlx.DB

func init() {
	// database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	// defer Db.Close()
}

func main() {
	http.HandleFunc("/", IndexHandler)

	http.HandleFunc("/insert/person", Insert)
	http.HandleFunc("/select/person", SelectPerson)
	http.HandleFunc("/update/person", UpdatePerson)
	http.HandleFunc("/delete/person", DeletePerson)
	http.HandleFunc("/transaction/person", TransactionPerson)
	http.ListenAndServe("127.0.0.1:9000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func Insert(w http.ResponseWriter, r *http.Request) {
	InsertPerson()
}

func SelectPerson(w http.ResponseWriter, r *http.Request) {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id = ?", 3)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select success:", person)
}

func InsertPerson() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "test1", "man", "test1@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert success:", id)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	res, err := Db.Exec("update person set username=? where user_id=?", "test2", 3)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update success:", row)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	res, err := Db.Exec("delete from person where user_id=?", 3)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("delete success: ", row)
}

func TransactionPerson(w http.ResponseWriter, r *http.Request) {
	TestTransaction()
}

// TestTransaction
// Db.Begin()        	开始事务
// Db.Commit()        	提交事务
// Db.Rollback()        回滚事务
func TestTransaction() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "test3", "man", "test3@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "test4", "man", "test4@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success:", id)
	conn.Commit()
}
