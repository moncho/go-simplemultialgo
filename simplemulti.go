package simplemultialgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

//Algorithm holds information about algorithm profitability
type Algorithm struct {
	Paying string `json:"paying"`
	Port   uint   `json:"port"`
	Name   string `json:"name"`
	Index  int    `json:"algo"`
}

type response struct {
	R result `json:"result"`
}
type result struct {
	Algos []Algorithm `json:"simplemultialgo"`
}

const niceHashURL = "https://api.nicehash.com/api?method=simplemultialgo.info"

//NiceHashMultiAlgo queries profitability of all algorithms on NiceHash and
//returns the most profitable.
//The given map is used to filter and weight algorithms. It can be empty.
//See https://www.nicehash.com/software-developers
func NiceHashMultiAlgo(algoSpeeds map[string]int) (*Algorithm, error) {

	resp, err := http.Get(niceHashURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r response
	err = json.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	return mostProfitable(r.R.Algos, algoSpeeds), nil
}

func mostProfitable(algos []Algorithm, algoSpeeds map[string]int) *Algorithm {

	if len(algos) == 0 {
		return nil
	}
	mostProf := 0.0
	var result Algorithm

	if len(algoSpeeds) == 0 {
		sort.SliceStable(algos, sortByPaying(algos))
		return &algos[0]
	}
	for _, algo := range algos {
		if speed, ok := algoSpeeds[algo.Name]; ok {
			pay, err := strconv.ParseFloat(algo.Paying, 64)
			if err != nil {
				continue
			}
			weightedPaying := pay * float64(speed)
			if weightedPaying > mostProf {
				mostProf = weightedPaying
				result = algo
			}
		}
	}

	return &result
}

func sortByPaying(algos []Algorithm) func(i, j int) bool {
	return func(i, j int) bool {
		return algos[i].Paying > algos[j].Paying
	}
}
