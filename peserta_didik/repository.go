package peserta_didik

import "gorm.io/gorm"

type Repository interface {
	Save(pesertaDidik PesertaDidik) (PesertaDidik, error)
	GetAll() ([]PesertaDidik, error)
	GetByNisn(Nisn string) (PesertaDidik, error)
	GetByNis(Nis string) (PesertaDidik, error)
	Update(Nisn string, NewDataPesertaDidik PesertaDidik) (PesertaDidik, error)
	Delete(Nisn string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(pesertaDidik PesertaDidik) (PesertaDidik, error) {
	err := r.db.Create(&pesertaDidik).Error

	if err != nil {
		return pesertaDidik, err
	}

	return pesertaDidik, nil
}

func (r *repository) GetAll() ([]PesertaDidik, error) {
	var pesertaDidik []PesertaDidik

	err := r.db.Find(&pesertaDidik).Error
	if err != nil {
		return pesertaDidik, err
	}

	return pesertaDidik, nil
}

func (r *repository) GetByNisn(Nisn string) (PesertaDidik, error) {
	var pesertaDidik PesertaDidik
	err := r.db.Where("nisn=?", Nisn).Find(&pesertaDidik).Error
	if err != nil {
		return pesertaDidik, err
	}

	return pesertaDidik, nil
}

func (r *repository) GetByNis(Nis string) (PesertaDidik, error) {
	var pesertaDidik PesertaDidik
	err := r.db.Where("nis=?", Nis).Find(&pesertaDidik).Error
	if err != nil {
		return pesertaDidik, err
	}

	return pesertaDidik, nil
}

func (r *repository) Update(Nisn string, NewDataPesertaDidik PesertaDidik) (PesertaDidik, error) {
	var pesertaDidik PesertaDidik
	err := r.db.Model(&pesertaDidik).Where("nisn=?", Nisn).Updates(NewDataPesertaDidik).Error
	if err != nil {
		return pesertaDidik, err
	}

	return NewDataPesertaDidik, nil
}

func (r *repository) Delete(Nisn string) error {
	var pesertaDidik PesertaDidik
	err := r.db.Where("nisn=?", Nisn).Delete(&pesertaDidik).Error
	return err
}
