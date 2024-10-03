package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.2"

type FileInfo struct {
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
	Path    string    `json:"path"`
}

var (
	inputDir  string
	recursive bool
	exclude   string
	include   string
	rootCmd   = &cobra.Command{
		Use:   "concatenator",
		Short: "A tool to concatenate file information into a JSON file",
		Long:  `Concatenator is a CLI tool that traverses a directory, collects file information, and outputs it as a structured JSON file.`,
	}
	concatenateCmd = &cobra.Command{
		Use:   "concatenate [output_file]",
		Short: "Concatenate file information into a JSON file",
		Long:  `Traverse a directory, collect file information, and output it as a structured JSON file.`,
		Args:  cobra.MaximumNArgs(1),
		Run:   run,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of concatenator",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("concatenator version %s\n", VERSION)
		},
	}
)

func init() {
	concatenateCmd.Flags().StringVarP(&inputDir, "dir", "d", ".", "Input directory")
	concatenateCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Traverse directory recursively")
	concatenateCmd.Flags().StringVarP(&exclude, "exclude", "e", "", "Exclude files matching pattern (comma-separated, supports wildcards)")
	concatenateCmd.Flags().StringVarP(&include, "include", "i", "", "Include only files matching pattern (comma-separated, supports wildcards)")
	concatenateCmd.Flags().Lookup("recursive").NoOptDefVal = "true"
	rootCmd.AddCommand(concatenateCmd, versionCmd)
}

func matchesPattern(path string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}
	}
	return false
}

func run(cmd *cobra.Command, args []string) {
	outputFile := "output.json"
	if len(args) > 0 {
		outputFile = args[0]
	}

	excludePatterns := strings.Split(exclude, ",")
	includePatterns := strings.Split(include, ",")
	for i, pattern := range excludePatterns {
		excludePatterns[i] = strings.TrimSpace(pattern)
	}
	for i, pattern := range includePatterns {
		includePatterns[i] = strings.TrimSpace(pattern)
	}

	var files []FileInfo
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !recursive && info.IsDir() && path != inputDir {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if len(includePatterns) > 0 && !matchesPattern(path, includePatterns) {
			return nil
		}

		if matchesPattern(path, excludePatterns) {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(inputDir, path)
		fileInfo := FileInfo{
			Name:    info.Name(),
			Content: string(content),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			Path:    relPath,
		}

		files = append(files, fileInfo)
		return nil
	})

	if err != nil {
		fmt.Printf("Error while traversing directory: %v\n", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(files, "", "  ")
	if err != nil {
		fmt.Printf("Error while creating JSON: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error while writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("JSON file created successfully: %s\n", outputFile)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
