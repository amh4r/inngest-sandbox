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

	"github.com/spf13/cobra"
)

func updateInngestCmd() *cobra.Command {
	var (
		only []string
		skip []string
	)

	cmd := &cobra.Command{
		Use:   "update-inngest",
		Short: "Update Inngest dependencies across all example projects",
		Long: `Update Inngest dependencies across all example projects.

By default, updates all languages (go, python, typescript).
Use --only to update specific languages, or --skip to exclude languages.

Examples:
  sandbox update-inngest                    # Update all
  sandbox update-inngest --only go,ts       # Update only Go and TypeScript
  sandbox update-inngest --skip python      # Update all except Python`,
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := findProjectRoot()
			if err != nil {
				return err
			}
			os.Chdir(root)

			languages := resolveLanguages(only, skip)

			if languages["go"] {
				if err := updateGo(); err != nil {
					return err
				}
			}

			if languages["python"] {
				if err := updatePython(); err != nil {
					return err
				}
			}

			if languages["typescript"] {
				if err := updateTypeScript(); err != nil {
					return err
				}
			}

			fmt.Println("==> Done")
			return nil
		},
	}

	cmd.Flags().StringSliceVar(&only, "only", nil, "Only update these languages (go, python/py, typescript/ts)")
	cmd.Flags().StringSliceVar(&skip, "skip", nil, "Skip these languages (go, python/py, typescript/ts)")

	return cmd
}

func resolveLanguages(only, skip []string) map[string]bool {
	languages := map[string]bool{
		"go":         true,
		"python":     true,
		"typescript": true,
	}

	// Normalize aliases
	normalize := func(lang string) string {
		switch strings.ToLower(lang) {
		case "py":
			return "python"
		case "ts":
			return "typescript"
		default:
			return strings.ToLower(lang)
		}
	}

	// If --only is specified, start with nothing enabled
	if len(only) > 0 {
		languages = map[string]bool{
			"go":         false,
			"python":     false,
			"typescript": false,
		}
		for _, lang := range only {
			languages[normalize(lang)] = true
		}
	}

	// Apply --skip
	for _, lang := range skip {
		languages[normalize(lang)] = false
	}

	return languages
}

func updateGo() error {
	fmt.Println("==> Updating Go")
	if err := runIn("go", "go", "get", "-u", "github.com/inngest/inngestgo@latest"); err != nil {
		return err
	}
	return runIn("go", "go", "mod", "tidy")
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

	return run("uv", "sync")
}

func updateTypeScript() error {
	fmt.Println("==> Updating TypeScript")
	return run("pnpm", "update", "inngest", "@inngest/*", "--recursive", "--latest")
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

func run(name string, args ...string) error {
	return runIn("", name, args...)
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
		// Look for markers that indicate project root
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
