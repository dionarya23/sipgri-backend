package kelas

import "gorm.io/gorm"

type Repository interface {
	Save(kelas Kelas) (Kelas, error)
	FindAll() ([]Kelas, error)
	FindByID(IDKelas int) (Kelas, error)
	FindByNipWali(NipWali string) (Kelas, error)
	Update(updatedKelas Kelas) (Kelas, error)
	Delete(IDKelas int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(kelas Kelas) (Kelas, error) {
	err := r.db.Select("id_kelas", "nama_kelas", "nip_wali").Create(&kelas).Error

	if err != nil {
		return kelas, err
	}

	return kelas, nil
}

func (r *repository) FindAll() ([]Kelas, error) {
	var result []Kelas
	err := r.db.Preload("Guru").Preload("PesertaDidik").Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) FindByID(IDKelas int) (Kelas, error) {
	var kelas Kelas

	err := r.db.Preload("Guru").Preload("PesertaDidik").First(&kelas, IDKelas).Error
	if err != nil {
		return kelas, err
	}

	return kelas, nil
}

func (r *repository) FindByNipWali(NipWali string) (Kelas, error) {
	var kelas Kelas

	err := r.db.Preload("Guru").Preload("PesertaDidik").First(&kelas, "nip_wali=?", NipWali).Error
	if err != nil {
		return kelas, err
	}

	return kelas, nil
}

func (r *repository) Update(updatedKelas Kelas) (Kelas, error) {
	err := r.db.Save(&updatedKelas).Error

	if err != nil {
		return updatedKelas, err
	}

	return updatedKelas, nil
}

func (r *repository) Delete(IDKelas int) error {
	err := r.db.Where("id_kelas=?", IDKelas).Delete(&Kelas{}).Error
	return err
}
