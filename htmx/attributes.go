package htmx

import (
	"github.com/go-andiamo/aitch"
	"sort"
	"strings"
)

// Get declares an "hx-get" attribute for issuing a GET request via htmx
//
// Reference: https://htmx.org/attributes/hx-get/
func Get(value ...any) aitch.Node {
	return aitch.NewAttribute(attrGet, value...)
}

// Post declares an "hx-post" attribute for issuing a POST request via htmx
//
// Reference: https://htmx.org/attributes/hx-post/
func Post(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPost, value...)
}

// Put declares an "hx-put" attribute for issuing a PUT request via htmx
//
// Reference: https://htmx.org/attributes/hx-put/
func Put(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPut, value...)
}

// Delete declares an "hx-delete" attribute for issuing a DELETE request via htmx
//
// Reference: https://htmx.org/attributes/hx-delete/
func Delete(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDelete, value...)
}

// Patch declares an "hx-patch" attribute for issuing a PATCH request via htmx
//
// Reference: https://htmx.org/attributes/hx-patch/
func Patch(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPatch, value...)
}

// Target declares an "hx-target" attribute to specify the element that will be updated by the response
//
// Reference: https://htmx.org/attributes/hx-target/
func Target(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTarget, value...)
}

// Swap declares an "hx-swap" attribute to control how the returned HTML is swapped into the DOM
//
// Reference: https://htmx.org/attributes/hx-swap/
func Swap(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSwap, value...)
}

// SwapOob declares an "hx-swap-oob" attribute for out-of-band swapping (e.g. "true", "beforebegin")
//
// Reference: https://htmx.org/attributes/hx-swap-oob/
func SwapOob(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSwapOob, value...)
}

// Trigger declares an "hx-trigger" attribute to define when the request should be triggered
//
// multiple declarations will be space-delimited and merged automatically
//
// Reference: https://htmx.org/attributes/hx-trigger/
func Trigger(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrTrigger, space, value...)
}

// Include declares an "hx-include" attribute to include additional inputs or elements in the request
//
// Reference: https://htmx.org/attributes/hx-include/
func Include(value ...any) aitch.Node {
	return aitch.NewAttribute(attrInclude, value...)
}

// Select declares an "hx-select" attribute to specify which part of the response to swap in
//
// Reference: https://htmx.org/attributes/hx-select/
func Select(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSelect, value...)
}

// Params declares an "hx-params" attribute to control which parameters are submitted with the request
//
// multiple declarations will be space-delimited and merged automatically
//
// Reference: https://htmx.org/attributes/hx-params/
func Params(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrParams, space, value...)
}

// PushURL declares an "hx-push-url" attribute to push the URL into the browser history
//
// Reference: https://htmx.org/attributes/hx-push-url/
func PushURL(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPushUrl, value...)
}

// ReplaceURL declares an "hx-replace-url" attribute to replace the current browser history entry
//
// Reference: https://htmx.org/attributes/hx-replace-url/
func ReplaceURL(value ...any) aitch.Node {
	return aitch.NewAttribute(attrReplaceUrl, value...)
}

// On sets a custom "hx-on:*" event attribute handler. The event name is concatenated with "hx-on:"
//
// # This is a raw escape hatch and is not type-safe — use with care
//
// Reference: https://htmx.org/attributes/hx-on/
func On(event string, value ...any) aitch.Node {
	return aitch.NewAttribute(append(attrOn, event...), value...)
}

// OnBeforeRequest declares an "hx-on:before-request" handler to run JS before the request is issued
//
// Reference: https://htmx.org/attributes/hx-on/
func OnBeforeRequest(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeRequest, value...)
}

// OnSend declares an "hx-on:send" handler to run JS when the request is sent
//
// Reference: https://htmx.org/attributes/hx-on/
func OnSend(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSend, value...)
}

// OnAfterRequest declares an "hx-on:after-request" handler to run JS after the response is received
//
// Reference: https://htmx.org/attributes/hx-on/
func OnAfterRequest(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAfterRequest, value...)
}

// OnBeforeSwap declares an "hx-on:before-swap" handler to run JS before the DOM is swapped
//
// Reference: https://htmx.org/attributes/hx-on/
func OnBeforeSwap(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeSwap, value...)
}

// OnAfterSwap declares an "hx-on:after-swap" handler to run JS after the DOM has been updated
//
// Reference: https://htmx.org/attributes/hx-on/
func OnAfterSwap(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAfterSwap, value...)
}

// OnResponseError declares an "hx-on:response-error" handler for failed HTTP response codes (e.g. 500)
//
// Reference: https://htmx.org/attributes/hx-on/
func OnResponseError(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnResponseError, value...)
}

