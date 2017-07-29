package engine

import (
	"github.com/360EntSecGroup-Skylar/goreporter/linters/unittest"
)

type StrategyImportPackages struct {
	Sync *Synchronizer `inject:""`
}

func (s *StrategyImportPackages) GetName() string {
	return "ImportPackages"
}

func (s *StrategyImportPackages) GetDescription() string {
	return "Check the project variables, functions, etc. naming spelling is wrong."
}

func (s *StrategyImportPackages) GetWeight() float64 {
	return 0.
}

// linterImportPackages is a function that scan the project contains all the
// package lists.It will extract from the linter need to convert
// the data.The result will be saved in the r's attributes.
func (s *StrategyImportPackages) Compute(parameters StrategyParameter) (summaries Summaries) {
	summaries = NewSummaries()

	importPkgs := unittest.GoListWithImportPackages(parameters.ProjectPath)
	for i := 0; i < len(importPkgs); i++ {
		summaries[importPkgs[i]] = Summary{Name: importPkgs[i]}
	}
	return
}

func (s *StrategyImportPackages) Percentage(summaries Summaries) float64 {
	return CountPercentage(len(summaries))
}
