package cmd

import (
	"fmt"
	"github.com/heyujiang/user/config"
	"github.com/heyujiang/user/server/biz"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "user server",
	Long:  "user server",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println(config.GetConfig())

	biz.StartUserBiz(config.GetConfig().UserBiz)

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}
