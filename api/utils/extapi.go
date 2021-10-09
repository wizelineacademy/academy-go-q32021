package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func FetchInfo(idParam int) string {

	id := strconv.Itoa(idParam)

	resp, err := http.Get("https://rupaulsdragrace.fandom.com/api.php/?action=query&prop=extracts&exlimit=1&explaintext=true&pageids=" + id + "&format=json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	plainBody := string(body)
	object := strings.Split(plainBody, "{")[4]
	extract := strings.Split(object, "\"extract\":")[1]
	bio := strings.Split(extract, "==")[0]
	trimmedBio := strings.Replace(
		strings.Replace(
			strings.Replace(bio, "\"", "", -1), "\\n", "", -1), ",", "", -1)

	return trimmedBio

}
