package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
)

func (r *bookRepository) FetchBooks(query *model.BookQuery) (*model.CurrencyResponse, error) {
	resp, err := http.Get(r.getUri(query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.CurrencyResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, fmt.Errorf("error al consultar  API")
	}

	return &response, nil
}

func (r *bookRepository) getUri(query *model.BookQuery) string {
	baseURL := os.Getenv("URI_CURRENCY")

	params := url.Values{}
	params.Add("access_key", os.Getenv("API_KEY"))
	params.Add("currencies", query.CurrencyTo)
	params.Add("source", query.CurrencyFrom)

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}
