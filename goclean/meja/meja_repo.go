package meja

import "goclean/model"

import "database/sql"

type MejaRepo interface {
	BeginTrans() (*sql.Tx, error)

	GetById(id int) (*model.Meja, error)
	Insert(meja *model.Meja, tx *sql.Tx) error
}
