package rank

import (
	"math"

	"github.com/hearchco/hearchco/src/config"
	"github.com/hearchco/hearchco/src/search/result"
)

func GetScore(result result.Result, rconf config.Ranking) float64 {
	retRankScore := float64(0)
	for _, er := range result.EngineRanks {
		seMul := rconf.Engines[er.SearchEngine.ToLower()].Mul
		seConst := rconf.Engines[er.SearchEngine.ToLower()].Const //these 2 could be preproced into array
		retRankScore += (100.0/math.Pow(float64(er.Rank)*rconf.A+rconf.B, rconf.REXP)*rconf.C+rconf.D)*seMul + seConst
	}
	retRankScore /= float64(result.TimesReturned)

	timesReturnedScore := math.Log(float64(result.TimesReturned)*rconf.TRA+rconf.TRB)*10*rconf.TRC + rconf.TRD

	score := retRankScore + timesReturnedScore
	return score
}