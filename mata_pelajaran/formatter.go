package mata_pelajaran

type MataPelajaranFormatter struct {
	IdMataPelajaran int    `json:"id_mata_pelajaran"`
	MataPelajaran   string `json:"mata_pelajaran"`
	Kelompok        string `json:"kelompok"`
}

func FormatMataPelajaranDetail(mata_pelajaran MataPelajaran) MataPelajaranFormatter {
	formatter := MataPelajaranFormatter{}
	formatter.IdMataPelajaran = mata_pelajaran.IdMataPelajaran
	formatter.MataPelajaran = mata_pelajaran.MataPelajaran
	formatter.Kelompok = mata_pelajaran.Kelompok
	return formatter
}

func FormatListMataPelajaran(mata_pelajaran []MataPelajaran) []MataPelajaranFormatter {
	mataPelajaran := []MataPelajaranFormatter{}

	for _, mata_pelajaran_ := range mata_pelajaran {
		mataPelajaranFormatter := FormatMataPelajaranDetail(mata_pelajaran_)
		mataPelajaran = append(mataPelajaran, mataPelajaranFormatter)
	}

	return mataPelajaran
}
