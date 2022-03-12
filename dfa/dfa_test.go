package dfa

import "testing"

type test struct {
    currentState   string
    inputSymbol    string
    wanted         string
}

func TestDelta(t *testing.T) {
    tests := []test{
        {currentState: "0", inputSymbol: "a", wanted: "1"},
        {currentState: "0", inputSymbol: "b", wanted: "0"},
    }
    // Nilsson, S. (2018). 5 basic for loop patterns. Yourbasic.org; yourbasic.org. https://yourbasic.org/golang/for-loop/    
    for _, tc := range tests {
     
        got := Delta(tc.currentState, tc.inputSymbol) 
        if got != tc.wanted {
            t.Errorf("Error, got %q, wanted %q.", got, tc.wanted)
        }

    }
}
