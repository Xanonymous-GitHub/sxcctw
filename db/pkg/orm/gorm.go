package orm

import (
	"fmt"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/vp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUserName = vp.Cvp.GetString("dbUserName")
	dbPassword = vp.Cvp.GetString("dbPassword")
	dbHost     = vp.Cvp.GetString("dbHost")
	dbPort     = vp.Cvp.GetInt("dbPort")
	dbName     = vp.Cvp.GetString("dbName")
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dataSourceName}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
