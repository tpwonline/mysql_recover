package main

import (
   "mysql_recover/mysql_recover/src/repository"
)

func main() {
   //1->2
   //mysqlInfo1 := make(map[string]string)
   ////mysqlInfo2 := make(map[string]string)
   //// 从标准输入流中接收输入数据
   //input := bufio.NewScanner(os.Stdin)
   //fmt.Printf("注：由 Pre 数据库导入到 Now 数据库（Pre -> Now）,Now中操作的表原记录会被清空 \n")
   //fmt.Printf("输入 Pre 数据库的信息\n")
   //fmt.Printf("输入 Pre 的 IP\n")
   //input.Scan()
   //mysqlInfo1["ip"] = input.Text()
   //fmt.Printf("输入 Pre 的 端口号\n")
   //input.Scan()
   //mysqlInfo1["port"] = input.Text()
   //fmt.Printf("输入 Pre 的 用户名\n")
   //input.Scan()
   //mysqlInfo1["user"] = input.Text()
   //fmt.Printf("输入 Pre 的 密码\n")
   //input.Scan()
   //mysqlInfo1["pw"] = input.Text()
   //fmt.Printf("输入 Pre 的 数据库名\n")
   //input.Scan()
   //mysqlInfo1["db_name"] = input.Text()
   //fmt.Printf("输入 Pre 的 需要操作的表\n")
   //input.Scan()
   //mysqlInfo1["op_table"] = input.Text()
   //connect1 := &repository.MysqlRepository{DbName: mysqlInfo1["db_name"], Ip: mysqlInfo1["ip"], Port: mysqlInfo1["port"], User: mysqlInfo1["user"], Pw: mysqlInfo1["pw"]}
   connect1 := &repository.MysqlRepository{DbName: "test1", Ip: "127.0.0.1", Port: "3306", User: "root", Pw: "root"}
   connect2 := &repository.MysqlRepository{DbName: "test2", Ip: "127.0.0.1", Port: "3306", User: "root", Pw: "root"}
   mysqlData1 := connect1.Query()
   connect2.Empty()
   connect2.Insert("x",mysqlData1)
   //fmt.Println(mysqlData1[0].Record["Id"])
}