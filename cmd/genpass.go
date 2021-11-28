/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// genpassCmd represents the genpass command
var genpassCmd = &cobra.Command{
	Use:   "genpass",
	Short: "Generate random password",
	Long: `Generate random password. If no argument provides the password will be 12 characters long`,
	Run: func(cmd *cobra.Command, args []string) {
		customFlag, err := cmd.Flags().GetBool("with_special_characters")
		if err != nil {
			fmt.Println(err)
		}
		generatePassword(args, customFlag)
	},
}

func init() {
	genpassCmd.Flags().BoolP("with_special_characters", "w", false, "Generate ramdom password with special characters")
	rootCmd.AddCommand(genpassCmd)
}

// generatePassword generates a random password 
func generatePassword(args []string, wsc bool) {
	var length int
	var chars []rune
	if len(args) == 0 {
		fmt.Println("Password lenght is set by default to 12 characters")
		length = 12
	} else {
		argsLength, err :=strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Argument needs to be an int")
			os.Exit(1)
		}
		length = argsLength
	}
	rand.Seed(time.Now().UnixNano())
	if wsc {
		chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "@#$*")
	} else {
		chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	}
	
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	generatedPassword := str
	fmt.Println(generatedPassword)
}
