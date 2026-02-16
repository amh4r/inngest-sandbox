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

type UpdateInngestCmd struct {
	Only []string `short:"o" help:"Only update these languages (go, py, ts)"`
	Skip []string `short:"s" help:"Skip these languages (go, py, ts)"`
}

func (c *UpdateInngestCmd) Run() error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}
	os.Chdir(root)

	languages := c.resolveLanguages()

	if languages["go"] {
		if err := updateGo(); err != nil {
			return err
		}
	}

	if languages["py"] {
		if err := updatePython(); err != nil {
			return err
		}
	}

	if languages["ts"] {
		if err := updateTypeScript(); err != nil {
			return err
		}
	}

	fmt.Println("==> Done")
	return nil
}

func (c *UpdateInngestCmd) resolveLanguages() map[string]bool {
	languages := map[string]bool{
		"go": true,
		"py": true,
		"ts": true,
	}

	if len(c.Only) > 0 {
		for k := range languages {
			languages[k] = false
		}
		for _, lang := range c.Only {
			languages[lang] = true
		}
	}

	for _, lang := range c.Skip {
		languages[lang] = false
	}

	return languages
}

func updateGo() error {
	fmt.Println("==> Updating Go")

	if err := runIn("./go", "go", "get", "-u", "github.com/inngest/inngestgo@latest"); err != nil {
		return err
	}

	if err := runIn("./go", "go", "mod", "tidy"); err != nil {
		return err
	}

	return runInRoot("go", "work", "vendor")
}

func updatePython() error {
	fmt.Println("==> Updating Python")

	version, err := fetchLatestPyPIVersion("inngest")
	if err != nil {
		return err
	}
	fmt.Printf("    Latest version: %s\n", version)

	dirs, _ := filepath.Glob("py-*")
	for _, dir := range dirs {
		updatePyprojectTOML(filepath.Join(dir, "pyproject.toml"), version)
	}

	return runInRoot("uv", "sync", "--all-packages")
}

func updateTypeScript() error {
	fmt.Println("==> Updating TypeScript")
	return runInRoot("pnpm", "update", "inngest", "@inngest/*", "--recursive", "--latest")
}

func updatePyprojectTOML(path, version string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	content := string(data)

	re := regexp.MustCompile(`"inngest==([^"]*)"`)
	content = re.ReplaceAllString(content, fmt.Sprintf(`"inngest==%s"`, version))

	reConnect := regexp.MustCompile(`"inngest\[connect\]==([^"]*)"`)
	content = reConnect.ReplaceAllString(content, fmt.Sprintf(`"inngest[connect]==%s"`, version))

	os.WriteFile(path, []byte(content), 0644)
}

func fetchLatestPyPIVersion(pkg string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://pypi.org/pypi/%s/json", pkg))
	if err != nil {
		return "", fmt.Errorf("failed to fetch PyPI info: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("failed to parse PyPI response: %w", err)
	}
	return data.Info.Version, nil
}

func runInRoot(name string, args ...string) error {
	return runIn("./", name, args...)
}

func runIn(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "pnpm-workspace.yaml")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root")
		}
		dir = parent
	}
}
