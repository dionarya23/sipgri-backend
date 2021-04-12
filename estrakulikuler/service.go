package estrakulikuler

import (
	"errors"
)

type Service interface {
	Create(inputData InputNewEskull) (Estrakulikuler, error)
	GetAll() ([]Estrakulikuler, error)
	GetByID(inputIDEskul InputIDEskul) (Estrakulikuler, error)
	GetByNipGuru(inputNipGuru InputNipGuru) (Estrakulikuler, error)
	Update(inputIDEskul InputIDEskul, inputData InputNewEskull) (Estrakulikuler, error)
	Delete(inputIDEskul InputIDEskul) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(inputData InputNewEskull) (Estrakulikuler, error) {
	eskul := Estrakulikuler{}

	eskul.Jenis = inputData.Jenis
	eskul.NipPembimbing = inputData.NipPembimbing

	newEskul, err := s.repository.Save(eskul)

	if err != nil {
		return newEskul, err
	}

	return newEskul, nil
}

func (s *service) GetAll() ([]Estrakulikuler, error) {
	listEskull, err := s.repository.FindAll()

	if err != nil {
		return listEskull, err
	}

	return listEskull, nil
}

func (s *service) GetByID(inputIDEskul InputIDEskul) (Estrakulikuler, error) {
	eskul, err := s.repository.FindByID(inputIDEskul.IDEstrakulikuler)

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (s *service) GetByNipGuru(inputNipGuru InputNipGuru) (Estrakulikuler, error) {
	eskul, err := s.repository.FindByNipPembimbing(inputNipGuru.NipPembimbing)

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (s *service) Update(inputIDEskul InputIDEskul, inputData InputNewEskull) (Estrakulikuler, error) {
	eskul, err := s.repository.FindByID(inputIDEskul.IDEstrakulikuler)
	if err != nil {
		return eskul, err
	}

	if eskul.IDEstrakulikuler == 0 {
		return eskul, errors.New("Eskull not found")
	}

	eskul.Jenis = inputData.Jenis
	eskul.NipPembimbing = inputData.NipPembimbing

	updatedEskul, err := s.repository.Update(eskul)

	if err != nil {
		return updatedEskul, err
	}

	return updatedEskul, nil

}

func (s *service) Delete(inputIDEskul InputIDEskul) error {
	err := s.repository.Delete(inputIDEskul.IDEstrakulikuler)
	return err
}
