package mata_pelajaran

import "errors"

type Service interface {
	CreateNewMataPelajaran(input InputNewMataPelajaran) (MataPelajaran, error)
	FindMataPelajaranByName(mataPelajaran string) (MataPelajaran, error)
	GetAllMataPelajaran() ([]MataPelajaran, error)
	GetOneMataPelajaran(IdMataPelajaran int) (MataPelajaran, error)
	UpdateMataPelajaranByID(inputID InputIDMataPelajaran, inputData InputNewMataPelajaran) (MataPelajaran, error)
	DeleteMataPelajaranById(IdMataPelajaran int) error
}

type service struct {
	repository Repostory
}

func NewService(repository Repostory) *service {
	return &service{repository}
}

func (s *service) CreateNewMataPelajaran(input InputNewMataPelajaran) (MataPelajaran, error) {
	mataPelajaran := MataPelajaran{}

	mataPelajaran.MataPelajaran = input.MataPelajaran
	mataPelajaran.Kelompok = input.Kelompok

	newMataPelajaran, err := s.repository.Save(mataPelajaran)

	if err != nil {
		return newMataPelajaran, err
	}

	return newMataPelajaran, nil
}

func (s *service) FindMataPelajaranByName(mataPelajaran string) (MataPelajaran, error) {
	mataPelajarn, err := s.repository.FindByMataPelajaran(mataPelajaran)

	if err != nil {
		return mataPelajarn, err
	}

	return mataPelajarn, nil
}

func (s *service) GetAllMataPelajaran() ([]MataPelajaran, error) {
	mataPelajaran, err := s.repository.FindAll()
	if err != nil {
		return mataPelajaran, err
	}

	return mataPelajaran, nil
}

func (s *service) GetOneMataPelajaran(IdMataPelajaran int) (MataPelajaran, error) {
	mataPelajaran, err := s.repository.FindById(IdMataPelajaran)
	if err != nil {
		return mataPelajaran, err
	}

	return mataPelajaran, nil
}

func (s *service) UpdateMataPelajaranByID(inputID InputIDMataPelajaran, inputData InputNewMataPelajaran) (MataPelajaran, error) {
	mataPelajaran, err := s.repository.FindById(inputID.IdMataPelajaran)

	if err != nil {
		return mataPelajaran, err
	}

	if mataPelajaran.IdMataPelajaran == 0 {
		return mataPelajaran, errors.New("Guru not found")
	}

	mataPelajaran.MataPelajaran = inputData.MataPelajaran
	mataPelajaran.Kelompok = inputData.Kelompok

	updatedMataPelajaran, err := s.repository.Update(mataPelajaran, mataPelajaran.IdMataPelajaran)

	if err != nil {
		return mataPelajaran, err
	}

	return updatedMataPelajaran, nil

}

func (s *service) DeleteMataPelajaranById(IdMataPelajaran int) error {
	err := s.repository.DeleteByID(IdMataPelajaran)

	return err
}
