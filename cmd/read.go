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
	"bytes"
	"io"
	"os"

	"github.com/chris-rock/pgbytea/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// readCmd represents the load command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a specific file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		s := &store.Store{ConnStr: os.Getenv("PG_CONN")}
		s.Connect()
		defer s.Close()

		data, err := s.Read(args[0])
		if err != nil {
			log.Error().Err(err).Msg("could not prepare insert statement")
			os.Exit(1)
		}
		reader := bytes.NewReader(data)
		io.Copy(os.Stdout, reader)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
