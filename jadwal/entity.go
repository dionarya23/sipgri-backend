package jadwal

import (
	"github.com/dionarya23/sipgri-backend/kelas"
	"github.com/dionarya23/sipgri-backend/mengajar"
)

type Jadwal struct {
	IDJadwal     int `gorm:"primaryKey"`
	KodeMengajar string
	IDKelas      int
	JamMulai     string
	JamSelesai   string
	Hari         string
	Kelas        kelas.Kelas       `gorm:"foreignKey:IDKelas"`
	Pengajar     mengajar.Mengajar `gorm:"foreignKey:KodeMengajar"`
}
