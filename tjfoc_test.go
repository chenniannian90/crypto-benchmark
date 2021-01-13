package crypto_benchmark

import (
	"bytes"
	"crypto/rand"
	"github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	"github.com/Hyperledger-TWGC/tjfoc-gm/sm3"
	"testing"
)

// 同济gm库 性能测试
func BenchmarkSign_TJ(b *testing.B) {
	hashed := []byte("testing")
	priv, _ := sm2.GenerateKey(rand.Reader)
	for i := 0; i < b.N; i++ {
		_, _ = priv.Sign(rand.Reader, hashed, nil)
	}
}

func BenchmarkVerify_TJ(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	origin := []byte("testing")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)

	sig,_ := priv.Sign(rand.Reader,hashed, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(&priv.PublicKey).Verify(hashed,sig)
	}
}

func TestEncAndDec_TJ(t *testing.T) {
	msg := []byte("sm2 encryption standard")

	sk, _ := sm2.GenerateKey(rand.Reader)
	pk := &sk.PublicKey

	//test encryption
	cipher, err := sm2.Encrypt(pk, msg, rand.Reader)
	if err != nil {
		t.Errorf("enc err:%s", err)
		return
	}

	//test decryption
	plain, err := sm2.Decrypt(sk, cipher)
	if err != nil {
		t.Errorf("dec err:%s", err)
		return
	}

	if !bytes.Equal(msg, plain) {
		t.Error("sm2 encryption is invalid")
		return
	}
}