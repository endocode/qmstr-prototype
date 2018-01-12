package analysis

import (
	"os/exec"
	"log"
	"encoding/json"
	"bytes"
	"bufio"
	"regexp"
)

type ScancodeAnalyzer struct {
	cmd string
	cmdargs []string
	result map[string]interface{}
}
func NewScancodeAnalyzer() *ScancodeAnalyzer {
	sc := ScancodeAnalyzer{"scancode",[]string{"--license", "--quiet"}, map[string]interface{}{}}
	return &sc
}

func (sc *ScancodeAnalyzer) analyzeDirectorySources(sourcefile string) error {
	log.Printf("scanning with scancode %s", sourcefile)
	licenses := []string{}
	cmd := exec.Command(sc.cmd,  append(sc.cmdargs, sourcefile)...)
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	cmd.Stdout = writer
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Could not read stdout of scancode: %v", err)
	}

	re := regexp.MustCompile("{.+")
	jsonBytes := re.Find(buf.Bytes())
	var data Data
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatalf("JSON parsing failed: %v\n", err)
	}
	l := data.Files[0].Licenses

	for idx := range l {
		licenses = append(licenses, l[idx].License)
	}

	sc.result["licenses"] = licenses
	log.Printf("Found the following licenses with scancode: %v", licenses)
	return nil
}

func (sc *ScancodeAnalyzer) GetName() string {
	return "Scancode analyzer"
}

func (sc *ScancodeAnalyzer) Analyze(a Analyzable) error {
	err := sc.analyzeDirectorySources(a.GetFile())
	if err != nil {
		return err
	}

	a.StoreResult(sc.result)

	return nil
}