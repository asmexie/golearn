package main

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
/*创建数据库表
CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`departname` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)
*/
func main() {
	//root:账户名
	//123456:账户密码
	//student:数据库
	db, err := sql.Open("mysql","root:123456@/student")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("xielechuan","研发部门","2012-12-09")
	checkErr(err)
	res, err = stmt.Exec("codeme","财务部门","2014-12-12")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("UPDATE userinfo set username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("asmexie",id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}