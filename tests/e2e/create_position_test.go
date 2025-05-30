package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	utils_test "github.com/handarudwiki/payroll-sistem/tests/utils"
)

func TestSuccessCreatePosition(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running TestCreatePosition_Success")

	nameTest := "test position"
	usernameTest := "testPosition"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest, models.RoleAdmin)

	loginRequest := map[string]string{
		"username": usernameTest,
		"password": passwordTest,
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/user/login", loginRequest, "")

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var loginResponse LoginResponse

	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	token := loginResponse.Data.Token

	createPositionRequest := map[string]interface{}{
		"name":        "Software Engineer",
		"base_salary": 5000,
	}

	w, _ = utils_test.MakeRequest(t, app, "POST", "/position/", createPositionRequest, token)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

}

func TestFailedCreatePositionWithoutToken(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running TestCreatePosition_FailedWithoutToken")

	createPositionRequest := map[string]interface{}{
		"name":        "Software Engineer",
		"base_salary": 5000,
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/position/", createPositionRequest, "")

	if w.Code != 401 {
		t.Errorf("Expected status code 401, got %d", w.Code)
	}

	var errorResponse ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if errorResponse.Message != "unauthorized" {
		t.Errorf("Expected error message 'unauthorized', got '%s'", errorResponse.Message)
	}
}

func TestFailedCreatePositionBadRequest(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running TestCreatePosition_Failed Bad Request")

	nameTest := "test position"
	usernameTest := "testPositionBadRequest"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest, models.RoleAdmin)

	loginRequest := map[string]string{
		"username": usernameTest,
		"password": passwordTest,
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/user/login", loginRequest, "")

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var loginResponse LoginResponse

	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	token := loginResponse.Data.Token

	createPositionRequest := map[string]string{
		"name":        "",
		"base_salary": "",
	}

	w, _ = utils_test.MakeRequest(t, app, "POST", "/position/", createPositionRequest, token)

	if w.Code != 400 {
		t.Errorf("Expected status code 400, got %d", w.Code)
	}

}

func TestFailedCreatePositionForbidden(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running TestCreatePosition_Failed Forbidden")

	nameTest := "test position"
	usernameTest := "testPositionForbidden"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest, models.RoleUser)

	loginRequest := map[string]string{
		"username": usernameTest,
		"password": passwordTest,
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/user/login", loginRequest, "")

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var loginResponse LoginResponse

	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	token := loginResponse.Data.Token

	createPositionRequest := map[string]interface{}{
		"name":        "Senior Software Engineer",
		"base_salary": 10000,
	}

	w, _ = utils_test.MakeRequest(t, app, "POST", "/position/", createPositionRequest, token)

	if w.Code != 403 {
		t.Errorf("Expected status code 403, got %d", w.Code)
	}

}
