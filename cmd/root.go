package cmd

import (
	"fmt"
	"os"

	"github.com/oasysgames/oasys-pos-cli/cmd/crypto"
	"github.com/oasysgames/oasys-pos-cli/cmd/validator"
	"github.com/spf13/cobra"
)

const (
	versionMajor = 1
	versionMinor = 0
	versionPatch = 0
	versionMeta  = "alpha0"
)

var rootCmd = &cobra.Command{
	Use: "oaspos",
	Long: fmt.Sprintf(`Name:
  oaspos - Command Line Tool for manage proof-of-stake of Oasys Blockchain.

  Copyright 2022 Oasys | Blockchain for The Games All Rights Reserved.
  
Version:
  %d.%d.%d-%s`, versionMajor, versionMinor, versionPatch, versionMeta),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	validator.AddCommand(rootCmd)
	crypto.AddCommand(rootCmd)
}
