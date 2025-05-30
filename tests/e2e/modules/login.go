package modules

import (
	"encoding/json"
	"testing"

	utils_test "github.com/handarudwiki/payroll-sistem/tests/utils"
)

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func TestAuth_Login_Success(t *testing.T) {
	app := utils_test.NewTestApp(t)

	nameTest := "testuser"
	usernameTest := "testuser"
	passwordTest := "testpassword"

	utils_test.CreateUserTest(app.DB, nameTest, usernameTest, passwordTest)

	loginRequest := map[string]string{
		"username": "testuser",
		"password": "testpassword",
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
