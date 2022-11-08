package api

import (
	"bytes"
	"chat/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Translate(msg []byte) []byte {
	//envファイル読み込み処理

	Key := os.Getenv("subscriptionKey")
	location := os.Getenv("location")
	endpoint := os.Getenv("endpoint")
	uri := endpoint + os.Getenv("uri")
	//envファイルで読み込んだものを代入
	//IMPORTANT PLEASE READ Check your subscriptionKey and location.

	u, _ := url.Parse(uri)
	v, _ := url.Parse(uri)
	q := u.Query()
	r := v.Query()
	q.Add("from", "ja")
	q.Add("to", "en")
	r.Add("from", "en")
	r.Add("to", "ja")
	u.RawQuery = q.Encode()
	v.RawQuery = r.Encode()
	//再翻訳するために２つ作成
	//Create an anonymous struct for your request body and encode it to JSON
	text := string(msg)

	body := []struct {
		Text string `json:"text"`
	}{
		{Text: text},
	}

	b, _ := json.Marshal(body)

	// Build the HTTP POST request
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Connected")
	// Add required headers to the request
	req.Header.Add("Ocp-Apim-Subscription-Key", Key)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	// Call the Translator API
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("ok")

	var arr []structs.TranslationRes

	translation, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(translation, &arr)
	if err != nil {
		log.Fatal(err)
	}

	text_str := arr[0].Translation[0].Text
	//fmt.Println(text_str)
	///////////////////////////////////////ここまで1回目翻訳///////////////////////////////////

	Datum := []struct { //再翻訳するためのjson構造体を用意する
		Text string `json:"text"`
	}{
		{Text: text_str},
	}
	c, _ := json.Marshal(Datum)

	// Build the HTTP POST request
	req2, err := http.NewRequest("POST", v.String(), bytes.NewBuffer(c))
	if err != nil {
		log.Fatal(err)
	}

	// Add required headers to the request
	req2.Header.Add("Ocp-Apim-Subscription-Key", Key)
	req2.Header.Add("Ocp-Apim-Subscription-Region", location)
	req2.Header.Add("Content-Type", "application/json")
	//翻訳と再翻訳で２回行われるため、Azure上の使用文字数は２倍になる

	// Call the Translator API
	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}

	translations2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Fatal(err)
	} //res2.Bodyからjsonを読み込む

	err = json.Unmarshal(translations2, &arr)
	if err != nil {
		log.Fatal(err)
	} //Jsonをstructに変換する
	text_str2 := arr[0].Translation[0].Text

	text_to_byte := []byte(text_str2)
	//Byte型で返ってくるため
	return text_to_byte
}
