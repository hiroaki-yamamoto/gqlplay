package gqlplay

// Configuration structures

// CursorShape enumerate strings to specify the sahpe of the cursor.
type CursorShape string

const (
	// CursorShapeLine set the cursor shape to line.
	CursorShapeLine CursorShape = "line"
	// CursorShapeBlock set the cursor shape to block.
	CursorShapeBlock CursorShape = "block"
	// CursorShapeUnderline set the cursor shape to underline.
	CursorShapeUnderline CursorShape = "underline"
)

// Theme enumerates strings to specify the theme of the playground
type Theme string

const (
	// DarkTheme sets the theme to be dark
	DarkTheme Theme = "dark"
	// LightTheme sets the theme to be Light
	LightTheme Theme = "light"
)

// Credentials enumerates strings to specify credentials options.
type Credentials string

const (
	// OmitCredentials omits the credentials.
	OmitCredentials Credentials = "omit"
	// IncludeCredentials accepts the credentials from anywhere.
	IncludeCredentials Credentials = "include"
	// SameOriginCredentials accepts the credentials from same-origin.
	SameOriginCredentials Credentials = "same-origin"
)

// Settings is the setting structure to configure GraphQL Playground.
type Settings struct {
	CursorShape           CursorShape `json:"editor.cursorShape,omitempty"`
	FontFamily            string      `json:"editor.fontFamily,omitempty"`
	FontSize              float32     `json:"editor.fontSize,omitempty"`
	ReuseHeaders          bool        `json:"editor.reuseHeaders"`
	Theme                 Theme       `json:"editor.theme,omitempty"`
	BetaUpdates           bool        `json:"general.betaUpdates"`
	PrintWidth            float32     `json:"prettier.printWidth,omitempty"`
	TabWidth              float32     `json:"prettier.tabWidth,omitempty"`
	UseTabs               bool        `json:"prettier.useTabs"`
	Credentials           Credentials `json:"request.credentials,omitmepty"`
	PollingEnabled        bool        `json:"schema.polling.enable"`
	PollingEndpointFilter string      `json:"schema.polling.endpointFilter,omitempty"`
	PollingInterval       float32     `json:"schema.polling.interval,omitempty"`
	DisableSchemaComments bool        `json:"schema.disableComments"`
	HideTraceResponse     bool        `json:"tracing.hideTracingResponse"`
}

// Option represents the configuration of playground that depends on:
// https://github.com/prisma/graphql-playground
type Option struct {
	Settings
	Endpoint             string            `json:",omitempty"`
	SubscriptionEndpoint string            `json:",omitempty"`
	Headers              map[string]string `json:",omitempty"`
}
