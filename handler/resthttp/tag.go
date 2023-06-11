package resthttp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gadhittana01/socialmedia/services"
)

type TagHandler struct {
	tagService TagService
}

func NewTagHandler(tagService TagService) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

func (p TagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	res, err := p.tagService.GetTags(context.Background())
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}

func (p TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type CreateUserReq struct {
		Tagname string `json:"tagname"`
	}

	reqBody := CreateUserReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Tagname == "" {
		resp.SetBadRequest("Tagname cannot be null", w)
		return
	}

	res, err := p.tagService.CreateTag(context.Background(), reqBody.Tagname)
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetCreated(res, w)
	return
}

func (p TagHandler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type UpdateTagReq struct {
		Tagname string `json:"tagname"`
	}

	reqBody := UpdateTagReq{}
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

	tid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Tagname == "" {
		resp.SetBadRequest("Tagname cannot be null", w)
		return
	}

	res, err := p.tagService.UpdateTag(context.Background(), services.UpdateTagParams{
		ID:      int32(tid),
		Tagname: reqBody.Tagname,
	})
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetCreated(res, w)
	return
}

func (p TagHandler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	id := r.URL.Query().Get("id")
	if id == "" {
		resp.SetBadRequest("Invalid Request Parameter", w)
		return
	}

	tid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	err = p.tagService.DeleteTag(context.Background(), int32(tid))
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(map[string]interface{}{
		"status": "success",
	}, w)
	return
}
