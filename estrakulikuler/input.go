package estrakulikuler

type InputNewEskull struct {
	Jenis   string `json:"jenis" binding:"required"`
	NipGuru string `json:"nip_guru" binding:"required"`
}

type InputIDEskul struct {
	IDEstrakulikuler int `uri:"id_estrakulikuler" binding:"required"`
}

type InputNipGuru struct {
	NipGuru string `uri:"nip_guru" binding:"required"`
}
