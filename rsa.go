package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

//（1）加密：采用sha1算法加密后转base64格式
func RsaEncryptWithSha1Base64(originalData, publicKey string) (string, error) {
	key, _ := base64.StdEncoding.DecodeString(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(key)

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

//（2）解密：对采用sha1算法加密后转base64格式的数据进行解密（私钥PKCS1格式）
func RsaDecryptWithSha1Base64(encryptedData, privateKey string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	prvKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return "", err
	}

	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey.(*rsa.PrivateKey), encryptedDecodeBytes)
	return string(originalData), err
}
