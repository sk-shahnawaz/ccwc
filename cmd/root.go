package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var (
	calculateCharacters bool
	calculateWords      bool
	calculateLines      bool
	calculateBytes      bool
)

var rootCommand = &cobra.Command{
	Use:   "ccwc",
	Short: "CLI utility to calculate words, lines, characters and bytes",
	Long:  "CLI utility to calculate words, lines, characters and bytes",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Determine which counts to calculate
		if calculateBytes {
			calculateCharacters, calculateWords, calculateLines = false, false, false
		} else if !calculateCharacters && !calculateWords && !calculateLines {
			// If no explicit flags are set, default to -c -l -w
			calculateCharacters, calculateWords, calculateLines = true, true, true
			calculateBytes = false
		}

		// Read content from file or stdin
		content, filePath, err := readInput(args)
		if err != nil {
			return err
		}

		// Calculate requested metrics
		counts := calculateCounts(content)

		// Build and print output
		fmt.Printf("%s %s", buildOutputBasedOnFlag(counts), filePath)

		return nil
	},
}

func init() {
	rootCommand.Flags().BoolVarP(&calculateCharacters, "characters", "c", false, "Prints the number of characters")
	rootCommand.Flags().BoolVarP(&calculateWords, "words", "w", false, "Prints the number of words")
	rootCommand.Flags().BoolVarP(&calculateLines, "lines", "l", false, "Prints the number of lines")
	rootCommand.Flags().BoolVarP(&calculateBytes, "bytes", "b", false, "Prints the number of bytes")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing ccwc '%s'\n", err)
		os.Exit(1)
	}
}

// counts holds the calculated metrics for the input
type counts struct {
	characters int
	words      int
	lines      int
	bytes      int
}

// readInput reads content from a file or stdin
func readInput(args []string) ([]byte, string, error) {
	if len(args) == 0 {
		content, err := io.ReadAll(os.Stdin)
		return content, "", err
	}

	filePath := args[0]
	if _, err := os.Stat(filePath); err != nil {
		return nil, "", err
	}

	content, err := os.ReadFile(filePath)
	return content, filePath, err
}

// calculateCounts calculates all metrics from the content
func calculateCounts(content []byte) counts {
	c := counts{
		bytes: len(content),
	}

	// Convert to string once for all text operations
	textContent := string(content)

	if calculateLines {
		// Count newlines (matches standard wc -l behavior)
		c.lines = strings.Count(textContent, "\n")
	}

	if calculateWords {
		c.words = len(strings.Fields(textContent))
	}

	if calculateCharacters {
		c.characters = utf8.RuneCountInString(textContent)
	}

	return c
}

// buildOutputBasedOnFlag formats the output based on active flags
func buildOutputBasedOnFlag(c counts) string {
	var sb strings.Builder

	if calculateCharacters {
		sb.WriteString(strconv.Itoa(c.characters))
		sb.WriteString("\t")
	}

	if calculateWords {
		sb.WriteString(strconv.Itoa(c.words))
		sb.WriteString("\t")
	}

	if calculateLines {
		sb.WriteString(strconv.Itoa(c.lines))
		sb.WriteString("\t")
	}

	if calculateBytes {
		sb.WriteString(strconv.Itoa(c.bytes))
		sb.WriteString("\t")
	}

	return sb.String()
}
