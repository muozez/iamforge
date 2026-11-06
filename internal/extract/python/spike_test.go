package python

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/muozez/iamforge/internal/extract"
)

func findNodeByType(n extract.Node, targetType string) bool {
	if n == nil {
		return false
	}
	if n.Type() == targetType {
		return true
	}
	for i := 0; i < n.ChildCount(); i++ {
		if findNodeByType(n.Child(i), targetType) {
			return true
		}
	}
	return false
}

func TestPureGoPythonParserSpike(t *testing.T) {
	tests := []struct {
		name               string
		filename           string
		expectedSyntaxType string
	}{
		{
			name:               "f-strings support",
			filename:           "spike_fstring.py",
			expectedSyntaxType: "interpolation", // or formatted_string
		},
		{
			name:               "decorator support",
			filename:           "spike_decorator.py",
			expectedSyntaxType: "decorated_definition",
		},
		{
			name:               "type hint support",
			filename:           "spike_typehint.py",
			expectedSyntaxType: "type",
		},
		{
			name:               "async function support",
			filename:           "spike_async.py",
			expectedSyntaxType: "function_definition",
		},
		{
			name:               "pattern match statement support (Python 3.10+)",
			filename:           "spike_match.py",
			expectedSyntaxType: "match_statement",
		},
	}

	p := NewParser()
	ctx := context.Background()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filePath := filepath.Join("testdata", tc.filename)
			src, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("failed to read test file %s: %v", tc.filename, err)
			}

			tree, err := p.Parse(ctx, src)
			if err != nil {
				t.Fatalf("failed to parse %s: %v", tc.filename, err)
			}
			defer tree.Release()

			root := tree.RootNode()
			if root == nil {
				t.Fatalf("expected non-nil root node for %s", tc.filename)
			}

			if root.HasError() {
				t.Errorf("syntax error detected in %s: %s", tc.filename, root.Text())
			}

			if tc.expectedSyntaxType != "" {
				if !findNodeByType(root, tc.expectedSyntaxType) {
					t.Errorf("expected node of type %q in %s, but none was found", tc.expectedSyntaxType, tc.filename)
				}
			}
		})
	}
}

