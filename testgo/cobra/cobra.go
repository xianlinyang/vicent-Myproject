package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
type command struct {
	// global configuration
	root             *cobra.Command
	globalConfig     *viper.Viper
	globalConfigFile string
	homeDir          string

}
func (c *command) initCheckCmd() (err error) {


	cmd := &cobra.Command{
		Use:   "check",
		Short: "boson.integration tests on a Aurora cluster",
		Long:  `runs integration tests on a Aurora cluster.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println("sdf")
			return nil
		},
	}
	c.root.AddCommand(cmd)
	return nil
}

func getcommand()(c *command){
	c = &command{
		root: &cobra.Command{
			Use:           "aurorakeeper",
			Short:         "Guass-Project/Aurora Keeper",
			SilenceErrors: true,
			SilenceUsage:  true,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("fd")
				fmt.Println(c.globalConfigFile)
				return nil
			},
		},
	}
	return c
}
func main(){
    //设置argument值为 check --config ssfd
	c := getcommand()
	c.initCheckCmd()
	globalFlags := c.root.PersistentFlags()
	globalFlags.StringVar(&c.globalConfigFile, "config", "", "config file (default is $HOME/.aurorakeeper.yaml)")

	c.root.Execute()

	//var cmdPull = &cobra.Command{
	//	Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	//	Short: "Pull an image or a repository from a registry",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Pull: " + strings.Join(args, " "))
	//	},
	//}
	//
	//var rootCmd = &cobra.Command{Use: "docker"}
	//rootCmd.AddCommand(cmdPull)
	//rootCmd.Execute()
}
