package main

import "DBCourse/db"

func Init() {
	db.Init()
}

func main() {
	// 砍掉检查表和员工表
	Init()

}
