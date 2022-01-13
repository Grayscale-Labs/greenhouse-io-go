package models

type RejectionReason struct {
	ID   int                 `json:"id"`
	Name string              `json:"name"`
	Type RejectionReasonType `json:"type"`
}

type RejectionReasonType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RejectionDetails struct {
	CustomFields      map[string]interface{}            `json:"custom_fields"`
	KeyedCustomFields map[string]map[string]interface{} `json:"keyed_custom_fields"`
}
