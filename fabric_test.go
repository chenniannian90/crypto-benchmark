package crypto_benchmark

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

// fabric 默认的签名/验签 性能测试
func BenchmarkSign_FABRIC(b *testing.B) {
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	for i := 0; i < b.N; i++ {
		_, _, _ = ecdsa.Sign(rand.Reader, priv, hashed)
	}
}

func BenchmarkVerify_FABRIC(b *testing.B) {
	//priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//origin := []byte("testing")
	//hash := sm3.New()
	//hash.Write(origin)
	//hashed := hash.Sum(nil)
	//
	//sig,_ := priv.Sign(rand.Reader,hashed, nil)
	//b.ResetTimer()
	//for i := 0; i < b.N; i++ {
	//	(&priv.PublicKey)..Verify(hashed,sig)
	//}
}

func TestEncAndDec_FABRIC(t *testing.T) {
	//msg := []byte("sm2 encryption standard")
	//
	//sk, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//pk := &sk.PublicKey
	//
	////test encryption
	//cipher, err := ecdsa.Encrypt(pk, msg, rand.Reader)
	//if err != nil {
	//	t.Errorf("enc err:%s", err)
	//	return
	//}
	//
	////test decryption
	//plain, err := sm2.Decrypt(sk, cipher)
	//if err != nil {
	//	t.Errorf("dec err:%s", err)
	//	return
	//}
	//
	//if !bytes.Equal(msg, plain) {
	//	t.Error("sm2 encryption is invalid")
	//	return
	//}
}
