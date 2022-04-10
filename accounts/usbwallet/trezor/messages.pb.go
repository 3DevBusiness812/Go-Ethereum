
package trezor

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// Mapping between TREZOR wire identifier (uint) and a protobuf message
type MessageType int32

const (
	// Management
	MessageType_MessageType_Initialize             MessageType = 0
	MessageType_MessageType_Ping                   MessageType = 1
	MessageType_MessageType_Success                MessageType = 2
	MessageType_MessageType_Failure                MessageType = 3
	MessageType_MessageType_ChangePin              MessageType = 4
	MessageType_MessageType_WipeDevice             MessageType = 5
	MessageType_MessageType_GetEntropy             MessageType = 9
	MessageType_MessageType_Entropy                MessageType = 10
	MessageType_MessageType_LoadDevice             MessageType = 13
	MessageType_MessageType_ResetDevice            MessageType = 14
	MessageType_MessageType_Features               MessageType = 17
	MessageType_MessageType_PinMatrixRequest       MessageType = 18
	MessageType_MessageType_PinMatrixAck           MessageType = 19
	MessageType_MessageType_Cancel                 MessageType = 20
	MessageType_MessageType_ClearSession           MessageType = 24
	MessageType_MessageType_ApplySettings          MessageType = 25
	MessageType_MessageType_ButtonRequest          MessageType = 26
	MessageType_MessageType_ButtonAck              MessageType = 27
	MessageType_MessageType_ApplyFlags             MessageType = 28
	MessageType_MessageType_BackupDevice           MessageType = 34
	MessageType_MessageType_EntropyRequest         MessageType = 35
	MessageType_MessageType_EntropyAck             MessageType = 36
	MessageType_MessageType_PassphraseRequest      MessageType = 41
	MessageType_MessageType_PassphraseAck          MessageType = 42
	MessageType_MessageType_PassphraseStateRequest MessageType = 77
	MessageType_MessageType_PassphraseStateAck     MessageType = 78
	MessageType_MessageType_RecoveryDevice         MessageType = 45
	MessageType_MessageType_WordRequest            MessageType = 46
	MessageType_MessageType_WordAck                MessageType = 47
	MessageType_MessageType_GetFeatures            MessageType = 55
	MessageType_MessageType_SetU2FCounter          MessageType = 63
	// Bootloader
	MessageType_MessageType_FirmwareErase   MessageType = 6
	MessageType_MessageType_FirmwareUpload  MessageType = 7
	MessageType_MessageType_FirmwareRequest MessageType = 8
	MessageType_MessageType_SelfTest        MessageType = 32
	// Bitcoin
	MessageType_MessageType_GetPublicKey     MessageType = 11
	MessageType_MessageType_PublicKey        MessageType = 12
	MessageType_MessageType_SignTx           MessageType = 15
	MessageType_MessageType_TxRequest        MessageType = 21
	MessageType_MessageType_TxAck            MessageType = 22
	MessageType_MessageType_GetAddress       MessageType = 29
	MessageType_MessageType_Address          MessageType = 30
	MessageType_MessageType_SignMessage      MessageType = 38
	MessageType_MessageType_VerifyMessage    MessageType = 39
	MessageType_MessageType_MessageSignature MessageType = 40
	// Crypto
	MessageType_MessageType_CipherKeyValue    MessageType = 23
	MessageType_MessageType_CipheredKeyValue  MessageType = 48
	MessageType_MessageType_SignIdentity      MessageType = 53
	MessageType_MessageType_SignedIdentity    MessageType = 54
	MessageType_MessageType_GetECDHSessionKey MessageType = 61
	MessageType_MessageType_ECDHSessionKey    MessageType = 62
	MessageType_MessageType_CosiCommit        MessageType = 71
	MessageType_MessageType_CosiCommitment    MessageType = 72
	MessageType_MessageType_CosiSign          MessageType = 73
	MessageType_MessageType_CosiSignature     MessageType = 74
	// Debug
	MessageType_MessageType_DebugLinkDecision    MessageType = 100
	MessageType_MessageType_DebugLinkGetState    MessageType = 101
	MessageType_MessageType_DebugLinkState       MessageType = 102
	MessageType_MessageType_DebugLinkStop        MessageType = 103
	MessageType_MessageType_DebugLinkLog         MessageType = 104
	MessageType_MessageType_DebugLinkMemoryRead  MessageType = 110
	MessageType_MessageType_DebugLinkMemory      MessageType = 111
	MessageType_MessageType_DebugLinkMemoryWrite MessageType = 112
	MessageType_MessageType_DebugLinkFlashErase  MessageType = 113
	// Ethereum
	MessageType_MessageType_EthereumGetPublicKey     MessageType = 450
	MessageType_MessageType_EthereumPublicKey        MessageType = 451
	MessageType_MessageType_EthereumGetAddress       MessageType = 56
	MessageType_MessageType_EthereumAddress          MessageType = 57
	MessageType_MessageType_EthereumSignTx           MessageType = 58
	MessageType_MessageType_EthereumTxRequest        MessageType = 59
	MessageType_MessageType_EthereumTxAck            MessageType = 60
	MessageType_MessageType_EthereumSignMessage      MessageType = 64
	MessageType_MessageType_EthereumVerifyMessage    MessageType = 65
	MessageType_MessageType_EthereumMessageSignature MessageType = 66
	// NEM
	MessageType_MessageType_NEMGetAddress       MessageType = 67
	MessageType_MessageType_NEMAddress          MessageType = 68
	MessageType_MessageType_NEMSignTx           MessageType = 69
	MessageType_MessageType_NEMSignedTx         MessageType = 70
	MessageType_MessageType_NEMDecryptMessage   MessageType = 75
	MessageType_MessageType_NEMDecryptedMessage MessageType = 76
	// Lisk
	MessageType_MessageType_LiskGetAddress       MessageType = 114
	MessageType_MessageType_LiskAddress          MessageType = 115
	MessageType_MessageType_LiskSignTx           MessageType = 116
	MessageType_MessageType_LiskSignedTx         MessageType = 117
	MessageType_MessageType_LiskSignMessage      MessageType = 118
	MessageType_MessageType_LiskMessageSignature MessageType = 119
	MessageType_MessageType_LiskVerifyMessage    MessageType = 120
	MessageType_MessageType_LiskGetPublicKey     MessageType = 121
	MessageType_MessageType_LiskPublicKey        MessageType = 122
	// Tezos
	MessageType_MessageType_TezosGetAddress   MessageType = 150
	MessageType_MessageType_TezosAddress      MessageType = 151
	MessageType_MessageType_TezosSignTx       MessageType = 152
	MessageType_MessageType_TezosSignedTx     MessageType = 153
	MessageType_MessageType_TezosGetPublicKey MessageType = 154
	MessageType_MessageType_TezosPublicKey    MessageType = 155
	// Stellar
	MessageType_MessageType_StellarSignTx               MessageType = 202
	MessageType_MessageType_StellarTxOpRequest          MessageType = 203
	MessageType_MessageType_StellarGetAddress           MessageType = 207
	MessageType_MessageType_StellarAddress              MessageType = 208
	MessageType_MessageType_StellarCreateAccountOp      MessageType = 210
	MessageType_MessageType_StellarPaymentOp            MessageType = 211
	MessageType_MessageType_StellarPathPaymentOp        MessageType = 212
	MessageType_MessageType_StellarManageOfferOp        MessageType = 213
	MessageType_MessageType_StellarCreatePassiveOfferOp MessageType = 214
	MessageType_MessageType_StellarSetOptionsOp         MessageType = 215
	MessageType_MessageType_StellarChangeTrustOp        MessageType = 216
	MessageType_MessageType_StellarAllowTrustOp         MessageType = 217
	MessageType_MessageType_StellarAccountMergeOp       MessageType = 218
	// omitted: StellarInflationOp is not a supported operation, would be 219
	MessageType_MessageType_StellarManageDataOp   MessageType = 220
	MessageType_MessageType_StellarBumpSequenceOp MessageType = 221
	MessageType_MessageType_StellarSignedTx       MessageType = 230
	// TRON
	MessageType_MessageType_TronGetAddress MessageType = 250
	MessageType_MessageType_TronAddress    MessageType = 251
	MessageType_MessageType_TronSignTx     MessageType = 252
	MessageType_MessageType_TronSignedTx   MessageType = 253
	// Cardano
	// dropped Sign/VerifyMessage ids 300-302
	MessageType_MessageType_CardanoSignTx       MessageType = 303
	MessageType_MessageType_CardanoTxRequest    MessageType = 304
	MessageType_MessageType_CardanoGetPublicKey MessageType = 305
	MessageType_MessageType_CardanoPublicKey    MessageType = 306
	MessageType_MessageType_CardanoGetAddress   MessageType = 307
	MessageType_MessageType_CardanoAddress      MessageType = 308
	MessageType_MessageType_CardanoTxAck        MessageType = 309
	MessageType_MessageType_CardanoSignedTx     MessageType = 310
	// Ontology
	MessageType_MessageType_OntologyGetAddress               MessageType = 350
	MessageType_MessageType_OntologyAddress                  MessageType = 351
	MessageType_MessageType_OntologyGetPublicKey             MessageType = 352
	MessageType_MessageType_OntologyPublicKey                MessageType = 353
	MessageType_MessageType_OntologySignTransfer             MessageType = 354
	MessageType_MessageType_OntologySignedTransfer           MessageType = 355
	MessageType_MessageType_OntologySignWithdrawOng          MessageType = 356
	MessageType_MessageType_OntologySignedWithdrawOng        MessageType = 357
	MessageType_MessageType_OntologySignOntIdRegister        MessageType = 358
	MessageType_MessageType_OntologySignedOntIdRegister      MessageType = 359
	MessageType_MessageType_OntologySignOntIdAddAttributes   MessageType = 360
	MessageType_MessageType_OntologySignedOntIdAddAttributes MessageType = 361
	// Ripple
	MessageType_MessageType_RippleGetAddress MessageType = 400
	MessageType_MessageType_RippleAddress    MessageType = 401
	MessageType_MessageType_RippleSignTx     MessageType = 402
	MessageType_MessageType_RippleSignedTx   MessageType = 403
	// Monero
	MessageType_MessageType_MoneroTransactionInitRequest              MessageType = 501
	MessageType_MessageType_MoneroTransactionInitAck                  MessageType = 502
	MessageType_MessageType_MoneroTransactionSetInputRequest          MessageType = 503
	MessageType_MessageType_MoneroTransactionSetInputAck              MessageType = 504
	MessageType_MessageType_MoneroTransactionInputsPermutationRequest MessageType = 505
	MessageType_MessageType_MoneroTransactionInputsPermutationAck     MessageType = 506
	MessageType_MessageType_MoneroTransactionInputViniRequest         MessageType = 507
	MessageType_MessageType_MoneroTransactionInputViniAck             MessageType = 508
	MessageType_MessageType_MoneroTransactionAllInputsSetRequest      MessageType = 509
	MessageType_MessageType_MoneroTransactionAllInputsSetAck          MessageType = 510
	MessageType_MessageType_MoneroTransactionSetOutputRequest         MessageType = 511
	MessageType_MessageType_MoneroTransactionSetOutputAck             MessageType = 512
	MessageType_MessageType_MoneroTransactionAllOutSetRequest         MessageType = 513
	MessageType_MessageType_MoneroTransactionAllOutSetAck             MessageType = 514
	MessageType_MessageType_MoneroTransactionSignInputRequest         MessageType = 515
	MessageType_MessageType_MoneroTransactionSignInputAck             MessageType = 516
	MessageType_MessageType_MoneroTransactionFinalRequest             MessageType = 517
	MessageType_MessageType_MoneroTransactionFinalAck                 MessageType = 518
	MessageType_MessageType_MoneroKeyImageExportInitRequest           MessageType = 530
	MessageType_MessageType_MoneroKeyImageExportInitAck               MessageType = 531
	MessageType_MessageType_MoneroKeyImageSyncStepRequest             MessageType = 532
	MessageType_MessageType_MoneroKeyImageSyncStepAck                 MessageType = 533
	MessageType_MessageType_MoneroKeyImageSyncFinalRequest            MessageType = 534
	MessageType_MessageType_MoneroKeyImageSyncFinalAck                MessageType = 535
	MessageType_MessageType_MoneroGetAddress                          MessageType = 540
	MessageType_MessageType_MoneroAddress                             MessageType = 541
	MessageType_MessageType_MoneroGetWatchKey                         MessageType = 542
	MessageType_MessageType_MoneroWatchKey                            MessageType = 543
	MessageType_MessageType_DebugMoneroDiagRequest                    MessageType = 546
	MessageType_MessageType_DebugMoneroDiagAck                        MessageType = 547
	MessageType_MessageType_MoneroGetTxKeyRequest                     MessageType = 550
	MessageType_MessageType_MoneroGetTxKeyAck                         MessageType = 551
	MessageType_MessageType_MoneroLiveRefreshStartRequest             MessageType = 552
	MessageType_MessageType_MoneroLiveRefreshStartAck                 MessageType = 553
	MessageType_MessageType_MoneroLiveRefreshStepRequest              MessageType = 554
	MessageType_MessageType_MoneroLiveRefreshStepAck                  MessageType = 555
	MessageType_MessageType_MoneroLiveRefreshFinalRequest             MessageType = 556
	MessageType_MessageType_MoneroLiveRefreshFinalAck                 MessageType = 557
	// EOS
	MessageType_MessageType_EosGetPublicKey    MessageType = 600
	MessageType_MessageType_EosPublicKey       MessageType = 601
	MessageType_MessageType_EosSignTx          MessageType = 602
	MessageType_MessageType_EosTxActionRequest MessageType = 603
	MessageType_MessageType_EosTxActionAck     MessageType = 604
	MessageType_MessageType_EosSignedTx        MessageType = 605
	// Binance
	MessageType_MessageType_BinanceGetAddress   MessageType = 700
	MessageType_MessageType_BinanceAddress      MessageType = 701
	MessageType_MessageType_BinanceGetPublicKey MessageType = 702
	MessageType_MessageType_BinancePublicKey    MessageType = 703
	MessageType_MessageType_BinanceSignTx       MessageType = 704
	MessageType_MessageType_BinanceTxRequest    MessageType = 705
	MessageType_MessageType_BinanceTransferMsg  MessageType = 706
	MessageType_MessageType_BinanceOrderMsg     MessageType = 707
	MessageType_MessageType_BinanceCancelMsg    MessageType = 708
	MessageType_MessageType_BinanceSignedTx     MessageType = 709
)

