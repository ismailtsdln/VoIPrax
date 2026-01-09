package main

import (
	"time"

	"github.com/ismailtsdln/VoIPrax/internal/fuzz"
	"github.com/ismailtsdln/VoIPrax/internal/sip"
	"github.com/spf13/cobra"
)

var (
	targetURI string
	fromURI   string
	toURI     string
	count     int
)

var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "Start SIP fuzzing",
	RunE: func(cmd *cobra.Command, args []string) error {
		stack := sip.NewStack(log)
		if err := stack.ListenUDP(":0"); err != nil {
			return err
		}
		defer stack.Close()

		fuzzer := fuzz.NewFuzzer()

		log.Info().Str("target", targetURI).Msg("Starting fuzzing session")

		for i := 0; i < count; i++ {
			msg := fuzz.GenerateInviteTemplate(targetURI, fromURI, toURI)

			// Fuzz headers
			fuzzer.FuzzHeader(msg, "Via")
			fuzzer.FuzzHeader(msg, "Contact")

			log.Debug().Int("iteration", i+1).Msg("Sending fuzzed message")

			if err := stack.SendUDP(targetURI, msg); err != nil {
				log.Error().Err(err).Msg("Failed to send fuzzed message")
			}

			// Small delay between packets
			time.Sleep(100 * time.Millisecond)
		}

		log.Info().Msg("Fuzzing session complete")
		return nil
	},
}

func init() {
	fuzzCmd.Flags().StringVarP(&targetURI, "target", "t", "", "Target SIP URI (e.g., 192.168.1.1:5060)")
	fuzzCmd.Flags().StringVar(&fromURI, "from", "sip:alice@atlanta.com", "From SIP URI")
	fuzzCmd.Flags().StringVar(&toURI, "to", "sip:bob@biloxi.com", "To SIP URI")
	fuzzCmd.Flags().IntVarP(&count, "count", "c", 100, "Number of fuzzed packets to send")
	fuzzCmd.MarkFlagRequired("target")

	rootCmd.AddCommand(fuzzCmd)
}
