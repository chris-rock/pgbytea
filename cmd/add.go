// Copyright © 2018 Christoph Hartmann chris@lollyrock.com
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
	"fmt"
	"os"

	"github.com/chris-rock/pgbytea/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// addCmd represents the store command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Stores a new file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		s := &store.Store{ConnStr: os.Getenv("PG_CONN")}
		s.Connect()
		defer s.Close()
		err := s.Save(args[0], args[1])
		if err != nil {
			log.Error().Err(err).Msg("could not insert file")
			os.Exit(1)
		}
		fmt.Println("Stored file successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
