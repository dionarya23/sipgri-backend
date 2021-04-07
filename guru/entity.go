package guru

type Guru struct {
	Nip          string `gorm:"primaryKey"`
	Nama         string
	NomorTelepon string
	Email        string
	Password     string
	Type         string
}
