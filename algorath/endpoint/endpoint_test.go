package endpoint

import (
	"algorath/algorath"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type HttpClientMock struct {
	mock.Mock
}

func (c *HttpClientMock) Get(url string) (resp *http.Response, err error) {
	args := c.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func (c *HttpClientMock) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	args := c.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func (c *HttpClientMock) Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	args := c.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

type databaseMock struct {
	mock.Mock
}

func (d databaseMock) GetCredential() (algorath.Credentials, error) {
	args := d.Called()
	return args.Get(0).(algorath.Credentials), args.Error(1)
}

func (d databaseMock) UpdateCredential(algorath.Credentials) error {
	args := d.Called()
	return args.Error(1)
}

type managerMock struct {
	mock.Mock
}

func (m managerMock) Launch() (error) {
	args := m.Called()
	return args.Error(1)
}

var api *mux.Router
var databaseMocked *databaseMock
var managerMocked *managerMock
var controller *Controller

func TestMain(m *testing.M){

	databaseMocked = new(databaseMock)
	controller = New(databaseMocked, managerMocked)

	routes := controller.Routes()

	api = &mux.Router{}
	controller.Routes()

	for _, route := range routes {
		api.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestControllerGetCredentials(t *testing.T){

	databaseMocked.On("GetCredential").Return(algorath.Credentials{
		APIKey:    "test",
		APISecret: "test",
	},nil).Once()


	req, err := http.NewRequest("GET", "/crendentials", nil)


	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	api.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("TestControllerGetCredentials returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

}

func TestControllerLaunch(t *testing.T){

	managerMocked.On("Launch").Return(nil).Once()


	req, err := http.NewRequest("GET", "/start", nil)


	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	api.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("TestControllerLaunch returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

}