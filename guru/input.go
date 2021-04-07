package guru

type RegisterGuruInput struct {
	Nip          string `json:"nip" binding:"required"`
	Nama         string `json:"nama" binding:"required"`
	NomorTelepon string `json:"nomor_telepon" binding:"required"`
	Email        string `json:"email" binding:"required, email"`
	Password     string `json:"password" binding:"required"`
	Type         string `json:"type" binding:"required"`
}
