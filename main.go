/*
Copyright © 2023 Julian Easterling julian@julianscorner.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var imageVersion string

func main() {
	args := os.Args[1:]
	pwd, _ := os.Getwd()
	data := strings.ReplaceAll(fmt.Sprintf("%s:/site", pwd), "\\", "/")
	docker := []string{
		"run",
		"--rm",
		"-it",
		"-v",
		data,
	}

	if strings.Count(strings.ToLower(strings.Join(args, " ")), "server") > 0 {
		port := 1313
		re := regexp.MustCompile("-p [0-9]+")
		p := re.FindAllString(strings.Join(args, " "), -1)
		if len(p) > 0 {
			re = regexp.MustCompile("[0-9]+")
			if re.MatchString(p[0]) {
				port, _ = strconv.Atoi(re.FindAllString(p[0], -1)[0])
			}
		}

		docker = append(docker, "-p")
		docker = append(docker, fmt.Sprintf("%d:1313/tcp", port))

		args = append(args, "--bind 0.0.0.0")
	}

	if len(args) > 0 {
		if args[0] == "--image-version" {
			fmt.Println(imageVersion)
			os.Exit(0)
		}
	}

	docker = append(docker, fmt.Sprintf("dcjulian29/hugo:%s", imageVersion))

	if len(args) > 0 {
		docker = append(docker, args...)
	} else {
		docker = append(docker, "")
	}

	cmd := exec.Command("docker", docker...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("\033[1;31m%s\033[0m\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
