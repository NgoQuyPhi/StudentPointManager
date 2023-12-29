package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateTable(user, pass string) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/studentpointmanager?charset=utf8mb4&parseTime=True&loc=Local", user, pass)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_khoa` (`mk` varchar(5) NOT NULL,`tenkhoa` varchar(50) NOT NULL,`dienthoai` varchar(20) NOT NULL)")

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_lopchuyennganh` (`mlcn` varchar(10) NOT NULL,`tenlop` varchar(50) NOT NULL,`nienkhoa` varchar(3) NOT NULL,`siso` int(3) NOT NULL,`mk` varchar(5) NOT NULL)")

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_lophocphan` ( `mlcn` varchar(10) NOT NULL, `mhp` varchar(10) NOT NULL, `nhom` varchar(3) NOT NULL, `hocki` varchar(1) NOT NULL, `namhoc` varchar(5) NOT NULL )")

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_sinhvien` ( `msv` varchar(10) NOT NULL, `holot` varchar(20) NOT NULL, `ten` varchar(10) NOT NULL, `mlcn` varchar(10) NOT NULL, `sdd` varchar(12) NOT NULL, `email` varchar(30) NOT NULL )")

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_diemhocphan` ( `msv` varchar(10) NOT NULL, `mhp` varchar(10) NOT NULL, `A` int(2) NOT NULL, `B` int(2) NOT NULL, `C` int(2) NOT NULL )")

	db.Exec("CREATE TABLE IF NOT EXISTS `tbl_hocphan` ( `mhp` varchar(10) NOT NULL, `tenhocphan` varchar(40) NOT NULL, `tinchi` int(1) NOT NULL )")

	return db
}
