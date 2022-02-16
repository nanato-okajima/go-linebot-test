package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/xerrors"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

var weatherUrl = "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json"

func GetWeather() (string, error) {
	body, err := httpGetBody(weatherUrl)
	if err != nil {
		return "", err
	}

	weather, err := formatWeather(body)
	if err != nil {
		return "", err
	}

	result := weather.ToS()

	return result, nil
}

func httpGetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.Errorf("Get Http Error: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("IO read Error:: %s", err)
	}

	return body, err
}

func formatWeather(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		return nil, xerrors.Errorf("JSON Unmarshal  error: %s", err)
	}

	return weather, nil
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}
