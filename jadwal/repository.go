package jadwal

import "gorm.io/gorm"

type Repository interface {
	Save(jadwal Jadwal) (Jadwal, error)
	FindAll() ([]Jadwal, error)
	FindOne(IDKelas int) (Jadwal, error)
	Update(UpdatedData Jadwal) (Jadwal, error)
	Delete(IDKelas int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(jadwal Jadwal) (Jadwal, error) {
	err := r.db.Omit("Kelas", "Pengajar").Create(&jadwal).Error

	if err != nil {
		return jadwal, err
	}

	return jadwal, nil
}

func (r *repository) FindAll() ([]Jadwal, error) {
	var jadwal []Jadwal
	err := r.db.Find(&jadwal).Error

	if err != nil {
		return jadwal, err
	}

	return jadwal, nil
}

func (r *repository) FindOne(IDKelas int) (Jadwal, error) {
	var jadwal Jadwal

	err := r.db.Where("id_kelas=?", IDKelas).Find(&jadwal).Error

	if err != nil {
		return jadwal, err
	}

	return jadwal, nil
}

func (r *repository) Update(UpdatedData Jadwal) (Jadwal, error) {
	err := r.db.Omit("Kelas", "Pengajar").Save(&UpdatedData).Error

	if err != nil {
		return UpdatedData, err
	}

	return UpdatedData, nil
}

func (r *repository) Delete(IDKelas int) error {
	err := r.db.Where("id_kelas=?", IDKelas).Delete(&Jadwal{}).Error

	return err
}
