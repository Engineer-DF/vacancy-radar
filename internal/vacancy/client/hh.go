package client

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "github.com/Engineer-DF/vacancy-radar/internal/vacancy"
)

// ТЕСТОВАЯ ВЕРСИЯ, УДАЛИТЬ СТРУКТУРУ ОТСЮДА
type HHResponse struct {
	Items []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Area struct {
			Name string `json:"name"`
		} `json:"area"`
		Salary struct {
			From int    `json:"from"`
			To   int    `json:"to"`
			Curr string `json:"currency"`
		} `json:"salary"`
	} `json:"items"`
	Found int `json:"found"`
}

type UserAgentTransport struct {
	//userAgent string
	base http.RoundTripper
}

func (t *UserAgentTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	//request.Header.Set("User-Agent", t.userAgent)

	return t.base.RoundTrip(request)
}

// достучаться через бота до хх ру ебаного блять

func PrototypeAPIRequest() error {
	u, err := url.Parse("https://hh.ru")
	if err != nil {
		fmt.Printf("Ошибка парсинга базового URL: %v\n", err)
		return err
	}

	params := url.Values{}
	params.Add("text", "Go")
	params.Add("area", "53")
	params.Add("salary", "150000")
	params.Add("label", "only_with_salary")
	params.Add("experience", "between1And3")
	params.Add("per_page", "1")
	u.RawQuery = params.Encode()

	client := &http.Client{
		Transport: &UserAgentTransport{
			//userAgent: "VacancyRadar/0.0.1 (throwing_knife@mail.ru)",
			base: http.DefaultTransport,
		},
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return err
	}
	request.Header.Set("User-Agent", "VacancyRadar/0.0.1 (throwing_knife@mail.ru)")
	request.Header.Set("Accept", "*/*")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return err
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return fmt.Errorf("WAF заблокировал запрос и вернул HTML вместо данных")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("API returned non-200 status: %s, body: %s", response.Status, string(body))
		return fmt.Errorf("bad status: %s", response.Status)
	}

	var hhResult HHResponse
	if err := json.Unmarshal(body, &hhResult); err != nil {
		fmt.Printf("Ошибка десериализации JSON: %v\n", err)
		return err
	}

	fmt.Printf("Успешно! Найдено вакансий в Краснодаре: %d\n", hhResult.Found)
	return nil
}
