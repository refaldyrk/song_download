package song_donwload

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func GetSong(keyword string) Song {
	res, err := http.Get(BASE_URL_SEARCH + keyword)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	var song Song
	if err := json.NewDecoder(res.Body).Decode(&song); err != nil {
		panic(err)
	}
	return song
}

func (i Item) PreDownload() Extractor {
	body := map[string]string{
		"ftype": "mp3",
		"url":   "https://www.youtube.com/watch?v=" + i.Id,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
		return Extractor{}
	}

	client := &http.Client{Timeout: 100 * time.Second}
	req, err := http.NewRequest("POST", BASE_URL_PRE_DOWNLOAD, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
		return Extractor{}
	}

	req.Header.Set("Origin", "https://apidl.net")
	req.Header.Set("Referer", "https://apidl.net/")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
		return Extractor{}
	}
	defer res.Body.Close()

	var extractor Extractor
	if err := json.NewDecoder(res.Body).Decode(&extractor); err != nil {
		panic(err)
		return Extractor{}
	}

	return extractor
}

func (t Task) Download(jwt string) string {
	body := map[string]string{
		"hash": t.Hash,
		"jwt":  jwt,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
		return ""
	}

	client := &http.Client{Timeout: 100 * time.Second}
	req, err := http.NewRequest("POST", BASE_URL_PRE_DOWNLOAD, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
		return ""
	}

	req.Header.Set("Origin", "https://apidl.net")
	req.Header.Set("Referer", "https://apidl.net/")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
		return ""
	}
	defer res.Body.Close()

	var extractor map[string]string
	if err := json.NewDecoder(res.Body).Decode(&extractor); err != nil {
		panic(err)
		return ""
	}

	bodyTask := map[string]string{
		"taskId": extractor["taskId"],
	}

	jsonBodyTask, err := json.Marshal(bodyTask)
	if err != nil {
		panic(err)
		return ""
	}

	clientTask := &http.Client{Timeout: 100 * time.Second}
	reqTask, err := http.NewRequest("POST", BASE_URL_DOWNLOAD, bytes.NewBuffer(jsonBodyTask))
	if err != nil {
		panic(err)
		return ""
	}

	reqTask.Header.Set("Origin", "https://apidl.net")
	reqTask.Header.Set("Referer", "https://apidl.net/")
	reqTask.Header.Set("Content-Type", "application/json")

	resTask, err := clientTask.Do(reqTask)
	if err != nil {
		panic(err)
		return ""
	}
	defer resTask.Body.Close()

	var task Download
	if err := json.NewDecoder(resTask.Body).Decode(&task); err != nil {
		panic(err)
		return ""
	}
	var url string

	for {
		startTime := time.Now()
		cbg := context.Background()
		context.WithTimeout(cbg, 10*time.Second)

		if task.Status == "finished" {
			url = task.Download
			break
		}

		time.Sleep(1 * time.Second)

		if time.Since(startTime) > 10*time.Second {
			break
		}
	}

	return url
}