var MessageType_name = map[int32]string{
	0:   "MessageType_Initialize",
	1:   "MessageType_Ping",
	2:   "MessageType_Success",
	3:   "MessageType_Failure",
	4:   "MessageType_ChangePin",
	5:   "MessageType_WipeDevice",
	9:   "MessageType_GetEntropy",
	10:  "MessageType_Entropy",
	13:  "MessageType_LoadDevice",
	14:  "MessageType_ResetDevice",
	17:  "MessageType_Features",
	18:  "MessageType_PinMatrixRequest",
	19:  "MessageType_PinMatrixAck",
	20:  "MessageType_Cancel",
	24:  "MessageType_ClearSession",
	25:  "MessageType_ApplySettings",
	26:  "MessageType_ButtonRequest",
	27:  "MessageType_ButtonAck",
	28:  "MessageType_ApplyFlags",
	34:  "MessageType_BackupDevice",
	35:  "MessageType_EntropyRequest",
	36:  "MessageType_EntropyAck",
	41:  "MessageType_PassphraseRequest",
	42:  "MessageType_PassphraseAck",
	77:  "MessageType_PassphraseStateRequest",
	78:  "MessageType_PassphraseStateAck",
	45:  "MessageType_RecoveryDevice",
	46:  "MessageType_WordRequest",
	47:  "MessageType_WordAck",
	55:  "MessageType_GetFeatures",
	63:  "MessageType_SetU2FCounter",
	6:   "MessageType_FirmwareErase",
	7:   "MessageType_FirmwareUpload",
	8:   "MessageType_FirmwareRequest",
	32:  "MessageType_SelfTest",
	11:  "MessageType_GetPublicKey",
	12:  "MessageType_PublicKey",
	15:  "MessageType_SignTx",
	21:  "MessageType_TxRequest",
	22:  "MessageType_TxAck",
	29:  "MessageType_GetAddress",
	30:  "MessageType_Address",
	38:  "MessageType_SignMessage",
	39:  "MessageType_VerifyMessage",
	40:  "MessageType_MessageSignature",
	23:  "MessageType_CipherKeyValue",
	48:  "MessageType_CipheredKeyValue",
	53:  "MessageType_SignIdentity",
	54:  "MessageType_SignedIdentity",
	61:  "MessageType_GetECDHSessionKey",
	62:  "MessageType_ECDHSessionKey",
	71:  "MessageType_CosiCommit",
	72:  "MessageType_CosiCommitment",
	73:  "MessageType_CosiSign",
	74:  "MessageType_CosiSignature",
	100: "MessageType_DebugLinkDecision",
	101: "MessageType_DebugLinkGetState",
	102: "MessageType_DebugLinkState",
	103: "MessageType_DebugLinkStop",
	104: "MessageType_DebugLinkLog",
	110: "MessageType_DebugLinkMemoryRead",
	111: "MessageType_DebugLinkMemory",
	112: "MessageType_DebugLinkMemoryWrite",
	113: "MessageType_DebugLinkFlashErase",
	450: "MessageType_EthereumGetPublicKey",
	451: "MessageType_EthereumPublicKey",
	56:  "MessageType_EthereumGetAddress",
	57:  "MessageType_EthereumAddress",
	58:  "MessageType_EthereumSignTx",
	59:  "MessageType_EthereumTxRequest",
	60:  "MessageType_EthereumTxAck",
	64:  "MessageType_EthereumSignMessage",
	65:  "MessageType_EthereumVerifyMessage",
	66:  "MessageType_EthereumMessageSignature",
	67:  "MessageType_NEMGetAddress",
	68:  "MessageType_NEMAddress",
	69:  "MessageType_NEMSignTx",
	70:  "MessageType_NEMSignedTx",
	75:  "MessageType_NEMDecryptMessage",
	76:  "MessageType_NEMDecryptedMessage",
	114: "MessageType_LiskGetAddress",
	115: "MessageType_LiskAddress",
	116: "MessageType_LiskSignTx",
	117: "MessageType_LiskSignedTx",
	118: "MessageType_LiskSignMessage",
	119: "MessageType_LiskMessageSignature",
	120: "MessageType_LiskVerifyMessage",
	121: "MessageType_LiskGetPublicKey",
	122: "MessageType_LiskPublicKey",
	150: "MessageType_TezosGetAddress",
	151: "MessageType_TezosAddress",
	152: "MessageType_TezosSignTx",
	153: "MessageType_TezosSignedTx",
	154: "MessageType_TezosGetPublicKey",
	155: "MessageType_TezosPublicKey",
	202: "MessageType_StellarSignTx",
	203: "MessageType_StellarTxOpRequest",
	207: "MessageType_StellarGetAddress",
	208: "MessageType_StellarAddress",
	210: "MessageType_StellarCreateAccountOp",
	211: "MessageType_StellarPaymentOp",
	212: "MessageType_StellarPathPaymentOp",
	213: "MessageType_StellarManageOfferOp",
	214: "MessageType_StellarCreatePassiveOfferOp",
	215: "MessageType_StellarSetOptionsOp",
	216: "MessageType_StellarChangeTrustOp",
	217: "MessageType_StellarAllowTrustOp",
	218: "MessageType_StellarAccountMergeOp",
	220: "MessageType_StellarManageDataOp",
	221: "MessageType_StellarBumpSequenceOp",
	230: "MessageType_StellarSignedTx",
	250: "MessageType_TronGetAddress",
	251: "MessageType_TronAddress",
	252: "MessageType_TronSignTx",
	253: "MessageType_TronSignedTx",
	303: "MessageType_CardanoSignTx",
	304: "MessageType_CardanoTxRequest",
	305: "MessageType_CardanoGetPublicKey",
	306: "MessageType_CardanoPublicKey",
	307: "MessageType_CardanoGetAddress",
	308: "MessageType_CardanoAddress",
	309: "MessageType_CardanoTxAck",
	310: "MessageType_CardanoSignedTx",
	350: "MessageType_OntologyGetAddress",
	351: "MessageType_OntologyAddress",
	352: "MessageType_OntologyGetPublicKey",
	353: "MessageType_OntologyPublicKey",
	354: "MessageType_OntologySignTransfer",
	355: "MessageType_OntologySignedTransfer",
	356: "MessageType_OntologySignWithdrawOng",
	357: "MessageType_OntologySignedWithdrawOng",
	358: "MessageType_OntologySignOntIdRegister",
	359: "MessageType_OntologySignedOntIdRegister",
	360: "MessageType_OntologySignOntIdAddAttributes",
	361: "MessageType_OntologySignedOntIdAddAttributes",
	400: "MessageType_RippleGetAddress",
	401: "MessageType_RippleAddress",
	402: "MessageType_RippleSignTx",
	403: "MessageType_RippleSignedTx",
	501: "MessageType_MoneroTransactionInitRequest",
	502: "MessageType_MoneroTransactionInitAck",
	503: "MessageType_MoneroTransactionSetInputRequest",
	504: "MessageType_MoneroTransactionSetInputAck",
	505: "MessageType_MoneroTransactionInputsPermutationRequest",
	506: "MessageType_MoneroTransactionInputsPermutationAck",
	507: "MessageType_MoneroTransactionInputViniRequest",
	508: "MessageType_MoneroTransactionInputViniAck",
	509: "MessageType_MoneroTransactionAllInputsSetRequest",
	510: "MessageType_MoneroTransactionAllInputsSetAck",
	511: "MessageType_MoneroTransactionSetOutputRequest",
	512: "MessageType_MoneroTransactionSetOutputAck",
	513: "MessageType_MoneroTransactionAllOutSetRequest",
	514: "MessageType_MoneroTransactionAllOutSetAck",
	515: "MessageType_MoneroTransactionSignInputRequest",
	516: "MessageType_MoneroTransactionSignInputAck",
	517: "MessageType_MoneroTransactionFinalRequest",
	518: "MessageType_MoneroTransactionFinalAck",
	530: "MessageType_MoneroKeyImageExportInitRequest",
	531: "MessageType_MoneroKeyImageExportInitAck",
	532: "MessageType_MoneroKeyImageSyncStepRequest",
	533: "MessageType_MoneroKeyImageSyncStepAck",
	534: "MessageType_MoneroKeyImageSyncFinalRequest",
	535: "MessageType_MoneroKeyImageSyncFinalAck",
	540: "MessageType_MoneroGetAddress",
	541: "MessageType_MoneroAddress",
	542: "MessageType_MoneroGetWatchKey",
	543: "MessageType_MoneroWatchKey",
	546: "MessageType_DebugMoneroDiagRequest",
	547: "MessageType_DebugMoneroDiagAck",
	550: "MessageType_MoneroGetTxKeyRequest",
	551: "MessageType_MoneroGetTxKeyAck",
	552: "MessageType_MoneroLiveRefreshStartRequest",
	553: "MessageType_MoneroLiveRefreshStartAck",
	554: "MessageType_MoneroLiveRefreshStepRequest",
	555: "MessageType_MoneroLiveRefreshStepAck",
	556: "MessageType_MoneroLiveRefreshFinalRequest",
	557: "MessageType_MoneroLiveRefreshFinalAck",
	600: "MessageType_EosGetPublicKey",
	601: "MessageType_EosPublicKey",
	602: "MessageType_EosSignTx",
	603: "MessageType_EosTxActionRequest",
	604: "MessageType_EosTxActionAck",
	605: "MessageType_EosSignedTx",
	700: "MessageType_BinanceGetAddress",
	701: "MessageType_BinanceAddress",
	702: "MessageType_BinanceGetPublicKey",
	703: "MessageType_BinancePublicKey",
	704: "MessageType_BinanceSignTx",
	705: "MessageType_BinanceTxRequest",
	706: "MessageType_BinanceTransferMsg",
	707: "MessageType_BinanceOrderMsg",
	708: "MessageType_BinanceCancelMsg",
	709: "MessageType_BinanceSignedTx",
}

