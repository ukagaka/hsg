package session

import (
	"hxsg/logger"
)

// Begin a transaction
func (s *Session) Begin() (err error) {
	logger.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		logger.Error(err.Error())
		return
	}
	return
}

// Commit a transaction
func (s *Session) Commit() (err error) {
	logger.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		logger.Error(err.Error())
	}
	return
}

// Rollback a transaction
func (s *Session) Rollback() (err error) {
	logger.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		logger.Error(err.Error())
	}
	return
}
