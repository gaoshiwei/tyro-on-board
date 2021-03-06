package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	// database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	// 使用docker-compose起的时候需要使用下面的这行去访问mysql
	database, err := sqlx.Open("mysql", "root:123456@tcp(mysql:3306)/test")
	if err != nil {
		log.Println("mysql connect fail", err)
		return
	}
	Db = database
	log.Println("mysql connect successful")
	// defer Db.Close()
}

func main() {
	http.HandleFunc("/", IndexHandler)

	http.HandleFunc("/insert/person", Insert)
	http.HandleFunc("/select/person", SelectPerson)
	http.HandleFunc("/update/person", UpdatePerson)
	http.HandleFunc("/delete/person", DeletePerson)
	http.HandleFunc("/transaction/person", TransactionPerson)
	log.Println("服务启动前")
	http.ListenAndServe("0.0.0.0:9000", nil)
	log.Println("服务启动成功")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var person Person
	json.Unmarshal(body, &person)
	fmt.Println(person)
	// TODO Param validate
	InsertPerson(&person)
}

func SelectPerson(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userId, err := strconv.Atoi(query.Get("user_id"))
	if err != nil {
		fmt.Println("url param wrong, ", err)
	}
	var person []Person
	err = Db.Select(&person, "select user_id, username, sex, email from person where user_id = ?", userId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select success:", person)
}

func InsertPerson(person *Person) {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", person.UserName, person.Sex, person.Email)
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
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}
	var person Person
	json.Unmarshal(body, &person)
	fmt.Println(person)
	res, err := Db.Exec("update person set username=? where user_id=?", person.UserName, person.UserId)
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
