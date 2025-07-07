package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/vsevolodGolangDev/bank-service/models"
)

func Convert(euro float32, currency string) (float32, error) {
	if currency == "" {
		return euro, nil
	}

	resp, err := http.Get("http://api.exchangeratesapi.io/v1/latest?access_key=5bb179314fdbfaa6a839358e571d426f&base=EUR&symbols=" + currency)
	if err != nil {
		return 0, err
	}

	var get models.Response
	json.NewDecoder(resp.Body).Decode(&get)
	result, err := strconv.ParseFloat(fmt.Sprintf("%.2f", euro*get.Rates[currency]), 32)
	if err != nil {
		return 0, err
	}

	return float32(result), nil
}

func ParseTime(value string, t *testing.T) time.Time {
	timeAt, err := time.Parse(time.DateTime, value)
	if err != nil {
		t.Error(err)
	}

	return timeAt
}
