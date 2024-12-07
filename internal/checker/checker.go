package checker

import (
	"fmt"
	"math"
)

func CheckLoad(load float64) string {
	if load > 30 {
		return fmt.Sprintf("Load Average is too high: %.2f", load)
	}
	return ""
}

func CheckMemory(totalMem, usedMem float64) string {
	if totalMem <= 0 {
		return ""
	}
	usage := usedMem / totalMem * 100
	if usage > 80 {
		return fmt.Sprintf("Memory usage too high: %d%%", int(math.Floor(usage)))
	}
	return ""
}

func CheckDisk(totalDisk, usedDisk float64) string {
	if totalDisk <= 0 {
		return ""
	}
	usage := usedDisk / totalDisk
	if usage > 0.9 {
		freeBytes := totalDisk - usedDisk
		freeMb := int(math.Floor(freeBytes / (1024 * 1024)))
		return fmt.Sprintf("Free disk space is too low: %d Mb left", freeMb)
	}
	return ""
}

func CheckNetwork(totalNet, usedNet float64) string {
	if totalNet <= 0 {
		return ""
	}
	usage := usedNet / totalNet
	if usage > 0.9 {
		freeBytes := totalNet - usedNet
		freeMbit := int(math.Floor((freeBytes * 8) / (1024 * 1024)))
		return fmt.Sprintf("Network bandwidth usage high: %d Mbit/s available", freeMbit)
	}
	return ""
}