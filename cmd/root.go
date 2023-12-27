package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// from https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
// also see https://en.wikipedia.org/wiki/ANSI_escape_code
const Reset = "\033[0m"
const Green = "\033[92m"
const Blue = "\033[94m"

// flags
var Lowercase bool
var OwO bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Yet another neofetch clone",

	Run: func(cmd *cobra.Command, args []string) {
		var h strings.Builder

		fmt.Print(Blue, Username(), Reset, "@", Blue, Hostname(), Reset, "\n")

		h.WriteString(Green)
		if OwO == true {
			h.WriteString("OwOS")
		} else {
			h.WriteString("OS")
		}
		write(&h, Reset, ": ", Distro(), "\n")

		write(&h, Green, "Arch", Reset, ": ", Arch(), "\n")
		write(&h, Green, "Kernel", Reset, ": ", Kernel(), "\n")
		write(&h, Green, "Shell", Reset, ": ", Shell(), "\n")
		write(&h, Green, "DE/WM", Reset, ": ", Desktop())

		// this probably negates any performance benefits from using
		// strings.Builder but I'm doing it because I can then just
		// lowercase/owoify the entire text at once :P
		text := h.String()

		if Lowercase == true {
			text = strings.ToLower(text)
		}

		if OwO == true {
			text = OwOify(text)
		}

		fmt.Println(text)
	},
}

// convenient way to add multiple strings to a
// strings.Builder at once
func write(b *strings.Builder, text ...string) {

	for _, t := range text {
		b.WriteString(t)
	}

}

func OwOify(text string) string {
	return strings.NewReplacer(
		"r", "w",
		"l", "w",
		"R", "W",
		"L", "W").Replace(text)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Lowercase, "lowercase", "l", false, "makes the text lowercase")
	rootCmd.PersistentFlags().BoolVarP(&OwO, "owo", "o", false, "makes the text mowe owo")
}
