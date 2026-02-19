package request

import (
	"io"
	"net/http"
	"webapp/app/cookies"
)

func RequestAuth(req *http.Request, method, url string, date io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, date)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.ReadToken(req)
	request.Header.Add("Authorization", "Bearer "+cookie)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
