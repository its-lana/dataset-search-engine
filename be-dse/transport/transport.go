package transport

import (
	"be-dse/datastruct"
	"be-dse/logging"
	"be-dse/service"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"


	"github.com/gorilla/mux"

)



func Insert_Dataset(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acce ss-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var datasets datastruct.Datasets
	

	err := json.NewDecoder(r.Body).Decode(&datasets)
	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}
	

	datasets.LastUpdated.String = logging.GetDateTimeNowInString()
	service.Create(datasets)

	logging.Log(fmt.Sprintf("Datasets dengan id %d berhasil ditambahkan!", datasets.Kode_data))

	res := datastruct.ResponsePost{
		Kode_data: datasets.Kode_data,
		Topik: datasets.Topik,
		Grup : datasets.Grup,
		Judul: datasets.Judul,
		TahunAwalData: datasets.TahunAwalData.Int64,
		TahunAkhirData: datasets.TahunAkhirData.Int64,
		SumberData: datasets.SumberData,
		TautanSumberData: datasets.TautanSumberData,
		TautanDatasetTerkait: datasets.TautanDatasetTerkait,
		OrganisasiTerkait: datasets.OrganisasiTerkait.String,
		FrekuensiPenerbitan: datasets.FrekuensiPenerbitan.String,
		LastUpdated: datasets.LastUpdated.String,
		JenisData: datasets.JenisData,
		Message:  "Berhasil menambahkan dataset!",
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}


func Update_Dataset(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)

	id_, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	var datasets datastruct.Datasets
	
	

	err = json.NewDecoder(r.Body).Decode(&datasets)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	datasets.LastUpdated.String = logging.GetDateTimeNowInString()
	updatedRows := service.Update(int64(id_), datasets)

	msg := fmt.Sprintf("Dataset telah berhasil diupdate. Status yang diupdate %v rows/record", updatedRows)

	res := datastruct.ResponseUpdate{
		Status: 200,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}


func Delete_Dataset(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)

	id_, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	var datasets datastruct.Datasets

	err = json.NewDecoder(r.Body).Decode(&datasets)

	deletedRows := service.Delete(int64(id_))

	msg := fmt.Sprintf("Dataset berhasil dihapus. Status yang dihapus %v rows/record", deletedRows)

	res := datastruct.ResponseDelete{
		Status: 200,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}


func Get_All_Datasets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	searchQuery := r.URL.Query().Get("search")
	

	 sortBy := r.URL.Query().Get("sortBy")
	 if sortBy == "" {
		 // id.asc is the default sort query
		 sortBy = "judul.asc"
	 }
	 sortQuery, err := logging.ValidateAndReturnSortQuery(sortBy)
	 if err != nil {
		 http.Error(w, err.Error(), http.StatusBadRequest)
		 return
	 }

    strPage := r.URL.Query().Get("page")
    page := 1
    if strPage != "" {
        page, err = strconv.Atoi(strPage)
        if err != nil || page < 1 {
            http.Error(w, "page query parameter is no valid number", http.StatusBadRequest)
            return
        }
    }

	filter := r.URL.Query().Get("filter")
	filterQuery := ""
    if filter != "" {
        filterQuery, err = logging.ValidateAndReturnFilterMap(filter)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    }

	datasets, err := service.Get_All(searchQuery, sortQuery, filterQuery, page)

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response datastruct.ResponseGetAll
	response.Status = 200
	response.Message = "Success"
	response.Data = datasets

	json.NewEncoder(w).Encode(response)
}







func Get_Dataset(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	data_set := service.Get_Data(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}


	logging.Log(fmt.Sprintf("%d menampilkan data like", id))

	response:= datastruct.ResponseGet{
		Status: 200,
		Message: "Success",
		Data: data_set,
	}
	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

// func TmplknSpecificLike(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	params := mux.Vars(r)

// 	user_id, err := strconv.Atoi(params["user_id"])

// 	if err != nil {
// 		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
// 	}

// 	feed_id, err := strconv.Atoi(params["feed_id"])

// 	if err != nil {
// 		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
// 	}

// 	SpecificLike  := service.ShowSpecificLike(int64(user_id),int64(feed_id))

// 	if err != nil {
// 		log.Fatalf("Tidak bisa mengambil data. %v", err)
// 	}



// 	response5:= datastruct.Response5{
// 		Status: 200,
// 		Message: "Success",
// 		LikeStatus: SpecificLike,
// 	}
	
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(response5)
// }
