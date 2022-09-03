package service

import (
	"be-dse/config"
	"be-dse/datastruct"
	"fmt"
	"log"
)

func Create(datasets datastruct.Datasets) int64 {

	db := config.CreateConnection()

	//defer db.Close()


	sqlStatement := `INSERT INTO datasets (kode_data, topik, grup, judul, 
					tahun_awal_data, tahun_akhir_data, sumber_data, tautan_sumber_data, tautan_dataset_terkait,
					organisasi_terkait, frekuensi_penerbitan, last_updated, jenis_data) VALUES ($1, $2, $3, $4,
					$5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING kode_data`

	err := db.QueryRow(sqlStatement, datasets.Kode_data, datasets.Topik, datasets.Grup, datasets.Judul,
							datasets.TahunAwalData.Int64, datasets.TahunAkhirData.Int64, datasets.SumberData,
							datasets.TautanSumberData, datasets.TautanDatasetTerkait,
							datasets.OrganisasiTerkait.String, datasets.FrekuensiPenerbitan.String, datasets.LastUpdated.String,
							datasets.JenisData).Scan(&datasets.Kode_data)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	return datasets.Kode_data
}

func Update(id int64, datasets datastruct.Datasets) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE datasets SET topik=$2, grup=$3, judul=$4,
					tahun_awal_data=$5, tahun_akhir_data=$6, sumber_data= $7,
					tautan_sumber_data=$8, tautan_dataset_terkait=$9, organisasi_terkait=$10,
					frekuensi_penerbitan=$11, last_updated=$12, jenis_data=$13 WHERE kode_data=$1`

	res, err := db.Exec(sqlStatement, id, datasets.Topik, datasets.Grup, datasets.Judul,
						datasets.TahunAwalData.Int64, datasets.TahunAkhirData.Int64,
						datasets.SumberData, datasets.TautanSumberData, datasets.TautanDatasetTerkait,
						datasets.OrganisasiTerkait.String, datasets.FrekuensiPenerbitan.String,
						datasets.LastUpdated.String, datasets.JenisData)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	return rowsAffected
}



func Delete(id int64) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM datasets WHERE kode_data=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang dihapus. %v", err)
	}

	return rowsAffected
}

func Get_All(search string, sort string, filter string, page int) ([]datastruct.Datasets, error) {
	db := config.CreateConnection()


	defer db.Close()

	var datasets []datastruct.Datasets
	perPage := 9


	sqlStatement := `SELECT * FROM datasets`

	if len(search) != 0{
		sqlStatement = fmt.Sprintf("%s WHERE judul LIKE '%%%s%%'" , sqlStatement, search)

		if len(filter) != 0 {
			sqlStatement = fmt.Sprintf("%s AND %s" , sqlStatement, filter)
		}
	} else {
			if len(filter) != 0 {
				sqlStatement = fmt.Sprintf("%s WHERE %s" , sqlStatement, filter)
			}
	} 

	if len(sort) != 0 {
		sqlStatement = fmt.Sprintf("%s ORDER BY  %s" , sqlStatement, sort)
	}

	if page >= 1{
		sqlStatement = fmt.Sprintf("%s LIMIT %d OFFSET %d" , sqlStatement, perPage, (page - 1) * perPage)
	}

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("sql : %s tidak bisa mengeksekusi query. %v", sqlStatement, err)
	}

	defer rows.Close()

	for rows.Next() {
		var data datastruct.Datasets

		err = rows.Scan(&data.Kode_data, &data.Topik, &data.Grup, &data.Judul, &data.TahunAwalData, &data.TahunAkhirData,
						&data.SumberData, &data.TautanSumberData, &data.TautanDatasetTerkait,
						&data.OrganisasiTerkait, &data.FrekuensiPenerbitan, &data.LastUpdated, &data.JenisData)

		
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		datasets = append(datasets, data)
	}
	return datasets, err
}


func Get_Data(id int64) (datastruct.Datasets) {
	db := config.CreateConnection()

	defer db.Close()

	var datasets datastruct.Datasets

	sqlStatement := `SELECT * FROM datasets WHERE kode_data=$1`

	err := db.QueryRow(sqlStatement, id).Scan(&datasets.Kode_data, &datasets.Topik, &datasets.Grup,
		&datasets.Judul, &datasets.TahunAwalData, &datasets.TahunAkhirData,
		&datasets.SumberData, &datasets.TautanSumberData, &datasets.TautanDatasetTerkait, 
		&datasets.OrganisasiTerkait, &datasets.FrekuensiPenerbitan, &datasets.LastUpdated, &datasets.JenisData)

	

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	return datasets
}


// func ShowLike(id int64) (int64) {
// 	db := config.CreateConnection()

// 	defer db.Close()

// 	var sqlStatement string
// 	var count int64

// 	sqlStatement = `SELECT COUNT(*) FROM menyukai WHERE feed_id=$1 AND status_like IS TRUE`


// 	err := db.QueryRow(sqlStatement, id).Scan(&count)

// 	if err != nil {
// 		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
// 	}

// 	return count
// }


// func ShowSpecificLike(user_id int64,feed_id int64)(bool) {
// 	db := config.CreateConnection()

// 	defer db.Close()

// 	var sqlStatement string
// 	var status_like bool

// 	sqlStatement = `SELECT status_like FROM menyukai WHERE user_id= $1 AND feed_id=$2`


// 	err := db.QueryRow(sqlStatement, user_id,feed_id).Scan(&status_like)

// 	if err != nil {
// 		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
// 	}

// 	return status_like
// }

