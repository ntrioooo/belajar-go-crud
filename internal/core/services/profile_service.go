package services

import (
	"context"
	"sort"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type profileService struct {
	users    ports.UserRepository
	posts    ports.PostRepository
	retweets ports.RetweetRepository
	comments ports.CommentRepository
}

func NewProfileService(
	users ports.UserRepository,
	posts ports.PostRepository,
	retweets ports.RetweetRepository,
	comments ports.CommentRepository,
) ports.ProfileService {
	return &profileService{
		users:    users,
		posts:    posts,
		retweets: retweets,
		comments: comments,
	}
}

// GetProfile menyatukan: post asli milik user + retweet milik user,
// lalu dinormalisasi jadi []PostView agar bisa ditampilkan sebagai satu timeline.
func (s *profileService) GetProfile(ctx context.Context, viewerID, profileUserID uint) (*domain.ProfileView, error) {
	u, err := s.users.FindByID(ctx, profileUserID)
	if err != nil || u == nil {
		return nil, err
	}

	// 1) Ambil semua post (pakai List lalu filter UserID = profileUserID; untuk produksi sebaiknya tambahkan ListByUser di PostRepository)
	allPosts, err := s.posts.List(ctx)
	if err != nil {
		return nil, err
	}
	selfPosts := make([]domain.Post, 0, len(allPosts))
	for i := range allPosts {
		if allPosts[i].UserID == profileUserID {
			selfPosts = append(selfPosts, allPosts[i])
		}
	}

	// 2) Ambil retweet user
	rts, err := s.retweets.ListByUser(ctx, profileUserID)
	if err != nil {
		return nil, err
	}

	// 3) Bangun PostView untuk post asli
	uids := make([]uint, 0, len(selfPosts))
	cids := make([]uint, 0, len(selfPosts))
	for i := range selfPosts {
		uids = append(uids, selfPosts[i].UserID)
		cids = append(cids, selfPosts[i].CategoryID)
	}
	authorMap, _ := s.posts.BatchAuthorUsernames(ctx, uids)
	catMap, _ := s.posts.BatchCategoryNames(ctx, cids)

	timeline := make([]domain.PostView, 0, len(selfPosts)+len(rts))

	for i := range selfPosts {
		p := selfPosts[i]
		likeCnt, _ := s.posts.CountLikes(ctx, p.ID)
		liked := false
		if viewerID != 0 {
			liked, _ = s.posts.IsLiked(ctx, viewerID, p.ID)
		}
		// agregasi komentar & retweet count
		cmtCnt, _ := s.comments.CountByPost(ctx, p.ID)
		rtCnt, _ := s.retweets.CountByPost(ctx, p.ID)

		timeline = append(timeline, domain.PostView{
			Post:               p,
			AuthorUsername:     authorMap[p.UserID],
			CategoryName:       catMap[p.CategoryID],
			LikeCount:          likeCnt,
			LikedByMe:          liked,
			CommentCount:       cmtCnt,
			RetweetCount:       rtCnt,
			RepostedByUsername: nil,
			QuoteBody:          nil,
		})
	}

	// 4) Normalisasi retweet → PostView (dengan RepostedByUsername & QuoteBody)
	for i := range rts {
		rt := rts[i]
		orig, _ := s.posts.FindByID(ctx, rt.OriginalPostID)
		if orig == nil {
			continue
		}
		authorMap, _ := s.posts.BatchAuthorUsernames(ctx, []uint{orig.UserID, profileUserID})
		catMap, _ := s.posts.BatchCategoryNames(ctx, []uint{orig.CategoryID})
		likeCnt, _ := s.posts.CountLikes(ctx, orig.ID)
		liked := false
		if viewerID != 0 {
			liked, _ = s.posts.IsLiked(ctx, viewerID, orig.ID)
		}
		cmtCnt, _ := s.comments.CountByPost(ctx, orig.ID)
		rtCnt, _ := s.retweets.CountByPost(ctx, orig.ID)

		repostedBy := authorMap[profileUserID]
		pv := domain.PostView{
			Post:           *orig,
			AuthorUsername: authorMap[orig.UserID],
			CategoryName:   catMap[orig.CategoryID],
			LikeCount:      likeCnt,
			LikedByMe:      liked,
			CommentCount:   cmtCnt,
			RetweetCount:   rtCnt,
		}
		pv.RepostedByUsername = &repostedBy
		if rt.QuoteBody != nil && *rt.QuoteBody != "" {
			pv.QuoteBody = rt.QuoteBody
		}
		// Trik kecil: set CreatedAt PostView mengikuti waktu retweet agar sort timeline terasa benar.
		// Kamu bisa tambahkan field ekstra di PostView jika ingin membedakan display timestamp.
		// Untuk MVP, biarkan urutan kita tentukan manual di step sort di bawah.
		timeline = append(timeline, pv)
	}

	// 5) Urutkan timeline terbaru dulu.
	// Untuk akurat, kamu bisa menambahkan "event time" ke PostView (post.CreatedAt vs retweet.CreatedAt).
	// Di MVP ini, kita sort berdasar UpdatedAt paling baru sebagai pendekatan sederhana.
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].UpdatedAt.After(timeline[j].UpdatedAt)
	})

	// 6) Hitung ringkasan profil
	postCount := int64(len(selfPosts))
	retweetCount := int64(len(rts))

	view := &domain.ProfileView{
		UserID:       u.ID,
		Username:     u.Username,
		PostCount:    postCount,
		RetweetCount: retweetCount,
		Posts:        timeline,
	}
	return view, nil
}
