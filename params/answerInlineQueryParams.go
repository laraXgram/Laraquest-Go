package params

type AnswerInlineQueryParams struct { 
    Inline_query_id any `json:"inline_query_id"`
    Results any `json:"results"`
    Cache_time *any `json:"cache_time"`
    Is_personal *any `json:"is_personal"`
    Next_offset *any `json:"next_offset"`
    Button *any `json:"button"`
}