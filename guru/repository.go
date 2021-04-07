package guru

import "gorm.io/gorm"

type Repository interface {
	Save(guru Guru) (Guru, error)
	FindByNip(nip string) (Guru, error)
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
