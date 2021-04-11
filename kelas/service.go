package kelas

import "errors"

type Service interface {
	CreateKelas(input InputNewKelas) (Kelas, error)
	GetAll() ([]Kelas, error)
	GetById(input InputIDKelas) (Kelas, error)
	GetByNipWali(input InputNipWali) (Kelas, error)
	UpdateById(inputID InputIDKelas, inputData InputNewKelas) (Kelas, error)
	DeleteById(input InputIDKelas) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateKelas(input InputNewKelas) (Kelas, error) {
	kelas := Kelas{}
	kelas.NamaKelas = input.NamaKelas
	kelas.NipWali = input.NipWali

	newKelas, err := s.repository.Save(kelas)

	if err != nil {
		return newKelas, err
	}

	return newKelas, err

}

func (s *service) GetAll() ([]Kelas, error) {
	kelas, err := s.repository.FindAll()
	if err != nil {
		return kelas, err
	}

	return kelas, err
}

func (s *service) GetById(input InputIDKelas) (Kelas, error) {
	kelas, err := s.repository.FindByID(input.IDKelas)

	if err != nil {
		return kelas, err
	}

	return kelas, nil
}

func (s *service) GetByNipWali(input InputNipWali) (Kelas, error) {
	kelas, err := s.repository.FindByNipWali(input.NipWali)

	if err != nil {
		return kelas, err
	}

	return kelas, nil
}

func (s *service) UpdateById(inputID InputIDKelas, inputData InputNewKelas) (Kelas, error) {
	kelas, err := s.repository.FindByID(inputID.IDKelas)

	if err != nil {
		return kelas, err
	}

	if kelas.IDKelas == 0 {
		return kelas, errors.New("Kelas not found")
	}

	kelas.NamaKelas = inputData.NamaKelas
	kelas.NipWali = inputData.NipWali

	updatedKelas, err := s.repository.Update(kelas)
	if err != nil {
		return updatedKelas, err
	}

	return updatedKelas, nil
}

func (s *service) DeleteById(input InputIDKelas) error {
	isKelasExist, err := s.repository.FindByID(input.IDKelas)

	if err != nil {
		return err
	}

	if isKelasExist.IDKelas == 0 {
		return errors.New("Kelas Not Found")
	}

	err_ := s.repository.Delete(input.IDKelas)

	return err_
}
