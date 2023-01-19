package commands

type AutoCompleteChoice struct {
	Name              string             `json:"name"`
	NameLocalizations *map[string]string `json:"name_localizations,omitempty"`
	Value             interface{}        `json:"value"`
}
