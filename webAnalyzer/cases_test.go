package webAnalyzer

var testCases = []struct {
	in   string
	want PageInformation
}{
	{"<!Doctype html public><html><head></head><body></body>,</html>",
		PageInformation{HtmlVersion: "HTML4", LoginForm: false}},
	{`<!Doctype html><html><head><title>Page Title</title></head><body></body></html>`,
		PageInformation{HtmlVersion: "HTML5", PageTitle: "Page Title", LoginForm: false}},
	{"<!Doctype html><head><metadata></metadata></head><body><h1></h1></body></html>",
		PageInformation{HtmlVersion: "HTML5", PageTitle: "", NumH1: 1, LoginForm: false}},
	{"<!Doctype html><head><metadata></metadata></head><body><h2></h2></body><html>",
		PageInformation{HtmlVersion: "HTML5", PageTitle: "", NumH2: 1, LoginForm: false}},
	{"<!Doctype html><head></head><body><h3></h3><body></html>",
		PageInformation{HtmlVersion: "HTML5", NumH3: 1, LoginForm: false}},
	{"<!Doctype html><html><head></head><body><h4></h4></body></html>",
		PageInformation{HtmlVersion: "HTML5", NumH4: 1, LoginForm: false}},
	{"<!Doctype html><html><head></head><body><h5></h5></body></html>",
		PageInformation{HtmlVersion: "HTML5", NumH5: 1, LoginForm: false}},
	{"<!Doctype html><html><head></head><body><h6></h6></body></html>",
		PageInformation{HtmlVersion: "HTML5", NumH6: 1, LoginForm: false}},
	{`<!Doctype html><html><head></head><body><a href="http://example.com/"></a></body></html>`,
		PageInformation{HtmlVersion: "HTML5", NumAnchors: 1, LoginForm: false}},
	{`<!Doctype html><html><head><link></link></head><body></body></html>`,
		PageInformation{HtmlVersion: "HTML5", NumLinks: 1, LoginForm: false}},
	{`<!Doctype html><head></head><body><a href="https://example"></a></body></html>`,
		PageInformation{HtmlVersion: "HTML5", NumAnchors: 1, NumInaccessibleLinks: 1,
			LoginForm: false}},
	{`<!Doctype html><head></head><body><form>
	<input type="password" data-val="true" data-val-required="The Password field is required." 
	id="Input_Password"/></form></body></html>`,
		PageInformation{HtmlVersion: "HTML5", LoginForm: true}},
}
