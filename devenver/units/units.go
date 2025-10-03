// taken from: https://github.com/docker/go-units/blob/master/size.go
package units

import (
	"fmt"
)

const (
	// Decimal
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB

	// Binary
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB
	PiB = 1024 * TiB
)

type unitMap map[byte]int64

var (
	decimapAbbrs = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
)

func getSizeAndUnit(size float64, base float64, _map []string) (float64, string) {
	i := 0
	unitsLimit := len(_map) - 1
	for size >= base && i < unitsLimit {
		size = size / base
		i++
	}
	return size, _map[i]
}

// HumanSizeWithPrecision allows the size to be in any precision,
// instead of 4 digit precision used in units.HumanSize.
func HumanSizeWithPrecision(size float64, precision int) string {
	size, unit := getSizeAndUnit(size, 1000.0, decimapAbbrs)
	return fmt.Sprintf("%.*g%s", precision, size, unit)
}

// HumanSize returns a human-readable approximation of a size
// capped at 4 valid numbers (eg. "2.746 MB", "796 KB").
func HumanSize(size float64) string {
	return HumanSizeWithPrecision(size, 4)
}
