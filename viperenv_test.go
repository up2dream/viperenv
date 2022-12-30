package viperenv

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	assert.Equal(t, Config.Get("app.db.mysql.user"), "root-dev")
	assert.Equal(t, Config.Get("app.db.mysql.password"), "pass")
	assert.Equal(t, Config.Get("app.inc1"), "yes")
	assert.Equal(t, Config.Get("app.inc2"), "yes")
}
