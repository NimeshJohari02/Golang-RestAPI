package unittesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	user "nimeshjohari02.com/restapi/user"
)

func TestFetchUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/getUserById", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "7a62cbbc-7f70-48db-b740-7be5fef57328")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.GetUserById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected :=`{"Email":"\"nimeshjohari95@gmail.com\"","Password":"$2a$14$C0pdzLah2gIHdttH2kbiOOf55mwHgEdlewV1Jlt2nyK8E8Jo.PSga","Posts":null,"_id":"6161ad1923c42f4188930f8a","id":"7a62cbbc-7f70-48db-b740-7be5fef57328","name":"\"NJ\""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
