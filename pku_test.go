package crypto_benchmark

import (
	"bytes"
	"github.com/Hyperledger-TWGC/pku-gm/gmssl"
	"testing"
)

// 北大gm库 性能测试
func BenchmarkSign_PKU(b *testing.B) {
	hashed := []byte(TestHashString)
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	priv, _ := gmssl.GeneratePrivateKey("EC", sm2keygenargs, nil)
	for i := 0; i < b.N; i++ {
		 _, _ = priv.Sign("priv", hashed, nil)
	}
}

func BenchmarkVerify_PKU(b *testing.B) {
	hashed := []byte(TestHashString)
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	priv, _ := gmssl.GeneratePrivateKey("EC", sm2keygenargs, nil)
	sm2pkpem, _ := priv.GetPublicKeyPEM()
	pk, _ := gmssl.NewPublicKeyFromPEM(sm2pkpem)

	sm3hash := gmssl.New()
	sm3hash.Write(hashed)
	digest := sm3hash.Sum(nil)
	sig,_ := priv.Sign("sm2sign", digest, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := pk.Verify("sm2sign", digest,sig, nil)
		if err != nil{
			b.Errorf("failed to call pk.Verify, err:%v", err)
		}

	}
}

func TestEncAndDec_PKU(t *testing.T) {
	msg := []byte(TestHashString)

	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	priv, _ := gmssl.GeneratePrivateKey("EC", sm2keygenargs, nil)
	sm2pkpem, _ := priv.GetPublicKeyPEM()
	pk, _ := gmssl.NewPublicKeyFromPEM(sm2pkpem)

	//test encryption
	cipher, err :=  pk.Encrypt("sm2encrypt-with-sm3", msg, nil)
	if err != nil {
		t.Errorf("enc err:%s", err)
		return
	}

	//test decryption
	plain, err := priv.Decrypt("sm2encrypt-with-sm3", cipher, nil)
	if err != nil {
		t.Errorf("dec err:%s", err)
		return
	}

	if !bytes.Equal(msg, plain) {
		t.Error("sm2 encryption is invalid")
		return
	}
}

func BenchmarkHash_PKU(b *testing.B){
	origin := []byte(TestHashString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm3hash := gmssl.New()
		sm3hash.Write(origin)
		  _ = sm3hash.Sum(nil)
	}
}