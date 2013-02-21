// Defines the Prolog standard library. All prelude predicates are
// implemented in pure Prolog.  Each variable in this package is a predicate
// definition.  At init time, they're combined into a single string
// in the Prelude var.
package prelude

import "strings"

// After init(), Prelude contains all prelude predicates combined into
// a single large string.  One rarely addresses this variable directly
// because golog.NewMachine() does it for you.
var Prelude string
func init() {
    Prelude = strings.Join([]string{
        Phrase3,
    }, "\n\n")
}

// phrase(:DCGBody, ?List, ?Rest) is nondet.
//
// True when DCGBody applies to the difference List/Rest.
var Phrase3 = `
phrase(Dcg, Head, Tail) :-
    call(Dcg, Head, Tail).
`
