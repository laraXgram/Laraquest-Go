package params

type GetUpdatesParams struct { 
    Offset *any `json:"offset"`
    Limit *any `json:"limit"`
    Timeout *any `json:"timeout"`
    Allowed_updates *any `json:"allowed_updates"`
}