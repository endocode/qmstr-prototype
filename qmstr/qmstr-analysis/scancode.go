package analysis

import (
	"fmt"
	"errors"
)

type PrescanScanCodeAnalyzer struct {
	ScanData map[string]interface{}
}

func (la *PrescanScanCodeAnalyzer) Configure(data map[string]interface{}) error {
	if value, ok := data["scancode"]; ok && value != nil{
		la.ScanData = data["scancode"].(map[string]interface{})
		return nil
	}

	return errors.New("Scancode is not initialized")
}

func (la *PrescanScanCodeAnalyzer) Analyze(a Analyzable) error {
	filename := a.GetFile()
	licenses := []string{}
	copyholders := []string{}
	authors := []string{}
	for _, file := range la.ScanData["files"].([]interface{}) {
		fileData := file.(map[string]interface{})
		if fileData["path"] == filename {
			fmt.Printf("Found %s", filename)
			for _, license := range fileData["licenses"].([]interface{}) {
				licenses = append(licenses, license.(map[string]interface{})["spdx_license_key"].(string))
			}
			for _, copyright := range fileData["copyrights"].([]interface{}) {
				holders := copyright.(map[string]interface{})
				for _, holder := range holders["holders"].([]interface{}) {
					copyholders = append(copyholders, holder.(string))
				}
				auths := copyright.(map[string]interface{})
				for _, auth := range auths["authors"].([]interface{}){
					authors = append(authors, auth.(string))
				}
			}
		}
	}
	result := map[string]interface{}{}
	result["licenses"] = licenses
	result["copyholders"] = copyholders
	result["authors"] = authors
	a.StoreResult(la.GetName(), result)
	return nil
}

func (la *PrescanScanCodeAnalyzer) GetName() string {
	return "PrescanScanCode Analyzer"
}
