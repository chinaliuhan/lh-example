package lh_common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

/**
谷歌验证码的生成和验证
*/
type googleAuth struct {
}

func NewGoogleAuth() *googleAuth {
	return &googleAuth{}
}

func (r *googleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (r *googleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (r *googleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (r *googleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (r *googleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (r *googleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (r *googleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := r.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := r.toUint32(hashParts)
	return number % 1000000
}

// 获取秘钥
func (r *googleAuth) GetSecret() string {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.BigEndian, r.un())
	return strings.ToUpper(r.base32encode(r.hmacSha1(buf.Bytes(), nil)))
}

// 获取动态码
func (r *googleAuth) GetCodeBySecret(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := r.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := r.oneTimePassword(secretKey, r.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

// 获取动态码二维码内容
func (r *googleAuth) GetQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// 验证动态码
func (r *googleAuth) VerifyCode(secret, code string) (bool, error) {
	getCode, err := r.GetCodeBySecret(secret)
	if err != nil {
		return false, err
	}
	return getCode == code, nil
}
