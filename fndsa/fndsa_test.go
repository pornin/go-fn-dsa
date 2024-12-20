package fndsa

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	sha3 "golang.org/x/crypto/sha3"
	"testing"
)

func TestFNDSA_Self(t *testing.T) {
	for logn := uint(2); logn <= uint(10); logn++ {
		fmt.Printf("[%d]", logn)
		for i := 0; i < 10; i++ {
			sk, vk, err := KeyGen(logn, nil)
			if err != nil {
				t.Fatal(err)
			}
			if len(sk) != SigningKeySize(logn) {
				t.Fatalf("wrong signing key size (logn=%d): %d\n",
					logn, len(sk))
			}
			if len(vk) != VerifyingKeySize(logn) {
				t.Fatalf("wrong verifying key size (logn=%d): %d\n",
					logn, len(vk))
			}
			data := []byte("test")
			var sig []byte
			if logn <= 8 {
				sig, err = SignWeak(nil, sk, DOMAIN_NONE, 0, data)
			} else {
				sig, err = Sign(nil, sk, DOMAIN_NONE, 0, data)
			}
			if err != nil {
				t.Fatal(err)
			}
			if len(sig) != SignatureSize(logn) {
				t.Fatalf("wrong signature size (logn=%d): %d\n",
					logn, len(sig))
			}
			var r bool
			if logn <= 8 {
				r = VerifyWeak(vk, DOMAIN_NONE, 0, data, sig)
			} else {
				r = Verify(vk, DOMAIN_NONE, 0, data, sig)
			}
			if !r {
				t.Fatalf("signature verification failed (logn=%d)\n", logn)
			}
			fmt.Print(".")
		}
	}
	fmt.Println()
}

