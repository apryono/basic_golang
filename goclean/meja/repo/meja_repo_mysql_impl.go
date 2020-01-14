package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"goclean/meja"
	"goclean/model"
)

var _ meja.MejaRepo = &MejaRepoMysqlImpl{}

type MejaRepoMysqlImpl struct {
	db *sql.DB
}

func (m *MejaRepoMysqlImpl) BeginTrans() (*sql.Tx, error) {
	return m.db.Begin()
}

func CreateMejaRepoMysqlImpl(db *sql.DB) meja.MejaRepo {
	return &MejaRepoMysqlImpl{db}
}

func (m *MejaRepoMysqlImpl) GetById(id int) (*model.Meja, error) {

	qry := "SELECT id, status FROM ms_meja WHERE id = ?"

	meja := model.Meja{}

	err := m.db.QueryRow(qry, id).Scan(&meja.ID, &meja.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[MejaRepoMysqlImpl.GetById] Error when query get meja by id : %w", err)
	}

	return &meja, nil
}

func (m *MejaRepoMysqlImpl) Insert(meja *model.Meja, tx *sql.Tx) error {

	qry := "INSERT INTO ms_meja(id, status) VALUES (?, ?)"

	var err error
	if tx == nil {
		_, err = tx.Exec(qry, meja.ID, meja.Status)
		if err != nil {
			return fmt.Errorf("[MejaRepoMysqlImpl.Insert] Error when insert with tx : %w", err)
		}
	} else {
		_, err = m.db.Exec(qry, meja.ID, meja.Status)
		if err != nil {
			return fmt.Errorf("[MejaRepoMysqlImpl.Insert] Error when insert without tx : %w", err)
		}
	}

	return nil
}
