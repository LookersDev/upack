package main

type UniversalPackageMetadata struct {
	Group        string   `json:"group,omitempty"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	IconURL      string   `json:"icon,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}

func (meta UniversalPackageMetadata) BareVersion() string {
	v, err := ParseUniversalPackageVersion(meta.Version)

	if err == nil {
		v.Prerelease = ""
		v.Build = ""
		return v.String()
	}

	return meta.Version
}

func (meta UniversalPackageMetadata) groupAndName() string {
	if meta.Group != "" {
		return meta.Group + "/" + meta.Name
	}
	return meta.Name
}
