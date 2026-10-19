package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type PoSTProof struct {
	NodeID     string
	SpaceUnits uint64
	Challenge  string
}

func (p *PoSTProof) Verify() string {
	h := sha256.New()
	h.Write([]byte(p.NodeID + p.Challenge))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	proof := PoSTProof{
		NodeID:     "smesher_node_sample_01",
		SpaceUnits: 4,
		Challenge:  "epoch_challenge_proof",
	}
	fmt.Printf("Spacemesh PoST Verification Hash: %s\n", proof.Verify())
}
