package handler

import (
	"bytes"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/exec"
	"io"
	"io/ioutil"
	"mime/multipart"
	"github.com/rs/xid"
	"fmt"
	"encoding/json"
)

type Response struct {
	IsSuccess bool
	Result interface{}
}

type Empath struct {
	Error int `json:"error"`
	Calm int `json:"calm"`
	Anger int `json:"anger"`
	Joy int `json:"joy"`
	Sorrow int `json:"sorrow"`
	Energy int `json:"energy"`
}

func IndexHandler(c echo.Context) error {
	file, _ := c.FormFile("mp3")
	if file == nil {
		return c.JSON(
		http.StatusOK,
		Response{
			IsSuccess: false,
			Result: "mp3 not found",
		})
	}

	filepath, err := Upload(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	filepath = "public/" + filepath
	fmt.Println(filepath)
	fmt.Println(filepath[:len(filepath) - 3] + "wav")
	s := []string{"-i", filepath, "-vn", "-ac", "1", "-ar", "11025",
								"-acodec", "pcm_s16le", "-f", "wav",
								filepath[:len(filepath) - 3] + "wav"}
	fmt.Println(s)
	err = exec.Command("ffmpeg", s...).Run()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	empath, err := InvokeEmpathAPI(filepath[:len(filepath) - 3] + "wav")

	fmt.Printf("%#v", empath)
	fmt.Printf("%#v", err)
	return c.JSON(
		http.StatusOK,
		Response{
			IsSuccess: true,
			Result: empath,
		})
}

func Upload(file *multipart.FileHeader) (string, error) {
	guid := xid.New()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := guid.String() + ".mp3"

	var filepath string
	basepath := "./public/"
	filepath = "uploads/single/"
	fullpath := basepath + filepath + filename

	os.MkdirAll(basepath+filepath, os.ModePerm)

	dst, err := os.Create(fullpath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filepath + filename, nil
}

func InvokeEmpathAPI(filepath string) (*Empath, error) {
	url := "https://api.webempath.net/v2/analyzeWav"
	file, err := os.Open(filepath)

	body := &bytes.Buffer{}

	mw := multipart.NewWriter(body)

	fw, err := mw.CreateFormFile("wav", filepath)
	_, err = io.Copy(fw, file)
	fmt.Println(err)

	label, err := mw.CreateFormField("apikey")
	if err != nil {
			return nil, err
	}
	label.Write([]byte("OhRfQkL4KFjw37ov77Gi-hw4-gSDnjiKAeExAFAgKUw"))
	contentType := mw.FormDataContentType()

	err = mw.Close()

	resp, err := http.Post(url, contentType, body)
	fmt.Println(err)
	fmt.Println(resp.Body)

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		jsonBytes := ([]byte)(string(bodyBytes))
    empath := new(Empath)

    if err := json.Unmarshal(jsonBytes, empath); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return nil, err
    }
		return empath, nil
	}
	return nil, err
}
