package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get user home directory: \n")
		fmt.Printf("无法获取用户目录：\n")
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	nuPath     := flag.String("nushell", "nu.exe", "")
	configPath := flag.String("config", filepath.Join(userHome, ".config"), "")
	cachePath  := flag.String("cache", filepath.Join(userHome, ".cache"), "")
	dataPath   := flag.String("data", filepath.Join(userHome, ".local", "share"), "")

	flag.Parse()

	extraArgs := flag.Args()

	cmd := exec.Command(*nuPath, extraArgs...)

	cmd.Env = os.Environ()

	_, existConfigHome := os.LookupEnv("XDG_CONFIG_HOME")
	if !existConfigHome {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("XDG_CONFIG_HOME=%s", *configPath),
			"NUX_INJECTED_CONFIG=1",
		)
	}

	_, existCacheHome := os.LookupEnv("XDG_CACHE_HOME")
	if !existCacheHome {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("XDG_CACHE_HOME=%s", *cachePath),
			"NUX_INJECTED_CACHE=1",
		)
	}

	_, existDataHome := os.LookupEnv("XDG_DATA_HOME")
	if !existDataHome {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("XDG_DATA_HOME=%s", *dataPath),
			"NUX_INJECTED_DATA=1",
		)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Printf("Nushell start failed: \n")
		fmt.Printf("Nushell 启动失败：\n")
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
