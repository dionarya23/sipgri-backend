package estrakulikuler

type InputNewEskull struct {
	Jenis         string `json:"jenis" binding:"required"`
	NipPembimbing string `json:"nip_pembimbing" binding:"required"`
}

type InputIDEskul struct {
	IDEstrakulikuler int `uri:"id_estrakulikuler" binding:"required"`
}

type InputNipGuru struct {
	NipPembimbing string `uri:"nip_pembimbing" binding:"required"`
}
