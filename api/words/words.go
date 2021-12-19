package words

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GetRandomEndpoint = "?number=%d"
)

type Api struct {
	BaseURL string
}

func NewApi(baseUrl string) Api {
	return Api{
		BaseURL: baseUrl,
	}
}

func (a Api) GetRandom(amount int) ([]string, error) {
	var words []string
	route := fmt.Sprintf(GetRandomEndpoint, amount)
	resp, err := http.Get(a.BaseURL + route)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &words); err != nil {
		return nil, err
	}

	return words, err
}
