package handler

import "fmt"

func errParamisRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create opening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Link == "" && r.Salary <= 0 && r.Remote == nil && r.Location == "" && r.Company == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamisRequired("role", "string")
	}
	if r.Company == "" {
		return errParamisRequired("company", "string")
	}
	if r.Location == "" {
		return errParamisRequired("location", "string")
	}
	if r.Link == "" {
		return errParamisRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamisRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamisRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Link != "" || r.Location != "" || r.Salary > 0 || r.Remote != nil {
		return nil
	}
	// If none of the fields were not provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
