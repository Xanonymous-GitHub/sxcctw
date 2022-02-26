package orm

import "github.com/Xanonymous-GitHub/sxcctw/db/pkg/schema"

func init() {
	var err error

	DB, err = Connect()
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&schema.Record{})
	if err != nil {
		panic(err)
	}
}
