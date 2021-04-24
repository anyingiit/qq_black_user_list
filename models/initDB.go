package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"qq_black_user_list/pkg/setting"
)

var Db *sql.DB

func InitDB() (err error) {
	fmt.Println("type:", setting.DatabaseType)
	Db, err = sql.Open(setting.DatabaseType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseUser,
			setting.DatabasePwd,
			setting.DatabaseHost,
			setting.DatabaseName))
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	return nil
}
