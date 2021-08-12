package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"codes/gocodes/dg/models"
)

type LoginData struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Token struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

func Login() (string, error) {
	login := &LoginData{
		Username: "admin",
		Password: "admin@2013",
	}
	loginJs, err := json.Marshal(login)
	if err != nil {
		return "", err
	}
	r := bytes.NewReader(loginJs)
	req, err := http.NewRequest(http.MethodPost, models.LoginURL, r)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	tk := &Token{}
	if err = json.Unmarshal(respBytes, tk); err != nil {
		return "", err
	}
	return tk.Token, nil
}
