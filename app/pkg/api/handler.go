package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	rt "github.com/joshiomkarj/inMemoryServer/app/pkg/runtime"
)

// Api response struct
type apiResponse struct {
	Status  int      `json:"status"`
	Message []Server `json:"servers"`
}

type Server struct {
	ID     string `json:"id"`
	VMName string `json:"vmname"`
	VMID   string `json:"vmid"`
	CPU    string `json:"cpuutilization"`
}

// Send Api response function
func sendApiResponse(w http.ResponseWriter, msg []Server, respCode int) {

	// Set content type
	w.Header().Set("Content-type", "application/json")

	// set response code
	w.WriteHeader(respCode)

	// Create response body
	response, err := json.Marshal(apiResponse{Status: respCode, Message: msg})
	if err != nil {
		log.Printf("Failed to create api response. Error: '%s'", err)
		return
	}

	// set response body
	w.Write(response)
}

// GetServers returns a list of vms
func GetServers(w http.ResponseWriter, r *http.Request) {

	log.Printf("GetServers")

	req := &rt.RegisterRequest{}
	defer r.Body.Close()

	log.Printf("Request body is received: '%s'", req)
	sendApiResponse(w, VMList, http.StatusOK)
}

// GetServer returns a server based on id
func GetServer(w http.ResponseWriter, r *http.Request) {

	log.Printf("GetServers")
	vars := mux.Vars(r)
	req := &rt.RegisterRequest{}
	defer r.Body.Close()

	log.Println("id of the VM you're looking for is: ", vars["id"])
	for _, vm := range VMList {
		if vm.ID == vars["id"] {
			log.Printf("The VM you're looking for is: %+v", vm)
		}
	}

	log.Printf("Request body is received: '%s'", req)
	sendApiResponse(w, VMList, http.StatusOK)
}
