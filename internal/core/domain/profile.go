package domain

// ProfileView: data untuk halaman profil user.
// Biasanya butuh pagination di layer handler (meta total/limit/offset).
type ProfileView struct {
	UserID       uint   `json:"user_id"`
	Username     string `json:"username"`
	PostCount    int64  `json:"post_count"`
	RetweetCount int64  `json:"retweet_count"`
	// items dapat berisi campuran PostView (post asli) dan RetweetView yang
	// di-"normalized" ke PostView (lihat field RepostedByUsername/QuoteBody).
	Posts []PostView `json:"posts"`
}
