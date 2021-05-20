package webAnalyzer

type WebPageContent struct {
	HtmlVersion          string
	PageTitle            string
	NumH1                int
	NumH2                int
	NumH3                int
	NumH4                int
	NumAnchors           int
	NumLinks             int
	NumInaccessibleLinks int
	LoginForm            bool
}
