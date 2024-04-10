package request

import (
	"time"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/Bruce960205/b7s/models/execute"
)

// Execute describes the `MessageExecute` request payload.
type Execute struct {
	Type  string  `json:"type,omitempty"`
	From  peer.ID `json:"from,omitempty"`
	Code  string  `json:"code,omitempty"`
	Topic string  `json:"topic,omitempty"`

	execute.Request // execute request is embedded.

	// RequestID may be set initially, if the execution request is relayed via roll-call.
	RequestID      string    `json:"request_id,omitempty"`
	ReportingPeers []peer.ID `json:"reporting_peers,omitempty"`

	// Execution request timestamp is a factor for PBFT.
	Timestamp time.Time `json:"timestamp,omitempty"`
}
