package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jwalton/gchalk"
)

// colors
var green = gchalk.BrightGreen
var blue = gchalk.BrightBlue

// flags
var Lowercase bool
var OwO bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Yet another neofetch clone",

	Run: func(cmd *cobra.Command, args []string) {
		var h strings.Builder

		fmt.Print(blue(Username()), "@", blue(Hostname()), "\n")

		if OwO {
			h.WriteString(green("OwOS"))
		} else {
			h.WriteString(green("OS"))
		}
		write(&h, ": ", Distro(), "\n")

		write(&h, green("Arch"), ": ", Arch(), "\n")
		write(&h, green("Kernel"), ": ", Kernel(), "\n")
		write(&h, green("Shell"), ": ", Shell(), "\n")
		write(&h, green("DE/WM"), ": ", Desktop())

		// this probably negates any performance benefits from using
		// strings.Builder but I'm doing it because I can then just
		// lowercase/owoify the entire text at once :P
		text := h.String()

		if Lowercase {
			text = strings.ToLower(text)
		}

		if OwO {
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
