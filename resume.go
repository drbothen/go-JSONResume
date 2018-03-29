package JSONResume

import (
	"os"
	"encoding/json"
)

// TODO: Add godoc style documentation
type Resume struct {
	Basics struct {
		Name     string `json:"name"`
		Label    string `json:"label"`
		Picture  string `json:"picture"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Website  string `json:"website"`
		Summary  string `json:"summary"`
		Location struct {
			Address     string `json:"address"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
			CountryCode string `json:"countryCode"`
			Region      string `json:"region"`
		} `json:"location"`
		Profiles []struct {
			Network  string `json:"network"`
			Username string `json:"username"`
			URL      string `json:"url"`
		} `json:"profiles"`
	} `json:"basics"`
	Work []struct {
		Company    string   `json:"company"`
		Position   string   `json:"position"`
		Website    string   `json:"website"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		Summary    string   `json:"summary"`
		Highlights []string `json:"highlights"`
	} `json:"work"`
	Volunteer []struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		Website      string   `json:"website"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	} `json:"volunteer"`
	Education []struct {
		Institution string   `json:"institution"`
		Area        string   `json:"area"`
		StudyType   string   `json:"studyType"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Gpa         string   `json:"gpa"`
		Courses     []string `json:"courses"`
	} `json:"education"`
	Awards []struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	} `json:"awards"`
	Publications []struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		Website     string `json:"website"`
		Summary     string `json:"summary"`
	} `json:"publications"`
	Skills []struct {
		Name     string   `json:"name"`
		Level    string   `json:"level"`
		Keywords []string `json:"keywords"`
	} `json:"skills"`
	Languages []struct {
		Name  string `json:"name"`
		Level string `json:"level"`
	} `json:"languages"`
	Interests []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"interests"`
	References []struct {
		Name      string `json:"name"`
		Reference string `json:"reference"`
	} `json:"references"`
}

// TODO: Add godoc style documentation
func LoadResumeFile( path string ) (Resume, error) {
	// Open Resume stored in .json file
	jr, err := os.Open(path)
	// TODO: add custom cant open file error?
	if err != nil {
		return Resume{}, err
	}
	var r Resume // instantiate a Resume struct
	// Pares the json into the Resume struct
	if err := json.NewDecoder(jr).Decode(&r); err != nil {
		return Resume{}, err
	}
	return r, err
}