package grok

import "regexp"

// Compile exposes the denormalised, compiled regular expression for a grok
// pattern so callers can run offset-aware matching (FindStringSubmatchIndex)
// and inspect capture-group names (SubexpNames) — neither of which the
// value-only Match/Parse surface provides.
//
// This is an additive accessor for the BreakpointVantage fork: it reuses the
// unexported compile() path (and its compiled-pattern cache) verbatim, so the
// returned *regexp.Regexp is byte-for-byte the one Match/Parse already use.
// It changes no matching behaviour.
func (g *Grok) Compile(pattern string) (*regexp.Regexp, error) {
	gr, err := g.compile(pattern)
	if err != nil {
		return nil, err
	}
	return gr.regexp, nil
}
