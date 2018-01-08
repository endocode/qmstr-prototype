package main

import (
	"os"
	"encoding/csv"
	"os/exec"
	"log"
	"qmstr-prototype/qmstr/qmstr-model"
)

// analyze sources with Ninka for all the Target Entities
func useNinka(targets []model.TargetEntity) map[string][]string {
	analysis := make(map[string][]string)
	for _, target := range targets {
		var s model.SourceEntity
		s.Path = target.Sources[0]
		s.Hash = "filehash"

		licenses := []string{}
		cmd := exec.Command("ninka", "-i", s.Path)
		err := cmd.Start()
		checkErr(err)
		if err := cmd.Wait(); err != nil {
			log.Fatalf("License analysis failed for %s", s.Path)
		}

		licenseFile, err := os.Open(s.Path + ".license")
		checkErr(err)
		r := csv.NewReader(licenseFile)
		r.Comma = ';'
		records, err := r.ReadAll()
		checkErr(err)

		for _, fields := range records {
			if len(fields) > 0 {
				licenses = append(licenses, fields[0])
			}
		}
		s.Licenses = licenses
		analysis[target.Name] = licenses
	}

	return analysis
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}