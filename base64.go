package base64

import (
	"unsafe"
)

const (
	alphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	maxUint16 = ^uint16(0)
)

// bits:      [1][2][3][4][5][6][7][8][1][2][3][4][5][6][7][8][1][2][3][4][5][6][7][8]
// bytes:     [          B0          ][          B1          ][          B2          ]
// b64 chars: [       C0       ][       C1       ][       C2       ][       C3       ]
var (
	b0dt [maxUint16]byte
	b1dt [maxUint16]byte
	b2dt [maxUint16]byte

	b0dtp = unsafe.Pointer(&b0dt)
	b1dtp = unsafe.Pointer(&b1dt)
	b2dtp = unsafe.Pointer(&b2dt)
)

func init() {
	// Each table needs 52 cache lines for two base64 char combinations.
	for i0, c0 := range alphabet {
		for i1, c1 := range alphabet {
			idx := uint16(c0) | uint16(c1)<<8
			v := i0*64 + i1
			b0dt[idx] = byte((v >> 4) & 0xff)
			b1dt[idx] = byte((v >> 2) & 0xff)
			b2dt[idx] = byte(v & 0xff)
		}
	}
}

func Decode(dst []byte, src []byte) int {
	if len(src) == 0 {
		return 0
	}

	sp := unsafe.Pointer(unsafe.SliceData(src))
	dp := unsafe.Pointer(unsafe.SliceData(dst))

	spi := uintptr(0)
	dpi := uintptr(0)
	sl := uintptr(len(src))
	spl := uintptr(sl - 4)

	// main loop
	for spi < spl {
		*(*uint32)(unsafe.Pointer(uintptr(dp) + dpi)) =
			uint32(*(*byte)(unsafe.Pointer(uintptr(b0dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi)))))) |
				uint32(*(*byte)(unsafe.Pointer(uintptr(b1dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi + 1))))))<<8 |
				uint32(*(*byte)(unsafe.Pointer(uintptr(b2dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi + 2))))))<<16
		dpi += 3
		spi += 4
	}

	// last block
	if *(*byte)(unsafe.Pointer(uintptr(sp) + sl - 2)) == '=' {
		// one byte left
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi)) = *(*byte)(unsafe.Pointer(uintptr(b0dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi)))))
		dpi++
	} else if *(*byte)(unsafe.Pointer(uintptr(sp) + sl - 1)) == '=' {
		// two bytes left
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi)) = *(*byte)(unsafe.Pointer(uintptr(b0dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi)))))
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi + 1)) = *(*byte)(unsafe.Pointer(uintptr(b1dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi + 1)))))
		dpi += 2
	} else {
		// three bytes left
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi)) = *(*byte)(unsafe.Pointer(uintptr(b0dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi)))))
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi + 1)) = *(*byte)(unsafe.Pointer(uintptr(b1dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi + 1)))))
		*(*byte)(unsafe.Pointer(uintptr(dp) + dpi + 2)) = *(*byte)(unsafe.Pointer(uintptr(b2dtp) + uintptr(*(*uint16)(unsafe.Pointer(uintptr(sp) + spi + 2)))))
		dpi += 3
	}

	return int(dpi) + 1
}
