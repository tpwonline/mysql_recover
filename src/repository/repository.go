package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	DbName,Ip,Port,User,Pw string  //数据库名,ip地址,端口号,用户名,密码
}

//数据库连接
func (r *MysqlRepository) Connect() *sql.DB {
	db, err := sql.Open("mysql", r.User+":"+r.Pw+"@tcp("+r.Ip+":"+r.Port+")/"+r.DbName+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

type Data struct {
	Record map[string]string
}

//查询数据
func (r *MysqlRepository) Query() []Data {
	db := r.Connect()
	rows, err := db.Query("SELECT * FROM x")
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var data []Data
	for rows.Next() {
		var d Data
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		d.Record = record
		data = append(data,d)
	}
	return data
}

//清空表
func (r *MysqlRepository) Empty(){
	db := r.Connect()
	_,err := db.Query("truncate table x")
	if err != nil{
		panic(err)
	}
}

//插入数据
//@param tab 表名
//@param data 要插入的数据集合
func (r *MysqlRepository) Insert(tab string,data []Data){
	//for _,col := range data {
	//	sqlKey := "("
	//	sqlValue := "("
	//	//sql := "INSERT "+tab    //Id,name) values (?,?)
	//	for k,col1 := range col.Record {
	//		sqlKey += k
	//		sqlValue += "?"
	//		//fmt.Println(k)
	//		//fmt.Println(col1)
	//	}
	//	db := r.Connect()
	//	stmt, err := db.Prepare(sql)
	//	_,err = stmt.Exec(1, "La")
	//	if err != nil{
	//		panic(err)
	//	}
	//}


	db := r.Connect()
	stmt, err := db.Prepare(`INSERT x (Id,name) values (?,?)`)
		var ddd []interface{}
		ddd = append(ddd,1)
		ddd = append(ddd,"La")
	_,err = stmt.Exec(ddd)
	if err != nil{
		panic(err)
	}
	stmt.Close()



	//fmt.Println(ddd)

	//db := r.Connect()
	//stmt, err := db.Prepare(`INSERT x (Id,name) values (?,?)`)
	//if err != nil{
	//	panic(err)
	//}
}