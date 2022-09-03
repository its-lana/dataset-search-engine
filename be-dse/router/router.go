package router

import (
	"be-dse/transport"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/insert_dataset", transport.Insert_Dataset).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update_dataset/{id}", transport.Update_Dataset).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete_dataset/{id}", transport.Delete_Dataset).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/get_all_datasets", transport.Get_All_Datasets).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/get_dataset/{id}", transport.Get_Dataset).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/search_by_judul/{judul}", transport.Search_By_Judul).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/filter_by_tahun/{tahun_awal}/{tahun_akhir}", transport.Filter_By_Tahun).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/filter_by_sumber_data/{sumber_data}", transport.Filter_By_Sumber_Data).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/filter_by_group/{group}", transport.Filter_By_Group).Methods("GET", "OPTIONS")
	
	

	return router
}