package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
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
//@param tab 表名
func (r *MysqlRepository) Query(tab string) []Data {
	db := r.Connect()
	sql := "SELECT * FROM "+tab
	rows, err := db.Query(sql)
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
//@param tab 表名
func (r *MysqlRepository) Empty(tab string){
	db := r.Connect()
	_,err := db.Query("truncate table "+tab)
	if err != nil{
		panic(err)
	}
}

//插入数据
//@param tab 表名
//@param data 要插入的数据集合
func (r *MysqlRepository) Insert(tab string,data []Data){
	db := r.Connect()
	for _,col := range data {
		sqlKey := "("
		sqlValue := "("
		sql := "INSERT "+tab+" "
		var insertVal []interface{}
		for k,col1 := range col.Record {
			sqlKey += k+","
			sqlValue += "?,"
			insertVal = append(insertVal,col1)
		}
		sqlKey = strings.TrimRight(sqlKey, ",")
		sqlValue = strings.TrimRight(sqlValue, ",")
		sqlKey += ")"
		sqlValue += ")"
		newSql := sql+sqlKey+" values "+sqlValue
		//入库
		stmt, err := db.Prepare(newSql)
		_,err = stmt.Exec(insertVal...)
		if err != nil{
			panic(err)
		}
	}
}