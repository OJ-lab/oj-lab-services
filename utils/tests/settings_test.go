package tests

import (
	"github.com/OJ-lab/oj-lab-services/utils"
	"log"
	"testing"
)

func TestIniBasicUsage(t *testing.T) {
	databaseSettings, _ := utils.GetDatabaseSettings("../../config/example.ini")
	log.Print(utils.GetDatabaseDSN(databaseSettings))
}
