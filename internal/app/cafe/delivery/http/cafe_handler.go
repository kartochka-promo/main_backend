package http

import (
	"2020_1_drop_table/configs"
	"2020_1_drop_table/internal/app"
	"2020_1_drop_table/internal/app/cafe"
	"2020_1_drop_table/internal/app/cafe/models"
	globalModels "2020_1_drop_table/internal/app/models"
	"2020_1_drop_table/internal/pkg/permissions"
	"2020_1_drop_table/internal/pkg/responses"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CafeHandler struct {
	CUsecase cafe.Usecase
}

func NewCafeHandler(r *mux.Router, us cafe.Usecase) {
	handler := CafeHandler{
		CUsecase: us,
	}

	r.HandleFunc("/api/v1/cafe", permissions.CheckCSRF(permissions.CheckAuthenticated(handler.AddCafeHandler))).Methods("POST")
	r.HandleFunc("/api/v1/cafe", permissions.SetCSRF(handler.GetByOwnerIDHandler)).Methods("GET")
	r.HandleFunc("/api/v1/cafe/{id:[0-9]+}", permissions.SetCSRF(handler.GetByIDHandler)).Methods("GET")
	r.HandleFunc("/api/v1/cafe/{id:[0-9]+}", permissions.CheckCSRF(permissions.CheckAuthenticated(handler.EditCafeHandler))).Methods("PUT")
	r.HandleFunc("/api/v1/cafe/get_all", permissions.SetCSRF(handler.GetAllCafes)).Methods("GET")
	r.HandleFunc("/api/v1/cafe/get_by_geo", permissions.SetCSRF(handler.GetCafeListByGeoAndRadius)).Methods("GET")
	r.HandleFunc("/api/v1/cafe/with_pass/{id:[0-9]+}", permissions.SetCSRF(handler.GetByIDWithPassInfoHandler)).Methods("GET")
}

func (c *CafeHandler) fetchCafe(r *http.Request) (models.Cafe, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return models.Cafe{}, globalModels.ErrBadRequest
	}

	jsonData := r.FormValue("jsonData")
	if jsonData == "" || jsonData == "null" {
		return models.Cafe{}, globalModels.ErrEmptyJSON
	}

	cafeObj := models.Cafe{}
	err = cafeObj.UnmarshalJSON([]byte(jsonData))
	if err != nil {
		return models.Cafe{}, globalModels.ErrBadJSON
	}
	if file, handler, err := r.FormFile("photo"); err == nil {
		filename, err := app.SaveFile(file, handler, "cafe")
		if err == nil {
			cafeObj.Photo = fmt.Sprintf("%s/%s", configs.ServerUrl, filename)
		}
	}

	return cafeObj, nil
}

func (c *CafeHandler) GetCafeListByGeoAndRadius(w http.ResponseWriter, r *http.Request) {
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	radius := r.FormValue("radius")

	res, err := c.CUsecase.GetCafeSortedByRadius(r.Context(), latitude, longitude, radius)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}
	responses.SendOKAnswer(res, w)
}

func (c *CafeHandler) AddCafeHandler(w http.ResponseWriter, r *http.Request) {
	cafeObj, err := c.fetchCafe(r)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	cafeObjWithGeo, err := c.CUsecase.Add(r.Context(), cafeObj)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	responses.SendOKAnswer(cafeObjWithGeo, w)
}

func (c *CafeHandler) EditCafeHandler(w http.ResponseWriter, r *http.Request) {
	cafeObj, err := c.fetchCafe(r)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		message := fmt.Sprintf("bad id: %s", mux.Vars(r)["id"])
		responses.SendSingleError(message, w)
		return
	}

	cafeObj.CafeID = id

	cafeDB, err := c.CUsecase.Update(r.Context(), cafeObj)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	responses.SendOKAnswer(cafeDB, w)
}

func (c *CafeHandler) GetByOwnerIDHandler(w http.ResponseWriter, r *http.Request) {
	cafesObj, err := c.CUsecase.GetByOwnerID(r.Context())
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	responses.SendOKAnswer(cafesObj, w)
}

func (c *CafeHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		message := fmt.Sprintf("bad id: %s", mux.Vars(r)["id"])
		responses.SendSingleError(message, w)
		return
	}
	cafeObj, err := c.CUsecase.GetByID(r.Context(), id)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}

	responses.SendOKAnswer(cafeObj, w)
}

func (c *CafeHandler) GetByIDWithPassInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		message := fmt.Sprintf("bad id: %s", mux.Vars(r)["id"])
		responses.SendSingleError(message, w)
		return
	}
	cafeObj, err := c.CUsecase.GetByIDWithPassInfo(r.Context(), id)
	if err != nil {
		responses.SendSingleError(err.Error(), w)
		return
	}
	responses.SendOKAnswer(cafeObj, w)
}

func (c *CafeHandler) GetAllCafes(writer http.ResponseWriter, request *http.Request) {
	limit, err := strconv.Atoi(request.FormValue("limit"))
	since, err2 := strconv.Atoi(request.FormValue("since"))
	search := request.FormValue("searchBy")
	if err != nil || err2 != nil {
		responses.SendSingleError("Bad GET params", writer)
		return
	}
	cafes, err := c.CUsecase.GetAllCafes(request.Context(), since, limit, search)
	if err != nil {
		responses.SendSingleError(err.Error(), writer)
		return
	}
	responses.SendOKAnswer(cafes, writer)
}
