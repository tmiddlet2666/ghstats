package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/tmiddlet2666/ghstats/pkg/config"
	"github.com/tmiddlet2666/ghstats/pkg/constants"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// GetAPIURL returns the API URL for a username and repository
func GetAPIURL(username, repository string) string {
	return fmt.Sprintf("%s/%s/%s", constants.APIPrefix, username, repository)
}

// GetReleases returns release details
func GetReleases(username, repository string) ([]config.Release, error) {
	var (
		err      error
		data     []byte
		status   int
		releases []config.Release
		repoURL  = GetAPIURL(username, repository) + "/releases"
	)

	data, status, err = HttpGETRequest(repoURL)
	if err != nil {
		return releases, fmt.Errorf("unable to get releases for user: %s, repo: %s - %v", username, repository, err)
	}

	if status != 200 {
		return releases, fmt.Errorf("status code %d returned for %s", status, repoURL)
	}

	err = json.Unmarshal(data, &releases)
	if err != nil {
		return releases, fmt.Errorf("unable to unmarshall releases. %v", err)
	}

	return releases, nil

}

// HttpGETRequest issues a GET request and returns the contents
func HttpGETRequest(URL string) ([]byte, int, error) {
	var (
		err    error
		req    *http.Request
		resp   *http.Response
		body   []byte
		buffer bytes.Buffer
	)
	cookies, _ := cookiejar.New(nil)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	client := &http.Client{Transport: tr,
		Timeout: time.Duration(constants.RequestTimeout) * time.Second,
		Jar:     cookies,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err = http.NewRequest("GET", URL, bytes.NewBuffer(constants.EmptyByte))
	if err != nil {
		return constants.EmptyByte, 0, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return constants.EmptyByte, 0, err
	}

	defer resp.Body.Close()

	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		return constants.EmptyByte, 0, err
	}

	body = buffer.Bytes()

	return body, resp.StatusCode, nil
}
