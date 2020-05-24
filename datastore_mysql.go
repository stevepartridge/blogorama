package blog

import (
	"fmt"

	_ "database/sql"

	"github.com/Boostport/migration"
	"github.com/Boostport/migration/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/db"
)

const (
	mysqlDbID = "blog"
)

// MySQL implementation of Datastore
type MySQL struct {
	Config DatastoreConfig
	conn   *sqlx.DB
}

// NewMySQL returns an instantiated Datastore for MySQL based on the provided DatastoreConfig
func NewMySQL(config DatastoreConfig) (Datastore, error) {

	log.Debug().Interface("config", config).Msg("NewMySQL")

	store := MySQL{
		Config: config,
	}

	db.MaxRetries = 25

	if config.Type == DatastoreTypeMySQLMock {

		db.AddMock(mysqlDbID)

		return store, nil
	}

	db.AddMySQL( // Add a new MySQL for easy connection reference
		mysqlDbID,
		store.Config.Host,
		store.Config.Port,
		store.Config.Name,
		store.Config.User,
		store.Config.Pass,
	)

	err := db.Check(mysqlDbID)
	if log.IfError(err) {
		return store, err
	}

	if config.Migrations {

		err = store.Migrations()
		if log.IfError(err) {
			return Store, err
		}

	}

	return store, err
}

func (store MySQL) Migrations() error {

	assetMigration := &migration.GoBindataMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "_sql/mysql",
	}

	list, err := assetMigration.ListMigrationFiles()
	log.IfError(err)

	// Create driver
	driver, err := mysql.NewFromDB(db.Conn(mysqlDbID))
	log.IfError(err)

	if err != nil {
		return err
	}

	// Run all up migrations
	applied, err := migration.Migrate(driver, assetMigration, migration.Up, 0)
	log.IfError(err)

	if applied > 0 {

		// List out the files that were used
		for i := range list {
			fmt.Println("\t", list[i])
		}

		if applied != 1 {
			log.Info().Msg(fmt.Sprintf("Applied %d migrations", applied))
		} else {
			log.Info().Msg(fmt.Sprintf("Applied %d migration", applied))
		}
	}

	return nil

}
