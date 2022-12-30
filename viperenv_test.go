package viperenv

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config := ReadConfig()
	assert.Equal(t, config.Get("app.db.mysql.user"), "root-dev")
	assert.Equal(t, config.Get("app.db.mysql.password"), "pass")
	assert.Equal(t, config.Get("app.inc1"), "yes")
	assert.Equal(t, config.Get("app.inc2"), "yes")
}
