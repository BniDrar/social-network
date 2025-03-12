package post

import (
	"net/http"
	"socialNetwork/entity"
)

func (p *post) CreatePostRepo (post entity.Post, id int) (int, int, error) {
	
	return 0, http.StatusCreated, nil
}

func (p *post) CreateGroup(post entity.Post, id int) (int, int, error) {
	return 0, http.StatusCreated, nil
}