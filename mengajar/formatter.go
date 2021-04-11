package mengajar

import (
	"fmt"

	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
)

type MengajarFormatter struct {
	KodeMengajar    string      `json:"kode_mengajar"`
	NipGuru         string      `json:"nip_guru"`
	IDMataPelajaran int         `json:"id_mata_pelajaran"`
	Guru            interface{} `json:"guru"`
	MataPelajaran   interface{} `json:"mata_pelajaran"`
}

func FormatMengajar(mengajar Mengajar) MengajarFormatter {
	formatter := MengajarFormatter{}

	formatter.KodeMengajar = mengajar.KodeMengajar
	formatter.NipGuru = mengajar.NipGuru
	formatter.IDMataPelajaran = mengajar.IDMataPelajaran

	if mengajar.Guru.Nip != "" {
		formatter.Guru = guru.FormatDetailGuru(mengajar.Guru)
	}

	fmt.Println(mengajar)

	if mengajar.MataPelajaran.IdMataPelajaran != 0 {
		formatter.MataPelajaran = mata_pelajaran.FormatMataPelajaranDetail(mengajar.MataPelajaran)
	}

	return formatter
}

func FormatListMengajar(mengajar []Mengajar) []MengajarFormatter {
	listFormatter := []MengajarFormatter{}

	for _, value := range mengajar {
		mengajarDetail := FormatMengajar(value)
		listFormatter = append(listFormatter, mengajarDetail)
	}

	return listFormatter
}
