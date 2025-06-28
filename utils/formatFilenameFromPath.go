package utils

import "path/filepath"

func FormatFilenameFromPath(from string) string {
	// to := from
	// to = strings.Replace(to, "\\", "___", -1)
	// to = strings.Replace(to, "/", "___", -1)
	// return to
	return filepath.Base(from)
}
