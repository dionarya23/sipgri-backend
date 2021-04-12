package estrakulikuler

import "gorm.io/gorm"

type Repository interface {
	Save(eskul Estrakulikuler) (Estrakulikuler, error)
	FindAll() ([]Estrakulikuler, error)
	FindByID(IDEskul int) (Estrakulikuler, error)
	FindByNipPembimbing(NipPembimbing string) (Estrakulikuler, error)
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
	err := r.db.Omit("Pembimbing").Create(&eskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindAll() ([]Estrakulikuler, error) {
	var eskul []Estrakulikuler

	err := r.db.Preload("Pembimbing").Find(&eskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindByID(IDEskul int) (Estrakulikuler, error) {
	var eskul Estrakulikuler

	err := r.db.Preload("Pembimbing").First(&eskul, IDEskul).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) FindByNipPembimbing(NipPembimbing string) (Estrakulikuler, error) {
	var eskul Estrakulikuler

	err := r.db.Preload("Pembimbing").First(&eskul, "nip_pembimbing=?", NipPembimbing).Error

	if err != nil {
		return eskul, err
	}

	return eskul, nil
}

func (r *repository) Update(updatedData Estrakulikuler) (Estrakulikuler, error) {
	err := r.db.Omit("Pembimbing").Save(&updatedData).Error

	if err != nil {
		return updatedData, err
	}

	return updatedData, nil
}

func (r *repository) Delete(IDEskul int) error {
	err := r.db.Where("id_estrakulikuler=?", IDEskul).Delete(&Estrakulikuler{}).Error
	return err
}
