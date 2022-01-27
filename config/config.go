package config

type Split struct {
	Review     uint8 `json:"review,string"`
	Contribute uint8 `json:"contribute,string"`
}

type IgnoreFiles struct {
	Names []string `json:"ignored"`
}

type DeterminationConfig struct {
	Split   Split       `json:"split"`
	Ignored IgnoreFiles `json:"ignored"`
}