var MessageType_value = map[string]int32{
	"MessageType_Initialize":                                0,
	"MessageType_Ping":                                      1,
	"MessageType_Success":                                   2,
	"MessageType_Failure":                                   3,
	"MessageType_ChangePin":                                 4,
	"MessageType_WipeDevice":                                5,
	"MessageType_GetEntropy":                                9,
	"MessageType_Entropy":                                   10,
	"MessageType_LoadDevice":                                13,
	"MessageType_ResetDevice":                               14,
	"MessageType_Features":                                  17,
	"MessageType_PinMatrixRequest":                          18,
	"MessageType_PinMatrixAck":                              19,
	"MessageType_Cancel":                                    20,
	"MessageType_ClearSession":                              24,
	"MessageType_ApplySettings":                             25,
	"MessageType_ButtonRequest":                             26,
	"MessageType_ButtonAck":                                 27,
	"MessageType_ApplyFlags":                                28,
	"MessageType_BackupDevice":                              34,
	"MessageType_EntropyRequest":                            35,
	"MessageType_EntropyAck":                                36,
	"MessageType_PassphraseRequest":                         41,
	"MessageType_PassphraseAck":                             42,
	"MessageType_PassphraseStateRequest":                    77,
	"MessageType_PassphraseStateAck":                        78,
	"MessageType_RecoveryDevice":                            45,
	"MessageType_WordRequest":                               46,
	"MessageType_WordAck":                                   47,
	"MessageType_GetFeatures":                               55,
	"MessageType_SetU2FCounter":                             63,
	"MessageType_FirmwareErase":                             6,
	"MessageType_FirmwareUpload":                            7,
	"MessageType_FirmwareRequest":                           8,
	"MessageType_SelfTest":                                  32,
	"MessageType_GetPublicKey":                              11,
	"MessageType_PublicKey":                                 12,
	"MessageType_SignTx":                                    15,
	"MessageType_TxRequest":                                 21,
	"MessageType_TxAck":                                     22,
	"MessageType_GetAddress":                                29,
	"MessageType_Address":                                   30,
	"MessageType_SignMessage":                               38,
	"MessageType_VerifyMessage":                             39,
	"MessageType_MessageSignature":                          40,
	"MessageType_CipherKeyValue":                            23,
	"MessageType_CipheredKeyValue":                          48,
	"MessageType_SignIdentity":                              53,
	"MessageType_SignedIdentity":                            54,
	"MessageType_GetECDHSessionKey":                         61,
	"MessageType_ECDHSessionKey":                            62,
	"MessageType_CosiCommit":                                71,
	"MessageType_CosiCommitment":                            72,
	"MessageType_CosiSign":                                  73,
	"MessageType_CosiSignature":                             74,
	"MessageType_DebugLinkDecision":                         100,
	"MessageType_DebugLinkGetState":                         101,
	"MessageType_DebugLinkState":                            102,
	"MessageType_DebugLinkStop":                             103,
	"MessageType_DebugLinkLog":                              104,
	"MessageType_DebugLinkMemoryRead":                       110,
	"MessageType_DebugLinkMemory":                           111,
	"MessageType_DebugLinkMemoryWrite":                      112,
	"MessageType_DebugLinkFlashErase":                       113,
	"MessageType_EthereumGetPublicKey":                      450,
	"MessageType_EthereumPublicKey":                         451,
	"MessageType_EthereumGetAddress":                        56,
	"MessageType_EthereumAddress":                           57,
	"MessageType_EthereumSignTx":                            58,
	"MessageType_EthereumTxRequest":                         59,
	"MessageType_EthereumTxAck":                             60,
	"MessageType_EthereumSignMessage":                       64,
	"MessageType_EthereumVerifyMessage":                     65,
	"MessageType_EthereumMessageSignature":                  66,
	"MessageType_NEMGetAddress":                             67,
	"MessageType_NEMAddress":                                68,
	"MessageType_NEMSignTx":                                 69,
	"MessageType_NEMSignedTx":                               70,
	"MessageType_NEMDecryptMessage":                         75,
	"MessageType_NEMDecryptedMessage":                       76,
	"MessageType_LiskGetAddress":                            114,
	"MessageType_LiskAddress":                               115,
	"MessageType_LiskSignTx":                                116,
	"MessageType_LiskSignedTx":                              117,
	"MessageType_LiskSignMessage":                           118,
	"MessageType_LiskMessageSignature":                      119,
	"MessageType_LiskVerifyMessage":                         120,
	"MessageType_LiskGetPublicKey":                          121,
	"MessageType_LiskPublicKey":                             122,
	"MessageType_TezosGetAddress":                           150,
	"MessageType_TezosAddress":                              151,
	"MessageType_TezosSignTx":                               152,
	"MessageType_TezosSignedTx":                             153,
	"MessageType_TezosGetPublicKey":                         154,
	"MessageType_TezosPublicKey":                            155,
	"MessageType_StellarSignTx":                             202,
	"MessageType_StellarTxOpRequest":                        203,
	"MessageType_StellarGetAddress":                         207,
	"MessageType_StellarAddress":                            208,
	"MessageType_StellarCreateAccountOp":                    210,
	"MessageType_StellarPaymentOp":                          211,
	"MessageType_StellarPathPaymentOp":                      212,
	"MessageType_StellarManageOfferOp":                      213,
	"MessageType_StellarCreatePassiveOfferOp":               214,
	"MessageType_StellarSetOptionsOp":                       215,
	"MessageType_StellarChangeTrustOp":                      216,
	"MessageType_StellarAllowTrustOp":                       217,
	"MessageType_StellarAccountMergeOp":                     218,
	"MessageType_StellarManageDataOp":                       220,
	"MessageType_StellarBumpSequenceOp":                     221,
	"MessageType_StellarSignedTx":                           230,
	"MessageType_TronGetAddress":                            250,
	"MessageType_TronAddress":                               251,
	"MessageType_TronSignTx":                                252,
	"MessageType_TronSignedTx":                              253,
	"MessageType_CardanoSignTx":                             303,
	"MessageType_CardanoTxRequest":                          304,
	"MessageType_CardanoGetPublicKey":                       305,
	"MessageType_CardanoPublicKey":                          306,
	"MessageType_CardanoGetAddress":                         307,
	"MessageType_CardanoAddress":                            308,
	"MessageType_CardanoTxAck":                              309,
	"MessageType_CardanoSignedTx":                           310,
	"MessageType_OntologyGetAddress":                        350,
	"MessageType_OntologyAddress":                           351,
	"MessageType_OntologyGetPublicKey":                      352,
	"MessageType_OntologyPublicKey":                         353,
	"MessageType_OntologySignTransfer":                      354,
	"MessageType_OntologySignedTransfer":                    355,
	"MessageType_OntologySignWithdrawOng":                   356,
	"MessageType_OntologySignedWithdrawOng":                 357,
	"MessageType_OntologySignOntIdRegister":                 358,
	"MessageType_OntologySignedOntIdRegister":               359,
	"MessageType_OntologySignOntIdAddAttributes":            360,
	"MessageType_OntologySignedOntIdAddAttributes":          361,
	"MessageType_RippleGetAddress":                          400,
	"MessageType_RippleAddress":                             401,
	"MessageType_RippleSignTx":                              402,
	"MessageType_RippleSignedTx":                            403,
	"MessageType_MoneroTransactionInitRequest":              501,
	"MessageType_MoneroTransactionInitAck":                  502,
	"MessageType_MoneroTransactionSetInputRequest":          503,
	"MessageType_MoneroTransactionSetInputAck":              504,
	"MessageType_MoneroTransactionInputsPermutationRequest": 505,
	"MessageType_MoneroTransactionInputsPermutationAck":     506,
	"MessageType_MoneroTransactionInputViniRequest":         507,
	"MessageType_MoneroTransactionInputViniAck":             508,
	"MessageType_MoneroTransactionAllInputsSetRequest":      509,
	"MessageType_MoneroTransactionAllInputsSetAck":          510,
	"MessageType_MoneroTransactionSetOutputRequest":         511,
	"MessageType_MoneroTransactionSetOutputAck":             512,
	"MessageType_MoneroTransactionAllOutSetRequest":         513,
	"MessageType_MoneroTransactionAllOutSetAck":             514,
	"MessageType_MoneroTransactionSignInputRequest":         515,
	"MessageType_MoneroTransactionSignInputAck":             516,
	"MessageType_MoneroTransactionFinalRequest":             517,
	"MessageType_MoneroTransactionFinalAck":                 518,
	"MessageType_MoneroKeyImageExportInitRequest":           530,
	"MessageType_MoneroKeyImageExportInitAck":               531,
	"MessageType_MoneroKeyImageSyncStepRequest":             532,
	"MessageType_MoneroKeyImageSyncStepAck":                 533,
	"MessageType_MoneroKeyImageSyncFinalRequest":            534,
	"MessageType_MoneroKeyImageSyncFinalAck":                535,
	"MessageType_MoneroGetAddress":                          540,
	"MessageType_MoneroAddress":                             541,
	"MessageType_MoneroGetWatchKey":                         542,
	"MessageType_MoneroWatchKey":                            543,
	"MessageType_DebugMoneroDiagRequest":                    546,
	"MessageType_DebugMoneroDiagAck":                        547,
	"MessageType_MoneroGetTxKeyRequest":                     550,
	"MessageType_MoneroGetTxKeyAck":                         551,
	"MessageType_MoneroLiveRefreshStartRequest":             552,
	"MessageType_MoneroLiveRefreshStartAck":                 553,
	"MessageType_MoneroLiveRefreshStepRequest":              554,
	"MessageType_MoneroLiveRefreshStepAck":                  555,
	"MessageType_MoneroLiveRefreshFinalRequest":             556,
	"MessageType_MoneroLiveRefreshFinalAck":                 557,
	"MessageType_EosGetPublicKey":                           600,
	"MessageType_EosPublicKey":                              601,
	"MessageType_EosSignTx":                                 602,
	"MessageType_EosTxActionRequest":                        603,
	"MessageType_EosTxActionAck":                            604,
	"MessageType_EosSignedTx":                               605,
	"MessageType_BinanceGetAddress":                         700,
	"MessageType_BinanceAddress":                            701,
	"MessageType_BinanceGetPublicKey":                       702,
	"MessageType_BinancePublicKey":                          703,
	"MessageType_BinanceSignTx":                             704,
	"MessageType_BinanceTxRequest":                          705,
	"MessageType_BinanceTransferMsg":                        706,
	"MessageType_BinanceOrderMsg":                           707,
	"MessageType_BinanceCancelMsg":                          708,
	"MessageType_BinanceSignedTx":                           709,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

var E_WireIn = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50002,
	Name:          "hw.trezor.messages.wire_in",
	Tag:           "varint,50002,opt,name=wire_in",
	Filename:      "messages.proto",
}

