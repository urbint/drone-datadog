package datadog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	SendMessage(*Event) error
}

type client struct {
	url string
}

func NewClient(api_key string) Client {
	return &client{"https://app.datadoghq.com/api/v1/events?api_key=" + api_key}
}

func (c *client) SendMessage(msg *Event) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	resp, err := http.Post(c.url, "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &Error{resp.StatusCode, string(t)}
	}

	return nil
}
