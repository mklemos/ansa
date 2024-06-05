package models

import (
    "fmt"
    "time"
)

type NFT struct {
    ID      string
    Details string
}

func generateNFTID() string {
    // Generate a unique NFT ID
    return fmt.Sprintf("nft-%d", time.Now().UnixNano())
}
