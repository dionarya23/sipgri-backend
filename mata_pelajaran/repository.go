package mata_pelajaran

import "gorm.io/gorm"

type Repostory interface {
	Save(mata_pelajaran MataPelajaran) (MataPelajaran, error)
	FindAll() ([]MataPelajaran, error)
	FindByMataPelajaran(mata_pelajaran string) (MataPelajaran, error)
	FindById(IdMataPelajaran int) (MataPelajaran, error)
	Update(newMataPelajaran MataPelajaran, IdMataPelajaran int) (MataPelajaran, error)
	DeleteByID(IdMataPelajaran int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(mata_pelajaran MataPelajaran) (MataPelajaran, error) {
	err := r.db.Create(&mata_pelajaran).Error

	if err != nil {
		return mata_pelajaran, err
	}

	return mata_pelajaran, nil
}

func (r *repository) FindAll() ([]MataPelajaran, error) {
	var mata_pelajaran []MataPelajaran

	err := r.db.Find(&mata_pelajaran).Error
	if err != nil {
		return mata_pelajaran, err
	}

	return mata_pelajaran, nil
}

func (r *repository) FindByMataPelajaran(mata_pelajaran string) (MataPelajaran, error) {
	var mataPelajaran MataPelajaran

	err := r.db.Where("mata_pelajaran=?", mata_pelajaran).Find(&mataPelajaran).Error
	if err != nil {
		return mataPelajaran, err
	}

	return mataPelajaran, nil
}

func (r *repository) FindById(IdMataPelajaran int) (MataPelajaran, error) {
	var mata_pelajaran MataPelajaran

	err := r.db.Where("id_mata_pelajaran=?", IdMataPelajaran).Find(&mata_pelajaran).Error
	if err != nil {
		return mata_pelajaran, err
	}

	return mata_pelajaran, nil
}

func (r *repository) Update(newMataPelajaran MataPelajaran, IdMataPelajaran int) (MataPelajaran, error) {
	var mata_pelajaran MataPelajaran

	err := r.db.Model(&mata_pelajaran).Where("id_mata_pelajaran=?", IdMataPelajaran).Updates(newMataPelajaran).Error
	if err != nil {
		return mata_pelajaran, err
	}

	return newMataPelajaran, nil
}

func (r *repository) DeleteByID(IdMataPelajaran int) error {
	var mata_pelajaran MataPelajaran

	err := r.db.Where("id_mata_pelajaran=?", IdMataPelajaran).Delete(&mata_pelajaran).Error
	return err
}
