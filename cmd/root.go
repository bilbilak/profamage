package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	app "github.com/bilbilak/profamage/config"
	"github.com/bilbilak/profamage/internal"
)

var (
	Help    bool
	Version bool
	License bool
)

var rootCmd = &cobra.Command{
	Use:   strings.ToLower(app.Name),
	Short: "Profile Image Generator",
	Long:  app.Name + ` is a CLI tool for generating profile images that fit perfectly within various social media platforms.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if Help {
			_ = cmd.Help()
			return
		}

		if Version {
			fmt.Println(app.Version)
			return
		}

		if License {
			fmt.Println(app.License)
			return
		}

		internal.Help()
	},
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().BoolVar(&Help, "help", false, "Show usage instructions")

	rootCmd.Flags().BoolVar(&Version, "version", false, "Display the installed version number")
	rootCmd.Flags().BoolVar(&License, "license", false, "Display the license name")

	rootCmd.Flags().StringVarP(&internal.Color, "color", "c", "transparent", "")
	rootCmd.Flags().IntVarP(&internal.Size, "size", "s", 512, "")
	rootCmd.Flags().IntVarP(&internal.Padding, "padding", "p", 10, "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		internal.FatalError(err)
	}
}
