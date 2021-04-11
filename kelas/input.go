package kelas

type InputNewKelas struct {
	NamaKelas string `json:"nama_kelas" binding:"required"`
	NipWali   string `json:"nip_wali" binding:"required"`
}

type InputIDKelas struct {
	IDKelas int `uri:"id_kelas" binding:"required"`
}

type InputNipWali struct {
	NipWali string `uri:"nip_wali" binding:"required"`
}
