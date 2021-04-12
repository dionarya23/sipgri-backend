package estrakulikuler

import "github.com/dionarya23/sipgri-backend/guru"

type Estrakulikuler struct {
	IDEstrakulikuler int `gorm:"primaryKey"`
	Jenis            string
	NipPembimbing    string
	Pembimbing       guru.Guru `gorm:"foreignKey:NipPembimbing"`
}
