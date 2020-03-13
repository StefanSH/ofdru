package ofdru

import (
	"fmt"
	"testing"
	"time"
)

func TestOfdRu(t *testing.T) {
	client := OfdRu("3245001416", "mikhail.merkulov@megafon.ru", "121212", "https://ofd.ru")
	date := time.Date(2018, 04, 14, 0, 0, 0, 0, time.UTC)
	rec, err := client.GetReceipts(date)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(rec)

}
