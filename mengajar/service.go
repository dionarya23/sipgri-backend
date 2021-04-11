package mengajar

import "errors"

type Service interface {
	Create(input InputNewMengajar) (Mengajar, error)
	GetAll() ([]Mengajar, error)
	GetByKodeMengajar(input InputKodeMengajar) (Mengajar, error)
	GetByNipGuru(input InputNipGuru) ([]Mengajar, error)
	UpdateByKodeMengajar(inputID InputKodeMengajar, inputData InputNewMengajar) (Mengajar, error)
	DeleteByKodeMengajar(inputID InputKodeMengajar) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputNewMengajar) (Mengajar, error) {
	mengajar := Mengajar{}

	mengajar.KodeMengajar = input.KodeMengajar
	mengajar.NipGuru = input.NipGuru
	mengajar.IDMataPelajaran = input.IDMataPelajaran

	newMengajar, err := s.repository.Save(mengajar)
	if err != nil {
		return newMengajar, err
	}

	return newMengajar, nil
}

func (s *service) GetAll() ([]Mengajar, error) {
	mengajar, err := s.repository.FindAll()

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (s *service) GetByKodeMengajar(input InputKodeMengajar) (Mengajar, error) {
	mengajar, err := s.repository.FindByKodeMengajar(input.KodeMengajar)

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (s *service) GetByNipGuru(input InputNipGuru) ([]Mengajar, error) {
	mengajar, err := s.repository.FindByKodeNipGuru(input.NipGuru)

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (s *service) UpdateByKodeMengajar(inputID InputKodeMengajar, inputData InputNewMengajar) (Mengajar, error) {
	mengajar, err := s.repository.FindByKodeMengajar(inputID.KodeMengajar)

	if err != nil {
		return mengajar, err
	}

	if mengajar.KodeMengajar == "" {
		return mengajar, errors.New("Mengajar not found")
	}

	mengajar.KodeMengajar = inputData.KodeMengajar
	mengajar.NipGuru = inputData.NipGuru
	mengajar.IDMataPelajaran = inputData.IDMataPelajaran

	updatedMengajar, err := s.repository.Update(inputID.KodeMengajar, mengajar)
	if err != nil {
		return updatedMengajar, err
	}

	return updatedMengajar, nil

}

func (s *service) DeleteByKodeMengajar(inputID InputKodeMengajar) error {
	err := s.repository.Delete(inputID.KodeMengajar)

	return err
}
