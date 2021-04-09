package peserta_didik

import "errors"

type Service interface {
	CreatePesertaDidik(input InputDataPesertaDidik) (PesertaDidik, error)
	GetAllPesertaDidik() ([]PesertaDidik, error)
	GetOnePesertaDidik(queryParams map[string][]string) (PesertaDidik, error)
	UpdatePesertaDidikByNisn(inputNisn InputNisn, InputData InputDataPesertaDidik) (PesertaDidik, error)
	DeleteByNisn(input InputNisn) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePesertaDidik(input InputDataPesertaDidik) (PesertaDidik, error) {
	pesertaDidik := PesertaDidik{}

	pesertaDidik.Nisn = input.Nisn
	pesertaDidik.Nis = input.Nis
	pesertaDidik.Nama = input.Nama
	pesertaDidik.JenisKelamin = input.JenisKelamin
	pesertaDidik.TempatLahir = input.TempatLahir
	pesertaDidik.TanggalLahir = input.TanggalLahir
	pesertaDidik.StatusDalamKeluarga = input.StatusDalamKeluarga
	pesertaDidik.AnakKe = input.AnakKe
	pesertaDidik.Alamat = input.Alamat
	pesertaDidik.NomorTelepon = input.NomorTelepon
	pesertaDidik.SekolahAsal = input.SekolahAsal
	pesertaDidik.TanggalDiterima = input.TanggalDiterima
	pesertaDidik.NamaAyah = input.NamaAyah
	pesertaDidik.NamaIbu = input.NamaIbu
	pesertaDidik.NamaWali = input.NamaWali
	pesertaDidik.NomorTeleponWali = input.NomorTeleponWali

	newPesertaDidik, err := s.repository.Save(pesertaDidik)

	if err != nil {
		return newPesertaDidik, err
	}

	return newPesertaDidik, nil
}

func (s *service) GetAllPesertaDidik() ([]PesertaDidik, error) {
	pesertaDidik, err := s.repository.GetAll()
	if err != nil {
		return pesertaDidik, err
	}

	return pesertaDidik, nil
}

func (s *service) GetOnePesertaDidik(queryParams map[string][]string) (PesertaDidik, error) {
	pesertaDidik := PesertaDidik{}

	if len(queryParams["nis"]) == 1 {
		peserta_didik, err := s.repository.GetByNis(queryParams["nis"][0])
		if err != nil {
			return peserta_didik, err
		}

		pesertaDidik = peserta_didik

	} else if len(queryParams["nisn"]) == 1 {
		peserta_didik, err := s.repository.GetByNisn(queryParams["nisn"][0])
		if err != nil {
			return peserta_didik, err
		}
		pesertaDidik = peserta_didik

	} else {
		return pesertaDidik, errors.New("something error in service peserta_didik")
	}

	return pesertaDidik, nil
}

func (s *service) UpdatePesertaDidikByNisn(inputNisn InputNisn, InputData InputDataPesertaDidik) (PesertaDidik, error) {
	peserta_didik, err := s.repository.GetByNisn(inputNisn.Nisn)
	if err != nil {
		return peserta_didik, err
	}

	if peserta_didik.Nisn == "" {
		return peserta_didik, errors.New("Peserta didik not found")
	}

	pesertaDidik := PesertaDidik{}

	pesertaDidik.Nisn = InputData.Nisn
	pesertaDidik.Nis = InputData.Nis
	pesertaDidik.Nama = InputData.Nama
	pesertaDidik.JenisKelamin = InputData.JenisKelamin
	pesertaDidik.TempatLahir = InputData.TempatLahir
	pesertaDidik.TanggalLahir = InputData.TanggalLahir
	pesertaDidik.StatusDalamKeluarga = InputData.StatusDalamKeluarga
	pesertaDidik.AnakKe = InputData.AnakKe
	pesertaDidik.Alamat = InputData.Alamat
	pesertaDidik.NomorTelepon = InputData.NomorTelepon
	pesertaDidik.SekolahAsal = InputData.SekolahAsal
	pesertaDidik.TanggalDiterima = InputData.TanggalDiterima
	pesertaDidik.NamaAyah = InputData.NamaAyah
	pesertaDidik.NamaIbu = InputData.NamaIbu
	pesertaDidik.NamaWali = InputData.NamaWali
	pesertaDidik.NomorTeleponWali = InputData.NomorTeleponWali

	updatedPesertaDidik, err := s.repository.Update(peserta_didik.Nisn, pesertaDidik)
	if err != nil {
		return peserta_didik, err
	}

	return updatedPesertaDidik, nil
}

func (s *service) DeleteByNisn(input InputNisn) error {
	err := s.repository.Delete(input.Nisn)

	return err
}
