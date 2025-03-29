package models

type LeaderboardEntry struct {
	Username     string `json:"username"`
	TotalScore   int    `json:"total_score"`
	QuizzesTaken int    `json:"quizzes_taken"`
	Rank         int    `json:"rank"`
}

type LeaderboardResponse struct {
	Top10    []LeaderboardEntry `json:"top_10,omitempty"`
	UserRank *LeaderboardEntry  `json:"user_rank,omitempty"`
}
