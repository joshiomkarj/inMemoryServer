package api

import (
	"encoding/json"
	"fmt"
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

// CreateServer returns a server based on id
func CreateServer(w http.ResponseWriter, r *http.Request) {

	log.Printf("Create Server")
	req := &rt.RegisterRequest{}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		msg := fmt.Sprintf("Failed to decode request body into JSON. Error: %s", err)
		log.Printf("%s", msg)
		return
	}

	log.Printf("Request body is received: '%+v'", req)
	var vm = Server{
		ID:     req.ID,
		VMName: req.VMName,
		VMID:   req.VMID,
		CPU:    req.CPU,
	}
	VMList = append(VMList, vm)
	sendApiResponse(w, VMList, http.StatusOK)
}

// DeleteServer returns a server based on id
func DeleteServer(w http.ResponseWriter, r *http.Request) {

	log.Printf("Delete Server")
	vars := mux.Vars(r)
	req := &rt.RegisterRequest{}
	defer r.Body.Close()

	var idx int
	for i, vm := range VMList {
		if vm.ID == vars["id"] {
			log.Printf("The VM you're looking to delete is: %+v", vm)
			idx = i
		}
	}

	log.Printf("VMList: '%+v'", VMList)
	VMList = append(VMList[:idx], VMList[idx+1:]...)
	log.Printf("VMList: '%+v'", VMList)
	log.Printf("Request body is received: '%s'", req)
	sendApiResponse(w, VMList, http.StatusOK)
}

// PatchServer returns a server based on id
func PatchServer(w http.ResponseWriter, r *http.Request) {

	// Using PATCH instead of PUT because PUT requires an entire body to be sent
	// effectively making PUT a subset of PATCH
	log.Printf("Updates the vm")
	vars := mux.Vars(r)
	req := &rt.RegisterRequest{}
	defer r.Body.Close()

	var idx int
	for i, vm := range VMList {
		if vm.ID == vars["id"] {
			log.Printf("The VM you're looking to update is: %+v", vm)
			idx = i
		}
	}

	log.Printf("r.Body: '%+v'", r.Body)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		msg := fmt.Sprintf("Failed to decode request body into JSON. Error: %s", err)
		log.Printf("%s", msg)
		return
	}

	vm := VMList[idx]
	log.Printf("req: '%+v'", req)
	log.Printf("VM: '%+v'", vm)
	if req.VMName != "" {
		vm.VMName = req.VMName
	}
	if req.CPU != "" {
		vm.CPU = req.CPU
	}

	VMList[idx] = vm

	log.Printf("Request body is received: '%+v'", req)
	sendApiResponse(w, VMList, http.StatusOK)
}
