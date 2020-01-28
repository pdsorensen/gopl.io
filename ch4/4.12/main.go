package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type xkcdResponse struct {
	Transcript string
	Num        int
}

type xkcdJsonItem struct {
	Transcript string `json:"Transcript"`
	Num        int    `json:"Num"`
}

func main() {
	// buildIndex()
	search(os.Args[1])
}

func search(word string) {
	file, _ := os.Open("big_marhsall.json")

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	// initialize our comics array
	var comics []xkcdJsonItem

	// unmarshal our byteArray into 'comics'
	err := json.Unmarshal(byteValue, &comics)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(comics); i++ {
		if strings.Contains(comics[i].Transcript, word) {
			fmt.Printf("Number %d has word %s in it: \n", comics[i].Num, word)
			fmt.Println(comics[i].Transcript)
		}
	}
}

func buildIndex() {
	baseInfo := getBaseInfo()
	var comics []xkcdResponse

	fmt.Printf("Need to fetch %d comics\n", baseInfo.Num)
	for i := 1; i < baseInfo.Num; i++ {

		if i == 404 { // that's a funny one
			continue
		}

		fmt.Println(i)
		comic := getInfo(i)
		comics = append(comics, comic)
	}

	// build file
	jsonString, _ := json.Marshal(comics)
	ioutil.WriteFile("big_marhsall.json", jsonString, os.ModePerm)
}

func getBaseInfo() xkcdResponse {
	resp, _ := http.Get("https://xkcd.com/info.0.json")
	var result xkcdResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}

	resp.Body.Close()

	return result
}

func getInfo(index int) xkcdResponse {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", index)
	resp, _ := http.Get(url)
	var result xkcdResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}

	resp.Body.Close()

	return result
}
