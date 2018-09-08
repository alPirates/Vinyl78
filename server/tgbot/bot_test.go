package tgbot

import (
	"log"
	"reflect"
	"testing"
)

func TestRegexp(t *testing.T) {
	table := map[string]map[string]string{
		"data@yandex.ru": map[string]string{
			"smtp": "smtp.yandex.ru",
			"port": "smtp.yandex.ru:25",
		},
		"mail.something@google.com": map[string]string{
			"smtp": "smtp.gmail.com",
			"port": "smtp.gmail.com:587",
		},
		"hmmm.somet.hing@google.com": map[string]string{
			"smtp": "smtp.gmail.com",
			"port": "smtp.gmail.com:587",
		},
	}

	for key, value := range table {
		r, err := getSmtpName(key)
		if err != nil {
			log.Println(err)
		}
		if !reflect.DeepEqual(r, value) {
			t.Errorf("Expected %v, on test %v, got %v", value, key, r)
		}
	}

}
