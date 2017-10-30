package wal

const (
	pageSize       = 4096
	pageMetaSize   = 64 // 32-byte checksum + 4 uint64s
	maxPayloadSize = pageSize - pageMetaSize

	checksumSize = 32
)

const (
	pageStatusInvalid = iota
	pageStatusOther
	pageStatusWritten
	pageStatusComitted
	pageStatusApplied
)

const (
	metadataHeader  = "WAL"
	metadataVersion = "1.0"
)

// metadata contains the header and version of the WAL.
type metadata struct {
	Header, Version string
}

// A checksum is a 256-bit blake2b hash.
type checksum [checksumSize]byte
