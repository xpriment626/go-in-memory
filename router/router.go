package router

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"memory-crud/models"
	"net/http"
	"strconv"

	"memory-crud/utils"

	"github.com/gorilla/mux"
)

func getTenants(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Tenants)
}
func deleteTenant(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, tenant := range models.Tenants {
        if tenant.ID == params["id"] {
            models.Tenants = append(models.Tenants[:index], models.Tenants[index+1:]...)
            break
        }
    }
}
func postTenant(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var tenant models.Tenant
    _ = json.NewDecoder(r.Body).Decode(&tenant)
    tenant.ID = strconv.Itoa(rand.Intn(9999999))
    models.Tenants = append(models.Tenants, tenant)
    json.NewEncoder(w).Encode(tenant)
}
func updateTenant(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, tenant := range models.Tenants {
        if tenant.ID == params["id"] {
            models.Tenants = append(models.Tenants[:index], models.Tenants[index + 1:]...)
            var tenant models.Tenant
            _ = json.NewDecoder(r.Body).Decode(&tenant)
            tenant.ID = strconv.Itoa(rand.Intn(9999999))
            models.Tenants = append(models.Tenants, tenant)
            json.NewEncoder(w).Encode(tenant)
            return
        }
    }
}
func getTenant(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, tenant := range models.Tenants {
        if tenant.ID == params["id"] {
            json.NewEncoder(w).Encode(tenant)
            return
        }
    }
}
func UseRouter() {
    utils.Populate()
	r := mux.NewRouter()
    r.HandleFunc("/tenants", getTenants).Methods("GET")
    r.HandleFunc("/tenants/{id}", getTenant).Methods("GET")
    r.HandleFunc("/tenants", postTenant).Methods("POST")
    r.HandleFunc("/tenants/{id}", updateTenant).Methods("PUT")
    r.HandleFunc("/tenants/{id}", deleteTenant).Methods("DELETE")
    fmt.Println("Starting server at port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
