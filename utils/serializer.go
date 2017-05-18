package utils

import (
	"bufio"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

type domain struct {
	OverallScore int `json:"Overall Score"`
	Facets       map[string]int
}

type personality struct {
	Name                 string `json:"NAME"`
	Extraversion         domain `json:"EXTRAVERSION"`
	Agreeableness        domain `json:"AGREEABLENESS"`
	Conscientiousness    domain `json:"CONSCIENTIOUSNESS"`
	Neuroticism          domain `json:"NEUROTICISM"`
	OpennessToExperience domain `json:"OPENNESS TO EXPERIENCE"`
}

// BigFiveResultsTextSerializer class
type BigFiveResultsTextSerializer struct {
	profile string
}

// Read function
func (s *BigFiveResultsTextSerializer) Read(text string) {
	s.profile = text
}

// Hash function
func (s BigFiveResultsTextSerializer) Hash() string {

	facets := make(map[string]int)

	d := domain{}

	p := personality{}
	var user, domainName, line string
	var domainScore int

	isFirstRun := true

	scanner := bufio.NewScanner(strings.NewReader(s.profile))

	for scanner.Scan() {

		line = scanner.Text()
		if strings.HasPrefix(line, "This report compares") {
			words := strings.Fields(line)
			user = words[3]
		} else if isDomainScoreLine(line) {

			if !isFirstRun {
				d.OverallScore = domainScore
				d.Facets = facets

				switch domainName {
				case "EXTRAVERSION":
					p.Extraversion = d
				case "AGREEABLENESS":
					p.Agreeableness = d
				case "CONSCIENTIOUSNESS":
					p.Conscientiousness = d
				case "NEUROTICISM":
					p.Neuroticism = d
				case "OPENNESS TO EXPERIENCE":
					p.OpennessToExperience = d
				}

				d = domain{}
				facets = make(map[string]int)
			}

			domainName, domainScore = getNameScore(line)

			isFirstRun = false

		} else if isFacetScoreLine(line) {
			facetName, facetScore := getNameScore(line)
			facets[facetName] = facetScore
		}
	}

	d.OverallScore = domainScore
	d.Facets = facets

	switch domainName {
	case "EXTRAVERSION":
		p.Extraversion = d
	case "AGREEABLENESS":
		p.Agreeableness = d
	case "CONSCIENTIOUSNESS":
		p.Conscientiousness = d
	case "NEUROTICISM":
		p.Neuroticism = d
	case "OPENNESS TO EXPERIENCE":
		p.OpennessToExperience = d
	}

	p.Name = user

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	jsonStr, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jsonStr))
	return strings.Replace(string(jsonStr), ":", "=>", -1)
}

func isDomainScoreLine(line string) bool {
	match, _ := regexp.MatchString("^[^.]+\\.+\\d+$", line)
	return match
}

func isFacetScoreLine(line string) bool {
	match, _ := regexp.MatchString("^\\.+[^.]+\\.+\\d+$", line)
	return match
}

func getNameScore(line string) (string, int) {
	r1, _ := regexp.Compile("[^.\\d]+")
	name := r1.FindString(line)

	r2, _ := regexp.Compile("\\d+")
	score, _ := strconv.Atoi(r2.FindString(line))

	return name, score
}
