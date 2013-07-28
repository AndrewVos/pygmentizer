package pygmentizer

import (
  "testing"
  "strings"
)

func TestHighlightsRubyCode(t *testing.T) {
  highlighted:= strings.TrimSpace(Highlight("ruby", "def hello\n  puts 'hi'\nend"))
  expected:= strings.TrimSpace(`
<div class="highlight"><pre><span class="k">def</span> <span class="nf">hello</span>
  <span class="nb">puts</span> <span class="s1">&#39;hi&#39;</span>
<span class="k">end</span>
</pre></div>
`)
  if highlighted != expected {
    t.Errorf("\nExpected:\n%q\ngot:\n%q\n", expected, highlighted)
  }
}
