package mata_pelajaran

type MataPelajaran struct {
	IdMataPelajaran int `gorm:"primaryKey"`
	MataPelajaran   string
	Kelompok        string
}
