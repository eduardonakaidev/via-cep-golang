package services

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/eduardofrnkdev/via-cep-golang/internal/cache"
	"github.com/eduardofrnkdev/via-cep-golang/pkg/models"
)

// GetCep fetches CEP information from cache or the ViaCEP API.
func GetCep(id string) (string, error) {
	if cached := cache.GetFromCache(id); cached != "" {
		return cached, nil
	}

	resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch CEP: %s", resp.Status)
	}

	var c models.Cep
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return "", err
	}

	res, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return cache.SaveOnCache(id, string(res)), nil
}
