package databases

import (
	"assesment/pkg/configs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB(config configs.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch config.Database.Driver {
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database)
		dialector = sqlserver.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Database, config.Database.Port)
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsuported database driver")
	}

	db, err := gorm.Open(
		dialector,
		&gorm.Config{
			FullSaveAssociations: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				TablePrefix:   fmt.Sprintf("%s.%s.", config.Database.Database, config.Database.Schema),
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
