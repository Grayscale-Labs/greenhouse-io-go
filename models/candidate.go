package models

import "time"

type Candidate struct {
	ID                   int                               `json:"id"`
	FirstName            string                            `json:"first_name"`
	LastName             string                            `json:"last_name"`
	Company              string                            `json:"company"`
	Title                string                            `json:"title"`
	CreatedAt            time.Time                         `json:"created_at"`
	UpdatedAt            time.Time                         `json:"updated_at"`
	LastActivity         time.Time                         `json:"last_activity"`
	IsPrivate            bool                              `json:"is_private"`
	PhotoURL             string                            `json:"photo_url"`
	Attachments          []Attachment                      `json:"attachments"`
	ApplicationIDs       []int                             `json:"application_ids"`
	PhoneNumbers         []PhoneNumber                     `json:"phone_numbers"`
	Addresses            []Address                         `json:"addresses"`
	EmailAddresses       []EmailAddress                    `json:"email_addresses"`
	WebsiteAddresses     []WebsiteAddress                  `json:"website_addresses"`
	SocialMediaAddresses []SocialMediaAddress              `json:"social_media_addresses"`
	Recruiter            Employee                          `json:"recruiter"`
	Coordinator          Employee                          `json:"coordinator"`
	CanEmail             bool                              `json:"can_email"`
	Tags                 []string                          `json:"tags"`
	Applications         []Application                     `json:"applications"`
	Educations           []Education                       `json:"educations"`
	Employments          []Employment                      `json:"employments"`
	LinkedUserIds        []int                             `json:"linked_user_ids"`
	CustomFields         map[string]interface{}            `json:"custom_fields"`
	KeyedCustomFields    map[string]map[string]interface{} `json:"keyed_custom_fields"`
}
