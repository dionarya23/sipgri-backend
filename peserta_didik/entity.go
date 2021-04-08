package peserta_didik

type PesertaDidik struct {
	Nisn                string `gorm:"primaryKey"`
	Nis                 string
	IDKelas             int
	Nama                string
	JenisKelamin        string
	TempatLahir         string
	TanggalLahir        string
	StatusDalamKeluarga string
	AnakKe              int
	Alamat              string
	NomorTelepon        string
	SekolahAsal         string
	TanggalDiterima     string
	NamaAyah            string
	NamaIbu             string
	NamaWali            string
	NomorTeleponWali    string
}
