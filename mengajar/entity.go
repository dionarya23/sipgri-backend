package mengajar

import (
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
)

type Mengajar struct {
	KodeMengajar    string `gorm:"primaryKey"`
	NipGuru         string
	IDMataPelajaran int
	Guru            guru.Guru                    `gorm:"foreignKey:NipGuru"`
	MataPelajaran   mata_pelajaran.MataPelajaran `gorm:"Foreignkey:IDMataPelajaran;association_foreignkey:IdMataPelajaran;"`
}
