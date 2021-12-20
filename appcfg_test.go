package qaecommon

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadAppCfg(t *testing.T) {
	filename := filepath.Join(os.Getenv("QAE_APPROOT"), "testdata", "app.yaml")
	cfg, err := LoadAppCfg(filename)
	assert.Nil(t, err)

	assert.Equal(t, cfg.AppName, "kubia")
	assert.Equal(t, cfg.Config[0].Local, "config/store-production.yaml")
	assert.Equal(t, cfg.Config[0].Remote, "config/store.yaml")
	assert.Equal(t, cfg.Tasks[0].Type, "web")
	assert.Equal(t, cfg.Tasks[0].Name, "default")
	assert.Equal(t, cfg.Tasks[0].Url, "/")
}

func Test_LoadErrAppCfg(t *testing.T) {
	filename := filepath.Join(os.Getenv("QAE_APPROOT"), "testdata", "errapp.yaml")
	_, err := LoadAppCfg(filename)
	assert.NotNil(t, err)
}
