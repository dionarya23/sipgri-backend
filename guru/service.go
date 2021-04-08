package guru

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterGuru(input RegisterGuruInput) (Guru, error)
	Login(input LoginGuruInput) (Guru, error)
	IsNipAvalaible(input CheckNipInput) (bool, error)
	IsEmailAvalaible(input CheckEmailInput) (bool, error)
	GetGuruByNip(nipGuru string) (Guru, error)
	GetAllGuru() ([]Guru, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterGuru(input RegisterGuruInput) (Guru, error) {
	guru := Guru{}
	guru.Nip = input.Nip
	guru.Nama = input.Nama
	guru.Email = input.Email
	guru.NomorTelepon = input.NomorTelepon
	guru.Type = input.Type

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return guru, err
	}

	guru.Password = string(passwordHash)
	newGuru, err := s.repository.Save(guru)

	if err != nil {
		return guru, err
	}

	return newGuru, nil
}

func (s *service) Login(input LoginGuruInput) (Guru, error) {
	email := input.Email
	password := input.Password

	guru, err := s.repository.FindByEmail(email)

	if err != nil {
		return guru, err
	}

	err_ := bcrypt.CompareHashAndPassword([]byte(guru.Password), []byte(password))

	if err != nil {
		return guru, err_
	}

	return guru, nil
}

func (s *service) IsNipAvalaible(input CheckNipInput) (bool, error) {
	guru, err := s.repository.FindByNip(input.Nip)
	if err != nil {
		return false, err
	}

	if guru.Nip == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) IsEmailAvalaible(input CheckEmailInput) (bool, error) {
	guru, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if guru.Nip == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) GetGuruByNip(nipGuru string) (Guru, error) {
	guru, err := s.repository.FindByNip(nipGuru)
	if err != nil {
		return guru, err
	}

	if guru.Nip == "" {
		return guru, errors.New("Guru not found")
	}

	return guru, nil
}

func (s *service) GetAllGuru() ([]Guru, error) {

	guru, err := s.repository.FindAll()
	if err != nil {
		return guru, err
	}

	return guru, nil
}
