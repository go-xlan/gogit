package gogitv5x

import (
	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func NewFmtActiveFilesOptions(root string) *GetActiveFilesOptions {
	return NewGetActiveFilesOptions(root).SetFileExtension(".go").SetRunOnEachPath(func(path string) error {
		zaplog.ZAPS.P1.LOG.Info("go-fmt-file", zap.String("path", path))
		if err := formatgo.FormatFile(path); err != nil {
			return erero.Wro(err)
		}
		return nil
	})
}
