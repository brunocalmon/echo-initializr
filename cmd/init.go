package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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
		structureData := make(map[string]string)

		structureData["namespace"], _ = cmd.Flags().GetString("namespace")
		structureData["version"], _ = cmd.Flags().GetString("version")
		structureData["outputDir"], _ = cmd.Flags().GetString("outputDir")
		structureData["dependencies"], _ = cmd.Flags().GetString("dependencies")

		port, _ := cmd.Flags().GetInt("port")
		files := data.Initializr(structureData["namespace"], "clean_architecture")

		structureData["port"] = strconv.Itoa(port)

		logic.CreateFiles(structureData, files)
		installAllDependencies(structureData["namespace"], structureData["outputDir"], structureData["dependencies"])

		fmt.Println("New project successfuly created at: " + structureData["outputDir"] + "/" + structureData["namespace"])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	HomeEnv := os.Getenv("HOME")

	var version string = checkGoVersion()

	initCmd.Flags().StringP("namespace", "n", "github.com/example/sample", "Set your project's namespace")
	initCmd.Flags().StringP("version", "v", version, "Set your project's go version, avoid use this flag.")
	initCmd.Flags().StringP("outputDir", "o", HomeEnv+"/echo_initializr", "Set the output directory to your project.")
	initCmd.Flags().StringP("dependencies", "d", "", "Set the dependencies of your project.")
	initCmd.Flags().IntP("port", "p", 8080, "Set the port of your project webserver.")
}

func installAllDependencies(namespace, outputDir string, dependencies string) {
	installDependence(outputDir+"/"+namespace, "github.com/labstack/echo/v4/...")

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
	completeVersion := strings.Split(string(out), " ")[2][2:]
	version := strings.Split(completeVersion, ".")[0] + "." + strings.Split(completeVersion, ".")[1]
	return version
}
