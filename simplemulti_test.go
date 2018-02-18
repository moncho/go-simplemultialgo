package simplemultialgo

import (
	"encoding/json"
	"reflect"
	"testing"
)

var testResult = `
{
	"result":{
	   "simplemultialgo":[
		  {
			 "paying":"0.00209085",
			 "port":3333,
			 "name":"scrypt",
			 "algo":0
		  },
		  {
			 "paying":"0.00000009",
			 "port":3334,
			 "name":"sha256",
			 "algo":1
		  },
		  {
			 "paying":"0",
			 "port":3335,
			 "name":"scryptnf",
			 "algo":2
		  },
		  {
			 "paying":"0.00002959",
			 "port":3336,
			 "name":"x11",
			 "algo":3
		  },
		  {
			 "paying":"0.00042363",
			 "port":3337,
			 "name":"x13",
			 "algo":4
		  },
		  {
			 "paying":"0",
			 "port":3338,
			 "name":"keccak",
			 "algo":5
		  },
		  {
			 "paying":"0.00040224",
			 "port":3339,
			 "name":"x15",
			 "algo":6
		  },
		  {
			 "paying":"0.00421801",
			 "port":3340,
			 "name":"nist5",
			 "algo":7
		  },
		  {
			 "paying":"0.21294808",
			 "port":3341,
			 "name":"neoscrypt",
			 "algo":8
		  },
		  {
			 "paying":"0",
			 "port":3342,
			 "name":"lyra2re",
			 "algo":9
		  },
		  {
			 "paying":"0",
			 "port":3343,
			 "name":"whirlpoolx",
			 "algo":10
		  },
		  {
			 "paying":"0.00038679",
			 "port":3344,
			 "name":"qubit",
			 "algo":11
		  },
		  {
			 "paying":"0.00034302",
			 "port":3345,
			 "name":"quark",
			 "algo":12
		  },
		  {
			 "paying":"0",
			 "port":3346,
			 "name":"axiom",
			 "algo":13
		  },
		  {
			 "paying":"0.0056219",
			 "port":3347,
			 "name":"lyra2rev2",
			 "algo":14
		  },
		  {
			 "paying":"0",
			 "port":3348,
			 "name":"scryptjanenf16",
			 "algo":15
		  },
		  {
			 "paying":"0.00001511",
			 "port":3349,
			 "name":"blake256r8",
			 "algo":16
		  },
		  {
			 "paying":"0",
			 "port":3350,
			 "name":"blake256r14",
			 "algo":17
		  },
		  {
			 "paying":"0",
			 "port":3351,
			 "name":"blake256r8vnl",
			 "algo":18
		  },
		  {
			 "paying":"146.92300974",
			 "port":3352,
			 "name":"hodl",
			 "algo":19
		  },
		  {
			 "paying":"0.00721243",
			 "port":3353,
			 "name":"daggerhashimoto",
			 "algo":20
		  },
		  {
			 "paying":"0.00001979",
			 "port":3354,
			 "name":"decred",
			 "algo":21
		  },
		  {
			 "paying":"194.78372349",
			 "port":3355,
			 "name":"cryptonight",
			 "algo":22
		  },
		  {
			 "paying":"0.0000684",
			 "port":3356,
			 "name":"lbry",
			 "algo":23
		  },
		  {
			 "paying":"515.19097229",
			 "port":3357,
			 "name":"equihash",
			 "algo":24
		  },
		  {
			 "paying":"0.00006442",
			 "port":3358,
			 "name":"pascal",
			 "algo":25
		  },
		  {
			 "paying":"0.00578291",
			 "port":3359,
			 "name":"x11gost",
			 "algo":26
		  },
		  {
			 "paying":"0.00000551",
			 "port":3360,
			 "name":"sia",
			 "algo":27
		  },
		  {
			 "paying":"0.00001042",
			 "port":3361,
			 "name":"blake2s",
			 "algo":28
		  },
		  {
			 "paying":"0.00663464",
			 "port":3362,
			 "name":"skunk",
			 "algo":29
		  }
	   ]
	}
 }
`

/*
func TestNiceHashMultiAlgo(t *testing.T) {
	type args struct {
		algos map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    *Algorithm
		wantErr bool
	}{{
		"",
		args{
			map[string]int{"scrypt": 1, "x11": 7, "quark": 6},
		},
		nil,
		false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NiceHashMultiAlgo(tt.args.algos)
			if (err != nil) != tt.wantErr {
				t.Errorf("NiceHashMultiAlgo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NiceHashMultiAlgo() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func Test_mostProfitable(t *testing.T) {
	var r response
	err := json.Unmarshal([]byte(testResult), &r)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		algos      []Algorithm
		algoSpeeds map[string]int
	}
	tests := []struct {
		name string
		args args
		want *Algorithm
	}{
		{
			"mostProfitable returns the most profitable algo taking into account the algo speeds",
			args{
				r.R.Algos,
				map[string]int{"scrypt": 1, "x11": 7, "quark": 10},
			},
			&Algorithm{Name: "quark",
				Paying: "0.00034302",
				Port:   3345,
				Index:  12,
			},
		},
		{
			"mostProfitable returns the most profitable algo",
			args{
				r.R.Algos,
				nil,
			},
			&Algorithm{Name: "equihash",
				Paying: "515.19097229",
				Port:   3357,
				Index:  24,
			},
		},
		{
			"empty algo list results in not finding a profitable algo",
			args{
				nil,
				map[string]int{"scrypt": 1, "x11": 7, "quark": 16},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostProfitable(tt.args.algos, tt.args.algoSpeeds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostProfitable() = %v, want %v", got, tt.want)
			}
		})
	}
}
