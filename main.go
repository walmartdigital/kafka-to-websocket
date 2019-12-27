package main

import (
	"github.com/spf13/cobra"
)

const topic = "topic-A"

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "kafka-poc",
		Short: "...",
		Long:  "...",
	}

	produceCmd = &cobra.Command{
		Use:   "produce",
		Short: "...",
		Long:  "...",
		Args:  cobra.ExactArgs(1),
		Run:   produceCommand,
	}

	consumeCmd = &cobra.Command{
		Use:   "consume",
		Short: "...",
		Long:  "...",
		Args:  cobra.ExactArgs(1),
		Run:   consumeCommand,
	}
)

func main() {
	rootCmd.Execute()
}

func produceCommand(cmd *cobra.Command, args []string) {
	mainProducer(args[0])
}

func consumeCommand(cmd *cobra.Command, args []string) {
	mainConsumer(args[0])
}

func init() {
	rootCmd.AddCommand(produceCmd)
	rootCmd.AddCommand(consumeCmd)
}
