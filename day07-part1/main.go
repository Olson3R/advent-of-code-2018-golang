package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Step struct {
	name          string
	requiredSteps []string
}

func (s *Step) Add(stepName string) {
	s.requiredSteps = append(s.requiredSteps, stepName)
}

func (s Step) RequirementsMet(processedSteps string) bool {
	if len(s.requiredSteps) == 0 {
		return true
	}
	for _, stepName := range s.requiredSteps {
		if !strings.Contains(processedSteps, stepName) {
			return false
		}
	}
	return true
}

type Steps struct {
	steps map[string]*Step
}

func (s *Steps) Add(name string) *Step {
	step, ok := s.steps[name]
	if !ok {
		step = &Step{name, []string{}}
		s.steps[name] = step
	}
	return step
}

func (s Steps) AddRequiredStep(parentName string, childName string) {
	s.Add(parentName)
	childStep := s.Add(childName)
	childStep.Add(parentName)
}

func (s Steps) Process() string {
	orderedSteps := s.OrderedSteps()
	var processedSteps bytes.Buffer

	for len(orderedSteps) > 0 {
		fmt.Println(len(orderedSteps))
		for i, stepName := range orderedSteps {
			step := s.steps[stepName]
			if step.RequirementsMet(processedSteps.String()) {
				processedSteps.WriteString(stepName)
				orderedSteps = append(orderedSteps[:i], orderedSteps[i+1:]...)
				break
			}
		}
	}

	return processedSteps.String()
}

func (s Steps) OrderedSteps() []string {
	orderedSteps := make([]string, len(s.steps))

	i := 0
	for key := range s.steps {
		orderedSteps[i] = key
		i++
	}
	sort.Strings(orderedSteps)
	return orderedSteps
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRow(row string) (parent string, child string) {
	re := regexp.MustCompile("Step (\\w) must be finished before step (\\w) can begin.")
	parts := re.FindStringSubmatch(row)
	return parts[1], parts[2]
}

func main() {
	file, err := os.Open("input/d07-input.txt")
	check(err)
	defer file.Close()

	steps := Steps{make(map[string]*Step)}

	// parse parent/child steps
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parentName, childName := parseRow(row)
		steps.AddRequiredStep(parentName, childName)
	}

	processedSteps := steps.Process()

	fmt.Println("STEPS", processedSteps)
	fmt.Println("Done")
}
