package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Curl struct {
	Token     string
	ApiServer string
}

func NewCurl(token string, apiServer string) *Curl {
	if apiServer == "" {
		apiServer = "https://api.telegram.org/"
	}
	return &Curl{
		Token:     token,
		ApiServer: apiServer,
	}
}

func (c *Curl) setURL(method string) string {
	if len(c.ApiServer) > 0 && c.ApiServer[len(c.ApiServer)-1] == '/' {
		return fmt.Sprintf("%sbot%s/%s", c.ApiServer, c.Token, method)
	}
	return fmt.Sprintf("%s/bot%s/%s", c.ApiServer, c.Token, method)
}

func (c *Curl) Endpoint(method string, content map[string]interface{}, post bool, response bool) (map[string]interface{}, error) {
	url := c.setURL(method)

	var jsonData []byte
	var err error
	if post {
		jsonData, err = json.Marshal(content)
		if err != nil {
			return nil, err
		}
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if post {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if !response {
		go func() {
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error during async request: %v", err)
				return
			}
			defer resp.Body.Close()
		}()
		return nil, nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}