package db

type Student struct {
	UserID      int64
	StudentName string
	Password    string
	Age         int
}

// Admin 管理员表
type Admin struct {
	AdminID   int64
	AdminName string
	Password  string
}

// Inspection 检查表  记录管理员和公寓关系
type Inspection struct {
	AdminID        int64
	ApartmentID    int64
	InspectionTime string
	Status         string
	Comment        string
}

// 发票表
type Invoice struct {
	StudentID   int64
	ApartmentID int64
	RoomID      int64
	Deadline    string
	Money       int64
	PayTime     string
	PayMethod   string
	Status      string
}

type Apartment struct {
	ApartmentID int64
	RoomID      int64
	AdminID     int64
	Address     string
	Count       int64 //寝室已经住了的人数，最大为4
}

// 提示表
type Remind struct {
}

// Lease 租聘关系表
type Lease struct {
	StudentID   int64
	ApartmentID int64
	RoomID      int64
	deadline    string
}
