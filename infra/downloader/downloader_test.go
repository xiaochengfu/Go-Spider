package downloader

import (
	"fmt"
	"github.com/xiaochengfu/Go-Spider/infra/url"
	"testing"
)

func TestGetHttpResponse(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.ChinaFilmURL,
		},
		{
			url: url.CboURL,
		},
		{
			url: url.NuoMiURL,
		},
		{
			url: url.PFangURL,
		},
	}

	for _, test := range tests {
		htmlByte, _ := GetHttpResponse(test.url, true)
		fmt.Println(string([]byte(htmlByte)))
	}
}
