package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type Problem struct {
	Number     string
	Title      string
	Difficulty string
	Tags       string
	URL        string
	Dir        string
}

const problemsDir = "problems"
const readmePath = "README.md"

const solutionTemplate = `// Problem: {{.Number}}. {{.Title}}
// Difficulty: {{.Difficulty}}
// Tags:
// URL: https://leetcode.com/problems/{{.Slug}}/
// Date: {{.Date}}

package {{.Package}}

func solution() {
	// TODO: implement
}
`

const testTemplate = `package {{.Package}}

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		// TODO: define fields
	}{
		{name: "example1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: implement
		})
	}
}
`

var metaPatterns = map[string]*regexp.Regexp{
	"number":     regexp.MustCompile(`^// Problem:\s*(\d+)\.\s*(.+)`),
	"difficulty": regexp.MustCompile(`^// Difficulty:\s*(\w+)`),
	"tags":       regexp.MustCompile(`^// Tags:\s*(.*)`),
	"url":        regexp.MustCompile(`^// URL:\s*(.*)`),
}

func parseMeta(path string) (*Problem, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := &Problem{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "//") {
			break
		}
		if m := metaPatterns["number"].FindStringSubmatch(line); m != nil {
			p.Number = m[1]
			p.Title = strings.TrimSpace(m[2])
		}
		if m := metaPatterns["difficulty"].FindStringSubmatch(line); m != nil {
			p.Difficulty = strings.TrimSpace(m[1])
		}
		if m := metaPatterns["tags"].FindStringSubmatch(line); m != nil {
			p.Tags = strings.TrimSpace(m[1])
		}
		if m := metaPatterns["url"].FindStringSubmatch(line); m != nil {
			p.URL = strings.TrimSpace(m[1])
		}
	}
	return p, scanner.Err()
}

func collectProblems() ([]Problem, error) {
	dirs, err := os.ReadDir(problemsDir)
	if err != nil {
		return nil, err
	}

	var problems []Problem
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		mainGo := filepath.Join(problemsDir, d.Name(), "main.go")
		p, err := parseMeta(mainGo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warn: skipping %s: %v\n", d.Name(), err)
			continue
		}
		p.Dir = d.Name()
		problems = append(problems, *p)
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})
	return problems, nil
}

func generateReadme(problems []Problem) error {
	input, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	content := string(input)

	// Build stats
	counts := map[string]int{"Easy": 0, "Medium": 0, "Hard": 0}
	for _, p := range problems {
		counts[p.Difficulty]++
	}

	// Build table rows
	var rows strings.Builder
	for _, p := range problems {
		link := fmt.Sprintf("[%s](%s/)", p.Title, filepath.Join(problemsDir, p.Dir))
		rows.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", p.Number, link, p.Difficulty, p.Tags))
	}

	// Replace stats
	statsTable := fmt.Sprintf(
		"| Difficulty | Count |\n|------------|-------|\n| Easy       | %-5d |\n| Medium     | %-5d |\n| Hard       | %-5d |",
		counts["Easy"], counts["Medium"], counts["Hard"],
	)
	statsRe := regexp.MustCompile(`(?s)\| Difficulty \| Count \|.*?\| Hard\s*\|[^\n]*`)
	content = statsRe.ReplaceAllString(content, statsTable)

	// Replace problems table
	beginMarker := "<!-- AUTO-GENERATED: DO NOT EDIT BELOW -->"
	endMarker := "<!-- AUTO-GENERATED: END -->"
	tableHeader := "| # | Title | Difficulty | Tags |\n|---|-------|------------|------|\n"
	re := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(beginMarker) + `.*?` + regexp.QuoteMeta(endMarker))
	replacement := beginMarker + "\n" + tableHeader + rows.String() + endMarker
	content = re.ReplaceAllString(content, replacement)

	return os.WriteFile(readmePath, []byte(content), 0644)
}

type scaffoldData struct {
	Number     string
	Title      string
	Difficulty string
	Slug       string
	Package    string
	Date       string
}

func scaffold(num int) error {
	slug := fmt.Sprintf("%04d-new-problem", num)
	dir := filepath.Join(problemsDir, slug)

	if _, err := os.Stat(dir); err == nil {
		return fmt.Errorf("directory already exists: %s", dir)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	pkg := strings.ReplaceAll(slug, "-", "")

	data := scaffoldData{
		Number:     fmt.Sprintf("%d", num),
		Title:      "New Problem",
		Difficulty: "Medium",
		Slug:       "new-problem",
		Package:    pkg,
		Date:       "YYYY-MM-DD",
	}

	// main.go
	tmpl := template.Must(template.New("sol").Parse(solutionTemplate))
	mainFile, err := os.Create(filepath.Join(dir, "main.go"))
	if err != nil {
		return err
	}
	defer mainFile.Close()
	if err := tmpl.Execute(mainFile, data); err != nil {
		return err
	}

	// main_test.go
	testTmpl := template.Must(template.New("test").Parse(testTemplate))
	testFile, err := os.Create(filepath.Join(dir, "main_test.go"))
	if err != nil {
		return err
	}
	defer testFile.Close()
	if err := testTmpl.Execute(testFile, data); err != nil {
		return err
	}

	// go.mod
	modContent := fmt.Sprintf("module %s\n\ngo 1.22\n", pkg)
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(modContent), 0644); err != nil {
		return err
	}

	fmt.Printf("created: %s\n", dir)
	fmt.Println("→ タイトル・難易度・タグ・URLを main.go のコメントに記入してください")
	return nil
}

func main() {
	newNum := flag.Int("new", 0, "scaffold a new problem by number")
	flag.Parse()

	if *newNum > 0 {
		if err := scaffold(*newNum); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	problems, err := collectProblems()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := generateReadme(problems); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("README updated: %d problems\n", len(problems))
}
