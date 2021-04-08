package guru

import "gorm.io/gorm"

type Repository interface {
	Save(guru Guru) (Guru, error)
	FindByNip(nip string) (Guru, error)
	FindByEmail(email string) (Guru, error)
	FindAll() ([]Guru, error)
	Update(guru Guru, nipGuru string) (Guru, error)
	Delete(nipGuru string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(guru Guru) (Guru, error) {
	err := r.db.Create(&guru).Error

	if err != nil {
		return guru, err
	}

	return guru, nil

}

func (r *repository) FindByNip(nip string) (Guru, error) {
	var guru Guru
	err := r.db.Where("nip=?", nip).Find(&guru).Error
	if err != nil {
		return guru, err
	}

	return guru, nil
}

func (r *repository) FindByEmail(email string) (Guru, error) {
	var guru Guru
	err := r.db.Where("email=?", email).Find(&guru).Error

	if err != nil {
		return guru, err
	}
	return guru, nil
}

func (r *repository) FindAll() ([]Guru, error) {
	var guru []Guru
	err := r.db.Find(&guru).Error
	if err != nil {
		return guru, err
	}
	return guru, nil
}

func (r *repository) Update(guru Guru, nipGuru string) (Guru, error) {
	var guru_ Guru
	err := r.db.Model(&guru_).Where("nip=?", nipGuru).Updates(guru).Error
	if err != nil {
		return guru, err
	}

	return guru, nil
}

func (r *repository) Delete(nipGuru string) error {
	var guru Guru
	err := r.db.Where("nip=?", nipGuru).Delete(&guru).Error
	return err
}
