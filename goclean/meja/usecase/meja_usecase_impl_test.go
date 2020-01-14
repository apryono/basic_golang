package usecase_test

import (
	"goclean/meja/mock"
	"goclean/meja/usecase"
	"goclean/model"
	"testing"
)

func TestGetById(t *testing.T) {

	t.Run("Test Normal Case", func(t *testing.T) {
		t.Run("Should Return Data", func(t *testing.T) {
			var mejaMock = model.Meja{
				ID:     5,
				Status: "OPEN",
			}

			mejaMockRepo := new(mock.MejaRepoMock)
			mejaMockRepo.On("GetById", mejaMock.ID).Return(&mejaMock, nil)

			mejaUsecase := usecase.CreateMejaUsecase(mejaMockRepo)
			meja, _ := mejaUsecase.GetById(5)
			if meja == nil {
				t.Fatalf("should return data but return nil")
			}
			if meja.ID != mejaMock.ID {
				t.Fatalf("Meja ID should return %v, got %v", mejaMock.ID, meja.ID)
			}
			if meja.Status != mejaMock.Status {
				t.Fatalf("Meja Status should return %v, got %v", mejaMock.Status, meja.Status)
			}
		})
	})

}
