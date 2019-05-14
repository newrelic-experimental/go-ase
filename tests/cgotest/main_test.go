package cgotest

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/SAP/go-ase/cgo"
	"github.com/SAP/go-ase/tests/libtest"
)

func TestMain(m *testing.M) {
	err := testMain(m)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

func testMain(m *testing.M) error {
	simpleDSN, simpleTeardown, err := libtest.DSN(false)
	if err != nil {
		return fmt.Errorf("Failed to setup simple DSN: %v", err)
	}
	defer simpleTeardown()

	err = libtest.RegisterDSN("username password", *simpleDSN, cgo.NewConnector)
	if err != nil {
		return fmt.Errorf("Failed to setup simple databases: %v", err)
	}

	userstoreDSN, userstoreTeardown, err := libtest.DSN(true)
	if err != nil {
		return fmt.Errorf("Failed to setup userstore DSN: %v", err)
	}
	defer userstoreTeardown()

	err = libtest.RegisterDSN("userstorekey", *userstoreDSN, cgo.NewConnector)
	if err != nil {
		return fmt.Errorf("Failed to setup userstorekey databases: %v", err)
	}

	rc := m.Run()
	if rc != 0 {
		return fmt.Errorf("Tests failed with %d", rc)
	}

	return nil
}

func TestInt64(t *testing.T) {
	libtest.DoTestInt64(t)
}

func TestUint64(t *testing.T) {
	libtest.DoTestUint64(t)
}

func TestFloat64(t *testing.T) {
	libtest.DoTestFloat64(t)
}

func TestBool(t *testing.T) {
	libtest.DoTestBool(t)
}

func TestBytes(t *testing.T) {
	libtest.DoTestBytes(t)
}

func TestString(t *testing.T) {
	libtest.DoTestString(t)
}

func TestTime(t *testing.T) {
	libtest.DoTestTime(t)
}

func TestSQLTx(t *testing.T) {
	libtest.DoTestSQLTx(t)
}

func TestSQLExec(t *testing.T) {
	libtest.DoTestSQLExec(t)
}
