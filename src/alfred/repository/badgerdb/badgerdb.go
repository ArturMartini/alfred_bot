package badgerdb

import (
	"alfred/canonical"
	log "github.com/sirupsen/logrus"
)
import "github.com/dgraph-io/badger"

type Database struct {
	db *badger.DB
}

func Initialize() *Database {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/alfred"
	opts.ValueDir = "/tmp/alfred"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &Database{
		db,
	}
}

func (r Database) Update(config canonical.Configuration) error {
	updates := make(map[string]string)

	rootKey := config.UserID + "-"

	updates[rootKey+"email"] = config.Email
	updates[rootKey+"pass"] = config.Pass
	updates[rootKey+"template"] = config.Template

	txn := r.db.NewTransaction(true)
	for k, v := range updates {
		err := txn.Set([]byte(k), []byte(v)); if err == nil {
			log.WithError(err).Error("Error when set the values in txn")
			return err
		}
	}
	err := txn.Commit(func(err error){
		log.WithError(err).Error("Error when try commit txn")
	})
	return err
}
