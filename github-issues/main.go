package main

import (
	"encoding/csv"
	"fmt"
	"liamfallon/rough-work/github-issues/types"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records, err
}

func writeCsvFile(filePath string, records [][]string) error {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Unable to create file "+filePath, err)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	err = csvWriter.WriteAll(records)
	if err != nil {
		log.Fatal("Unable to write file as CSV for "+filePath, err)
	}

	return err
}

func main() {
	types.IssueMap = make(map[string]*types.GithubRecord)

	records, err := readCsvFile("/Users/liam/work/PorchIssues/porch-issues.csv")
	if err != nil {
		return
	}

	for i := 1; i < len(records); i++ {
		ghr := types.NewGithubRecord(records[i])

		if !strings.Contains(ghr.GetLabels(), "area/porch") {
			continue
		}

		if ghr.GetState() != "closed" {
			continue
		}

		storedGhr, found := types.IssueMap[ghr.IssueUrl()]
		if found {
			if err := storedGhr.Merge(ghr); err != nil {
				fmt.Println("merge failed " + err.Error())
			}
		} else if !ghr.IsBlank() {
			types.IssueMap[ghr.IssueUrl()] = ghr
			storedGhr = ghr
		}
	}

	var outRecords [][]string

	outRecords = append(outRecords, []string{"title", "body", "labels"})

	for _, issue := range types.IssueMap {
		outRecords = append(outRecords, issue.OutStringArray())
	}

	writeCsvFile("/Users/liam/work/PorchIssues/OutClosedIssues.csv", outRecords)

	fmt.Printf("%d issues extracted\n", len(types.IssueMap))
}
