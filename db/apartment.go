package db

import (
	"DBCourse/req"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 获取所有空闲宿舍
func MGetApartment() ([]Apartment, error) {
	res := make([]Apartment, 0)

	err := DB.Model(&Apartment{}).Where("count < ?", 4).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 学生申请宿舍
func ApplyApartment(r req.ApplyApartment) error {
	// 1. 插入租聘表记录
	// 2. 插入发票表记录
	// 3. apartment表入住人数加1

	t, err := ParseTime(r.BeginTime)
	if err != nil {
		return err
	}
	t.Add(time.Hour * 24 * time.Duration(r.Duration))
	return DB.Transaction(func(tx *gorm.DB) error {
		//1.
		err = tx.Model(&Lease{}).Create(&Lease{
			StudentID:   r.StudentID,
			ApartmentID: r.ApartmentID,
			RoomID:      r.RoomID,
			deadline:    t.Format("2006-01-02"),
		}).Error
		if err != nil {
			return err
		}
		//2.
		err = tx.Model(&Invoice{}).Create(&Invoice{
			StudentID:   r.StudentID,
			ApartmentID: r.ApartmentID,
			RoomID:      r.RoomID,
			Deadline:    t.Format("2006-01-02"),
			Money:       r.Duration * 10,
			PayTime:     "",
			PayMethod:   "",
			Status:      "未支付",
		}).Error
		if err != nil {
			return err
		}
		// 3.
		err = tx.Model(&Apartment{}).Where(&Apartment{ApartmentID: r.ApartmentID, RoomID: r.RoomID}).Update("count", "count+1").Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 根据学号查询学生的宿舍信息
func GetStudentApartment(studentID int64) (req.ApartmentInfoResp, error) {
	//  查出住在哪个宿舍
	a := Apartment{}
	err := DB.Model(&Lease{}).Where(&Lease{StudentID: studentID}).First(&a).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}
	// 根据宿舍查室友
	rm := make([]Student, 0)
	err = DB.Model(&Lease{}).Where(&Lease{ApartmentID: a.ApartmentID, RoomID: a.RoomID}).Find(&rm).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}
	rmResp := make([]string, 0)
	for _, s := range rm {
		rmResp = append(rmResp, fmt.Sprintf("%d%s", s.UserID, s.StudentName))
	}
	// 查管理员
	adminer := Admin{}
	err = DB.Model(&Admin{}).Where(&Admin{AdminID: a.AdminID}).First(&adminer).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}

	// 查询过期时间
	l := Lease{}
	DB.Model(&Lease{}).Where(&Lease{StudentID: studentID}).First(&l)
	return req.ApartmentInfoResp{
		ApartmentID: a.ApartmentID,
		RoomID:      a.RoomID,
		Roommate:    rmResp,
		AdminName:   adminer.AdminName,
		Deadline:    l.deadline,
	}, nil
}

// 管理员查询宿舍信息
func AdminQueryApartment(apartmentID int64, roomID int64) (req.ApartmentInfoResp, error) {
	// 查询管理者姓名
	a := Apartment{}
	err := DB.Model(&Apartment{}).Where(&Apartment{ApartmentID: apartmentID, RoomID: roomID}).First(&a).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}
	ad := &Admin{}
	err = DB.Model(&Admin{}).Where(&Admin{AdminID: a.AdminID}).First(&ad).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}
	// 查询租聘室友
	rm := make([]Student, 0)
	err = DB.Model(&Lease{}).Where(&Lease{ApartmentID: apartmentID, RoomID: roomID}).Find(&rm).Error
	if err != nil {
		return req.ApartmentInfoResp{}, err
	}
	rmResp := make([]string, 0)
	for _, s := range rm {
		rmResp = append(rmResp, fmt.Sprintf("%d%s", s.UserID, s.StudentName))
	}

	return req.ApartmentInfoResp{
		AdminName: ad.AdminName,
		Roommate:  rmResp,
	}, nil
}
