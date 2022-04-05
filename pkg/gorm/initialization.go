package gorm

import "github.com/Xanonymous-GitHub/sxcctw/pkg/schema"

func InitializeDBTables(dbClient DBClient) error {
	if err := dbClient.PrepareTables([]any{
		&schema.Record{},
	}); err != nil {
		return err
	}

	return nil
}
