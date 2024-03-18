package api

import (
	"encoding/json"
	"fem-go/cryptomasters/datatypes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (*datatypes.Rate, error) {
	resp, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code received: %v", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cexResp CexResponse

	err = json.Unmarshal(bodyBytes, &cexResp)
	if err != nil {
		return nil, err
	}

	return &datatypes.Rate{
		Price:    cexResp.Bid,
		Currency: currency,
	}, nil
}
