package jadwal

type Service interface {
	Create(inputData InputJadwal) (Jadwal, error)
	FindAll() ([]Jadwal, error)
	FindByIdJadwal(input InputParamsIDJadwal) (Jadwal, error)
	UpdateById(inputID InputParamsIDJadwal, inputData InputJadwal) (Jadwal, error)
	Delete(inputId InputParamsIDJadwal) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(inputData InputJadwal) (Jadwal, error) {
	jadwal := Jadwal{}

	jadwal.KodeMengajar = inputData.KodeMengajar
	jadwal.IDKelas = inputData.IDKelas
	jadwal.JamMulai = inputData.JamMulai
	jadwal.JamSelesai = inputData.JamSelesai
	jadwal.Hari = inputData.Hari

	newJadwal, err := s.repository.Save(jadwal)

	if err != nil {
		return newJadwal, err
	}

	return newJadwal, nil
}

func (s *service) FindAll() ([]Jadwal, error) {
	jadwal, err := s.repository.FindAll()
	if err != nil {
		return jadwal, err
	}

	return jadwal, nil
}

func (s *service) FindByIdJadwal(input InputParamsIDJadwal) (Jadwal, error) {
	jadwal, err := s.repository.FindOne(input.IDJadwal)
	if err != nil {
		return jadwal, err
	}

	return jadwal, nil
}

func (s *service) UpdateById(inputID InputParamsIDJadwal, inputData InputJadwal) (Jadwal, error) {
	jadwal, err := s.repository.FindOne(inputID.IDJadwal)

	if err != nil {
		return jadwal, err
	}

	jadwal.KodeMengajar = inputData.KodeMengajar
	jadwal.IDKelas = inputData.IDKelas
	jadwal.JamMulai = inputData.JamMulai
	jadwal.JamSelesai = inputData.JamSelesai
	jadwal.Hari = inputData.Hari

	updatedJadwal, err := s.repository.Update(jadwal)

	if err != nil {
		return updatedJadwal, err
	}

	return updatedJadwal, nil
}

func (s *service) Delete(inputId InputParamsIDJadwal) error {
	err := s.repository.Delete(inputId.IDJadwal)

	return err
}
