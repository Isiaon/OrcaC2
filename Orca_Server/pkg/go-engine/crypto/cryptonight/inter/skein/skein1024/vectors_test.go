// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package skein1024

import (
	"bytes"
	"encoding/hex"
	"Orca_Server/pkg/go-engine/crypto/cryptonight/inter/skein"
	"testing"
)

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

var testVectors = []struct {
	hashsize  int
	conf      *skein.Config
	msg, hash string
}{
	{
		hashsize: 128,
		conf:     nil,
		msg:      "",
		hash: "0FFF9563BB3279289227AC77D319B6FFF8D7E9F09DA1247B72A0A265CD6D2A62" +
			"645AD547ED8193DB48CFF847C06494A03F55666D3B47EB4C20456C9373C86297" +
			"D630D5578EBD34CB40991578F9F52B18003EFA35D3DA6553FF35DB91B81AB890" +
			"BEC1B189B7F52CB2A783EBB7D823D725B0B4A71F6824E88F68F982EEFC6D19C6",
	},
	{
		hashsize: 20,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "2E6A4CBF2EF05EA9C24B93E8D1DE732DDF2739EB",
	},
	{
		hashsize: 28,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "1D6DE19F37F7A3C265440EECB4B9FBD3300BB5AC60895CFC0D4D3C72",
	},
	{
		hashsize: 32,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "986A4D472B123E8148731A8EAC9DB23325F0058C4CCBC44A5BB6FE3A8DB672D7",
	},
	{
		hashsize: 48,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "9C3D0648C11F31C18395D5E6C8EBD73F43D189843FC45235E2C35E345E12D62B" +
			"C21A41F65896DDC6A04969654C2E2CE9",
	},
	{
		hashsize: 128,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "96CA81F586C825D0360AEF5ACAEC49AD55289E1797072EEE198B64F349CE65B6" +
			"E6ED804FE38F05135FE769CC56240DDDA5098F620865CE4A4278C77FA2EC6BC3" +
			"1C0F354CA78C7CA81665BFCC5DC54258C3B8310ED421D9157F36C093814D9B25" +
			"103D83E0DDD89C52D0050E13A64C6140E6388431961685734B1F138FE2243086",
	},
	{
		hashsize: 257,
		conf:     nil,
		msg: "FBD17C26B61A82E12E125F0D459B96C91AB4837DFF22B39B78439430CDFC5DC8" +
			"78BB393A1A5F79BEF30995A85A12923339BA8AB7D8FC6DC5FEC6F4ED22C122BB" +
			"E7EB61981892966DE5CEF576F71FC7A80D14DAB2D0C03940B95B9FB3A727C66A" +
			"6E1FF0DC311B9AA21A3054484802154C1826C2A27A0914152AEB76F1168D4410",
		hash: "56A0CAB1AD315859DA7A6CFC35807CBFE039AF06CA4B8671C053360BDA0B17C1" +
			"4A9EB5EB2ABB01B0DB3F45C03CD30C69D7C1B70C5C9EF74C06FB3AEF0C843CE9" +
			"B4C1BA2294DDB5C71CAB692CEDC1E6F908C471B38C0C583418B55AEFDDFE08AB" +
			"A4055D0D19EDB5CCBA16C3E288471EFE463E6BF6CC346CC74F6C013E0293E6DF" +
			"D2E4AA66A92242FD395B6D91AAAD5A071C449D77EA00E44ECC75073890AC50D4" +
			"F4210E8C9DA45385A46D214A0FCCC131DB3F842F955E6E76AC311B3BF439DD51" +
			"9BEDD691785ADF7540F3163AD1216CF2ADB7D4BF40D93BE3184AEF51B651CA26" +
			"C7EC44073F43AD689D269EA9FF02F8D2C8932FE6CED0292F97FB5F07CA276D6B" +
			"43",
	},
	{
		hashsize: 128,
		conf: &skein.Config{Key: fromHex("CB41F1706CDE09651203C2D0EFBADDF847A0D315CB2E53FF8BAC41DA0002672E" +
			"920244C66E02D5F0DAD3E94C42BB65F0D14157DECF4105EF5609D5B0984457C1" +
			"935DF3061FF06E9F204192BA11E5BB2CAC0430C1C370CB3D113FEA5EC1021EB8" +
			"75E5946D7A96AC69A1626C6206B7252736F24253C9EE9B85EB852DFC81463134")},
		msg: "",
		hash: "BCF37B3459C88959D6B6B58B2BFE142CEF60C6F4EC56B0702480D7893A2B0595" +
			"AA354E87102A788B61996B9CBC1EADE7DAFBF6581135572C09666D844C90F066" +
			"B800FC4F5FD1737644894EF7D588AFC5C38F5D920BDBD3B738AEA3A3267D161E" +
			"D65284D1F57DA73B68817E17E381CA169115152B869C66B812BB9A84275303F0",
	},
	{
		hashsize: 128,
		conf: &skein.Config{Key: fromHex("CB41F1706CDE09651203C2D0EFBADDF847A0D315CB2E53FF8BAC41DA0002672E" +
			"920244C66E02D5F0DAD3E94C42BB65F0D14157DECF4105EF5609D5B0984457C1" +
			"935DF3061FF06E9F204192BA11E5BB2CAC0430C1C370CB3D113FEA5EC1021EB8" +
			"75E5946D7A96AC69A1626C6206B7252736F24253C9EE9B85EB852DFC81463134" +
			"6C")},
		msg: "D3090C72",
		hash: "DF0596E5808835A3E304AA27923DB05F61DAC57C0696A1D19ABF188E70AA9DBC" +
			"C659E9510F7C9A37FBC025BD4E5EA293E78ED7838DD0B08864E8AD40DDB3A880" +
			"31EBEFC21572A89960D1916107A7DA7AC0C067E34EC46A86A29CA63FA250BD39" +
			"8EB32EC1ED0F8AC8329F26DA018B029E41E2E58D1DFC44DE81615E6C987ED9C9",
	},
	{
		hashsize: 128,
		conf:     &skein.Config{Key: fromHex("")},
		msg:      "D3",
		hash: "F1FBB54F260D0FB9D49A29EEC184B265EDC663668A9720AA61661E43659B3CD6" +
			"97C700CE1E3E535E0C69801220B5DA975138E7CB1EC8D8E3018F078A32CAE28B" +
			"C189350B68EE67785623B372EF7811BB06BA6C67E5847596FB72F2B51994EB8E" +
			"E079B960E228F7026E1BFE8CEA0877496F986FD13DB82E132CC45F70BB010F27",
	},
}

