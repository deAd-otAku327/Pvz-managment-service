package integration

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"pvz-service/internal/dto"
	"pvz-service/internal/enum"
	"pvz-service/internal/middleware"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dbName *string = flag.String("db", "", "test db container naming")

const AddProductCount = 50

func TestMain(m *testing.M) {
	flag.Parse()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Failed to create docker pool: %v", err)
	}

	postgresResource, err := createPostgresDB(pool, *dbName)
	if err != nil {
		log.Fatalf("Failed to run docker container with postgres: %v", err)
	}

	err = initTestServer(postgresResource.GetPort("5432/tcp"))
	if err != nil {
		log.Fatalf("Failed to init test app: %v", err)
	}

	defer func() {
		err = teardown(pool, postgresResource)
		if err != nil {
			log.Fatalf("Failed to purge resources: %v", err)
		}
	}()

	m.Run()
}

func TestDummyLogin(t *testing.T) {
	type dummyLoginRequest struct {
		Role string `json:"role"`
	}

	for k := range roles {
		requestData := dummyLoginRequest{
			Role: k,
		}

		requestDataJSON, err := json.Marshal(&requestData)
		if err != nil {
			t.Fatalf("Failed to marshall request: %v", err)
		}

		request := httptest.NewRequest("POST", "/dummyLogin", bytes.NewBuffer(requestDataJSON))
		response := httptest.NewRecorder()

		server.Server.Handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		require.Equal(t, 1, len(response.Result().Cookies()))

		cookie := response.Result().Cookies()[0]

		assert.Equal(t, middleware.CookieName, cookie.Name)

		roles[k] = cookie.Value
	}
}

func TestCreatePvz(t *testing.T) {
	requestBody := dto.CreatePvzRequestDTO{
		City: enum.Moscow.String(),
	}

	expectedResponse := dto.PvzResponseDTO{
		ID:   1,
		City: requestBody.City,
	}

	actualResponse := dto.PvzResponseDTO{}

	requestBodyJSON, err := json.Marshal(&requestBody)
	if err != nil {
		t.Fatalf("Failed to marshall request: %v", err)
	}

	request := httptest.NewRequest("POST", "/pvz", bytes.NewBuffer(requestBodyJSON))
	request.Header.Set("Cookie", fmt.Sprintf("%s=%s", middleware.CookieName, roles[enum.Moderator.String()]))

	response := httptest.NewRecorder()

	server.Server.Handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)

	err = json.NewDecoder(response.Result().Body).Decode(&actualResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshall response: %v", err)
	}

	assert.Equal(t, expectedResponse.ID, actualResponse.ID)
	assert.Equal(t, expectedResponse.City, actualResponse.City)
}

func TestCreateReception(t *testing.T) {
	requestBody := dto.CreateReceptionRequestDTO{
		PvzID: 1,
	}

	expectedResponse := dto.ReceptionResponseDTO{
		ID:     1,
		PvzID:  requestBody.PvzID,
		Status: enum.StatusInProgress.String(),
	}

	actualResponse := dto.ReceptionResponseDTO{}

	requestBodyJSON, err := json.Marshal(&requestBody)
	if err != nil {
		t.Fatalf("Failed to marshall request: %v", err)
	}

	request := httptest.NewRequest("POST", "/receptions", bytes.NewBuffer(requestBodyJSON))
	request.Header.Set("Cookie", fmt.Sprintf("%s=%s", middleware.CookieName, roles[enum.Employye.String()]))

	response := httptest.NewRecorder()

	server.Server.Handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)

	err = json.NewDecoder(response.Result().Body).Decode(&actualResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshall response: %v", err)
	}

	assert.Equal(t, expectedResponse.ID, actualResponse.ID)
	assert.Equal(t, expectedResponse.PvzID, actualResponse.PvzID)
	assert.Equal(t, expectedResponse.Status, actualResponse.Status)
}

func TestAddProduct(t *testing.T) {
	requestBody := dto.AddProductRequestDTO{
		Type:  enum.TypeElectrinics.String(),
		PvzID: 1,
	}

	expectedResponse := dto.ProductResponseDTO{
		ID:          1,
		Type:        requestBody.Type,
		ReceptionID: 1,
	}

	for i := 0; i < AddProductCount; i, expectedResponse.ID = i+1, expectedResponse.ID+1 {
		actualResponse := dto.ProductResponseDTO{}

		requestBodyJSON, err := json.Marshal(&requestBody)
		if err != nil {
			t.Fatalf("Failed to marshall request: %v", err)
		}

		request := httptest.NewRequest("POST", "/products", bytes.NewBuffer(requestBodyJSON))
		request.Header.Set("Cookie", fmt.Sprintf("%s=%s", middleware.CookieName, roles[enum.Employye.String()]))

		response := httptest.NewRecorder()

		server.Server.Handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Result().StatusCode)

		err = json.NewDecoder(response.Result().Body).Decode(&actualResponse)
		if err != nil {
			t.Fatalf("Failed to unmarshall response: %v", err)
		}

		assert.Equal(t, expectedResponse.ID, actualResponse.ID)
		assert.Equal(t, expectedResponse.Type, actualResponse.Type)
		assert.Equal(t, expectedResponse.ReceptionID, actualResponse.ReceptionID)
	}
}

func TestCloseReception(t *testing.T) {
	expectedResponse := dto.ReceptionResponseDTO{
		ID:     1,
		PvzID:  1,
		Status: enum.StatusClose.String(),
	}

	actualResponse := dto.ReceptionResponseDTO{}

	request := httptest.NewRequest("POST", fmt.Sprintf("/pvz/%d/close_last_reception", 1), nil)
	request.Header.Set("Cookie", fmt.Sprintf("%s=%s", middleware.CookieName, roles[enum.Employye.String()]))

	response := httptest.NewRecorder()

	server.Server.Handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	err := json.NewDecoder(response.Result().Body).Decode(&actualResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshall response: %v", err)
	}

	assert.Equal(t, expectedResponse.ID, actualResponse.ID)
	assert.Equal(t, expectedResponse.PvzID, actualResponse.PvzID)
	assert.Equal(t, expectedResponse.Status, actualResponse.Status)
}
