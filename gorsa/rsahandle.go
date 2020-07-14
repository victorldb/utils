package gorsa

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

// stdRsa --
var stdRsa = &RSASecurity{}

// RSASecurity --
type RSASecurity struct {
	pubkey *rsa.PublicKey  //公钥
	prikey *rsa.PrivateKey //私钥
}

// InitPubPriKey --
func InitPubPriKey(pubBytes, priBytes []byte) (err error) {
	if len(pubBytes) == 0 && len(priBytes) == 0 {
		return fmt.Errorf("pubkey and prikey is empty")
	}
	if len(pubBytes) > 0 {
		stdRsa.pubkey, err = getPubKey(pubBytes)
		if err != nil {
			return err
		}
	}
	if len(priBytes) > 0 {
		stdRsa.prikey, err = getPriKey(priBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewRSAPKCSHandle --
func NewRSAPKCSHandle(pubBytes, priBytes []byte) (rsaHandle *RSASecurity, err error) {
	rsaHandle = &RSASecurity{}
	if len(pubBytes) == 0 && len(priBytes) == 0 {
		return nil, fmt.Errorf("pubkey and prikey is empty")
	}
	if len(pubBytes) > 0 {
		rsaHandle.pubkey, err = getPubKey(pubBytes)
		if err != nil {
			return nil, err
		}
	}
	if len(priBytes) > 0 {
		rsaHandle.prikey, err = getPriKey(priBytes)
		if err != nil {
			return nil, err
		}
	}
	return rsaHandle, nil
}

// PubKeyENCTYPT 公钥加密
func (c *RSASecurity) PubKeyENCTYPT(input []byte) ([]byte, error) {
	if c.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(c.pubkey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// PubKeyENCTYPTWithBase64 --
func (c *RSASecurity) PubKeyENCTYPTWithBase64(input []byte) ([]byte, error) {
	rsadata, err := c.PubKeyENCTYPT(input)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(rsadata)))
	base64.StdEncoding.Encode(buf, rsadata)
	return buf, nil
}

// PubKeyDECRYPT 公钥解密
func (c *RSASecurity) PubKeyDECRYPT(input []byte) ([]byte, error) {
	if c.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(c.pubkey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// PubKeyDECRYPTWithBase64 --
func (c *RSASecurity) PubKeyDECRYPTWithBase64(input []byte) ([]byte, error) {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	n, err := base64.StdEncoding.Decode(dbuf, input)
	if err != nil {
		return nil, err
	}
	return c.PubKeyDECRYPT(dbuf[:n])
}

// PriKeyENCTYPT 私钥加密
func (c *RSASecurity) PriKeyENCTYPT(input []byte) ([]byte, error) {
	if c.prikey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(c.prikey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// PriKeyENCTYPTWithBase64 --
func (c *RSASecurity) PriKeyENCTYPTWithBase64(input []byte) ([]byte, error) {
	rsadata, err := c.PriKeyENCTYPT(input)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(rsadata)))
	base64.StdEncoding.Encode(buf, rsadata)
	return buf, nil
}

// PriKeyDECRYPT 私钥解密
func (c *RSASecurity) PriKeyDECRYPT(input []byte) ([]byte, error) {
	if c.prikey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(c.prikey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}

	return ioutil.ReadAll(output)

}

// PriKeyDECRYPTWithBase64 --
func (c *RSASecurity) PriKeyDECRYPTWithBase64(input []byte) ([]byte, error) {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	n, err := base64.StdEncoding.Decode(dbuf, input)
	if err != nil {
		return nil, err
	}
	return c.PriKeyDECRYPT(dbuf[:n])

}
