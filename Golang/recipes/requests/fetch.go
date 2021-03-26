package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func Fetch(){
	req, err := http.NewRequest("GET", "https://swapi.dev/api/people/1",nil)
		if err != nil {
		log.Println(err)
		return
	}

	req.Close = true // closes the request once the request has been made

	// Transport caches connections for future use 
	var transport = &http.Transport{DisableKeepAlives: true}
	var myClient = &http.Client{Timeout: 10 * time.Second, Transport: transport}
	resp, err := myClient.Do(req)

	p := make([]byte, 500)
	a1, err := resp.Body.Read(p)
	// a, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(a)
	fmt.Println(a1)

	fmt.Println(err)
	fmt.Println(p)
	fmt.Println(string(p))
}

func FetchVariable(i int) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://swapi.dev/api/people/%d", i), nil)
	if err != nil {
		log.Println(err)
		return 
	}
	req.Close = true

}

// Taken from the starwars api wrapper, kinda cool ngl
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}

func TestClientDo(t *testing.T) {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Film{Title: "Example"})
	})

	ts := httptest.NewServer(hf)
	defer ts.Close()

	c := NewClient(BaseURL(ts.URL))

	r, err := c.newRequest("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var v Film

	if _, err := c.do(r, &v); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := v.Title, "Example"; got != want {
		t.Fatalf("v.Title = %q, want %q", got, want)
	}
}