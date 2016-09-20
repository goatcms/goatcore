package mscope

import (
	"fmt"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/scope"
)

// Insert new data to entity in context
func Insert(s scope.Scope, daoID string, entity interface{}) (interface{}, error) {
	// Get dao
	daoIns, err := s.DP().Get(daoID)
	if err != nil {
		return nil, err
	}
	dao, ok := daoIns.(db.DAO)
	if !ok {
		return nil, fmt.Errorf("%v is not a dao instance", daoIns)
	}
	// Get scope transaction
	txIns, err := s.DP().Get(scope.TX)
	if err != nil {
		return nil, err
	}
	tx, ok := txIns.(db.TX)
	if !ok {
		return nil, fmt.Errorf("%v is not a db.TX instance", txIns)
	}
	return dao.Insert(tx, entity)
}

// Delete delete entity in context
func Delete(s scope.Scope, daoID string, id int64) error {
	// Get dao
	daoIns, err := s.DP().Get(daoID)
	if err != nil {
		return err
	}
	dao, ok := daoIns.(db.DAO)
	if !ok {
		return fmt.Errorf("%v is not a dao instance", daoIns)
	}
	// Get scope transaction
	txIns, err := s.DP().Get(scope.TX)
	if err != nil {
		return err
	}
	tx, ok := txIns.(db.TX)
	if !ok {
		return fmt.Errorf("%v is not a db.TX instance", txIns)
	}
	return dao.Delete(tx, id)
}
