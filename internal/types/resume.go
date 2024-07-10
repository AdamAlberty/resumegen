package types

type PersonalInfo struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	LinkedIn string `json:"linked_in"`
	GitHub   string `json:"github"`
	Website  string `json:"website"`
}

type WorkExperience struct {
	Company   string   `json:"company"`
	Role      string   `json:"role"`
	Location  string   `json:"location"`
	TimeFrame string   `json:"time_frame"`
	Bullets   []string `json:"bullets"`
}

type Education struct {
	Institution    string   `json:"institution"`
	Degree         string   `json:"degree"`
	Location       string   `json:"location"`
	GraduationYear string   `json:"graduation_year"`
	Bullets        []string `json:"bullets"`
}

type ResumeData struct {
	PersonalInfo   PersonalInfo     `json:"personal_info"`
	WorkExperience []WorkExperience `json:"work_experience"`
	Education      []Education      `json:"education"`
	Certifications string           `json:"certifications"`
	Skills         string           `json:"skills"`
	Interests      string           `json:"interests"`
}
