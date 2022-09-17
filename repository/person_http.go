package repository

import (
	"net/http"
	"os"
)

func GetPersonsFromHTTP(path string) (*http.Request, error) {
	req, errReq := http.NewRequest(http.MethodGet, os.Getenv("HOST_PERSON")+path, nil)
	return req, errReq
}
