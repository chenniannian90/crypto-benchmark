package crypto_benchmark

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
	"testing"
)

func signECDSA(k *ecdsa.PrivateKey, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}

	//s, err = utils.ToLowS(&k.PublicKey, s)
	//if err != nil {
	//	return nil, err
	//}

	return utils.MarshalECDSASignature(r, s)
}

func verifyECDSA(k *ecdsa.PublicKey, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	r, s, err := utils.UnmarshalECDSASignature(signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	//lowS, err := utils.IsLowS(k, s)
	//if err != nil {
	//	return false, err
	//}
	//
	//if !lowS {
	//	return false, fmt.Errorf("Invalid S. Must be smaller than half the order [%s][%s].", s, utils.GetCurveHalfOrdersAt(k.Curve))
	//}

	return ecdsa.Verify(k, digest, r, s), nil
}


// fabric 默认的签名/验签 性能测试
func BenchmarkSign_FABRIC(b *testing.B) {
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	for i := 0; i < b.N; i++ {
		 _, _ = signECDSA(priv, hashed, nil)
	}
}

func BenchmarkVerify_FABRIC(b *testing.B) {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	origin := []byte("testing")

	hash := sha256.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)

	sig,_ := priv.Sign(rand.Reader,hashed, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = verifyECDSA(&priv.PublicKey, sig, hashed, nil)
	}
}

// go 标准库没有椭圆曲线加/解密
func TestEncAndDec_FABRIC(t *testing.T) {

}

func BenchmarkHash_Fabric(b *testing.B){
	origin := []byte(TestHashString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := sha256.New()
		hash.Write(origin)
		_ = hash.Sum(nil)
	}
}