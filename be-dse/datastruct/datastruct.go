package datastruct

import "be-dse/config"

// import (
// 	"database/sql"
// )

type Datasets struct {
	Kode_data            int64             `json:"kode_data"`
	Topik                string            `json:"topik"`
	Grup                 string            `json:"grup"`
	Judul                string            `json:"judul"`
	TahunAwalData        config.NullInt64  `json:"tahun_awal_data"`
	TahunAkhirData       config.NullInt64  `json:"tahun_akhir_data"`
	SumberData           string            `json:"sumber_data"`
	TautanSumberData     string            `json:"tautan_sumber_data"`
	TautanDatasetTerkait string            `json:"tautan_dataset_terkait"`
	OrganisasiTerkait    config.NullString `json:"organisasi_terkait"`
	FrekuensiPenerbitan  config.NullString `json:"frekuensi_penerbitan"`
	LastUpdated          config.NullString `json:"last_updated"`
	JenisData            string            `json:"jenis_data"`
}

type ResponsePost struct {
	Kode_data            int64  `json:"kode_data,omitempty"`
	Topik                string `json:"topik,omitempty"`
	Grup                 string `json:"grup,omitempty"`
	Judul                string `json:"judul,omitempty"`
	TahunAwalData        int64  `json:"tahun_awal_data"`
	TahunAkhirData       int64  `json:"tahun_akhir_data"`
	SumberData           string `json:"sumber_data"`
	TautanSumberData     string `json:"tautan_sumber_data"`
	TautanDatasetTerkait string `json:"tautan_dataset_terkait"`
	OrganisasiTerkait    string `json:"organisasi_terkait"`
	FrekuensiPenerbitan  string `json:"frekuensi_penerbitan"`
	LastUpdated          string `json:"last_updated"`
	JenisData            string `json:"jenis_data"`
	Message              string `json:"message,omitempty"`
}

type ResponseUpdate struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseDelete struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseGetAll struct {
<<<<<<< HEAD
	Status    int        `json:"status"`
	TotalData int        `json:"total_data"`
	TotalPage int        `json:"total_page"`
	Message   string     `json:"message"`
	Data      []Datasets `json:"data"`
}

type ResponseGetTotal struct {
	Status    int    `json:"status"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Message   string `json:"message"`
=======
	Status  int        `json:"status"`
	TotalData int     `json:"total data"`
	TotalPage int     `json:"total page"`
	Message string     `json:"message"`
	Data    []Datasets `json:"data"`
>>>>>>> 7e99f55408fe3d9b4f6ea53e7ae5f36e43f1c6f0
}

type ResponseGet struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Datasets `json:"data"`
}

type Response5 struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	LikeStatus bool   `json:"like_status"`
}
