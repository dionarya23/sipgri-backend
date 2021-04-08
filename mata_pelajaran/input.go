package mata_pelajaran

type InputNewMataPelajaran struct {
	MataPelajaran string `json:"mata_pelajaran" binding:"required"`
	Kelompok      string `json:"kelompok" binding:"required"`
}

type InputIDMataPelajaran struct {
	IdMataPelajaran int `uri:"id_mata_pelajaran" binding:"required"`
}
