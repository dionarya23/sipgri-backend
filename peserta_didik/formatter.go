package peserta_didik

type PesertaDidikFormatter struct {
	Nisn                string `json:"nisn"`
	Nis                 string `json:"nis"`
	Nama                string `json:"nama"`
	JenisKelamin        string `json:"jenis_kelamin"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	StatusDalamKeluarga string `json:"status_dalam_keluarga"`
	AnakKe              int    `json:"anak_ke"`
	Alamat              string `json:"alamat"`
	NomorTelepon        string `json:"nomor_telepon"`
	SekolahAsal         string `json:"sekolah_asal"`
	TanggalDiterima     string `json:"tanggal_diterima"`
	NamaAyah            string `json:"nama_ayah"`
	NamaIbu             string `json:"nama_ibu"`
	NamaWali            string `json:"nama_wali"`
	NomorTeleponWali    string `json:"nomor_telepon_wali"`
}

func FormatPesertaDidikDetail(persetaDidik PesertaDidik) PesertaDidikFormatter {
	peserta_didik := PesertaDidikFormatter{}
	peserta_didik.Nisn = persetaDidik.Nisn
	peserta_didik.Nis = persetaDidik.Nis
	peserta_didik.Nama = persetaDidik.Nama
	peserta_didik.JenisKelamin = persetaDidik.JenisKelamin
	peserta_didik.TempatLahir = persetaDidik.TempatLahir
	peserta_didik.TanggalLahir = persetaDidik.TanggalLahir
	peserta_didik.StatusDalamKeluarga = persetaDidik.StatusDalamKeluarga
	peserta_didik.AnakKe = persetaDidik.AnakKe
	peserta_didik.Alamat = persetaDidik.Alamat
	peserta_didik.NomorTelepon = persetaDidik.NomorTelepon
	peserta_didik.SekolahAsal = persetaDidik.SekolahAsal
	peserta_didik.TanggalDiterima = persetaDidik.TanggalDiterima
	peserta_didik.NamaAyah = persetaDidik.NamaAyah
	peserta_didik.NamaIbu = persetaDidik.NamaIbu
	peserta_didik.NamaWali = persetaDidik.NamaWali
	peserta_didik.NomorTeleponWali = persetaDidik.NomorTeleponWali

	return peserta_didik
}

func FormatPesertaDidikList(pesertaDidik []PesertaDidik) []PesertaDidikFormatter {
	peserta_didik_ := []PesertaDidikFormatter{}

	for _, perseta := range pesertaDidik {
		pesertaDidik_ := FormatPesertaDidikDetail(perseta)
		peserta_didik_ = append(peserta_didik_, pesertaDidik_)
	}

	return peserta_didik_
}
