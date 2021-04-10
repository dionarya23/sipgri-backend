package kelas

import (
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/peserta_didik"
)

type KelasFormatter struct {
	IDKelas      int         `json:"id_kelas"`
	NamaKelas    string      `json:"nama_kelas"`
	NipWali      string      `json:"nip_wali"`
	Guru         interface{} `json:"wali_kelas"`
	PesertaDidik interface{} `json:"peserta_didik"`
}

func FormatKelasDetail(kelas Kelas) KelasFormatter {
	formatter := KelasFormatter{}
	formatter.IDKelas = kelas.IDKelas
	formatter.NamaKelas = kelas.NamaKelas
	formatter.NipWali = kelas.NipWali

	if kelas.Guru.Nip != "" {
		formatter.Guru = guru.FormatDetailGuru(kelas.Guru)
	}

	if len(kelas.PesertaDidik) != 0 {
		formatter.PesertaDidik = peserta_didik.FormatPesertaDidikList(kelas.PesertaDidik)
	} else {
		formatter.PesertaDidik = make([]string, 0)
	}

	return formatter
}

func FormatKelasList(kelas []Kelas) []KelasFormatter {
	listKelas := []KelasFormatter{}

	for _, kelas := range kelas {
		formatKelas := FormatKelasDetail(kelas)
		listKelas = append(listKelas, formatKelas)
	}

	return listKelas
}
