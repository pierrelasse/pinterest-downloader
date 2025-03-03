package utils

func Size_formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return Fmt("%d B", bytes)
	}
	div, exp := unit, 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return Fmt("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
