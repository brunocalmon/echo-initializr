package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/brunocalmon/echo-initializr/data"
	"github.com/brunocalmon/echo-initializr/logic"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a basic project using echo framework",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		version, _ := cmd.Flags().GetString("version")
		outputDir, _ := cmd.Flags().GetString("outputDir")
		dependencies, _ := cmd.Flags().GetString("dependencies")
		port, _ := cmd.Flags().GetInt("port")
		files := data.Initializr(namespace, "clean_architecture")

		logic.CreateFiles(namespace, version, outputDir, port, files)
		installAllDependencies(namespace, outputDir, dependencies)

		fmt.Println("New project successfuly created at: " + outputDir + "/" + namespace)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	HOME_ENV := os.Getenv("HOME")

	var version string = checkGoVersion()

	initCmd.Flags().StringP("namespace", "n", "github.com/example/sample", "Set your project's name")
	initCmd.Flags().StringP("version", "v", version, "Set your project's version")
	initCmd.Flags().StringP("outputDir", "o", HOME_ENV+"/echo_initializr", "Set the output directory to your project.")
	initCmd.Flags().StringP("dependencies", "d", "", "Set the dependencies of your project.")
	initCmd.Flags().IntP("port", "p", 8080, "Set the port of your project webserver.")
}

func installAllDependencies(namespace, outputDir string, dependencies string) {
	installDependence(outputDir+"/"+namespace, "github.com/labstack/echo/v4")

	if dependencies != "" {
		splitted := strings.Split(dependencies, ",")
		for _, dependence := range splitted {
			installDependence(outputDir+"/"+namespace, dependence)
		}
	}
}

func installDependence(dir, dependence string) {
	cmd := exec.Command("go", "get", dependence)
	cmd.Dir = dir

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		panic("go get dependence error:" + string(err.Error()))
	}
}

func checkGoVersion() string {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		panic("go version not runnig")
	}
	complete_version := strings.Split(string(out), " ")[2][2:]
	version := strings.Split(complete_version, ".")[0] + "." + strings.Split(complete_version, ".")[1]
	return version
}
