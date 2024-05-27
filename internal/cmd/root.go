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

var lowercase bool
var OwO bool
var noAscii bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Yet another minimal system info *fetch ",

	Run: func(cmd *cobra.Command, args []string) {
		format()
	},
}

const cat = `
            %s
 ╱|、       %s
(˚ˎ 。7     %s
 |、˜〵     %s    
 じしˍ,)ノ  %s
            %s
`

func format() {
	var sb strings.Builder

	switch noAscii {
	case false:
		text := fmt.Sprintf(cat,
			usernameAndHostname(),
			distroLine(), archLine(), kernelLine(), shellLine(), desktopLine())
		sb.WriteString(text)
	case true:
		sb.WriteString(usernameAndHostname() + "\n")
		sb.WriteString(distroLine() + "\n")
		sb.WriteString(archLine() + "\n")
		sb.WriteString(kernelLine() + "\n")
		sb.WriteString(shellLine() + "\n")
		sb.WriteString(desktopLine())
	}

	switch lowercase {
	case true:
		text := strings.ToLower(sb.String())
		fmt.Println(text)
	case false:
		fmt.Println(sb.String())
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
	rootCmd.PersistentFlags().BoolVarP(&lowercase, "lowercase", "l", false,
		"makes the text lowercase")
	rootCmd.PersistentFlags().BoolVarP(&OwO, "owo", "o", false,
		"makes the text mowe owo")
	rootCmd.PersistentFlags().BoolVarP(&noAscii, "no-ascii", "n", false,
		"whether to not show the ascii art cat",
	)
}
