// Ex13 demonstrates const declarations for KB, MB, ..., YB.
package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("1 KiB = %d, 1 KB = %d, difference = %d%%\n", KiB, KB,
		100*(KiB-KB)/KB)
	fmt.Printf("1 MiB = %d, 1 MB = %d, difference = %d%%\n", MiB, MB,
		100*(MiB-MB)/MB)
	fmt.Printf("1 GiB = %d, 1 GB = %d, difference = %d%%\n", GiB, GB,
		100*(GiB-GB)/GB)
	fmt.Printf("1 TiB = %d, 1 TB = %d, difference = %d%%\n", TiB, TB,
		100*(TiB-TB)/TB)
	fmt.Printf("1 PiB = %d, 1 PB = %d, difference = %d%%\n", PiB, PB,
		100*(PiB-PB)/PB)
	fmt.Printf("1 EiB = %d, 1 EB = %d, difference = %d%%\n", EiB, EB,
		100*(EiB-EB)/EB)
	fmt.Printf("ZiB/ZB difference = %d%%\n",
		100*(ZiB-ZB)/ZB)
	fmt.Printf("YiB/YB difference = %d%%\n",
		100*(YiB-YB)/YB)
}
