package gorsa

import (
	"encoding/base64"
	"fmt"
)

// PublicEncrypt 公钥加密
func PublicEncrypt(data []byte) ([]byte, error) {
	if stdRsa.pubkey == nil {
		return nil, fmt.Errorf("pubkey is empty;Please run [InitPubPriKey]")
	}
	rsadata, err := stdRsa.PubKeyENCTYPT(data)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(rsadata)))
	base64.StdEncoding.Encode(buf, rsadata)
	return buf, nil
}

// PriKeyEncrypt 私钥加密
func PriKeyEncrypt(data []byte) ([]byte, error) {
	if stdRsa.prikey == nil {
		return nil, fmt.Errorf("prikey is empty;Please run [InitPubPriKey]")
	}
	rsadata, err := stdRsa.PriKeyENCTYPT(data)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(rsadata)))
	base64.StdEncoding.Encode(buf, rsadata)
	return buf, nil
}

// PublicDecrypt 公钥解密
func PublicDecrypt(data []byte) ([]byte, error) {
	if stdRsa.pubkey == nil {
		return nil, fmt.Errorf("pubkey is empty;Please run [InitPubPriKey]")
	}
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	if err != nil {
		return nil, err
	}
	return stdRsa.PubKeyDECRYPT(dbuf[:n])
}

// PriKeyDecrypt 私钥解密
func PriKeyDecrypt(data []byte) ([]byte, error) {
	if stdRsa.prikey == nil {
		return nil, fmt.Errorf("prikey is empty;Please run [InitPubPriKey]")
	}
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	if err != nil {
		return nil, err
	}
	return stdRsa.PriKeyDECRYPT(dbuf[:n])
}
