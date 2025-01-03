package updates

type Poll struct {
    Id string `json:"id"`
    Question string `json:"question"`
    Question_entities *[]MessageEntity `json:"question_entities"`
    Options *[]PollOption `json:"options"`
    Total_voter_count int64 `json:"total_voter_count"`
    Is_closed bool `json:"is_closed"`
    Is_anonymous bool `json:"is_anonymous"`
    Type string `json:"type"`
    Allows_multiple_answers bool `json:"allows_multiple_answers"`
    Correct_option_id int64 `json:"correct_option_id"`
    Explanation string `json:"explanation"`
    Explanation_entities *[]MessageEntity `json:"explanation_entities"`
    Open_period int64 `json:"open_period"`
    Close_date int64 `json:"close_date"`
}