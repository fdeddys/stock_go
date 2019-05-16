package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	Dbcon    *gorm.DB
	Errdb    error
	dbuser   string
	dbpass   string
	dbname   string
	dbaddres string
	dbport   string
	dbdebug  bool
	dbtype   string
	sslmode  string
)

func init() {

	fmt.Println("Starting connect to POSTGRES")
	dbuser = beego.AppConfig.DefaultString("db.postgres.user", "deddy")
	dbpass = beego.AppConfig.DefaultString("db.postgres.pass", "123")
	dbname = beego.AppConfig.DefaultString("db.postgres.name", "dbstock")
	dbaddres = beego.AppConfig.DefaultString("db.postgres.address", "127.0.0.1")
	dbport = beego.AppConfig.DefaultString("db.postgres.port", "5432")
	sslmode = beego.AppConfig.DefaultString("db.postgres.sslmode", "disable")
	dbdebug = beego.AppConfig.DefaultBool("db.postgres.debug", true)
	if dbOpen() != nil {
		fmt.Println("Failed connect to database...!")
		return
	}
	fmt.Println("connected successfull ... ")
}

// DbOpen ...
func dbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbaddres, dbport, dbuser, dbpass, dbname, sslmode)
	Dbcon, Errdb = gorm.Open("postgres", args)

	if Errdb != nil {
		logs.Error("open db Err ", Errdb)
		return Errdb
	}
	if errping := Dbcon.DB().Ping(); errping != nil {
		return errping
	}
	fmt.Println("connection ready [", Dbcon.CommonDB(), "]")
	return nil
}

// GetDbCon ...
func GetDbCon() *gorm.DB {

	if errping := Dbcon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = dbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}
	Dbcon.LogMode(dbdebug)
	return Dbcon
}
