package types

type (
	Options []Option

	Option struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	OptionsRes Options
)
