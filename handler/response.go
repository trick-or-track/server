package handler

import (
	"github.com/trick-or-track/server/model"
	"github.com/trick-or-track/server/utils"
)

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	Token string `json:"token"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = utils.Proper(u.Username)
	r.User.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}

type dataResponseData struct {
	UserID int `json:"userId"`
	Year   int `json:"year"`
	One    int `json:"one"`
	Two    int `json:"two"`
	Three  int `json:"three"`
	Four   int `json:"four"`
	Five   int `json:"five"`
	Six    int `json:"six"`
	Seven  int `json:"seven"`
	Eight  int `json:"eight"`
	Nine   int `json:"nine"`
	Ten    int `json:"ten"`
}

type dataResponse struct {
	Data []dataResponseData `json:"data"`
}

func newDataResponse(data []*model.Data) *dataResponse {
	r := new(dataResponse)
	for _, d := range data {
		nextD := new(dataResponseData)

		nextD.UserID = d.UserID
		nextD.Year = d.Year
		nextD.One = d.One
		nextD.Two = d.Two
		nextD.Three = d.Three
		nextD.Four = d.Four
		nextD.Five = d.Five
		nextD.Six = d.Six
		nextD.Seven = d.Seven
		nextD.Eight = d.Eight
		nextD.Nine = d.Nine
		nextD.Ten = d.Ten
		r.Data = append(r.Data, *nextD)
	}
	return r
}