// Test vectors:
// kat_n[] contains 10 vectors for n = 2^logn
// For test vector kat_n[j]:
//
//	Let seed1 = 0x00 || logn || j
//	Let seed2 = 0x01 || logn || j
//	(logn over one byte, j over 4 bytes, little-endian)
//	seed1 is used with FakeRng1 in keygen.
//	seed2 is used with FakeRng2 for signing.
//	A key pair (sk, vk) is generated. A message is signed:
//	    domain context: "domain" (6 bytes)
//	    message: "message" (7 bytes)
//	    if n is odd, message is pre-hashed with SHA3-256; raw otherwise
//	KAT_n[j] is SHA3-256(sk || vk || sig)
var kat_4 = []string{
	"feeb4bde204cb40cbe06c7e5834abdfcec199219197e603883dbe47028bbfbf2",
	"4f7d1867e9e02ee571a45b6d6d24b8f02b68b2e59441d1e341d06bbf36bf668e",
	"8bd38088f833b66d1a5a4319e48c0efd2b1578fd7fc3bb7d20e167f4cd52e8de",
	"24e37763e19942bb1acc6b5e5a4867170d07741fe055e8e3c2411f1b754bbd1b",
	"9679a55739e76b66a475fe94053606bf07b930d47cc05377444f19f2c85ef2e6",
	"3435ac75ffeb8c72df5e5d2c8619ef2a991de0fe9864014306a9af16630b41f3",
	"8913b2791a76a746242160a800737459dc6457d1420317d7b21043ae286c5798",
	"56413a0307b574b7bff2b6f9f9b59e346f6ab16c2c75fe1c64949a025dc40534",
	"570e6fe189c45ab50e039eaa0ac3c5f2f50efbffa08e006368d3364e4d49f7fd",
	"60b307e72b295b3fb13bd7c2f5926b521c34fbbd4d9ee3cdfe89eed9ffb2d2af",
}
var kat_8 = []string{
	"956766887db48fd1f9cac47a93a12c9e55de6e47006457eceee523d3566f3dec",
	"9f41d30fad1bee288928b1f78a376a46dc06a0edc869bdb6cce0acc36583e92f",
	"8389ba7095343bd222c9818da07ac7e66b73dfdeafb6cdc10377242874c27ece",
	"7fd7ba114d952c9afe2c1dd4ee30e644b2e6caed13aed4e7e969260962a25c58",
	"5a65e67783352ade4a5cfc7d0a48849fecbdbefffdcd8d25d425c3f013f9f019",
	"f3044d077d30621ac7735fb3f95c35a58e15a3aa1c391467b6c33e05d8240c28",
	"cf012db9b469ada96be790b8050b68d531fbdd2f4940d0ac07b8ffc02310f8e3",
	"1e1c251797a4b27f4849ab34dfb9b21b3a84a52c4b0c11b93cf07305da26134f",
	"6e051f873582f6d94c93b335f059588acb00722a40e09b310a0c00894fdf05af",
	"d025acba6daf2b1de7d82d423b6eecb946e98cd7f7125f150e302ac8fccc3af2",
}
var kat_16 = []string{
	"7e2561ddd8664383b2e03bcb4da2409d4c43676ed021dee59766e72890a4509b",
	"7f284169006a71440cc27cace9cfeab56440d357ee42b47609e1b76513281b21",
	"46c05f015b609826c310a098a2105a0e94ad271313031b307a5ff6af09b14de2",
	"ed689cbfd26b8d3f4785d2622df343ef6ef11bf7d883d41f570416a632213fe1",
	"3b8c717aa4b2c5ea95b8df2af003e97d982e20230058ccaaa3d465a3239b05ca",
	"71e64b14011712731f7e02dee789d8c76cbc0d5f16c983b044067b30d47971d1",
	"3bc7443e28014cd78cb31eb7e5283aa9e23827d21b1317a8fe4fbb031cedcac5",
	"4ffe1c59cfe27ecbf233710bdc535a4a332c68e741a0a9a1b684d773cfc031f3",
	"49adb0cb6ed7af916adb4f213016d862a88ab284f9a61fc11e12a1828540b1b4",
	"27d2d2558117e4861207851dcc5f51322fb5e21cad7ace06390f5132f4c0ec17",
}
var kat_32 = []string{
	"97517f9cfe9641fbb06b08afa09be14096b13573960f6790ba1119eb01a8f723",
	"66c8669fe31f434582a465705dafea2a09c4acaf5c2c9d5975b4ec72d556c80b",
	"7207b9f036d9b7a40f5d3647f03fb4ebc373719f240791cd65f9f35fc471ef35",
	"bb9f9ebe61c5db1d72ebfaf2d699cc4c70e4c899f896b4f331fae004cd7a9b59",
	"90484e94c5bb5c6c2f5c48bfeec4ce15b4935d09bc55b1fdaa6ad71e3e03e194",
	"ea8822e989b8bf3484eaac010d77275d7d953cd0d16a51dbde9dc43ccf4bed0a",
	"afb52381c81b8b5fa7b48bbd8262e450bc69161e6c31112678a3743b5efbf58b",
	"96be92ccc265fa68564593baa4fe4f3cbac2f4a0c85c81f80ca28b2f3a3c099b",
	"05af0cb90b923f778c7f88b0e6747861da0a0f73481fe2b1587b16417ed7101f",
	"fc3201c8a5763e6b9919c54044aa7c302dc11344ab629917ef14680d3dce82fa",
}
var kat_64 = []string{
	"dc6efcd8382f2ec32a5d0048ccfecd7d0aa2804ed31f9ca7b3b7fe80a1f278d6",
	"8b96fe42791a4ddd3f426ea35d278830d0d688a2259355e568e63a88afe8093a",
	"6fd98e52e33c89a20dda23f4f25744350fd69f3fec640c06590866b004f3799d",
	"b0696877b0de7a9b82b74038b4be03d8a4669de8aa39845c36bc969ec8cdd4a5",
	"e0047e262bfd3df4874587d3966d12191835d27a84935d4f28ee6551c4d56db9",
	"3b61bc4d990adf23afbef5e0366d4d3328f776e74173792de0ebb1ac9d87412b",
	"358eeb3cfa720339970489378e1418cb618f927b47065e580f8c56b74f92f46b",
	"4384664a9f6ef03ddb96b77e09349ce951480ac0e0666e9f4236b213c69cfb2f",
	"8cd018e4d9add2fd5f12dc3015e9ff5ef6195154d4c09f4dfa8436681899db6d",
	"37e523a85668c4ea1ea59eb44e44bb1872d0ce8ec9571e329a1b2a9a60eacd05",
}
var kat_128 = []string{
	"22195e02f65e0906245eaedd12bedd89a89afcb68c62d27ded954a72fdfa1547",
	"dc2051d21719a1276c7a1f860e334c632ea0b1b15ff5203aac6fb93fe11ee123",
	"adc6b4de01547c5d6b382534fbc715fe7c434cd5c213f7bfd2d1d5056e7618a6",
	"191a5490b1a8fb166e3337ffd15b2b9d99dc31ebb07f69c8fa527e5e4878edf6",
	"60592b75cfb9bc459b99ada2e6b357b8b2a0796316a97efdf7d42d49ac8a20ac",
	"ff9660ef3e4f918ed588bc315e5f295421e0e8ff88d3c787d8c587396ab8e881",
	"2dd9b7c1632f64ad88da054db0488324d00f4ef550bddbd5961b963400f824b9",
	"c3d68100de903315d7a7ae47ce3d33ba9da7f9d6a27d563ccce997771a13974f",
	"b793a6fec199c60455ea22cf3b9cf0987a3c1157b4729f522498fdfe1e8f6043",
	"23f66127ac55cfab218a9a4b199fa42bc64056bb040ae653e90e63cc882eff60",
}
var kat_256 = []string{
	"0e19693ced586519efd7ff4cb45b8013d2f300b60eba2d291599d366bb03f1d9",
	"30c926ee6237d407cff189c2baeb3171872aebc461b919484cf30d93250fcde0",
	"3d4268db567841caa0e360e2d6c79c354b659f521509243381b494b4eec2b4af",
	"0c504032ffbdd2f2b26cb8d0c478fbf645e2fe3bdacd1fe25a5d15fd3830edc0",
	"3c9bcb09a3b6b54264068bf1df32051065f1099d4fa0b90ffb14e5391e7af564",
	"2c2efd441d9733dc3c14b1e62444856d9ffe12e4ef5104dd30e891c2c16237f0",
	"5b87a753c041dc60f938b0971e066d6feb6055f1a021db3036ca64741280a116",
	"f38f57b7cda123e36a03ebb7c0bb196a86dda4abd66a038cd054f7a4bd61e50a",
	"755a29fb4dcf7808399f501fde4c0e23d11b9face58c9f6681f1c636b2256989",
	"af091e60104821510b28599068fa84fd814af62d978f6830e7fa2fc51fedcf9b",
}
var kat_512 = []string{
	"a32f07baf6b7ff6bc7c3c4f8c638871ff8c4803b0e54bedb9363f5672011077b",
	"1794cfad199c20879d1ffe10ce263334095e51f0ed191ed74e4cba635e233d80",
	"d16188abb5502eae81e6e03750123e156d8ed7dfa830a0c879560b383a5dc53a",
	"14c03d690bf39bed73ac024a2b94adc1ff276d0c11e35d3455b9ea13c361b96c",
	"c3bdbab8e434c5264c1d6fb523777d5bab1e23a1c292066e3cb731742230b042",
	"d315c931fde38bdaeb83e6378d322f33ec9a36915ea5ed05e84ec3debddafa55",
	"3587e5d75e2f0de5e2116c3a136d1a559e58ffd4a10328060ce9a430e47bd87c",
	"b622711852cdc9893aec144ed635d2ae775778c6f4152e106b7b6b2842c8055d",
	"32b3c2ed31f11795dde312b0574164dc4d00712f4736d1c5142a49cab4261ed4",
	"e2bd2350de9bdab72d3a517251217d8fdbd7ea6e386ad2ff1da19c7c2111bcb2",
}
var kat_1024 = []string{
	"16ef63f9dc51b66565bb05ac525f3668fa48186b973a95599e0c963cfd6a4297",
	"f62ac74368b2f8b80b6e12f13e026c9ba493c59b9eb2225a2626dc773e257dba",
	"3f4de163f9a44137c52b0d9d6042a236fb8a05f9bd6617e12fbbd32bb0f2120c",
	"77d567ae787dae191cdcf406f5e6a88e16b6a3729b814ac49f7d182b6cd624d8",
	"d19a28ba50359df8d119fa4557116d45dffec6f422ae9aa563186270a6a36ee0",
	"cbda1bcc23e33ff63864cbb44db9e618c76214a91e8a4f57ea1170b468181728",
	"ebf8388ba558660ffc67ac6d14709b7ffd096603ba23660c761b603767b469d7",
	"233f0de0b9f70c2b7de870fc2f3d0b0d1fa37224a3264525d2d8537862c353d8",
	"9fdf2626bcb2e5a8622dd1fcc78ce78db3a2aceeff030def85574259ae41e555",
	"979346e3d31abf04f815ffd1d7bd44da03c636172b46ab260e365c4a4672445e",
}

