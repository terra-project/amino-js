package main

import (
	gopher "github.com/gopherjs/gopherjs/js"
	lib "github.com/jordansexton/go-amino-js/go"
	"github.com/tendermint/go-amino"
)

func main() {
	gopher.Global.Set("Amino", map[string]interface{}{
		// Basic encoding
		"encodeByte":      lib.EncodeByte,
		"encodeByteSlice": lib.EncodeByteSlice,
		"encodeInt8":      lib.EncodeInt8,
		"encodeInt16":     lib.EncodeInt16,
		"encodeInt32":     lib.EncodeInt32,
		"encodeInt64":     lib.EncodeInt64,
		"encodeVarint":    lib.EncodeVarint,
		"encodeUint8":     lib.EncodeUint8,
		"encodeUint16":    lib.EncodeUint16,
		"encodeUint32":    lib.EncodeUint32,
		"encodeUint64":    lib.EncodeUint64,
		"encodeUvarint":   lib.EncodeUvarint,
		"encodeFloat32":   lib.EncodeFloat32,
		"encodeFloat64":   lib.EncodeFloat64,
		"encodeBool":      lib.EncodeBool,
		"encodeString":    lib.EncodeString,
		"encodeTime":      lib.EncodeTime,

		// Basic decoding
		"decodeByte":      lib.DecodeByte,
		"decodeByteSlice": lib.DecodeByteSlice,
		"decodeInt8":      lib.DecodeInt8,
		"decodeInt16":     lib.DecodeInt16,
		"decodeInt32":     lib.DecodeInt32,
		"decodeInt64":     lib.DecodeInt64,
		"decodeVarint":    lib.DecodeVarint,
		"decodeUint8":     lib.DecodeUint8,
		"decodeUint16":    lib.DecodeUint16,
		"decodeUint32":    lib.DecodeUint32,
		"decodeUint64":    lib.DecodeUint64,
		"decodeUvarint":   lib.DecodeUvarint,
		"decodeFloat32":   lib.DecodeFloat32,
		"decodeFloat64":   lib.DecodeFloat64,
		"decodeBool":      lib.DecodeBool,
		"decodeString":    lib.DecodeString,
		"decodeTime":      lib.DecodeTime,

		// Meta
		"decodeDisambPrefixBytes": lib.DecodeDisambPrefixBytes,
		"nameToDisfix":            amino.NameToDisfix,
		"byteSliceSize":           amino.ByteSliceSize,
		"uvarintSize":             amino.UvarintSize,
		"varintSize":              amino.VarintSize,

		// Typed encoding
		"encodeMultiStoreProofOp":              lib.EncodeMultiStoreProofOp,
		"encodeIAVLAbsenceOp":                  lib.EncodeIAVLAbsenceOp,
		"encodeIAVLValueOp":                    lib.EncodeIAVLValueOp,
		"encodePrivKeyLedgerSecp256k1":         lib.EncodePrivKeyLedgerSecp256k1,
		"encodeInfo":                           lib.EncodeInfo,
		"encodeBIP44Params":                    lib.EncodeBIP44Params,
		"encodeLocalInfo":                      lib.EncodeLocalInfo,
		"encodeLedgerInfo":                     lib.EncodeLedgerInfo,
		"encodeOfflineInfo":                    lib.EncodeOfflineInfo,
		"encodeMultiInfo":                      lib.EncodeMultiInfo,
		"encodeMsg":                            lib.EncodeMsg,
		"encodeTx":                             lib.EncodeTx,
		"encodeAccount":                        lib.EncodeAccount,
		"encodeVestingAccount":                 lib.EncodeVestingAccount,
		"encodeBaseAccount":                    lib.EncodeBaseAccount,
		"encodeBaseVestingAccount":             lib.EncodeBaseVestingAccount,
		"encodeContinuousVestingAccount":       lib.EncodeContinuousVestingAccount,
		"encodeDelayedVestingAccount":          lib.EncodeDelayedVestingAccount,
		"encodeStdTx":                          lib.EncodeStdTx,
		"encodeMsgSend":                        lib.EncodeMsgSend,
		"encodeMsgMultiSend":                   lib.EncodeMsgMultiSend,
		"encodeMsgVerifyInvariant":             lib.EncodeMsgVerifyInvariant,
		"encodeMsgWithdrawDelegatorReward":     lib.EncodeMsgWithdrawDelegatorReward,
		"encodeMsgWithdrawValidatorCommission": lib.EncodeMsgWithdrawValidatorCommission,
		"encodeMsgSetWithdrawAddress":          lib.EncodeMsgSetWithdrawAddress,
		"encodeContent":                        lib.EncodeContent,
		"encodeMsgSubmitProposal":              lib.EncodeMsgSubmitProposal,
		"encodeMsgDeposit":                     lib.EncodeMsgDeposit,
		"encodeMsgVote":                        lib.EncodeMsgVote,
		"encodeTextProposal":                   lib.EncodeTextProposal,
		"encodeSoftwareUpgradeProposal":        lib.EncodeSoftwareUpgradeProposal,
		"encodeMsgIBCTransfer":                 lib.EncodeMsgIBCTransfer,
		"encodeMsgIBCReceive":                  lib.EncodeMsgIBCReceive,
		"encodeParameterChangeProposal":        lib.EncodeParameterChangeProposal,
		"encodeMsgUnjail":                      lib.EncodeMsgUnjail,
		"encodeMsgCreateValidator":             lib.EncodeMsgCreateValidator,
		"encodeMsgEditValidator":               lib.EncodeMsgEditValidator,
		"encodeMsgDelegate":                    lib.EncodeMsgDelegate,
		"encodeMsgUndelegate":                  lib.EncodeMsgUndelegate,
		"encodeMsgBeginRedelegate":             lib.EncodeMsgBeginRedelegate,
		"encodeBlockchainMessage":              lib.EncodeBlockchainMessage,
		"encodeBcBlockRequestMessage":          lib.EncodeBcBlockRequestMessage,
		"encodeBcBlockResponseMessage":         lib.EncodeBcBlockResponseMessage,
		"encodeBcNoBlockResponseMessage":       lib.EncodeBcNoBlockResponseMessage,
		"encodeBcStatusResponseMessage":        lib.EncodeBcStatusResponseMessage,
		"encodeBcStatusRequestMessage":         lib.EncodeBcStatusRequestMessage,
		"encodeConsensusMessage":               lib.EncodeConsensusMessage,
		"encodeNewRoundStepMessage":            lib.EncodeNewRoundStepMessage,
		"encodeNewValidBlockMessage":           lib.EncodeNewValidBlockMessage,
		"encodeProposalMessage":                lib.EncodeProposalMessage,
		"encodeProposalPOLMessage":             lib.EncodeProposalPOLMessage,
		"encodeBlockPartMessage":               lib.EncodeBlockPartMessage,
		"encodeVoteMessage":                    lib.EncodeVoteMessage,
		"encodeHasVoteMessage":                 lib.EncodeHasVoteMessage,
		"encodeVoteSetMaj23Message":            lib.EncodeVoteSetMaj23Message,
		"encodeVoteSetBitsMessage":             lib.EncodeVoteSetBitsMessage,
		"encodeWALMessage":                     lib.EncodeWALMessage,
		"encodeMsgInfo":                        lib.EncodeMsgInfo,
		"encodeTimeoutInfo":                    lib.EncodeTimeoutInfo,
		"encodeEndHeightMessage":               lib.EncodeEndHeightMessage,
		"encodePubKey":                         lib.EncodePubKey,
		"encodePrivKey":                        lib.EncodePrivKey,
		"encodePubKeyEd25519":                  lib.EncodePubKeyEd25519,
		"encodePrivKeyEd25519":                 lib.EncodePrivKeyEd25519,
		"encodePubKeySecp256k1":                lib.EncodePubKeySecp256k1,
		"encodePrivKeySecp256k1":               lib.EncodePrivKeySecp256k1,
		"encodePubKeyMultisigThreshold":        lib.EncodePubKeyMultisigThreshold,
		"encodeEvidenceMessage":                lib.EncodeEvidenceMessage,
		"encodeEvidenceListMessage":            lib.EncodeEvidenceListMessage,
		"encodeMempoolMessage":                 lib.EncodeMempoolMessage,
		"encodeTxMessage":                      lib.EncodeTxMessage,
		"encodePacket":                         lib.EncodePacket,
		"encodePacketPing":                     lib.EncodePacketPing,
		"encodePacketPong":                     lib.EncodePacketPong,
		"encodePacketMsg":                      lib.EncodePacketMsg,
		"encodePexMessage":                     lib.EncodePexMessage,
		"encodePexRequestMessage":              lib.EncodePexRequestMessage,
		"encodePexAddrsMessage":                lib.EncodePexAddrsMessage,
		"encodeRemoteSignerMsg":                lib.EncodeRemoteSignerMsg,
		"encodePubKeyRequest":                  lib.EncodePubKeyRequest,
		"encodePubKeyResponse":                 lib.EncodePubKeyResponse,
		"encodeSignVoteRequest":                lib.EncodeSignVoteRequest,
		"encodeSignedVoteResponse":             lib.EncodeSignedVoteResponse,
		"encodeSignProposalRequest":            lib.EncodeSignProposalRequest,
		"encodeSignedProposalResponse":         lib.EncodeSignedProposalResponse,
		"encodePingRequest":                    lib.EncodePingRequest,
		"encodePingResponse":                   lib.EncodePingResponse,
		"encodeTMEventData":                    lib.EncodeTMEventData,
		"encodeEventDataNewBlock":              lib.EncodeEventDataNewBlock,
		"encodeEventDataNewBlockHeader":        lib.EncodeEventDataNewBlockHeader,
		"encodeEventDataTx":                    lib.EncodeEventDataTx,
		"encodeEventDataRoundState":            lib.EncodeEventDataRoundState,
		"encodeEventDataNewRound":              lib.EncodeEventDataNewRound,
		"encodeEventDataCompleteProposal":      lib.EncodeEventDataCompleteProposal,
		"encodeEventDataVote":                  lib.EncodeEventDataVote,
		"encodeEventDataValidatorSetUpdates":   lib.EncodeEventDataValidatorSetUpdates,
		"encodeEventDataString":                lib.EncodeEventDataString,
		"encodeEvidence":                       lib.EncodeEvidence,
		"encodeDuplicateVoteEvidence":          lib.EncodeDuplicateVoteEvidence,
		"encodeMockGoodEvidence":               lib.EncodeMockGoodEvidence,
		"encodeMockRandomGoodEvidence":         lib.EncodeMockRandomGoodEvidence,
		"encodeMockBadEvidence":                lib.EncodeMockBadEvidence,

		// Typed decoding
		"decodeMultiStoreProofOp":              lib.DecodeMultiStoreProofOp,
		"decodeIAVLAbsenceOp":                  lib.DecodeIAVLAbsenceOp,
		"decodeIAVLValueOp":                    lib.DecodeIAVLValueOp,
		"decodePrivKeyLedgerSecp256k1":         lib.DecodePrivKeyLedgerSecp256k1,
		"decodeInfo":                           lib.DecodeInfo,
		"decodeBIP44Params":                    lib.DecodeBIP44Params,
		"decodeLocalInfo":                      lib.DecodeLocalInfo,
		"decodeLedgerInfo":                     lib.DecodeLedgerInfo,
		"decodeOfflineInfo":                    lib.DecodeOfflineInfo,
		"decodeMultiInfo":                      lib.DecodeMultiInfo,
		"decodeMsg":                            lib.DecodeMsg,
		"decodeTx":                             lib.DecodeTx,
		"decodeAccount":                        lib.DecodeAccount,
		"decodeVestingAccount":                 lib.DecodeVestingAccount,
		"decodeBaseAccount":                    lib.DecodeBaseAccount,
		"decodeBaseVestingAccount":             lib.DecodeBaseVestingAccount,
		"decodeContinuousVestingAccount":       lib.DecodeContinuousVestingAccount,
		"decodeDelayedVestingAccount":          lib.DecodeDelayedVestingAccount,
		"decodeStdTx":                          lib.DecodeStdTx,
		"decodeMsgSend":                        lib.DecodeMsgSend,
		"decodeMsgMultiSend":                   lib.DecodeMsgMultiSend,
		"decodeMsgVerifyInvariant":             lib.DecodeMsgVerifyInvariant,
		"decodeMsgWithdrawDelegatorReward":     lib.DecodeMsgWithdrawDelegatorReward,
		"decodeMsgWithdrawValidatorCommission": lib.DecodeMsgWithdrawValidatorCommission,
		"decodeMsgSetWithdrawAddress":          lib.DecodeMsgSetWithdrawAddress,
		"decodeContent":                        lib.DecodeContent,
		"decodeMsgSubmitProposal":              lib.DecodeMsgSubmitProposal,
		"decodeMsgDeposit":                     lib.DecodeMsgDeposit,
		"decodeMsgVote":                        lib.DecodeMsgVote,
		"decodeTextProposal":                   lib.DecodeTextProposal,
		"decodeSoftwareUpgradeProposal":        lib.DecodeSoftwareUpgradeProposal,
		"decodeMsgIBCTransfer":                 lib.DecodeMsgIBCTransfer,
		"decodeMsgIBCReceive":                  lib.DecodeMsgIBCReceive,
		"decodeParameterChangeProposal":        lib.DecodeParameterChangeProposal,
		"decodeMsgUnjail":                      lib.DecodeMsgUnjail,
		"decodeMsgCreateValidator":             lib.DecodeMsgCreateValidator,
		"decodeMsgEditValidator":               lib.DecodeMsgEditValidator,
		"decodeMsgDelegate":                    lib.DecodeMsgDelegate,
		"decodeMsgUndelegate":                  lib.DecodeMsgUndelegate,
		"decodeMsgBeginRedelegate":             lib.DecodeMsgBeginRedelegate,
		"decodeBlockchainMessage":              lib.DecodeBlockchainMessage,
		"decodeBcBlockRequestMessage":          lib.DecodeBcBlockRequestMessage,
		"decodeBcBlockResponseMessage":         lib.DecodeBcBlockResponseMessage,
		"decodeBcNoBlockResponseMessage":       lib.DecodeBcNoBlockResponseMessage,
		"decodeBcStatusResponseMessage":        lib.DecodeBcStatusResponseMessage,
		"decodeBcStatusRequestMessage":         lib.DecodeBcStatusRequestMessage,
		"decodeConsensusMessage":               lib.DecodeConsensusMessage,
		"decodeNewRoundStepMessage":            lib.DecodeNewRoundStepMessage,
		"decodeNewValidBlockMessage":           lib.DecodeNewValidBlockMessage,
		"decodeProposalMessage":                lib.DecodeProposalMessage,
		"decodeProposalPOLMessage":             lib.DecodeProposalPOLMessage,
		"decodeBlockPartMessage":               lib.DecodeBlockPartMessage,
		"decodeVoteMessage":                    lib.DecodeVoteMessage,
		"decodeHasVoteMessage":                 lib.DecodeHasVoteMessage,
		"decodeVoteSetMaj23Message":            lib.DecodeVoteSetMaj23Message,
		"decodeVoteSetBitsMessage":             lib.DecodeVoteSetBitsMessage,
		"decodeWALMessage":                     lib.DecodeWALMessage,
		"decodeMsgInfo":                        lib.DecodeMsgInfo,
		"decodeTimeoutInfo":                    lib.DecodeTimeoutInfo,
		"decodeEndHeightMessage":               lib.DecodeEndHeightMessage,
		"decodePubKey":                         lib.DecodePubKey,
		"decodePrivKey":                        lib.DecodePrivKey,
		"decodePubKeyEd25519":                  lib.DecodePubKeyEd25519,
		"decodePrivKeyEd25519":                 lib.DecodePrivKeyEd25519,
		"decodePubKeySecp256k1":                lib.DecodePubKeySecp256k1,
		"decodePrivKeySecp256k1":               lib.DecodePrivKeySecp256k1,
		"decodePubKeyMultisigThreshold":        lib.DecodePubKeyMultisigThreshold,
		"decodeEvidenceMessage":                lib.DecodeEvidenceMessage,
		"decodeEvidenceListMessage":            lib.DecodeEvidenceListMessage,
		"decodeMempoolMessage":                 lib.DecodeMempoolMessage,
		"decodeTxMessage":                      lib.DecodeTxMessage,
		"decodePacket":                         lib.DecodePacket,
		"decodePacketPing":                     lib.DecodePacketPing,
		"decodePacketPong":                     lib.DecodePacketPong,
		"decodePacketMsg":                      lib.DecodePacketMsg,
		"decodePexMessage":                     lib.DecodePexMessage,
		"decodePexRequestMessage":              lib.DecodePexRequestMessage,
		"decodePexAddrsMessage":                lib.DecodePexAddrsMessage,
		"decodeRemoteSignerMsg":                lib.DecodeRemoteSignerMsg,
		"decodePubKeyRequest":                  lib.DecodePubKeyRequest,
		"decodePubKeyResponse":                 lib.DecodePubKeyResponse,
		"decodeSignVoteRequest":                lib.DecodeSignVoteRequest,
		"decodeSignedVoteResponse":             lib.DecodeSignedVoteResponse,
		"decodeSignProposalRequest":            lib.DecodeSignProposalRequest,
		"decodeSignedProposalResponse":         lib.DecodeSignedProposalResponse,
		"decodePingRequest":                    lib.DecodePingRequest,
		"decodePingResponse":                   lib.DecodePingResponse,
		"decodeTMEventData":                    lib.DecodeTMEventData,
		"decodeEventDataNewBlock":              lib.DecodeEventDataNewBlock,
		"decodeEventDataNewBlockHeader":        lib.DecodeEventDataNewBlockHeader,
		"decodeEventDataTx":                    lib.DecodeEventDataTx,
		"decodeEventDataRoundState":            lib.DecodeEventDataRoundState,
		"decodeEventDataNewRound":              lib.DecodeEventDataNewRound,
		"decodeEventDataCompleteProposal":      lib.DecodeEventDataCompleteProposal,
		"decodeEventDataVote":                  lib.DecodeEventDataVote,
		"decodeEventDataValidatorSetUpdates":   lib.DecodeEventDataValidatorSetUpdates,
		"decodeEventDataString":                lib.DecodeEventDataString,
		"decodeEvidence":                       lib.DecodeEvidence,
		"decodeDuplicateVoteEvidence":          lib.DecodeDuplicateVoteEvidence,
		"decodeMockGoodEvidence":               lib.DecodeMockGoodEvidence,
		"decodeMockRandomGoodEvidence":         lib.DecodeMockRandomGoodEvidence,
		"decodeMockBadEvidence":                lib.DecodeMockBadEvidence,
	})
}
