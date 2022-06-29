// Package util go实用库
// RSA的签名及验签
// SHA256 With RSA
package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RSA RSA签名
func RSA(src, privateKey string) (s string, err error) {
	//1.加载RSA的私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = errors.New("private_key error")
		return
	}
	//pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	rsaPri, _ := pri.(*rsa.PrivateKey)
	hash := crypto.SHA256
	//2.计算代签名字串的SHA256哈希值
	hashed, err := SHA256([]byte(src))
	if err != nil {
		return
	}
	//3. base64 encode
	sig, _ := rsa.SignPKCS1v15(rand.Reader, rsaPri, hash, hashed)
	s = base64.StdEncoding.EncodeToString(sig)
	return
}

// RSACheck RSA验签
func RSACheck(src, sign, publicKey string) (ok bool, err error) {
	//1. 加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		err = errors.New("private_key error")
		return
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)

	//2. 计算代签名字串的SHA1哈希
	digest, err := SHA256([]byte(src))
	if err != nil {
		return
	}
	//3. base64 decode
	data, err := base64.StdEncoding.DecodeString(string(sign))
	if err != nil {
		return
	}
	//4. 调用rsa包的VerifyPKCS1v15验证签名有效性
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, digest, data)
	if err != nil {
		err = fmt.Errorf("Verify sig error, reason: %v", err)
		return
	}

	return true, nil
}
