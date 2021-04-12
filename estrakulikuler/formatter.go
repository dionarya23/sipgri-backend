package estrakulikuler

import "github.com/dionarya23/sipgri-backend/guru"

type EskulDetailFormatter struct {
	IDEstrakulikuler int         `json:"id_estrakulikuler"`
	Jenis            string      `json:"jenis"`
	NipPembimbing    string      `json:"nip_pembimbing"`
	Pembimbing       interface{} `json:"pembimbing"`
}

func FormatEskulDetail(eskul Estrakulikuler) EskulDetailFormatter {
	formatter := EskulDetailFormatter{}
	formatter.IDEstrakulikuler = eskul.IDEstrakulikuler
	formatter.Jenis = eskul.Jenis
	formatter.NipPembimbing = eskul.NipPembimbing

	if eskul.Pembimbing.Nip != "" {
		formatter.Pembimbing = guru.FormatDetailGuru(eskul.Pembimbing)
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
