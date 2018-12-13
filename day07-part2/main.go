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

const ASCII_OFFSET = 64
const BASE_TIME = 60

type Step struct {
	name          string
	time          int
	requiredSteps []string
}

func (s *Step) Add(stepName string) {
	s.requiredSteps = append(s.requiredSteps, stepName)
}

func (s *Step) SetTime() {
	s.time = BASE_TIME + int(s.name[0]) - ASCII_OFFSET
}

func (s *Step) Tick() bool {
	s.time--
	return s.time <= 0
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
		step = &Step{name, -1, []string{}}
		step.SetTime()
		s.steps[name] = step
	}
	return step
}

func (s Steps) AddRequiredStep(parentName string, childName string) {
	s.Add(parentName)
	childStep := s.Add(childName)
	childStep.Add(parentName)
}

func (s Steps) Process() int {
	orderedSteps := s.OrderedSteps()
	var processedSteps bytes.Buffer
	workers := []*Step{nil, nil, nil, nil, nil}
	ticks := 0
	for processedSteps.Len() < 26 {
		ticks++
		fmt.Println("TICK", ticks, processedSteps.Len(), len(orderedSteps))
		for i, step := range workers {
			if step == nil {
				step = s.NextStep(orderedSteps, processedSteps.String())
				if step != nil {
					workers[i] = step
					orderedSteps = strings.Replace(orderedSteps, step.name, "", 1)
				}
			}
		}

		for i, step := range workers {
			if step != nil && step.Tick() {
				fmt.Println("WORKER DONE", step.name)
				processedSteps.WriteString(step.name)
				workers[i] = nil
			}
		}
	}

	return ticks
}

func (s Steps) NextStep(orderedSteps string, processedSteps string) *Step {
	for _, stepName := range orderedSteps {
		step := s.steps[string(stepName)]
		if step.RequirementsMet(processedSteps) {
			return step
		}
	}
	return nil
}

func (s Steps) OrderedSteps() string {
	orderedSteps := make([]string, len(s.steps))

	i := 0
	for key := range s.steps {
		orderedSteps[i] = key
		i++
	}
	sort.Strings(orderedSteps)
	return strings.Join(orderedSteps, "")
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
