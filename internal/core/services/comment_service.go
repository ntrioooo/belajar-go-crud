package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type commentService struct {
	comments ports.CommentRepository
	posts    ports.PostRepository
}

func NewCommentService(comments ports.CommentRepository, posts ports.PostRepository) ports.CommentService {
	return &commentService{comments: comments, posts: posts}
}

func (s *commentService) Create(ctx context.Context, userID, postID uint, parentID *uint, body string) (*domain.Comment, error) {
	if userID == 0 {
		return nil, errors.New("unauthorized")
	}
	body = strings.TrimSpace(body)
	if len(body) < 1 {
		return nil, errors.New("comment body required")
	}

	// pastikan post ada
	if p, _ := s.posts.FindByID(ctx, postID); p == nil {
		return nil, errors.New("post not found")
	}

	// (opsional) jika parentID != nil, bisa validasi parent exist dan parent.PostID == postID
	c := &domain.Comment{
		PostID:   postID,
		UserID:   userID,
		ParentID: parentID,
		Body:     body,
	}
	if err := s.comments.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *commentService) GetByPost(ctx context.Context, viewerID, postID uint) ([]domain.CommentView, error) {
	items, err := s.comments.ListByPost(ctx, postID)
	if err != nil {
		return nil, err
	}

	// bangun view
	out := make([]domain.CommentView, 0, len(items))
	for i := range items {
		c := items[i]
		liked := false
		if viewerID != 0 {
			liked, _ = s.comments.IsLiked(ctx, viewerID, c.ID)
		}
		likeCnt, _ := s.comments.CountLikes(ctx, c.ID)
		replyCnt, _ := s.comments.CountReplies(ctx, c.ID)

		// ambil username author (reuse batch di PostRepository kalau kamu mau,
		// atau biarkan repo Comment menambah batch function; untuk MVP, kita simple)
		authorUsername := "" // fallback
		// jika kamu sudah punya batch map via PostRepository, bisa dipakai di sini.
		// atau, tambahkan CommentRepository.BatchAuthorUsernames jika mau optimal.

		out = append(out, domain.CommentView{
			Comment:        c,
			AuthorUsername: authorUsername,
			LikeCount:      likeCnt,
			LikedByMe:      liked,
			ReplyCount:     replyCnt,
		})
	}
	return out, nil
}

func (s *commentService) GetReplies(ctx context.Context, viewerID, parentID uint) ([]domain.CommentView, error) {
	items, err := s.comments.ListReplies(ctx, parentID)
	if err != nil {
		return nil, err
	}
	out := make([]domain.CommentView, 0, len(items))
	for i := range items {
		c := items[i]
		liked := false
		if viewerID != 0 {
			liked, _ = s.comments.IsLiked(ctx, viewerID, c.ID)
		}
		likeCnt, _ := s.comments.CountLikes(ctx, c.ID)
		replyCnt, _ := s.comments.CountReplies(ctx, c.ID)

		out = append(out, domain.CommentView{
			Comment:        c,
			AuthorUsername: "",
			LikeCount:      likeCnt,
			LikedByMe:      liked,
			ReplyCount:     replyCnt,
		})
	}
	return out, nil
}

func (s *commentService) Delete(ctx context.Context, userID, commentID uint) error {
	if userID == 0 {
		return errors.New("unauthorized")
	}
	c, err := s.comments.FindByID(ctx, commentID)
	if err != nil || c == nil {
		return errors.New("comment not found")
	}
	// aturan sederhana: author bisa hapus komennya sendiri
	if c.UserID != userID {
		return errors.New("forbidden")
	}
	return s.comments.Delete(ctx, commentID)
}

func (s *commentService) ToggleLike(ctx context.Context, userID, commentID uint) (bool, int64, error) {
	if userID == 0 {
		return false, 0, errors.New("unauthorized")
	}
	liked, err := s.comments.IsLiked(ctx, userID, commentID)
	if err != nil {
		return false, 0, err
	}
	if liked {
		if err := s.comments.Unlike(ctx, userID, commentID); err != nil {
			return false, 0, err
		}
	} else {
		if err := s.comments.Like(ctx, userID, commentID); err != nil {
			return false, 0, err
		}
	}
	nowLiked := !liked
	count, _ := s.comments.CountLikes(ctx, commentID)
	return nowLiked, count, nil
}
