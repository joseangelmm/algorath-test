package endpoint

import (
	"algorath/algorath"
	"algorath/manager"
	"algorath/repository"
	"encoding/json"
	"net/http"
)

type Controller struct {
	db repository.DatabaseI
	manager manager.ManagerI
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(db repository.DatabaseI, manager manager.ManagerI) *Controller {

	newController := new(Controller)

	newController.db = db
	newController.manager = manager

	return newController
}

func (c *Controller) Routes() []Route {
	return []Route{
		{
			Path:    "/crendentials",
			Handler: c.GetCredentials,
			Method:  http.MethodGet,
		},
		{
			Path:    "/start",
			Handler: c.StartProcedure,
			Method:  http.MethodGet,
		},
		{
			Path:    "/crendentials",
			Handler: c.SetCredentials,
			Method:  http.MethodPut,
		},
	}
}

func (c *Controller) GetCredentials(w http.ResponseWriter, _ *http.Request) {

	cred, err := c.db.GetCredential()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		error := Error{500, "Internal Server Error -" + err.Error()}
		_ = json.NewEncoder(w).Encode(error)

	} else {

		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(cred)
	}

}


func (c *Controller) StartProcedure(w http.ResponseWriter, _ *http.Request) {

	err := c.manager.Launch()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		error := Error{500, "Internal Server Error -" + err.Error()}
		_ = json.NewEncoder(w).Encode(error)

	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

func (c *Controller) SetCredentials(w http.ResponseWriter, r *http.Request) {

	var res algorath.Credentials

	if r.Body == nil {

		w.WriteHeader(http.StatusBadRequest)

		error := Error{400, "Bad Request - The body is empty"}

		_ = json.NewEncoder(w).Encode(error)

		return
	}

	json.NewDecoder(r.Body).Decode(&res)

	err := c.db.UpdateCredential(res)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		error := Error{500, "Internal Server Error -" + errint.Error()}
		_ = json.NewEncoder(w).Encode(error)

		return
	} else {

		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(res)

	}

}
