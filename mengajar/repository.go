package mengajar

import "gorm.io/gorm"

type Repository interface {
	Save(mengajar Mengajar) (Mengajar, error)
	FindAll() ([]Mengajar, error)
	FindByKodeMengajar(KodeMengajar string) (Mengajar, error)
	FindByKodeNipGuru(NipGuru string) ([]Mengajar, error)
	Update(KodeMengajar string, mengajar Mengajar) (Mengajar, error)
	Delete(KodeMengajar string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(mengajar Mengajar) (Mengajar, error) {
	err := r.db.Select("kode_mengajar", "nip_guru", "id_mata_pelajaran").Create(&mengajar).Error

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (r *repository) FindAll() ([]Mengajar, error) {
	var mengajar []Mengajar

	err := r.db.Preload("Guru").Preload("MataPelajaran").Find(&mengajar).Error

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (r *repository) FindByKodeMengajar(KodeMengajar string) (Mengajar, error) {
	var mengajar Mengajar

	err := r.db.Preload("Guru").Preload("MataPelajaran").First(&mengajar, "kode_mengajar=?", KodeMengajar).Error

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (r *repository) FindByKodeNipGuru(NipGuru string) ([]Mengajar, error) {
	var mengajar []Mengajar
	err := r.db.Preload("Guru").Preload("MataPelajaran").Find(&mengajar, "nip_guru=?", NipGuru).Error

	if err != nil {
		return mengajar, err
	}

	return mengajar, nil
}

func (r *repository) Update(KodeMengajar string, mengajar Mengajar) (Mengajar, error) {

	err := r.db.Model(&Mengajar{}).Select("kode_mengajar", "nip_guru", "id_mata_pelajaran").Where("kode_mengajar=?", KodeMengajar).Updates(mengajar).Error
	if err != nil {
		return mengajar, err
	}

	return mengajar, err
}

func (r *repository) Delete(KodeMengajar string) error {
	err := r.db.Where("kode_mengajar=?", KodeMengajar).Delete(&Mengajar{}).Error
	return err
}