// OnSendError declares an "hx-on:send-error" handler for network or client-side request failures
//
// Reference: https://htmx.org/attributes/hx-on/
func OnSendError(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSendError, value...)
}

// Confirm declares an "hx-confirm" attribute to show a confirmation dialog before issuing the request
//
// Reference: https://htmx.org/attributes/hx-confirm/
func Confirm(value ...any) aitch.Node {
	return aitch.NewAttribute(attrConfirm, value...)
}

// Disable declares an "hx-disable" attribute to disable htmx processing on the element
//
// Reference: https://htmx.org/attributes/hx-disable/
func Disable(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDisable, value...)
}

// DisabledElt declares an "hx-disabled-elt" attribute to specify the element to disable during the request
//
// Reference: https://htmx.org/attributes/hx-disabled-elt/
func DisabledElt(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDisabledElt, value...)
}

// Boost declares an "hx-boost" attribute to enable or disable boosted navigation on anchors and forms
//
// Reference: https://htmx.org/attributes/hx-boost/
func Boost(value ...any) aitch.Node {
	return aitch.NewAttribute(attrBoost, value...)
}

// Ext declares an "hx-ext" attribute to enable one or more htmx extensions
//
// Reference: https://htmx.org/attributes/hx-ext/
func Ext(value ...any) aitch.Node {
	return aitch.NewAttribute(attrExt, value...)
}

// Encoding declares an "hx-encoding" attribute to set the encoding type (e.g., "multipart/form-data")
//
// Reference: https://htmx.org/attributes/hx-encoding/
func Encoding(value ...any) aitch.Node {
	return aitch.NewAttribute(attrEncoding, value...)
}

// HistoryElt declares an "hx-history-elt" attribute to control which element updates browser history
//
// Reference: https://htmx.org/attributes/hx-history-elt/
func HistoryElt(value ...any) aitch.Node {
	return aitch.NewAttribute(attrHistoryElt, value...)
}

// Indicator declares an "hx-indicator" attribute to show/hide a loading indicator element during requests
//
// Reference: https://htmx.org/attributes/hx-indicator/
func Indicator(value ...any) aitch.Node {
	return aitch.NewAttribute(attrIndicator, value...)
}

// Request declares an "hx-request" attribute to control fine-grained request options (e.g., credentials)
//
// Reference: https://htmx.org/attributes/hx-request/
func Request(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRequest, value...)
}

// Vals declares an "hx-vals" attribute to include additional parameters in the request as JSON.
// Accepts a pre-serialized JSON string, or a Go map that your code converts to JSON.
//
// Reference: https://htmx.org/attributes/hx-vals/
func Vals(value ...any) aitch.Node {
	return aitch.NewAttribute(attrVals, value...)
}

// Vars declares an "hx-vars" attribute to define local JS variables for htmx processing.
//
// multiple declarations will be comma+space-delimited and merged automatically
//
// Reference: https://htmx.org/attributes/hx-vars/
func Vars(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrVars, commaSpace, value...)
}

// Var is used within Vars to denote a var name and value
//
// the value can be dynamic (e.g. aitch.DynamicValueKey or aitch.DynamicValueFunc)
func Var(name string, value any) aitch.Value {
	return aitch.NewConcatValue(name, colonSpace, value)
}

// Headers declares an "hx-headers" attribute to send additional HTTP headers as a JSON object.
//
// the value arg can be either normal any values (strings, int, aitch.DynamicValueKey etc.) or
// one or more HeadersMap (but you cannot mix HeadersMap with other types)
//
// Reference: https://htmx.org/attributes/hx-headers/
func Headers(value ...any) aitch.Node {
	foundHeadersMap := false
	foundOthers := false
	mergedHeadersMap := make(HeadersMap)
	for _, v := range value {
		if hm, ok := v.(HeadersMap); ok {
			foundHeadersMap = true
			if foundOthers {
				break
			}
			for k, val := range hm {
				mergedHeadersMap[k] = val
			}
		} else {
			foundOthers = true
			if foundHeadersMap {
				break
			}
		}
	}
	if foundOthers && foundHeadersMap {
		panic("htmx.Headers cannot mix HeadersMap with other value types")
	}
	if foundHeadersMap {
		return aitch.NewAttribute(attrHeaders, aitch.NewConcatValue(mergedHeadersMap.serialize(nil)...))
	}
	return aitch.NewAttribute(attrHeaders, value...)
}

// HeadersJS declares an "hx-headers" attribute that uses a JavaScript expression.
//
// The expression value is prefixed with "js:" in the rendered attribute.
//
// Only HeadersMap values are accepted (multiple ones are merged).
//
// Reference: https://htmx.org/attributes/hx-headers/
func HeadersJS(value ...HeadersMap) aitch.Node {
	var mergedHeadersMap HeadersMap
	if len(value) == 1 {
		mergedHeadersMap = value[0]
	} else {
		mergedHeadersMap = make(HeadersMap)
		for _, hm := range value {
			for k, v := range hm {
				mergedHeadersMap[k] = v
			}
		}
	}
	return aitch.NewAttribute(attrHeaders, aitch.NewConcatValue(mergedHeadersMap.serialize(jsPrefix)...))
}

