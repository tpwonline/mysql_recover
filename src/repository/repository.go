package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


//数据库连接
//@param dbName 数据库名
//@param ip ip地址
//@param port 端口号
//@param user 用户名
//@param pw 密码
func Connect(dbName string,ip string,port string,user string,pw string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+pw+"@tcp("+ip+":"+port+")/"+dbName+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

//查询数据
func Query(){
	db := Connect("","","","","")
	rows, err := db.Query("SELECT * FROM x")
	//rows, err := db
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	//var res []interface{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record["Id"])
	}
}

////插入demo
//func insert() {
//
//	checkErr(err)
//
//	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
//	checkErr(err)
//	res, err := stmt.Exec("tony", 20, 1)
//	checkErr(err)
//	id, err := res.LastInsertId()
//	checkErr(err)
//	fmt.Println(id)
//}
//
////查询demo
//func query() {
//	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
//	checkErr(err)
//
//	rows, err := db.Query("SELECT * FROM user")
//	checkErr(err)
//
//	//普通demo
//	//for rows.Next() {
//	//	var userId int
//	//	var userName string
//	//	var userAge int
//	//	var userSex int
//
//	//	rows.Columns()
//	//	err = rows.Scan(&userId, &userName, &userAge, &userSex)
//	//	checkErr(err)
//
//	//	fmt.Println(userId)
//	//	fmt.Println(userName)
//	//	fmt.Println(userAge)
//	//	fmt.Println(userSex)
//	//}
//
//	//字典类型
//	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
//	columns, _ := rows.Columns()
//	scanArgs := make([]interface{}, len(columns))
//	values := make([]interface{}, len(columns))
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//
//	for rows.Next() {
//		//将行数据保存到record字典
//		err = rows.Scan(scanArgs...)
//		record := make(map[string]string)
//		for i, col := range values {
//			if col != nil {
//				record[columns[i]] = string(col.([]byte))
//			}
//		}
//		fmt.Println(record)
//	}
//}
//
////更新数据
//func update() {
//	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
//	checkErr(err)
//
//	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
//	checkErr(err)
//	res, err := stmt.Exec(21, 2, 1)
//	checkErr(err)
//	num, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println(num)
//}
//
////删除数据
//func remove() {
//	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
//	checkErr(err)
//
//	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
//	checkErr(err)
//	res, err := stmt.Exec(1)
//	checkErr(err)
//	num, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println(num)
//}
//
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}