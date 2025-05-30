package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	utils_test "github.com/handarudwiki/payroll-sistem/tests/utils"
)

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func TestSuccessLogin(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running TestAuth_Login_Success")

	nameTest := "testuser"
	usernameTest := "testuser"
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

	if loginResponse.Data.Token == "" {
		t.Error("Expected token to be present in response, got empty token")
	}

}

func TestFailedLoginWrongUsernameOrpassword(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running failed login wrong username or password")

	nameTest := "testuser"
	usernameTest := "testuser2"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest, models.RoleUser)

	loginRequest := map[string]string{
		"username": usernameTest + "wrong",
		"password": passwordTest,
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/user/login", loginRequest, "")

	if w.Code != 401 {
		t.Errorf("Expected status code 401, got %d", w.Code)
	}

	var errorResponse ErrorResponse

	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if errorResponse.Message != "invalid credentials" {
		t.Errorf("Expected error message 'Invalid credentials', got '%s'", errorResponse.Message)
	}

}

func TestFailedBadRequest(t *testing.T) {
	app := utils_test.NewTestApp(t)

	fmt.Println("Running failed login bad request")

	nameTest := "testuser"
	usernameTest := "testuser3"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest, models.RoleUser)

	loginRequest := map[string]string{
		"username": "",
		"password": "",
	}

	w, _ := utils_test.MakeRequest(t, app, "POST", "/user/login", loginRequest, "")

	if w.Code != 400 {
		t.Errorf("Expected status code 400, got %d", w.Code)
	}

	var errorResponse ErrorResponse

	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if errorResponse.Message != "Validation error" {
		t.Errorf("Expected error message 'Bad request', got '%s'", errorResponse.Message)
	}
}
