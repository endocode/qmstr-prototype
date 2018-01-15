package main

import (
	"bytes"
	"log"
	model "qmstr-prototype/qmstr/qmstr-model"
	"strings"
	"text/template"
)

type report struct {
	SPDXVersion, DataLicense, Name string
	License                        string
}

// CreateReport renders an SPDX document for the given TargetEntity
func CreateReport(toolName string, target model.TargetEntity) string {

	//Define the template
	const reportTemplate = "SPDXVersion: {{.SPDXVersion}}\\nDataLicense: {{.DataLicense}}\\nPackageName: {{.Name}}\\nPackageLicenseDeclared: {{.License}}"
	//Create a new template and parse the data
	r := template.Must(template.New("report").Parse(reportTemplate))

	licenses := extractLicenses(toolName, target.Sources)

	// TODO: if licenses[0] == nil
	report := report{"SPDX-2.0", "CCO-1.0", target.Name, strings.Join(licenses, " AND ")}

	//Execute the template
	b := bytes.Buffer{}
	err := r.Execute(&b, report)
	if err != nil {
		log.Println("Failed to render report template:", err)
	}
	return b.String()
}

func extractLicenses(toolName string, sources []string) []string {
	licenseSet := map[string]struct{}{}
	for _, v := range sources {
		s, err := Model.GetSourceEntity(v)
		if err != nil {
			return []string{}
		}
		if s.Licenses == nil || len(s.Licenses) == 0 {
			// Find corresponding target entity
			t, err := Model.GetTargetEntityByPath(v)
			if err != nil {
				return []string{}
			}
			for _, source := range t.Sources {
				ts, err := Model.GetSourceEntity(source)
				if err != nil {
					return []string{}
				}
				for t, licenses := range ts.Licenses {
					if t == toolName {
						for _,license := range licenses {
							licenseSet[license] = struct{}{}
						}
					}
				}
			}
		} else {
			for t, licenses := range s.Licenses {
				if t == toolName {
					for _,license := range licenses {
						licenseSet[license] = struct{}{}
					}
				}
			}
		}
	}
	license := []string{}
	for k := range licenseSet {
		license = append(license, k)
	}
	return license
}
