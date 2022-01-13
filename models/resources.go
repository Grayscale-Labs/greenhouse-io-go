package models

// Smaller resource models that are sub-fields of larger resources.

import "time"

type Address struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Answer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Attachment struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Type     string `json:"type"`
}

type Education struct {
	ID         int       `json:"id"`
	SchoolName string    `json:"school_name"`
	Degree     string    `json:"degree"`
	Discipline string    `json:"discipline"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type EmailAddress struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Employee struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Name       string  `json:"name"`
	EmployeeID *string `json:"employee_id"`
}

type Employment struct {
	ID          int       `json:"id"`
	CompanyName string    `json:"company_name"`
	Title       string    `json:"title"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	Address string `json:"address"`
}

type PhoneNumber struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type ProspectDetail struct {
	ProspectPool  interface{} `json:"prospect_pool"`
	ProspectStage interface{} `json:"prospect_stage"`
	ProspectOwner interface{} `json:"prospect_owner"`
}

type SocialMediaAddress struct {
	Value string `json:"value"`
}

type Source struct {
	ID         int    `json:"id"`
	PublicName string `json:"public_name"`
}

type Stage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WebsiteAddress struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
