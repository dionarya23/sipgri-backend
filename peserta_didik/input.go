package peserta_didik

type InputDataPesertaDidik struct {
	Nisn                string `json:"nisn" binding:"required"`
	Nis                 string `json:"nis" binding:"required"`
	IDKelas             int    `json:"id_kelas" binding:"required"`
	Nama                string `json:"nama" binding:"required"`
	JenisKelamin        string `json:"jenis_kelamin" binding:"required"`
	TempatLahir         string `json:"tempat_lahir" binding:"required"`
	TanggalLahir        string `json:"tanggal_lahir" binding:"required"`
	StatusDalamKeluarga string `json:"status_dalam_keluarga" binding:"required"`
	AnakKe              int    `json:"anak_ke" binding:"required"`
	Alamat              string `json:"alamat" binding:"required"`
	NomorTelepon        string `json:"nomor_telepon" binding:"required"`
	SekolahAsal         string `json:"sekolah_asal" binding:"required"`
	TanggalDiterima     string `json:"tanggal_diterima" binding:"required"`
	NamaAyah            string `json:"nama_ayah" binding:"required"`
	NamaIbu             string `json:"nama_ibu" binding:"required"`
	NamaWali            string `json:"nama_wali" binding:"required"`
	NomorTeleponWali    string `json:"nomor_telepon_wali" binding:"required"`
}

type InputNisn struct {
	Nisn string `uri:"nisn" binding:"required"`
}
