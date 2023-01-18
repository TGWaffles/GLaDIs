package components

type ModalSubmitData struct {
	CustomId   string             `json:"custom_id"`
	Components []MessageComponent `json:"components"`
}
