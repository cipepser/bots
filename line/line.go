package line

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"

	"github.com/cipepser/bots/util"
)

const (
	// URL is an endpoint of LINE notify
	URL = "https://notify-api.line.me/api/notify"
	// URL = "http://localhost:8080"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

// TODO: 説明を書く。引数と返り値
// TODO: アクセストークンの取得とclientの作成まではNewClientの責務にしたほうが設計きれいになりそう
func SendMessage(msg string) error {
	// TODO: 複数のトークルームに贈りたい場合もある。
	t := &Token{}
	err := util.GetToken("./token/line_test.json", t)
	if err != nil {
		return err
	}

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	c := &http.Client{}

	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+t.AccessToken)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to send message, get http status code: " + resp.Status)
	}

	// ***********for debug***********
	// fmt.Println(resp.Header)
	// fmt.Println("*******")
	// str, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(str))
	// *******************************

	return nil
}

// SendImage sent the image to Talkroom
func SendImage(msg string, img io.Reader) error {
	t := &Token{}
	err := util.GetToken("./token/line_test.json", t)
	if err != nil {
		return err
	}

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	c := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("message")
	if err != nil {
		return err
	}
	if _, err = fw.Write([]byte(msg)); err != nil {
		return err
	}

	part := make(textproto.MIMEHeader)
	part.Set("Content-Disposition", `form-data; name="imageFile"; filename="sample.jpg"`)
	part.Set("Content-Type", "image/jpeg")
	fw, err = w.CreatePart(part)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(fw, img)

	w.Close()

	req, err := http.NewRequest("POST", u.String(), &b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+t.AccessToken)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to send image, get http status code: " + resp.Status)
	}

	// ***********for debug***********
	// fmt.Println(resp.Header)
	// fmt.Println("*******")
	// str, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(str))
	// *******************************

	return nil
}
