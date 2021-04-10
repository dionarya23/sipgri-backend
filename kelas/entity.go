package kelas

import (
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/peserta_didik"
)

type Kelas struct {
	IDKelas      int `gorm:"primaryKey"`
	NamaKelas    string
	NipWali      string
	Guru         guru.Guru                    `gorm:"foreignKey:NipWali"`
	PesertaDidik []peserta_didik.PesertaDidik `gorm:"foreignKey:IDKelas"`
}
