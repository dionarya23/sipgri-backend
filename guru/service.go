package guru

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterGuru(input RegisterGuruInput) (Guru, error)
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
