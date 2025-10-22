package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type postService struct {
	posts ports.PostRepository
}

func NewPostService(posts ports.PostRepository) ports.PostService {
	return &postService{posts: posts}
}

func (s *postService) Create(ctx context.Context, authorID uint, title, body string, categoryID uint) (*domain.Post, error) {
	title = strings.TrimSpace(title)
	body = strings.TrimSpace(body)
	if authorID == 0 {
		return nil, errors.New("unauthorized")
	}
	if len(title) < 3 || len(body) < 3 {
		return nil, errors.New("title/body too short")
	}

	p := &domain.Post{Title: title, Body: body, UserID: authorID, CategoryID: categoryID}
	if err := s.posts.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

// func (s *postService) GetByID(ctx context.Context, id uint) (*domain.Post, error) {
// 	return s.posts.FindByID(ctx, id)
// }

// func (s *postService) List(ctx context.Context) ([]domain.Post, error) {
// 	return s.posts.List(ctx)
// }

func (s *postService) Update(ctx context.Context, id uint, title, body string, categoryID uint) (*domain.Post, error) {
	p, err := s.posts.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	p.Title, p.Body, p.CategoryID = title, body, categoryID
	if err := s.posts.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *postService) Delete(ctx context.Context, id uint) error {
	return s.posts.Delete(ctx, id)
}

func (s *postService) List(ctx context.Context, viewerID uint) ([]domain.PostView, error) {
	posts, err := s.posts.List(ctx)
	if err != nil {
		return nil, err
	}
	// kumpulkan userIDs
	uids := make([]uint, 0, len(posts))
	cids := make([]uint, 0, len(posts))
	for i := range posts {
		uids = append(uids, posts[i].UserID)
		cids = append(cids, posts[i].CategoryID)
	}

	authorMap, _ := s.posts.BatchAuthorUsernames(ctx, uids)
	catMap, _ := s.posts.BatchCategoryNames(ctx, cids)

	out := make([]domain.PostView, 0, len(posts))
	for i := range posts {
		p := posts[i]
		cnt, _ := s.posts.CountLikes(ctx, p.ID)
		liked := false
		if viewerID != 0 {
			liked, _ = s.posts.IsLiked(ctx, viewerID, p.ID)
		}
		out = append(out, domain.PostView{
			Post:           p,
			AuthorUsername: authorMap[p.UserID],
			CategoryName:   catMap[p.CategoryID], // <-- diisi dari map
			LikeCount:      cnt,
			LikedByMe:      liked,
		})
	}
	return out, nil
}

func (s *postService) GetByID(ctx context.Context, viewerID, id uint) (*domain.PostView, error) {
	p, err := s.posts.FindByID(ctx, id)
	if err != nil || p == nil {
		return nil, err
	}
	cnt, _ := s.posts.CountLikes(ctx, p.ID)
	liked := false
	if viewerID != 0 {
		liked, _ = s.posts.IsLiked(ctx, viewerID, p.ID)
	}
	authorMap, _ := s.posts.BatchAuthorUsernames(ctx, []uint{p.UserID})
	catMap, _ := s.posts.BatchCategoryNames(ctx, []uint{p.CategoryID})
	view := &domain.PostView{
		Post: *p, AuthorUsername: authorMap[p.UserID], CategoryName: catMap[p.CategoryID], LikeCount: cnt, LikedByMe: liked,
	}
	return view, nil
}

func (s *postService) ToggleLike(ctx context.Context, userID, postID uint) (bool, int64, error) {
	if userID == 0 {
		return false, 0, errors.New("unauthorized")
	}
	liked, err := s.posts.IsLiked(ctx, userID, postID)
	if err != nil {
		return false, 0, err
	}
	if liked {
		if err := s.posts.Unlike(ctx, userID, postID); err != nil {
			return false, 0, err
		}
	} else {
		if err := s.posts.Like(ctx, userID, postID); err != nil {
			return false, 0, err
		}
	}
	likedNow := !liked
	count, _ := s.posts.CountLikes(ctx, postID)
	return likedNow, count, nil
}
