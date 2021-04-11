package mengajar

type InputNewMengajar struct {
	KodeMengajar    string `json:"kode_mengajar" binding:"required"`
	NipGuru         string `json:"nip_guru" binding:"required"`
	IDMataPelajaran int    `json:"id_mata_pelajaran" binding:"required"`
}

type InputKodeMengajar struct {
	KodeMengajar string `uri:"kode_mengajar" binding:"required"`
}

type InputNipGuru struct {
	NipGuru string `uri:"nip_guru" binding:"required"`
}
