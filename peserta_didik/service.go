package peserta_didik

type Service interface {
	CreatePesertaDidik(input InputDataPesertaDidik) (PesertaDidik, error)
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
