package analysis

type Data struct {
	Notice string `json:"scancode_notice"`
	Version string `json:"scancode_version"`
	Options Options `json:"scancode_options"`
	Files_count int `json:"files_count"`
	Files []Files `json:"files"`
}

type Files struct {
	Path string `json:"path"`
	Scan_errors []string `json:"scan_errors"`
	Licenses []Licenses `json:"licenses"`
}

type Licenses struct{
	Key string `json:"key"`
	Score float32 `json:"score"`
	S_name string `json:"short_name"`
	Category string `json:"category"`
	Owner string `json:"owner"`
	Homepage_url string `json:"homepage_url"`
	Text_url string `json:"text_url"`
	Reference_url string `json:"reference_url"`
	License string `json:"spdx_license_key"`
	Spdx_url string `json:"spdx_url"`
	Start_line int `json:"start_line"`
	End_line int `json:"end_line"`
	Matched_rule Macher `json:"matched_rule"`
}

type Options struct {
	License bool `json:"--license"`
	Score int `json:"--license-score"`
	Format string `json:"--format"`
}

type Macher struct {
	Identifier string `json:"identifier"`
	License_choice bool `json:"license_choice"`
	Licenses []string `json:"licenses"`
}