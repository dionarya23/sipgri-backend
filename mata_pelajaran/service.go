package mata_pelajaran

type Service interface {
	CreateNewMataPelajaran(input InputNewMataPelajaran) (MataPelajaran, error)
	FindMataPelajaranByName(mataPelajaran string) (MataPelajaran, error)
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
