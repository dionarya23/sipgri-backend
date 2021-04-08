package guru

type GuruAuthFormatter struct {
	Nip          string `json:"nip"`
	Nama         string `json:"nama"`
	NomorTelepon string `json:"nomor_telepon"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	Token        string `json:"token"`
}

func FormatAuthGuru(guru Guru, token string) GuruAuthFormatter {
	formatter := GuruAuthFormatter{}
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

type DetailGuruFormatter struct {
	Nip          string `json:"nip"`
	Nama         string `json:"nama"`
	NomorTelepon string `json:"nomor_telepon"`
	Email        string `json:"email"`
	Type         string `json:"type"`
}

func FormatDetailGuru(guru Guru) DetailGuruFormatter {
	formatter := DetailGuruFormatter{}
	formatter.Nip = guru.Nip
	formatter.Nama = guru.Nama
	formatter.NomorTelepon = guru.NomorTelepon
	formatter.Email = guru.Email
	formatter.Type = guru.Type

	return formatter
}

func FormatListGuru(guru []Guru) []DetailGuruFormatter {
	gurusFormatter := []DetailGuruFormatter{}
	for _, guru := range guru {
		guruFormatter := FormatDetailGuru(guru)
		gurusFormatter = append(gurusFormatter, guruFormatter)
	}

	return gurusFormatter
}
