package handles

import (
	"api/models"
	"encoding/json"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request){
	page, err := os.ReadFile("../doc/index.html")

	if err != nil {
		response := models.Response {Message: "failed to load page"}
		w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(response)
	}else{
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(string(page)))
	}

}