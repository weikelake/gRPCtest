package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	CompanyName string
	DirectorFIO string
	Inn         string
	Kpp         string
}

func ParseRusProfile(inn string) (Data, error) {
	id, err := getPageID(inn)
	if err != nil {
		return Data{}, err
	}
	rpData, err := getRusProfileData(id)
	return rpData, err
}

//в запросе нет данных о кпп, поэтому вытаскиваем айди страницы, для дальнейшего парсинга
func getPageID(inn string) (string, error) {
	urlJson := fmt.Sprintf("https://www.rusprofile.ru/ajax.php?query=%s&action=search", inn)
	jsonRes, err := http.Get(urlJson)
	defer jsonRes.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(jsonRes.Body)
	if err != nil {
		return "", err
	}

	//структура для анмаршала json (только необходимые данные).
	//остюда можно было вытащить и другую инфу, но т.к. кпп тут нет
	//и всё равно придётся парсить страницу, решил не делать
	result := struct {
		UlCount int `json:"ul_count"`
		Ul      []struct {
			URL string `json:"url"`
		} `json:"ul"`
		Message string `json:"message"`
	}{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	if result.Message != "OK" {
		err = errors.New("result message != OK")
		return "", err
	}
	if result.UlCount == 0 {
		err = errors.New("inn does not exist")
		return "", err
	}
	var id string
	id = result.Ul[0].URL
	return id, nil
}

//парсинг страницы rusProfile
func getRusProfileData(id string) (Data, error) {
	url := fmt.Sprintf("https://www.rusprofile.ru%s", id)
	res, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Data{}, err
	}
	companyNameData, _ := doc.Find(".company-name").Html()
	companyNameData = strings.ReplaceAll(companyNameData, "&#34;", "\"")
	directorFioData, _ := doc.Find(".company-info__text").Find("a").Find("span").Html()
	innData, _ := doc.Find("#clip_inn").Html()
	kppData, _ := doc.Find("#clip_kpp").Html()

	return Data{
		CompanyName: companyNameData,
		DirectorFIO: directorFioData,
		Inn:         innData,
		Kpp:         kppData,
	}, nil
}
