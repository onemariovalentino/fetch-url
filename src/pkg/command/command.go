package command

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/onemariovalentino/fetch-webpage/src/pkg/di"
	"github.com/spf13/cobra"
)

type (
	Command struct {
		comm *cobra.Command
	}
)

var metadataFlag bool

func New() *Command {
	services := di.New()

	fetchCommand := &cobra.Command{
		Use: "fetch",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "fetch [--metadata] {url}...",
		},
		Short: "Fetch web page or metadata based on url",
		RunE: func(cmd *cobra.Command, args []string) error {
			cx := context.Background()
			if !metadataFlag {
				err := services.FetchUsecase.DownloadPage(cx, args)
				if err != nil {
					return err
				}
			} else {
				meta, err := services.FetchUsecase.GetMetadata(cx, args[0])
				if err != nil {
					return err
				}
				fmt.Printf("site: %s\n", meta.URL)
				fmt.Printf("num_links: %d\n", meta.NumLinks)
				fmt.Printf("images: %d\n", meta.NumImages)
				fmt.Printf("last_fetch: %s\n", meta.LastFetch.Format("Mon Jan 02 2006 15:04 MST"))
			}
			return nil
		},
	}
	fetchCommand.Flags().BoolVarP(&metadataFlag, "metadata", "m", false, "fetch metadata only")
	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(fetchCommand)
	return &Command{comm: rootCmd}
}

func (c *Command) Run() {
	if err := c.comm.Execute(); err != nil {
		log.Fatalf("err:%s", err)
		os.Exit(2)
	}
}
