package models

import "time"

type Application struct {
	ID                    int               `json:"id"`
	CandidateID           int               `json:"candidate_id"`
	Prospect              bool              `json:"prospect"`
	AppliedAt             time.Time         `json:"applied_at"`
	RejectedAt            time.Time         `json:"rejected_at"`
	LastActivityAt        time.Time         `json:"last_activity_at"`
	Location              Location          `json:"location"`
	Source                Source            `json:"source"`
	CreditedTo            Employee          `json:"credited_to"`
	RejectionReason       *RejectionReason  `json:"rejection_reason"`
	RejectionDetails      *RejectionDetails `json:"rejection_details"`
	Jobs                  []Job             `json:"jobs"`
	JobPostID             int               `json:"job_post_id,omitempty"`
	Status                string            `json:"status"`
	CurrentStage          Stage             `json:"current_stage"`
	Answers               []Answer          `json:"answers"`
	ProspectiveOffice     interface{}       `json:"prospective_office"`
	ProspectiveDepartment interface{}       `json:"prospective_department"`
	ProspectDetail        ProspectDetail    `json:"prospect_detail"`
	Attachments           []Attachment      `json:"attachments"`
}
