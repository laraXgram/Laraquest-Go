package updates

type ReactionCount struct {
    Type *ReactionType `json:"type"`
    Total_count int64 `json:"total_count"`
}