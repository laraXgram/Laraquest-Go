package updates

type WriteAccessAllowed struct {
    From_request bool `json:"from_request"`
    Web_app_name string `json:"web_app_name"`
    From_attachment_menu bool `json:"from_attachment_menu"`
}