package main

import (
	"github.com/gorook/rook/fs"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rook",
	Short: "build site",
	RunE:  buildSite,
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "run development server",
	Aliases: []string{"s", "serve"},
	RunE:    startServer,
}

var newSiteCmd = &cobra.Command{
	Use:   "new [site]",
	Short: "create new rook site",
	Args:  cobra.ExactArgs(1),
	RunE:  createNewSite,
}

var addContentCmd = &cobra.Command{
	Use:   "add [page]",
	Short: "add new content page",
	Args:  cobra.ExactArgs(1),
	RunE:  addNewContent,
}

func init() {
	cmd.AddCommand(serverCmd)
	cmd.AddCommand(newSiteCmd)
	cmd.AddCommand(addContentCmd)
}

func buildSite(c *cobra.Command, args []string) error {
	filesys := fs.New(afero.NewOsFs(), afero.NewOsFs())
	a := newApplication(filesys)
	err := a.init()
	if err != nil {
		return err
	}
	return a.build()
}

func startServer(c *cobra.Command, args []string) error {
	return nil
}

func createNewSite(c *cobra.Command, args []string) error {
	return nil
}

func addNewContent(c *cobra.Command, args []string) error {
	return nil
}
