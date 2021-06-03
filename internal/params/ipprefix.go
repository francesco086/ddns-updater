package params

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var ErrParsePrefix = errors.New("cannot parse IP prefix")

func ipv6DecimalPrefixToMask(prefixDecimal string) (ipMask net.IPMask, err error) {
	if prefixDecimal == "" {
		return nil, fmt.Errorf("%w: empty prefix", ErrParsePrefix)
	}

	prefixDecimal = strings.TrimPrefix(prefixDecimal, "/")

	const bits = 8 * net.IPv6len

	ones, consumed, ok := decimalToInteger(prefixDecimal)
	if !ok || consumed != len(prefixDecimal) || ones < 0 || ones > bits {
		return nil, fmt.Errorf("%w: %s", ErrParsePrefix, prefixDecimal)
	}

	return net.CIDRMask(ones, bits), nil
}

func decimalToInteger(s string) (ones int, i int, ok bool) {
	const big = 0xFFFFFF // Bigger than we need, not too big to worry about overflow
	const ten = 10

	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		ones = ones*ten + int(s[i]-'0')
		if ones >= big {
			return big, i, false
		}
	}

	return ones, i, true
}
