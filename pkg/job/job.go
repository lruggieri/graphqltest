package job

type Job struct {
	ID             int      `json:"id"`
	Position       string   `json:"position"`
	CompanyID      int      `json:"companyID"`
	Description    string   `json:"description"`
	SkillsRequired []string `json:"skillsRequired"`
	Location       string   `json:"location"`
	EmploymentType string   `json:"employmentType"`
}