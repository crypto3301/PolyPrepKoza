package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type KeycloakUser struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Enabled     bool   `json:"enabled"`
	Credentials []struct {
		Type      string `json:"type"`
		Value     string `json:"value"`
		Temporary bool   `json:"temporary"`
	} `json:"credentials"`
}

func AuthenticateUser(username, password string) (string, error) {
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	realm := os.Getenv("REALM")

	data := fmt.Sprintf("client_id=%s&client_secret=%s&username=%s&password=%s&grant_type=password", clientID, clientSecret, username, password)
	req, err := http.NewRequest("POST", keycloakURL+"/realms/"+realm+"/protocol/openid-connect/token", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка Keycloak: %s", string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if result["access_token"] == nil {
		return "", fmt.Errorf("токен доступа не найден в ответе")
	}

	return result["access_token"].(string), nil
}

func CreateUser(username, email, password string) error {
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	realm := os.Getenv("REALM")

	adminToken, err := getAdminToken()
	if err != nil {
		return fmt.Errorf("ошибка получения токена администратора: %v", err)
	}

	user := KeycloakUser{
		Username: username,
		Email:    email,
		Enabled:  true,
		Credentials: []struct {
			Type      string `json:"type"`
			Value     string `json:"value"`
			Temporary bool   `json:"temporary"`
		}{
			{Type: "password", Value: password, Temporary: false},
		},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга пользователя: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/realms/%s/users", keycloakURL, realm), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("ошибка создания пользователя: %d", resp.StatusCode)
	}

	return nil
}
func getAdminToken() (string, error) {
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	data := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret)
	req, err := http.NewRequest("POST", keycloakURL+"/realms/master/protocol/openid-connect/token", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка получения токена администратора: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка Keycloak: %s", string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if result["access_token"] == nil {
		return "", fmt.Errorf("токен администратора не найден в ответе")
	}

	return result["access_token"].(string), nil
}