type HeadersMap map[string]any

func (m HeadersMap) serialize(prefix []byte) []any {
	result := make([]any, 0, (len(m)*2)+3)
	if len(prefix) > 0 {
		result = append(result, prefix)
	}
	result = append(result, []byte{'{'})
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	delimit := ""
	for _, key := range keys {
		result = append(result, []byte(delimit+`'`+strings.ReplaceAll(key, `"`, `&quot;`)+`': `), m[key])
		delimit = ", "
	}
	result = append(result, []byte{'}'})
	return result
}

// Inherit declares an "hx-inherit" attribute to control which attributes are inherited from parent elements.
//
// multiple declarations will be space-delimited and merged automatically
//
// Reference: https://htmx.org/attributes/hx-inherit/
func Inherit(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrInherit, space, value...)
}

// Disinherit declares an "hx-disinherit" attribute to prevent specified attributes from being inherited.
//
// multiple declarations will be space-delimited and merged automatically
//
// Reference: https://htmx.org/attributes/hx-disinherit/
func Disinherit(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrDisinherit, space, value...)
}

// Sync declares an "hx-sync" attribute to control request synchronization (e.g. "this:drop" or "closest form:abort").
//
// Reference: https://htmx.org/attributes/hx-sync/
func Sync(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSync, value...)
}

// Preserve declares an "hx-preserve" attribute to preserve the element across DOM swaps.
//
// Reference: https://htmx.org/attributes/hx-preserve/
func Preserve(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPreserve, value...)
}

// Prompt declares an "hx-prompt" attribute to display a prompt dialog before sending the request.
// The result of the prompt is included in the request payload.
//
// Reference: https://htmx.org/attributes/hx-prompt/
func Prompt(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPrompt, value...)
}

// Validate declares an "hx-validate" attribute to control whether form validation should be run
// before the request is submitted. Accepts "true" or "false" (as string or bool).
//
// Reference: https://htmx.org/attributes/hx-validate/
func Validate(value ...any) aitch.Node {
	return aitch.NewAttribute(attrValidate, value...)
}

var (
	space               = []byte{' '}
	commaSpace          = []byte{',', ' '}
	colonSpace          = []byte{':', ' '}
	jsPrefix            = []byte("js:")
	attrGet             = []byte(AttrGet)
	attrPost            = []byte(AttrPost)
	attrPut             = []byte(AttrPut)
	attrDelete          = []byte(AttrDelete)
	attrPatch           = []byte(AttrPatch)
	attrTarget          = []byte(AttrTarget)
	attrSwap            = []byte(AttrSwap)
	attrSwapOob         = []byte(AttrSwapOob)
	attrTrigger         = []byte(AttrTrigger)
	attrInclude         = []byte(AttrInclude)
	attrSelect          = []byte(AttrSelect)
	attrParams          = []byte(AttrParams)
	attrPushUrl         = []byte(AttrPushURL)
	attrReplaceUrl      = []byte(AttrReplaceURL)
	attrOn              = []byte("hx-on:")
	attrOnBeforeRequest = []byte("hx-on:before-request")
	attrOnSend          = []byte("hx-on:send")
	attrOnAfterRequest  = []byte("hx-on:after-request")
	attrOnBeforeSwap    = []byte("hx-on:before-swap")
	attrOnAfterSwap     = []byte("hx-on:after-swap")
	attrOnResponseError = []byte("hx-on:response-error")
	attrOnSendError     = []byte("hx-on:send-error")
	attrConfirm         = []byte(AttrConfirm)
	attrDisable         = []byte(AttrDisable)
	attrDisabledElt     = []byte(AttrDisabledElt)
	attrBoost           = []byte(AttrBoost)
	attrExt             = []byte(AttrExt)
	attrEncoding        = []byte(AttrEncoding)
	attrHistoryElt      = []byte(AttrHistoryElt)
	attrIndicator       = []byte(AttrIndicator)
	attrRequest         = []byte(AttrRequest)
	attrVals            = []byte(AttrVals)
	attrVars            = []byte(AttrVars)
	attrHeaders         = []byte(AttrHeaders)
	attrInherit         = []byte(AttrInherit)
	attrDisinherit      = []byte(AttrDisinherit)
	attrSync            = []byte(AttrSync)
	attrPreserve        = []byte(AttrPreserve)
	attrPrompt          = []byte(AttrPrompt)
	attrValidate        = []byte(AttrValidate)
)