func TestFNDSA_KAT(t *testing.T) {
	testFNDSA_KAT_inner(t, 2, kat_4)
	testFNDSA_KAT_inner(t, 3, kat_8)
	testFNDSA_KAT_inner(t, 4, kat_16)
	testFNDSA_KAT_inner(t, 5, kat_32)
	testFNDSA_KAT_inner(t, 6, kat_64)
	testFNDSA_KAT_inner(t, 7, kat_128)
	testFNDSA_KAT_inner(t, 8, kat_256)
	testFNDSA_KAT_inner(t, 9, kat_512)
	testFNDSA_KAT_inner(t, 10, kat_1024)
	fmt.Println()
}

func testFNDSA_KAT_inner(t *testing.T, logn uint, kat []string) {
	fmt.Printf("[%d]", logn)
	for j := 0; j < len(kat); j++ {
		var seed [6]byte
		seed[0] = 0x00
		seed[1] = byte(logn)
		seed[2] = byte(j)
		seed[3] = byte(j >> 8)
		seed[4] = byte(j >> 16)
		seed[5] = byte(j >> 24)

		var seed_kgen [32]byte
		sh := sha3.NewShake256()
		sh.Write(seed[:])
		sh.Read(seed_kgen[:])
		skey, vkey, err := KeyGen(logn, bytes.NewReader(seed_kgen[:]))
		if err != nil {
			t.Fatal(err)
		}

		seed[0] = 0x01
		ctx := DomainContext([]byte("domain"))
		var id crypto.Hash
		msg := []byte("message")
		if (j & 1) == 0 {
			id = 0
		} else {
			id = crypto.SHA3_256
			hv := sha3.Sum256(msg)
			msg = hv[:]
		}
		sig, err := sign_inner_seeded(logn, logn, seed[:], skey, ctx, id, msg)
		if err != nil {
			t.Fatal(err)
		}
		var r bool
		if logn <= 8 {
			r = VerifyWeak(vkey, ctx, id, msg, sig)
		} else {
			r = Verify(vkey, ctx, id, msg, sig)
		}
		if !r {
			t.Fatalf("signature verification failed (logn=%d, j=%d)", logn, j)
		}

		sc := sha3.New256()
		sc.Write(skey)
		sc.Write(vkey)
		sc.Write(sig)
		tmp := sc.Sum(nil)
		ref, _ := hex.DecodeString(kat[j])
		if !bytes.Equal(tmp, ref) {
			t.Fatalf("KAT failed (logn=%d, j=%d): wrong hash\n", logn, j)
		}

		fmt.Print(".")
	}
}

