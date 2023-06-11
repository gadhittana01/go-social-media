package resthttp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gadhittana01/socialmedia/services"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (p UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	res, err := p.userService.GetUsers(context.Background())
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}

func (p UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type CreateUserReq struct {
		Fullname string `json:"fullname"`
	}

	reqBody := CreateUserReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Fullname == "" {
		resp.SetBadRequest("fullname cannot be null", w)
		return
	}

	res, err := p.userService.CreateUser(context.Background(), reqBody.Fullname)
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetCreated(res, w)
	return
}

func (p UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type UpdateUserReq struct {
		Fullname string `json:"fullname"`
	}

	reqBody := UpdateUserReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		resp.SetBadRequest("Invalid Request Parameter", w)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Fullname == "" {
		resp.SetBadRequest("fullname cannot be null", w)
		return
	}

	res, err := p.userService.UpdateUser(context.Background(), services.UpdateUserParams{
		ID:       int32(uid),
		Fullname: reqBody.Fullname,
	})
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}

func (p UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	id := r.URL.Query().Get("id")
	if id == "" {
		resp.SetBadRequest("Invalid Request Parameter", w)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	err = p.userService.DeleteUser(context.Background(), int32(uid))
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(map[string]interface{}{
		"status": "success",
	}, w)
	return
}
