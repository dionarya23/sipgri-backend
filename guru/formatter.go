package guru

type GuruFormatter struct {
	Nip          string `json:"nip" binding:"required"`
	Nama         string `json:"nama" binding:"required"`
	NomorTelepon string `json:"nomor_telepon" binding:"required"`
	Email        string `json:"email" binding:"required, email"`
	Type         string `json:"type" binding:"required"`
}

func FormatGuru(guru Guru) GuruFormatter {
	formatter := GuruFormatter{
		Nip:          guru.Nip,
		Nama:         guru.Nama,
		NomorTelepon: guru.NomorTelepon,
		Email:        guru.Email,
		Type:         guru.Type,
	}

	return formatter
}
