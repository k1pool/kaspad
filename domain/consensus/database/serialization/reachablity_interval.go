package serialization

import (
	"github.com/k1pool/kaspad/domain/consensus/model"
)

func reachablityIntervalToDBReachablityInterval(reachabilityInterval *model.ReachabilityInterval) *DbReachabilityInterval {
	return &DbReachabilityInterval{
		Start: reachabilityInterval.Start,
		End:   reachabilityInterval.End,
	}
}

func dbReachablityIntervalToReachablityInterval(dbReachabilityInterval *DbReachabilityInterval) *model.ReachabilityInterval {
	return &model.ReachabilityInterval{
		Start: dbReachabilityInterval.Start,
		End:   dbReachabilityInterval.End,
	}
}