func TestVectors(t *testing.T) {
	for i, v := range testVectors {
		conf, msg, ref := v.conf, fromHex(v.msg), fromHex(v.hash)

		h := New(v.hashsize, conf)

		h.Write(msg)
		sum := h.Sum(nil)
		if !bytes.Equal(sum, ref) {
			t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(sum), hex.EncodeToString(ref))
		}

		sum = Sum(msg, v.hashsize, conf)
		if !bytes.Equal(sum, ref) {
			t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(sum), hex.EncodeToString(ref))
		}

		var key []byte
		if conf != nil {
			key = conf.Key
		}

		switch v.hashsize {
		case 64:
			{
				var out [64]byte
				Sum512(&out, msg, key)
				if !bytes.Equal(out[:], ref) {
					t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(out[:]), hex.EncodeToString(ref))
				}
			}
		case 48:
			{
				var out [48]byte
				Sum384(&out, msg, key)
				if !bytes.Equal(out[:], ref) {
					t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(out[:]), hex.EncodeToString(ref))
				}
			}
		case 32:
			{
				var out [32]byte
				Sum256(&out, msg, key)
				if !bytes.Equal(out[:], ref) {
					t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(out[:]), hex.EncodeToString(ref))
				}
			}
		case 20:
			{
				var out [20]byte
				Sum160(&out, msg, key)
				if !bytes.Equal(out[:], ref) {
					t.Fatalf("Test vector %d : Hash does not match:\nFound:      %s\nExpected: %s", i, hex.EncodeToString(out[:]), hex.EncodeToString(ref))
				}
			}
		}
	}
}
