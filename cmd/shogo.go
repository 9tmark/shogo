// Copyright (C) 2022 The shogo Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	Domain  = ""
	Key     = ""
	Version = "0.0.0"
)

func main() {
	verbose := false
	if len(os.Args) > 2 {
		for _, arg := range os.Args {
			if arg == "--verbose" {
				verbose = true
			}
		}
	}

	if verbose {
		fmt.Printf("shogo v%s\n", Version)
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: shogo <url>")
		os.Exit(1)
	}

	var domain string
	if _, err := os.Stat("domain.txt"); err == nil {
		f, _ := os.Open("domain.txt")
		defer f.Close()
		_, _ = fmt.Fscanf(f, "%s", &domain)
	}
	if domain == "" {
		domain = os.Getenv("SHG_DOMAIN")
		if domain == "" {
			domain = Domain
			if domain == "" {
				fmt.Println("no domain found")
				os.Exit(1)
			}
		}
	}

	var key string
	if _, err := os.Stat("key.txt"); err == nil {
		f, _ := os.Open("key.txt")
		defer f.Close()
		_, _ = fmt.Fscanf(f, "%s", &key)
	}
	if key == "" {
		key = os.Getenv("SHORT_IO_KEY")
		if key == "" {
			key = Key
			if key == "" {
				fmt.Println("no api key found")
				os.Exit(1)
			}
		}
	}

	var target string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "https") {
			target = arg
		}
	}
	if target == "" || strings.Contains(target, "short.io") {
		fmt.Println("invalid target: must not contain 'short.io'")
		os.Exit(1)
	}
	if strings.Contains(target, "dropbox.com") {
		if !strings.Contains(target, "?dl=0") {
			target += "?dl=0"
		}
	}

	// not a free feature
	// expiry := fmt.Sprintf("%d", (time.Now().UnixNano()/1e6)+86400000)

	if verbose {
		fmt.Printf("setting up request to shorten '%s'\n", target)
	}

	client := &http.Client{}

	req, errCr := http.NewRequest("POST", "https://api.short.io/links", nil)
	if errCr != nil {
		fmt.Println("error creating link request")
		os.Exit(1)
	}
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`{"allowDuplicates":false,"domain":"%s","originalURL":"%s"}`, domain, target)))

	resp, errSe := client.Do(req)
	if errSe != nil {
		fmt.Println("error sending link request")
		os.Exit(1)
	}
	defer resp.Body.Close()

	var bBytes []byte
	if resp.Body != nil {
		bBytes, _ = io.ReadAll(resp.Body)
	} else {
		fmt.Println("error reading link response")
		os.Exit(1)
	}
	resultAsStr := string(bBytes)

	if verbose {
		fmt.Println("got response:")
		cmdLo := exec.Command("jq", ".")
		cmdLo.Stdin = strings.NewReader(resultAsStr)
		cmdLo.Stdout = os.Stdout
		cmdLo.Stderr = os.Stderr
		_ = cmdLo.Run()
	}
	if resp.StatusCode >= 299 {
		fmt.Println("http error creating link")
		os.Exit(1)
	}

	var result map[string]any
	errDe := json.NewDecoder(strings.NewReader(resultAsStr)).Decode(&result)
	if errDe != nil {
		fmt.Println("error decoding link")
		os.Exit(1)
	}
	shortUrl := result["shortURL"].(string)
	if shortUrl == "" {
		fmt.Println("error getting short url")
		os.Exit(1)
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(shortUrl)
	err := cmd.Run()
	if err != nil && verbose {
		fmt.Println("error copying url to clipboard")
	} else if err == nil && verbose {
		fmt.Println("short link created successfully and copied to clipboard:")
	}

	fmt.Println(shortUrl)
}
