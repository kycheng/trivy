package cargo

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/go-dep-parser/pkg/rust/cargo"
	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/aquasecurity/trivy/pkg/fanal/analyzer"
	"github.com/aquasecurity/trivy/pkg/fanal/analyzer/language"
	"github.com/aquasecurity/trivy/pkg/fanal/types"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	analyzer.RegisterAnalyzer(&cargoLibraryAnalyzer{})
}

const version = 1

type cargoLibraryAnalyzer struct{}

func (a cargoLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(types.Cargo, input.FilePath, input.Content, cargo.NewParser())
	if err != nil {
		return nil, xerrors.Errorf("error with Cargo.lock: %w", err)
	}
	return res, nil
}

func (a cargoLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return fileName == types.CargoLock
}

func (a cargoLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeCargo
}

func (a cargoLibraryAnalyzer) Version() int {
	return version
}
