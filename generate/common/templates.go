package common

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

func ParseTemplates(templateDir, outputDir string, doc any) {
	// Log progress
	parseLog := log.With().Str("template_dir", templateDir).Str("output_dir", outputDir).Logger()
	parseLog.Info().Msg("Generating templates ...")
	parseLog.Debug().Interface("data", doc).Msg("Generating templates with data...")

	// Init template and parse components
	rootTemplate := template.New("")
	rootTemplate.Funcs(template.FuncMap{
		"openapiParamsToGin": openapiParamsToGin,
	})
	rootTemplate.ParseGlob(path.Join(templateDir, "*snippet*.go.tmpl"))

	// Parse and execute pages
	filepath.Walk(templateDir, func(templPath string, info os.FileInfo, err error) error {
		// Check if walk successful
		if err != nil {
			log.Fatal().Err(err).Str("template", templPath).Msg("Failed to walk template directory")
		}

		// Strip template dir prefix
		if templPath == templateDir {
			return nil
		}
		templName := strings.TrimPrefix(templPath, templateDir)[1:] // Trim template dir and path separator

		// Create duplicate in output if directory
		walkLog := log.With().Str("template", templName).Logger()
		walkLog.Debug().Msg("Processing template ...")
		if info.IsDir() {
			walkLog.Debug().Msg("Template is a directory, duplicating in output directory")
			EnsureDirExists(filepath.Join(outputDir, templName))
			return nil
		}

		// Validate file type and name
		if !strings.HasSuffix(templPath, ".go.tmpl") {
			return nil
		}
		templ, err := rootTemplate.Clone()
		if err != nil {
			walkLog.Fatal().Err(err).Msg("Failed to clone root template")
			return err
		}
		templContent, err := os.ReadFile(templPath)
		if err != nil {
			walkLog.Fatal().Err(err).Msg("Failed to read template file")
			return err
		}
		if _, err = templ.Parse(string(templContent)); err != nil {
			walkLog.Fatal().Err(err).Msg("Failed to parse template")
			return err
		}
		outputName := strings.TrimSuffix(templName, ".go.tmpl") + ".go"
		file, err := os.Create(path.Join(outputDir, outputName))
		defer func() {
			file.Sync()
			if fileErr := file.Close(); fileErr != nil {
				walkLog.Fatal().Err(err).Msg("Failed to close parsed page output file")
			}
		}()
		if err != nil {
			walkLog.Fatal().Err(err).Msg("Failed to create parsed page output file")
			return err
		}
		walkLog.Info().Msg("Executing template ...")
		err = templ.Execute(file, doc)
		if err != nil {
			walkLog.Fatal().Err(err).Msg("Failed to execute template")
			return err
		}
		return nil
	})
}

var openapiParamsToGinRegex = regexp.MustCompile(`\{([-_a-zA-Z]+)\}`)

func openapiParamsToGin(input string) string {
	return openapiParamsToGinRegex.ReplaceAllString(input, ":$1")
}
