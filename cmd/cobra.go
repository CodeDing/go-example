package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var echoTimes int
	var cfgPath string

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
		       For many years people have printed back to the screen.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
			   Echo works a lot like print, except it has a child command.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo Print: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [-t times] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
		       a count and a string.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("len(args) = %d\n", len(args))
			fmt.Printf("echoTimes = %d\n", echoTimes)
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, ", "))
			}
		},
	}

	var cmdPath = &cobra.Command{
		Use:   "path [-k8s-cfg]",
		Short: "Echo absolute path",
		Long:  "echo k8s config path",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}
			fmt.Printf("k8s absolute path is %s\n", cfgPath)
		},
	}

	kubeCfg := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 2, "times to echo the input")
	//cfgpath := cmd.Flags().String("k8s-cfg", kubeCfg, "kubernetes config")
	cmdPath.Flags().StringVar(&cfgPath, "k8s-cfg", kubeCfg, "kubernetes config")

	//var rootCmd = &cobra.Command{Use: "app"}
	var rootCmd = &cobra.Command{Use: "appXXX"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes, cmdPath)
	rootCmd.Execute()
}
