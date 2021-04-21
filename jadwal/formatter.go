package jadwal

type JadwalFormatter struct {
	IDJadwal     int         `json:"id_jadwal"`
	KodeMengajar string      `json:"kode_mengajar"`
	IDKelas      int         `json:"id_kelas"`
	JamMulai     string      `json:"jam_mulai"`
	JamSelesai   string      `json:"jam_selesai"`
	Hari         string      `json:"hari"`
	Kelas        interface{} `json:"kelas"`
	Pengajar     interface{} `json:"pengajar"`
}

func FormatJadwalDetail(jadwal Jadwal) JadwalFormatter {
	formatter := JadwalFormatter{}

	formatter.IDJadwal = jadwal.IDJadwal
	formatter.KodeMengajar = jadwal.KodeMengajar
	formatter.IDKelas = jadwal.IDKelas
	formatter.JamMulai = jadwal.JamMulai
	formatter.JamSelesai = jadwal.JamSelesai
	formatter.Hari = jadwal.Hari

	if jadwal.Kelas.IDKelas != 0 {
		formatter.Kelas = jadwal.Kelas
	}

	if jadwal.Pengajar.KodeMengajar != "" {
		formatter.Pengajar = jadwal.Pengajar
	}

	return formatter
}

func FormatJadwalList(jadwal []Jadwal) []JadwalFormatter {
	formatterList := []JadwalFormatter{}

	for _, value := range jadwal {
		formatter := FormatJadwalDetail(value)
		formatterList = append(formatterList, formatter)
	}

	return formatterList
}
