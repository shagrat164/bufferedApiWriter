package writer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/shagrat164/bufferedApiWriter/internal/buffer"
	"github.com/shagrat164/bufferedApiWriter/internal/model"
)

const (
	apiURL      = "https://development.kpi-drive.ru/_api/facts/save_fact"
	bearerToken = "48ab34464a5573519725deb5865cc74c"
)

// sendFact отправляет запись на сервер.
func sendFact(fact model.Fact) error {
	data := url.Values{}
	data.Set("period_start", fact.PeriodStart)
	data.Set("period_end", fact.PeriodEnd)
	data.Set("period_key", fact.PeriodKey)
	data.Set("indicator_to_mo_id", fact.IndicatorToMoID)
	data.Set("indicator_to_mo_fact_id", fact.IndicatorToMoFactID)
	data.Set("value", fact.Value)
	data.Set("fact_time", fact.FactTime)
	data.Set("is_plan", fact.IsPlan)
	data.Set("auth_user_id", fact.AuthUserID)
	data.Set("comment", fact.Comment)

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode())) //nolint:noctx
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send fact, status: %d", resp.StatusCode)
	}
	return nil
}

func Worker(buffer *buffer.Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fact, hasData := buffer.Get()
		if !hasData { // Если буфер пуст
			fmt.Println("go Worker stop...")
			return
		}
		if err := sendFact(fact); err != nil {
			fmt.Println("Error sending fact:", err)
			buffer.Put(fact) // Возвращаем в очередь для повторной отправки
		}
	}
}
