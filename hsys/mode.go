package hsys

import "strings"

const SysRunModeCfg = "sys.run.mode"

// Mode Run Mode Defined
type Mode uint8

const (
	LOCAL   Mode = 1
	DEV     Mode = 2
	TEST    Mode = 3
	PRE     Mode = 4
	PROD    Mode = 5
	UNKNOWN Mode = 0
)

const (
	localName   = "LOCAL"
	devName     = "DEV"
	testName    = "TEST"
	preName     = "PRE"
	prodName    = "PROD"
	unknownName = "UNKNOWN"
)

func ModeOf(modeVal string) Mode {
	mode := strings.ToUpper(modeVal)
	switch mode {
	case localName:
		return LOCAL
	case devName:
		return DEV
	case testName:
		return TEST
	case preName:
		return PRE
	case prodName:
		return PROD
	}
	return UNKNOWN
}

func (m Mode) String() string {
	switch m {
	case LOCAL:
		return localName
	case DEV:
		return devName
	case TEST:
		return testName
	case PRE:
		return preName
	case PROD:
		return prodName
	}
	return unknownName
}

func (m Mode) IsRd() bool {
	return m.IsLocal() ||
		m.IsDev() ||
		m.IsTest()
}

func (m Mode) IsLocal() bool {
	return m == LOCAL
}

func (m Mode) IsDev() bool {
	return m == DEV
}

func (m Mode) IsTest() bool {
	return m == TEST
}

func (m Mode) IsPre() bool {
	return m == PRE
}

func (m Mode) IsProd() bool {
	return m == PROD
}
