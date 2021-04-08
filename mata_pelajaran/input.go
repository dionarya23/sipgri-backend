package mata_pelajaran

type InputNewMataPelajaran struct {
	MataPelajaran string `json:"mata_pelajaran" binding:"required"`
	Kelompok      string `json:"kelompok" binding:"required"`
}
