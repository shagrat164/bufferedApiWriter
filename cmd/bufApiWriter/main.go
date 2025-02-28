package main

import (
	"sync"

	"github.com/shagrat164/bufferedApiWriter/internal/buffer"
	"github.com/shagrat164/bufferedApiWriter/internal/model"
	"github.com/shagrat164/bufferedApiWriter/internal/writer"
)

func main() {
	buffer := buffer.NewBuffer()
	var wg sync.WaitGroup

	// Заполняем буфер 10 записями
	for i := 0; i < 10; i++ {
		buffer.Put(model.Fact{
			PeriodStart:         "2024-12-01",
			PeriodEnd:           "2024-12-31",
			PeriodKey:           "month",
			IndicatorToMoID:     "227373",
			IndicatorToMoFactID: "0",
			Value:               "1",
			FactTime:            "2024-12-31",
			IsPlan:              "0",
			AuthUserID:          "40",
			Comment:             "buffer Protsvetov Danila",
		})
	}

	wg.Add(1)
	go writer.Worker(buffer, &wg) // Запускаем горутину для обработки буфера

	wg.Wait() // Ожидание завершения
}
