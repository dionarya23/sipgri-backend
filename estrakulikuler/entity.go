package estrakulikuler

import "github.com/dionarya23/sipgri-backend/guru"

type Estrakulikuler struct {
	IDEstrakulikuler int `gorm:"primaryKey"`
	Jenis            string
	NipGuru          string
	Pembimbing       guru.Guru `gorm:"Foreignkey:NipGuru;association_foreignkey:Nip;"`
}
