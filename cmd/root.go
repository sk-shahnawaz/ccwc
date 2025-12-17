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

		if calculateBytes {
			calculateCharacters, calculateWords, calculateLines = false, false, false

		} else if !calculateCharacters && !calculateWords && !calculateLines {
			// Step 5
			// If no explicit flags are set, default to -c -l -w
			calculateCharacters, calculateWords, calculateLines = true, true, true
			calculateBytes = false
		}

		var (
			content  []byte
			filePath string
		)

		if len(args) == 0 {
			// Step 5
			stdInContent, err := io.ReadAll(os.Stdin)

			if err != nil {
				return err
			}

			content = stdInContent
		} else {
			filePath = args[0]
			_, err := os.Stat(filePath)

			if err != nil {
				return err
			}

			content, err = os.ReadFile(filePath)

			if err != nil {
				return err
			}
		}

		var (
			charactersCount, wordsCount, linesCount, bytesCount = 0, 0, 0, 0
		)

		if calculateBytes {
			// Step 1: Calculate number of bytes
			bytesCount = len(content)
		}

		if calculateLines {
			// Step 2: Calculate number of lines
			textContent := string(content)
			linesCount = len(strings.Split(textContent, "\n"))
		}

		if calculateWords {
			// Step 3: Calculate number of words
			textContent := string(content)
			wordsCount = len(strings.Fields(textContent))
		}

		if calculateCharacters {
			// Step 4: Calculate number of characters
			// utf8.RuneCountInString() counts the actual number of Unicode characters (runes).
			charactersCount = utf8.RuneCountInString(string(content))
		}

		fmt.Printf("%s %s", buildOutputBasedOnFlag(charactersCount, wordsCount, linesCount, bytesCount), filePath)

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

func buildOutputBasedOnFlag(charactersCount int, wordsCount int, linesCount int, bytesCount int) string {
	var sb strings.Builder

	if calculateCharacters {
		sb.WriteString(strconv.Itoa(charactersCount))
		sb.WriteString("\t")
	}

	if calculateWords {
		sb.WriteString(strconv.Itoa(wordsCount))
		sb.WriteString("\t")
	}

	if calculateLines {
		sb.WriteString(strconv.Itoa(linesCount))
		sb.WriteString("\t")
	}

	if calculateBytes {
		sb.WriteString(strconv.Itoa(bytesCount))
		sb.WriteString("\t")
	}

	return sb.String()
}
