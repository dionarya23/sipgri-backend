package jadwal

type InputJadwal struct {
	KodeMengajar string `json:"kode_mengajar" binding:"required"`
	IDKelas      int    `json:"id_kelas" binding:"required"`
	JamMulai     string `json:"jam_mulai" binding:"required"`
	JamSelesai   string `json:"jam_selesai" binding:"required"`
	Hari         string `json:"hari" binding:"required"`
}

type InputParamsIDJadwal struct {
	IDJadwal int `uri:"id_jadwal" binding:"required"`
}
