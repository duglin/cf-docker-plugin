package main

import (
    "os"
    "os/exec"
    "github.com/cloudfoundry/cli/plugin"
)

type DockerPlugin struct{}

func (c *DockerPlugin) Run(cliConnection plugin.CliConnection, args []string) {
    cmd := exec.Command("docker", args[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    err := cmd.Run()
    if err != nil {
        panic( err )
    }
}

func (c *DockerPlugin) GetMetadata() plugin.PluginMetadata {
    return plugin.PluginMetadata{
        Name: "DockerPlugin",
        Commands: []plugin.Command{
            {
                Name:     "docker",
                HelpText: "Try: cf docker help ",
            },
        },
    }
}

func main() {
    plugin.Start(new(DockerPlugin))
}
