package guru

type GuruFormatter struct {
	Nip          string `json:"nip" binding:"required"`
	Nama         string `json:"nama" binding:"required"`
	NomorTelepon string `json:"nomor_telepon" binding:"required"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	Token        string `json:"token"`
}

func FormatGuru(guru Guru, token string) GuruFormatter {
	formatter := GuruFormatter{}
	formatter.Nip = guru.Nip
	formatter.Nama = guru.Nama
	formatter.NomorTelepon = guru.NomorTelepon
	formatter.Email = guru.Email
	formatter.Type = guru.Type
	formatter.Token = ""

	if token != "" {
		formatter.Token = token
	}

	return formatter
}
