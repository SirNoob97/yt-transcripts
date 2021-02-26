package transcript

// PlayerCaptionsTracklistRenderer ...
type PlayerCaptionsTracklistRenderer struct {
	CaptionTracks []CaptionTrack `json:"captionTracks"`
}

// Name ...
type Name struct {
	SimpleText string `json:"simpleText"`
}

// CaptionTrack ...
type CaptionTrack struct {
	BaseURL        string `json:"baseUrl"`
	Name           Name   `json:"name"`
	VssID          string `json:"vssId"`
	LanguageCode   string `json:"languageCode"`
	Kind           string `json:"kind,omitempty"`
	IsTranslatable bool   `json:"isTranslatable"`
}

// Caption ...
type Caption struct {
	PCTR PlayerCaptionsTracklistRenderer `json:"playerCaptionsTracklistRenderer"`
}
