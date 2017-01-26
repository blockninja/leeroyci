package notification

import (
	"os"
	"testing"

	"github.com/nii236/leeroyci/database"
)

func TestMain(m *testing.M) {
	database.NewDatabase("sqlite3", ":memory:")

	i := m.Run()

	os.Exit(i)
}
