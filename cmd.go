package main

import (
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

var listen string

func init() {
	serverCmd.Flags().StringVarP(&listen, "listen", "l", "localhost:1414", "Address to listen on")

	cmd.AddCommand(serverCmd)
	cmd.AddCommand(newSiteCmd)
	cmd.AddCommand(addContentCmd)
}

func buildSite(c *cobra.Command, args []string) error {
	a := newApplication(appDefault)
	err := a.init("")
	if err != nil {
		return err
	}
	return a.build()
}

func startServer(c *cobra.Command, args []string) error {
	a := newApplication(appRenderToMemory)
	err := a.init(listen)
	if err != nil {
		return err
	}
	return a.startServer(listen)
}

func createNewSite(c *cobra.Command, args []string) error {
	a := newApplication(appDefault)
	return a.createSite(args[0])
}

func addNewContent(c *cobra.Command, args []string) error {
	a := newApplication(appDefault)
	return a.createPost(args[0])
}
