// Usage: go run scripts/update-inngest.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	root := findRoot()
	os.Chdir(root)

	updateGo()
	updatePython()
	updateTypeScript()

	fmt.Println("==> Done")
}

func updateGo() {
	fmt.Println("==> Updating Go")
	runIn("go", "go", "get", "-u", "github.com/inngest/inngestgo@latest")
	runIn("go", "go", "mod", "tidy")
}

func updatePython() {
	fmt.Println("==> Updating Python")

	version := fetchLatestPyPIVersion("inngest")
	fmt.Printf("    Latest version: %s\n", version)

	dirs, _ := filepath.Glob("py-*")
	for _, dir := range dirs {
		updatePyprojectTOML(filepath.Join(dir, "pyproject.toml"), version)
	}

	run("uv", "sync")
}

func updateTypeScript() {
	fmt.Println("==> Updating TypeScript")
	run("pnpm", "update", "inngest", "@inngest/*", "--recursive", "--latest")
}

func updatePyprojectTOML(path, version string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	content := string(data)

	// Update inngest==X.Y.Z
	re := regexp.MustCompile(`"inngest==([^"]*)"`)
	content = re.ReplaceAllString(content, fmt.Sprintf(`"inngest==%s"`, version))

	// Update inngest[connect]==X.Y.Z
	reConnect := regexp.MustCompile(`"inngest\[connect\]==([^"]*)"`)
	content = reConnect.ReplaceAllString(content, fmt.Sprintf(`"inngest[connect]==%s"`, version))

	os.WriteFile(path, []byte(content), 0644)
}

func fetchLatestPyPIVersion(pkg string) string {
	resp, err := http.Get(fmt.Sprintf("https://pypi.org/pypi/%s/json", pkg))
	if err != nil {
		fatal("Failed to fetch PyPI info: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
	}
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Info.Version
}

func run(name string, args ...string) {
	runIn("", name, args...)
}

func runIn(dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fatal("Command failed: %s %s\n%v", name, strings.Join(args, " "), err)
	}
}

func findRoot() string {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)

	// If running with "go run", use working directory
	if strings.Contains(exe, "go-build") {
		dir, _ = os.Getwd()
	}

	// Walk up to find scripts directory or root
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		if _, err := os.Stat(filepath.Join(dir, "scripts")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			fatal("Could not find project root")
		}
		dir = parent
	}
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
