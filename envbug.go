package godebug

import "os"

var envCache map[string]string

func init() {
	envCache = make(map[string]string)
}

func ChkEnv(envVar string) bool {
	v, ok := envCache[envVar]
	if ok {
		return ParseBool(v)
	}
	v = os.Getenv(envVar)
	envCache[envVar] = v
	x := ChkEnv(envVar)
	return x
}

var trueValues map[string]bool

func init() {

	trueValues = make(map[string]bool)
	trueValues["t"] = true
	trueValues["T"] = true
	trueValues["yes"] = true
	trueValues["Yes"] = true
	trueValues["YES"] = true
	trueValues["1"] = true
	trueValues["true"] = true
	trueValues["True"] = true
	trueValues["TRUE"] = true
	trueValues["on"] = true
	trueValues["On"] = true
	trueValues["ON"] = true
}

func ParseBool(s string) (b bool) {
	_, b = trueValues[s]
	return
	//if InArray(s, []string{"t", "T", "yes", "Yes", "YES", "1", "true", "True", "TRUE", "on", "On", "ON"}) {
	//	return true
	//}
	//return false
}
