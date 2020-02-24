package main

import (
   "bufio"
   "fmt"
   "mysql_recover/mysql_recover/src/repository"
   "os"
   "strings"
)

func main() {
   //1->2
   mysqlInfo1 := make(map[int]string)
   mysqlInfo2 := make(map[int]string)

   input := bufio.NewScanner(os.Stdin)
   fmt.Printf("注：由 Pre 数据库导入到 Now 数据库（Pre -> Now）,Now中操作的表原记录会被清空 \n")
   fmt.Printf("请输入Pre表的信息，格式：Ip,端口号,用户名,密码,数据库名,需要操作的表 \n")
   input.Scan()
   m1 := strings.Split(input.Text(), ",")
   for i,col := range m1 {
      mysqlInfo1[i] = col
   }
   fmt.Printf("请输入Now表的信息，格式：Ip,端口号,用户名,密码,数据库名,需要操作的表 \n")
   input.Scan()
   m2 := strings.Split(input.Text(), ",")
   for i,col := range m2 {
      mysqlInfo2[i] = col
   }
   fmt.Printf("【注意】按下回车确定执行！ \n")
   input.Scan()
   connect1 := &repository.MysqlRepository{DbName: mysqlInfo1[4], Ip: mysqlInfo1[0], Port: mysqlInfo1[1], User: mysqlInfo1[2], Pw: mysqlInfo1[3]}
   mysqlData1 := connect1.Query(mysqlInfo1[5])
   connect2 := &repository.MysqlRepository{DbName: mysqlInfo2[4], Ip: mysqlInfo2[0], Port: mysqlInfo2[1], User: mysqlInfo2[2], Pw: mysqlInfo2[3]}
   //connect1 := &repository.MysqlRepository{DbName: "test1", Ip: "127.0.0.1", Port: "3306", User: "root", Pw: "root"}
   //connect2 := &repository.MysqlRepository{DbName: "test2", Ip: "127.0.0.1", Port: "3306", User: "root", Pw: "root"}
   connect2.Empty(mysqlInfo2[5])
   connect2.Insert(mysqlInfo2[5],mysqlData1)
   fmt.Printf("执行成功！请按回车退出... \n")
   input.Scan()
}