package usecase

import (
	"errors"
	"fmt"
	"goclean/meja"
	"goclean/model"
)

type MejaUsecaseImpl struct {
	mejaRepo meja.MejaRepo
	// kursiRepo kursi.KursiRepo
}

func CreateMejaUsecase(mejaRepo meja.MejaRepo /*, kursiRepo kursi.KursiRepo*/) meja.MejaUsecase {
	return &MejaUsecaseImpl{mejaRepo /*, kursiRepo*/}
}

func (m *MejaUsecaseImpl) GetById(id int) (*model.Meja, error) {
	return m.mejaRepo.GetById(id)
}

func (m *MejaUsecaseImpl) Insert(meja *model.Meja) error {

	mejaVal, err := m.mejaRepo.GetById(meja.ID)
	if err != nil {
		return fmt.Errorf("[CreateMejaUsecase.Insert] Error when get meja by id' : %w", err)
	}

	fmt.Println(mejaVal)

	if mejaVal != nil {
		return errors.New("ID Meja sudah ada, silahkan masukkan id lain")
	}

	// tx, err := m.mejaRepo.BeginTrans()
	// if err != nil {
	// 	return fmt.Errorf("[CreateMejaUsecase.Insert] Error when Begin Trans to Insert Meja : %w", err)
	// }

	err = m.mejaRepo.Insert(meja, nil)
	// err = m.mejaRepo.Insert(meja, tx)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("[CreateMejaUsecase.Insert] Error when Insert Meja : %w", err)
	}

	// err = m.kursiRepo.Insert(menu, tx)
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("[CreateMejaUsecase.Insert] Error when Insert Menu : %w", err)
	// }

	// tx.Commit()
	// if err != nil {
	// 	return fmt.Errorf("[CreateMejaUsecase.Insert] Error when Commit Menu : %w", err)
	// }

	return nil

}
