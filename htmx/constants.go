package htmx

const (
	// SwapOuterHTML replaces the entire target element including its tag
	SwapOuterHTML = "outerHTML"
	// SwapInnerHTML replaces the contents inside the target element
	SwapInnerHTML = "innerHTML"
	// SwapTextContent replaces only the text content of the element, without parsing HTML
	SwapTextContent = "textContent"
	// SwapNone prevents swapping any content — useful for background requests
	SwapNone = "none"
	// SwapDelete removes the target element regardless of response
	SwapDelete = "delete"
	// SwapBeforeBegin inserts the response HTML before the target element
	SwapBeforeBegin = "beforebegin"
	// SwapAfterBegin inserts the response HTML at the start of the target element’s content
	SwapAfterBegin = "afterbegin"
	// SwapBeforeEnd inserts the response HTML at the end of the target element’s content
	SwapBeforeEnd = "beforeend"
	// SwapAfterEnd inserts the response HTML after the target element
	SwapAfterEnd = "afterend"
)

const (
	// TriggerClick triggers the request on click
	TriggerClick = "click"
	// TriggerChange triggers when an input or select element changes
	TriggerChange = "change"
	// TriggerSubmit triggers on form submission
	TriggerSubmit = "submit"
	// TriggerLoad triggers when the element is loaded
	TriggerLoad = "load"
	// TriggerEvery is a polling pattern (e.g., "every 2s")
	TriggerEvery = "every"
	// TriggerRevealed triggers when an element is revealed (e.g., scrolled into view)
	TriggerRevealed = "revealed"
	// TriggerIntersect triggers when the element intersects the viewport (IntersectionObserver)
	TriggerIntersect = "intersect"
)

const (
	// EventBeforeRequest fires before the request is issued
	EventBeforeRequest = "before-request"
	// EventSend fires when the request is sent
	EventSend = "send"
	// EventAfterRequest fires after the response is received
	EventAfterRequest = "after-request"
	// EventBeforeSwap fires before the DOM is updated with the response
	EventBeforeSwap = "before-swap"
	// EventAfterSwap fires after the DOM has been updated with the response
	EventAfterSwap = "after-swap"
	// EventResponseError fires when the server returns an error (e.g., 500)
	EventResponseError = "response-error"
	// EventSendError fires if the request failed to send (e.g., network error)
	EventSendError = "send-error"
)

const (
	// SyncAbort cancels the previous request and sends the new one (default)
	SyncAbort = "abort"
	// SyncDrop drops the new request if a previous one is in flight
	SyncDrop = "drop"
	// SyncReplace replaces the last queued request with the new one
	SyncReplace = "replace"
	// SyncQueue queues the new request to run after the current one finishes
	SyncQueue = "queue"
	// SyncQueueFirst queue the first request to show up while a request is in flight
	SyncQueueFirst = "queue first"
	// SyncQueueLast queue the last request to show up while a request is in flight
	SyncQueueLast = "queue last"
	// SyncQueueAll queue all requests that show up while a request is in flight
	SyncQueueAll = "queue all"
)

const (
	AttrGet         = "hx-get"
	AttrPost        = "hx-post"
	AttrPut         = "hx-put"
	AttrDelete      = "hx-delete"
	AttrPatch       = "hx-patch"
	AttrTarget      = "hx-target"
	AttrSwap        = "hx-swap"
	AttrSwapOob     = "hx-swap-oob"
	AttrTrigger     = "hx-trigger"
	AttrInclude     = "hx-include"
	AttrSelect      = "hx-select"
	AttrParams      = "hx-params"
	AttrPushURL     = "hx-push-url"
	AttrReplaceURL  = "hx-replace-url"
	AttrConfirm     = "hx-confirm"
	AttrDisable     = "hx-disable"
	AttrDisabledElt = "hx-disabled-elt"
	AttrBoost       = "hx-boost"
	AttrExt         = "hx-ext"
	AttrEncoding    = "hx-encoding"
	AttrHistoryElt  = "hx-history-elt"
	AttrIndicator   = "hx-indicator"
	AttrRequest     = "hx-request"
	AttrOn          = "hx-on"
	AttrVals        = "hx-vals"
	AttrVars        = "hx-vars"
	AttrHeaders     = "hx-headers"
	AttrInherit     = "hx-inherit"
	AttrDisinherit  = "hx-disinherit"
	AttrPreserve    = "hx-preserve"
	AttrPrompt      = "hx-prompt"
	AttrSync        = "hx-sync"
	AttrValidate    = "hx-validate"
)
