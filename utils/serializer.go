package utils

import (
	"bufio"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// keep domain data
type domain struct {
	OverallScore int `json:"Overall Score"`
	Facets       map[string]int
}

// keep profile data
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

// Initialize function
func (s *BigFiveResultsTextSerializer) Initialize(text string) {
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

	// read profile data
	scanner := bufio.NewScanner(strings.NewReader(s.profile))

	// read lines
	for scanner.Scan() {
		line = scanner.Text()

		// get user name
		if strings.HasPrefix(line, "This report compares") {
			words := strings.Fields(line)
			user = words[3]
		} else if isDomainScoreLine(line) {
			// get overall score data
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
			// get facet and score data
			facetName, facetScore := getNameScore(line)
			facets[facetName] = facetScore
		}
	}

	// push previous saved state after loop
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
	// make json string and replace : with => to match requested format
	return strings.Replace(string(jsonStr), ":", "=>", -1)
}

// determine if line is domain score line
func isDomainScoreLine(line string) bool {
	match, _ := regexp.MatchString("^[^.]+\\.+\\d+$", line)
	return match
}

// determine if line is facet score line
func isFacetScoreLine(line string) bool {
	match, _ := regexp.MatchString("^\\.+[^.]+\\.+\\d+$", line)
	return match
}

// get domain or facet name and score out of line
func getNameScore(line string) (string, int) {
	r1, _ := regexp.Compile("[^.\\d]+")
	name := r1.FindString(line)

	r2, _ := regexp.Compile("\\d+")
	score, _ := strconv.Atoi(r2.FindString(line))

	return name, score
}
