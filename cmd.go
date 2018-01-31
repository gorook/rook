package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rook",
	Short: "build site",
	Run: func(c *cobra.Command, args []string) {
		a := newApplication()
		a.buildSite()
	},
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "run development server",
	Aliases: []string{"s", "serve"},
	Run: func(c *cobra.Command, args []string) {
		a := newApplication(renderToMemory)
		a.startServer()
	},
}

var newSiteCmd = &cobra.Command{
	Use:   "new [site]",
	Short: "create new rook site",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		a := newApplication()
		a.createNewSite(args[0])
	},
}

var addContentCmd = &cobra.Command{
	Use:   "add [page]",
	Short: "add new content page",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		a := newApplication()
		a.addNewContent(args[0])
	},
}

func init() {
	cmd.AddCommand(serverCmd)
	cmd.AddCommand(newSiteCmd)
	cmd.AddCommand(addContentCmd)
}
