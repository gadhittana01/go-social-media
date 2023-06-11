package resthttp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gadhittana01/socialmedia/services"
)

type PostHandler struct {
	postService PostService
}

func NewPostHandler(postService PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (p PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	res, err := p.postService.GetPosts(context.Background())
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}

func (p PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type CreatePostReq struct {
		Userid      int32   `json:"user_id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		TagIDs      []int32 `json:"tag_ids"`
	}

	reqBody := CreatePostReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Userid == 0 {
		resp.SetBadRequest("user_id cannot be null", w)
		return
	}

	if reqBody.Title == "" {
		resp.SetBadRequest("title cannot be null", w)
		return
	}

	if reqBody.Description == "" {
		resp.SetBadRequest("description cannot be null", w)
		return
	}

	res, err := p.postService.CreatePost(context.Background(), services.CreatePostParams{
		Userid:      reqBody.Userid,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		TagID:       reqBody.TagIDs,
	})
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetCreated(res, w)
	return
}

func (p PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	type UpdatePostReq struct {
		ID          int32   `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		TagIDs      []int32 `json:"tag_ids"`
	}

	reqBody := UpdatePostReq{}
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

	pid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	if reqBody.Title == "" {
		resp.SetBadRequest("title cannot be null", w)
		return
	}

	if reqBody.Description == "" {
		resp.SetBadRequest("description cannot be null", w)
		return
	}

	res, err := p.postService.UpdatePost(context.Background(), services.UpdatePostParams{
		ID:          int32(pid),
		Title:       reqBody.Title,
		Description: reqBody.Description,
		TagID:       reqBody.TagIDs,
	})
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetCreated(res, w)
	return
}

func (p PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	id := r.URL.Query().Get("id")
	if id == "" {
		resp.SetBadRequest("Invalid Request Parameter", w)
		return
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		resp.SetBadRequest(err.Error(), w)
		return
	}

	err = p.postService.DeletePost(context.Background(), int32(pid))
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(map[string]interface{}{
		"status": "success",
	}, w)
	return
}
