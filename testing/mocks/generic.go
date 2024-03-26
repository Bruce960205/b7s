package mocks

import (
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/rs/zerolog"

	"github.com/RedBird96/b7s/models/blockless"
	"github.com/RedBird96/b7s/models/codes"
	"github.com/RedBird96/b7s/models/execute"
)

// Global variables that can be used for testing. They are valid non-nil values for commonly needed types.
var (
	NoopLogger = zerolog.New(io.Discard)

	GenericError = errors.New("dummy error")

	GenericPeerID = peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x56, 0x77, 0x86, 0x82, 0x76, 0xa, 0xc5, 0x9, 0x63, 0xde, 0xe4, 0x31, 0xfc, 0x44, 0x75, 0xdd, 0x5a, 0x27, 0xee, 0x6b, 0x94, 0x13, 0xed, 0xe2, 0xa3, 0x6d, 0x8a, 0x1d, 0x57, 0xb6, 0xb8, 0x91})

	GenericAddress = "/ip4/127.0.0.1/tcp/9000/p2p/12D3KooWRp3AVk7qtc2Av6xiqgAza1ZouksQaYcS2cvN94kHSCoa"

	GenericString = "test"

	GenericUUID = uuid.UUID{0xd1, 0xc2, 0x44, 0xaf, 0xa3, 0x1d, 0x48, 0x87, 0x93, 0x9d, 0xd6, 0xc7, 0xf, 0xe, 0x4f, 0xd0}

	GenericExecutionResult = execute.Result{
		Code: codes.Unknown,
		Result: execute.RuntimeOutput{
			Stdout:   "generic-execution-result",
			Stderr:   "generic-execution-log",
			ExitCode: 0,
		},
		RequestID: GenericUUID.String(),
	}

	GenericExecutionRequest = execute.Request{
		FunctionID: "generic-function-id",
		Method:     "wasm",
		Parameters: []execute.Parameter{
			{
				Name:  "generic-param-name",
				Value: "generic-param-value",
			},
		},
	}

	GenericManifest = blockless.FunctionManifest{
		ID:          "generic-id",
		Name:        "generic-name",
		Description: "generic-description",
		Function: blockless.Function{
			ID:      "function-id",
			Name:    "function-name",
			Runtime: "generic-runtime",
		},
		Deployment: blockless.Deployment{
			CID:      "generic-cid",
			Checksum: "1234567890",
			URI:      "generic-uri",
		},
		FSRootPath: "/var/tmp/blockless/",
		Entry:      "/var/tmp/blockless/app.wasm",
	}

	// List of a few peer IDs in case multiple are required.
	GenericPeerIDs = []peer.ID{
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0xe5, 0xc, 0xbd, 0xb8, 0xff, 0xed, 0x5a, 0x74, 0x48, 0x4, 0x2, 0x33, 0x4e, 0x42, 0xc, 0x40, 0xab, 0x28, 0x3a, 0x28, 0xdb, 0x5, 0x7e, 0xc6, 0xf5, 0x0, 0x6f, 0x36, 0xa2, 0x8d, 0x82, 0x48}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x19, 0x2a, 0x38, 0x8d, 0xf9, 0x66, 0xa1, 0x14, 0xea, 0x6c, 0xce, 0x1f, 0xf2, 0x3b, 0xee, 0x3b, 0x56, 0xe9, 0x56, 0x27, 0x7e, 0x70, 0x1b, 0x49, 0x9c, 0x25, 0x5, 0x8e, 0xb, 0xda, 0xa3, 0x87}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x68, 0xbc, 0x89, 0x26, 0x87, 0xd5, 0x10, 0x62, 0xa8, 0x6, 0x83, 0xb4, 0xae, 0x62, 0xe1, 0x87, 0xe7, 0xcd, 0x4e, 0x2b, 0x58, 0xa5, 0x82, 0x3b, 0x6a, 0xf6, 0x57, 0x83, 0x38, 0x80, 0x5b, 0xc2}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x26, 0xa8, 0xd0, 0xc6, 0x7c, 0x0, 0xcb, 0xf, 0x7e, 0x23, 0xd9, 0x1c, 0x88, 0xef, 0xc6, 0x2a, 0xd5, 0xa3, 0x18, 0xb5, 0xde, 0xa7, 0x21, 0x44, 0xb0, 0x38, 0xa7, 0xc9, 0x18, 0x6f, 0xd1, 0x25}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x5e, 0x5d, 0x44, 0xfe, 0xde, 0xbc, 0xd7, 0xbd, 0x82, 0xcd, 0x49, 0x23, 0xe8, 0x48, 0x46, 0x56, 0xca, 0x61, 0x13, 0x9d, 0xf4, 0x14, 0x5, 0x8a, 0x87, 0xdd, 0xd2, 0xd9, 0xd6, 0x1c, 0xc4, 0xc2}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x37, 0xf2, 0xf0, 0x1a, 0x4f, 0x1a, 0xaf, 0xd3, 0x48, 0xf4, 0xe8, 0xa7, 0x4b, 0xb, 0xfe, 0x5, 0xbb, 0x18, 0x1, 0xcb, 0x44, 0x1d, 0xe4, 0x4, 0x31, 0x5c, 0x55, 0xf, 0xbd, 0xae, 0x77, 0x95}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x9, 0xb1, 0x4c, 0x68, 0xd5, 0x17, 0x29, 0x76, 0xbc, 0xca, 0xe8, 0xa8, 0x76, 0x7c, 0x2b, 0x82, 0x68, 0xa1, 0xae, 0xe8, 0x35, 0x4b, 0x42, 0xff, 0x5f, 0xaa, 0xe9, 0x6, 0x2d, 0x46, 0xa1, 0x9c}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0xc3, 0x36, 0xd9, 0xea, 0x3, 0x29, 0x85, 0x17, 0x8e, 0x60, 0xa, 0xc5, 0xf, 0x5d, 0xe3, 0xe8, 0xd, 0x1c, 0x53, 0x7b, 0x31, 0x82, 0x58, 0xc9, 0x4e, 0x80, 0x97, 0x8d, 0x6d, 0x1c, 0x97, 0x86}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x49, 0xbd, 0x82, 0x34, 0x7d, 0x97, 0xd, 0xb5, 0x52, 0xf5, 0x82, 0x47, 0x8b, 0xc, 0x42, 0x16, 0x22, 0xa5, 0x24, 0x32, 0xf2, 0x24, 0xfb, 0xc7, 0x44, 0xd1, 0xfc, 0xe1, 0x2e, 0xd3, 0x70, 0xa7}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x24, 0x49, 0x1b, 0x67, 0x6, 0x91, 0x6, 0x94, 0xa, 0xcb, 0x5d, 0x1c, 0xe5, 0x69, 0x37, 0xbe, 0x7c, 0x6b, 0x7e, 0x97, 0x4b, 0x44, 0xd7, 0xbe, 0x94, 0x22, 0x9f, 0xfa, 0x1e, 0x7e, 0x2d, 0xcf}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x15, 0xf8, 0xc7, 0x4b, 0x28, 0x5e, 0x1e, 0xf9, 0x96, 0xd8, 0xbd, 0x15, 0xf9, 0xde, 0x46, 0x39, 0x30, 0xe1, 0xa1, 0x2e, 0xa4, 0x17, 0x5f, 0xef, 0xbd, 0x2d, 0xe4, 0xd1, 0x43, 0xe8, 0xcb, 0x53}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0xd0, 0xca, 0xd3, 0x5c, 0x95, 0xf5, 0xdd, 0xb7, 0x73, 0x4e, 0xe3, 0x6a, 0x6b, 0xfc, 0x73, 0xe5, 0x55, 0x8d, 0xbf, 0x78, 0x2f, 0xa8, 0x42, 0xc9, 0x1d, 0x70, 0xd6, 0xce, 0x2b, 0x9e, 0x4f, 0xf7}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0x9a, 0xc4, 0x96, 0xd, 0x16, 0x58, 0x7d, 0x93, 0x40, 0xa, 0x7a, 0xf, 0xdf, 0x48, 0x12, 0x19, 0x46, 0xc5, 0x4d, 0x2d, 0x8e, 0x11, 0x96, 0xb2, 0xf6, 0xb7, 0x4e, 0x51, 0xff, 0xee, 0x0, 0xb5}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0xaa, 0xed, 0x63, 0x5d, 0xa1, 0xdf, 0x41, 0x2b, 0xe8, 0x9c, 0x49, 0xed, 0xe8, 0x0, 0x5c, 0xa8, 0x64, 0x58, 0x1d, 0x3, 0xf3, 0x59, 0x41, 0x74, 0xff, 0x2b, 0xcd, 0xde, 0x37, 0xfe, 0x15, 0xc6}),
		peer.ID([]byte{0x0, 0x24, 0x8, 0x1, 0x12, 0x20, 0xc6, 0x8f, 0x95, 0xd3, 0x98, 0x66, 0x40, 0x6b, 0xc4, 0x6c, 0x19, 0x5e, 0x80, 0xe0, 0x8c, 0x9b, 0x15, 0x4f, 0x8c, 0x6b, 0xd0, 0x1d, 0x5b, 0x83, 0x23, 0x7b, 0x9a, 0x97, 0xc0, 0x9b, 0x9d, 0x9b}),
	}
)
