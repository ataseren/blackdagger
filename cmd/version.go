package cmd

import (
	"fmt"

	"github.com/ErdemOzgen/blackdagger/internal/config"
	"github.com/ErdemOzgen/blackdagger/internal/constants"
	"github.com/spf13/cobra"
)

var AsciiArt = `

██████╗░██╗░░░░░░█████╗░░█████╗░██╗░░██╗██████╗░░█████╗░░██████╗░░██████╗░███████╗██████╗░
██╔══██╗██║░░░░░██╔══██╗██╔══██╗██║░██╔╝██╔══██╗██╔══██╗██╔════╝░██╔════╝░██╔════╝██╔══██╗
██████╦╝██║░░░░░███████║██║░░╚═╝█████═╝░██║░░██║███████║██║░░██╗░██║░░██╗░█████╗░░██████╔╝
██╔══██╗██║░░░░░██╔══██║██║░░██╗██╔═██╗░██║░░██║██╔══██║██║░░╚██╗██║░░╚██╗██╔══╝░░██╔══██╗
██████╦╝███████╗██║░░██║╚█████╔╝██║░╚██╗██████╔╝██║░░██║╚██████╔╝╚██████╔╝███████╗██║░░██║
╚═════╝░╚══════╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝╚═════╝░╚═╝░░╚═╝░╚═════╝░░╚═════╝░╚══════╝╚═╝░░╚═╝              
`

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display the binary version",
		Long:  `blackdagger version`,
		PreRun: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(config.LoadConfig())
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(AsciiArt)
			fmt.Println(constants.Version)
		},
	}
}
