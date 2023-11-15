package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "path/filepath"

  "github.com/spf13/cobra"
)

func main() {
  var rootCommand = &cobra.Command{}
  var projectName, projectPath string

  var cmd = &cobra.Command{
    Use:   "create",
    Short: "Create boilerplate for a new project",
    Run: func(cmd *cobra.Command, args []string) {
      // validations
      if projectName == "" {
        fmt.Println("You must supply a project name.")
        return
      }
      if projectPath == "" {
        fmt.Println("You must supply a project path.")
        return
      }
      fmt.Println("Creating project...")

      //set global path
      globalPath := filepath.Join(projectPath, projectName)

      // check if project directory already exists
      if _, err := os.Stat(globalPath); err == nil {
        fmt.Println("Project directory already exists.")
        return
      }

      // create project directory by project name
      // Create main project directory
      if err := os.Mkdir(globalPath, os.ModePerm); err != nil {
        log.Fatal(err)
      }

      // start golang project and create go.mod
      startGo := exec.Command("go", "mod", "init", projectName)
      startGo.Dir = globalPath
      startGo.Stdout = os.Stdout
      startGo.Stderr = os.Stderr
      err := startGo.Run()
      if err != nil {
        log.Fatal(err)
      }

      // create directory structure
      // create cmd
      cmdPath := filepath.Join(globalPath, "cmd")
      if err := os.Mkdir(cmdPath, os.ModePerm); err != nil {
        log.Fatal(err)
      }
      // create internal
      internalPath := filepath.Join(globalPath, "internal")
      if err := os.Mkdir(internalPath, os.ModePerm); err != nil {
        log.Fatal(err)
      }
      // create handler
      handlerPath := filepath.Join(internalPath, "handler")
      if err := os.Mkdir(handlerPath, os.ModePerm); err != nil {
        log.Fatal(err)
      }
      // create routes
      routesPath := filepath.Join(handlerPath, "routes")
      fmt.Println(routesPath)
      if err := os.Mkdir(routesPath, os.ModePerm); err != nil {
        log.Fatal(err)
      }

      // create files
      // create main.go
      mainPath := filepath.Join(cmdPath, "main.go")
      mainFile, err := os.Create(mainPath)
      if err != nil {
        log.Fatal(err)
      }
      defer mainFile.Close()
      if err := WriteMainFile(mainPath); err != nil {
        log.Fatal(err)
      }

      // create routes.go
      routesFilePath := filepath.Join(routesPath, "routes.go")
      routesFile, err := os.Create(routesFilePath)
      if err != nil {
        log.Fatal(err)
      }
      defer routesFile.Close()
      if err := WriteRoutesFile(routesFilePath); err != nil {
        log.Fatal(err)
      }
    },
  }

  cmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
  cmd.Flags().StringVarP(&projectPath, "path", "p", "", "Path where the project will be created")

  rootCommand.AddCommand(cmd)
  rootCommand.Execute()
}

func WriteMainFile(mainPath string) error {
  // code to write in routes.go
  packageContent := []byte(`package main

import "fmt"

func main() {
  fmt.Println("Hello World!")
}
`)

  mainFile, err := os.OpenFile(mainPath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer mainFile.Close()

  // write in main.go
  _, err = mainFile.Write(packageContent)
  if err != nil {
    return err
  }

  return nil
}

func WriteRoutesFile(routesFilePath string) error {
  // code to write in routes.go
  packageContent := []byte(`package routes

// Seu código aqui
`)

  routesFile, err := os.OpenFile(routesFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer routesFile.Close()

  // write in routes.go
  _, err = routesFile.Write(packageContent)
  if err != nil {
    return err
  }

  return nil
}
