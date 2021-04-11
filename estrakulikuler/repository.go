package estrakulikuler

import "gorm.io/gorm"

type Repository interface {
	Save(eskul Estrakulikuler) (Estrakulikuler, error)
	FindAll() ([]Estrakulikuler, error)
	FindByID(IDEskul int) (Estrakulikuler, error)
	FindByNipGuru(NipGuru string) (Estrakulikuler, error)
	Update(updatedData Estrakulikuler) (Estrakulikuler, error)
	Delete(IDEskul int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(eskul Estrakulikuler) (Estrakulikuler, error) {
	err := r.db.Omit("Guru").Create(&eskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindAll() ([]Estrakulikuler, error) {
	var eskul []Estrakulikuler

	err := r.db.Preload("Guru").Find(&eskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindByID(IDEskul int) (Estrakulikuler, error) {
	var eskul Estrakulikuler

	err := r.db.Preload("Guru").First(&eskul, IDEskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindByNipGuru(NipGuru string) (Estrakulikuler, error) {
	var eskul Estrakulikuler

	err := r.db.Preload("Guru").First(&eskul, "nip_guru=?", NipGuru).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) Update(updatedData Estrakulikuler) (Estrakulikuler, error) {
	err := r.db.Omit("Guru").Save(&updatedData).Error

	if err != nil {
		return updatedData, err
	}

	return updatedData, nil
}

func (r *repository) Delete(IDEskul int) error {
	err := r.db.Where("id_eskrakulikuler=?", IDEskul).Delete(&Estrakulikuler{}).Error
	return err
}
