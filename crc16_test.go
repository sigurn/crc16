package crc16

import (
	"testing"
)

var testData = []byte("123456789")

func testSelectedCRC(params Params, t *testing.T) {
	table := MakeTable(params)
	if table == nil {
		t.Errorf("Failed to create %q computer", params.Name)
	}

	crc := Checksum(testData, table)

	if crc != table.params.Check {
		t.Errorf("Invalid %q sample calculation, expected: %X, actual: %X", table.params.Name, table.params.Check, crc)
	}
}

func TestCRC16_ARC(t *testing.T) {
	testSelectedCRC(CRC16_ARC, t)
}

func TestCRC16_AUG_CCIT(t *testing.T) {
	testSelectedCRC(CRC16_AUG_CCITT, t)
}

func TestCRC16_BUYPASS(t *testing.T) {
	testSelectedCRC(CRC16_BUYPASS, t)
}

func TestCRC16_CCITT_FALSE(t *testing.T) {
	testSelectedCRC(CRC16_CCITT_FALSE, t)
}

func TestCRC16_CDMA2000(t *testing.T) {
	testSelectedCRC(CRC16_CDMA2000, t)
}

func TestCRC16_DDS_110(t *testing.T) {
	testSelectedCRC(CRC16_DDS_110, t)
}

func TestCRC16_DECT_R(t *testing.T) {
	testSelectedCRC(CRC16_DECT_R, t)
}

func TestCRC16_DECT_X(t *testing.T) {
	testSelectedCRC(CRC16_DECT_X, t)
}

func TestCRC16_DNP(t *testing.T) {
	testSelectedCRC(CRC16_DNP, t)
}

func TestCRC16_EN_13757(t *testing.T) {
	testSelectedCRC(CRC16_EN_13757, t)
}

func TestCRC16_GENIBUS(t *testing.T) {
	testSelectedCRC(CRC16_GENIBUS, t)
}

func TestCRC16_MAXIM(t *testing.T) {
	testSelectedCRC(CRC16_MAXIM, t)
}

func TestCRC16_MCRF4XX(t *testing.T) {
	testSelectedCRC(CRC16_MCRF4XX, t)
}

func TestCRC16_RIELLO(t *testing.T) {
	testSelectedCRC(CRC16_RIELLO, t)
}

func TestCRC16_T10_DIF(t *testing.T) {
	testSelectedCRC(CRC16_T10_DIF, t)
}

func TestCRC16_TELEDISK(t *testing.T) {
	testSelectedCRC(CRC16_TELEDISK, t)
}

func TestCRC16_TMS37157(t *testing.T) {
	testSelectedCRC(CRC16_TMS37157, t)
}

func TestCRC16_USB(t *testing.T) {
	testSelectedCRC(CRC16_USB, t)
}

func TestCRC16_CRC_A(t *testing.T) {
	testSelectedCRC(CRC16_CRC_A, t)
}

func TestCRC16_KERMIT(t *testing.T) {
	testSelectedCRC(CRC16_KERMIT, t)
}

func TestCRC16_MODBUS(t *testing.T) {
	testSelectedCRC(CRC16_MODBUS, t)
}

func TestCRC16_X_25(t *testing.T) {
	testSelectedCRC(CRC16_X_25, t)
}

func TestCRC16_XMODEM(t *testing.T) {
	testSelectedCRC(CRC16_XMODEM, t)
}