func BenchmarkKeyGen512(b *testing.B) {
	bench_keygen_inner(b, 9)
}

func BenchmarkKeyGen1024(b *testing.B) {
	bench_keygen_inner(b, 10)
}

func bench_keygen_inner(b *testing.B, logn uint) {
	for i := 0; i < b.N; i++ {
		KeyGen(logn, nil)
	}
}

func BenchmarkSign512(b *testing.B) {
	bench_sign_inner(b, 9)
}

func BenchmarkSign1024(b *testing.B) {
	bench_sign_inner(b, 10)
}

func bench_sign_inner(b *testing.B, logn uint) {
	// Make a key pair.
	sk, vk, _ := KeyGen(logn, nil)

	// Data is a raw message, not pre-hashed, and context is empty.
	data := []byte("test")

	// A few blank signatures for "warm-up".
	for i := 0; i < 10; i++ {
		sig, err := Sign(nil, sk, DOMAIN_NONE, 0, data)
		if err != nil {
			b.Fatalf("failure, err = %v", err)
		}
		if !Verify(vk, DOMAIN_NONE, 0, data, sig) {
			b.Fatalf("ERR: signature verification failed")
		}
		data = sig[len(sig)-32:]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sig, _ := Sign(nil, sk, DOMAIN_NONE, 0, data)
		data = sig[len(sig)-32:]
	}
}

func BenchmarkVerify512(b *testing.B) {
	bench_verify_inner(b, 9)
}

func BenchmarkVerify1024(b *testing.B) {
	bench_verify_inner(b, 10)
}

func bench_verify_inner(b *testing.B, logn uint) {
	// Make a key pair.
	sk, vk, _ := KeyGen(logn, nil)

	// Data is a raw message, not pre-hashed, and context is empty.
	data := []byte("test")

	// Compute some signatures.
	var sigs [10][]byte
	for i := 0; i < 10; i++ {
		sigs[i], _ = Sign(nil, sk, DOMAIN_NONE, 0, data)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !Verify(vk, DOMAIN_NONE, 0, data, sigs[i%len(sigs)]) {
			b.Fatal("signature verification failed")
		}
	}
}
