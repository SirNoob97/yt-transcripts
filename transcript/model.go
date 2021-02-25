package transcript

// Name ...
type Name struct {
	SimpleText string `json:"simpleText"`
}

// CaptionTrack ...
type CaptionTrack struct {
	BaseURL        string `json:"baseUrl"`
	Name           Name   `json:"name"`
	VssID          string `json:"vssId"`
	LanguageCode   string `json:"language_code"`
	Kind           string `json:"kind,omitempty"`
	IsTranslatable bool   `json:"isTranslatable"`
}

// Caption ...
type Caption struct {
	CaptionTracks CaptionTrack `json:"captionTracks"`
}
