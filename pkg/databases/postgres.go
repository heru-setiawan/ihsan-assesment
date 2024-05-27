package databases

import (
	"assesment/pkg/configs"
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgres(config configs.Postgres, traceProvider trace.TracerProvider) (*postgresDB, error) {
	database := new(postgresDB)
	database.config = config
	if err := database.connect(traceProvider); err != nil {
		return nil, err
	}

	return database, nil
}

type postgresDB struct {
	Conn   *gorm.DB
	config configs.Postgres
}

func (db *postgresDB) connect(traceProvider trace.TracerProvider) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", db.config.Host, db.config.User, db.config.Password, db.config.Database, db.config.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   fmt.Sprintf("%s.%s.", db.config.Database, db.config.Schema),
		},
	})
	if err != nil {
		return err
	}

	conn.Use(otelgorm.NewPlugin(otelgorm.WithTracerProvider(traceProvider)))

	db.Conn = conn
	return nil
}

func (db *postgresDB) Migrate(models ...any) error {
	db.Conn.Exec("CREATE SCHEMA IF NOT EXISTS " + db.config.Schema)
	db.Conn.AutoMigrate(models...)
	return nil
}
