package util

import (
	"fmt"
	"testing"
)

func TestRsaVerySignWithSha1Base64(t *testing.T) {

	pubkey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCF5fg1AK1aXstUNw8HAcWwk+qznxmP9RiR6WR5BsNOp8Ax16CZTY71bClJMIKoxZsWE8RqzkUOBMSZnSTXWR3k5aAQCkUvO6lRlhCoFBX1Uvh+Ruas3IvDcYrDr6Rr2PETAdNVws9cuL6QQiz5oat1vsDAkS2TsEK/SOpKQZ9AzwIDAQAB"

	privateKey := "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAIXl+DUArVpey1Q3DwcBxbCT6rOfGY/1GJHpZHkGw06nwDHXoJlNjvVsKUkwgqjFmxYTxGrORQ4ExJmdJNdZHeTloBAKRS87qVGWEKgUFfVS+H5G5qzci8NxisOvpGvY8RMB01XCz1y4vpBCLPmhq3W+wMCRLZOwQr9I6kpBn0DPAgMBAAECgYAermBefMY8I+aBaJBX8a8D1BD+kaFA3E+B97HfDY6pMkUW5qrVFfGBCQS40ZXO6GCuAV+LfbJTEQKflGOBuNjhaJxy2JjEzVn8vPZA3NlL4MDCj922TxSwSFJxMBieh9puGvwBifUGugvCUby0rxbrlhWrCmydOa+9ih/tm+VqUQJBAONQrOrEnbnx9P02hSWR4/vMlmpzGJs3SrzzgrWkLSVd5jY4f/kmsgec7C3FDMO9i06N9yaEHEuXwBjYq8xgOAcCQQCWy4BKUI5m6+IPxkd/G9hMm7seRECxUGQYUo+1hmn1pq2HopmLBGvE2M72Xjw1cKyomo1FpBo5b7yElONZbK75AkAjEexLoB/xcynn8wRhwntY+rxuGem+8K3gLvWIjpbEgBMnZFoiF106HXS2rwMEI/cdHHv3/kPbScNCUhNXSbT7AkBHhhzCovT4Ulf2XXjaDG4K4C7fy0XYFKZ1duudETU/BCD43aHwc1deleuMpePvAROUIUJyzsR1i88iH7C6YLeZAkBE5uqRL6qAkou9iHcv0DSLc4yOtbyBTTH4TbRl2BfcOa1WuWf2mq6d9hmDhGqt0mpoPcd2AAejb+TEZsIFTmBu"

	data, err := RsaEncryptWithSha1Base64("123456", pubkey)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("密文：%s \n", data)

	base64, err := RsaDecryptWithSha1Base64(data, privateKey)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("解密：%s \n", base64)

}
