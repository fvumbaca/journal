package fs

import (
	"fmt"
	"path/filepath"
	"time"
)

// func WithFSOptions(ctx context.Context, baseFS string) context.Context {
// 	return context.WithValue(ctx, "fs_base", baseFS)
// }

func DayFilename(base string, t time.Time) string {
	return filepath.Join(base, fmt.Sprintf("%d/%s/%02d", t.Year(), t.Month().String(), t.Day())) + ".md"
}
