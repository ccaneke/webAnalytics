package webAnalyzer

type Want struct {
	NumAnchors      int
	NumH1           int
	NumH2           int
	NumH3           int
	NumH4           int
	PageTitle       string
	NumInaccessible int
	HTMLVersion     string
}

var testCases = []struct {
	in   string
	want PageInformation
}{
	{"<!Doctype html public><html><head></head><body></body>,</html>",
		PageInformation{"", "", 0, 0, 0, 0, 1, 0, 0, false}},
	{`<!Doctype html><html><head><title>Page Title</title></head><body></body></html>`,
		PageInformation{"", "Page Title", 0, 0, 0, 0, 0, 0, 0, false}},
	{"<!Doctype html><head><metadata></metadata></head><body><h1></h1></body></html>",
		PageInformation{"", "", 1, 0, 0, 0, 0, 0, 0, false}},
	{"<!Doctype html><head><metadata></metadata></head><body><h2></h2></body><html>",
		PageInformation{"", "", 0, 1, 0, 0, 0, 0, 0, false}},
	{"<!Doctype html><head></head><body><h3></h3><body></html>",
		PageInformation{"", "", 0, 0, 1, 0, 0, 0, 0, false}},
	{"<!Doctype html><html><head></head><body><h4></h4></body></html>",
		PageInformation{"", "", 0, 0, 0, 1, 0, 0, 0, false}},
	{`<!Doctype html><html><head></head><body><a href="example.com"></a></body></html>`,
		PageInformation{"", "", 0, 0, 0, 0, 1, 0, 0, false}},
	{`<!Doctype html><html><head><link></link></head><body></body></html>`,
		PageInformation{"", "", 0, 0, 0, 0, 0, 1, 0, false}},
	{`<!Doctype html><head></head><body><a href="https://example"></a></body></html>`,
		PageInformation{"", "", 0, 0, 0, 0, 0, 0, 1, false}},
	{`<!Doctype html><head></head><body><form>
	<input type="password" data-val="true" data-val-required="The Password field is required." 
	id="Input_Password"/></form></body></html>`,
		PageInformation{"", "", 0, 0, 0, 0, 0, 0, 1, true}},
}
