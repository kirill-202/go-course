package oak

import (
	"bytes"

	"oak/token"
	"reflect"
	"testing"
)


type scanTest struct {
	name string
	input string
	want []token.Token
}

func (st scanTest) run(t *testing.T) {
	b := bytes.NewBufferString(st.input)
	c := ScanConfig{}
	s := NewScanner(c, st.name, b)

	var got []token.Token

	for tok := s.Next(); tok.Type != token.EOF; tok = s.Next() {
		got = append(got, tok)
	}

	if !reflect.DeepEqual(st.want, got) {
		t.Errorf("line %q, wanted %v, got %v", st.input, st.want, got)

	}

	var scanTests = []scanTest{
		{
			name: "simple-add",
			input: "2 -1 + # comment",
			want: []token.Token{
				{Type: token.Number, Line:1, Text: "2"},
				{Type: token.Number,Line: 1, Text: "-1"},
				{Type: token.Operator, Line: 1, Text: "+"},
				{Type: token.Comment, Line: 1},
			},
		},
		
	}
}

func TestScanner(t *testing.T) {
	for _, st := range scanTest {
		t.Run(st.name, st.run)
	}
}