package handles

import (
	DB "api/db"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mrz1836/go-sanitize"
)

// Handler function for the postjob endpoint
func AddPostal(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    postal := models.PostalCode{
        Id: uuid.NewString(),
        Area: sanitize.AlphaNumeric(sanitize.XSS(r.FormValue("area")),true),
        Postal_code: sanitize.XSS(r.FormValue("postal-code")),
        District: sanitize.XSS(r.FormValue("district")),
        Region: sanitize.XSS(r.FormValue("region")),
        Settlement: sanitize.AlphaNumeric(sanitize.XSS(r.FormValue("settlement")),true),
    }

    // initialize a db connection
    db := DB.Connect()

    result := db.Create(&postal)

    if result.Error == nil{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
    }else{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
    }

}

// Handler function for the getjob endpoint
func GetPostal(w http.ResponseWriter, r *http.Request){
    // initialize a db connection
    db := DB.Connect()

    postal := []models.PostalCode{}
    fields := `
        postal_codes.id AS id,
        area,  
        code AS Postal_code, 
        district.name AS district,
        region.name AS region, 
        location.type AS settlement
    `
    result := db.Select(fields).
                 Joins("INNER JOIN location ON postal_codes.location_type_id = location.id").
                 Joins("INNER JOIN district ON postal_codes.district_id = district.id").
                 Joins("INNER JOIN region ON district.region_id = region.id ").Find(&postal)

    res, err := json.Marshal(postal)

    if err != nil {
        log.Fatal(err)
    }

    if result.Error == nil{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusAccepted)
        w.Write(res)
        
    }else{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
    }
}

func GetPostalBySlug(w http.ResponseWriter, r *http.Request){
    // initialize a db connection
    db := DB.Connect()

    slug := sanitize.AlphaNumeric(sanitize.XSS(chi.URLParam(r,"slug")),true)
    
    postal := []models.PostalCode{}

    sql := `SELECT postal_codes.id AS id,
        area,  
        code AS Postal_code, 
        district.name AS district,
        region.name AS region, 
        location.type AS settlement 
        FROM postal_codes
        INNER JOIN location ON postal_codes.location_type_id = location.id
        INNER JOIN district ON postal_codes.district_id = district.id
        INNER JOIN region ON district.region_id = region.id `

    sql = fmt.Sprintf("%s WHERE area LIKE '%%%s%%' OR district.name LIKE '%%%s%%' OR region.name LIKE '%%%s%%' OR code LIKE '%%%s%%' OR location.type LIKE '%%%s%%'",sql,slug,slug,slug,slug,slug)

    result := db.Raw(sql).Scan(&postal)

    res, err := json.Marshal(postal)

    if err != nil {
        log.Fatal(err)
    }

    if result.Error == nil{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusAccepted)
        w.Write(res)
        
    }else{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
    }
}