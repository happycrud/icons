// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package material

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Attractions(size string, classes ...string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{classes}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" enable-background=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("new 0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 122}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs("0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 161}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><g><rect fill=\"none\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 215}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 5, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></rect></g><g><path d=\"M10.43,18.75C10.8,18.29,11.37,18,12,18c0.63,0,1.19,0.29,1.56,0.75c0.39-0.09,0.76-0.21,1.12-0.36l-1.42-3.18 c-0.39,0.15-0.82,0.23-1.26,0.23c-0.46,0-0.9-0.09-1.3-0.25l-1.43,3.19C9.65,18.54,10.03,18.67,10.43,18.75z M5.15,10 c-0.16,0.59-0.25,1.21-0.25,1.85c0,0.75,0.12,1.47,0.33,2.15c0.63,0.05,1.22,0.4,1.56,0.99c0.33,0.57,0.35,1.23,0.11,1.79 c0.27,0.27,0.56,0.53,0.87,0.76l1.52-3.39v0c-0.47-0.58-0.75-1.32-0.75-2.13c0-1.89,1.55-3.41,3.46-3.41 c1.91,0,3.46,1.53,3.46,3.41c0,0.82-0.29,1.57-0.78,2.16l1.5,3.35c0.32-0.24,0.62-0.5,0.9-0.79c-0.22-0.55-0.2-1.2,0.12-1.75 c0.33-0.57,0.9-0.92,1.52-0.99c0.22-0.68,0.34-1.41,0.34-2.16c0-0.64-0.09-1.27-0.25-1.86c-0.64-0.04-1.26-0.39-1.6-1 c-0.36-0.62-0.35-1.36-0.03-1.95c-0.91-0.98-2.1-1.71-3.44-2.05C13.39,5.6,12.74,6,12,6c-0.74,0-1.39-0.41-1.74-1.01 C8.92,5.33,7.73,6.04,6.82,7.02C7.15,7.62,7.17,8.37,6.8,9C6.45,9.62,5.81,9.97,5.15,10z M3.85,9.58C3.07,8.98,2.83,7.88,3.34,7 c0.51-0.88,1.58-1.23,2.49-0.85c1.11-1.17,2.56-2.03,4.18-2.42C10.15,2.75,10.99,2,12,2c1.01,0,1.85,0.75,1.98,1.73 c1.63,0.39,3.07,1.24,4.18,2.42c0.91-0.38,1.99-0.03,2.49,0.85c0.51,0.88,0.27,1.98-0.51,2.58c0.23,0.77,0.35,1.58,0.35,2.42 s-0.12,1.65-0.35,2.42c0.78,0.6,1.02,1.7,0.51,2.58c-0.51,0.88-1.58,1.23-2.49,0.85c-0.4,0.43-0.85,0.81-1.34,1.15l1.34,3H16.3 l-0.97-2.17c-0.43,0.18-0.88,0.33-1.34,0.44C13.85,21.25,13.01,22,12,22c-1.01,0-1.85-0.75-1.98-1.73 C9.54,20.15,9.08,20,8.64,19.81L7.66,22H5.78l1.36-3.03c-0.47-0.33-0.91-0.71-1.3-1.12C4.92,18.23,3.85,17.88,3.34,17 c-0.51-0.88-0.27-1.98,0.51-2.58C3.62,13.65,3.5,12.84,3.5,12S3.62,10.35,3.85,9.58z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AttractionsOutlined(size string, classes ...string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var11 = []any{classes}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" enable-background=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs("new 0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 122}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs("0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 161}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><g><rect fill=\"none\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 215}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 9, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" y=\"0\"></rect></g><g><g><g><path d=\"M20.15,14.42c0.23-0.77,0.35-1.58,0.35-2.42s-0.12-1.65-0.35-2.42c0.78-0.6,1.02-1.7,0.51-2.58s-1.58-1.23-2.49-0.85 c-1.11-1.17-2.56-2.03-4.18-2.42C13.85,2.75,13.01,2,12,2s-1.85,0.75-1.98,1.73C8.39,4.12,6.95,4.98,5.83,6.15 C4.92,5.77,3.85,6.12,3.34,7S3.07,8.98,3.85,9.58C3.62,10.35,3.5,11.16,3.5,12s0.12,1.65,0.35,2.42c-0.78,0.6-1.02,1.7-0.51,2.58 s1.58,1.23,2.49,0.85c0.4,0.42,0.83,0.79,1.3,1.12L5.78,22h1.88l0.98-2.19c0.44,0.19,0.9,0.34,1.38,0.46 C10.15,21.25,10.99,22,12,22s1.85-0.75,1.98-1.73c0.46-0.11,0.91-0.26,1.34-0.44L16.3,22h1.88l-1.34-3 c0.48-0.34,0.93-0.72,1.34-1.15c0.91,0.38,1.99,0.03,2.49-0.85C21.17,16.12,20.93,15.02,20.15,14.42z M13.56,18.75 C13.19,18.29,12.63,18,12,18s-1.2,0.29-1.57,0.75c-0.4-0.09-0.79-0.21-1.16-0.37l1.43-3.19c0.4,0.16,0.84,0.25,1.3,0.25 c0.44,0,0.87-0.08,1.26-0.23l1.42,3.18C14.32,18.54,13.95,18.66,13.56,18.75z M10.48,12.02c0-0.83,0.67-1.5,1.5-1.5 s1.5,0.67,1.5,1.5s-0.67,1.5-1.5,1.5S10.48,12.85,10.48,12.02z M18.71,14.01c-0.61,0.07-1.18,0.41-1.52,0.99 c-0.32,0.56-0.34,1.2-0.12,1.75c-0.28,0.29-0.58,0.55-0.9,0.79l-1.5-3.35c0.49-0.59,0.78-1.34,0.78-2.16 c0-1.89-1.55-3.41-3.46-3.41s-3.46,1.53-3.46,3.41c0,0.8,0.28,1.54,0.75,2.13l0,0l-1.52,3.39c-0.31-0.23-0.6-0.48-0.87-0.76 C7.15,16.23,7.13,15.57,6.8,15c-0.34-0.59-0.93-0.94-1.56-0.99c-0.22-0.68-0.33-1.4-0.33-2.15c0-0.64,0.09-1.26,0.25-1.85 c0.66-0.03,1.3-0.38,1.65-1c0.37-0.63,0.35-1.38,0.01-1.98C7.74,6.05,8.93,5.34,10.27,5c0.34,0.59,0.99,1,1.73,1s1.39-0.4,1.73-1 c1.34,0.34,2.53,1.07,3.44,2.05C16.85,7.64,16.84,8.38,17.2,9c0.35,0.6,0.96,0.95,1.6,1c0.16,0.59,0.25,1.21,0.25,1.86 C19.05,12.61,18.93,13.33,18.71,14.01z\"></path></g></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AttractionsRound(size string, classes ...string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var20 = []any{classes}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var20...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" enable-background=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs("new 0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 string
		templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 122}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs("0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 161}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var25 string
		templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var20).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><g><rect fill=\"none\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var26 string
		templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 215}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var27 string
		templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 13, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" y=\"0\"></rect></g><g><g><path d=\"M10.44,18.75c0.37-0.46,0.94-0.75,1.57-0.75s1.19,0.29,1.56,0.75c0.39-0.09,0.76-0.21,1.12-0.36l-1.42-3.18 c-0.39,0.15-0.82,0.23-1.26,0.23c-0.46,0-0.9-0.09-1.3-0.25l-1.43,3.19C9.65,18.54,10.03,18.67,10.44,18.75z M5.16,10 C5,10.59,4.91,11.21,4.91,11.85c0,0.75,0.12,1.47,0.33,2.15c0.63,0.05,1.22,0.4,1.56,0.99c0.33,0.57,0.35,1.23,0.11,1.79 c0.27,0.27,0.56,0.53,0.87,0.76l1.52-3.39l0,0c-0.47-0.58-0.75-1.32-0.75-2.13c0-1.89,1.55-3.41,3.46-3.41s3.46,1.53,3.46,3.41 c0,0.82-0.29,1.57-0.78,2.16l1.5,3.35c0.32-0.24,0.62-0.5,0.9-0.79c-0.22-0.55-0.2-1.2,0.12-1.75c0.33-0.57,0.9-0.92,1.52-0.99 c0.22-0.68,0.34-1.41,0.34-2.16c0-0.64-0.09-1.27-0.25-1.86c-0.64-0.04-1.26-0.39-1.6-1c-0.36-0.62-0.35-1.36-0.03-1.95 c-0.91-0.98-2.1-1.71-3.44-2.05C13.4,5.6,12.74,6,12.01,6s-1.39-0.41-1.74-1.01C8.93,5.33,7.74,6.04,6.83,7.02 C7.16,7.62,7.18,8.37,6.81,9C6.45,9.62,5.82,9.97,5.16,10z M3.86,9.58C3.08,8.98,2.84,7.88,3.35,7s1.58-1.23,2.49-0.85 c1.11-1.17,2.56-2.03,4.18-2.42C10.15,2.75,10.99,2,12.01,2s1.85,0.75,1.98,1.73c1.63,0.39,3.07,1.24,4.18,2.42 c0.91-0.38,1.99-0.03,2.49,0.85c0.51,0.88,0.27,1.98-0.51,2.58c0.23,0.77,0.35,1.58,0.35,2.42s-0.12,1.65-0.35,2.42 c0.78,0.6,1.02,1.7,0.51,2.58s-1.58,1.23-2.49,0.85c-0.4,0.43-0.85,0.81-1.34,1.15l0.81,1.8c0.25,0.56-0.16,1.2-0.78,1.2h0 c-0.33,0-0.64-0.2-0.78-0.5l-0.75-1.67c-0.43,0.18-0.88,0.33-1.34,0.44C13.86,21.25,13.02,22,12.01,22s-1.85-0.75-1.98-1.73 C9.55,20.15,9.09,20,8.65,19.81L7.89,21.5C7.75,21.8,7.45,22,7.11,22H7.1c-0.62,0-1.03-0.64-0.77-1.2l0.82-1.83 c-0.47-0.33-0.91-0.71-1.3-1.12c-0.92,0.38-1.99,0.03-2.5-0.85s-0.27-1.98,0.51-2.58C3.62,13.65,3.51,12.84,3.51,12 S3.62,10.35,3.86,9.58z\"></path></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AttractionsSharp(size string, classes ...string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var28 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var28 == nil {
			templ_7745c5c3_Var28 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var29 = []any{classes}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var29...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" enable-background=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var30 string
		templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs("new 0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var31 string
		templ_7745c5c3_Var31, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var31))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var32 string
		templ_7745c5c3_Var32, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 122}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var32))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var33 string
		templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs("0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 161}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var34 string
		templ_7745c5c3_Var34, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var29).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var34))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><g><rect fill=\"none\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var35 string
		templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 215}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var36 string
		templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 17, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" y=\"0\"></rect></g><g><g><path d=\"M10.44,18.75c0.37-0.46,0.94-0.75,1.57-0.75s1.19,0.29,1.56,0.75c0.39-0.09,0.76-0.21,1.12-0.36l-1.42-3.18 c-0.39,0.15-0.82,0.23-1.26,0.23c-0.46,0-0.9-0.09-1.3-0.25l-1.43,3.19C9.65,18.54,10.03,18.67,10.44,18.75z M5.16,10 C5,10.59,4.91,11.21,4.91,11.85c0,0.75,0.12,1.47,0.33,2.15c0.63,0.05,1.22,0.4,1.56,0.99c0.33,0.57,0.35,1.23,0.11,1.79 c0.27,0.27,0.56,0.53,0.87,0.76l1.52-3.39l0,0c-0.47-0.58-0.75-1.32-0.75-2.13c0-1.89,1.55-3.41,3.46-3.41s3.46,1.53,3.46,3.41 c0,0.82-0.29,1.57-0.78,2.16l1.5,3.35c0.32-0.24,0.62-0.5,0.9-0.79c-0.22-0.55-0.2-1.2,0.12-1.75c0.33-0.57,0.9-0.92,1.52-0.99 c0.22-0.68,0.34-1.41,0.34-2.16c0-0.64-0.09-1.27-0.25-1.86c-0.64-0.04-1.26-0.39-1.6-1c-0.36-0.62-0.35-1.36-0.03-1.95 c-0.91-0.98-2.1-1.71-3.44-2.05C13.4,5.6,12.74,6,12.01,6s-1.39-0.41-1.74-1.01C8.93,5.33,7.74,6.04,6.83,7.02 C7.16,7.62,7.18,8.37,6.81,9C6.45,9.62,5.82,9.97,5.16,10z M3.86,9.58C3.08,8.98,2.84,7.88,3.35,7s1.58-1.23,2.49-0.85 c1.11-1.17,2.56-2.03,4.18-2.42C10.15,2.75,10.99,2,12.01,2s1.85,0.75,1.98,1.73c1.63,0.39,3.07,1.24,4.18,2.42 c0.91-0.38,1.99-0.03,2.49,0.85c0.51,0.88,0.27,1.98-0.51,2.58c0.23,0.77,0.35,1.58,0.35,2.42s-0.12,1.65-0.35,2.42 c0.78,0.6,1.02,1.7,0.51,2.58s-1.58,1.23-2.49,0.85c-0.4,0.43-0.85,0.81-1.34,1.15l1.34,3h-1.86l-0.97-2.17 c-0.43,0.18-0.88,0.33-1.34,0.44C13.86,21.25,13.02,22,12.01,22s-1.85-0.75-1.98-1.73C9.55,20.15,9.09,20,8.65,19.81L7.66,22H5.79 l1.36-3.03c-0.47-0.33-0.91-0.71-1.3-1.12c-0.92,0.38-1.99,0.03-2.5-0.85s-0.27-1.98,0.51-2.58C3.62,13.65,3.51,12.84,3.51,12 S3.62,10.35,3.86,9.58z\"></path></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AttractionsTwotone(size string, classes ...string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var37 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var37 == nil {
			templ_7745c5c3_Var37 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var38 = []any{classes}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var38...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" enable-background=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var39 string
		templ_7745c5c3_Var39, templ_7745c5c3_Err = templ.JoinStringErrs("new 0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var39))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var40 string
		templ_7745c5c3_Var40, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var40))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var41 string
		templ_7745c5c3_Var41, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 122}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var41))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var42 string
		templ_7745c5c3_Var42, templ_7745c5c3_Err = templ.JoinStringErrs("0 0 " + size + " " + size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 161}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var42))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var43 string
		templ_7745c5c3_Var43, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var38).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var43))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><g><rect fill=\"none\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var44 string
		templ_7745c5c3_Var44, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 215}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var44))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var45 string
		templ_7745c5c3_Var45, templ_7745c5c3_Err = templ.JoinStringErrs(size)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `attractions.templ`, Line: 21, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var45))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" y=\"0\"></rect></g><g><g><circle cx=\"11.98\" cy=\"12.02\" opacity=\".3\" r=\"1.5\"></circle><path d=\"M20.15,14.42c0.23-0.77,0.35-1.58,0.35-2.42s-0.12-1.65-0.35-2.42c0.78-0.6,1.02-1.7,0.51-2.58 c-0.51-0.88-1.58-1.23-2.49-0.85c-1.11-1.17-2.56-2.03-4.18-2.42C13.85,2.75,13.01,2,12,2s-1.85,0.75-1.98,1.73 C8.39,4.12,6.95,4.98,5.83,6.15C4.92,5.77,3.85,6.12,3.34,7C2.83,7.88,3.07,8.98,3.85,9.58C3.62,10.35,3.5,11.16,3.5,12 s0.12,1.65,0.35,2.42c-0.78,0.6-1.02,1.7-0.51,2.58c0.51,0.88,1.58,1.23,2.49,0.85c0.4,0.42,0.83,0.79,1.3,1.12L5.78,22h1.88 l0.98-2.19c0.44,0.19,0.9,0.34,1.38,0.46C10.15,21.25,10.99,22,12,22s1.85-0.75,1.98-1.73c0.46-0.11,0.91-0.26,1.34-0.44L16.3,22 h1.88l-1.34-3c0.48-0.34,0.93-0.72,1.34-1.15c0.91,0.38,1.99,0.03,2.49-0.85S20.93,15.02,20.15,14.42z M13.56,18.75 C13.19,18.29,12.63,18,12,18s-1.2,0.29-1.57,0.75c-0.4-0.09-0.79-0.21-1.16-0.37l1.43-3.19c0.4,0.16,0.84,0.25,1.3,0.25 c0.44,0,0.87-0.08,1.26-0.23l1.42,3.18C14.32,18.54,13.95,18.66,13.56,18.75z M10.48,12.02c0-0.83,0.67-1.5,1.5-1.5 c0.83,0,1.5,0.67,1.5,1.5s-0.67,1.5-1.5,1.5C11.15,13.52,10.48,12.85,10.48,12.02z M18.71,14.01c-0.61,0.07-1.18,0.41-1.52,0.99 c-0.32,0.56-0.34,1.2-0.12,1.75c-0.28,0.29-0.58,0.55-0.9,0.79l-1.5-3.35c0.49-0.59,0.78-1.34,0.78-2.16 c0-1.89-1.55-3.41-3.46-3.41c-1.91,0-3.46,1.53-3.46,3.41c0,0.8,0.28,1.54,0.75,2.13l-1.52,3.39c-0.31-0.23-0.6-0.48-0.87-0.76 C7.15,16.23,7.13,15.57,6.8,15c-0.34-0.59-0.93-0.94-1.56-0.99c-0.22-0.68-0.33-1.4-0.33-2.15c0-0.64,0.09-1.26,0.25-1.85 c0.66-0.03,1.3-0.38,1.65-1c0.37-0.63,0.35-1.38,0.01-1.98C7.74,6.05,8.93,5.34,10.27,5c0.34,0.59,0.99,1,1.73,1s1.39-0.4,1.73-1 c1.34,0.34,2.53,1.07,3.44,2.05C16.85,7.64,16.84,8.38,17.2,9c0.35,0.6,0.96,0.95,1.6,1c0.16,0.59,0.25,1.21,0.25,1.86 C19.05,12.61,18.93,13.33,18.71,14.01z\"></path></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate