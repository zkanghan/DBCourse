package db

func CheckStudent(s Student) error {
	res := &Student{}
	return DB.Model(&Student{}).Where(&Student{UserID: s.UserID, Password: s.Password}).First(res).Error
}

func CheckAdmin(a Admin) error {
	res := &Admin{}
	return DB.Model(&Admin{}).Where(&Admin{AdminID: a.AdminID, Password: a.Password}).First(res).Error
}
