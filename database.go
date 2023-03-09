//
//
////import (
////	"database/sql"
////	"fmt"
////	"time"
////
////	_ "github.com/go-sql-driver/mysql"
////)
//
//var db *sql.DB
//
//func initMySQL() (err error) {
//	dsn := "root:mysql@tcp(127.0.0.1:3306)/user"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	err = db.Ping()
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	db.SetConnMaxLifetime(time.Second * 10)
//	db.SetMaxOpenConns(200)
//	return
//}
//
//type user struct {
//	id   int
//	age  int
//	name string
//}
//
//func queryRowDemo() {
//	sqlStr := "select id, name, age from user where id =?"
//	var u user
//	row := db.QueryRow(sqlStr, 2)
//	err := row.Scan(&u.id, &u.name, &u.age)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	fmt.Printf("u.id, u.name, u.age =====> 🚀🚀🚀 id:%v\n,name: %v\n,age: %v\n", u.id, u.name, u.age)
//}
//
//func queryMultiRowDemo() {
//	sqlStr := "select id, name, age from user where id > ?"
//	rows, err := db.Query(sqlStr, 0)
//	if err != nil {
//		fmt.Printf("query failed  =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var u user
//		err := rows.Scan(&u.id, &u.name, &u.age)
//		if err != nil {
//			fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//			return
//		}
//		fmt.Printf("u.id, u.name, u.age =====> 🚀🚀🚀 %v\n %v\n %v\n ", u.id, u.name, u.age)
//	}
//}
//
//func insertRowDemo() {
//	sqlStr := "insert into user(name, age) values (?, ?)"
//	ret, err := db.Exec(sqlStr, "haha", 111)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	theID, err := ret.LastInsertId()
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	fmt.Printf("theID =====> 🚀🚀🚀 %v\n", theID)
//}
//
//func updateRowDemo() {
//	sqlStr := "update user set age=? where id=?"
//	ret, err := db.Exec(sqlStr, 391111, 3)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected()
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	fmt.Printf("110 =====> 🚀🚀🚀 %v\n", n)
//}
//
//func deleteRowDemo() {
//	sqlStr := "delete from user where id=?"
//	ret, err := db.Exec(sqlStr, 3)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected()
//	fmt.Printf("n =====> 🚀🚀🚀 %v\n", n)
//}
//
//func prepareQueryDemo() {
//	sqlStr := "select id, name, age from user where id > ?"
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(0)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var u user
//		err := rows.Scan(&u.id, &u.name, &u.age)
//		if err != nil {
//			fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//			return
//		}
//		fmt.Printf("u.id, u.name, u.age =====> 🚀🚀🚀 %v\n %v\n %v\n", u.id, u.name, u.age)
//	}
//
//}
//
//func mySelect(name string) {
//	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
//	fmt.Printf("SQL:%s\n", sqlStr)
//	var u user
//	err := db.QueryRow(sqlStr).Scan(&u.age, &u.id, &u.name)
//	if err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//		return
//	}
//	fmt.Printf("my success =====> 🚀🚀🚀")
//}
//
//func main() {
//	if err := initMySQL(); err != nil {
//		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
//	}
//	defer db.Close()
//	fmt.Printf("connect to db success =====> 🚀🚀🚀")
//	updateRowDemo()
//	queryRowDemo()
//	deleteRowDemo()
//	queryMultiRowDemo()
//	insertRowDemo()
//	prepareQueryDemo()
//	mySelect("cao")
//}
