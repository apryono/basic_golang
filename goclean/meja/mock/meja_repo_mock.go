package mock

import (
	"database/sql"
	"goclean/meja"
	"goclean/model"

	"github.com/stretchr/testify/mock"
)

var _ meja.MejaRepo = &MejaRepoMock{}

type MejaRepoMock struct {
	mock.Mock // Type Embedding
}

func (m *MejaRepoMock) BeginTrans() (*sql.Tx, error) {
	panic("Not Implemented Yet")
}

func (m *MejaRepoMock) GetById(id int) (*model.Meja, error) {
	ret := m.Called(id)

	return ret.Get(0).(*model.Meja), ret.Error(1)
}

func (m *MejaRepoMock) Insert(meja *model.Meja, tx *sql.Tx) error {
	panic("Not Implemented Yet")
}
