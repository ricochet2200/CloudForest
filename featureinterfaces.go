package CloudForest

import ()

const maxExhaustiveCats = 5
const maxNonRandomExahustive = 10
const maxNonBigCats = 30
const minImp = 0.0
const constant_cutoff = 1e-7

//Feature contains all methods needed for a predictor feature.
type Feature interface {
	NCats() (n int)
	Length() (l int)
	GetStr(i int) (value string)
	IsMissing(i int) bool
	MissingVals() bool
	GoesLeft(i int, splitter *Splitter) bool
	PutMissing(i int)
	PutStr(i int, v string)
	SplitImpurity(l *[]int, r *[]int, m *[]int, allocs *BestSplitAllocs) (impurityDecrease float64)
	UpdateSImpFromAllocs(l *[]int, r *[]int, m *[]int, allocs *BestSplitAllocs, movedRtoL *[]int) (impurityDecrease float64)
	Impurity(cases *[]int, counter *[]int) (impurity float64)
	FindPredicted(cases []int) (pred string)
	BestSplit(target Target,
		cases *[]int,
		parentImp float64,
		leafSize int,
		randomSplit bool,
		allocs *BestSplitAllocs) (codedSplit interface{}, impurityDecrease float64, constant bool)
	DecodeSplit(codedSplit interface{}) (s *Splitter)
	ShuffledCopy() (fake Feature)
	Copy() (copy Feature)
	CopyInTo(copy Feature)
	Shuffle()
	ShuffleCases(cases *[]int)
	ImputeMissing()
	GetName() string
	Append(v string)
	Split(codedSplit interface{}, cases []int) (l []int, r []int, m []int)
	SplitPoints(codedSplit interface{}, cases *[]int) (lastl int, firstr int)
}

//NumFeature contains the methods of Feature plus methods needed to implement
//diffrent types of regression. It is usually embeded by regression targets to
//provide access to the underlying data.
type NumFeature interface {
	Feature
	Span(cases *[]int) float64
	Get(i int) float64
	Put(i int, v float64)
	Predicted(cases *[]int) float64
	Mean(cases *[]int) float64
	Norm(i int, v float64) float64
	Error(cases *[]int, predicted float64) (e float64)
	Less(i int, j int) bool
}

//CatFeature contains the methods of Feature plus methods needed to implement
//diffrent types of classification. It is usually embeded by classification targets to
//provide access to the underlying data.
type CatFeature interface {
	Feature
	CountPerCat(cases *[]int, counter *[]int)
	MoveCountsRtoL(allocs *BestSplitAllocs, movedRtoL *[]int)
	DistinctCats(cases *[]int, counter *[]int) int
	CatToNum(value string) (numericv int)
	NumToCat(i int) (value string)
	Geti(i int) int
	Puti(i int, v int)
	Modei(cases *[]int) int
	Mode(cases *[]int) string
	Gini(cases *[]int) float64
	GiniWithoutAlocate(cases *[]int, counts *[]int) (e float64)
	EncodeToNum() (fs []Feature)
}

//Target abstracts the methods needed for a feature to be predictable
//as either a catagroical or numerical feature in a random forest.
type Target interface {
	GetName() string
	NCats() (n int)
	SplitImpurity(l *[]int, r *[]int, m *[]int, allocs *BestSplitAllocs) (impurityDecrease float64)
	UpdateSImpFromAllocs(l *[]int, r *[]int, m *[]int, allocs *BestSplitAllocs, movedRtoL *[]int) (impurityDecrease float64)
	Impurity(cases *[]int, counter *[]int) (impurity float64)
	FindPredicted(cases []int) (pred string)
}

//BoostingTarget augments Target with a "Boost" method that will be called after each
//tree is grown with the partion generated by that tree. It will return the weigh the
//tree should be given and boost the target for the next tree.
type BoostingTarget interface {
	Target
	Boost(partition *[][]int) (weight float64)
}
