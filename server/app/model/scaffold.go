package model

type (
	Database struct {
		Name      string `json:"name"`
		Character string `json:"character"`
	}

	Table struct {
		Name   string `json:"name"`
		Engine string `json:"engine"`
	}

	Column struct {
		Name         string `json:"name"`
		Length       int    `json:"length"`
		Type         string `json:"type"`
		Nullable     bool   `json:"nullable"`
		DefaultValue string `json:"default_value"`
		Comment      string `json:"comment"`
	}

	Fields []Field

	Field struct {
		Name        string `json:"name"`
		JsonName    string `json:"json_name"`
		Translation string `json:"translation"`
		Type        string `json:"type"`
		Nullable    bool   `json:"nullable"`
		Default     string `json:"default"`
		Comment     string `json:"comment"`
		Hidden      bool   `json:"hidden"`
	}
)
