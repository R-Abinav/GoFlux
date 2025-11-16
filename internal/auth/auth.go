package auth

import (
	"errors"
	"net/http"
	"strings"
)

//Extracts an API Key from the headers of an HTTP Request
//Example:
/*
	Authorization: ApiKey <API KEY HERE>
*/

func GetApiKey(headers http.Header) (string, error){
	val := headers.Get("Authorization");
	if val == ""{
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2{
		return "", errors.New("malformed Auth header");
	}

	if vals[0] != "ApiKey"{
		return "", errors.New("malformed First Part Auth header");
	}

	return vals[1], nil;
}