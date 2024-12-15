package model

type ThreadVariants struct {
	Name      string `json:"name,omitempty"`
	DateAdded string `json:"dateAdded,omitempty"`
}

type Thread struct {
	ThreatName    string           `json:"threatName,omitempty"`
	Category      string           `json:"category,omitempty"`
	Size          int              `json:"size,omitempty"`
	DetectionDate string           `json:"detectionDate,omitempty"`
	Variants      []ThreadVariants `json:"variants,omitempty"`
}
