package estrakulikuler

type EskulDetailFormatter struct {
	IDEstrakulikuler int         `json:"id_estrakulikuler"`
	Jenis            string      `json:"jenis"`
	NipGuru          string      `json:"nip_guru"`
	Pembimbing       interface{} `json:"pembimbing"`
}

func FormatEskulDetail(eskul Estrakulikuler) EskulDetailFormatter {
	formatter := EskulDetailFormatter{}
	formatter.IDEstrakulikuler = eskul.IDEstrakulikuler
	formatter.Jenis = eskul.Jenis
	formatter.NipGuru = eskul.NipGuru

	if eskul.Pembimbing.Nip != "" {
		formatter.Pembimbing = eskul.Pembimbing
	}

	return formatter

}

func FormatEskulList(eskul []Estrakulikuler) []EskulDetailFormatter {
	listFormatter := []EskulDetailFormatter{}

	for _, value := range eskul {
		detailEskul := FormatEskulDetail(value)
		listFormatter = append(listFormatter, detailEskul)
	}

	return listFormatter
}
