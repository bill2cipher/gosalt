// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/jellybean4/gosalt/release"
	"github.com/spf13/cobra"
)

var (
	release_version string
	release_types   []string
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "pack release server packages",
	Long: `pack the specific release server packages, version defined which version to release, which is passed
as the third argument into release script; types defined which server type to release, which is given as the
last arguments passed to release script`,
	Run: release_runner,
}

func init() {
	RootCmd.AddCommand(releaseCmd)
	releaseCmd.Flags().StringVar(&release_version, "version", "trunk", "code version to release (default is trunk)")
	releaseCmd.Flags().StringSliceVar(&release_types, "types", []string{"all"}, "types to release (default is all)")
}

func release_runner(cmd *cobra.Command, args []string) {
	log.WithFields(log.Fields{
		"version": release_version,
		"types":   release_types,
	}).Info("release server package starting...")
	Release(release_version, release_types...)
}