var E_WireOut = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50003,
	Name:          "hw.trezor.messages.wire_out",
	Tag:           "varint,50003,opt,name=wire_out",
	Filename:      "messages.proto",
}

var E_WireDebugIn = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50004,
	Name:          "hw.trezor.messages.wire_debug_in",
	Tag:           "varint,50004,opt,name=wire_debug_in",
	Filename:      "messages.proto",
}

var E_WireDebugOut = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50005,
	Name:          "hw.trezor.messages.wire_debug_out",
	Tag:           "varint,50005,opt,name=wire_debug_out",
	Filename:      "messages.proto",
}

var E_WireTiny = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50006,
	Name:          "hw.trezor.messages.wire_tiny",
	Tag:           "varint,50006,opt,name=wire_tiny",
	Filename:      "messages.proto",
}

var E_WireBootloader = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50007,
	Name:          "hw.trezor.messages.wire_bootloader",
	Tag:           "varint,50007,opt,name=wire_bootloader",
	Filename:      "messages.proto",
}

var E_WireNoFsm = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50008,
	Name:          "hw.trezor.messages.wire_no_fsm",
	Tag:           "varint,50008,opt,name=wire_no_fsm",
	Filename:      "messages.proto",
}

func init() {
	proto.RegisterEnum("hw.trezor.messages.MessageType", MessageType_name, MessageType_value)
	proto.RegisterExtension(E_WireIn)
	proto.RegisterExtension(E_WireOut)
	proto.RegisterExtension(E_WireDebugIn)
	proto.RegisterExtension(E_WireDebugOut)
	proto.RegisterExtension(E_WireTiny)
	proto.RegisterExtension(E_WireBootloader)
	proto.RegisterExtension(E_WireNoFsm)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 2430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x9a, 0xd9, 0x73, 0x1c, 0xc5,
	0x1d, 0xc7, 0xb3, 0xab, 0x11, 0x88, 0xf6, 0x41, 0x23, 0xb0, 0x2d, 0xaf, 0x2f, 0xf9, 0xc0, 0x96,
	0x2f, 0xd9, 0x10, 0x0c, 0x44, 0x38, 0x60, 0x69, 0xb5, 0x12, 0x8a, 0xb5, 0x5a, 0x97, 0x76, 0xb1,
	0x1f, 0x5d, 0xa3, 0x9d, 0xd6, 0x6e, 0x97, 0x67, 0x67, 0x86, 0x9e, 0x1e, 0x49, 0xeb, 0xa7, 0x9c,
	0x3c, 0x13, 0x48, 0xc0, 0xb9, 0xa9, 0xa4, 0x2a, 0x21, 0x57, 0x85, 0x1c, 0x4e, 0x25, 0x55, 0x39,
	0x08, 0x24, 0x2f, 0xc9, 0x43, 0x52, 0x9c, 0x86, 0x40, 0xee, 0x90, 0xe4, 0x0f, 0xc8, 0xc5, 0x91,
	0xa4, 0x7a, 0xa6, 0xbb, 0xe7, 0xd8, 0xdf, 0xae, 0x36, 0x6f, 0x58, 0xf3, 0xf9, 0x7d, 0x7f, 0x47,
	0xff, 0xfa, 0x37, 0xdd, 0xb3, 0xa0, 0xcd, 0x2d, 0xe2, 0xfb, 0x66, 0x83, 0xf8, 0xe3, 0x1e, 0x73,
	0xb9, 0x3b, 0x3c, 0xdc, 0x5c, 0x1d, 0xe7, 0x8c, 0x5c, 0x76, 0xd9, 0xb8, 0x7a, 0x52, 0x18, 0x6d,
	0xb8, 0x6e, 0xc3, 0x26, 0x27, 0x42, 0x62, 0x29, 0x58, 0x3e, 0x61, 0x11, 0xbf, 0xce, 0xa8, 0xc7,
	0x5d, 0x16, 0x59, 0x1d, 0xf9, 0xfe, 0x7d, 0x68, 0x43, 0x39, 0xc2, 0x6b, 0x6d, 0x8f, 0x0c, 0x1f,
	0x40, 0x5b, 0x13, 0xff, 0xbc, 0x38, 0xe7, 0x50, 0x4e, 0x4d, 0x9b, 0x5e, 0x26, 0xf8, 0x5d, 0x85,
	0xa1, 0x87, 0xaf, 0x8e, 0xe4, 0x9e, 0xba, 0x3a, 0x92, 0x1b, 0x2e, 0x20, 0x9c, 0xa4, 0xce, 0x51,
	0xa7, 0x81, 0x73, 0x05, 0x43, 0x3c, 0x1f, 0xde, 0x85, 0x6e, 0x4e, 0x3e, 0xab, 0x06, 0xf5, 0x3a,
	0xf1, 0x7d, 0x9c, 0x2f, 0x18, 0x57, 0x80, 0xc7, 0x33, 0x26, 0xb5, 0x03, 0x46, 0xf0, 0x80, 0x7c,
	0xbc, 0x07, 0x6d, 0x49, 0x3e, 0x2e, 0x36, 0x4d, 0xa7, 0x41, 0xce, 0x51, 0x07, 0x1b, 0x52, 0x7e,
	0x34, 0x1d, 0xe0, 0x05, 0xea, 0x91, 0x69, 0xb2, 0x42, 0xeb, 0x04, 0x0f, 0xc2, 0xc4, 0x2c, 0xe1,
	0x25, 0x87, 0x33, 0xd7, 0x6b, 0xe3, 0x1b, 0xe0, 0x10, 0xd5, 0x63, 0x24, 0x63, 0xc8, 0x08, 0xcc,
	0xbb, 0xa6, 0x25, 0x5d, 0x6c, 0x92, 0x02, 0x7b, 0xd1, 0xb6, 0x24, 0xb1, 0x48, 0x7c, 0xc2, 0x25,
	0xb2, 0x59, 0x22, 0xbb, 0xd1, 0x2d, 0xa9, 0x3c, 0x89, 0xc9, 0x03, 0x46, 0x7c, 0x7c, 0x93, 0x74,
	0x72, 0x10, 0xed, 0xcc, 0x94, 0xb0, 0x6c, 0x72, 0x46, 0xd7, 0x16, 0xc9, 0x83, 0x01, 0xf1, 0x39,
	0x1e, 0x96, 0xdc, 0x11, 0x34, 0x02, 0x72, 0x93, 0xf5, 0x4b, 0xf8, 0xe6, 0xc2, 0x46, 0xb5, 0x24,
	0x4f, 0x47, 0x81, 0x0f, 0xa7, 0x8a, 0x67, 0x3a, 0x75, 0x62, 0xe3, 0x5b, 0x12, 0x0b, 0xb7, 0x2f,
	0xad, 0x56, 0xb4, 0x89, 0xc9, 0xaa, 0xc4, 0xf7, 0xa9, 0xeb, 0xe0, 0x11, 0x19, 0xf9, 0x7e, 0xb4,
	0x3d, 0xc9, 0x4c, 0x7a, 0x9e, 0xdd, 0xae, 0x12, 0xce, 0xa9, 0xd3, 0xf0, 0xf1, 0x76, 0x18, 0x9a,
	0x0a, 0x38, 0x77, 0x1d, 0x15, 0x7b, 0x41, 0xc6, 0x7e, 0x28, 0xbd, 0x98, 0x11, 0x24, 0x02, 0xdf,
	0xd1, 0x11, 0xf8, 0xd6, 0x0e, 0x97, 0x33, 0xb6, 0xd9, 0xf0, 0xf1, 0x4e, 0xe9, 0x2f, 0x13, 0xf8,
	0x94, 0x59, 0xbf, 0x14, 0x78, 0xb2, 0xe4, 0xfb, 0x24, 0x73, 0x00, 0x15, 0x80, 0x65, 0x55, 0x41,
	0xed, 0x87, 0x57, 0x57, 0x52, 0x22, 0xaa, 0x03, 0x52, 0xe7, 0x10, 0xda, 0x95, 0x2a, 0xb9, 0xe9,
	0xfb, 0x5e, 0x93, 0x99, 0x3e, 0x51, 0x52, 0x87, 0xa5, 0xd4, 0xd1, 0x74, 0x11, 0x62, 0x50, 0xa8,
	0x1d, 0xc9, 0xe4, 0x78, 0x0c, 0xed, 0x83, 0xe1, 0x2a, 0x37, 0xb9, 0x96, 0x2e, 0x4b, 0xe9, 0x93,
	0x68, 0x77, 0x0f, 0x5a, 0xe8, 0x2f, 0x64, 0xf4, 0x33, 0xd9, 0x2f, 0x92, 0xba, 0xbb, 0x42, 0x58,
	0x5b, 0xd6, 0xe8, 0x38, 0xdc, 0xb9, 0x17, 0x5c, 0x66, 0x29, 0xd7, 0xe3, 0xf0, 0x0e, 0x15, 0x88,
	0xf0, 0x77, 0x02, 0x56, 0x98, 0x25, 0x5c, 0xf7, 0xf6, 0x5d, 0x70, 0x73, 0x54, 0x09, 0x7f, 0xe0,
	0xf6, 0x99, 0xa2, 0x1b, 0x38, 0x9c, 0x30, 0x7c, 0x9f, 0xae, 0x72, 0x0a, 0x9a, 0xa1, 0xac, 0xb5,
	0x6a, 0x32, 0x52, 0x12, 0x49, 0xe2, 0xeb, 0xa2, 0x9e, 0xfd, 0x9e, 0x00, 0xc7, 0xd2, 0x89, 0x29,
	0xf0, 0x01, 0xcf, 0x76, 0x4d, 0x0b, 0x5f, 0x9f, 0x20, 0x0f, 0xa3, 0x1d, 0x10, 0xa9, 0x12, 0x1c,
	0x2a, 0x0c, 0x5d, 0x51, 0xe8, 0xbe, 0xf4, 0xf6, 0xac, 0x12, 0x7b, 0xb9, 0x26, 0x98, 0xd1, 0x84,
	0x5c, 0xa6, 0xe7, 0x66, 0x09, 0x3f, 0x17, 0x2c, 0xd9, 0xb4, 0x7e, 0x96, 0xb4, 0xf1, 0x06, 0x99,
	0x45, 0x66, 0x5e, 0xc5, 0xc0, 0x46, 0x59, 0xcd, 0x9d, 0xe9, 0x3d, 0x59, 0xa5, 0x0d, 0xa7, 0xb6,
	0x86, 0x6f, 0x84, 0xcd, 0x6b, 0x7a, 0xfb, 0x6f, 0x91, 0xe6, 0x3b, 0xd0, 0x4d, 0x69, 0x40, 0x2c,
	0xc5, 0xd6, 0xae, 0x93, 0x6e, 0xd2, 0xb2, 0x98, 0x98, 0xb6, 0xbb, 0xe0, 0x49, 0xa7, 0x1e, 0xef,
	0x96, 0xea, 0x99, 0xb5, 0x14, 0xc1, 0xc9, 0x7f, 0xe3, 0x83, 0xf0, 0x5a, 0x9e, 0x27, 0x8c, 0x2e,
	0xb7, 0x15, 0x74, 0x48, 0x42, 0x99, 0x61, 0x26, 0xff, 0x5b, 0xc8, 0x85, 0x9d, 0x81, 0xc7, 0xa4,
	0xbf, 0x4c, 0x8f, 0x16, 0xa9, 0xd7, 0x24, 0xec, 0x2c, 0x69, 0x9f, 0x37, 0xed, 0x80, 0xe0, 0x6d,
	0xb0, 0x5a, 0x44, 0x11, 0x4b, 0x73, 0x27, 0xa5, 0x5a, 0x66, 0x7d, 0x84, 0xbb, 0x39, 0x8b, 0x38,
	0x9c, 0xf2, 0x36, 0x3e, 0x05, 0xcf, 0x04, 0xc1, 0x10, 0x4b, 0x53, 0x77, 0xea, 0x41, 0xb5, 0x2b,
	0xfb, 0xca, 0x28, 0x4e, 0xdf, 0x2f, 0x07, 0xa3, 0x58, 0xcd, 0xf7, 0x76, 0x19, 0x31, 0x69, 0xea,
	0x5e, 0x78, 0xc4, 0x14, 0x5d, 0x9f, 0x16, 0xdd, 0x56, 0x8b, 0x72, 0x3c, 0x0b, 0xeb, 0xc4, 0x44,
	0x8b, 0x38, 0x1c, 0xdf, 0x2f, 0x75, 0x32, 0xef, 0x10, 0x41, 0x89, 0x04, 0xf0, 0x1c, 0xbc, 0x36,
	0xea, 0x79, 0x54, 0xf3, 0xf7, 0x49, 0x91, 0x13, 0xe9, 0xdc, 0xa6, 0xc9, 0x52, 0xd0, 0x98, 0xa7,
	0xce, 0xa5, 0x69, 0x52, 0xa7, 0xe1, 0xdc, 0xb7, 0x0a, 0x1b, 0x9f, 0x48, 0x0e, 0x92, 0xa3, 0x5d,
	0x0c, 0x66, 0x09, 0x0f, 0x87, 0x0f, 0x26, 0x85, 0x21, 0x65, 0x90, 0x4d, 0x44, 0xc3, 0x11, 0xb9,
	0x5c, 0x30, 0x9e, 0x04, 0x02, 0x4d, 0x50, 0xae, 0x87, 0x1b, 0x05, 0xe3, 0x09, 0x60, 0x39, 0x35,
	0x34, 0xef, 0x36, 0x70, 0x53, 0x0a, 0x1d, 0x46, 0x7b, 0x40, 0xa6, 0x4c, 0x5a, 0x2e, 0x6b, 0x2f,
	0x12, 0xd3, 0xc2, 0x8e, 0x94, 0xbb, 0x35, 0x3d, 0x0c, 0x32, 0x28, 0x76, 0xa5, 0xe2, 0x11, 0x34,
	0xda, 0x03, 0xbb, 0xc0, 0x28, 0x27, 0xd8, 0x93, 0x92, 0xdd, 0xbc, 0xcf, 0xd8, 0xa6, 0xdf, 0x8c,
	0x06, 0xd7, 0x83, 0x12, 0x3d, 0x9a, 0x96, 0x2d, 0x71, 0xd1, 0xc2, 0x41, 0x2b, 0x35, 0x43, 0x9e,
	0x19, 0x90, 0xeb, 0x38, 0x96, 0xae, 0xb8, 0x82, 0x63, 0xf2, 0x59, 0x75, 0x3c, 0x1a, 0x4b, 0xbf,
	0x16, 0x12, 0xb2, 0x6a, 0x6b, 0xdf, 0x2d, 0x35, 0x33, 0xe9, 0x2b, 0x52, 0x61, 0xef, 0x81, 0x77,
	0xa4, 0xc2, 0xe4, 0x98, 0x9a, 0x80, 0xdf, 0x88, 0x8a, 0x8a, 0xc7, 0xd5, 0x3d, 0x52, 0x2e, 0xb3,
	0xd0, 0x31, 0x28, 0xc6, 0xd6, 0x69, 0xa9, 0x96, 0x29, 0x63, 0xd2, 0xa7, 0x1a, 0x2c, 0x67, 0x24,
	0x7a, 0x14, 0xed, 0x85, 0xd0, 0xf4, 0x14, 0x9a, 0x94, 0xf0, 0x38, 0x3a, 0x00, 0xc1, 0x1d, 0xd3,
	0x68, 0x0a, 0x0e, 0x76, 0xa1, 0x54, 0x4e, 0xd4, 0xb1, 0x08, 0xcf, 0xd8, 0x85, 0x52, 0x59, 0x11,
	0xd3, 0xf0, 0x91, 0x75, 0xa1, 0x54, 0x96, 0xd5, 0x2b, 0xc1, 0x6f, 0x4c, 0x09, 0x10, 0xab, 0xb6,
	0x86, 0x67, 0xe0, 0x01, 0xb4, 0x50, 0x2a, 0x4f, 0x93, 0x3a, 0x6b, 0x7b, 0x5c, 0xe5, 0x78, 0x16,
	0xae, 0x5d, 0x0c, 0x12, 0x4b, 0xa1, 0xf3, 0xf0, 0xd2, 0xce, 0x53, 0xff, 0x52, 0x22, 0x3f, 0x06,
	0x07, 0x27, 0x28, 0x85, 0xf8, 0x5d, 0xce, 0xc3, 0xd4, 0xbf, 0x24, 0x33, 0xe4, 0xf0, 0xe9, 0x4c,
	0x11, 0x61, 0x8a, 0x81, 0x54, 0xc9, 0x34, 0xa4, 0x62, 0x54, 0xd4, 0x2b, 0x52, 0x2a, 0xb3, 0x1f,
	0x05, 0xd6, 0xb1, 0x80, 0xab, 0x70, 0xd5, 0x04, 0x9b, 0xee, 0x8c, 0x35, 0xf8, 0x8d, 0x22, 0x4b,
	0x11, 0xef, 0xaf, 0x36, 0x3c, 0x50, 0x05, 0x17, 0x43, 0x97, 0xf5, 0xc9, 0x3d, 0x95, 0x48, 0x8d,
	0x5c, 0x76, 0xfd, 0x44, 0x61, 0x1f, 0xcb, 0x69, 0xb1, 0x91, 0x0e, 0x4e, 0x41, 0x8f, 0xe7, 0xf4,
	0x3b, 0x6c, 0x5b, 0x07, 0x24, 0x8b, 0x7b, 0x25, 0xa7, 0x5f, 0x16, 0xdb, 0x41, 0x26, 0x2c, 0xef,
	0x27, 0x72, 0x7a, 0x34, 0xec, 0x82, 0xc2, 0x8a, 0xe3, 0xff, 0x64, 0x4e, 0x8f, 0x86, 0x42, 0x07,
	0x19, 0x63, 0x9f, 0xca, 0xe9, 0xfe, 0x49, 0x9f, 0xe2, 0x38, 0xb1, 0x6d, 0x93, 0xc9, 0xe0, 0x7e,
	0x9e, 0xd3, 0x0d, 0xb9, 0x1b, 0xa0, 0x6a, 0x6b, 0x15, 0x4f, 0xcd, 0x86, 0x5f, 0x74, 0x89, 0x50,
	0xa2, 0x89, 0xd2, 0xfd, 0xb2, 0x4b, 0x84, 0x92, 0x54, 0xd8, 0xaf, 0x94, 0xe0, 0xf1, 0xf4, 0x91,
	0x5a, 0x62, 0x45, 0x46, 0xc2, 0x23, 0x72, 0x5d, 0x1c, 0x38, 0x2b, 0x1e, 0x7e, 0x2e, 0xa7, 0xa7,
	0xd8, 0x4e, 0x00, 0x3f, 0x67, 0xb6, 0xc5, 0x4b, 0xb7, 0xe2, 0xe1, 0xe7, 0x73, 0x7a, 0xea, 0x8c,
	0x82, 0x20, 0x6f, 0xc6, 0xf0, 0x0b, 0xbd, 0xe1, 0xb2, 0xe9, 0x98, 0x0d, 0x52, 0x59, 0x5e, 0x26,
	0xac, 0xe2, 0xe1, 0x17, 0x15, 0x7c, 0x3b, 0x3a, 0xd4, 0x35, 0x62, 0x71, 0xc6, 0xa7, 0x2b, 0xda,
	0xe6, 0xa5, 0x9c, 0xde, 0x11, 0x7b, 0xa0, 0x75, 0x20, 0xbc, 0xe2, 0x71, 0xea, 0x3a, 0x7e, 0xc5,
	0xc3, 0x2f, 0xf7, 0x0e, 0x26, 0xba, 0x45, 0xd7, 0x58, 0xe0, 0x8b, 0xc8, 0xaf, 0xf5, 0x16, 0x9e,
	0xb4, 0x6d, 0x77, 0x55, 0xb1, 0xaf, 0x28, 0xf6, 0x58, 0x7a, 0x10, 0x2b, 0x36, 0x2a, 0x72, 0x99,
	0xb0, 0x06, 0xa9, 0x78, 0xf8, 0xd5, 0xde, 0xca, 0x51, 0x4d, 0xa6, 0x4d, 0x6e, 0x56, 0x3c, 0xfc,
	0x5a, 0x6f, 0xe5, 0xa9, 0xa0, 0xe5, 0x55, 0x45, 0x03, 0x39, 0x75, 0xa1, 0xfc, 0x7a, 0x4e, 0xef,
	0xe4, 0x1d, 0x5d, 0x9a, 0x32, 0xdc, 0x0d, 0x6f, 0xe4, 0xf4, 0xb4, 0x49, 0xf7, 0x38, 0x73, 0x9d,
	0x44, 0xa3, 0xbd, 0x99, 0xd3, 0x83, 0x6b, 0x5b, 0x16, 0x53, 0xcc, 0x5b, 0x39, 0x7d, 0x48, 0xde,
	0x9a, 0x65, 0xe4, 0x26, 0x78, 0xbb, 0xdb, 0x56, 0x97, 0x48, 0x18, 0xd2, 0x3b, 0x5d, 0xf6, 0x53,
	0xd1, 0x64, 0x96, 0xe9, 0xb8, 0x52, 0xea, 0x1b, 0x79, 0xb8, 0x49, 0x25, 0x15, 0xbf, 0x69, 0x9f,
	0xca, 0xeb, 0x0f, 0x03, 0x7b, 0x00, 0x30, 0xb5, 0xe3, 0xbf, 0xd9, 0x5b, 0x34, 0x06, 0xbf, 0x95,
	0x87, 0xb7, 0x68, 0x2c, 0xaa, 0xaa, 0xf2, 0xed, 0x3c, 0xbc, 0x45, 0x25, 0xa9, 0xb0, 0xef, 0xe4,
	0xf5, 0x3b, 0x76, 0x04, 0x4c, 0x47, 0x9c, 0x07, 0xae, 0xe6, 0xe1, 0x45, 0x4d, 0x54, 0x26, 0xac,
	0xe0, 0x77, 0x95, 0x58, 0x66, 0xd6, 0x54, 0x1c, 0xee, 0xda, 0x6e, 0xa3, 0x9d, 0x08, 0xef, 0x37,
	0x5d, 0x24, 0x15, 0xaa, 0xb8, 0xdf, 0xe6, 0xf5, 0x15, 0x7e, 0xb4, 0x8b, 0x64, 0x5c, 0x9d, 0xdf,
	0xe5, 0xe1, 0x73, 0x9a, 0x82, 0x63, 0xf2, 0xf7, 0xeb, 0xc8, 0x86, 0x8b, 0xcd, 0x4c, 0xc7, 0x5f,
	0x26, 0x0c, 0xff, 0x41, 0xc9, 0x66, 0xc6, 0x58, 0x12, 0x26, 0x96, 0xc6, 0xff, 0xa8, 0xb4, 0xc7,
	0xd1, 0xfe, 0x6e, 0xf8, 0x05, 0xca, 0x9b, 0x16, 0x33, 0x57, 0x2b, 0x4e, 0x03, 0xff, 0x49, 0xc9,
	0x9f, 0x44, 0xb7, 0x76, 0x97, 0x4f, 0x5a, 0xfc, 0x39, 0xaf, 0x3f, 0x3e, 0x74, 0xb5, 0xa8, 0x38,
	0x7c, 0xce, 0x5a, 0x24, 0x0d, 0xea, 0x8b, 0xbb, 0xfc, 0x1b, 0x79, 0x78, 0xae, 0xa5, 0x7d, 0xa4,
	0x6d, 0xfe, 0xa2, 0xbc, 0x9c, 0x42, 0x47, 0x7a, 0x7a, 0x99, 0xb4, 0xac, 0x49, 0xce, 0x19, 0x5d,
	0x0a, 0x38, 0xf1, 0xf1, 0x5f, 0x95, 0xab, 0xbb, 0xd0, 0xb1, 0x75, 0x5c, 0xa5, 0x0d, 0xff, 0x96,
	0xd7, 0xa7, 0x85, 0xd4, 0x26, 0x58, 0xa4, 0x9e, 0x67, 0x93, 0x44, 0xef, 0x3c, 0x3c, 0x00, 0xbf,
	0x6f, 0x23, 0x50, 0x51, 0x1f, 0x1d, 0x80, 0x3b, 0x3b, 0xa2, 0xe4, 0x6e, 0x7e, 0x64, 0x00, 0xde,
	0x25, 0x31, 0x14, 0x36, 0xf6, 0xa3, 0x0a, 0x7b, 0x37, 0x1a, 0x4b, 0xdd, 0x9f, 0x5d, 0x87, 0x30,
	0x37, 0x5c, 0x79, 0xb3, 0x2e, 0x66, 0xfc, 0x9c, 0x43, 0xb9, 0x1a, 0x00, 0x7f, 0x1f, 0xd0, 0x17,
	0xbb, 0x03, 0xeb, 0x1a, 0x89, 0x6d, 0xf6, 0x0f, 0x65, 0x90, 0xa9, 0x5c, 0x87, 0x41, 0x95, 0xf0,
	0x39, 0xc7, 0x0b, 0xb4, 0xa7, 0x7f, 0x2a, 0xc3, 0xf5, 0xc2, 0x53, 0x86, 0xc2, 0xdb, 0xbf, 0x94,
	0xd1, 0x19, 0x74, 0x6a, 0x9d, 0xf0, 0xbc, 0x80, 0xfb, 0xe7, 0x08, 0x6b, 0x05, 0xdc, 0x14, 0x7f,
	0x50, 0x6e, 0xff, 0xad, 0x14, 0x4e, 0xa3, 0xdb, 0xfe, 0x3f, 0x05, 0xe1, 0xff, 0x4d, 0x65, 0x7d,
	0x37, 0x3a, 0xbe, 0xbe, 0xf5, 0x79, 0xea, 0x50, 0xe5, 0xf7, 0x2d, 0x65, 0x79, 0x07, 0x3a, 0xdc,
	0x9f, 0xa5, 0xf0, 0xf7, 0xb6, 0xb2, 0xba, 0x07, 0x9d, 0xec, 0x69, 0x35, 0x69, 0xdb, 0x51, 0xc0,
	0x55, 0xa2, 0x2b, 0xfc, 0x4e, 0xbf, 0x4b, 0x93, 0x34, 0x16, 0x5e, 0xff, 0xd3, 0x6f, 0x96, 0xe2,
	0x98, 0x10, 0xf0, 0xc4, 0xa2, 0xfe, 0xb7, 0xdf, 0x2c, 0xb5, 0xa5, 0xf0, 0xf7, 0x7e, 0xa3, 0x4f,
	0x7f, 0x93, 0xb6, 0x5d, 0x09, 0x78, 0x22, 0xc5, 0x0f, 0x18, 0x7d, 0xfa, 0xd3, 0x96, 0xc2, 0xdf,
	0x07, 0xfb, 0xf5, 0x17, 0x7e, 0xf4, 0x49, 0x36, 0xed, 0x87, 0xfa, 0xf5, 0xa7, 0x2d, 0x85, 0xbf,
	0x0f, 0xf7, 0x6b, 0x35, 0x43, 0x1d, 0xd3, 0x56, 0xbe, 0x3e, 0x62, 0xc0, 0x03, 0x13, 0xb6, 0x12,
	0x7e, 0x1e, 0x52, 0x16, 0x77, 0xa2, 0xa3, 0x9d, 0x16, 0x67, 0x49, 0x7b, 0xae, 0x65, 0x36, 0x48,
	0x69, 0xcd, 0x73, 0x19, 0x4f, 0x6e, 0xfa, 0x47, 0x94, 0x5d, 0x66, 0xd0, 0x76, 0xb3, 0x13, 0xbe,
	0x1e, 0xed, 0x99, 0x93, 0xb2, 0xa9, 0xb6, 0x9d, 0x7a, 0x95, 0x13, 0x7d, 0x5a, 0xff, 0x58, 0xcf,
	0x9c, 0xb2, 0x56, 0xc2, 0xcf, 0xc7, 0x0d, 0x78, 0xa0, 0x77, 0x5a, 0xa4, 0x8a, 0xf7, 0x98, 0x32,
	0xbb, 0x0d, 0x1d, 0xec, 0xc3, 0x4c, 0x78, 0x7a, 0xdc, 0x80, 0x47, 0x79, 0x64, 0x92, 0x18, 0xe5,
	0x9f, 0x36, 0xe0, 0x51, 0x1e, 0x81, 0x8a, 0xfa, 0x8c, 0x01, 0x9f, 0x7a, 0xb4, 0xdc, 0x05, 0x93,
	0xd7, 0x9b, 0xe2, 0xbd, 0xfe, 0x59, 0x03, 0x9e, 0xe7, 0x11, 0xa9, 0xb1, 0xcf, 0x19, 0xf0, 0xc5,
	0x24, 0xfc, 0x50, 0x14, 0xb1, 0xd3, 0xd4, 0x6c, 0xa8, 0x0a, 0x7c, 0xde, 0x80, 0xef, 0x50, 0x19,
	0x5c, 0x64, 0xfe, 0x05, 0xa5, 0x9c, 0x39, 0x2d, 0xeb, 0x50, 0x6b, 0x6b, 0x67, 0x89, 0xfe, 0xa9,
	0xe3, 0x8b, 0x06, 0x7c, 0x60, 0x49, 0xd3, 0x42, 0xf7, 0x4b, 0x3d, 0x7b, 0x64, 0x9e, 0xae, 0x90,
	0x45, 0xb2, 0xcc, 0x88, 0xdf, 0xac, 0x72, 0x93, 0xe9, 0x6e, 0x7c, 0xd2, 0x80, 0x8f, 0x16, 0xb0,
	0x95, 0xf0, 0xf3, 0x65, 0xa3, 0xd7, 0xab, 0x24, 0x65, 0x11, 0xb7, 0xe2, 0x57, 0x94, 0x1b, 0xf0,
	0x4d, 0x97, 0x31, 0x12, 0x5e, 0xbe, 0xda, 0x6f, 0x36, 0xa9, 0x46, 0xfc, 0x5a, 0xbf, 0xd9, 0xe8,
	0x3e, 0xfc, 0xba, 0x01, 0x7f, 0x0a, 0x28, 0x65, 0x6e, 0xdc, 0xd7, 0x0c, 0xf8, 0x7e, 0x50, 0x4a,
	0xde, 0xb7, 0x5f, 0x31, 0xf4, 0x67, 0x96, 0x2d, 0x19, 0x48, 0x9e, 0x26, 0x5e, 0xed, 0xd2, 0x27,
	0x25, 0xd7, 0x17, 0x07, 0xe9, 0xe4, 0xbb, 0xf3, 0xd7, 0x06, 0x7c, 0xff, 0x49, 0xa0, 0x22, 0x81,
	0xd7, 0x0c, 0xf8, 0xfe, 0x53, 0x4a, 0x7c, 0x58, 0x78, 0xbd, 0xcb, 0xee, 0x98, 0xa2, 0x8e, 0xe9,
	0xd4, 0x93, 0x07, 0xa7, 0x1f, 0x0c, 0xc2, 0xbb, 0x43, 0x92, 0x0a, 0xfb, 0xe1, 0x20, 0x7c, 0x73,
	0x89, 0x05, 0xe3, 0xa2, 0xfc, 0x68, 0x10, 0xbe, 0xb9, 0x48, 0x36, 0x06, 0x7f, 0x3c, 0x08, 0xdf,
	0xae, 0x24, 0x28, 0x2b, 0xf8, 0x74, 0x6f, 0xb9, 0xf8, 0x76, 0xf5, 0x93, 0x41, 0xf8, 0xaa, 0xa1,
	0x40, 0x79, 0x18, 0x2f, 0xfb, 0x0d, 0xfc, 0xcc, 0x20, 0x7c, 0xd5, 0x90, 0x68, 0x85, 0x59, 0x11,
	0xf7, 0x6c, 0x6f, 0xdf, 0xd1, 0x8f, 0xb4, 0x02, 0xfc, 0x69, 0x6f, 0x41, 0xbd, 0x30, 0x3f, 0x93,
	0x31, 0x4e, 0x9c, 0x46, 0xd7, 0xaf, 0x52, 0x46, 0x2e, 0x52, 0x67, 0x78, 0xef, 0x78, 0xf4, 0x4b,
	0xff, 0xb8, 0xfa, 0xa5, 0x7f, 0xbc, 0xe4, 0x04, 0xad, 0xf0, 0xe7, 0x12, 0xf9, 0x95, 0x60, 0xe4,
	0xb9, 0x87, 0x06, 0x46, 0x73, 0x63, 0x43, 0x8b, 0xd7, 0x09, 0x9b, 0x39, 0x67, 0xe2, 0x5e, 0x34,
	0x14, 0x5a, 0xbb, 0x01, 0xef, 0xc7, 0xfc, 0x79, 0x69, 0x1e, 0xba, 0xac, 0x04, 0x7c, 0x62, 0x16,
	0x6d, 0x0a, 0xed, 0x2d, 0x31, 0xad, 0xfa, 0x8c, 0xe1, 0x05, 0x29, 0xb2, 0x41, 0x58, 0x86, 0x63,
	0x6e, 0xce, 0x99, 0x98, 0x43, 0x9b, 0x13, 0x42, 0x7d, 0x86, 0xf3, 0xa2, 0x54, 0xda, 0xa8, 0x95,
	0x44, 0x4c, 0x67, 0xd0, 0x0d, 0xa1, 0x14, 0xa7, 0x4e, 0xbb, 0x1f, 0x95, 0x97, 0xa4, 0x4a, 0x58,
	0x89, 0x1a, 0x75, 0xda, 0x13, 0xf3, 0xe8, 0xc6, 0x50, 0x61, 0xc9, 0x75, 0xb9, 0xed, 0x9a, 0x16,
	0x61, 0xfd, 0xe8, 0xbc, 0x2c, 0x75, 0xc2, 0x44, 0xa6, 0xb4, 0xe9, 0x44, 0x11, 0x85, 0x99, 0x5e,
	0x74, 0xdc, 0x8b, 0xcb, 0x7e, 0xab, 0x1f, 0xa5, 0x6b, 0x52, 0x29, 0xcc, 0x63, 0xc1, 0x9d, 0xf1,
	0x5b, 0x53, 0x77, 0xa0, 0xfd, 0x75, 0xb7, 0x35, 0xee, 0x9b, 0xdc, 0xf5, 0x9b, 0xd4, 0x36, 0x97,
	0x7c, 0xf5, 0xff, 0x79, 0xd8, 0x74, 0x49, 0x4b, 0x4d, 0x6d, 0xaa, 0x85, 0x7f, 0x94, 0x9d, 0xf3,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x69, 0x67, 0x5d, 0x1f, 0x22, 0x00, 0x00,
}
