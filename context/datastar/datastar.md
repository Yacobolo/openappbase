### POST /endpoint with Options Example

Source: https://data-star.dev/docs

Example of using @get action with various options.

```APIDOC
## Example Request with Options

```html
<button data-on:click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})" data-request-id="1"></button>
```
```

--------------------------------

### Configure Datastar Action Options

Source: https://data-star.dev/docs

This example demonstrates how to configure various options for a Datastar action, including filtering signals, setting custom headers, enabling open-when-hidden, and disabling request cancellation. It shows a button triggering a GET request with these configurations.

```html
<button data-on:click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})" ></button>
```

--------------------------------

### JavaScript SSE Event Handling Setup

Source: https://data-star.dev/docs

This JavaScript snippet shows the initial setup for creating a `ServerSentEventGenerator` instance, which also sends the necessary headers for SSE. It's the starting point for sending SSE events from a server.

```javascript
// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
ServerSentEventGenerator.stream(req, res, (stream) => {

```

--------------------------------

### C# SDK for Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar C# SDK to add Datastar as a service and patch signals asynchronously. It demonstrates patching signals with a delay between updates.

```csharp
 1using StarFederation.Datastar.DependencyInjection;
 2
 3// Adds Datastar as a service
 4builder.Services.AddDatastar();
 5
 6app.MapGet("/hal", async (IDatastarService datastarService) =>
 7{
 8    // Patches signals.
 9    await datastarService.PatchSignalsAsync(new { hal = "Affirmative, Dave. I read you" });
10
11    await Task.Delay(TimeSpan.FromSeconds(3));
12
13    await datastarService.PatchSignalsAsync(new { hal = "..." });
14});

```

--------------------------------

### Reading Signals in Ruby

Source: https://data-star.dev/docs

An example of setting up and reading signals in Ruby using the Datastar library. It initializes Datastar with the request and response objects and then accesses signals via `datastar.signals[:some_signal]`.

```ruby
1# Setup with request
2datastar = Datastar.new(request:, response:)
3
4# Read signals
5some_signal = datastar.signals[:some_signal]
```

--------------------------------

### Ruby SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Ruby SDK to create a streaming response and patch signals. It demonstrates how to instantiate a Datastar::Dispatcher and send patch signals with a delay.

```ruby
 1require 'datastar'
 2
 3# Create a Datastar::Dispatcher instance
 4
 5datastar = Datastar.new(request:, response:)
 6
 7# In a Rack handler, you can instantiate from the Rack env
 8# datastar = Datastar.from_rack_env(env)
 9
10# Start a streaming response
11datastar.stream do |sse|
12  # Patches signals
13  sse.patch_signals(hal: 'Affirmative, Dave. I read you.')
14
15  sleep 1
16  
17  sse.patch_signals(hal: '...')
18end

```

--------------------------------

### Go SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Go SDK to create an SSE generator and patch signals. It shows how to send patch signal data and include a delay.

```go
 1import (
 2    "github.com/starfederation/datastar-go/datastar"
 3)
 4
 5// Creates a new `ServerSentEventGenerator` instance.
 6sse := datastar.NewSSE(w, r)
 7
 8// Patches signals
 9sse.PatchSignals([]byte(`{hal: 'Affirmative, Dave. I read you.'}`))
10
11time.Sleep(1 * time.Second)
12
13sse.PatchSignals([]byte(`{hal: '...'}`))

```

--------------------------------

### Datastar Runtime Error Example

Source: https://data-star.dev/docs

Presents an example of a Datastar runtime error due to an invalid attribute key (`data-text-foo`), showing the console output with error details and a link to more information.

```log
Uncaught datastar runtime error: textKeyNotAllowed
More info: https://data-star.dev/errors/key_not_allowed?metadata=%7B%22plugin%22%3A%7B%22name%22%3A%22text%22%2C%22type%22%3A%22attribute%22%7D%2C%22element%22%3A%7B%22id%22%3A%22%22%2C%22tag%22%3A%22DIV%22%7D%2C%22expression%22%3A%7B%22rawKey%22%3A%22textFoo%22%2C%22key%22%3A%22foo%22%2C%22value%22%3A%22%22%2C%22fnContent%22%3A%22%22%7D%7D
Context: {
    "plugin": {
        "name": "text",
        "type": "attribute"
    },
    "element": {
        "id": "",
        "tag": "DIV"
    },
    "expression": {
        "rawKey": "textFoo",
        "key": "foo",
        "value": "",
        "fnContent": ""
    }
}
```

--------------------------------

### Datastar Action Options

Source: https://data-star.dev/docs

Explains the available options for Datastar actions, which are passed as a second argument.

```APIDOC
## Datastar Action Options

All Datastar actions accept an `options` object as their second argument to customize behavior.

### Options Available:

- **`contentType`** (`string`): Specifies the content type for the request. Accepts `json` (default) or `form`. `form` content type triggers form validation and submission without sending signals.
- **`filterSignals`** (`object`): Filters signals based on inclusion and exclusion criteria.
    - `include` (`RegExp`): Regular expression to match signal paths to include. Defaults to all signals (`/.*/`).
    - `exclude` (`RegExp`): Regular expression to exclude specific signal paths. Defaults to signals without a `_` prefix (`/(^_|._).*/`).
- **`selector`** (`string` | `null`): Specifies a form element to use when `contentType` is `form`. If `null`, the closest form is used. Defaults to `null`.
- **`headers`** (`object`): An object containing custom headers to send with the request.
- **`openWhenHidden`** (`boolean`): Determines if the connection should remain open when the page is hidden. Defaults to `false`.
- **`retry`** (`string`): Configures retry behavior for requests. Options include `'auto'` (default), `'error'`, `'always'`, and `'never'`.
- **`retryInterval`** (`number`): The interval in milliseconds between retries. Defaults to `1000`.
- **`retryScaler`** (`number`): A multiplier for scaling retry wait times. Defaults to `2`.
- **`retryMaxWaitMs`** (`number`): The maximum wait time in milliseconds between retries. Defaults to `30000`.
- **`retryMaxCount`** (`number`): The maximum number of retry attempts. Defaults to `10`.
- **`requestCancellation`** (`string` | `AbortController`): Controls request cancellation behavior. Options include `'auto'` (default), `'disabled'`, or an `AbortController` instance.
```

--------------------------------

### Datastar Backend SDKs

Source: https://data-star.dev/docs

Overview of the available backend Software Development Kits (SDKs) for Datastar, simplifying the generation of SSE events.

```APIDOC
# SDKs

Datastar provides backend SDKs to simplify the process of generating SSE events. Below is a list of available SDKs by language:

- **Clojure**: SDK and helper libraries available.
- **C# (.NET)**: SDK for working with Datastar.
- **Go**: SDK available.
- **Java**: SDK and examples available.
- **Kotlin**: SDK available.
- **PHP**: SDK available.
  - **Craft CMS**: Plugin for integrating Datastar with Craft CMS.
  - **Laravel**: Package for integrating Datastar with Laravel.
- **Python**: SDK and PyPI package available.
- **Ruby**: SDK available.
- **Rust**: SDK available.
  - **Rama**: Module for integrating Datastar with Rama (Rust-based HTTP proxy).
```

--------------------------------

### Java SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Java SDK to send SSE patch signals. It demonstrates creating a `ServerSentEventGenerator` and sending patch signal data with a delay.

```java
 1import starfederation.datastar.utils.ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
 6
 7// Patches signals.
 8generator.send(PatchSignals.builder()
 9    .data("{\"hal\": \"Affirmative, Dave. I read you.\"}")
10    .build()
11);
12
13Thread.sleep(1000);
14
15generator.send(PatchSignals.builder()
16    .data("{\"hal\": \"...\"}")
17    .build()
18);

```

--------------------------------

### Patch DOM Elements with Python datastar_py

Source: https://data-star.dev/guide/getting_started

Shows how to patch HTML elements using the Python datastar_py library. This asynchronous example yields SSE patch elements, waits for a second, and then yields another patch element to update the DOM.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response
import asyncio

@app.get('/open-the-bay-doors')
@datastar_response
async def open_doors(request):
    yield SSE.patch_elements('<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>')
    await asyncio.sleep(1)
    yield SSE.patch_elements('<div id="hal">Waiting for an order...</div>')
```

--------------------------------

### PHP SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar PHP SDK to generate SSE patch signals. It demonstrates creating a `ServerSentEventGenerator` and sending patch signal data with a delay.

```php
 1use starfederation\datastar\ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4$sse = new ServerSentEventGenerator();
 5
 6// Patches signals.
 7$sse->patchSignals(['hal' => 'Affirmative, Dave. I read you.']);
 8
 9sleep(1);
10
11$sse->patchSignals(['hal' => '...']);

```

--------------------------------

### Send GET Request with Datastar

Source: https://data-star.dev/guide/getting_started

This snippet demonstrates how to use the `@get()` action in Datastar to send a GET request to a specified URL. The response, if HTML, will be morphed into the DOM based on element IDs. This is useful for fetching data or triggering backend actions that update the UI.

```html
<button data-on:click="@get('/endpoint')">
    Open the pod bay doors, HAL.
</button>
<div id="hal"></div>
```

--------------------------------

### Datastar Attribute Order and Initialization

Source: https://data-star.dev/docs

Demonstrates the order of attribute processing in Datastar, specifically highlighting the importance of `data-indicator` preceding `data-init` for fetch requests.

```html
<div data-indicator:fetching data-init="@get('/endpoint')"></div>
```

--------------------------------

### Datastar Attribute Casing Conventions

Source: https://data-star.dev/docs

Explains Datastar's handling of attribute casing, converting `data-*` attributes to `camelCase` for signals and `kebab-case` for other attributes by default. Includes examples of manual casing control using `__case` and object syntax.

```html
<div data-signals:my-signal></div>
<div data-class:text-blue-700></div>
<div data-on:rocket-launched></div>
<div data-on:widget-loaded__case.camel></div>
```

--------------------------------

### GET Request with @get()

Source: https://data-star.dev/docs

Sends a GET request to a specified URI using the Fetch API. The response must contain Datastar SSE events. Supports options for controlling request behavior, like `openWhenHidden` and `contentType`.

```html
<button data-on:click="@get('/endpoint')"></button>

<button data-on:click="@get('/endpoint', {openWhenHidden: true})"></button>

<button data-on:click="@get('/endpoint', {contentType: 'form'})"></button>

<form enctype="multipart/form-data">
    <input type="file" name="file" />
    <button data-on:click="@get('/endpoint', {contentType: 'form'})"></button>
</form>
```

--------------------------------

### Kotlin SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Kotlin SDK to generate SSE patch signals. It shows how to create a `ServerSentEventGenerator` and send patch signal data with a delay.

```kotlin
 1val generator = ServerSentEventGenerator(response)
 2
 3generator.patchSignals(
 4    signals = """{"hal": "Affirmative, Dave. I read you."} """,
 5)
 6
 7Thread.sleep(ONE_SECOND)
 8
 9generator.patchSignals(
10    signals = """{"hal": "..."} """,
11)

```

--------------------------------

### Customizing Datastar Request Cancellation

Source: https://data-star.dev/docs

Shows how to customize request cancellation behavior in Datastar. The first example disables automatic cancellation, allowing multiple concurrent requests. The second example demonstrates using a custom AbortController for fine-grained control over request cancellation.

```html
<!-- Allow concurrent requests (no automatic cancellation) -->
<button data-on:click="@get('/endpoint', {requestCancellation: 'disabled'})">Allow Multiple</button>

<!-- Custom abort controller for fine-grained control -->
<div data-signals:controller="new AbortController()">
    <button data-on:click="@get('/endpoint', {requestCancellation: $controller})">Start Request</button>
    <button data-on:click="$controller.abort()">Cancel Request</button>
</div>
```

--------------------------------

### Response Handling

Source: https://data-star.dev/docs

Details how Datastar handles different response content types from the backend.

```APIDOC
## Response Handling

Datastar automatically handles various response content types:

- **`text/event-stream`**: Processes standard Server-Sent Events (SSE) with Datastar SSE events.
- **`text/html`**: Patches HTML elements into the DOM.
- **`application/json`**: Patches JSON encoded signals.
- **`text/javascript`**: Executes JavaScript code in the browser.

### `text/html` Response Headers:

When returning `text/html`, the server can optionally include the following headers to control patching behavior:

- **`datastar-selector`** (`string`): A CSS selector specifying the target elements for patching.
- **`datastar-mode`** (`string`): The patching mode. Options include `outer`, `inner`, `remove`, `replace`, `prepend`, `append`, `before`, `after`. Defaults to `outer`.
- **`datastar-use-view-transition`** (`boolean`): Enables the use of the View Transition API for patching. Defaults to `false`.

#### Example HTML Response Headers:

```javascript
response.headers.set('Content-Type', 'text/html');
response.headers.set('datastar-selector', '#my-element');
response.headers.set('datastar-mode', 'inner');
response.body = '<p>New content</p>';
```
```

--------------------------------

### Including Datastar Aliased Bundle

Source: https://data-star.dev/docs

Shows how to include the Datastar framework with aliased attributes using a script tag pointing to a CDN.

```html
<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar-aliased.js"></script>
```

--------------------------------

### Python SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Python SDK with a Sanic web framework integration to send SSE patch signals. It shows how to yield SSE events and include asynchronous delays.

```python
1from datastar_py import ServerSentEventGenerator as SSE
2from datastar_py.sanic import datastar_response
3
4@app.get('/do-you-read-me')
5@datastar_response
6async def open_doors(request):
7    yield SSE.patch_signals({"hal": "Affirmative, Dave. I read you."})
8    await asyncio.sleep(1)
9    yield SSE.patch_signals({"hal": "..."})

```

--------------------------------

### Patch Elements using SSE (C#)

Source: https://data-star.dev/guide/getting_started

This C# code demonstrates how to use Datastar's SDK to patch elements into the DOM via Server-Sent Events (SSE). It sets up Datastar as a service and then, within an HTTP GET handler, sends an initial message and updates it after a delay. This allows for dynamic UI updates in response to backend events.

```csharp
using StarFederation.Datastar.DependencyInjection;

// Adds Datastar as a service
builder.Services.AddDatastar();

app.MapGet("/", async (IDatastarService datastarService) =>
{
    // Patches elements into the DOM.
    await datastarService.PatchElementsAsync(@"<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>");

    await Task.Delay(TimeSpan.FromSeconds(1));

    await datastarService.PatchElementsAsync(@"<div id=\"hal\">Waiting for an order...</div>");
});
```

--------------------------------

### Content Security Policy Configuration

Source: https://data-star.dev/docs

This example demonstrates how to configure a Content Security Policy (CSP) to allow DataStar to evaluate expressions using the `Function()` constructor. The `unsafe-eval` directive must be included in the `script-src` directive.

```html
<meta http-equiv="Content-Security-Policy" 
    content="script-src 'self' 'unsafe-eval';"
>
```

--------------------------------

### Python SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Python example shows how to use Datastar's `ServerSentEventGenerator` within a Litestar endpoint to return SSE events. It demonstrates returning a `DatastarResponse` containing both patched elements and signals.

```python
1from datastar_py import ServerSentEventGenerator as SSE
2from datastar_py.litestar import DatastarResponse
3
4async def endpoint():
5    return DatastarResponse([
6        SSE.patch_elements('<div id="question">What do you put in a toaster?</div>'),
7        SSE.patch_signals({"response": "", "answer": "bread"})
8    ])
```

--------------------------------

### Datastar Expressions with Signals and Element Access

Source: https://data-star.dev/docs

Illustrates the use of Datastar expressions within `data-*` attributes, demonstrating how to access signals (prefixed with `$`) and the element context (`el`).

```html
<div id="bar" data-text="$foo + el.id"></div>
```

--------------------------------

### Datastar Expression: Multiple Statements

Source: https://data-star.dev/docs

This HTML example shows how to execute multiple statements within a single Datastar expression by separating them with a semicolon, useful for updating signals and triggering actions.

```html
<div data-signals:foo="1">
    <button data-on:click="$landingGearRetracted = true; @post('/launch')">
        Force launch
    </button>
</div>
```

--------------------------------

### Clojure SDK for SSE Patch Signals

Source: https://data-star.dev/docs

Example using the Datastar Clojure SDK to generate SSE events for patching signals. It shows how to create an SSE response and send patch signals with a delay.

```clojure
 1;; Import the SDK's api and your adapter
 2(require
 3  '[starfederation.datastar.clojure.api :as d*]
 4  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
 5
 6;; in a ring handler
 7(defn handler [request]
 8  ;; Create an SSE response
 9  (->sse-response request
10                  {on-open
11                   (fn [sse] 
12                     ;; Patches signal.
13                     (d*/patch-signals! sse "{hal: 'Affirmative, Dave. I read you.'}")
14                     (Thread/sleep 1000)
15                     (d*/patch-signals! sse "{hal: '...'}"))}))

```

--------------------------------

### HTML for Nesting Signals in Datastar

Source: https://data-star.dev/docs

Demonstrates how to use HTML attributes to define nested signals in Datastar. It shows examples using dot-notation, object syntax, and two-way binding for attributes like `data-signals` and `data-bind`.

```html
1<div data-signals:foo.bar="1"></div>
```

```html
1<div data-signals="{foo: {bar: 1}}"></div>
```

```html
1<input data-bind:foo.bar />
```

```html
1<div data-signals="{menu: {isOpen: {desktop: false, mobile: false}}}">
2    <button data-on:click="@toggleAll({include: /^menu\.isOpen\./})">
3        Open/close menu
4    </button>
5</div>
```

--------------------------------

### Show Loading Indicator During Request with Data-Star

Source: https://data-star.dev/docs

This example uses the `data-indicator` attribute to manage a loading state. The `data-indicator:fetching` attribute sets a signal to true while a request is in flight and false otherwise. This signal can then be used with `data-show` or `data-class` to display loading indicators. It requires an element with `id='question'` and a button to trigger the request.

```html
<div id="question"></div>
<button
    data-on:click="@get('/actions/quiz')"
    data-indicator:fetching
>
    Fetch a question
</button>
<div data-class:loading="$fetching" class="indicator"></div>
```

--------------------------------

### Reading Nested Signals in Kotlin (Ktor)

Source: https://data-star.dev/docs

Provides an example of reading signals in Kotlin using the Ktor framework. It defines a `Signals` data class and uses `Json.decodeFromString` with a `jsonUnmarshaller` to parse the request body.

```kotlin
 1@Serializable
 2data class Signals(
 3    val foo: String,
 4)
 5
 6val jsonUnmarshaller: JsonUnmarshaller<Signals> = { json -> Json.decodeFromString(json) }
 7
 8val request: Request =
 9    postRequest(
10        body =
11            """
12            {
13                "foo": "bar"
14            }
15            """.trimIndent()
16)
17
18val signals = readSignals(request, jsonUnmarshaller)
```

--------------------------------

### Install Datastar from local file

Source: https://data-star.dev/guide

If you prefer to host Datastar yourself, include the script from its local path after downloading or bundling it.

```html
<script type="module" src="/path/to/datastar.js"></script>
```

--------------------------------

### Send Multiple SSE Events (Patch Elements and Signals)

Source: https://data-star.dev/docs

This showcases sending multiple Server-Sent Events (SSE) in a single response. It includes examples for patching HTML elements and sending patch signals, demonstrating the flexibility of SSE for real-time updates.

```datastar
(d*/patch-elements! sse "<div id=\"question\">...</div>")
(d*/patch-elements! sse "<div id=\"instructions\">...</div>")
(d*/patch-signals! sse "{answer: '...', prize: '...'}")
```

```csharp
datastarService.PatchElementsAsync(@"<div id=\"question\">...</div>");
datastarService.PatchElementsAsync(@"<div id=\"instructions\">...</div>");
datastarService.PatchSignalsAsync(new { answer = "...", prize = "..." } );
```

```javascript
sse.PatchElements(`<div id="question">...</div>`)
sse.PatchElements(`<div id="instructions">...</div>`)
sse.PatchSignals([]byte(`{answer: '...', prize: '...'}`))
```

```java
generator.send(PatchElements.builder()
    .data("<div id=\"question\">...</div>")
    .build()
);
generator.send(PatchElements.builder()
    .data("<div id=\"instructions\">...</div>")
    .build()
);
generator.send(PatchSignals.builder()
    .data("{\"answer\": \"...\", \"prize\": \"...\"}")
    .build()
);
```

```python
generator.patchElements(
    elements = """<div id=\"question\">...</div>""",
)
generator.patchElements(
    elements = """<div id=\"instructions\">...</div>""",
)
generator.patchSignals(
    signals = "{\"answer\": \"...\", \"prize\": \"...\"}",
)
```

```php
$sse->patchElements('<div id="question">...</div>');
$sse->patchElements('<div id="instructions">...</div>');
$sse->patchSignals(['answer' => '...', 'prize' => '...']);
```

```ruby
return DatastarResponse([
    SSE.patch_elements('<div id="question">...</div>'),
    SSE.patch_elements('<div id="instructions">...</div>'),
    SSE.patch_signals({"answer": "...", "prize": "..."})
])
```

```ruby
datastar.stream do |sse|
  sse.patch_elements('<div id="question">...</div>')
  sse.patch_elements('<div id="instructions">...</div>')
  sse.patch_signals(answer: '...', prize: '...')
end
```

```ruby
yield PatchElements::new("<div id='question'>...</div>").into()
yield PatchElements::new("<div id='instructions'>...</div>").into()
yield PatchSignals::new("{answer: '...', prize: '...'}").into()
```

```javascript
stream.patchElements('<div id="question">...</div>');
stream.patchElements('<div id="instructions">...</div>');
stream.patchSignals({'answer': '...', 'prize': '...'});
```

--------------------------------

### data-init with Delay Modifier

Source: https://data-star.dev/docs

Demonstrates using the `__delay` modifier with `data-init` to introduce a delay before executing the initialization expression. This allows for controlled timing of initial actions.

```html
<div data-init__delay.500ms="$count = 1"></div>
```

--------------------------------

### Kotlin SDK Example

Source: https://data-star.dev/how_tos/load_more_list_items

Example implementation in Kotlin demonstrating how to read signals and generate Server-Sent Events for updating the UI.

```APIDOC
## Kotlin SDK Example

### Description
This Kotlin code snippet shows how to read signals from a request and use `ServerSentEventGenerator` to send patch events for elements and signals, or to remove elements.

### Method
N/A (Server-side logic)

### Endpoint
N/A (Server-side logic)

### Parameters
#### Query Parameters
N/A

#### Request Body
N/A (Signals are read from request context)

### Request Example
```kotlin
@Serializable
data class OffsetSignals(
    val offset: Int,
)

val signals = 
    readSignals(
        request,
        { json: String -> Json.decodeFromString<OffsetSignals>(json) },
    )

val max = 5
val limit = 1
val offset = signals.offset

val generator = ServerSentEventGenerator(response)

// Logic to patch elements, signals, or remove elements based on offset
// (Illustrative, actual implementation details may vary)
if (offset < max) {
    // Example: Patching new item
    // generator.patchElements(htmlElement, PatchMode.Append, "#list")
    // Example: Patching signals
    // generator.patchSignals(OffsetSignals(offset + limit))
} else {
    // Example: Removing button
    // generator.removeElement("#load-more")
}
```

### Response
#### Success Response (200)
Server-Sent Events (SSE) containing `datastar-patch-elements` and `datastar-patch-signals` events.

#### Response Example
```
# See Backend Response section for SSE examples
```
```

--------------------------------

### Trigger GET Request on Click with Data-Star

Source: https://data-star.dev/docs

This code illustrates how to use the `data-on:click` attribute to trigger a `GET` request to a specified endpoint when an element is clicked. It also shows how to bind signals and computed properties for dynamic UI updates. The server response can modify DOM elements and signals.

```html
<div
    data-signals="{response: '', answer: ''}"
    data-computed:correct="$response.toLowerCase() == $answer"
>
    <div id="question"></div>
    <button data-on:click="@get('/actions/quiz')">Fetch a question</button>
    <button
        data-show="$answer != ''"
        data-on:click="$response = prompt('Answer:') ?? ''"
    >
        BUZZ
    </button>
    <div data-show="$response != ''">
        You answered “<span data-text="$response"></span>”.
        <span data-show="$correct">That is correct ✅</span>
        <span data-show="!$correct">
        The correct answer is “<span data-text="$answer"></span>” 🤷
        </span>
    </div>
</div>
```

--------------------------------

### Go SDK Example

Source: https://data-star.dev/how_tos/load_more_list_items

Example implementation in Go using the Datastar SDK for handling HTTP requests and sending Server-Sent Events.

```APIDOC
## Go SDK Example

### Description
This Go code demonstrates how to use the Datastar SDK to handle incoming requests, read signals, and send SSE events to patch elements, append content, and update signals or remove elements.

### Method
GET

### Endpoint
`/how_tos/load_more/data` (or similar, triggered by client)

### Parameters
#### Query Parameters
N/A

#### Request Body
N/A (Signals are read from request context)

### Request Example
```go
import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/starfederation/datastar-go/datastar"
)

type OffsetSignals struct {
    Offset int `json:"offset"`
}

signals := &OffsetSignals{}
if err := datastar.ReadSignals(r, signals); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
}

max := 5
limit := 1
offset := signals.Offset

sse := datastar.NewSSE(w, r)

if offset < max {
    newOffset := offset + limit
    sse.PatchElements(fmt.Sprintf(`<div>Item %d</div>`, newOffset),
        datastar.WithSelectorID("list"),
        datastar.WithModeAppend(),
    )
    if newOffset < max {
        sse.PatchSignals([]byte(fmt.Sprintf(`{offset: %d}`, newOffset)))
    } else {
        sse.RemoveElements(`#load-more`)
    }
}
```

### Response
#### Success Response (200)
Server-Sent Events (SSE) containing `datastar-patch-elements` and `datastar-patch-signals` events.

#### Response Example
```
# See Backend Response section for SSE examples
```
```

--------------------------------

### C#: Registering and Using Datastar Service for DOM Patching

Source: https://data-star.dev/docs

This C# code demonstrates how to register Datastar as a service and then use its `PatchElementsAsync` method to update the DOM. It simulates sending an initial message and then updating it after a delay.

```csharp
1using StarFederation.Datastar.DependencyInjection;
2
3// Adds Datastar as a service
4builder.Services.AddDatastar();
5
6app.MapGet("/", async (IDatastarService datastarService) =>
7{
8    // Patches elements into the DOM.
9    await datastarService.PatchElementsAsync(@"<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>");
10
11    await Task.Delay(TimeSpan.FromSeconds(1));
12
13    await datastarService.PatchElementsAsync(@"<div id=\"hal\">Waiting for an order...</div>");
14});

```

--------------------------------

### Reading Signals in Python (FastAPI)

Source: https://data-star.dev/docs

Shows how to read signals in Python using FastAPI and the Datastar library. It utilizes the `@datastar_response` decorator and the `read_signals` utility function to get signal data from the request.

```python
1from datastar_py.fastapi import datastar_response, read_signals
2
3@app.get("/updates")
4@datastar_response
5async def updates(request: Request):
6    # Retrieve a dictionary with the current state of the signals from the frontend
7    signals = await read_signals(request)
```

--------------------------------

### data-init for Initialization Expressions

Source: https://data-star.dev/docs

Explains the `data-init` attribute, which executes an expression upon initialization. This can occur during page load, DOM patching, or attribute modification, useful for setting initial states.

```html
<div data-init="$count = 1"></div>
```

--------------------------------

### Server-Sent Events (SSE) for Datastar Patch Elements with Script

Source: https://data-star.dev/docs

An example of Server-Sent Events (SSE) formatted for Datastar. This specifically shows how to send a 'datastar-patch-elements' event containing HTML, including a script tag, to be rendered and executed on the client.

```text
event: datastar-patch-elements
data: elements <div id="hal">
data:     <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
data: </div>

```

--------------------------------

### Patch DOM Elements with PHP ServerSentEventGenerator

Source: https://data-star.dev/guide/getting_started

Demonstrates patching HTML elements into the DOM using the PHP ServerSentEventGenerator class. It sends an initial message, waits for 1 second, and then sends a subsequent message to update the same element.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
$sse = new ServerSentEventGenerator();

// Patches elements into the DOM.
$sse->patchElements(
    '<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>'
);

sleep(1);

$sse->patchElements(
    '<div id="hal">Waiting for an order...</div>'
);
```

--------------------------------

### Patching Elements and Signals

Source: https://data-star.dev/docs

Demonstrates how to patch HTML elements into the DOM and update frontend signals using the `stream.patchElements()` and `stream.patchSignals()` functions.

```APIDOC
## PATCH /actions/quiz (Simulated)

### Description
This endpoint is simulated via frontend actions. It's triggered by a user click and demonstrates patching HTML elements into the DOM and updating frontend signals. The `stream.patchElements()` function updates a specific DOM element, and `stream.patchSignals()` updates signal values.

### Method
GET (Triggered by frontend action)

### Endpoint
/actions/quiz

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```javascript
// Frontend JavaScript triggering the action
stream.patchElements(`<div id="question">What do you put in a toaster?</div>`);
stream.patchSignals({'response':  '', 'answer': 'bread'});
```

### Response
#### Success Response (200)
Signals and DOM elements are updated directly by the frontend functions. The server's response is interpreted as events to modify the DOM and signals.

#### Response Example
```json
// Example of server-sent events that would trigger frontend updates
// These are not direct JSON responses but instructions for the frontend.
{
  "type": "patch_elements",
  "payload": {
    "html": "<div id=\"question\">What do you put in a toaster?</div>"
  }
}
{
  "type": "patch_signals",
  "payload": {
    "signals": {
      "response": "",
      "answer": "bread"
    }
  }
}
```
```

--------------------------------

### Clojure SDK Example

Source: https://data-star.dev/how_tos/load_more_list_items

Example implementation in Clojure using the Datastar SDK to handle the 'load more' logic, including patching elements and signals.

```APIDOC
## Clojure SDK Example

### Description
This Clojure code demonstrates how to handle the `load_more` request. It reads the current offset, patches the new item into the `#list` container using append mode, and either patches the updated offset signal or removes the load more button if the maximum offset is reached.

### Method
POST

### Endpoint
`/how_tos/load_more/data` (or similar, triggered by client)

### Parameters
#### Query Parameters
N/A

#### Request Body
N/A (Signals are read from request headers/context)

### Request Example
```clojure
(require
  '[starfederation.datastar.clojure.api :as d*]
  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]]
  '[some.hiccup.library :refer [html]]
  '[some.json.library :refer [read-json-str write-json-str]]))

(def max-offset 5)

(defn handler [ring-request]
  (->sse-response ring-request
    {on-open
     (fn [sse]
       (let [d*-signals (-> ring-request d*/get-signals read-json-str)
             offset (get d*-signals "offset")
             limit 1
             new-offset (+ offset limit)]

         (d*/patch-elements! sse
                             (html [:div "Item " new-offset])
                             {d*/selector   "#list"
                              d*/merge-mode d*/mm-append})

         (if (< new-offset max-offset)
           (d*/patch-signals! sse (write-json-str {"offset" new-offset}))
           (d*/remove-fragment! sse "#load-more")

           (d*/close-sse! sse)))}
    ))
```

### Response
#### Success Response (200)
Server-Sent Events (SSE) containing `datastar-patch-elements` and `datastar-patch-signals` events.

#### Response Example
```
# See Backend Response section for SSE examples
```
```

--------------------------------

### Generate SSE for Patching Elements (Kotlin)

Source: https://data-star.dev/guide/getting_started

This Kotlin code shows how to use Datastar's `ServerSentEventGenerator` to patch elements into the DOM via SSE. It sends an initial HTML element and then updates it after a delay, illustrating real-time DOM manipulation. This approach is effective for dynamic user interfaces.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>""",
)

Thread.sleep(ONE_SECOND)


```

--------------------------------

### Data Indicator for Loading States

Source: https://data-star.dev/docs

Shows how to use the `data-indicator` attribute to manage loading states. The `data-indicator` attribute sets a signal to `true` while a request is in flight, allowing for visual loading indicators.

```APIDOC
## GET /actions/quiz (with Loading Indicator)

### Description
This endpoint demonstrates using the `data-indicator` attribute to visually indicate when a request is in progress. The `data-indicator:<signal_name>` attribute on an element will set the specified signal to `true` during the request and `false` afterwards. This signal can then be used to conditionally display loading indicators.

### Method
GET

### Endpoint
/actions/quiz

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```html
<button
    data-on:click="@get('/actions/quiz')"
    data-indicator:fetching
>
    Fetch a question
</button>
<div data-class:loading="$fetching" class="indicator"></div>
```

### Response
#### Success Response (200)
Upon successful completion, the `fetching` signal will be set to `false`, and any element bound to this signal (e.g., via `data-class:loading`) will update accordingly.

#### Response Example
```json
// Server response for a successful request (content depends on the action)
{
  "message": "Question fetched successfully"
}
```
```

--------------------------------

### Configure Datastar GET Request with Options

Source: https://data-star.dev/reference/actions

This example demonstrates how to configure a Datastar GET request using various options. It includes signal filtering, custom headers, enabling open-when-hidden, and disabling request cancellation. This is useful for fine-tuning network requests and their behavior.

```html
<button data-on:click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})" data-on-click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})"></button>
```

--------------------------------

### SSE Event: datastar-patch-signals

Source: https://data-star.dev/docs

Handles patching of signals into the existing signals on the page. Allows for conditional updates and setting signal values.

```APIDOC
## SSE Event: datastar-patch-signals

### Description
Patches signals into the existing signals on the page. The `onlyIfMissing` directive controls whether to update signals only if they do not already exist. The `signals` directive expects a valid `data-signals` attribute value.

### Event Type
`datastar-patch-signals`

### Data Directives
- **`onlyIfMissing <boolean>`**: If true, updates signals only if a signal with that name does not yet exist. Defaults to `false`.
- **`signals <json_object>`**: A JSON object representing the signals to patch, where keys are signal names and values are their new values. Setting a value to `null` removes the signal.

### Request Example (Patching Signals)
```
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}
```

### Request Example (Removing Signals)
```
event: datastar-patch-signals
data: signals {foo: null, bar: null}
```

### Request Example (Conditional Patching)
```
event: datastar-patch-signals
data: onlyIfMissing true
data: signals {foo: 1, bar: 2}
```

### Response
N/A (This is a client-side event triggered by the server).
```

--------------------------------

### Set Multiple HTML Attributes using data-attr

Source: https://data-star.dev/docs

This example shows how to use the `data-attr` attribute to set multiple HTML attributes on an element simultaneously. It accepts a key-value pair object where keys are attribute names and values are Datastar expressions.

```html
<div data-attr="{title: $foo, disabled: $bar}"></div>
```

--------------------------------

### Request Cancellation Behavior

Source: https://data-star.dev/docs

Explains the default request cancellation behavior and how to control it.

```APIDOC
## Request Cancellation

By default, Datastar automatically cancels any ongoing request on the same element when a new request is initiated. This prevents conflicts from rapid user interactions.

### Default Behavior Example:

```html
<!-- Clicking this button multiple times will cancel previous requests (default behavior) -->
<button data-on:click="@get('/slow-endpoint')">Load Data</button>
```

This cancellation is element-specific; requests on different elements can run concurrently.

### Controlling Request Cancellation:

You can modify this behavior using the `requestCancellation` option:

- **Allow Concurrent Requests:**
```html
<!-- Allow concurrent requests (no automatic cancellation) -->
<button data-on:click="@get('/endpoint', {requestCancellation: 'disabled'})">Allow Multiple</button>
```

- **Custom Abort Controller:**
```html
<div data-signals:controller="new AbortController()">
    <button data-on:click="@get('/endpoint', {requestCancellation: $controller})">Start Request</button>
    <button data-on:click="$controller.abort()">Cancel Request</button>
</div>
```
```

--------------------------------

### C# SDK Example

Source: https://data-star.dev/how_tos/load_more_list_items

Example implementation in C# using the Datastar SDK for ASP.NET Core to manage the 'load more' functionality.

```APIDOC
## C# SDK Example

### Description
This C# code demonstrates a minimal ASP.NET Core application using Datastar. The `/more` endpoint reads signals, patches new elements with append mode, and updates signals or removes the button based on the offset.

### Method
GET

### Endpoint
`/more`

### Parameters
#### Query Parameters
N/A

#### Request Body
N/A (Signals are read from request context)

### Request Example
```csharp
using System.Text.Json;
using StarFederation.Datastar;
using StarFederation.Datastar.DependencyInjection;

public class Program
{
    public record OffsetSignals(int offset);

    public static void Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);
        builder.Services.AddDatastar();
        var app = builder.Build();

        app.MapGet("/more", async (IDatastarService datastarService) =>
        {
            var max = 5;
            var limit = 1;
            var signals = await datastarService.ReadSignalsAsync<OffsetSignals>();
            var offset = signals.offset;
            if (offset < max)
            {
                var newOffset = offset + limit;
                await datastarService.PatchElementsAsync($"<div>Item {newOffset}</div>", new()
                {
                    Selector = "#list",
                    PatchMode = PatchElementsMode.Append,
                });
                if (newOffset < max)
                    await datastarService.PatchSignalsAsync(new OffsetSignals(newOffset));
                else
                    await datastarService.RemoveElementAsync("#load-more");
            }
        });

        app.Run();
    }
}
```

### Response
#### Success Response (200)
HTTP 200 OK with response body handled by Datastar client-side logic based on signals and patched elements.

#### Response Example
N/A (Response is typically SSE or handled via Datastar's client-side state management)
```

--------------------------------

### Reading Signals in PHP

Source: https://data-star.dev/docs

Demonstrates how to read all signals from an incoming request in PHP using the `ServerSentEventGenerator::readSignals()` helper function provided by Datastar.

```php
1use starfederation\datastar\ServerSentEventGenerator;
2
3// Reads all signals from the request.
4$signals = ServerSentEventGenerator::readSignals();
```

--------------------------------

### Generate SSE for Patching Elements (Java)

Source: https://data-star.dev/guide/getting_started

This Java snippet demonstrates generating Server-Sent Events (SSE) to patch elements into the DOM using Datastar's Java library. It initializes an SSE generator and then sends an initial patch, followed by an update after a one-second delay, facilitating dynamic content updates.

```java
import starfederation.datastar.utils.ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);

// Patches elements into the DOM.
generator.send(PatchElements.builder()
    .data("<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
    .build()
);

Thread.sleep(1000);

generator.send(PatchElements.builder()
    .data("<div id=\"hal\">Waiting for an order...</div>")
    .build()
);
```

--------------------------------

### Datastar Declarative Signals and Event Handling (HTML)

Source: https://data-star.dev/docs

This snippet demonstrates Datastar's declarative approach to managing frontend state and user interactions. It uses `data-signals` to initialize reactive variables, `data-on:click` to trigger actions, and `data-text` to display signal values, enabling dynamic UI updates based on user input.

```html
<div data-signals:hal="'...'">
    <button data-on:click="$hal = 'Affirmative, Dave. I read you.'">
        HAL, do you read me?
    </button>
    <div data-text="$hal"></div>
</div>
```

--------------------------------

### Install Datastar using CDN

Source: https://data-star.dev/guide

Include Datastar in your project by adding this script tag to your HTML. This fetches the latest version from the CDN.

```html
<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js"></script>
```

--------------------------------

### Ruby SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Ruby code demonstrates how to use the Datastar gem to stream SSE events within a Rack handler. It shows how to instantiate `Datastar` and use the `stream` method to send patched elements and signals.

```ruby
1require 'datastar'
2
3# Create a Datastar::Dispatcher instance
4
datastar = Datastar.new(request:, response:)
5
6# In a Rack handler, you can instantiate from the Rack env
7# datastar = Datastar.from_rack_env(env)
8
9# Start a streaming response
10datastar.stream do |sse|
11  # Patches elements into the DOM
12  sse.patch_elements %(<div id="question">What do you put in a toaster?</div>)
13
14  # Patches signals
15  sse.patch_signals(response: '', answer: 'bread')
16end
```

--------------------------------

### Patch Elements using Server-Sent Events (SSE)

Source: https://data-star.dev/guide/getting_started

This code illustrates how to send Server-Sent Events (SSE) to patch multiple HTML elements into the DOM. It shows how to send an initial response and then update it after a delay, demonstrating real-time content updates. The `datastar-patch-elements` event type is used for this purpose.

```sse
event: datastar-patch-elements
data: elements <div id="hal">
    I’m sorry, Dave. I’m afraid I can’t do that.
</div>

event: datastar-patch-elements
data: elements <div id="hal">
    Waiting for an order...
</div>


```

--------------------------------

### Reading Nested Signals in Go

Source: https://data-star.dev/docs

Illustrates how to read nested signals in Go using the `datastar-go` library. It defines a `Signals` struct with a nested `Foo` struct and uses `datastar.ReadSignals` to parse the incoming request.

```go
 1import ("github.com/starfederation/datastar-go/datastar")
 2
 3type Signals struct {
 4    Foo struct {
 5        Bar string `json:"bar"`
 6    } `json:"foo"`
 7}
 8
 9signals := &Signals{}
10if err := datastar.ReadSignals(request, signals); err != nil {
11    http.Error(w, err.Error(), http.StatusBadRequest)
12    return
13}
```

--------------------------------

### Java: Generating SSE Events for Datastar DOM Patching (JSP)

Source: https://data-star.dev/docs

This Java code example demonstrates using Datastar's `ServerSentEventGenerator` within a JSP context to send SSE events for patching elements. It shows how to send an initial HTML update and then a subsequent update after a delay.

```java
1import starfederation.datastar.utils.ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
6
7// Patches elements into the DOM.
8generator.send(PatchElements.builder()
9    .data("<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
10    .build()
11);
12
13Thread.sleep(1000);
14
15generator.send(PatchElements.builder()
16    .data("<div id=\"hal\">Waiting for an order...</div>")
17    .build()
18);

```

--------------------------------

### Datastar: Patching Elements with HTML Content-Type

Source: https://data-star.dev/docs

This example shows how Datastar patches HTML elements into the DOM when the backend response has a 'text/html' content-type. The morphing strategy updates existing DOM elements based on their IDs, ensuring efficient rendering.

```html
1<div id="hal">
2    I’m sorry, Dave. I’m afraid I can’t do that.
3</div>
```

--------------------------------

### Generate SSE for Patching Elements (Go)

Source: https://data-star.dev/guide/getting_started

This Go snippet shows how to generate Server-Sent Events (SSE) for patching elements into the DOM using Datastar's Go SDK. It creates an SSE generator, sends an initial HTML response, waits for a second, and then sends an updated response. This enables real-time UI modifications.

```go
import (
	"github.com/starfederation/datastar-go/datastar"
	time
)

// Creates a new `ServerSentEventGenerator` instance.
sse := datastar.NewSSE(w,r)

// Patches elements into the DOM.
sse.PatchElements(
    `<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`
)

time.Sleep(1 * time.Second)

sse.PatchElements(
    `<div id="hal">Waiting for an order...</div>`
)
```

--------------------------------

### C# SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This C# code demonstrates how to add Datastar as a service and handle SSE requests using `MapGet`. It shows asynchronous patching of DOM elements and signals within an ASP.NET Core application.

```csharp
1using StarFederation.Datastar.DependencyInjection;
2
3// Adds Datastar as a service
4builder.Services.AddDatastar();
5
6app.MapGet("/", async (IDatastarService datastarService) =>
7{
8    // Patches elements into the DOM.
9    await datastarService.PatchElementsAsync(@"<div id=\"question\">What do you put in a toaster?</div>");
10
11    // Patches signals.
12    await datastarService.PatchSignalsAsync(new { response = "", answer = "bread" });
13});
```

--------------------------------

### Server-Sent Events (SSE) for Appending Script to Body

Source: https://data-star.dev/docs

This SSE example demonstrates how to use Datastar's 'append' mode to insert a script tag directly into the document's body. This is useful for executing scripts that don't require specific element patching.

```text
event: datastar-patch-elements
data: mode append
data: selector body
data: elements <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>

```

--------------------------------

### GET /redirect (Go with Redirect Helper)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to initiate a redirect after patching elements using Server-Sent Events and a redirect helper.

```APIDOC
## GET /redirect

### Description
Handles a redirect request using Server-Sent Events. It patches elements to show a redirect message, waits for 3 seconds, and then calls a `Redirect` method to navigate the client to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```go
sse := datastar.NewSSE(w, r)
sse.PatchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`)
time.Sleep(3 * time.Second)
sse.Redirect("/guide")
```

### Response
#### Success Response (200)
A successful response initiates the SSE stream for UI updates and redirection.

#### Response Example
(No specific JSON response, SSE stream is used)
```

--------------------------------

### GET /redirect (C# with Redirect Helper)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to initiate a redirect using a helper method after patching elements.

```APIDOC
## GET /redirect

### Description
Handles a GET request to redirect the client. It first updates the UI by patching elements to display a message, then uses a `Redirect` helper method to navigate the client to the specified URL after a short delay.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```csharp
await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
await Task.Delay(TimeSpan.FromSeconds(3));
await datastarService.Redirect("/guide");
```

### Response
#### Success Response (200)
Returns an HTTP 200 OK response, and the client will be redirected after a delay.

#### Response Example
(No specific JSON response body, the client is redirected)
```

--------------------------------

### GET /redirect (C#)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to initiate a redirect after patching elements and executing a script.

```APIDOC
## GET /redirect

### Description
Initiates a redirect to a new URL after updating the client's UI. It first patches specific HTML elements to display a message, waits for a few seconds, and then executes a script to perform the redirect.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```csharp
await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
await Task.Delay(TimeSpan.FromSeconds(3));
await datastarService.ExecuteScriptAsync("setTimeout(() => window.location = \"/guide\");");
```

### Response
#### Success Response (200)
Returns an HTTP 200 OK response, and the client will be redirected after a delay.

#### Response Example
(No specific JSON response body, the client is redirected)
```

--------------------------------

### Display Initial State with Data Signals

Source: https://data-star.dev/examples/on_signal_patch

This HTML snippet demonstrates the initial setup of a Data-Star application, defining initial signal values for counter and message directly within the `data-signals` attribute. It also includes buttons to trigger updates and displays for current values.

```html
<div data-signals="{counter: 0, message: 'Hello World', allChanges: [], counterChanges: []}">
    <div class="actions">
        <button data-on:click="$message = `Updated: ${performance.now().toFixed(2)}`">
            Update Message
        </button>
        <button data-on:click="$counter++">
            Increment Counter
        </button>
        <button
            class="error"
            data-on:click="$allChanges.length = 0; $counterChanges.length = 0"
        >
            Clear All Changes
        </button>
    </div>
    <div>
        <h3>Current Values</h3>
        <p>Counter: <span data-text="$counter"></span></p>
        <p>Message: <span data-text="$message"></span></p>
    </div>
    <div
        data-on-signal-patch="$counterChanges.push(patch)"
        data-on-signal-patch-filter="{include: /^counter$/}"
    >
        <h3>Counter Changes Only</h3>
        <pre data-json-signals__terse="{include: /^counterChanges/}"></pre>
    </div>
    <div
        data-on-signal-patch="$allChanges.push(patch)"
        data-on-signal-patch-filter="{exclude: /allChanges|counterChanges/}"
    >
        <h3>All Signal Changes</h3>
        <pre data-json-signals__terse="{include: /^allChanges/}"></pre>
    </div>
</div>
```

--------------------------------

### Rust SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Rust code snippet shows how to generate SSE events using Datastar's prelude and the `async_stream` crate. It demonstrates yielding `PatchElements` and `PatchSignals` within a stream.

```rust
1use datastar::prelude::*;
2use async_stream::stream;
3
4Sse(stream! {
5    // Patches elements into the DOM.
6    yield PatchElements::new("<div id='question'>What do you put in a toaster?</div>").into();
7
8    // Patches signals.
9    yield PatchSignals::new("{response: '', answer: 'bread'}").into();
10})
```

--------------------------------

### PHP SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This PHP code demonstrates using Datastar's `ServerSentEventGenerator` to send SSE events, including patching DOM elements and signals. It shows how to initialize the generator and send data.

```php
1use starfederation\datastar\ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4$sse = new ServerSentEventGenerator();
5
6// Patches elements into the DOM.
7$sse->patchElements(
8    '<div id="question">What do you put in a toaster?</div>'
9);
10
11// Patches signals.
12$sse->patchSignals(['response' => '', 'answer' => 'bread']);
```

--------------------------------

### Implement SSE Stream with Rust

Source: https://data-star.dev/docs

This Rust code snippet demonstrates how to create a Server-Sent Events (SSE) stream using the `async_stream` crate. It yields patch signals at intervals, simulating real-time updates.

```rust
use async_stream::stream;
use datastar::prelude::*;
use std::thread;
use std::time::Duration;

Sse(stream! {
    // Patches signals.
    yield PatchSignals::new("{hal: 'Affirmative, Dave. I read you.'}").into();

    thread::sleep(Duration::from_secs(1));
    
    yield PatchSignals::new("{hal: '...'} ").into();
})
```

--------------------------------

### Kotlin SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Kotlin snippet shows how to instantiate and use the `ServerSentEventGenerator` to patch elements and signals. It provides a concise way to handle SSE responses in Kotlin applications.

```kotlin
1val generator = ServerSentEventGenerator(response)
2
generator.patchElements(
3    elements = """<div id=\"question\">What do you put in a toaster?</div>""",
4)
5
generator.patchSignals(
6    signals = "{"response": "", "answer": "bread"}",
7)
```

--------------------------------

### data-json-signals for Reactive JSON Output

Source: https://data-star.dev/docs

Shows how `data-json-signals` can be used to display a reactive, JSON stringified version of signals within an element's text content, primarily for debugging purposes.

```html
<!-- Display all signals -->
<pre data-json-signals></pre>
```

```html
<!-- Only show signals that include "user" in their path -->
<pre data-json-signals="{include: /user/}"></pre>

<!-- Show all signals except those ending with "temp" -->
<pre data-json-signals="{exclude: /temp$/}"></pre>

<!-- Combine include and exclude filters -->
<pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>
```

--------------------------------

### Generate SSE for Patching Elements (Clojure)

Source: https://data-star.dev/guide/getting_started

This Clojure snippet shows how to generate Server-Sent Events (SSE) to patch elements into the DOM using Datastar's SDK. It demonstrates sending an initial response with a message and then updating it after a short delay. This is typically used within a Ring handler.

```clojure
;; Import the SDK's api and your adapter
(require
 '[starfederation.datastar.clojure.api :as d*]
 '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])

;; in a ring handler
(defn handler [request]
 ;; Create an SSE response
 (->sse-response request
                 {on-open
                  (fn [sse]
                    ;; Patches elements into the DOM
                    (d*/patch-elements! sse
                                        "<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
                    (Thread/sleep 1000)

                    (d*/patch-elements! sse
                                        "<div id=\"hal\">Waiting for an order...</div>"))})
```

--------------------------------

### GET /redirect (Go)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to initiate a redirect after patching elements and executing a script using Server-Sent Events.

```APIDOC
## GET /redirect

### Description
Handles a redirect request by first patching elements on the client-side to show a "Redirecting..." message. After a 3-second delay, it executes a JavaScript script to navigate the user to the "/guide" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```go
sse := datastar.NewSSE(w, r)
sse.PatchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`)
time.Sleep(3 * time.Second)
sse.ExecuteScript(`
    setTimeout(() => window.location = "/guide")
`)
```

### Response
#### Success Response (200)
A successful response initiates the SSE stream for UI updates and redirection.

#### Response Example
(No specific JSON response, SSE stream is used)
```

--------------------------------

### Go SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Go code snippet shows how to create a `ServerSentEventGenerator` instance and use it to patch elements and signals into the DOM. It's designed for handling SSE requests within a Go web server.

```go
1import ("github.com/starfederation/datastar-go/datastar")
2
3// Creates a new `ServerSentEventGenerator` instance.
4sse := datastar.NewSSE(w,r)
5
6// Patches elements into the DOM.
7sse.PatchElements(
8    `<div id="question">What do you put in a toaster?</div>`
9)
10
11// Patches signals.
12sse.PatchSignals([]byte(`{response: '', answer: 'bread'}`))
```

--------------------------------

### SSE Event: datastar-patch-elements

Source: https://data-star.dev/docs

Handles patching of DOM elements using Server-Sent Events. It supports various morphing modes and options for selecting and updating elements.

```APIDOC
## SSE Event: datastar-patch-elements

### Description
Patches one or more elements in the DOM. By default, Datastar morphs elements by matching top-level elements based on their ID. IDs should be placed on top-level elements for morphing and on elements within them to preserve state.

### Event Type
`datastar-patch-elements`

### Data Directives
- **`mode <mode>`**: Specifies how to morph elements. Supported modes include `outer` (default), `inner`, `replace`, `prepend`, `append`, `before`, `after`, `remove`.
- **`selector <css_selector>`**: Selects the target element(s) for patching using a CSS selector. Not required for `outer` or `replace` modes.
- **`useViewTransition <boolean>`**: Enables view transitions during patching. Defaults to `false`.
- **`elements <html_content>`**: The HTML content to patch into the DOM.

### Request Example (Patching Element)
```
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>
```

### Request Example (Removing Element)
```
event: datastar-patch-elements
data: mode remove
data: selector #foo
```

### Request Example (Multi-line Elements with Options)
```
event: datastar-patch-elements
data: mode inner
data: selector #foo
data: useViewTransition true
data: elements <div>
data: elements        Hello world!
data: elements </div>
```

### Response
N/A (This is a client-side event triggered by the server).
```

--------------------------------

### Java SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Java code illustrates how to use Datastar's `ServerSentEventGenerator` to send SSE events, specifically patching DOM elements and signals. It utilizes an `HttpServletResponseAdapter` for integration with Java Servlets.

```java
1import starfederation.datastar.utils.ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
6
7// Patches elements into the DOM.
8generator.send(PatchElements.builder()
9    .data("<div id=\"question\">What do you put in a toaster?</div>")
10    .build()
11);
12
13// Patches signals.
14generator.send(PatchSignals.builder()
15    .data("{\"response\": \"\", \"answer\": \"\"}")
16    .build()
17);
```

--------------------------------

### GET /redirect (Kotlin with Redirect Helper)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating the UI with a message using a redirect helper.

```APIDOC
## GET /redirect

### Description
This endpoint handles a redirect by first updating the client's view with an indicator message, pausing execution for 3 seconds, and then calling a `redirect` method to navigate the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent(),
)

Thread.sleep(3 * ONE_SECOND)

generator.redirect("/guide")
```

### Response
#### Success Response (200)
Upon successful execution, the client UI is updated, and a redirection is initiated.

#### Response Example
(No JSON response body; client-side actions are performed)
```

--------------------------------

### Clojure SSE Event Handling with Datastar

Source: https://data-star.dev/docs

This Clojure code snippet demonstrates how to create a Server-Sent Events (SSE) response using Datastar's http-kit adapter. It shows how to set up a Ring handler that sends SSE events, including patching DOM elements and signals.

```clojure
1;; Import the SDK's api and your adapter
2(require
3 '[starfederation.datastar.clojure.api :as d*]
4 '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
5
6;; in a ring handler
7(defn handler [request]
8  ;; Create an SSE response
9  (->sse-response request
10                  {
11                   (fn [sse] 
12                     ;; Patches elements into the DOM
13                     (d*/patch-elements! sse
14                                         "<div id=\"question\">What do you put in a toaster?</div>")
15
16                     ;; Patches signals
17                     (d*/patch-signals! sse "{response: '', answer: 'bread'}"))}))
```

--------------------------------

### GET /redirect (Kotlin)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating the UI with a message and executing a script.

```APIDOC
## GET /redirect

### Description
This endpoint handles a redirect by first updating the client's view with an indicator message, pausing execution for 3 seconds, and then executing a JavaScript script to redirect the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent(),
)

Thread.sleep(3 * ONE_SECOND)

generator.executeScript(
    script = "setTimeout(() => window.location = '/guide')",
)
```

### Response
#### Success Response (200)
Upon successful execution, the client UI is updated, and a redirection script is sent.

#### Response Example
(No JSON response body; client-side actions are performed)
```

--------------------------------

### GET /redirect (Ruby)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to initiate a redirect after patching elements and executing a script via Server-Sent Events.

```APIDOC
## GET /redirect

### Description
Handles a redirect request by sending Server-Sent Events. It first updates the client's UI with a message indicating redirection. After a 3-second delay, it executes a JavaScript command to navigate the client to the \"/guide\" path.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'

  sleep 3

  sse.execute_script <<~JS
    setTimeout(() => {
      window.location = '/guide'
    })
  JS
end
```

### Response
#### Success Response (200)
Successful execution sends SSE events to the client for UI update and redirection.

#### Response Example
(No specific JSON response; SSE stream is used)
```

--------------------------------

### Generic SDK: Backend Redirect with Stream

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This example demonstrates a backend redirect using a generic Datastar SDK, employing a streaming approach to patch elements and then initiate a redirect. It assumes a Datastar object is initialized with request and response objects.

```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'

  sleep 3

  sse.redirect '/guide'
end
```

--------------------------------

### Set All Signals with @setAll()

Source: https://data-star.dev/docs

Sets the value of all matching signals using a regular expression filter. Supports including and excluding patterns. Useful for bulk updates to signal values.

```html
<div data-signals:foo="false">
    <button data-on:click="@setAll(true, {include: /^foo$/})"></button>
</div>

<div data-signals="{user: {name: '', nickname: ''}}">
    <button data-on:click="@setAll('johnny', {include: /^user\./})"></button>
</div>

<div data-signals="{data: '', data_temp: '', info: '', info_temp: ''}">
    <button data-on:click="@setAll('reset', {include: /.*/, exclude: /_temp$/})"></button>
</div>
```

--------------------------------

### Reading Nested Signals in C# (ASP.NET Core)

Source: https://data-star.dev/docs

Shows how to read nested signals from an incoming request in C# using Datastar's `IDatastarService`. It defines a `Signals` record with nested properties and retrieves the 'bar' signal.

```csharp
 1using StarFederation.Datastar.DependencyInjection;
 2
 3// Adds Datastar as a service
 4builder.Services.AddDatastar();
 5
 6public record Signals
 7{
 8    [JsonPropertyName("foo")] [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
 9    public FooSignals? Foo { get; set; } = null;
10
11    public record FooSignals
12    {
13        [JsonPropertyName("bar")] [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
14        public string? Bar { get; set; }
15    }
16}
17
18app.MapGet("/read-signals", async (IDatastarService datastarService) =>
19{
20    Signals? mySignals = await datastarService.ReadSignalsAsync<Signals>();
21    var bar = mySignals?.Foo?.Bar;
22});
```

--------------------------------

### HTML for Templ Counter Interface

Source: https://data-star.dev/examples/templ_counter

This HTML code sets up the user interface for the templ counter. It includes two buttons, one for a global counter and another for a user-specific counter, with styles for display and click event handling.

```html
<div
    style="display: flex; gap: var(--size-6)"
    data-init="@get('/examples/templ_counter/updates')"
>
    <!-- Global Counter -->
    <button
        id="global"
        class="info"
        data-on:click="@patch('/examples/templ_counter/global')"
    >
        Global Clicks: 0
    </button>

    <!-- User Counter -->
    <button
        id="user"
        class="success"
        data-on:click="@patch('/examples/templ_counter/user')"
    >
        User Clicks: 0
    </button>
</div>
```

--------------------------------

### GET /redirect (PHP)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating the UI with a message and executing a script.

```APIDOC
## GET /redirect

### Description
Handles a redirect request by sending server-sent events. It first updates the client's HTML with an indicator, waits for 3 seconds, and then executes a JavaScript command to redirect the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```php
$sse = new ServerSentEventGenerator();
$sse->patchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`);
sleep(3);
$sse->executeScript(`
    setTimeout(() => window.location = "/guide")
`);
```

### Response
#### Success Response (200)
Returns a successful response, and the client will receive SSE updates for redirection.

#### Response Example
(No JSON response body; SSE stream is used)
```

--------------------------------

### data-indicator with Casing Modifiers

Source: https://data-star.dev/docs

Shows how the `__case` modifier can be applied to the `data-indicator` attribute to control the casing of the generated indicator signal name, supporting camel, kebab, snake, and pascal cases.

```html
<button data-on:click="@get('/endpoint')"
        data-indicator:fetching
></button>
```

--------------------------------

### Datastar Default Request Cancellation

Source: https://data-star.dev/docs

Illustrates the default behavior of Datastar where clicking a button multiple times will cancel previous requests on the same element. This is useful for preventing conflicting states when the backend action is slow.

```html
<!-- Clicking this button multiple times will cancel previous requests (default behavior) -->
<button data-on:click="@get('/slow-endpoint')">Load Data</button>
```

--------------------------------

### GET /redirect (Python)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating the UI and executing a script using Server-Sent Events.

```APIDOC
## GET /redirect

### Description
Handles a redirect after a delay. It first sends a Server-Sent Event to update the client's UI with a redirect indicator. After a 3-second pause, it sends another event to execute a JavaScript snippet that redirects the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```python
@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.execute_script('setTimeout(() => window.location = "/guide")')
```

### Response
#### Success Response (200)
The client receives SSE events to update its UI and trigger a redirection.

#### Response Example
(No JSON response; SSE events are sent to the client)
```

--------------------------------

### Datastar Expression: Using JavaScript Operators

Source: https://data-star.dev/docs

This HTML demonstrates using JavaScript operators like the ternary operator `?:`, logical OR `||`, and logical AND `&&` within Datastar expressions for conditional rendering and actions.

```html
<!-- Output one of two values, depending on the truthiness of a signal -->
<div data-text="$landingGearRetracted ? 'Ready' : 'Waiting'"></div>

<!-- Show a countdown if the signal is truthy or the time remaining is less than 10 seconds -->
<div data-show="$landingGearRetracted || $timeRemaining < 10">
    Countdown
</div>

<!-- Only send a request if the signal is truthy -->
<button data-on:click="$landingGearRetracted && @post('/launch')">
    Launch
</button>
```

--------------------------------

### JavaScript for Synchronous Function in External Script

Source: https://data-star.dev/docs

A simple synchronous JavaScript function that takes data as an argument and returns a formatted string. This function can be called from Datastar HTML using data attributes.

```javascript
function myfunction(data) {
    return `You entered: ${data}`;
}
```

--------------------------------

### File Uploads with data-bind

Source: https://data-star.dev/docs

Handles file uploads using input fields of type 'file' with the `data-bind` attribute. File contents are automatically base64 encoded, eliminating the need for a form. The resulting signal format is an array of objects containing name, contents, and mime type. For server uploads, use a traditional form with `multipart/form-data`.

```html
<input type="file" data-bind:files multiple />
```

--------------------------------

### Datastar Expression: Signal Property Access

Source: https://data-star.dev/docs

This HTML snippet illustrates accessing a property of a signal (e.g., 'length' for a string signal) using Datastar expressions within a `data-text` attribute.

```html
<div data-text="$foo.length"></div>
```

--------------------------------

### Clojure - Redirect with Datastar API

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Provides a Clojure example for handling redirects using Datastar's API. It patches elements, pauses, and then redirects the client to a new URL.

```clojure
(require
  '[starfederation.datastar.clojure.api :as d*]
  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]]
  '[some.hiccup.library :refer [html]])


(defn handler [ring-request]
  (->sse-response ring-request
    {on-open
      (fn [sse]
        (d*/patch-elements! sse
          (html [:div#indicator "Redirecting in 3 seconds..."]))
        (Thread/sleep 3000)
        (d*/redirect! sse "/guide")
        (d*/close-sse! sse))}))

```

--------------------------------

### CQRS Pattern for Resilient SSE Updates

Source: https://data-star.dev/how_tos/prevent_sse_connections_closing

This example illustrates a CQRS (Command Query Responsibility Segregation) approach for managing SSE updates. It shows how to initialize a main content area with a GET request and suggests sending the complete state with each update to ensure resilience against connection interruptions.

```html
<div data-init="@get('/cqrs_endpoint')"></div>
<div id="main">
    ...
</div>
```

--------------------------------

### Clojure: Generating SSE Events for Datastar Patching

Source: https://data-star.dev/docs

This Clojure code snippet shows how to generate Server-Sent Events (SSE) for patching elements into the DOM using Datastar's SDK. It demonstrates sending an initial message and then updating it after a short delay.

```clojure
1;; Import the SDK's api and your adapter
2(require
3 '[starfederation.datastar.clojure.api :as d*]
4 '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
5
6;; in a ring handler
7(defn handler [request]
8  ;; Create an SSE response
9  (->sse-response request
10                  {on-open
11                   (fn [sse] 
12                     ;; Patches elements into the DOM
13                     (d*/patch-elements! sse
14                                         "<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
15                     (Thread/sleep 1000)
16                     (d*/patch-elements! sse
17                                         "<div id=\"hal\">Waiting for an order...</div>"))}))

```

--------------------------------

### Filter Signals for data-on-signal-patch with DataStar

Source: https://data-star.dev/docs

The `data-on-signal-patch-filter` attribute refines which signals trigger the `data-on-signal-patch` handler. It accepts an object with `include` and/or `exclude` properties, which are regular expressions. This allows for precise control over signal reactivity.

```html
<!-- Only react to counter signal changes -->
<div data-on-signal-patch-filter="{include: /^counter$/}"></div>
```

```html
<!-- React to all changes except those ending with "changes" -->
<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>
```

```html
<!-- Combine include and exclude filters -->
<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
```

--------------------------------

### Go: Generating SSE Events for Datastar DOM Patching

Source: https://data-star.dev/docs

This Go code snippet shows how to create a Server-Sent Event (SSE) generator using Datastar's Go SDK. It then uses the `PatchElements` method to send HTML content to update the DOM, including a timed update.

```go
1import (
2    "github.com/starfederation/datastar-go/datastar"
3    time
4)
5
6// Creates a new `ServerSentEventGenerator` instance.
7sse := datastar.NewSSE(w,r)
8
9// Patches elements into the DOM.
10sse.PatchElements(
11    `<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`
12)
13
14time.Sleep(1 * time.Second)
15
16sse.PatchElements(
17    `<div id="hal">Waiting for an order...</div>`
18)

```

--------------------------------

### Kotlin: Generating SSE Events for Datastar DOM Patching

Source: https://data-star.dev/docs

This Kotlin code snippet illustrates generating Server-Sent Events (SSE) for Datastar's DOM patching functionality. It shows sending an initial HTML update to the DOM and then a subsequent update after a one-second delay.

```kotlin
1val generator = ServerSentEventGenerator(response)
2
3generator.patchElements(
4    elements = """<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>""",
5)
6
7Thread.sleep(ONE_SECOND)
8

```

--------------------------------

### Basic Repeated Datastar Action

Source: https://data-star.dev/how_tos/keep_datastar_code_dry

This snippet shows a basic example of repeating a Datastar backend action (`@get('/endpoint')`) across multiple buttons. It highlights the problem of repetition that the following solutions aim to address.

```html
<button data-on:click="@get('/endpoint')">Click me</button>
<button data-on:click="@get('/endpoint')">No, click me!</button>
<button data-on:click="@get('/endpoint')">Click us all!</button>
```

--------------------------------

### GET /redirect (Rust)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating UI elements and executing a script using Server-Sent Events.

```APIDOC
## GET /redirect

### Description
Handles a GET request to perform a client-side redirect. It first sends a Server-Sent Event to update the HTML content, then waits for 3 seconds, and finally sends another event to execute JavaScript that redirects the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```rust
use datastar::prelude::*;
use async_stream::stream;
use core::time::Duration;

Sse(stream! {
    yield PatchElements::new("<div id='indicator'>Redirecting in 3 seconds...</div>").into();
    tokio::time::sleep(core::time::Duration::from_secs(3)).await;
    yield ExecuteScript::new("setTimeout(() => window.location = '/guide')").into();
});
```

### Response
#### Success Response (200)
Returns a successful response, initiating an SSE stream that updates the client and redirects it.

#### Response Example
(No JSON response; SSE stream is utilized)
```

--------------------------------

### Datastar Expression: Multi-line Statements

Source: https://data-star.dev/docs

This HTML snippet illustrates how Datastar expressions can span multiple lines. A semicolon is required to separate statements, as line breaks alone are not sufficient.

```html
<div data-signals:foo="1">
    <button data-on:click="
        $landingGearRetracted = true; 
        @post('/launch')
    ">
        Force launch
    </button>
</div>
```

--------------------------------

### Conditional Element Visibility with data-show

Source: https://data-star.dev/docs

The `data-show` attribute controls element visibility based on an expression's truthiness. An optional `style='display: none;'` can prevent content flash during initial rendering.

```html
<input data-bind:foo />
<button data-show="$foo != ''">Save</button>
```

```html
<input data-bind:foo />
<button data-show="$foo != ''" style="display: none;">
    Save
</button>
```

--------------------------------

### Datastar Click Event to Backend Endpoint

Source: https://data-star.dev/index

This code snippet demonstrates how to attach a click event listener to a button. When clicked, it triggers a GET request to the `/endpoint` using Datastar's `@get()` helper, allowing backend interaction with minimal user-side JavaScript.

```html
<button data-on:click="@get('/endpoint')">
    Open the pod bay doors, HAL.
</button>

<div id="hal">Waiting for an order...</div>
```

--------------------------------

### Compact JSON Output with `__terse` Modifier

Source: https://data-star.dev/docs

The `__terse` modifier for `data-json-signals` outputs JSON in a compact format without extra whitespace. This is useful for displaying filtered data inline efficiently.

```html
<pre data-json-signals__terse="{include: /counter/}"></pre>
```

--------------------------------

### Patch Signals with Datastar SSE

Source: https://data-star.dev/docs

Patches signals into existing signals on a page using the `datastar-patch-signals` SSE event. It allows updating signals with new values, optionally only if they don't exist (`onlyIfMissing`). Signals can be removed by setting their values to `null`. The `signals` line must contain a valid `data-signals` attribute.

```html
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}


```

```html
event: datastar-patch-signals
data: signals {foo: null, bar: null}


```

```html
event: datastar-patch-signals
data: onlyIfMissing true
data: signals {foo: 1, bar: 2}


```

--------------------------------

### HTML Table for Bulk Update

Source: https://data-star.dev/examples/bulk_update

This HTML structure sets up a table with selectable rows for bulk operations. It includes checkboxes for individual row selection and a master checkbox for selecting all rows. Buttons for 'Activate' and 'Deactivate' trigger `PUT` requests to the server, with visual feedback for ongoing operations.

```html
<div
    id="demo"
    data-signals__ifmissing="{_fetching: false, selections: Array(4).fill(false)}"
>
    <table>
        <thead>
            <tr>
                <th>
                    <input
                        type="checkbox"
                        data-bind:_all
                        data-on:change="$selections = Array(4).fill($_all)"
                        data-effect="$selections; $_all = $selections.every(Boolean)"
                        data-attr:disabled="$_fetching"
                    />
                </th>
                <th>Name</th>
                <th>Email</th>
                <th>Status</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>
                    <input
                        type="checkbox"
                        data-bind:selections
                        data-attr:disabled="$_fetching"
                    />
                </td>
                <td>Joe Smith</td>
                <td>joe@smith.org</td>
                <td>Active</td>
            </tr>
            <!-- More rows... -->
        </tbody>
    </table>
    <div role="group">
        <button
            class="success"
            data-on:click="@put('/examples/bulk_update/activate')"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
        >
            <i class="pixelarticons:user-plus"></i>
            Activate
        </button>
        <button
            class="error"
            data-on:click="@put('/examples/bulk_update/deactivate')"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
        >
            <i class="pixelarticons:user-x"></i>
            Deactivate
        </button>
    </div>
</div>
```

--------------------------------

### Implement Click-to-Load Button

Source: https://data-star.dev/examples/click_to_load

This HTML snippet defines a button that, when clicked, triggers the loading of more data. It utilizes data attributes for state management (fetching status) and event handling for the click action. The button dynamically updates its ARIA attribute based on the fetching state. It relies on a server endpoint '/examples/click_to_load/more' to provide the next set of data.

```html
<button
    class="info wide"
    data-indicator:_fetching
    data-attr:aria-disabled="`${$_fetching}`"
    data-on:click="!$_fetching && @get('/examples/click_to_load/more')"
>
    Load More
</button>
```

--------------------------------

### Toggle All Signals with @toggleAll()

Source: https://data-star.dev/docs

Toggles the boolean value of all matching signals using a regular expression filter. Supports including and excluding patterns. Ideal for managing boolean states across multiple signals.

```html
<div data-signals:foo="false">
    <button data-on:click="@toggleAll({include: /^foo$/})"></button>
</div>

<div data-signals="{isOpen: false, isActive: true, isEnabled: false}">
    <button data-on:click="@toggleAll({include: /^is/})"></button>
</div>

<div data-signals="{settings: {darkMode: false, autoSave: true}}">
    <button data-on:click="@toggleAll({include: /^settings\./})"></button>
</div>
```

--------------------------------

### Integrate data-signals and data-on for Reactivity

Source: https://data-star.dev/guide/reactive_signals

This example demonstrates frontend reactivity by combining `data-signals` for initial state, `data-on` for user interaction, and `data-text` to display signal values. It includes a button to update a signal and a div to show its current value.

```html
<div data-signals:hal="'...' '>
    <button data-on:click="$hal = 'Affirmative, Dave. I read you.'">
        HAL, do you read me?
    </button>
    <div data-text="$hal"></div>
</div>
```

--------------------------------

### Datastar Expression: Display Signal Value

Source: https://data-star.dev/docs

This HTML snippet demonstrates how to display the value of a Datastar signal named 'foo' using a `data-text` attribute. The signal's initial value is set to '1'.

```html
<div data-signals:foo="1">
    <div data-text="$foo"></div>
</div>
```

--------------------------------

### Datastar Expression: Access Element ID

Source: https://data-star.dev/docs

This HTML snippet shows how to access the `id` attribute of the current element using the `el` variable within a Datastar expression in a `data-text` attribute.

```html
<div id="foo" data-text="el.id"></div>
```

--------------------------------

### JavaScript for Asynchronous Function in External Script

Source: https://data-star.dev/docs

An asynchronous JavaScript function that simulates a delay using `setTimeout` and then dispatches a custom event containing the result. This is the pattern to use when dealing with promises or other non-blocking operations in Datastar external scripts.

```javascript
async function myfunction(element, data) {
    const value = await new Promise((resolve) => {
        setTimeout(() => resolve(`You entered: ${data}`), 1000);
    });
    element.dispatchEvent(
        new CustomEvent('mycustomevent', {detail: {value}})
    );
}
```

--------------------------------

### JavaScript (Node.js) Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

A Node.js example using DataStar's Server-Sent Events to manage redirects from the backend. It patches elements for a visual cue and then executes a script to navigate the user.

```javascript
import { createServer } from "node:http";
import { ServerSentEventGenerator } from "../npm/esm/node/serverSentEventGenerator.js";

const server = createServer(async (req, res) => {

  ServerSentEventGenerator.stream(req, res, async (sse) => {
    sse.patchElements(`
      <div id="indicator">Redirecting in 3 seconds...</div>
    `);

    setTimeout(() => {
      sse.executeScript(`window.location = "/guide"`);
    }, 3000);
  });
});
```

--------------------------------

### Submit Form Data on Submit Event with contentType: form

Source: https://data-star.dev/examples/form_data

This example illustrates submitting form data directly when the form's submit event is triggered. The '@get()' action is configured within the 'data-on:submit' attribute, ensuring that form data is automatically collected and sent upon submission. This is useful for forms that should submit their data without needing a separate submit button.

```html
<form data-on:submit="@get('/endpoint', {contentType: 'form'})">
    foo: <input type="text" name="foo" required />
    <button>
        Submit form
    </button>
</form>
```

--------------------------------

### GET /redirect (Node.js)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles a GET request to redirect the client after updating UI elements and executing a script using Server-Sent Events.

```APIDOC
## GET /redirect

### Description
Handles a redirect request by sending Server-Sent Events. It updates the client's HTML to show a redirect indicator, then uses `setTimeout` to wait for 3 seconds before executing a script that redirects the user to the \"/guide\" page.

### Method
GET

### Endpoint
/redirect

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```javascript
const server = createServer(async (req, res) => {

  ServerSentEventGenerator.stream(req, res, async (sse) => {
    sse.patchElements(`
      <div id="indicator">Redirecting in 3 seconds...</div>
    `);

    setTimeout(() => {
      sse.executeScript(`setTimeout(() => window.location = "/guide")`);
    }, 3000);
  });
});
```

### Response
#### Success Response (200)
Returns a successful response, initiating an SSE stream that updates the client and redirects it.

#### Response Example
(No JSON response; SSE stream is utilized)
```

--------------------------------

### Run Expressions on Animation Frames with data-on-raf

Source: https://data-star.dev/docs

The Pro attribute `data-on-raf` executes a given expression on every `requestAnimationFrame` event. This is useful for animations or continuously updating elements synchronized with the browser's rendering cycle. Changes to signals used within the expression will trigger re-execution.

```html
<div data-on-raf="$count++"></div>
```

--------------------------------

### Go: Reading Nested Signals from Incoming Requests

Source: https://data-star.dev/guide/backend_requests

Provides a Go example for reading nested signals from incoming HTTP requests using the DataStar Go SDK. It shows how to define the signal structure and use the `ReadSignals` helper function.

```go
import ("github.com/starfederation/datastar-go/datastar")

type Signals struct {
    Foo struct {
        Bar string `json:"bar"`
    } `json:"foo"`
}

signals := &Signals{}
if err := datastar.ReadSignals(request, signals); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
```

--------------------------------

### Web Server SSE: Stream Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This JavaScript code snippet shows a basic server-side setup for streaming Server-Sent Events (SSE) using Datastar. It initializes a `ServerSentEventGenerator` with the request and response objects, which also handles sending the necessary headers. This is a foundational example for server-side SSE handling.

```javascript
1// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
2ServerSentEventGenerator.stream(req, res, (stream) => {
3
4});

```

--------------------------------

### Datastar HTML for Synchronous Function Call in External Script

Source: https://data-star.dev/docs

This HTML snippet demonstrates how to bind an input's 'input' event to a local JavaScript function ('myfunction') using Datastar's data attributes. The function's return value is then displayed in a span. It requires a 'myfunction' defined in the same scope.

```html
<div data-signals:result>
    <input data-bind:foo 
        data-on:input="$result = myfunction($foo)"
    >
    <span data-text="$result"></span>
</div>
```

--------------------------------

### Set HTML Attribute Value using data-attr

Source: https://data-star.dev/docs

This demonstrates how to dynamically set an HTML attribute's value using Datastar's `data-attr` directive. It allows binding an attribute to a Datastar expression, keeping it synchronized with the expression's value.

```html
<div data-attr:title="$foo"></div>
```

--------------------------------

### Python (Sanic): `DatastarResponse` with SSE Patching

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Python example, using the Sanic web framework, demonstrates sending Server-Sent Events to patch client-side elements. It formats the current time and returns it within a `DatastarResponse`.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import DatastarResponse

@app.get("/endpoint")
async def endpoint():
    current_time = datetime.now()

    return DatastarResponse(SSE.patch_elements(f"""
        <div id="time" data-on-interval__duration.5s="@get('/endpoint')">
            {current_time:%Y-%m-%d %H:%M:%S}
        </div>
    """))

```

--------------------------------

### Python: Backend Redirect with Datastar Response

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Python example utilizes datastar_py to patch elements and perform a redirect. It requires the ServerSentEventGenerator and datastar_response from datastar_py, along with asyncio for the sleep functionality.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response
import asyncio

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.redirect("/guide")
```

--------------------------------

### CSS Color Throb Animation Example

Source: https://data-star.dev/examples/animations

This example demonstrates a simple color throb animation by maintaining a stable element ID across content swaps. Datastar facilitates CSS transitions by preserving the element's ID, allowing for smooth visual changes between old and new versions. This snippet shows a div with initial styles.

```html
<div
    id="color-throb"
    style="color: var(--blue-8); background-color: var(--orange-5);"
>
    blue on orange
</div>
```

--------------------------------

### PHP Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

A PHP example utilizing DataStar's Server-Sent Events to achieve a backend redirect. It displays a temporary message and then executes a JavaScript redirect after a 3-second delay.

```php
use starfederation\datastar\ServerSentEventGenerator;

$sse = new ServerSentEventGenerator();
$sse->patchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`);
sleep(3);
$sse->executeScript(`
    window.location = "/guide"
`);
```

--------------------------------

### Signal Casing Modifiers with data-bind and data-class

Source: https://data-star.dev/docs

Applies casing modifiers (__case) to signal names or class names when used with data-bind or data-class attributes. Supports camelCase, kebab-case, snake_case, and PascalCase. This allows for flexible naming conventions in templating languages.

```html
<input data-bind:my-signal__case.kebab />
```

```html
<div data-class:my-class__case.camel="$foo"></div>
```

--------------------------------

### Python SDK for SSE Patch Signals (Sanic)

Source: https://data-star.dev/guide/reactive_signals

This Python example uses the DataStar SDK with the Sanic web framework to demonstrate patching signals asynchronously via SSE, including a sleep interval.

```python
1from datastar_py import ServerSentEventGenerator as SSE
2from datastar_py.sanic import datastar_response
3
4@app.get('/do-you-read-me')
5@datastar_response
6async def open_doors(request):
7    yield SSE.patch_signals({"hal": "Affirmative, Dave. I read you."})
8    await asyncio.sleep(1)
9    yield SSE.patch_signals({"hal": "..."})
```

--------------------------------

### Ruby: Reading Signals with DataStar Gem

Source: https://data-star.dev/guide/backend_requests

Provides a Ruby example for reading signals using the DataStar gem. It initializes DataStar with the request and response objects and accesses signals via a dedicated attribute.

```ruby
# Setup with request
datastar = Datastar.new(request:, response:)

# Read signals
some_signal = datastar.signals[:some_signal]
```

--------------------------------

### HTML File Upload Form with Fetch

Source: https://data-star.dev/examples/file_upload

This HTML snippet defines a file upload interface. It utilizes custom attributes like 'data-bind:files' for file input and 'data-on:click' for form submission via fetch to a specified URL. Files are automatically base64 encoded, and a size limit is enforced, with console feedback for errors.

```html
<label>
    <p>Pick anything less than 1MB</p>
    <input type="file" data-bind:files multiple/>
</label>
<button
    class="warning"
    data-on:click="$files.length && @post('/examples/file_upload')"
    data-attr:aria-disabled="`${!$files.length}`">
    Submit
</button>
```

--------------------------------

### Datastar Starfield Component Attributes

Source: https://data-star.dev/index

This example shows how to configure a Datastar web component (`ds-starfield`) using reactive signals for its attributes (`center-x`, `center-y`, `speed`). The attributes are bound to backend variables using the `data-attr:*` syntax.

```html
<ds-starfield
    data-attr:center-x="$x"
    data-attr:center-y="$y"
    data-attr:speed="$speed"
></ds-starfield>
```

--------------------------------

### Create Element References with DataStar data-ref

Source: https://data-star.dev/docs

The `data-ref` attribute generates a signal that references the element it's attached to. The signal name can be defined in the attribute's key or value. Modifiers like `__case` can alter the signal name's casing.

```html
<div data-ref:foo></div>
```

```html
<div data-ref="foo"></div>
```

```html
$foo is a reference to a <span data-text="$foo.tagName"></span> element
```

```html
<div data-ref:my-signal__case.kebab></div>
```

--------------------------------

### Datastar Backend DOM Patching with SSE

Source: https://data-star.dev/index

This example illustrates how to update the DOM from the backend using Datastar's Server-Sent Events (SSE) capabilities. It patches a specific element (`#hal`) with new content and then reverts it after a short delay, showcasing real-time UI manipulation.

```go
sse.PatchElements(`
    <div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>
`)
time.Sleep(1 * time.Second)
sse.PatchElements(`<div id="hal">Waiting for an order...</div>`)
```

--------------------------------

### HTML with DataStar and SortableJS Integration

Source: https://data-star.dev/examples/sortable

This HTML structure sets up a sortable list using SortableJS and integrates with DataStar for reactive updates. It binds a signal '$orderInfo' to display the current order and uses 'data-on:reordered' to listen for sorting completion, updating the signal with event details.

```html
<div data-signals:order-info="'Initial order'" data-text="$orderInfo"></div>
<div id="sortContainer" data-on:reordered="$orderInfo = event.detail.orderInfo">
    <button>Item 1</button>
    <button>Item 2</button>
    <button>Item 3</button>
    <button>Item 4</button>
    <button>Item 5</button>
</div>
```

--------------------------------

### GET Request

Source: https://data-star.dev/reference/actions

Sends a GET request to the specified URI. Supports various options for controlling request behavior, including signal handling, background tab behavior, and content type.

```APIDOC
## GET /endpoint

### Description
Sends a `GET` request to the backend using the Fetch API. The URI can be any valid endpoint and the response must contain zero or more Datastar SSE events.
By default, requests are sent with a `Datastar-Request: true` header, and a `{datastar: *}` object containing all existing signals, except those beginning with an underscore. This behavior can be changed using the `filterSignals` option, which allows you to include or exclude specific signals using regular expressions.
When using a `get` request, the signals are sent as a query parameter, otherwise they are sent as a JSON body.
When a page is hidden (in a background tab, for example), the default behavior is for the SSE connection to be closed, and reopened when the page becomes visible again. To keep the connection open when the page is hidden, set the `openWhenHidden` option to `true`.
It’s possible to send form encoded requests by setting the `contentType` option to `form`. This sends requests using `application/x-www-form-urlencoded` encoding.
It’s also possible to send requests using `multipart/form-data` encoding by specifying it in the `form` element’s `enctype` attribute. This should be used when uploading files.

### Method
GET

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint URI.
- **options** (object) - Optional - Configuration options for the request.
  - **openWhenHidden** (boolean) - Optional - Keep the SSE connection open when the page is hidden. Defaults to `false`.
  - **contentType** (string) - Optional - Set the content type of the request. Can be 'form' for `application/x-www-form-urlencoded`.
  - **filterSignals** (object) - Optional - RegExp object to filter signals.

### Request Example
```html
<button data-on:click="@get('/endpoint')"></button>
<button data-on:click="@get('/endpoint', {openWhenHidden: true})"></button>
<button data-on:click="@get('/endpoint', {contentType: 'form'})</button>
```

### Response
#### Success Response (200)
- **Datastar SSE events** - The response contains zero or more Datastar SSE events.
```

--------------------------------

### Patch DOM Elements with Datastar SSE

Source: https://data-star.dev/docs

Patches one or more elements in the DOM using the `datastar-patch-elements` SSE event. It supports various morphing modes like 'outer', 'inner', 'replace', 'prepend', 'append', 'before', and 'after'. Elements can be targeted using CSS selectors, and view transitions can be enabled. IDs on top-level elements are recommended for proper morphing.

```html
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>


```

```html
event: datastar-patch-elements
data: mode remove
data: selector #foo


```

```html
event: datastar-patch-elements
data: mode inner
data: selector #foo
data: useViewTransition true
data: elements <div>
data: elements        Hello world!
data: elements </div>


```

--------------------------------

### Event Handling with `data-on` Attribute

Source: https://data-star.dev/docs

The `data-on` attribute attaches event listeners to elements, executing expressions when events are triggered. An `evt` variable representing the event object is available within the expression. It supports custom events and the `data-on:submit` event listener prevents default form submission behavior.

```html
<button data-on:click="$foo = ''">Reset</button>
```

```html
<div data-on:myevent="$foo = evt.detail"></div>
```

--------------------------------

### `data-on` Event Modifiers for Event Handling

Source: https://data-star.dev/docs

This snippet demonstrates various modifiers for the `data-on` attribute, controlling event listener behavior. Modifiers include `__once`, `__passive`, `__capture`, `__case` (with sub-options like `.camel`, `.kebab`), `__delay`, `__debounce` (with options like `.leading`, `.notrailing`), `__throttle` (with options like `.noleading`, `.trailing`), `__viewtransition`, `__window`, `__outside`, `__prevent`, and `__stop`.

```html
<button data-on:click__window__debounce.500ms.leading="$foo = ''"></button>
```

```html
<div data-on:my-event__case.camel="$foo = ''"></div>
```

--------------------------------

### HTML: Practical Use-Case of Nested Signals for Menu State

Source: https://data-star.dev/guide/backend_requests

Illustrates a practical scenario of nested signals to manage the open/closed state of multiple menus (desktop and mobile) and a function to toggle them all. This example highlights efficient state management.

```html
<div data-signals="{menu: {isOpen: {desktop: false, mobile: false}}}">
    <button data-on:click="@toggleAll({include: /^menu\.isOpen\./})">
        Open/close menu
    </button>
</div>
```

--------------------------------

### HTML Button for DataStar Redirect

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

An HTML button with a `data-on:click` attribute that triggers a GET request to a specified endpoint. It also includes a `div` for displaying a redirect indicator.

```html
<button data-on:click="@get('/endpoint')">
    Click to be redirected from the backend
</button>
<div id="indicator"></div>
```

--------------------------------

### Datastar HTML for Asynchronous Function Call in External Script

Source: https://data-star.dev/docs

This HTML snippet shows how to handle asynchronous functions in Datastar. The input event triggers an async function ('myfunction'), which dispatches a custom event ('mycustomevent') with the result. The component listens for this event to update the UI. Datastar does not await async calls directly within expressions.

```html
<div data-signals:result>
    <input data-bind:foo 
           data-on:input="myfunction(el, $foo)"
           data-on:mycustomevent__window="$result = evt.detail.value"
    >
    <span data-text="$result"></span>
</div>
```

--------------------------------

### Patch DOM Elements with Server-Sent Events (SSE)

Source: https://data-star.dev/docs

This functionality allows updating specific HTML elements in the browser using Server-Sent Events. It is implemented in PHP, Python (with Sanic), Ruby, Rust, and JavaScript. The core idea is to send an SSE event that targets an element by its ID and provides new HTML content.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
$sse = new ServerSentEventGenerator();

// Patches elements into the DOM.
$sse->patchElements(
    '<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>'
);

sleep(1);

$sse->patchElements(
    '<div id="hal">Waiting for an order...</div>'
);
```

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response
import asyncio

@app.get('/open-the-bay-doors')
@datastar_response
async def open_doors(request):
    yield SSE.patch_elements('<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>')
    await asyncio.sleep(1)
    yield SSE.patch_elements('<div id="hal">Waiting for an order...</div>')
```

```ruby
require 'datastar'

# Create a Datastar::Dispatcher instance
datastar = Datastar.new(request:, response:)

# Start a streaming response
datastar.stream do |sse|
  # Patches elements into the DOM.
  sse.patch_elements %(<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>)

  sleep 1
  
  sse.patch_elements %(<div id="hal">Waiting for an order...</div>)
end
```

```rust
use async_stream::stream;
use datastar::prelude::*;
use std::thread;
use std::time::Duration;

Sse(stream! {
    // Patches elements into the DOM.
    yield PatchElements::new("<div id='hal'>I’m sorry, Dave. I’m afraid I can’t do that.</div>").into();

    thread::sleep(Duration::from_secs(1));
    
    yield PatchElements::new("<div id='hal'>Waiting for an order...</div>").into();
})
```

```javascript
// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
ServerSentEventGenerator.stream(req, res, (stream) => {
    // Patches elements into the DOM.
    stream.patchElements(`<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`);

    setTimeout(() => {
        stream.patchElements(`<div id="hal">Waiting for an order...</div>`);
    }, 1000);
});
```

--------------------------------

### Datastar Backend Actions

Source: https://data-star.dev/reference/actions

Provides backend actions for common HTTP requests like GET, POST, PUT, PATCH, and DELETE. These actions are designed to interact with server-side APIs.

```APIDOC
## Backend Actions

### Description
These actions facilitate making HTTP requests to backend services.

### Methods
- `@get()`
- `@post()`
- `@put()`
- `@patch()`
- `@delete()`

### Endpoint
N/A (Used within Datastar expressions to trigger backend requests)

### Parameters
(Specific parameters depend on the chosen HTTP method, typically including URL, data, and options.)

### Request Example
```html
<!-- Example using @get() -->
<button data-on:click="@get('/api/users')">Fetch Users</button>

<!-- Example using @post() -->
<button data-on:click="@post('/api/users', { name: 'John Doe' })">Create User</button>
```

### Response
(Responses will vary based on the backend API endpoint being called.)

#### Success Response (200)
(Details of the successful response payload from the backend.)

#### Response Example
```json
{
  "message": "Success"
}
```
```

--------------------------------

### Go: Backend Logic for Loading More

Source: https://data-star.dev/how_tos/load_more_list_items

This Go implementation demonstrates handling Datastar requests to load more items. It reads signals, constructs new elements, and sends `PatchElements`, `PatchSignals`, or `RemoveElements` events via Server-Sent Events.

```go
1import (
2    "fmt"
3    "net/http"
4
5    "github.com/go-chi/chi/v5"
6    "github.com/starfederation/datastar-go/datastar"
7)
8
9type OffsetSignals struct {
10    Offset int `json:"offset"`
11}
12
13signals := &OffsetSignals{}
14if err := datastar.ReadSignals(r, signals); err != nil {
15    http.Error(w, err.Error(), http.StatusBadRequest)
16}
17
18max := 5
19limit := 1
20offset := signals.Offset
21
22sse := datastar.NewSSE(w, r)
23
24if offset < max {
25    newOffset := offset + limit
26    sse.PatchElements(fmt.Sprintf(`<div>Item %d</div>`, newOffset),
27        datastar.WithSelectorID("list"),
28        datastar.WithModeAppend(),
29    )
30    if newOffset < max {
31        sse.PatchSignals([]byte(fmt.Sprintf(`{offset: %d}`, newOffset)))
32    } else {
33        sse.RemoveElements(`#load-more`)
34    }
35}
```

--------------------------------

### Set Element Text Content with `data-text` Attribute

Source: https://data-star.dev/docs

The `data-text` attribute binds the text content of an HTML element to the value of a Datastar signal. It can also evaluate Datastar expressions, allowing for dynamic text manipulation like calling JavaScript methods on the signal.

```html
<input data-bind:foo />
<div data-text="$foo"></div>
<div data-text="$foo.toUpperCase()"></div>
```

--------------------------------

### HTML for Progress Bar and Completion Button

Source: https://data-star.dev/examples/progress_bar

This HTML code defines a progress bar using SVG and a completion button. The progress bar visually indicates completion percentage, and the button appears once the progress is 100%, allowing the user to restart the process. It utilizes custom data attributes for Datastar integration, including SSE event handling.

```html
<div
    id="progress-bar"
    data-init="@get('/examples/progress_bar/updates', {openWhenHidden: true})"
>
    <!-- When progress is less than 100% -->
    <svg
        width="200"
        height="200"
        viewbox="-25 -25 250 250"
        style="transform: rotate(-90deg)"
    >
        <circle
            r="90"
            cx="100"
            cy="100"
            fill="transparent"
            stroke="#e0e0e0"
            stroke-width="16px"
            stroke-dasharray="565.48px"
            stroke-dashoffset="565px"
        ></circle>
        <circle
            r="90"
            cx="100"
            cy="100"
            fill="transparent"
            stroke="#6bdba7"
            stroke-width="16px"
            stroke-linecap="round"
            stroke-dashoffset="282px"
            stroke-dasharray="565.48px"
        ></circle>
        <text
            x="44px"
            y="115px"
            fill="#6bdba7"
            font-size="52px"
            font-weight="bold"
            style="transform:rotate(90deg) translate(0px, -196px)"
        >50%</text>
    </svg>
    
    <!-- When progress is 100% -->
    <button
        data-indicator:_fetching
        data-attr:aria-disabled="`${$_fetching}`"
        data-on:click="
            !$_fetching && @get('/examples/progress_bar/updates', {openWhenHidden: true})
        "
    >
        <i class="material-symbols:check-circle"></i>
        Completed! Try again?
    </button>
</div>
```

--------------------------------

### Backend Response: `datastar-patch-elements` Event

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This example shows the structure of a backend response using a `datastar-patch-elements` event. The event data contains the updated HTML for the element, allowing the frontend to dynamically refresh content.

```http
event: datastar-patch-elements
data: elements <div id="time" data-on-interval__duration.5s="@get('/endpoint')">
data: elements     {{ now }}
data: elements </div>


```

--------------------------------

### Linearly Interpolate Values

Source: https://data-star.dev/docs

Uses the `@fit()` action to linearly interpolate a value from one range to another. This is useful for scaling values, such as mapping a slider input to an RGB color component or converting temperature units. Optional parameters allow clamping the output to the new range and rounding to the nearest integer.

```html
<!-- Convert a 0-100 slider to 0-255 RGB value -->
<div>
    <input type="range" min="0" max="100" value="50" data-bind:slider-value>
    <div data-computed:rgb-value="@fit($sliderValue, 0, 100, 0, 255)">
        RGB Value: <span data-text="$rgbValue"></span>
    </div>
</div>

<!-- Convert Celsius to Fahrenheit -->
<div>
    <input type="number" data-bind:celsius value="20" />
    <div data-computed:fahrenheit="@fit($celsius, 0, 100, 32, 212)">
        <span data-text="$celsius"></span>°C = <span data-text="$fahrenheit.toFixed(1)"></span>°F
    </div>
</div>

<!-- Map mouse position to element opacity (clamped) -->
<div
    data-signals:mouse-x="0"
    data-computed:opacity="@fit($mouseX, 0, window.innerWidth, 0, 1, true)"
    data-on:mousemove__window="$mouseX = evt.clientX"
    data-attr:style="'opacity: ' + $opacity"
>
    Move your mouse horizontally to change opacity
</div>
```

--------------------------------

### C# SDK for HTTP Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This C# code demonstrates how to use the DataStar SDK within an ASP.NET Core application to patch signals asynchronously over HTTP GET requests, including delays.

```csharp
 1using StarFederation.Datastar.DependencyInjection;
 2
 3// Adds Datastar as a service
 4builder.Services.AddDatastar();
 5
 6app.MapGet("/hal", async (IDatastarService datastarService) =>
 7{
 8    // Patches signals.
 9    await datastarService.PatchSignalsAsync(new { hal = "Affirmative, Dave. I read you" });
10
11    await Task.Delay(TimeSpan.FromSeconds(3));
12
13    await datastarService.PatchSignalsAsync(new { hal = "..." });
14});
```

--------------------------------

### Datastar HTML for Bad Apple Benchmark

Source: https://data-star.dev/examples/bad_apple

HTML structure using Datastar signals to display the Bad Apple video benchmark. It includes a label to manage animation state and percentage, and a pre tag to display the video frames. The animation state is initialized via a GET request.

```html
<label
    data-signals="{_percentage: 0, _contents: 'bad apple frames go here'}"
    data-init="@get('/examples/bad_apple/updates')"
>
    <span data-text="`Percentage: ${$_percentage.toFixed(2)}%`"></span>
    <input
        type="range"
        min="0"
        max="100"
        step="0.01"
        disabled
        style="cursor: default"
        data-attr:value="$_percentage"
    />
</label>
<pre style="line-height: 100%" data-text="$_contents"></pre>
```

--------------------------------

### C#: `IDatastarService.PatchElementsAsync` Example

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This C# snippet shows how to use the `IDatastarService` to asynchronously patch elements on the client. It fetches the current time and sends an updated HTML div to the client at a 5-second interval.

```csharp
using StarFederation.Datastar.DependencyInjection;

app.MapGet("/endpoint", async (IDatastarService datastarService) =>
{
    var currentTime = DateTime.Now.ToString("yyyy-MM-dd hh:mm:ss");
    await datastarService.PatchElementsAsync($"\
        <div id=\"time\" data-on-interval__duration.5s=\"@get('/endpoint')\">\
            {currentTime}\
        </div>
    ");
});

```

--------------------------------

### Send GET Request with Datastar

Source: https://data-star.dev/reference/actions

Sends a GET request to a specified URI using the Fetch API. Supports options for controlling signal filtering, background tab behavior, and content type. Signals are sent as query parameters by default.

```html
<button data-on:click="@get('/endpoint')"></button>
```

```html
<button data-on:click="@get('/endpoint', {openWhenHidden: true})"></button>
```

```html
<button data-on:click="@get('/endpoint', {contentType: 'form'})></button>
```

```html
<form enctype="multipart/form-data">
    <input type="file" name="file" />
    <button data-on:click="@get('/endpoint', {contentType: 'form'})></button>
</form>
```

--------------------------------

### DataStar Interval: Immediate Execution with `.leading`

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This example shows how to modify the interval attribute to execute the expression immediately upon page load in addition to the regular interval. This is achieved by appending the `.leading` modifier.

```html
<div id="time"
     data-on-interval__duration.5s.leading="@get('/endpoint')"
></div>
```

--------------------------------

### Log all signal patches on change (HTML)

Source: https://data-star.dev/reference

The `data-on-signal-patch` attribute allows you to run an expression whenever any signals are patched. This example logs the details of the signal patch to the console using the available `patch` variable.

```html
<div data-on-signal-patch="console.log('A signal changed!')"></div>
```

```html
<div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
```

--------------------------------

### Send GET Request with Datastar

Source: https://data-star.dev/guide

This snippet demonstrates how to send a GET request to a specified URL using Datastar's `@get()` action. It's useful for fetching data from the backend and updating the DOM based on the response. Ensure the backend returns HTML content for morphing.

```html
1<button data-on:click="@get('/endpoint')">
2    Open the pod bay doors, HAL.
3</button>
4<div id="hal"></div>
```

--------------------------------

### Replicating Datastar Pro Attributes with Free Version (HTML)

Source: https://data-star.dev/essays/greedy_developer

These HTML snippets demonstrate how to achieve functionality similar to Datastar Pro's convenience plugins using the free version. The first snippet replaces the current URL on load and when the '$page' variable changes, while the second scrolls an element into view.

```html
1<!-- Replaces the current URL on load and whenever $page changes. -->
<div data-effect="window.history.replaceState({}, '', '/page/' + $page)"></div>

<!-- Scrolls the element into view. -->
<div data-init="el.scrollIntoView()"></div>
```

--------------------------------

### Predefined Signal Types with data-bind

Source: https://data-star.dev/docs

Demonstrates how predefined signal types (e.g., number, array) are preserved during two-way data binding. When an element's value changes, it's automatically converted to match the original signal type. This is useful for maintaining data integrity, especially with form inputs like select options and checkboxes.

```html
<div data-signals:foo="0">
    <select data-bind:foo>
        <option value="10">10</option>
    </select>
</div>
```

```html
<div data-signals:foo="[]">
    <input data-bind:foo type="checkbox" value="bar" />
    <input data-bind:foo type="checkbox" value="baz" />
</div>
```

--------------------------------

### Kotlin: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Kotlin code example demonstrates using the Datastar SDK to stream Server-Sent Events (SSE). It shows the initialization of a `ServerSentEventGenerator` and calls to `patchElements` and `patchSignals` methods. This snippet is concise and leverages Kotlin's syntax for building SSE responses.

```kotlin
1val generator = ServerSentEventGenerator(response)
2
generator.patchElements(
3    elements = """<div id=\"question\">What do you put in a toaster?</div>""",
4)
5
generator.patchSignals(
6    signals = "{"response": "", "answer": "bread"}",
7)

```

--------------------------------

### Implement Alert Action Plugin

Source: https://data-star.dev/examples/custom_plugin

This JavaScript code defines a custom action plugin named 'alert'. It takes a context and a value, and applies the value by triggering a browser alert. This allows for custom event handling that displays messages to the user.

```javascript
1action({
    name: 'alert',
    apply(ctx, value) {
        alert(value)
    }
})
```

--------------------------------

### Rust: SSE with `PatchElements` using `async-stream`

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Rust example demonstrates creating Server-Sent Events to patch client elements. It uses `chrono` for time formatting and `async-stream` to generate the SSE stream containing updated HTML.

```rust
use datastar::prelude::*;
use chrono::Local;
use async_stream::stream;

let current_time = Local::now().format("%Y-%m-%d %H:%M:%S").to_string();

Sse(stream! {
    yield PatchElements::new(
        format!(
            "<div id='time' data-on-interval__duration.5s='@get(\"/endpoint\")'>{}</div>",
            current_time
        )
    ).into();
})

```

--------------------------------

### Send Multiple SSE Events in PHP using Datastar

Source: https://data-star.dev/guide/backend_requests

Demonstrates sending SSE events, including patching elements and signals, using PHP with the Datastar library. This example is for PHP backend implementations.

```php
$sse->patchElements('<div id="question">...</div>');
$sse->patchElements('<div id="instructions">...</div>');
$sse->patchSignals(['answer' => '...', 'prize' => '...']);
```

--------------------------------

### Patch Signals with data-signals Attribute

Source: https://data-star.dev/docs

The data-signals attribute allows patching (adding, updating, or removing) signals. Values defined later in the DOM override earlier ones. Nested signals use dot-notation, and multiple signals can be patched using key-value pairs in JavaScript or JSON format. Setting a signal to null or undefined removes it. Keys are converted to camel case.

```html
<div data-signals:foo="1"></div>
<div data-signals:foo.bar="1"></div>
<div data-signals="{foo: {bar: 1, baz: 2}}"></div>
<div data-signals="{foo: null}"></div>
<div data-signals:my-signal="1"></div>
<div data-signals="{mySignal: 1}"></div>
```

--------------------------------

### Combine include and exclude filters for signal patches (HTML)

Source: https://data-star.dev/reference

This example shows how to use both `include` and `exclude` properties within `data-on-signal-patch-filter` to create a more specific filter. It will react to signals containing 'user' but exclude those containing 'password'.

```html
<!-- Combine include and exclude filters -->
<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
```

--------------------------------

### Two-Way Data Binding with data-bind Attribute

Source: https://data-star.dev/docs

Establishes two-way data binding between an HTML element's value and a signal. Updates occur in both directions: element changes update the signal, and signal changes update the element. Supports input, select, textarea, and web components, listening for 'change' and 'input' events. Signal names can be specified in the attribute key or value.

```html
<input data-bind:foo />
```

```html
<input data-bind="foo" />
```

```html
<input data-bind:foo value="bar" />
```

```html
<div data-signals:foo="baz">
    <input data-bind:foo value="bar" />
</div>
```

--------------------------------

### HTML Structure for Web Component Binding

Source: https://data-star.dev/examples/web_component

This HTML snippet sets up the structure for a web component, including input fields and display areas. It utilizes custom attributes (`data-bind`, `data-signals`, `data-text`, `data-on`, `data-attr`) for data binding and event handling, commonly found in declarative UI frameworks.

```html
<label>
    Reversed
    <input type="text" value="Your Name" data-bind:_name/>
</label>
<span data-signals:_reversed data-text="$ _reversed"></span>
<reverse-component
    data-on:reverse="$_reversed = evt.detail.value"
    data-attr:name="$_name">
</reverse-component>
```

--------------------------------

### Submit Form Data via GET/POST with contentType: form

Source: https://data-star.dev/examples/form_data

This snippet shows how to submit form data using both GET and POST requests with the 'contentType' option set to 'form'. It demonstrates submitting data from within the form and from an external button, leveraging automatic form selection and validation. The 'selector' option can be used to target specific forms.

```html
<form id="myform">
    foo:<input type="checkbox" name="checkboxes" value="foo" />
    bar:<input type="checkbox" name="checkboxes" value="bar" />
    baz:<input type="checkbox" name="checkboxes" value="baz" />
    <button data-on:click="@get('/endpoint', {contentType: 'form'})">
        Submit GET request
    </button>
    <button data-on:click="@post('/endpoint', {contentType: 'form'})">
        Submit POST request
    </button>
</form>

<button data-on:click="@get('/endpoint', {contentType: 'form', selector: '#myform'})">
    Submit GET request from outside the form
</button>
```

--------------------------------

### Kotlin: `ServerSentEventGenerator` for Patching Elements

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Kotlin example utilizes `ServerSentEventGenerator` to patch elements on the client. It formats the current time and includes it in an HTML `div` with a specified 5-second update interval.

```kotlin
val now: LocalDateTime = currentTime()

val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements =
        """
        <div id="time" data-on-interval__duration.5s="@get('/endpoint')">
            $now
        </div>
        """.trimIndent()
)

```

--------------------------------

### C# Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates a backend redirect using C# with the DataStar library. It patches elements to show a redirect message, waits for 3 seconds, and then executes a script to navigate the user.

```csharp
using StarFederation.Datastar.DependencyInjection;

app.MapGet("/redirect", async (IDatastarService datastarService) =>
{
    await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
    await Task.Delay(TimeSpan.FromSeconds(3));
    await datastarService.ExecuteScriptAsync("window.location = \"/guide\";");
});
```

--------------------------------

### SSE Event for Patching Signals

Source: https://data-star.dev/guide/reactive_signals

This example shows how to format a Server-Sent Event (SSE) to trigger the 'datastar-patch-signals' event, updating frontend signals with new data.

```sse
1event: datastar-patch-signals
2data: signals {hal: 'Affirmative, Dave. I read you.'}
3

```

--------------------------------

### Display Contact Details and Edit/Reset Buttons (HTML)

Source: https://data-star.dev/examples/click_to_edit

This HTML snippet displays contact information (First Name, Last Name, Email) and provides 'Edit' and 'Reset' buttons. The 'Edit' button triggers a GET request to '/examples/click_to_edit/edit', while the 'Reset' button triggers a PATCH request to '/examples/click_to_edit/reset'. It utilizes custom attributes for data binding and event handling.

```html
<div id="demo">
    <p>First Name: John</p>
    <p>Last Name: Doe</p>
    <p>Email: joe@blow.com</p>
    <div role="group">
        <button
            class="info"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
            data-on:click="@get('/examples/click_to_edit/edit')"
        >
            Edit
        </button>
        <button
            class="warning"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
            data-on:click="@patch('/examples/click_to_edit/reset')"
        >
            Reset
        </button>
    </div>
</div>
```

--------------------------------

### JavaScript for SortableJS Initialization and Event Handling

Source: https://data-star.dev/examples/sortable

This JavaScript code initializes the SortableJS library on the 'sortContainer' element. It configures the Sortable instance to dispatch a custom 'reordered' event upon completion of a drag-and-drop operation, passing the old and new indices of the moved item.

```javascript
import Sortable from 'https://cdn.jsdelivr.net/npm/sortablejs/+esm'
new Sortable(sortContainer, {
    animation: 150,
    ghostClass: 'opacity-25',
    onEnd: (evt) => {
        sortContainer.dispatchEvent(
            new CustomEvent('reordered', {detail: {
                orderInfo: `Moved from position ${evt.oldIndex + 1} to ${evt.newIndex + 1}`
            }})
        )
    }
})
```

--------------------------------

### Send Multiple SSE Events in Ruby using Datastar

Source: https://data-star.dev/guide/backend_requests

Provides an example of sending SSE events, patching elements, and signals using Ruby with the Datastar framework. This is useful for Ruby on Rails or other Ruby backend projects.

```ruby
return DatastarResponse([
    SSE.patch_elements('<div id="question">...</div>'),
    SSE.patch_elements('<div id="instructions">...</div>'),
    SSE.patch_signals({"answer": "...", "prize": "..."})
])
```

--------------------------------

### Setting HTML Response Headers for Datastar DOM Patching

Source: https://data-star.dev/reference/actions

This example shows how to set response headers to control how HTML content is patched into the DOM using Datastar. It specifies the content type, the target selector, and the patching mode, allowing for dynamic updates of specific page elements.

```javascript
response.headers.set('Content-Type', 'text/html')
response.headers.set('datastar-selector', '#my-element')
response.headers.set('datastar-mode', 'inner')
response.body = '<p>New content</p>'
```

--------------------------------

### Ruby: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Ruby code snippet shows how to implement Server-Sent Events (SSE) using the Datastar gem within a Rack application. It demonstrates initializing a `Datastar` dispatcher and using the `stream` block to send patched elements and signals. This example highlights Datastar's integration with the Ruby Rack ecosystem.

```ruby
1require 'datastar'
2
3# Create a Datastar::Dispatcher instance
4
datastar = Datastar.new(request:, response:)
5
6# In a Rack handler, you can instantiate from the Rack env
7# datastar = Datastar.from_rack_env(env)
8
9# Start a streaming response
10datastar.stream do |sse|
11  # Patches elements into the DOM
12  sse.patch_elements %(<div id="question">What do you put in a toaster?</div>)
13
14  # Patches signals
15  sse.patch_signals(response: '', answer: 'bread')
16end

```

--------------------------------

### Data-Star Attributes for Interactive Elements

Source: https://data-star.dev/guide/backend_requests

This snippet shows how to use Data-Star attributes to create interactive elements. `data-on:click` triggers a GET request to a server endpoint. `data-signals` initializes signals, `data-computed` creates computed properties, `data-show` conditionally displays elements, and `data-text` sets element text content.

```html
<div
    data-signals="{response: '', answer: ''}"
    data-computed:correct="$response.toLowerCase() == $answer">
    <div id="question"></div>
    <button data-on:click="@get('/actions/quiz')">Fetch a question</button>
    <button
        data-show="$answer != ''"
        data-on:click="$response = prompt('Answer:') ?? ''">
        BUZZ
    </button>
    <div data-show="$response != ''">
        You answered “<span data-text="$response"></span>”.
        <span data-show="$correct">That is correct ✅</span>
        <span data-show="!$correct">
        The correct answer is “<span data-text="$answer"></span>” 🤷
        </span>
    </div>
</div>
```

--------------------------------

### Node.js: `ServerSentEventGenerator` for Streaming Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Node.js example uses the `ServerSentEventGenerator` to stream updates to the client. It captures the current time and sends it within an HTML structure, configured for a 5-second interval update.

```javascript
import { createServer } from "node:http";
import { ServerSentEventGenerator } from "../npm/esm/node/serverSentEventGenerator.js";

const server = createServer(async (req, res) => {
  const currentTime = new Date().toISOString();
  
  ServerSentEventGenerator.stream(req, res, (sse) => {
    sse.patchElements(`
       <div id="time"
          data-on-interval__duration.5s="@get('/endpoint')"
       >
         ${currentTime}
       </div>
    `);
  });
});

```

--------------------------------

### Datastar Expressions: Multiple Statements with Semicolon

Source: https://data-star.dev/guide/datastar_expressions

Illustrates how multiple statements can be executed within a single Datastar expression by separating them with a semicolon. This example shows updating a signal and then triggering an action.

```html
1<div data-signals:foo="1">
2    <button data-on:click="$landingGearRetracted = true; @post('/launch')">
3        Force launch
4    </button>
5</div>
```

--------------------------------

### Send PATCH Request with Datastar

Source: https://data-star.dev/reference/actions

Sends a PATCH request to a specified URI. Operates identically to `@get()` but employs the PATCH HTTP method. It accepts configuration options for the request.

```html
<button data-on:click="@patch('/endpoint')"></button>
```

--------------------------------

### Handle DataStar Fetch Events

Source: https://data-star.dev/reference/actions

Listens for 'datastar-fetch' events on a DOM element to react to different stages of a fetch request. The event detail includes the type of event (e.g., 'started', 'finished', 'error').

```html
<div data-on:datastar-fetch="
    evt.detail.type === 'error' && console.log('Fetch error encountered')
"></div>
```

--------------------------------

### Kotlin: Reading Signals with JSON Unmarshaller

Source: https://data-star.dev/guide/backend_requests

Illustrates reading signals in Kotlin using a JSON unmarshaller. This example defines a `Signals` data class and uses `Json.decodeFromString` to parse incoming JSON request bodies.

```kotlin
@Serializable
data class Signals(
    val foo: String,
)

val jsonUnmarshaller: JsonUnmarshaller<Signals> = { json -> Json.decodeFromString(json) }

val request: Request =
    postRequest(
        body =
            """
            {
                "foo": "bar"
            }
            """.trimIndent(),
    )

val signals = readSignals(request, jsonUnmarshaller)
```

--------------------------------

### Send Multiple SSE Events in Java using Datastar

Source: https://data-star.dev/guide/backend_requests

Provides an example of sending SSE events, including patching elements and signals, using Java with the Datastar library. This is intended for Java-based backend applications.

```java
generator.send(PatchElements.builder()
    .data("<div id=\"question\">...</div>")
    .build()
);
generator.send(PatchElements.builder()
    .data("<div id=\"instructions\">...</div>")
    .build()
);
generator.send(PatchSignals.builder()
    .data("{\"answer\": \"...\", \"prize\": \"...\"}")
    .build()
);
```

--------------------------------

### POST Request

Source: https://data-star.dev/reference/actions

Sends a POST request to the specified URI. Similar to the GET request, it supports options for controlling request behavior, but signals are sent as a JSON body by default.

```APIDOC
## POST /endpoint

### Description
Sends a `POST` request to the backend using the Fetch API. Works the same as `@get()` but sends a `POST` request to the backend.

### Method
POST

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint URI.
- **options** (object) - Optional - Configuration options for the request.
  - **openWhenHidden** (boolean) - Optional - Keep the SSE connection open when the page is hidden. Defaults to `false`.
  - **contentType** (string) - Optional - Set the content type of the request. Can be 'form' for `application/x-www-form-urlencoded`.
  - **filterSignals** (object) - Optional - RegExp object to filter signals.

### Request Example
```html
<button data-on:click="@post('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response body** - The response from the POST request.
```

--------------------------------

### PATCH Request

Source: https://data-star.dev/reference/actions

Sends a PATCH request to the specified URI. Operates similarly to the GET request, with signals transmitted as a JSON body by default.

```APIDOC
## PATCH /endpoint

### Description
Sends a `PATCH` request to the backend using the Fetch API. Works the same as `@get()` but sends a `PATCH` request to the backend.

### Method
PATCH

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint URI.
- **options** (object) - Optional - Configuration options for the request.
  - **openWhenHidden** (boolean) - Optional - Keep the SSE connection open when the page is hidden. Defaults to `false`.
  - **contentType** (string) - Optional - Set the content type of the request. Can be 'form' for `application/x-www-form-urlencoded`.
  - **filterSignals** (object) - Optional - RegExp object to filter signals.

### Request Example
```html
<button data-on:click="@patch('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response body** - The response from the PATCH request.
```

--------------------------------

### Send POST Request with Datastar

Source: https://data-star.dev/reference/actions

Sends a POST request to a specified URI. This function behaves similarly to `@get()` but utilizes the POST HTTP method. Options for request configuration are supported.

```html
<button data-on:click="@post('/endpoint')"></button>
```

--------------------------------

### PUT Request

Source: https://data-star.dev/reference/actions

Sends a PUT request to the specified URI. Functionality mirrors the GET request, with signals sent as a JSON body by default.

```APIDOC
## PUT /endpoint

### Description
Sends a `PUT` request to the backend using the Fetch API. Works the same as `@get()` but sends a `PUT` request to the backend.

### Method
PUT

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint URI.
- **options** (object) - Optional - Configuration options for the request.
  - **openWhenHidden** (boolean) - Optional - Keep the SSE connection open when the page is hidden. Defaults to `false`.
  - **contentType** (string) - Optional - Set the content type of the request. Can be 'form' for `application/x-www-form-urlencoded`.
  - **filterSignals** (object) - Optional - RegExp object to filter signals.

### Request Example
```html
<button data-on:click="@put('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response body** - The response from the PUT request.
```

--------------------------------

### Send Multiple SSE Events in Crystal using Datastar

Source: https://data-star.dev/guide/backend_requests

Shows how to send SSE events for patching elements and signals using Crystal and the Datastar library. This example is suitable for Crystal-based backend applications.

```crystal
yield PatchElements::new("<div id='question'>...</div>").into()
yield PatchElements::new("<div id='instructions'>...</div>").into()
yield PatchSignals::new("{answer: '...', prize: '...'}").into()
```

--------------------------------

### Editable Contact Form (HTML)

Source: https://data-star.dev/examples/click_to_edit

This HTML snippet presents an editable form for contact details, including input fields for First Name, Last Name, and Email. It also includes 'Save' and 'Cancel' buttons. The 'Save' button triggers a PUT request to '/examples/click_to_edit', and the 'Cancel' button triggers a GET request to '/examples/click_to_edit/cancel'. Input fields are bound to signals, and buttons manage fetching states.

```html
<div id="demo">
    <label>
        First Name
        <input
            type="text"
            data-bind:first-name
            data-attr:disabled="$_fetching"
        >
    </label>
    <label>
        Last Name
        <input
            type="text"
            data-bind:last-name
            data-attr:disabled="$_fetching"
        >
    </label>
    <label>
        Email
        <input
            type="email"
            data-bind:email
            data-attr:disabled="$_fetching"
        >
    </label>
    <div role="group">
        <button
            class="success"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
            data-on:click="@put('/examples/click_to_edit')"
        >
            Save
        </button>
        <button
            class="error"
            data-indicator:_fetching
            data-attr:disabled="$_fetching"
            data-on:click="@get('/examples/click_to_edit/cancel')"
        >
            Cancel
        </button>
    </div>
</div>
```

--------------------------------

### HTML: Read-only Row Structure

Source: https://data-star.dev/examples/edit_row

This snippet shows the HTML structure of a table row before editing. It includes table data cells for Name and Email, and an 'Edit' button.

```html
<tr>
    <td>Joe Smith</td>
    <td>joe@smith.org</td>
    <td>
        <button data-on:click="@get('/examples/edit_row/0')">
            Edit
        </button>
    </td>
</tr>
```

--------------------------------

### Use Alert Action in HTML

Source: https://data-star.dev/examples/custom_plugin

This HTML snippet demonstrates how to use the custom 'alert' action plugin. The `@alert` syntax is used within a `data-on:click` attribute to execute the alert action with a provided string message when the button is clicked.

```html
1<button data-on:click="@alert('Hello from an action')">
2    Alert using an action
3</button>
```

--------------------------------

### Debounce event listener for signal patches (HTML)

Source: https://data-star.dev/reference

This example uses `data-on-signal-patch` with the `__debounce` modifier and a `.500ms` timing to delay the execution of the `doSomething()` function. This ensures the function is only called after a period of inactivity, preventing rapid, repeated calls.

```html
<div data-on-signal-patch__debounce.500ms="doSomething()"></div>
```

--------------------------------

### Python Litestar: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Python code snippet demonstrates streaming Server-Sent Events (SSE) within a Litestar application using the `datastar-py` library. It shows how to create a `DatastarResponse` that includes both patched elements and signals, directly within an async endpoint function. This example leverages the library's integration with modern Python web frameworks.

```python
1from datastar_py import ServerSentEventGenerator as SSE
2from datastar_py.litestar import DatastarResponse
3
4async def endpoint():
5    return DatastarResponse([
6        SSE.patch_elements('<div id="question">What do you put in a toaster?</div>'),
7        SSE.patch_signals({"response": "", "answer": "bread"})
8    ])

```

--------------------------------

### DELETE Request

Source: https://data-star.dev/reference/actions

Sends a DELETE request to the specified URI. Functionality is consistent with the GET request, with signals sent as a JSON body by default.

```APIDOC
## DELETE /endpoint

### Description
Sends a `DELETE` request to the backend using the Fetch API. Works the same as `@get()` but sends a `DELETE` request to the backend.

### Method
DELETE

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint URI.
- **options** (object) - Optional - Configuration options for the request.
  - **openWhenHidden** (boolean) - Optional - Keep the SSE connection open when the page is hidden. Defaults to `false`.
  - **contentType** (string) - Optional - Set the content type of the request. Can be 'form' for `application/x-www-form-urlencoded`.
  - **filterSignals** (object) - Optional - RegExp object to filter signals.

### Request Example
```html
<button data-on:click="@delete('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response body** - The response from the DELETE request.
```

--------------------------------

### HTML Structure for Lazy Tabs

Source: https://data-star.dev/examples/lazy_tabs

This HTML structure defines a set of tabs using ARIA roles for accessibility. Each tab button has a `data-on:click` attribute to trigger Datastar's event handling, which fetches tab content. The selected tab is indicated by `aria-selected='true'`, and tab content is displayed within a `role='tabpanel'` element.

```html
<div id="demo">
    <div role="tablist">
        <button
            role="tab"
            aria-selected="true"
            data-on:click="@get('/examples/lazy_tabs/0')"
        >
            Tab 0
        </button>
        <button
            role="tab"
            aria-selected="false"
            data-on:click="@get('/examples/lazy_tabs/1')"
        >
            Tab 1
        </button>
        <button
            role="tab"
            aria-selected="false"
            data-on:click="@get('/examples/lazy_tabs/2')"
        >
            Tab 2
        </button>
        <!-- More tabs... -->
    </div>
    <div role="tabpanel">
        <p>Lorem ipsum dolor sit amet...</p>
        <p>Consectetur adipiscing elit...</p>
        <!-- Tab content -->
    </div>
</div>
```

--------------------------------

### HTML Table with Delete Button

Source: https://data-star.dev/examples/delete_row

This snippet displays an HTML table with rows containing data and an 'Actions' column. Each row includes a 'Delete' button that triggers a confirmation dialog and initiates a delete request upon confirmation. It utilizes custom data attributes for click handling and disabling the button during the fetching state.

```html
<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Joe Smith</td>
            <td>joe@smith.org</td>
            <td>
                <button
                    class="error"
                    data-on:click="confirm('Are you sure?') && @delete('/examples/delete_row/0')"
                    data-indicator:_fetching
                    data-attr:disabled="$_fetching"
                >
                    Delete
                </button>
            </td>
        </tr>
    </tbody>
</table>
```

--------------------------------

### Receive Multiple Patch Elements SSE Events

Source: https://data-star.dev/guide

This example demonstrates receiving multiple `datastar-patch-elements` Server-Sent Events (SSE) in sequence. The client receives an initial update for the 'hal' div, waits for a few seconds, and then receives a subsequent update to reset the text. This pattern allows for timed or sequenced DOM manipulations.

```html
1event: datastar-patch-elements
2data: elements <div id="hal">
3data: elements     I’m sorry, Dave. I’m afraid I can’t do that.
4data: elements </div>
5
6event: datastar-patch-elements
7data: elements <div id="hal">
8data: elements     Waiting for an order...
9data: elements </div>
10

```

--------------------------------

### Go Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

A Go implementation for backend-initiated redirects using DataStar's SSE. It first patches elements to display a message, then waits, and finally executes a script to redirect the page.

```go
import (
	"time"
	"github.com/starfederation/datastar-go/datastar"
)

sse := datastar.NewSSE(w, r)
sse.PatchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`)
time.Sleep(3 * time.Second)
sse.ExecuteScript(`
    window.location = "/guide"
`)
```

--------------------------------

### DataStar Button Triggering Backend Action

Source: https://data-star.dev/guide/datastar_expressions

An HTML button element configured with DataStar's 'data-on:click' attribute to trigger a GET request to '/endpoint' upon being clicked. This demonstrates initiating backend actions directly from user interactions.

```html
<button data-on:click="@get('/endpoint')">
    What are you talking about, HAL?
</button>
```

--------------------------------

### Transform to Random SVG Shape with Datastar

Source: https://data-star.dev/examples/svg_morphing

Demonstrates morphing to different SVG shape types dynamically. This Go example uses Datastar SSE to replace the content of 'shape-demo' with a randomly selected SVG shape.

```go
svgMorphingRouter.Get("/shape_transform", func(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    shape := svgShapes[rand.N(len(svgShapes))]
    sse.PatchElements(fmt.Sprintf(`<svg id="shape-demo">%s</svg>`, shape))
})
```

--------------------------------

### Data-Star `data-indicator` for Loading States

Source: https://data-star.dev/guide/backend_requests

This snippet illustrates the use of the `data-indicator` attribute in Data-Star to manage loading states. When a request is in flight (triggered by `data-on:click` and `@get`), the specified signal (e.g., `fetching`) is set to `true`. This can be used to visually indicate loading with attributes like `data-class:loading`.

```html
<div id="question"></div>
<button
    data-on:click="@get('/actions/quiz')"
    data-indicator:fetching>
    Fetch a question
</button>
<div data-class:loading="$fetching" class="indicator"></div>
```

--------------------------------

### Preserve multiple attributes during DOM morphing (HTML)

Source: https://data-star.dev/reference

You can preserve multiple attributes by listing them separated by spaces in the `data-preserve-attr` value. This example preserves both the `open` and `class` attributes on a `<details>` element.

```html
<details open class="foo" data-preserve-attr="open class">
    <summary>Title</summary>
    Content
</details>
```

--------------------------------

### Send PUT Request with Datastar

Source: https://data-star.dev/reference/actions

Sends a PUT request to a specified URI. This function mirrors the functionality of `@get()` but exclusively uses the PUT HTTP method. Request configuration options are available.

```html
<button data-on:click="@put('/endpoint')"></button>
```

--------------------------------

### HTML: Editable Row Structure

Source: https://data-star.dev/examples/edit_row

This snippet shows the HTML structure of a table row when it is in an editable state. It replaces data cells with input fields and changes action buttons to 'Cancel' and 'Save'.

```html
<tr>
    <td>
        <input type="text" data-bind:name>
    </td>
    <td>
        <input type="text" data-bind:email>
    </td>
    <td>
        <button data-on:click="@get('/examples/edit_row/cancel')">
            Cancel
        </button>
        <button data-on:click="@patch('/examples/edit_row/0')">
            Save
        </button>
    </td>
</tr>
```

--------------------------------

### Set custom interval duration with data-on-interval (HTML)

Source: https://data-star.dev/reference

This example demonstrates how to use the `__duration` modifier with `data-on-interval` to set a custom interval, in this case, 500 milliseconds. This allows for more fine-grained control over the frequency of expression execution.

```html
<div data-on-interval__duration.500ms="$count++"></div>
```

--------------------------------

### C# ASP.NET Core: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This C# code snippet illustrates how to integrate Datastar into an ASP.NET Core application to stream Server-Sent Events (SSE). It shows how to add Datastar as a service and then use the `IDatastarService` to patch elements and signals asynchronously. This example is suitable for minimal API endpoints in ASP.NET Core.

```csharp
1using StarFederation.Datastar.DependencyInjection;
2
3// Adds Datastar as a service
4builder.Services.AddDatastar();
5
6app.MapGet("/", async (IDatastarService datastarService) =>
7{
8    // Patches elements into the DOM.
9    await datastarService.PatchElementsAsync(@"<div id=\"question\">What do you put in a toaster?</div>");
10
11    // Patches signals.
12    await datastarService.PatchSignalsAsync(new { response = "", answer = "bread" });
13});

```

--------------------------------

### Kotlin: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Kotlin snippet shows how to generate server-sent events with a dynamic interval duration. It uses `LocalDateTime` to get the current time and determines the duration based on the seconds.

```kotlin
1val now: LocalDateTime = currentTime()
 2val currentSeconds = now.second
 3val duration = if (currentSeconds < 50) 5 else 1
 4
 5val generator = ServerSentEventGenerator(response)
 6
 7generator.patchElements(
 8    elements =
 9        """
10        <div id="time" data-on-interval__duration.${duration}s="@get('/endpoint')">
11            $now
12        </div>
13        ".trimIndent()
14)
```

--------------------------------

### Rust: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Rust code snippet demonstrates how to stream Server-Sent Events (SSE) using the `datastar` crate. It utilizes the `async_stream` crate to yield `PatchElements` and `PatchSignals` events within an `Sse` stream. This example is suitable for asynchronous Rust applications.

```rust
1use datastar::prelude::*;
2use async_stream::stream;
3
4Sse(stream! {
5    // Patches elements into the DOM.
6    yield PatchElements::new("<div id='question'>What do you put in a toaster?</div>").into();
7
8    // Patches signals.
9    yield PatchSignals::new("{response: '', answer: 'bread'}").into();
10})

```

--------------------------------

### Animated SVG Morph Sequence with Datastar

Source: https://data-star.dev/examples/svg_morphing

An example of creating a smooth animation sequence by performing multiple SVG morphs in succession using Datastar SSE. This Go code updates a circle's properties with delays, creating a visual effect.

```go
svgMorphingRouter.Get("/animated_morph", func(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    
    // First morph
    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="30" fill="red" /></svg>`)
    time.Sleep(500 * time.Millisecond)
    
    // Second morph
    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="45" fill="orange" /></svg>`)
    time.Sleep(500 * time.Millisecond)
    
    // Third morph
    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="60" fill="yellow" /></svg>`)
    time.Sleep(500 * time.Millisecond)
    
    // Reset
    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="20" fill="green" /></svg>`)
})
```

--------------------------------

### Send DELETE Request with Datastar

Source: https://data-star.dev/reference/actions

Sends a DELETE request to a specified URI. This function is analogous to `@get()` but exclusively uses the DELETE HTTP method. It supports various request configuration options.

```html
<button data-on:click="@delete('/endpoint')"></button>
```

--------------------------------

### DRY Datastar Action with Templating Variable

Source: https://data-star.dev/how_tos/keep_datastar_code_dry

This example demonstrates how to make Datastar code DRY using a templating language's variable assignment. By defining the action once and referencing the variable, repetition is eliminated.

```html
{% set action = "@get('/endpoint')" %}
<button data-on:click="{{ action }}">Click me</button>
<button data-on:click="{{ action }}">No, click me!</button>
<button data-on:click="{{ action }}">Click us all!</button>
```

--------------------------------

### Load More Button and HTML Structure

Source: https://data-star.dev/how_tos/load_more_list_items

Demonstrates the HTML structure for a 'load more' button and a list container, utilizing `data-signals` for initial offset and `data-on:click` for triggering a backend request.

```APIDOC
## HTML Structure

### Description
This HTML sets up a list container (`#list`) and a button (`#load-more`). The button uses `data-signals:offset="1"` to initialize the offset and `data-on:click="@get('/how_tos/load_more/data')"` to trigger a GET request to the backend when clicked.

### Method
N/A (HTML)

### Endpoint
N/A (HTML)

### Request Example
```html
<div id="list">
  <div>Item 1</div>
</div>
<button id="load-more" 
        data-signals:offset="1" 
        data-on:click="@get('/how_tos/load_more/data')">
Click to load another item
</button>
```

### Response
N/A (HTML)
```

--------------------------------

### Replicating Datastar Pro Attributes with Free Datastar

Source: https://data-star.dev/essays

These HTML snippets demonstrate how to replicate two specific Datastar Pro attributes using the free version. The first replaces the current URL on load and when a '$page' variable changes, while the second scrolls an element into view.

```html
<!-- Replaces the current URL on load and whenever $page changes. -->
<div data-effect="window.history.replaceState({}, '', '/page/' + $page)"></div>

<!-- Scrolls the element into view. -->
<div data-init="el.scrollIntoView()"></div>
```

--------------------------------

### Use Alert Attribute in HTML

Source: https://data-star.dev/examples/custom_plugin

This HTML snippet shows how to utilize the custom 'alert' attribute plugin. The `data-alert` attribute is applied to a button, and it expects a string value that will be displayed when the button is clicked. This provides a declarative way to add alert functionality.

```html
1<button data-alert="'Hello from an attribute'">
2    Alert using an attribute
3</button>
```

--------------------------------

### HTML Input for Active Search with Debounce

Source: https://data-star.dev/examples/active_search

This HTML input element is configured to trigger an active search. The `data-bind:search` attribute suggests data binding for the input's value, while `data-on:input__debounce.200ms` specifies that an event listener should be attached to the 'input' event. This listener will execute a GET request to '/examples/active_search/search' after a 200ms delay of user inactivity, using the input's value as a search parameter.

```html
<input
    type="text"
    placeholder="Search..."
    data-bind:search
    data-on:input__debounce.200ms="@get('/examples/active_search/search')"
/>
```

--------------------------------

### PHP: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This PHP code snippet shows how to generate Server-Sent Events (SSE) using the Datastar PHP library. It initializes a `ServerSentEventGenerator` and then uses its `patchElements` and `patchSignals` methods to send data. This example is suitable for PHP backend implementations requiring SSE capabilities.

```php
1use starfederation\datastar\ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4$sse = new ServerSentEventGenerator();
5
6// Patches elements into the DOM.
7$sse->patchElements(
8    '<div id="question">What do you put in a toaster?</div>'
9);
10
11// Patches signals.
12$sse->patchSignals(['response' => '', 'answer' => 'bread']);

```

--------------------------------

### Datastar Event Bubbling with Dynamic IDs

Source: https://data-star.dev/how_tos/keep_datastar_code_dry

This advanced snippet demonstrates event bubbling in Datastar, similar to the previous example, but also includes capturing a `data-id` attribute from the clicked button. This allows for sending specific identifiers with the backend action.

```html
<div data-on:click="evt.target.tagName == 'BUTTON' 
    && ($id = evt.target.dataset.id)
    && @get('/endpoint')
">
    <button data-id="1">Click me</button>
    <button data-id="2">No, click me!</button>
    <button data-id="3">Click us all!</button>
</div>
```

--------------------------------

### Python (Sanic) - Redirect with SSE

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates implementing a redirect in Python using the Sanic framework and Datastar's SSE. It yields SSE events to patch elements, waits, and then executes a script for redirection.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response
import asyncio

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.execute_script('setTimeout(() => window.location = "/guide")')

```

--------------------------------

### Implement Alert Attribute Plugin

Source: https://data-star.dev/examples/custom_plugin

This JavaScript code defines a custom attribute plugin named 'alert'. It's configured to require a value and return a value, and its `apply` method adds a click event listener to the element. This listener triggers an alert with the value returned by an expression when the element is clicked.

```javascript
1attribute({
    name: 'alert',
    requirement: {
        key: 'denied',
        value: 'must',
    },
    returnsValue: true,
    apply({ el, rx }) {
        const callback = () => alert(rx())
        el.addEventListener('click', callback)
        return () => el.removeEventListener('click', callback)
    }
})
```

--------------------------------

### Clojure Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a backend redirect using Clojure and DataStar. It sends SSE events to update an indicator, pauses for 3 seconds, and then executes a script to redirect the user.

```clojure
(require
  '[starfederation.datastar.clojure.api :as d*]
  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]]
  '[some.hiccup.library :refer [html]])


(defn handle [ring-request]
  (->sse-response ring-request
    {on-open
      (fn [sse]
        (d*/patch-elements! sse
          (html [:div#indicator "Redirecting in 3 seconds..."]))
        (Thread/sleep 3000)
        (d*/execute-script! sse "window.location = \"/guide\"")
        (d*/close-sse! sse)}))
```

--------------------------------

### Preserve a single attribute during DOM morphing (HTML)

Source: https://data-star.dev/reference

The `data-preserve-attr` attribute ensures that specified attributes on an element are not overwritten during DOM updates. This example preserves the `open` attribute on a `<details>` element, maintaining its expanded state.

```html
<details open data-preserve-attr="open">
    <summary>Title</summary>
    Content
</details>
```

--------------------------------

### HTML Structure for DBMon

Source: https://data-star.dev/examples/dbmon

This HTML snippet defines the main container for the DBMon demo, including elements for displaying render time, input fields for mutation rate and FPS, and a table for database cluster information. It utilizes DataStar's custom attributes for dynamic updates and event handling.

```html
<div
    id="demo"
    data-init="@get('/examples/dbmon/updates')"
    data-signals:_editing__ifmissing="false"
>
    <p>
        Average render time for entire page: { renderTime }
    </p>
    <div role="group">
        <label>
            Mutation Rate %
            <input
                type="number"
                min="0"
                max="100"
                value="20"
                data-on:focus="$_editing = true"
                data-on:blur="@put('/examples/dbmon/inputs'); $_editing = false"
                data-attr:data-bind:mutation-rate="$_editing"
                data-attr:data-bind:_mutation-rate="!$_editing"
            />
        </label>
        <label>
            FPS
            <input
                type="number"
                min="1"
                max="144"
                value="60"
                data-on:focus="$_editing = true"
                data-on:blur="@put('/examples/dbmon/inputs'); $_editing = false"
                data-attr:data-bind:fps="$_editing"
                data-attr:data-bind:_fps="!$_editing"
            />
        </label>
    </div>
    <table style="table-layout: fixed; width: 100%; word-break: break-all">
        <tbody>
            <!-- Dynamic rows generated by server -->
            <tr>
                <td>cluster1</td>
                <td style="background-color: var(--_active-color)" class="success">
                    8
                </td>
                <td aria-description="SELECT blah from something">
                    12ms
                </td>
                <!-- More query cells... -->
            </tr>
            <!-- More database rows... -->
        </tbody>
    </table>
</div>
```

--------------------------------

### Filter signal patches by name using include regex (HTML)

Source: https://data-star.dev/reference

The `data-on-signal-patch-filter` attribute with the `include` property allows you to specify a regular expression to only watch for changes in signals whose names match the pattern. This example specifically targets signals named 'counter'.

```html
<!-- Only react to counter signal changes -->
<div data-on-signal-patch-filter="{include: /^counter$/}"></div>
```

--------------------------------

### Receive Patch Elements SSE Event

Source: https://data-star.dev/guide

This example shows how a client-side application can receive and process `datastar-patch-elements` Server-Sent Events (SSE). When such an event is received, the provided HTML data is used to morph elements in the DOM, enabling dynamic updates driven by the server. This is a fundamental part of Datastar's frontend interaction model.

```html
1event: datastar-patch-elements
2data: elements <div id="hal">
3data: elements     I’m sorry, Dave. I’m afraid I can’t do that.
4data: elements </div>
5

```

--------------------------------

### C#: Backend Logic for Loading More

Source: https://data-star.dev/how_tos/load_more_list_items

Provides a C# implementation for the 'load more' functionality using Datastar's SDK. It configures the web application, handles incoming requests, and uses `IDatastarService` to send patch and remove events.

```csharp
1using System.Text.Json;
2using StarFederation.Datastar;
3using StarFederation.Datastar.DependencyInjection;
4
5public class Program
6{
7    public record OffsetSignals(int offset);
8
9    public static void Main(string[] args)
10    {
11        var builder = WebApplication.CreateBuilder(args);
12        builder.Services.AddDatastar();
13        var app = builder.Build();
14
15        app.MapGet("/more", async (IDatastarService datastarService) =>
16        {
17            var max = 5;
18            var limit = 1;
19            var signals = await datastarService.ReadSignalsAsync<OffsetSignals>();
20            var offset = signals.offset;
21            if (offset < max)
22            {
23                var newOffset = offset + limit;
24                await datastarService.PatchElementsAsync($"<div>Item {newOffset}</div>", new()
25                {
26                    Selector = "#list",
27                    PatchMode = PatchElementsMode.Append,
28                });
29                if (newOffset < max)
30                    await datastarService.PatchSignalsAsync(new OffsetSignals(newOffset));
31                else
32                    await datastarService.RemoveElementAsync("#load-more");
33            }
34        });
35
36        app.Run();
37    }
38}
```

--------------------------------

### JavaScript Web Component for String Reversal

Source: https://data-star.dev/examples/web_component

This JavaScript code defines a custom HTML element `ReverseComponent` that observes changes to its 'name' attribute. Upon detection of a change, it reverses the attribute's value and dispatches a 'reverse' custom event with the reversed string in its details.

```javascript
class ReverseComponent extends HTMLElement {
    static get observedAttributes() {
        return ["name"];
    }

    attributeChangedCallback(name, oldValue, newValue) {
        const len = newValue.length;
        let value = Array(len);
        let i = len - 1;
        for (const char of newValue) {
            value[i--] = char.codePointAt(0);
        }
        value = String.fromCodePoint(...value);
        this.dispatchEvent(new CustomEvent("reverse", { detail: { value } }));
    }
}

customElements.define("reverse-component", ReverseComponent);
```

--------------------------------

### Change Circle Color with Datastar

Source: https://data-star.dev/examples/svg_morphing

An example using Datastar's SSE to morph an SVG circle's color dynamically. It selects a random color from a predefined list and applies it to the SVG element identified by 'circle-demo'.

```go
svgMorphingRouter.Get("/circle_color", func(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    color := svgColors[rand.N(len(svgColors))]
    sse.PatchElements(fmt.Sprintf(`<svg id="circle-demo"><circle cx="50" cy="50" r="40" fill="%s" /></svg>`, color))
})
```

--------------------------------

### Datastar Event Bubbling for Single Listener

Source: https://data-star.dev/how_tos/keep_datastar_code_dry

This example utilizes event bubbling in Datastar to attach a single event listener to a parent element. It checks if the target is a button before executing the backend action, reducing the number of listeners.

```html
<div data-on:click="evt.target.tagName == 'BUTTON' 
    && @get('/endpoint')
">
    <button>Click me</button>
    <button>No, click me!</button>
    <button>Click us all!</button>
</div>
```

--------------------------------

### Keep SSE Connection Open When Page is Hidden

Source: https://data-star.dev/how_tos/prevent_sse_connections_closing

This example demonstrates how to configure an SSE connection to remain open even when the browser tab is inactive. It utilizes a custom attribute `data-on:click` to trigger an API call with the `openWhenHidden: true` option.

```html
<button data-on:click="@get('/endpoint', {openWhenHidden: true})"></button>
```

--------------------------------

### Go - Redirect with Server-Sent Events

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates performing a redirect using Datastar's SSE in Go. It involves patching HTML elements, pausing for a duration, and then executing a JavaScript script to redirect.

```go
import (
    "time"
    "github.com/starfederation/datastar-go/datastar"
)

sse := datastar.NewSSE(w, r)
sse.PatchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`)
time.Sleep(3 * time.Second)
sse.ExecuteScript(`
    setTimeout(() => window.location = "/guide")
`)

```

```go
import (
    "time"
    "github.com/starfederation/datastar-go/datastar"
)

sse := datastar.NewSSE(w, r)
sse.PatchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`)
time.Sleep(3 * time.Second)
sse.Redirect("/guide")

```

--------------------------------

### Filter signal patches by name using exclude regex (HTML)

Source: https://data-star.dev/reference

Using `data-on-signal-patch-filter` with the `exclude` property, you can prevent the event listener from triggering for signals matching a given regular expression. This example excludes changes to signals ending with 'changes'.

```html
<!-- React to all changes except those ending with "changes" -->
<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>
```

--------------------------------

### Trigger Alert on 'Ctrl + L' Keydown (HTML)

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

This example demonstrates triggering an alert when the 'Ctrl' and 'L' keys are pressed simultaneously. It checks both `evt.ctrlKey` and `evt.key` properties within the `data-on:keydown__window` attribute to identify the specific key combination. The `evt` object is crucial for accessing these event details.

```html
<div data-on:keydown__window="evt.ctrlKey && evt.key === 'l' && alert('Key pressed')"></div>
```

--------------------------------

### HTML: Load More Button and List Container

Source: https://data-star.dev/how_tos/load_more_list_items

Sets up the initial HTML structure with a list container and a 'load more' button. The button uses data attributes to manage the offset and trigger a backend request on click.

```html
1<div id="list">
2<div>Item 1</div>
3</div>
4<button id="load-more" 
5        data-signals:offset="1" 
6        data-on:click="@get('/how_tos/load_more/data')">
7Click to load another item
8</button>
```

--------------------------------

### POST /websites/data-star_dev

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This endpoint handles server-sent events, including patching HTML elements, executing scripts, and closing the SSE connection.

```APIDOC
## POST /websites/data-star_dev

### Description
Handles server-sent events to update the client. This includes patching HTML elements, executing scripts, and closing the SSE connection.

### Method
POST

### Endpoint
/websites/data-star_dev

### Parameters
#### Request Body
- **sse** (object) - Required - An object representing the server-sent event stream.

### Request Example
```clojure
(d*/patch-elements! sse
  (html [:div#indicator "Redirecting in 3 seconds..."]))
(Thread/sleep 3000)
(d*/execute-script! sse
  "setTimeout(() => window.location = \"/guide\""
(d*/close-sse! sse))
```

### Response
#### Success Response (200)
An SSE stream is initiated and maintained.

#### Response Example
(No specific JSON response body, as it's an SSE stream)
```

--------------------------------

### HTML: Nesting Signals using Dot Notation and Object Syntax

Source: https://data-star.dev/guide/backend_requests

Demonstrates how to nest signals in HTML using both dot notation and object syntax for granular targeting. Useful for managing complex state.

```html
<div data-signals:foo.bar="1"></div>
<div data-signals="{foo: {bar: 1}}"></div>
```

--------------------------------

### Ruby Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

A Ruby implementation for backend-initiated redirects using DataStar. It streams SSE events to patch elements, pauses, and then executes a script for redirection.

```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'
  sleep 3
  sse.execute_script 'window.location = "/guide"'
end
```

--------------------------------

### Import Datastar using npm, Deno, or Bun

Source: https://data-star.dev/guide

Import Datastar directly into your project using a package manager. This is suitable for projects managed with npm, Deno, or Bun.

```javascript
// @ts-expect-error (only required for TypeScript projects)
import 'https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js'
```

--------------------------------

### HTML structure for Event Bubbling Demo

Source: https://data-star.dev/examples/event_bubbling

This HTML sets up a div with a container for buttons. It includes a span to display the key pressed and a parent div with a `data-on:click` attribute to handle button clicks. The `pointer-events: none;` style is applied to the button container.

```html
<div id="demo">
    Key pressed: <span data-text="$key"></span>
    <div id="button-container" data-on:click="$key = evt.target.dataset.id">
        <button data-id="KEY ELSE" class="gray">KEY<br/>ELSE</button>
        <button data-id="CM">CM</button>
        <button data-id="OM">OM</button>
        <button data-id="FETCH">FETCH</button>
        <button data-id="SET">SET</button>
        <button data-id="EXEC">EXEC</button>
        <button data-id="TEST ALARM" class="gray">TEST<br/>ALARM</button>
        <button data-id="3">3</button>
        <button data-id="2">2</button>
        <button data-id="1">1</button>
        <button data-id="ENTER">ENTER</button>
        <button data-id="CLEAR">CLEAR</button>
    </div>
</div>

<style>
    #button-container {
        pointer-events: none;
    }
</style>
```

--------------------------------

### DataStar Interval: Run Expression Every 5 Seconds

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This snippet demonstrates setting up an interval to run an expression every 5 seconds using the `data-on-interval__duration.5s` attribute. It shows a basic implementation without immediate execution.

```html
<div id="time"
     data-on-interval__duration.5s="@get('/endpoint')"
></div>
```

--------------------------------

### Rust - Redirect with SSE and Tokio

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a redirect in Rust using Datastar's SSE and the Tokio runtime for asynchronous operations. It yields events to patch elements, sleeps, and then executes a redirect script.

```rust
use datastar::prelude::*;
use async_stream::stream;
use core::time::Duration;

Sse(stream! {
    yield PatchElements::new("<div id='indicator'>Redirecting in 3 seconds...</div>").into();
    tokio::time::sleep(core::time::Duration::from_secs(3)).await;
    yield ExecuteScript::new("setTimeout(() => window.location = '/guide')").into();
});

```

--------------------------------

### Python Backend Redirect with DataStar (Sanic)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Python snippet demonstrates a backend redirect using DataStar with the Sanic framework. It yields SSE events to update the UI and then execute a redirect script.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.execute_script('window.location = "/guide"')
```

--------------------------------

### Kotlin Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Kotlin code snippet shows how to perform a backend redirect using DataStar. It updates the DOM with an indicator, pauses execution, and then runs a script to change the window location.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements =
        """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent(),
)

Thread.sleep(3 * ONE_SECOND)

generator.executeScript(
    script = "window.location.href = '/success'",
)
```

--------------------------------

### Capture All Signal Changes with Signal Patch

Source: https://data-star.dev/examples/on_signal_patch

This HTML section demonstrates capturing all signal patches, excluding those specifically related to 'allChanges' and 'counterChanges' themselves, using `data-on-signal-patch` and a filter. The captured changes are displayed in a preformatted tag.

```html
<div
    data-on-signal-patch="$allChanges.push(patch)"
    data-on-signal-patch-filter="{exclude: /allChanges|counterChanges/}"
>
    <h3>All Signal Changes</h3>
    <pre data-json-signals__terse="{include: /^allChanges/}"></pre>
</div>
```

--------------------------------

### Kotlin - Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Shows how to implement a redirect in Kotlin using the ServerSentEventGenerator. It patches elements, pauses, and then redirects.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent(),
)

Thread.sleep(3 * ONE_SECOND)

generator.executeScript(
    script = "setTimeout(() => window.location = '/guide')",
)

```

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements =
        """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent(),
)

Thread.sleep(3 * ONE_SECOND)

generator.redirect("/guide")

```

--------------------------------

### PHP - Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Illustrates redirecting using the ServerSentEventGenerator in PHP. It involves patching elements, pausing, and then executing a script for redirection.

```php
use starfederation\datastar\ServerSentEventGenerator;

$sse = new ServerSentEventGenerator();
$sse->patchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`);
sleep(3);
$sse->executeScript(`
    setTimeout(() => window.location = "/guide")
`);

```

--------------------------------

### Go: SSE `PatchElements` Implementation

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Go code demonstrates sending Server-Sent Events to patch elements. It generates the current time and embeds it within an HTML `div` that is configured for a 5-second update interval.

```go
import (
    "time"
    "github.com/starfederation/datastar-go/datastar"
)

currentTime := time.Now().Format("2006-01-02 15:04:05")

sse := datastar.NewSSE(w, r)
sse.PatchElements(fmt.Sprintf(`
    <div id="time" data-on-interval__duration.5s="@get('/endpoint')">
        %s
    </div>
`, currentTime))

```

--------------------------------

### Kotlin: Server-Side Logic for Loading More

Source: https://data-star.dev/how_tos/load_more_list_items

A Kotlin snippet demonstrating the server-side handling of Datastar signals for a 'load more' feature. It deserializes signals, determines the next offset, and uses `ServerSentEventGenerator` to send UI update events.

```kotlin
1@Serializable
2data class OffsetSignals(
3    val offset: Int,
4)
5
6val signals =
7    readSignals(
8        request,
9        { json: String -> Json.decodeFromString<OffsetSignals>(json) },
10    )
11
12val max = 5
13val limit = 1
14val offset = signals.offset
15
16val generator = ServerSentEventGenerator(response)
17
18if (offset < max) {

```

--------------------------------

### C#: Reading Nested Signals with DataStar Service

Source: https://data-star.dev/guide/backend_requests

Demonstrates how to read nested signals like 'foo.bar' on the backend using C# and the DataStar service. It includes setting up DataStar as a service and deserializing signals from requests.

```csharp
using StarFederation.Datastar.DependencyInjection;

// Adds Datastar as a service
builder.Services.AddDatastar();

public record Signals
{
    [JsonPropertyName("foo")] [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public FooSignals? Foo { get; set; } = null;

    public record FooSignals
    {
        [JsonPropertyName("bar")] [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
        public string? Bar { get; set; }
    }
}

app.MapGet("/read-signals", async (IDatastarService datastarService) =>
{
    Signals? mySignals = await datastarService.ReadSignalsAsync<Signals>();
    var bar = mySignals?.Foo?.Bar;
});
```

--------------------------------

### Clojure: Server-Side Logic for Loading More

Source: https://data-star.dev/how_tos/load_more_list_items

Implements the backend logic in Clojure using Datastar's API. It handles incoming signals, generates new list items, and sends events to append elements or remove the load more button.

```clojure
1(require
2  '[starfederation.datastar.clojure.api :as d*]
3  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]]
4  '[some.hiccup.library :refer [html]]
5  '[some.json.library :refer [read-json-str write-json-str]]))
6
7
8(def max-offset 5)
9
10(defn handler [ring-request]
11  (->sse-response ring-request
12    {on-open
13     (fn [sse]
14       (let [d*-signals (-> ring-request d*/get-signals read-json-str)
15             offset (get d*-signals "offset")
16             limit 1
17             new-offset (+ offset limit)]
18
19         (d*/patch-elements! sse
20                             (html [:div "Item " new-offset])
21                             {d*/selector   "#list"
22                              d*/merge-mode d*/mm-append})
23
24         (if (< new-offset max-offset)
25           (d*/patch-signals! sse (write-json-str {"offset" new-offset}))
26           (d*/remove-fragment! sse "#load-more"))
27
28         (d*/close-sse! sse)))})

```

--------------------------------

### Ruby - Redirect with Datastar Stream

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Shows how to perform a redirect in Ruby using Datastar's stream functionality. It patches elements, sleeps, and then executes a JavaScript redirect script.

```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'

  sleep 3

  sse.execute_script <<~JS
    setTimeout(() => {
      window.location = '/guide'
    })
  JS
end

```

--------------------------------

### Backend SSE Events for Redirect (Conceptual)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Conceptual representation of backend Server-Sent Events (SSE) to first update a visible indicator and then append a script tag to redirect the page after a delay.

```sse
event: datastar-patch-elements
data: elements <div id="indicator">Redirecting in 3 seconds...</div>

// Wait 3 seconds

event: datastar-patch-elements
data: selector body
data: mode append
data: elements <script>window.location.href = "/guide"</script>


```

--------------------------------

### Patch HTML Elements via SSE in Go

Source: https://data-star.dev/guide

This Go code snippet illustrates how to patch HTML elements into the DOM using Datastar's Server-Sent Events (SSE) functionality. It initializes an `SSE` generator and uses the `PatchElements` method to send updates, including a one-second delay between them. This is useful for real-time web updates.

```go
1import (
2    "github.com/starfederation/datastar-go/datastar"
3    time
4)
5
6// Creates a new `ServerSentEventGenerator` instance.
7sse := datastar.NewSSE(w,r)
8
9// Patches elements into the DOM.
10sse.PatchElements(
11    `<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`
12)
13
14time.Sleep(1 * time.Second)
15
16sse.PatchElements(
17    `<div id="hal">Waiting for an order...</div>`
18)
```

--------------------------------

### Response Handling

Source: https://data-star.dev/reference/actions

Details how Datastar automatically handles different response content types and specific headers for `text/html` responses.

```APIDOC
## Response Handling

Datastar automatically handles the following response content types:

*   **`text/event-stream`**: Standard SSE events.
*   **`text/html`**: HTML elements to patch into the DOM.
*   **`application/json`**: JSON encoded signals for patching.
*   **`text/javascript`**: JavaScript code for browser execution.

### `text/html` Response Headers:

When returning `text/html`, the server can include these optional response headers:

*   **`datastar-selector`** (string): A CSS selector for the target DOM elements.
*   **`datastar-mode`** (string): Specifies the patching method (`outer`, `inner`, `remove`, `replace`, `prepend`, `append`, `before`, `after`). Defaults to `outer`.
*   **`datastar-use-view-transition`** (boolean): Enables the View Transition API for patching.

### Example Server-Side `text/html` Response:
```javascript
response.headers.set('Content-Type', 'text/html');
response.headers.set('datastar-selector', '#my-element');
response.headers.set('datastar-mode', 'inner');
response.body = '<p>New content</p>';
```
```

--------------------------------

### PHP: Reading Signals using ServerSentEventGenerator

Source: https://data-star.dev/guide/backend_requests

Demonstrates how to read all signals from an incoming request in PHP using the `ServerSentEventGenerator` class. This is a straightforward method for accessing frontend signal data.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Reads all signals from the request.
$signals = ServerSentEventGenerator::readSignals();
```

--------------------------------

### Clojure: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Clojure code snippet demonstrates how to create a Server-Sent Events (SSE) response within a Ring handler using the Datastar Clojure SDK. It shows how to set up an SSE response and then patch elements and signals into the DOM. This approach simplifies SSE implementation by using the provided SDK to handle headers and event formatting.

```clojure
1;; Import the SDK's api and your adapter
2(require
3 '[starfederation.datastar.clojure.api :as d*]
4 '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
5
6;; in a ring handler
7(defn handler [request]
8  ;; Create an SSE response
9  (->sse-response request
10                  {
11                   on-open
12                   (fn [sse]
13                     ;; Patches elements into the DOM
14                     (d*/patch-elements! sse
15                                         "<div id=\"question\">What do you put in a toaster?</div>")
16
17                     ;; Patches signals
18                     (d*/patch-signals! sse "{response: '', answer: 'bread'}"))}))

```

--------------------------------

### C# - Redirect with Datastar Service

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a redirect using Datastar's IDatastarService in C#. It patches HTML elements, delays execution, and then redirects the user to a new page.

```csharp
using StarFederation.Datastar.DependencyInjection;

app.MapGet("/redirect", async (IDatastarService datastarService) =>
{
    await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
    await Task.Delay(TimeSpan.FromSeconds(3));
    await datastarService.ExecuteScriptAsync("setTimeout(() => window.location = \"/guide\")");
});

```

```csharp
using StarFederation.Datastar.DependencyInjection;
using StarFederation.Datastar.Scripts;

app.MapGet("/redirect", async (IDatastarService datastarService) =>
{
    await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
    await Task.Delay(TimeSpan.FromSeconds(3));
    await datastarService.Redirect("/guide");
});

```

--------------------------------

### Node.js - Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates implementing a redirect in Node.js using the http module and Datastar's ServerSentEventGenerator. It streams SSE events to patch elements and trigger a redirect after a delay.

```javascript
import { createServer } from "node:http";
import { ServerSentEventGenerator } from "../npm/esm/node/serverSentEventGenerator.js";

const server = createServer(async (req, res) => {

  ServerSentEventGenerator.stream(req, res, async (sse) => {
    sse.patchElements(`
      <div id="indicator">Redirecting in 3 seconds...</div>
    `);

    setTimeout(() => {
      sse.executeScript(`setTimeout(() => window.location = "/guide")`);
    }, 3000);
  });
});

```

--------------------------------

### Load More Data with Append and Signal Patching in Python (FastAPI)

Source: https://data-star.dev/how_tos/load_more_list_items

This Python snippet, designed for a FastAPI application, demonstrates how to implement a 'load more' functionality using Data Star. It yields Server-Sent Events to append new items, update the offset signal, and remove the load more button when all items are displayed.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.consts import ElementPatchMode
from datastar_py.fastapi import datastar_response, ReadSignals

MAX_ITEMS = 5

@app.get("/how_tos/load_more/data")
@datastar_response
async def load_data(signals: ReadSignals):
    if signals["offset"] < MAX_ITEMS:
        new_offset = signals["offset"] + 1
        yield SSE.patch_elements(
            f"<div>Item {new_offset}</div>",
            mode=ElementPatchMode.APPEND,
            selector="#list"
        )
        if new_offset < MAX_ITEMS:
            yield SSE.patch_signals({"offset": new_offset})
        else:
            yield SSE.remove_elements("#load-more")
```

--------------------------------

### JSON Output for Signal Patch Data

Source: https://data-star.dev/examples/on_signal_patch

These JSON snippets represent the expected structure for displaying signal changes. The first shows an empty array for `counterChanges`, and the second shows an empty array for `allChanges`, typically seen before any patching occurs or after a clear operation.

```json
{"counterChanges":[]}
```

```json
{"allChanges":[]}
```

--------------------------------

### Go: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Go snippet demonstrates how to create server-sent events with a dynamic interval duration using the `datastar-go` library. It formats the current time and calculates the duration based on the current seconds.

```go
1import (
 2    "time"
 3    "github.com/starfederation/datastar-go/datastar"
 4)
 5
 6currentTime := time.Now().Format("2006-01-02 15:04:05")
 7currentSeconds := time.Now().Format("05")
 8duration := 1
 9if currentSeconds < "50" {
10    duration = 5
11}
12
13sse := datastar.NewSSE(w, r)
14sse.PatchElements(fmt.Sprintf(`
15    <div id="time" data-on-interval__duration.%ds="@get('/endpoint')">
16        %s
17    </div>
18`, duration, currentTime))
```

--------------------------------

### Go SDK for SSE Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This Go code snippet shows how to use the DataStar Go SDK to create a Server-Sent Events (SSE) generator and patch signals with a delay between updates.

```go
 1import (
 2    "github.com/starfederation/datastar-go/datastar"
 3)
 4
 5// Creates a new `ServerSentEventGenerator` instance.
 6sse := datastar.NewSSE(w, r)
 7
 8// Patches signals
 9sse.PatchSignals([]byte(`{hal: 'Affirmative, Dave. I read you.'}`))
10
11time.Sleep(1 * time.Second)
12
13sse.PatchSignals([]byte(`{hal: '...'}`))
```

--------------------------------

### Backend Event Handling for Appending Data

Source: https://data-star.dev/how_tos/load_more_list_items

Explains how the backend should respond with `datastar-patch-elements` to append new items to the list and `datastar-patch-signals` to update the offset. It also covers removing the button when all items are loaded.

```APIDOC
## Backend Response for Loading More Items

### Description
When the backend receives a request to load more items, it should send `datastar-patch-elements` with `mode: append` to add the new element to the `#list` container. It then sends `datastar-patch-signals` to update the offset. If the maximum number of items is reached, the `#load-more` button is removed.

### Method
N/A (Backend Events)

### Endpoint
N/A (Backend Events)

### Response
#### Success Response (200)
- **`event: datastar-patch-elements`**
  - **`data: selector #list`**: Specifies the container to patch.
  - **`data: mode append`**: Indicates that new elements should be appended.
  - **`data: elements <div>Item 2</div>`**: The HTML for the new item.
- **`event: datastar-patch-signals`**
  - **`data: signals {offset: 2}`**: Updates the client-side signal for the offset.
- **`event: datastar-patch-elements`**
  - **`data: selector #load-more`**: Specifies the element to remove.
  - **`data: mode remove`**: Indicates that the element should be removed.

### Response Example
```
event: datastar-patch-elements
data: selector #list
data: mode append
data: elements <div>Item 2</div>

event: datastar-patch-signals
data: signals {offset: 2}


# (Or if max items reached)
event: datastar-patch-elements
data: selector #load-more
data: mode remove
```
```

--------------------------------

### Go: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Go code demonstrates how to generate Server-Sent Events (SSE) using the Datastar Go SDK. It shows the creation of a `ServerSentEventGenerator` instance and its usage for patching elements and signals. This method is useful for directly handling HTTP requests and responses to stream SSE.

```go
1import ("github.com/starfederation/datastar-go/datastar")
2
3// Creates a new `ServerSentEventGenerator` instance.
4sse := datastar.NewSSE(w,r)
5
6// Patches elements into the DOM.
7sse.PatchElements(
8    `<div id="question">What do you put in a toaster?</div>`
9)
10
11// Patches signals.
12sse.PatchSignals([]byte(`{response: '', answer: 'bread'}`))

```

--------------------------------

### Patch HTML Elements via SSE in C#

Source: https://data-star.dev/guide

This C# code demonstrates patching HTML elements into the DOM using Datastar's `IDatastarService`. It shows how to register Datastar as a service and then use `PatchElementsAsync` to send updates, including a delay between them. This is suitable for ASP.NET Core applications.

```csharp
1using StarFederation.Datastar.DependencyInjection;
2
3// Adds Datastar as a service
4builder.Services.AddDatastar();
5
6app.MapGet("/", async (IDatastarService datastarService) =>
7{
8    // Patches elements into the DOM.
9    await datastarService.PatchElementsAsync(@"<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>");
10
11    await Task.Delay(TimeSpan.FromSeconds(1));
12
13    await datastarService.PatchElementsAsync(@"<div id=\"hal\">Waiting for an order...</div>");
14});
```

--------------------------------

### Datastar Patch Elements: Advanced Options

Source: https://data-star.dev/reference/sse_events

Illustrates advanced options for patching elements, including 'inner' morphing, using CSS selectors, enabling view transitions, and handling multi-line element content. This allows for more granular control over DOM updates.

```html
event: datastar-patch-elements
data: mode inner
data: selector #foo
data: useViewTransition true
data: elements <div>
data: elements        Hello world!
data: elements </div>


```

--------------------------------

### Rust Backend Redirect with DataStar

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Rust code uses DataStar to perform a backend redirect. It generates SSE events to update the page content and execute a JavaScript redirect after a specified delay.

```rust
use datastar::prelude::*;
use async_stream::stream;
use core::time::Duration;

Sse(stream! { //
    yield PatchElements::new("<div id='indicator'>Redirecting in 3 seconds...</div>").into();
    tokio::time::sleep(core::time::Duration::from_secs(3)).await;
    yield ExecuteScript::new("window.location = '/guide'").into();
});
```

--------------------------------

### Python (FastAPI): Reading Signals with DataStar

Source: https://data-star.dev/guide/backend_requests

Shows how to read signals from a FastAPI request in Python using the DataStar library. The `@datastar_response` decorator and `read_signals` function simplify signal retrieval.

```python
from datastar_py.fastapi import datastar_response, read_signals

@app.get("/updates")
@datastar_response
async def updates(request: Request):
    # Retrieve a dictionary with the current state of the signals from the frontend
    signals = await read_signals(request)
```

--------------------------------

### Datastar Expression: Basic Signal Evaluation

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates a basic Datastar expression where a signal '$foo' with an initial value of '1' is displayed in a `data-text` attribute. This shows how Datastar directly evaluates signal values.

```html
1<div data-signals:foo="1">
2    <div data-text="$foo"></div>
3</div>
```

--------------------------------

### Ruby: `Datastar` for Patching Elements

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Ruby snippet shows how to use the `Datastar` library to send element patches via Server-Sent Events. It formats the current time and injects it into an HTML structure for dynamic updates.

```ruby
datastar = Datastar.new(request:, response:)

current_time = Time.now.strftime('%Y-%m-%d %H:%M:%S')

datastar.patch_elements <<~FRAGMENT
    <div id="time"
         data-on-interval__duration.5s="@get('/endpoint')"
    >
        #{current_time}
    </div>
FRAGMENT

```

--------------------------------

### Datastar JS Options

Source: https://data-star.dev/reference/actions

Details the available options for Datastar JS actions, including contentType, filterSignals, selector, headers, openWhenHidden, and retry configurations.

```APIDOC
## Datastar JS Options

All Datastar actions accept an `options` argument for customization.

### Options Available:

*   **`contentType`** (string) - Specifies the type of content to send. Accepts `json` (default) or `form`. If `form`, it validates and sends form elements of the closest form.
*   **`filterSignals`** (object) - Filters signals. Accepts an `include` property (RegExp, defaults to `/.*/`) and an optional `exclude` property (RegExp, defaults to `/(^_|._).*/`).
*   **`selector`** (string|null) - Specifies a form to send when `contentType` is `form`. Defaults to `null` (uses the closest form).
*   **`headers`** (object) - An object containing custom headers for the request.
*   **`openWhenHidden`** (boolean) - Keeps the connection open when the page is hidden. Defaults to `false`.
*   **`retry`** (string) - Determines retry behavior. Options: `'auto'` (default), `'error'`, `'always'`, `'never'`.
*   **`retryInterval`** (number) - The retry interval in milliseconds. Defaults to `1000`.
*   **`retryScaler`** (number) - Multiplier for scaling retry wait times. Defaults to `2`.
*   **`retryMaxWaitMs`** (number) - Maximum wait time between retries in milliseconds. Defaults to `30000`.
*   **`retryMaxCount`** (number) - Maximum number of retry attempts. Defaults to `10`.
*   **`requestCancellation`** (string|object) - Controls request cancellation. Options: `'auto'` (default), `'disabled'`, or an `AbortController` instance.

### Request Example (GET with options):
```html
<button data-on:click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})"></button>
```
```

--------------------------------

### Clojure: Server-Sent Events with `patch-elements!`

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Clojure code demonstrates how to generate Server-Sent Events (SSE) to patch elements on the client side. It uses `datastar.clojure.api` to send an updated time within an HTML `div` at a 5-second interval.

```clojure
(require
  '[starfederation.datastar.clojure.api :as d*]
  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
  '[some.hiccup.library :refer [html]])

(import
  'java.time.format.DateTimeFormatter
  'java.time.LocalDateTime)

(def formatter (DateTimeFormatter/ofPattern "YYYY-MM-DD HH:mm:ss"))

(defn handle [ring-request]
   (->sse-response ring-request
     {on-open
      (fn [sse]
        (d*/patch-elements! sse
          (html [:div#time {:data-on-interval__duration.5s (d*/sse-get "/endpoint")} 
                  (LocalDateTime/.format (LocalDateTime/now) formatter)])))
        (d*/close-sse! sse)}))

```

--------------------------------

### Copy Text to Clipboard with DataStar

Source: https://data-star.dev/reference/actions

Utilizes the `@clipboard()` Pro Action to copy text to the system clipboard. It supports plain text and Base64 encoded text, which is decoded before copying.

```html
<!-- Copy plain text -->
<button data-on:click="@clipboard('Hello, world!')"></button>

<!-- Copy base64 encoded text (will decode before copying) -->
<button data-on:click="@clipboard('SGVsbG8sIHdvcmxkIQ==', true)"></button>
```

--------------------------------

### HTML Structure for DataStar Signals

Source: https://data-star.dev/guide/reactive_signals

This HTML snippet demonstrates how to set up elements to receive and display patched signals from the backend using DataStar's custom attributes.

```html
1<div data-signals:hal="'...'" >
2    <button data-on:click="@get('/endpoint')">
3        HAL, do you read me?
4    </button>
5    <div data-text="$hal"></div>
6</div>
```

--------------------------------

### Patch HTML Elements via SSE in Java

Source: https://data-star.dev/guide

This Java code demonstrates patching HTML elements into the DOM using Datastar's `ServerSentEventGenerator`. It sets up the generator with an `HttpServletResponseAdapter` and uses the `PatchElements.builder()` to send updates, including a one-second delay between them. This is typically used within a Java web application context.

```java
1import starfederation.datastar.utils.ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
6
7// Patches elements into the DOM.
8generator.send(PatchElements.builder()
9    .data("<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
10    .build()
11);
12
13Thread.sleep(1000);
14
15generator.send(PatchElements.builder()
16    .data("<div id=\"hal\">Waiting for an order...</div>")
17    .build()
18);
```

--------------------------------

### Datastar Pro Actions

Source: https://data-star.dev/reference/actions

Additional actions available in Datastar Pro, offering extended functionality.

```APIDOC
## Pro Actions

### Description
Datastar Pro extends functionality with additional helpful actions.

### Methods
- `@clipboard()`
- `@fit()`

### Endpoint
N/A (Used within Datastar expressions)

### Parameters
(Specific parameters depend on the action. Refer to Datastar Pro documentation for details.)

### Request Example
```html
<!-- Example using @clipboard() -->
<button data-on:click="@clipboard('Copy this text')">Copy</button>
```

### Response
(Responses will vary based on the action.)
```

--------------------------------

### Generate SSE Stream in Rust with Datastar

Source: https://data-star.dev/guide/reactive_signals

This Rust code snippet demonstrates how to create a Server-Sent Event (SSE) stream using the `async_stream` crate and Datastar. It yields `PatchSignals` at intervals, simulating real-time updates. Dependencies include `async_stream`, `datastar`, and `std::thread` for pausing execution.

```rust
use async_stream::stream;
use datastar::prelude::*;
use std::thread;
use std::time::Duration;

Sse(stream! {
    // Patches signals.
    yield PatchSignals::new("{hal: 'Affirmative, Dave. I read you.'}").into();

    thread::sleep(Duration::from_secs(1));
    
    yield PatchSignals::new("{hal: '...'});
})
```

--------------------------------

### Create Loading Indicators with data-indicator

Source: https://data-star.dev/reference/attributes

The `data-indicator` attribute creates a signal that is `true` when a fetch request is in flight and `false` otherwise. This signal can be used to display loading states, disable elements, etc.

```html
<button data-on:click="@get('/endpoint')"
        data-indicator:fetching
></button>
```

```html
<button data-on:click="@get('/endpoint')"
        data-indicator:fetching
        data-attr:disabled="$fetching"
></button>
<div data-show="$fetching">Loading...</div>
```

```html
<button data-indicator="fetching"></button>
```

```html
<div data-indicator:fetching data-init="@get('/endpoint')"></div>
```

--------------------------------

### Datastar Patch Signals: Basic Update

Source: https://data-star.dev/reference/sse_events

Demonstrates the basic usage of 'datastar-patch-signals' to update signal values. A 'signals' line provides a JSON object mapping signal names to their new values.

```html
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}


```

--------------------------------

### Send Multiple SSE Events in Python using Datastar

Source: https://data-star.dev/guide/backend_requests

Shows how to send SSE events, such as patching HTML elements and sending signal data, using Python with the Datastar framework. This is applicable for Python backend development.

```python
generator.patchElements(
    elements = """<div id=\"question\">...</div>""",
)
generator.patchElements(
    elements = """<div id=\"instructions\">...</div>""",
)
generator.patchSignals(
    signals = "{\"answer\": \"...\", \"prize\": \"...\"}",
)
```

--------------------------------

### Send Multiple SSE Events in Ruby (Streaming) using Datastar

Source: https://data-star.dev/guide/backend_requests

Illustrates sending SSE events, including patching elements and signals, using Ruby's streaming capabilities with Datastar. This approach is efficient for real-time updates.

```ruby
datastar.stream do |sse|
  sse.patch_elements('<div id="question">...</div>')
  sse.patch_elements('<div id="instructions">...</div>')
  sse.patch_signals(answer: '...', prize: '...')
end
```

--------------------------------

### PHP: Backend Redirect with Server-Sent Events

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This snippet shows how to use ServerSentEventGenerator in PHP to patch an element and then redirect the user after a delay. It relies on the ServerSentEventGenerator class.

```php
$sse = new ServerSentEventGenerator();
$sse->patchElements('
    <div id="indicator">Redirecting in 3 seconds...</div>
');
sleep(3);
$sse->location('/guide');
```

--------------------------------

### DataStar HTML with Web Component Integration

Source: https://data-star.dev/guide/datastar_expressions

An HTML structure demonstrating DataStar's declarative approach to integrating a custom web component. It shows how to bind input values to component attributes and capture custom events dispatched by the component to update a signal.

```html
<div data-signals:result="''">
    <input data-bind:foo />
    <my-component
        data-attr:src="$foo"
        data-on:mycustomevent="$result = evt.detail.value"
    ></my-component>
    <span data-text="$result"></span>
</div>
```

--------------------------------

### Go SDK for Generating SSE Script Execution Event

Source: https://data-star.dev/guide/datastar_expressions

A Go code snippet utilizing the DataStar SDK to programmatically generate a Server-Sent Event (SSE) that executes a JavaScript script on the client-side. This demonstrates server-side control over frontend script execution.

```go
sse := datastar.NewSSE(writer, request)
sse.ExecuteScript(`alert('This mission is too important for me to allow you to jeopardize it.')`)
```

--------------------------------

### Ruby: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Ruby snippet demonstrates creating server-sent events with a dynamic interval duration using the `Datastar` library. It calculates the duration based on the current seconds and formats the current time.

```ruby
1datastar = Datastar.new(request:, response:)
 2
 3now = Time.now
 4current_time = now.strftime('%Y-%m-%d %H:%M:%S')
 5current_seconds = now.strftime('%S').to_i
 6duration = current_seconds < 50 ? 5 : 1
 7
 8datastar.patch_elements <<~FRAGMENT
 9    <div id="time"
10         data-on-interval__duration.#{duration}s="@get('/endpoint')"
11    >
12        #{current_time}
13</div>
14FRAGMENT
```

--------------------------------

### Patch DOM Elements with Server-Sent Events (Rust)

Source: https://data-star.dev/guide

Demonstrates patching DOM elements using Server-Sent Events in Rust. It utilizes the `datastar` crate and `async_stream` to yield `PatchElements` commands, including a one-second sleep between them.

```rust
use async_stream::stream;
use datastar::prelude::*;
use std::thread;
use std::time::Duration;

Sse(stream! {
    // Patches elements into the DOM.
    yield PatchElements::new("<div id='hal'>I’m sorry, Dave. I’m afraid I can’t do that.</div>").into();

    thread::sleep(Duration::from_secs(1));
    
    yield PatchElements::new("<div id='hal'>Waiting for an order...</div>").into();
})
```

--------------------------------

### HTML: Two-Way Binding for Nested Signals

Source: https://data-star.dev/guide/backend_requests

Shows how to use two-way binding with nested signals in HTML. This allows for direct synchronization of input elements with nested signal states.

```html
<input data-bind:foo.bar />
```

--------------------------------

### PHP SDK for SSE Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This PHP code illustrates using the DataStar SDK's `ServerSentEventGenerator` to send SSE events for patching signals, with a one-second sleep between updates.

```php
 1use starfederation\datastar\ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4$sse = new ServerSentEventGenerator();
 5
 6// Patches signals.
 7$sse->patchSignals(['hal' => 'Affirmative, Dave. I read you.']);
 8
 9sleep(1);
10
11$sse->patchSignals(['hal' => '...']);
```

--------------------------------

### SSE Stream for Sequential Signal Updates

Source: https://data-star.dev/guide/reactive_signals

This SSE stream demonstrates sending multiple 'datastar-patch-signals' events sequentially, including a delay, to update and then reset a signal on the frontend.

```sse
1event: datastar-patch-signals
2data: signals {hal: 'Affirmative, Dave. I read you.'}
3
4// Wait 1 second
5
6event: datastar-patch-signals
7data: signals {hal: '...'}
8

```

--------------------------------

### Datastar TodoMVC HTML Structure

Source: https://data-star.dev/examples/todomvc

The HTML structure for the TodoMVC application implemented with Datastar. It includes elements for adding new todos, displaying the todo list, and managing todo status (All, Pending, Completed). This structure utilizes Datastar's custom attributes for event handling and data binding.

```html
<section
    id="todomvc"
    data-init="@get('/examples/todomvc/updates')"
>
    <header id="todo-header">
        <input
            type="checkbox"
            data-on:click__prevent="@post('/examples/todomvc/-1/toggle')"
            data-init="el.checked = false"
        />
        <input
            id="new-todo"
            type="text"
            placeholder="What needs to be done?"
            data-signals:input
            data-bind:input
            data-on:keydown="
                evt.key === 'Enter' && $input.trim() && @patch('/examples/todomvc/-1') && ($input = '');
            "
        />
    </header>
    <ul id="todo-list">
        <!-- Todo items are dynamically rendered here -->
    </ul>
    <div id="todo-actions">
        <span>
            <strong>0</strong> items pending
        </span>
        <button class="small info" data-on:click="@put('/examples/todomvc/mode/0')">
            All
        </button>
        <button class="small" data-on:click="@put('/examples/todomvc/mode/1')">
            Pending
        </button>
        <button class="small" data-on:click="@put('/examples/todomvc/mode/2')">
            Completed
        </button>
        <button class="error small" aria-disabled="true">
            Delete
        </button>
        <button class="warning small" data-on:click="@put('/examples/todomvc/reset')">
            Reset
        </button>
    </div>
</section>
```

--------------------------------

### Datastar Events: Patching Signals

Source: https://data-star.dev/how_tos/load_more_list_items

Demonstrates the `datastar-patch-signals` event used to update client-side signals. This event carries new signal values, such as the incremented offset, to maintain the application's state.

```text
1event: datastar-patch-signals
2data: signals {offset: 2}
3

```

--------------------------------

### Dispatch and Listen to Custom Events (HTML & JavaScript)

Source: https://data-star.dev/examples/custom_event

This snippet shows how to dispatch a custom event named 'myevent' every second from a paragraph element. It also demonstrates how to listen to this custom event using the `data-on` attribute and update the paragraph's text content with the event details. The `evt.detail` is accessible within the event handler.

```html
<p
    id="foo"
    data-signals:_event-details
    data-on:myevent="$_eventDetails = evt.detail"
    data-text="`Last Event Details: ${$_eventDetails}`"
></p>
<script>
    const foo = document.getElementById("foo");
    setInterval(() => {
        foo.dispatchEvent(
            new CustomEvent("myevent", {
                detail: JSON.stringify({
                    eventTime: new Date().toLocaleTimeString(),
                }),
            })
        );
    }, 1000);
</script>
```

--------------------------------

### Datastar Expression: Evaluating Signal Properties

Source: https://data-star.dev/guide/datastar_expressions

Shows how Datastar expressions can evaluate properties of signals, like `.length` for a string signal, by first converting the signal to its value and then evaluating the expression in a sandboxed context.

```html
1<div data-text="$foo.length"></div>
```

--------------------------------

### Send Multiple SSE Events in Python (Stream) using Datastar

Source: https://data-star.dev/guide/backend_requests

Demonstrates sending SSE events, including patching elements and signals, using Python's streaming API with Datastar. This is efficient for real-time data pushes.

```python
stream.patchElements('<div id="question">...</div>');
stream.patchElements('<div id="instructions">...</div>');
stream.patchSignals({'answer': '...', 'prize': '...'});
```

--------------------------------

### Java Servlet: Stream SSE Events with Datastar

Source: https://data-star.dev/guide/backend_requests

This Java code snippet shows how to implement Server-Sent Events (SSE) using the Datastar Java SDK within a Servlet environment. It demonstrates creating a `ServerSentEventGenerator` and using its `send` method with `PatchElements` and `PatchSignals` builders. This approach allows for fine-grained control over SSE generation in Java web applications.

```java
1import starfederation.datastar.utils.ServerSentEventGenerator;
2
3// Creates a new `ServerSentEventGenerator` instance.
4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
6
7// Patches elements into the DOM.
8generator.send(PatchElements.builder()
9    .data("<div id=\"question\">What do you put in a toaster?</div>")
10    .build()
11);
12
13// Patches signals.
14generator.send(PatchSignals.builder()
15    .data("{\"response\": \"\", \"answer\": \"\"}")
16    .build()
17);

```

--------------------------------

### Listen for Global Keydown Events (HTML)

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

This snippet demonstrates how to listen for any keydown event globally using the `data-on:keydown__window` attribute. The `__window` modifier ensures the event listener is attached to the `window` object. It triggers an alert when any key is pressed.

```html
<div data-on:keydown__window="alert('Key pressed')"></div>
```

--------------------------------

### HTML Structure for Progressive Load

Source: https://data-star.dev/examples/progressive_load

The HTML structure defines a button to trigger content loading and a main content area divided into sections. It uses custom data attributes for Datastar's directives to manage load states and progressive rendering. No external JS dependencies are explicitly mentioned for this HTML structure.

```html
<div>
    <div class="actions">
        <button
            id="load-button"
            data-signals:load-disabled="false"
            data-on:click="$loadDisabled=true; @get('/examples/progressive_load/updates')"
            data-attr:disabled="$loadDisabled"
            data-indicator:progressive-Load
        >
            Load
        </button>
        <!-- Indicator element -->
    </div>
    <p>
        Each part is loaded randomly and progressively.
    </p>
</div>
<div id="Load">
    <header id="header">Welcome to my blog</header>
    <section id="article">
        <h4>This is my article</h4>
        <section id="articleBody">
            <p>
                Lorem ipsum dolor sit amet...
            </p>
        </section>
    </section>
    <section id="comments">
        <h5>Comments</h5>
        <p>
            This is the comments section. It will also be progressively loaded as you scroll down.
        </p>
        <ul id="comments-list">
            <li id="1">
                <img src="https://avatar.iran.liara.run/username?username=example" alt="Avatar" class="avatar"/>
                This is a comment...
            </li>
            <!-- More comments loaded progressively -->
        </ul>
    </section>
    <div id="footer">Hope you like it</div>
</div>
```

--------------------------------

### Java SDK for SSE Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This Java code demonstrates using the DataStar SDK's `ServerSentEventGenerator` to send SSE events for patching signals, including a pause between updates.

```java
 1import starfederation.datastar.utils.ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
5ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);
 6
 7// Patches signals.
 8generator.send(PatchSignals.builder()
 9    .data("{\"hal\": \"Affirmative, Dave. I read you.\"}")
10    .build()
11);
12
13Thread.sleep(1000);
14
15generator.send(PatchSignals.builder()
16    .data("{\"hal\": \"...\"}")
17    .build()
18);
```

--------------------------------

### Patch DOM Elements and Signals using Data-Star

Source: https://data-star.dev/guide/backend_requests

This snippet demonstrates patching HTML elements and frontend signals into the DOM using the Data-Star library. It requires an existing DOM element with the ID 'question' to patch into. The function updates the DOM and frontend signals based on the provided data.

```javascript
stream.patchElements(`<div id="question">What do you put in a toaster?</div>`);
stream.patchSignals({'response':  '', 'answer': 'bread'});
```

--------------------------------

### Initialize Element on DOM Load/Modification (`data-init`)

Source: https://data-star.dev/reference

The `data-init` attribute executes an expression when the element is initialized in the DOM. This occurs on page load, DOM patching, or attribute changes. Modifiers like `__delay` and `__viewtransition` can alter the execution timing and behavior.

```html
<div data-init="$count = 1"></div>
<div data-init__delay.500ms="$count = 1"></div>
```

--------------------------------

### Send Multiple SSE Events in JavaScript using Datastar

Source: https://data-star.dev/guide/backend_requests

Demonstrates sending SSE events for patching elements and signals using JavaScript and Datastar's SSE helper. This is suitable for frontend or backend JavaScript environments.

```javascript
sse.PatchElements(`<div id="question">...</div>`)
sse.PatchElements(`<div id="instructions">...</div>`)
sse.PatchSignals([]byte(`{answer: '...', prize: '...'}`))
```

--------------------------------

### Send Multiple SSE Events in C# using Datastar

Source: https://data-star.dev/guide/backend_requests

Shows how to send SSE events, including patching elements and signals, using C# with the Datastar library. This is useful for backend integrations where C# is the primary language.

```csharp
datastarService.PatchElementsAsync(@"<div id=\"question\">...</div>");
datastarService.PatchElementsAsync(@"<div id=\"instructions\">...</div>");
datastarService.PatchSignalsAsync(new { answer = "...", prize = "..." } );
```

--------------------------------

### Integrating External JavaScript: Basic Function Definition

Source: https://data-star.dev/guide/datastar_expressions

Provides the definition for a simple, synchronous JavaScript function (`myfunction`) that takes data as input and returns a formatted string. This function is intended to be called from Datastar expressions.

```javascript
1function myfunction(data) {
2    return `You entered: ${data}`;
3}
```

--------------------------------

### Interactive Quiz using Datastar Attributes

Source: https://data-star.dev/guide/reactive_signals

This snippet implements an interactive quiz using `data-signals` for managing the user's response and the correct answer, `data-computed` to determine correctness, and `data-on` to handle user input via a prompt. `data-show` conditionally displays feedback.

```html
<div
    data-signals="{response: '', answer: 'bread'}"
    data-computed:correct="$response.toLowerCase() == $answer"
>
    <div id="question">What do you put in a toaster?</div>
    <button data-on:click="$response = prompt('Answer:') ?? ''">BUZZ</button>
    <div data-show="$response != ''">
        You answered “<span data-text="$response"></span>”.
        <span data-show="$correct">That is correct ✅</span>
        <span data-show="!$correct">
        The correct answer is “
        <span data-text="$answer"></span>
        ” 🤷
        </span>
    </div>
</div>
```

--------------------------------

### Datastar Actions - @setAll()

Source: https://data-star.dev/reference/actions

The `@setAll()` action sets the value of all matching signals to a specified value. It supports filtering signals by regular expressions to include or exclude specific patterns.

```APIDOC
## @setAll()

### Description
Sets the value of all matching signals (or all signals if no filter is used) to the expression provided in the first argument. The second argument is an optional filter object with an `include` property that accepts a regular expression to match signal paths. You can optionally provide an `exclude` property to exclude specific patterns.

### Method
Expression

### Endpoint
N/A (Used within Datastar expressions)

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```html
<!-- Sets the `foo` signal only -->
<div data-signals:foo="false">
    <button data-on:click="@setAll(true, {include: /^foo$/})"></button>
</div>

<!-- Sets all signals starting with `user.` -->
<div data-signals="{user: {name: '', nickname: ''}}">
    <button data-on:click="@setAll('johnny', {include: /^user\./})"></button>
</div>

<!-- Sets all signals except those ending with `_temp` -->
<div data-signals="{data: '', data_temp: '', info: '', info_temp: ''}">
    <button data-on:click="@setAll('reset', {include: /.*/, exclude: /_temp$/})"></button>
</div>
```

### Response
#### Success Response (200)
N/A (Modifies signal values directly)

#### Response Example
N/A
```

--------------------------------

### Linearly Interpolate Value with @fit()

Source: https://data-star.dev/reference/actions

Employs the `@fit()` Pro Action for linear interpolation between two ranges. It can optionally clamp the output to the new range and round to the nearest integer, useful for scaling and unit conversions.

```html
<!-- Convert a 0-100 slider to 0-255 RGB value -->
<div>
    <input type="range" min="0" max="100" value="50" data-bind:slider-value>
    <div data-computed:rgb-value="@fit($sliderValue, 0, 100, 0, 255)">
        RGB Value: <span data-text="$rgbValue"></span>
    </div>
</div>

<!-- Convert Celsius to Fahrenheit -->
<div>
    <input type="number" data-bind:celsius value="20" />
    <div data-computed:fahrenheit="@fit($celsius, 0, 100, 32, 212)">
        <span data-text="$celsius"></span>°C = <span data-text="$fahrenheit.toFixed(1)"></span>°F
    </div>
</div>

<!-- Map mouse position to element opacity (clamped) -->
<div
    data-signals:mouse-x="0"
    data-computed:opacity="@fit($mouseX, 0, window.innerWidth, 0, 1, true)"
    data-on:mousemove__window="$mouseX = evt.clientX"
    data-attr:style="'opacity: ' + $opacity">
    Move your mouse horizontally to change opacity
</div>
```

--------------------------------

### DataStar Interval: Rendering Current Time with Backend Data

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This snippet illustrates how to combine the interval attribute with backend templating to display dynamic content, such as the current time, which is updated periodically.

```html
<div id="time"
     data-on-interval__duration.5s="@get('/endpoint')"
>
    {{ now }}
</div>
```

--------------------------------

### Kotlin SDK for SSE Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This Kotlin code snippet shows how to use the `ServerSentEventGenerator` to patch signals via SSE, including a one-second delay between the signal updates.

```kotlin
 1val generator = ServerSentEventGenerator(response)
 2
 3generator.patchSignals(
 4    signals = """{"hal": "Affirmative, Dave. I read you."} """,
 5)
 6
 7Thread.sleep(ONE_SECOND)
 8
 9generator.patchSignals(
10    signals = """{"hal": "..."} """,
11)
```

--------------------------------

### Handle SSE Stream in JavaScript with Datastar

Source: https://data-star.dev/guide/reactive_signals

This JavaScript code snippet shows how to create and manage a Server-Sent Event (SSE) stream using Datastar's client-side SDK. It initializes an `ServerSentEventGenerator`, sends patch signals, and uses `setTimeout` to introduce delays between events. The `stream.patchSignals` method is used to send data.

```javascript
// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
ServerSentEventGenerator.stream(req, res, (stream) => {
    // Patches signals.
    stream.patchSignals({'hal': 'Affirmative, Dave. I read you.'});

    setTimeout(() => {
        stream.patchSignals({'hal': '...'});
    }, 1000);
});
```

--------------------------------

### Datastar Expressions: Ternary, Logical OR, and AND Operators

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates the use of JavaScript operators within Datastar expressions, including the ternary operator for conditional output, logical OR for combining conditions, and logical AND for conditional actions like sending a request.

```html
1// Output one of two values, depending on the truthiness of a signal
2<div data-text="$landingGearRetracted ? 'Ready' : 'Waiting'"></div>
3
4// Show a countdown if the signal is truthy or the time remaining is less than 10 seconds
5<div data-show="$landingGearRetracted || $timeRemaining < 10">
6    Countdown
7</div>
8
9// Only send a request if the signal is truthy
10<button data-on:click="$landingGearRetracted && @post('/launch')">
11    Launch
12</button>
```

--------------------------------

### Clojure: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Clojure snippet configures server-sent events to dynamically adjust the interval duration based on the current second of the minute. It uses `starfederation.datastar.clojure.api` and `http-kit` for SSE handling.

```clojure
1(require
 2  '[starfederation.datastar.clojure.api :as d*]
 3  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
 4  '[some.hiccup.library :refer [html]])
 5
 6(import
 7  'java.time.format.DateTimeFormatter
 8  'java.time.LocalDateTime)
 9
10(def date-time-formatter (DateTimeFormatter/ofPattern "YYYY-MM-DD HH:mm:ss"))
11(def seconds-formatter (DateTimeFormatter/ofPattern "ss"))
12
13(defn handle [ring-request]
14  (->sse-response ring-request
15    {on-open
16     (fn [sse]
17       (let [now (LocalDateTime/now)
18             current-time (LocalDateTime/.format now date-time-formatter)
19             seconds (LocalDateTime/.format now seconds-formatter)
20             duration (if (neg? (compare seconds "50"))
21                         "5"
22                         "1")]
23         (d*/patch-elements! sse
24           (html [:div#time {(str "data-on-interval__duration." duration "s")
25                             (d*/sse-get "/endpoint")}
26                   current-time]))))}))
27
28         (d*/close-sse! sse))}))
```

--------------------------------

### PHP: `ServerSentEventGenerator` Usage

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This PHP code shows how to use `ServerSentEventGenerator` to send Server-Sent Events for element patching. It captures the current date and time and embeds it within an HTML structure for periodic updates.

```php
use starfederation\datastar\ServerSentEventGenerator;

$currentTime = date('Y-m-d H:i:s');

$sse = new ServerSentEventGenerator();
$sse->patchElements(`
    <div id="time"
         data-on-interval__duration.5s="@get('/endpoint')"
    >
        $currentTime
    </div>
`);

```

--------------------------------

### Ruby SDK for SSE Patch Signals (Rack)

Source: https://data-star.dev/guide/reactive_signals

This Ruby code shows how to use the DataStar gem within a Rack application to stream SSE events for patching signals, including a sleep duration between updates.

```ruby
 1require 'datastar'
 2
 3# Create a Datastar::Dispatcher instance
 4
 5datastar = Datastar.new(request:, response:)
 6
 7# In a Rack handler, you can instantiate from the Rack env
 8# datastar = Datastar.from_rack_env(env)
 9
10# Start a streaming response
11datastar.stream do |sse|
12  # Patches signals
13  sse.patch_signals(hal: 'Affirmative, Dave. I read you.')
14
15  sleep 1
16  
17  sse.patch_signals(hal: '...')
18end
```

--------------------------------

### Send Data via POST Request using Datastar

Source: https://data-star.dev/guide/backend_requests

This snippet demonstrates how to trigger a POST request to a specified backend endpoint when a button is clicked, using Datastar's `@post()` directive. It's a simple way to send data for server-side processing.

```html
<button data-on:click="@post('/actions/quiz')">
    Submit answer
</button>
```

--------------------------------

### Clojure SDK for SSE Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This Clojure code snippet utilizes the DataStar SDK to generate Server-Sent Events (SSE) for patching signals, including a timed delay between updates.

```clojure
 1;; Import the SDK's api and your adapter
 2(require
 3  '[starfederation.datastar.clojure.api :as d*]
 4  '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
 5
 6;; in a ring handler
 7(defn handler [request]
 8  ;; Create an SSE response
 9  (->sse-response request
10                  {on-open
11                   (fn [sse]
12                     ;; Patches signal.
13                     (d*/patch-signals! sse "{hal: 'Affirmative, Dave. I read you.'}")
14                     (Thread/sleep 1000)
15                     (d*/patch-signals! sse "{hal: '...'}"))}))
```

--------------------------------

### Request Cancellation

Source: https://data-star.dev/reference/actions

Explains Datastar's default request cancellation behavior and how to customize it using the `requestCancellation` option.

```APIDOC
## Request Cancellation

By default, Datastar automatically cancels any ongoing request on the same element when a new one is initiated. This prevents conflicts from rapid user interactions.

### Default Behavior Example:
```html
<!-- Clicking multiple times cancels previous requests -->
<button data-on:click="@get('/slow-endpoint')">Load Data</button>
```

This behavior is element-specific; requests on different elements run concurrently.

### Customizing Request Cancellation:

Use the `requestCancellation` option to control this behavior:

*   **`'disabled'`**: Allows concurrent requests on the same element.
*   **`AbortController` instance**: Provides custom control over cancellation.

### Examples:

*   **Allowing concurrent requests:**
    ```html
    <!-- Allow multiple concurrent requests -->
    <button data-on:click="@get('/endpoint', {requestCancellation: 'disabled'})">Allow Multiple</button>
    ```

*   **Using a custom `AbortController`:**
    ```html
    <div data-signals:controller="new AbortController()">
        <button data-on:click="@get('/endpoint', {requestCancellation: $controller})">Start Request</button>
        <button data-on:click="$controller.abort()">Cancel Request</button>
    </div>
    ```
```

--------------------------------

### Capture Counter Changes with Signal Patch

Source: https://data-star.dev/examples/on_signal_patch

This HTML fragment utilizes the `data-on-signal-patch` attribute to push detected changes to the 'counter' signal into the `counterChanges` array. It includes a filter `data-on-signal-patch-filter` to specifically target changes related to the 'counter' signal.

```html
<div
    data-on-signal-patch="$counterChanges.push(patch)"
    data-on-signal-patch-filter="{include: /^counter$/}"
>
    <h3>Counter Changes Only</h3>
    <pre data-json-signals__terse="{include: /^counterChanges/}"></pre>
</div>
```

--------------------------------

### Set JavaScript Response Headers

Source: https://data-star.dev/reference/actions

Sets the Content-Type to text/javascript and optionally includes 'datastar-script-attributes' for customizing script element attributes. This allows for server-side configuration of client-side scripts.

```javascript
response.headers.set('Content-Type', 'text/javascript')
response.headers.set('datastar-script-attributes', JSON.stringify({ type: 'module' }))
response.body = 'console.log("Hello from server!");'
```

--------------------------------

### Datastar Actions - @peek()

Source: https://data-star.dev/reference/actions

The `@peek()` action allows accessing signals without subscribing to their changes in expressions. This is useful for reading signal values in contexts where you don't want to trigger re-evaluations on signal updates.

```APIDOC
## @peek()

### Description
Allows accessing signals without subscribing to their changes in expressions.

### Method
Expression

### Endpoint
N/A (Used within Datastar expressions)

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```html
<div data-text="$foo + @peek(() => $bar)"></div>
```

### Response
#### Success Response (200)
N/A (Returns the evaluated value of the signal within the expression)

#### Response Example
N/A
```

--------------------------------

### Integrating External JavaScript: Basic Function Call

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates how to call an external JavaScript function from a Datastar expression, passing data via arguments and capturing the return value to update a signal. This promotes code encapsulation.

```html
1<div data-signals:result>
2    <input data-bind:foo 
3        data-on:input="$result = myfunction($foo)"
4    >
5    <span data-text="$result"></span>
6</div>
```

--------------------------------

### Patch HTML Elements via SSE in Kotlin

Source: https://data-star.dev/guide

This Kotlin code snippet shows how to patch HTML elements into the DOM using Datastar's `ServerSentEventGenerator`. It demonstrates sending updates to a specific div with a delay, utilizing Kotlin's concise syntax for multiline strings and function calls. This is applicable in Kotlin-based web applications.

```kotlin
1val generator = ServerSentEventGenerator(response)
2
3generator.patchElements(
4    elements = """<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>""",
5)
6
7Thread.sleep(ONE_SECOND)
8
9generator.patchElements(
10    elements = """<div id=\"hal\">Waiting for an order...</div>""",
11)
```

--------------------------------

### Append Element and Update Signals in PHP

Source: https://data-star.dev/how_tos/load_more_list_items

This PHP snippet illustrates appending new items to an HTML list and updating signals using the ServerSentEventGenerator. It dynamically calculates the new offset and conditionally removes a 'load more' button based on the current offset and maximum items.

```php
use starfederation\datastar\enums\ElementPatchMode;
use starfederation\datastar\ServerSentEventGenerator;

$signals = ServerSentEventGenerator::readSignals();

$max = 5;
$limit = 1;
$offset = $signals['offset'] ?? 1;

$sse = new ServerSentEventGenerator();

if ($offset < $max) {
    $newOffset = $offset + $limit;
    $sse->patchElements("<div>Item $newOffset</div>", [
        'selector' => '#list',
        'mode' => ElementPatchMode::Append,
    ]);
    if (newOffset < $max) {
        $sse->patchSignals(['offset' => $newOffset]);
    } else {
        $sse->removeElements('#load-more');
    }
}
```

--------------------------------

### Handle click events with data-on

Source: https://data-star.dev/guide

Use the `data-on` attribute to attach an event listener to an HTML element. When the specified event occurs, the associated Datastar expression is executed.

```html
<button data-on:click="alert('I’m sorry, Dave. I’m afraid I can’t do that.')">
    Open the pod bay doors, HAL.
</button>
```

--------------------------------

### Rust Actix: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Rust snippet using Actix demonstrates sending server-sent events with a dynamically adjusted interval duration. It leverages the `chrono` crate for time and `async_stream` for stream processing.

```rust
1use datastar::prelude::*
 2use chrono::Local;
 3use async_stream::stream;
 4
 5let current_time = Local::now().format("%Y-%m-%d %H:%M:%S").to_string();
 6let current_seconds = Local::now().second();
 7let duration = if current_seconds < 50 {
 8    5
 9} else {
10    1
11};
12
13Sse(stream! {
14    yield PatchElements::new(
15        format!(
16            "<div id='time' data-on-interval__duration.{}s='@get(\"/endpoint\")'>{}</div>",
17            duration,
18            current_time,
19        )
20    ).into();
21})
```

--------------------------------

### Send Multiple SSE Events (Patch Elements and Signals)

Source: https://data-star.dev/guide/backend_requests

Illustrates sending multiple Server-Sent Events (SSE) in a single response. This includes patching HTML elements and sending signal data, showcasing SSE's capability to update the UI and application state efficiently.

```datastar-dsl
(d*/patch-elements! sse "<div id=\"question\">...</div>")
(d*/patch-elements! sse "<div id=\"instructions\">...</div>")
(d*/patch-signals! sse "{answer: '...', prize: '...'}")
```

--------------------------------

### data-on-interval: Execute expressions at regular intervals

Source: https://data-star.dev/reference/attributes

The `data-on-interval` attribute allows expressions to be executed at a set interval, defaulting to one second. The `__duration` modifier can change this interval, and `.leading` can execute the first interval immediately. View transitions can also be integrated.

```html
<div data-on-interval="$count++"></div>
<div data-on-interval__duration.500ms="$count++"></div>
```

--------------------------------

### Datastar: Two-Way Data Binding with data-bind

Source: https://data-star.dev/guide/reactive_signals

Sets up two-way data binding for input elements. It creates a signal that is bound to the element's value, updating automatically when either changes. Can use attribute shorthand or value assignment.

```html
<input data-bind:foo />
```

```html
<input data-bind="foo" />
```

--------------------------------

### Trigger Alert on 'Enter' or 'Ctrl + L' (HTML)

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

This snippet combines conditions to trigger an alert for either the 'Enter' key press or the 'Ctrl + L' key combination. It uses logical OR (`||`) to check multiple conditions within the `data-on:keydown__window` attribute, leveraging `evt.key` and `evt.ctrlKey`.

```html
<div data-on:keydown__window="(evt.key === 'Enter' || (evt.ctrlKey && evt.key === 'l')) && alert('Key pressed')"></div>
```

--------------------------------

### Datastar: Binding HTML Attributes with data-attr

Source: https://data-star.dev/guide/reactive_signals

Binds the value of any HTML attribute to an expression. Can set multiple attributes using key-value pairs.

```html
<input data-bind:foo />
<button data-attr:disabled="$foo == ''">
    Save
</button>
```

```html
<button data-attr="{disabled: $foo == '', title: $foo}">Save</button>
```

--------------------------------

### SVG Structure for Namespacing

Source: https://data-star.dev/examples/svg_morphing

Demonstrates the required SVG structure for morphing, ensuring inner SVG elements are correctly namespaced by wrapping them in an outer <svg> tag. This is crucial for Datastar's SVG handling.

```html
<svg>
    <svg id="target">
        <circle cx="50" cy="100" r="50" fill="red" />
    </svg>
    <circle cx="150" cy="100" r="50" fill="red" />
</svg>
```

--------------------------------

### Datastar Events: Appending and Removing Elements

Source: https://data-star.dev/how_tos/load_more_list_items

Defines the Datastar server-sent events used to update the UI. `datastar-patch-elements` with `append` mode adds new items to the list, and `remove` mode hides the load more button when all items are shown.

```text
1event: datastar-patch-elements
2data: selector #list
3data: mode append
4data: elements <div>Item 2</div>
5

```

```text
1event: datastar-patch-elements
2data: selector #load-more
3data: mode remove
4

```

--------------------------------

### Set JSON Response Headers

Source: https://data-star.dev/reference/actions

Sets the Content-Type to application/json and optionally includes the 'datastar-only-if-missing' header for patching signals. This is useful for controlling how updates are applied to existing data.

```javascript
response.headers.set('Content-Type', 'application/json')
response.headers.set('datastar-only-if-missing', 'true')
response.body = JSON.stringify({ foo: 'bar' })
```

--------------------------------

### HTML Element for Real-time Data Display

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This HTML snippet defines a div element with specific data attributes used for controlling interval-based data fetching and displaying dynamic content. The `data-on-interval__duration` attribute likely configures the polling frequency, and the `${currentTime.toISOString()}` placeholder indicates where the updated data will be rendered.

```html
<div id="time"
     data-on-interval__duration.${duration}s="@get('/endpoint]')"
  >
    ${currentTime.toISOString()}
</div>
```

--------------------------------

### Set Multiple Signals with @setAll() in DataStar

Source: https://data-star.dev/reference/actions

The @setAll() action sets the value of multiple signals, optionally filtered by a regular expression. It accepts a value and an optional filter object with 'include' and 'exclude' RegExp properties to target specific signals. This is useful for batch updates.

```html
<!-- Sets the `foo` signal only -->
<div data-signals:foo="false">
    <button data-on:click="@setAll(true, {include: /^foo$/})"></button>
</div>

<!-- Sets all signals starting with `user.` -->
<div data-signals="{user: {name: '', nickname: ''}}">
    <button data-on:click="@setAll('johnny', {include: /^user\./})"></button>
</div>

<!-- Sets all signals except those ending with `_temp` -->
<div data-signals="{data: '', data_temp: '', info: '', info_temp: ''}">
    <button data-on:click="@setAll('reset', {include: /.*/, exclude: /_temp$/})"></button>
</div>
```

--------------------------------

### Datastar Actions - @toggleAll()

Source: https://data-star.dev/reference/actions

The `@toggleAll()` action toggles the boolean value of all matching signals. It also supports filtering signals using regular expressions for inclusion and exclusion.

```APIDOC
## @toggleAll()

### Description
Toggles the boolean value of all matching signals (or all signals if no filter is used). The argument is an optional filter object with an `include` property that accepts a regular expression to match signal paths. You can optionally provide an `exclude` property to exclude specific patterns.

### Method
Expression

### Endpoint
N/A (Used within Datastar expressions)

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```html
<!-- Toggles the `foo` signal only -->
<div data-signals:foo="false">
    <button data-on:click="@toggleAll({include: /^foo$/})"></button>
</div>

<!-- Toggles all signals starting with `is` -->
<div data-signals="{isOpen: false, isActive: true, isEnabled: false}">
    <button data-on:click="@toggleAll({include: /^is/})"></button>
</div>

<!-- Toggles signals starting with `settings.` -->
<div data-signals="{settings: {darkMode: false, autoSave: true}}">
    <button data-on:click="@toggleAll({include: /^settings\./})"></button>
</div>
```

### Response
#### Success Response (200)
N/A (Modifies signal values directly)

#### Response Example
N/A
```

--------------------------------

### Datastar Expressions: Multi-line Statements

Source: https://data-star.dev/guide/datastar_expressions

Shows how Datastar expressions can span multiple lines, requiring semicolons to separate statements. This is useful for organizing complex logic within `data-*` attributes.

```html
1<div data-signals:foo="1">
2    <button data-on:click="
3        $landingGearRetracted = true; 
4        @post('/launch')
5    ">
6        Force launch
7    </button>
8</div>
```

--------------------------------

### Patch HTML Elements via SSE in Clojure

Source: https://data-star.dev/guide

This Clojure code snippet shows how to generate Server-Sent Events (SSE) to patch HTML elements into the DOM. It uses `datastar-patch-elements` events to update a div with ID 'hal' and demonstrates sending multiple updates with a delay. Requires Datastar SDK and an http-kit adapter.

```clojure
1;; Import the SDK's api and your adapter
2(require
3 '[starfederation.datastar.clojure.api :as d*]
4 '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])
5
6;; in a ring handler
7(defn handler [request]
8  ;; Create an SSE response
9  (->sse-response request
10                  {:on-open
11                   (fn [sse] 
12                     ;; Patches elements into the DOM
13                     (d*/patch-elements! sse
14                                         "<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>")
15                     (Thread/sleep 1000)
16                     (d*/patch-elements! sse
17                                         "<div id=\"hal\">Waiting for an order...</div>"))}))
```

--------------------------------

### HTML Form for Inline Email Validation

Source: https://data-star.dev/examples/inline_validation

This HTML code defines a form with input fields for Email Address, First Name, and Last Name. The email input is configured for inline validation, triggering a server-side POST request on keydown events with a debounce of 500ms. Validation results are displayed to the user.

```html
<div id="demo">
    <label>
        Email Address
        <input
            type="email"
            required
            aria-live="polite"
            aria-describedby="email-info"
            data-bind:email
            data-on:keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <p id="email-info" class="info">The only valid email address is "test@test.com".</p>
    <label>
        First Name
        <input
            type="text"
            required
            aria-live="polite"
            data-bind:first-name
            data-on:keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <label>
        Last Name
        <input
            type="text"
            required
            aria-live="polite"
            data-bind:last-name
            data-on:keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <button
        class="success"
        data-on:click="@post('/examples/inline_validation')"
    >
        <i class="material-symbols:person-add"></i>
        Sign Up
    </button>
</div>
```

--------------------------------

### Python Sanic: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Python snippet for the Sanic web framework uses `datastar_py` to send server-sent events with a dynamic interval duration. It determines the duration based on the current seconds of the minute.

```python
1from datastar_py import ServerSentEventGenerator as SSE
 2from datastar_py.sanic import DatastarResponse
 3
 4@app.get("/endpoint")
 5async def endpoint():
 6    current_time = datetime.now()
 7    duration = 5 if current_time.seconds < 50 else 1
 8
 9    return DatastarResponse(SSE.patch_elements(f"""
10        <div id="time" data-on-interval__duration.{duration}s="@get('/endpoint')">
11            {current_time:%Y-%m-%d %H:%M:%S}
12        </div>
13    """))
```

--------------------------------

### PHP: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This PHP snippet generates server-sent events with a dynamic interval duration using the `ServerSentEventGenerator` class. It calculates the duration based on the current seconds of the minute.

```php
1use starfederation\datastar\ServerSentEventGenerator;
 2
 3$currentTime = date('Y-m-d H:i:s');
 4$currentSeconds = date('s');
 5$duration = $currentSeconds < 50 ? 5 : 1;
 6
 7$sse = new ServerSentEventGenerator();
 8$sse->patchElements(`
 9    <div id="time"
10         data-on-interval__duration.${duration}s="@get('/endpoint')"
11    >
12        $currentTime
13    </div>
14`);
```

--------------------------------

### Run expression at regular intervals (HTML)

Source: https://data-star.dev/reference

The `data-on-interval` attribute executes a given expression at a set interval, defaulting to one second. This is ideal for periodic updates or counters. The interval duration can be modified using the `__duration` modifier.

```html
<div data-on-interval="$count++"></div>
```

--------------------------------

### Infinite Scroll Trigger with data-on-intersect

Source: https://data-star.dev/examples/infinite_scroll

This HTML snippet shows how to implement infinite scroll. The `data-on-intersect` attribute on the last element triggers a request to a specified URL when that element scrolls into view. The response is then appended to the DOM. This method is useful for progressively loading large datasets without overwhelming the user or the initial page load.

```html
<div data-on-intersect="@get('/examples/infinite_scroll/more')">
    Loading...
</div>
```

--------------------------------

### Patch DOM Elements with Server-Sent Events (Python)

Source: https://data-star.dev/guide

Employs the ServerSentEventGenerator (aliased as SSE) in Python within a Sanic web framework context. It yields SSE patch_elements calls to update a DOM element, with an asynchronous delay between updates.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response

@app.get('/open-the-bay-doors')
@datastar_response
async def open_doors(request):
    yield SSE.patch_elements('<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>')
    await asyncio.sleep(1)
    yield SSE.patch_elements('<div id="hal">Waiting for an order...</div>')
```

--------------------------------

### datastar-patch-signals SSE Event

Source: https://data-star.dev/reference/sse_events

Handles patching of signals. It allows updating existing signals or adding new ones, with an option to only update if a signal is missing.

```APIDOC
## datastar-patch-signals SSE Event

### Description

This SSE event type is used to patch signals into the existing signals on the page. It allows updating signal values and provides an option to only update a signal if it does not already exist.

### Method

Server-Sent Events (SSE)

### Endpoint

N/A (SSE stream endpoint)

### Parameters

#### Event Data (`data` lines)

- **`onlyIfMissing [true|false]`** - Optional - Determines whether to update each signal with the new value only if a signal with that name does not yet exist. Defaults to `false`.
- **`signals {<signal_name>: <value>}`** - Required - A valid `data-signals` attribute string representing the signals to patch. Signal values can be set to `null` to remove them.

### Request Example

```
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}
```

```
event: datastar-patch-signals
data: signals {foo: null, bar: null}
```

```
event: datastar-patch-signals
data: onlyIfMissing true
data: signals {foo: 1, bar: 2}
```

### Response

#### Success Response (200 OK)

SSE stream containing `datastar-patch-signals` events.

#### Response Example

```
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}
```
```

--------------------------------

### Two-Way Data Binding with data-bind

Source: https://data-star.dev/reference/attributes

The `data-bind` attribute creates a two-way data binding between an element's value and a signal. It supports various input types, predefined signal types (like numbers and arrays), and file uploads, automatically encoding files in base64. Modifiers like `__case` can alter signal naming conventions.

```html
<input data-bind:foo />
```

```html
<input data-bind="foo" />
```

```html
<input data-bind:foo value="bar" />
```

```html
<div data-signals:foo="baz">
    <input data-bind:foo value="bar" />
</div>
```

```html
<div data-signals:foo="0">
    <select data-bind:foo>
        <option value="10">10</option>
    </select>
</div>
```

```html
<div data-signals:foo="[]">
    <input data-bind:foo type="checkbox" value="bar" />
    <input data-bind:foo type="checkbox" value="baz" />
</div>
```

```html
<input type="file" data-bind:files multiple />
```

```html
<input data-bind:my-signal__case.kebab />
```

--------------------------------

### Direct JavaScript Execution Response

Source: https://data-star.dev/guide/datastar_expressions

A simple JavaScript alert message. This code snippet represents a response from a backend action with a 'text/javascript' content-type, intended to be executed directly in the browser.

```javascript
alert('This mission is too important for me to allow you to jeopardize it.')
```

--------------------------------

### Manage Conditional and Default Styles with data-style

Source: https://data-star.dev/reference/attributes

data-style handles conditional styling effectively. Setting a style value to falsy (empty string, null, undefined, false) restores the original inline style or removes the property. This preserves existing styles and manages dynamic changes.

```html
<!-- When $x is false, color remains red from inline style -->
<div style="color: red;" data-style:color="$x && 'green'"></div>

<!-- When $hiding is true, display becomes none; when false, reverts to flex from inline style -->
<div style="display: flex;" data-style:display="$hiding && 'none'"></div>
```

--------------------------------

### Patch DOM Elements with Server-Sent Events (PHP)

Source: https://data-star.dev/guide

Uses the ServerSentEventGenerator class in PHP to patch HTML content into a DOM element. It sends an initial message, waits for one second, and then sends a second message to the same element.

```php
require 'datastar'

# Create a Datastar::Dispatcher instance

datastar = Datastar.new(request:, response:)

# In a Rack handler, you can instantiate from the Rack env
# datastar = Datastar.from_rack_env(env)

# Start a streaming response
datastar.stream do |sse|
  # Patches elements into the DOM.
  sse.patch_elements %(<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>)

  sleep 1
  
  sse.patch_elements %(<div id="hal">Waiting for an order...</div>)
end
```

--------------------------------

### Datastar: Setting Text Content with data-text

Source: https://data-star.dev/guide/reactive_signals

Sets the text content of an element to the value of a signal or a Datastar expression. Requires the '$' prefix for signals and allows JavaScript expressions for dynamic content.

```html
<input data-bind:foo />
<div data-text="$foo"></div>
```

```html
<input data-bind:foo />
<div data-text="$foo.toUpperCase()"></div>
```

--------------------------------

### Prevent Default on 'Enter' Keydown and Alert (HTML)

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

This code illustrates how to prevent the default action of a keydown event, such as form submission, while also showing an alert. It uses `evt.preventDefault()` within the `data-on:keydown__window` expression, executed conditionally for the 'Enter' key. The `evt` object provides the `preventDefault` method.

```html
<div data-on:keydown__window="evt.key === 'Enter' && (evt.preventDefault(), alert('Key pressed'))"></div>
```

--------------------------------

### Datastar Patch Elements: Default Morphing

Source: https://data-star.dev/reference/sse_events

Demonstrates the default 'outer' morphing mode for patching elements in the DOM. Elements are matched by ID and their outer HTML is replaced. Ensure top-level elements have IDs for correct matching.

```html
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>


```

--------------------------------

### Append Element and Update Signals in Kotlin

Source: https://data-star.dev/how_tos/load_more_list_items

This snippet demonstrates how to append new items to a list and update pagination signals using Data Star's generator in Kotlin. It includes logic to remove a 'load more' button when the maximum number of items is reached.

```kotlin
fun updateList(generator: ServerSentEventGenerator, offset: Int, limit: Int, max: Int) {
    val newOffset = offset + limit

    generator.patchElements(
        elements = "<div>Item $newOffset</div>",
        options =
            PatchElementsOptions(
                selector = "#list",
                mode = ElementPatchMode.Append,
            ),
    )

    if (newOffset < max) {
        generator.patchSignals(
            signals = "{\"offset\": $newOffset}"",
        )
    } else {
        generator.patchElements(
            options =
                PatchElementsOptions(
                    selector = "#load-more",
                    mode = ElementPatchMode.Remove,
                ),
        )
    }
}
```

--------------------------------

### Integrating External JavaScript: Asynchronous Function Definition

Source: https://data-star.dev/guide/datastar_expressions

Defines an asynchronous JavaScript function (`myfunction`) that simulates an async operation using `setTimeout` and `Promise`. Upon completion, it dispatches a custom event containing the result, allowing Datastar expressions to capture it.

```javascript
1async function myfunction(element, data) {
2    const value = await new Promise((resolve) => {
3        setTimeout(() => resolve(`You entered: ${data}`), 1000);
4    });
5    element.dispatchEvent(
6        new CustomEvent('mycustomevent', {detail: {value}})
7    );
8}
```

--------------------------------

### Datastar Request Cancellation Behaviors

Source: https://data-star.dev/reference/actions

This snippet illustrates Datastar's request cancellation. The default behavior automatically cancels previous requests on the same element when a new one is initiated, preventing conflicts. The 'disabled' option allows concurrent requests, and an AbortController provides custom control.

```html
<!-- Clicking this button multiple times will cancel previous requests (default behavior) -->
<button data-on:click="@get('/slow-endpoint')">Load Data</button>
```

```html
<!-- Allow concurrent requests (no automatic cancellation) -->
<button data-on:click="@get('/endpoint', {requestCancellation: 'disabled'})">Allow Multiple</button>

<!-- Custom abort controller for fine-grained control -->
<div data-signals:controller="new AbortController()">
    <button data-on:click="@get('/endpoint', {requestCancellation: $controller})">Start Request</button>
    <button data-on:click="$controller.abort()">Cancel Request</button>
</div>
```

--------------------------------

### Toggle Boolean Signals with @toggleAll() in DataStar

Source: https://data-star.dev/reference/actions

The @toggleAll() action toggles the boolean value of multiple signals, with optional filtering using regular expressions. It accepts a filter object with 'include' and 'exclude' RegExp properties to specify which signals to toggle. This simplifies managing boolean states across signals.

```html
<!-- Toggles the `foo` signal only -->
<div data-signals:foo="false">
    <button data-on:click="@toggleAll({include: /^foo$/})"></button>
</div>

<!-- Toggles all signals starting with `is` -->
<div data-signals="{isOpen: false, isActive: true, isEnabled: false}">
    <button data-on:click="@toggleAll({include: /^is/})"></button>
</div>

<!-- Toggles signals starting with `settings.` -->
<div data-signals="{settings: {darkMode: false, autoSave: true}}">
    <button data-on:click="@toggleAll({include: /^settings\./})"></button>
</div>
```

--------------------------------

### Create and Patch Signals using data-signals Attribute

Source: https://data-star.dev/guide/reactive_signals

The `data-signals` attribute is used to create, update, or remove signals. It can be applied directly with a value, use dot-notation for nested signals, or accept a JSON object for patching multiple signals simultaneously.

```html
<div data-signals:foo="1"></div>
<div data-signals:form.foo="2"></div>
<div data-signals="{foo: 1, form: {foo: 2}}"></div>
```

--------------------------------

### Trigger Alert on 'Enter' Keydown (HTML)

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

This code snippet shows how to trigger an alert specifically when the 'Enter' key is pressed. It utilizes the `evt.key` property within the `data-on:keydown__window` expression to check the pressed key. The `evt` variable provides access to the event object's properties.

```html
<div data-on:keydown__window="evt.key === 'Enter' && alert('Key pressed')"></div>
```

--------------------------------

### DRY Datastar Action with Templating Loop

Source: https://data-star.dev/how_tos/keep_datastar_code_dry

This snippet shows another way to achieve DRY code in Datastar using a templating language's loop. It iterates over a list of labels to generate buttons, each performing the same backend action.

```html
{% set labels = ['Click me', 'No, click me!', 'Click us all!'] %}
{% for label in labels %}
    <button data-on:click="@get('/endpoint')">{{ label }}</button>
{% endfor %}
```

--------------------------------

### Apply Dynamic CSS Styles with data-style

Source: https://data-star.dev/reference/attributes

The data-style directive sets inline CSS styles based on expressions, keeping them synchronized with your application's state. It supports individual style properties and multiple properties defined in an object.

```html
<div data-style:background-color="$usingRed ? 'red' : 'blue'"></div>
<div data-style:display="$hiding && 'none'"></div>
<div data-style="{
    display: $hiding ? 'none' : 'flex',
    flexDirection: 'column',
    color: $usingRed ? 'red' : 'green'
}"></div>
```

--------------------------------

### Custom Web Component with DataStar Event Dispatch

Source: https://data-star.dev/guide/datastar_expressions

A JavaScript class defining a custom HTML element 'my-component' that extends HTMLElement. It observes the 'src' attribute and dispatches a 'mycustomevent' with the new value when the attribute changes, enabling communication back to DataStar bindings.

```javascript
class MyComponent extends HTMLElement {
    static get observedAttributes() {
        return ['src'];
    }

    attributeChangedCallback(name, oldValue, newValue) {
        const value = `You entered: ${newValue}`;
        this.dispatchEvent(
            new CustomEvent('mycustomevent', {detail: {value}})
        );
    }
}

customElements.define('my-component', MyComponent);
```

--------------------------------

### Apply Casing to Class Names with data-class Modifier

Source: https://data-star.dev/reference/attributes

The `__case` modifier for `data-class` allows you to convert the casing of the class name. Supported cases include camel, kebab (default), snake, and pascal.

```html
<div data-class:my-class__case.camel="$foo"></div>
```

--------------------------------

### Create Element Reference with data-ref

Source: https://data-star.dev/reference/attributes

The data-ref directive creates a signal that references an HTML element. The signal name can be specified either as the attribute key or value. This is useful for obtaining a direct reference to an element for manipulation.

```html
<div data-ref:foo></div>
<div data-ref="foo"></div>
$foo is a reference to a <span data-text="$foo.tagName"></span> element
```

--------------------------------

### data-on-signal-patch: React to signal changes

Source: https://data-star.dev/reference/attributes

This attribute executes an expression whenever signals are patched, useful for tracking data changes. The `patch` variable provides details. Modifiers support delaying, debouncing, and throttling the listener, with options for leading/trailing edges and view transitions.

```html
<div data-on-signal-patch="console.log('A signal changed!')"></div>
<div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
<div data-on-signal-patch__debounce.500ms="doSomething()"></div>
```

--------------------------------

### Sync Query String Params with data-query-string

Source: https://data-star.dev/reference/attributes

The `data-query-string` attribute synchronizes URL query parameters with Datastar signal values. It updates signals on load and query params on change. This Pro feature supports filtering and history management via `__filter` and `__history` modifiers.

```html
<div data-query-string></div>
```

```html
<div data-query-string="{include: /foo/, exclude: /bar/}"></div>
```

```html
<div data-query-string__filter__history></div>
```

--------------------------------

### Run Expressions on Element Resize with data-on-resize

Source: https://data-star.dev/reference/attributes

The `data-on-resize` attribute triggers a Datastar expression when an element's dimensions change. This Pro feature is useful for responsive UIs. It supports `__debounce` and `__throttle` modifiers for controlling event firing.

```html
<div data-on-resize="$count++"></div>
```

```html
<div data-on-resize__debounce.10ms="$count++"></div>
```

--------------------------------

### data-on-signal-patch-filter: Filter signals for patching

Source: https://data-star.dev/reference/attributes

The `data-on-signal-patch-filter` attribute allows filtering which signals trigger the `data-on-signal-patch` attribute. It accepts an object with `include` and/or `exclude` properties, which are regular expressions to match signal names.

```html
<!-- Only react to counter signal changes -->
<div data-on-signal-patch-filter="{include: /^counter$/}"></div>

<!-- React to all changes except those ending with "changes" -->
<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>

<!-- Combine include and exclude filters -->
<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
```

--------------------------------

### Datastar: Conditional Visibility with data-show

Source: https://data-star.dev/guide/reactive_signals

Shows or hides an element based on a boolean expression. Recommended to set initial `display: none;` to prevent flash of unwanted content.

```html
<input data-bind:foo />
<button data-show="$foo != ''">
    Save
</button>
```

```html
<input data-bind:foo />
<button data-show="$foo != ''" style="display: none;">
    Save
</button>
```

--------------------------------

### Attach Event Listeners (`data-on`)

Source: https://data-star.dev/reference

The `data-on` attribute attaches an event listener to an element, executing an expression when the event fires. The `evt` object is available in the expression. It supports custom events, default prevention for form submissions, and numerous modifiers for advanced control over event handling, including timing, casing, and target.

```html
<button data-on:click="$foo = ''">Reset</button>
<div data-on:myevent="$foo = evt.detail"></div>
<button data-on:click__window__debounce.500ms.leading="$foo = ''"></button>
<div data-on:my-event__case.camel="$foo = ''"></div>
```

--------------------------------

### Datastar Patch Elements: Remove Mode

Source: https://data-star.dev/reference/sse_events

Shows how to remove elements from the DOM using the 'remove' mode and a CSS selector. This is useful for dynamically clearing specific parts of the page.

```html
event: datastar-patch-elements
data: mode remove
data: selector #foo


```

--------------------------------

### Run Expressions on requestAnimationFrame with data-on-raf

Source: https://data-star.dev/reference/attributes

The `data-on-raf` attribute executes a Datastar expression on each `requestAnimationFrame` event. This Pro feature is ideal for animations and performance-sensitive updates. It supports `__throttle` modifiers for controlling execution frequency.

```html
<div data-on-raf="$count++"></div>
```

```html
<div data-on-raf__throttle.10ms="$count++"></div>
```

--------------------------------

### Node.js: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Node.js snippet uses the built-in `http` module and a custom `ServerSentEventGenerator` to create server-sent events with a dynamic interval duration. It calculates the duration based on the current seconds.

```javascript
1import { createServer } from "node:http";
 2import { ServerSentEventGenerator } from "../npm/esm/node/serverSentEventGenerator.js";
 3
 4const server = createServer(async (req, res) => {
 5  const currentTime = new Date();
 6  const duration = currentTime.getSeconds > 50 ? 5 : 1;
 7
 8  ServerSentEventGenerator.stream(req, res, (sse) => {
 9    sse.patchElements(`
```

--------------------------------

### Datastar Patch Signals: Conditional Update

Source: https://data-star.dev/reference/sse_events

Illustrates using the 'onlyIfMissing' option to update signals only if they do not already exist on the page. This prevents overwriting existing signal values with default ones.

```html
event: datastar-patch-signals
data: onlyIfMissing true
data: signals {foo: 1, bar: 2}


```

--------------------------------

### Apply Casing to Indicator Signal Name with data-indicator Modifier

Source: https://data-star.dev/reference

Similar to other attributes, the `__case` modifier for `data-indicator` allows you to specify the casing for the indicator signal name, with options for camel, kebab, snake, and pascal.

```html
<button data-on:click="@get('/endpoint')"
        data-indicator:fetching__case.pascal
></button>
```

--------------------------------

### Patch Signals with data-signals

Source: https://data-star.dev/reference/attributes

The data-signals directive allows you to add, update, or remove signals directly from your HTML. It supports single signal updates, nested signals using dot notation, and multiple signal patches using JavaScript/JSON object notation.

```html
<div data-signals:foo="1"></div>
<div data-signals:foo.bar="1"></div>
<div data-signals="{foo: {bar: 1, baz: 2}}"></div>
<div data-signals="{foo: null}"></div>
```

--------------------------------

### datastar-patch-elements SSE Event

Source: https://data-star.dev/reference/sse_events

Handles patching of DOM elements. It morphs elements based on their IDs or CSS selectors and supports various morphing modes.

```APIDOC
## datastar-patch-elements SSE Event

### Description

This SSE event type is used to patch one or more elements in the DOM. Datastar, by default, morphs elements by matching top-level elements based on their ID. Special handling is required for SVG morphing due to XML namespaces.

### Method

Server-Sent Events (SSE)

### Endpoint

N/A (SSE stream endpoint)

### Parameters

#### Event Data (`data` lines)

- **`elements <HTML>`** - Required - The HTML elements to patch. Can span multiple lines.
- **`mode [outer|inner|replace|prepend|append|before|after|remove]`** - Optional - Specifies how to morph the elements. Defaults to `outer`.
  - `outer`: Morphs the outer HTML of the elements.
  - `inner`: Morphs the inner HTML of the elements.
  - `replace`: Replaces the outer HTML of the elements.
  - `prepend`: Prepends the elements to the target’s children.
  - `append`: Appends the elements to the target’s children.
  - `before`: Inserts the elements before the target as siblings.
  - `after`: Inserts the elements after the target as siblings.
  - `remove`: Removes the target elements from the DOM.
- **`selector <CSS Selector>`** - Optional - Selects the target element of the patch using a CSS selector. Not required when using `outer` or `replace` modes.
- **`useViewTransition [true|false]`** - Optional - Whether to use view transitions when patching elements. Defaults to `false`.

### Request Example

```
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>
```

```
event: datastar-patch-elements
data: mode remove
data: selector #foo
```

```
event: datastar-patch-elements
data: mode inner
data: selector #foo
data: useViewTransition true
data: elements <div>
  data: elements        Hello world!
  data: elements </div>
```

### Response

#### Success Response (200 OK)

SSE stream containing `datastar-patch-elements` events.

#### Response Example

```
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>
```
```

--------------------------------

### Patch DOM Elements with Server-Sent Events (JavaScript)

Source: https://data-star.dev/guide

Shows how to patch DOM elements using Server-Sent Events in JavaScript. The ServerSentEventGenerator streams updates to the DOM, with a one-second timeout separating the two updates.

```javascript
// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
ServerSentEventGenerator.stream(req, res, (stream) => {
    // Patches elements into the DOM.
    stream.patchElements(`<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`);

    setTimeout(() => {
        stream.patchElements(`<div id="hal">Waiting for an order...</div>`);
    }, 1000);
});
```

--------------------------------

### Datastar Patch Signals: Remove Signals

Source: https://data-star.dev/reference/sse_events

Shows how to remove signals by setting their values to null within the 'signals' data line. This is useful for clearing specific state on the page.

```html
event: datastar-patch-signals
data: signals {foo: null, bar: null}


```

--------------------------------

### Integrating External JavaScript: Asynchronous Function with Custom Event

Source: https://data-star.dev/guide/datastar_expressions

Shows how to handle asynchronous JavaScript functions within Datastar expressions. Since Datastar doesn't await async code, the function dispatches a custom event with the result, which is then captured by another `data-*` attribute.

```html
1<div data-signals:result>
2    <input data-bind:foo 
3           data-on:input="myfunction(el, $foo)"
4           data-on:mycustomevent__window="$result = evt.detail.value"
5    >
6    <span data-text="$result"></span>
7</div>
```

--------------------------------

### Morph Multiple SVG Elements with Datastar

Source: https://data-star.dev/examples/svg_morphing

Shows how to update multiple SVG elements simultaneously using Datastar SSE. This Go code morphs three circles, each with a random color and radius, identified by the 'multi-demo' SVG.

```go
svgMorphingRouter.Get("/multiple_elements", func(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    color1 := svgColors[rand.N(len(svgColors))]
    color2 := svgColors[rand.N(len(svgColors))]
    color3 := svgColors[rand.N(len(svgColors))]
    r1 := 10 + rand.N(20) // radius 10-30
    r2 := 10 + rand.N(20)
    r3 := 10 + rand.N(20)
    sse.PatchElements(fmt.Sprintf(`<svg id="multi-demo">
        <circle cx="30" cy="30" r="%d" fill="%s" />
        <circle cx="70" cy="30" r="%d" fill="%s" />
        <circle cx="50" cy="70" r="%d" fill="%s" />
    </svg>`, r1, color1, r2, color2, r3, color3))
})
```

--------------------------------

### Attach Event Listeners with data-on Attribute

Source: https://data-star.dev/guide/reactive_signals

The `data-on` attribute allows attaching event listeners to DOM elements. When an event is triggered, an associated expression is evaluated, enabling dynamic updates to signals or other reactive states.

```html
<input data-bind:foo />
<button data-on:click="$foo = ''">
    Reset
</button>
```

--------------------------------

### C# ASP.NET Core: Dynamic SSE Interval Duration

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This C# snippet for ASP.NET Core uses `IDatastarService` to send server-sent events with a dynamically adjusted interval duration. It formats the current time and determines the duration based on the current second.

```csharp
1using StarFederation.Datastar.DependencyInjection;
 2
 3app.MapGet("/endpoint", async (IDatastarService datastarService) =>
 4{
 5    var currentTime = DateTime.Now.ToString("yyyy-MM-dd hh:mm:ss");
 6    var currentSeconds = DateTime.Now.Second;
 7    var duration = currentSeconds < 50 ? 5 : 1;
 8    await datastarService.PatchElementsAsync($"""
 9        <div id="time" data-on-interval__duration.{duration}s="@get('/endpoint')">
10            {currentTime}
11        </div>
12    """);
13});
```

--------------------------------

### Persist Signals in Storage with data-persist

Source: https://data-star.dev/reference/attributes

The `data-persist` attribute stores Datastar signals in local storage, preserving values across page loads. This Pro feature can be configured with include/exclude regex filters and custom storage keys. It supports `__session` modifier for using session storage.

```html
<div data-persist></div>
```

```html
<div data-persist="{include: /foo/, exclude: /bar/}"></div>
```

```html
<div data-persist:mykey></div>
```

```html
<!-- Persists signals using a custom key `mykey` in session storage -->
<div data-persist:mykey__session></div>
```

--------------------------------

### Display Reactive JSON Signals (`data-json-signals`)

Source: https://data-star.dev/reference

The `data-json-signals` attribute displays a reactive, JSON-stringified version of signals in an element's text content. It's useful for debugging. Optional filters (`include`, `exclude`) can specify which signals to display, and the `__terse` modifier can compact the output.

```html
<!-- Display all signals -->
<pre data-json-signals></pre>

<!-- Only show signals that include "user" in their path -->
<pre data-json-signals="{include: /user/}"></pre>

<!-- Show all signals except those ending with "temp" -->
<pre data-json-signals="{exclude: /temp$/}"></pre>

<!-- Combine include and exclude filters -->
<pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>

<!-- Display filtered signals in a compact format -->
<pre data-json-signals__terse="{include: /counter/}"></pre>
```

--------------------------------

### Run Expression on Viewport Intersection (`data-on-intersect`)

Source: https://data-star.dev/reference

The `data-on-intersect` attribute runs an expression when the element intersects with the viewport. This is useful for triggering actions or animations when an element becomes visible.

```html
<div data-on-intersect="$intersected = true"></div>
```

--------------------------------

### Datastar: Conditional Class Addition with data-class

Source: https://data-star.dev/guide/reactive_signals

Adds or removes an element's class based on an expression. Can manage multiple classes using key-value pairs.

```html
<input data-bind:foo />
<button data-class:success="$foo != ''">
    Save
</button>
```

```html
<button data-class="{success: $foo != '', 'font-bold': $foo == 'bar'}">
    Save
</button>
```

--------------------------------

### Execute Side Effects with data-effect

Source: https://data-star.dev/reference/attributes

The `data-effect` attribute executes an expression on page load and whenever signals within the expression change. This is ideal for side effects like updating other signals, making API requests, or DOM manipulation.

```html
<div data-effect="$foo = $bar + $baz"></div>
```

--------------------------------

### DataStar SSE for Appending Script to Body

Source: https://data-star.dev/guide/datastar_expressions

A Server-Sent Event (SSE) instructing DataStar to append a script tag directly to the document's body. This is a method for executing JavaScript without necessarily patching existing DOM elements.

```sse
event: datastar-patch-elements
data: mode append
data: selector body
data: elements <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>


```

--------------------------------

### Access Signals Safely with @peek() in DataStar

Source: https://data-star.dev/reference/actions

The @peek() action allows accessing signal values within DataStar expressions without subscribing to their changes. This prevents the expression from re-evaluating when the peeked signal updates. It takes a callable function as an argument that returns the signal value.

```html
<div data-text="$foo + @peek(() => $bar)"></div>
```

--------------------------------

### Datastar Expression: Using the 'el' Variable

Source: https://data-star.dev/guide/datastar_expressions

Illustrates how the 'el' variable, available in every Datastar expression, can be used to access properties of the element the attribute is attached to, such as its ID.

```html
1<div id="foo" data-text="el.id"></div>
```

--------------------------------

### DataStar SSE for Script Execution within Elements

Source: https://data-star.dev/guide/datastar_expressions

A Server-Sent Event (SSE) formatted response from DataStar. This specific event, 'datastar-patch-elements', includes instructions to replace or update elements, embedding a script tag that will execute JavaScript in the browser.

```sse
event: datastar-patch-elements
data: elements <div id="hal">
data: elements     <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
data: elements </div>


```

--------------------------------

### Set HTML Attribute with data-attr

Source: https://data-star.dev/reference/attributes

The `data-attr` attribute sets the value of any HTML attribute to an expression and keeps it in sync. It can be used to set a single attribute or multiple attributes using a key-value pair object.

```html
<div data-attr:title="$foo"></div>
```

```html
<div data-attr="{title: $foo, disabled: $bar}"></div>
```

--------------------------------

### Change Circle Radius with Datastar

Source: https://data-star.dev/examples/svg_morphing

This Go code snippet illustrates morphing an SVG circle's radius using Datastar. It generates a random radius and updates the 'size-demo' SVG element, changing the circle's size.

```go
svgMorphingRouter.Get("/circle_size", func(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    radius := 15 + rand.N(45) // Random radius between 15-60
    sse.PatchElements(fmt.Sprintf(`<svg id="size-demo"><circle cx="50" cy="50" r="%d" fill="green" /></svg>`, radius))
})
```

--------------------------------

### Scroll Element into View with data-scroll-into-view

Source: https://data-star.dev/reference/attributes

The `data-scroll-into-view` attribute automatically scrolls the associated element into the viewport. This Pro feature is particularly useful when dynamically adding content via backend requests, ensuring the new content is visible to the user.

```html
<div data-scroll-into-view></div>
```

--------------------------------

### Skip Morphing Elements with data-ignore-morph

Source: https://data-star.dev/reference/attributes

The `data-ignore-morph` attribute instructs the `PatchElements` watcher to skip morphing a specific element and its children. To remove it, re-patch the element without the attribute.

```html
<div data-ignore-morph>
    This element will not be morphed.
</div>
```

--------------------------------

### Bind Element Text Content with data-text

Source: https://data-star.dev/reference/attributes

The `data-text` attribute binds the text content of an HTML element to a Datastar expression. This allows for dynamic updates of an element's text based on signal changes. No specific dependencies are required beyond the Datastar framework itself.

```html
<div data-text="$foo"></div>
```

--------------------------------

### JSON Payload for Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This JSON object represents the payload used to patch signals. It follows the JSON Merge Patch RFC 7396 standard and can contain multiple signal updates.

```json
1{"hal": "Affirmative, Dave. I read you."}
```

--------------------------------

### Replace URL without Reloading with data-replace-url

Source: https://data-star.dev/reference/attributes

The `data-replace-url` attribute replaces the browser's current URL without triggering a page reload. The value of the attribute is an evaluated Datastar expression that resolves to the new URL. This Pro feature allows for dynamic, client-side navigation.

```html
<div data-replace-url="`/page${page}`"></div>
```

--------------------------------

### data-on-intersect: Trigger once or when element is half/fully visible

Source: https://data-star.dev/reference/attributes

The `data-on-intersect` attribute triggers an expression when an element intersects with the viewport. Modifiers like `__once`, `__half`, and `__full` control the trigger conditions. It also supports debouncing, throttling, and delaying the event listener, along with view transition integration.

```html
<div data-on-intersect__once__full="$fullyIntersected = true"></div>
```

--------------------------------

### Modify data-signals Naming and Behavior with Modifiers

Source: https://data-star.dev/reference/attributes

data-signals supports the __case modifier for signal name casing (camel, kebab, snake, pascal) and __ifmissing to conditionally patch signals only if they don't already exist, useful for setting defaults.

```html
<div data-signals:my-signal__case.kebab="1"
     data-signals:foo__ifmissing="1"
></div>
```

--------------------------------

### Modify data-ref Signal Naming with __case

Source: https://data-star.dev/reference/attributes

The __case modifier for data-ref allows you to control the casing of the signal name. Supported cases include camel, kebab, snake, and pascal, ensuring signal names match your preferred convention.

```html
<div data-ref:my-signal__case.kebab></div>
```

--------------------------------

### Conditionally Show/Hide Elements with data-show

Source: https://data-star.dev/reference/attributes

The data-show directive controls the visibility of an element based on a boolean expression. To prevent flickering, it's recommended to initially hide the element using 'display: none'.

```html
<div data-show="$foo"></div>
<div data-show="$foo" style="display: none"></div>
```

--------------------------------

### Datastar: Derived Signals with data-computed

Source: https://data-star.dev/guide/reactive_signals

Creates a read-only signal derived from a reactive expression. Its value updates automatically when signals within the expression change. Useful for memoizing expressions.

```html
<input data-bind:foo />
<div data-computed:repeated="$foo.repeat(2)" data-text="$repeated"></div>
```

--------------------------------

### data-preserve-attr: Maintain attribute values during DOM morphing

Source: https://data-star.dev/reference/attributes

The `data-preserve-attr` attribute ensures that specified attribute values are maintained when DOM elements are morphed. Multiple attributes can be preserved by separating their names with a space.

```html
<details open data-preserve-attr="open">
    <summary>Title</summary>
    Content
</details>
<details open class="foo" data-preserve-attr="open class">
    <summary>Title</summary>
    Content
</details>
```

--------------------------------

### Ignore Elements and Descendants with data-ignore

Source: https://data-star.dev/reference/attributes

The `data-ignore` attribute prevents Datastar from processing an element and its descendants. This is useful for third-party libraries or to avoid conflicts with user-generated content.

```html
<div data-ignore data-show-thirdpartylib="">
    <div>
        Datastar will not process this element.
    </div>
</div>
```

--------------------------------

### Create Computed Signals with data-computed

Source: https://data-star.dev/reference/attributes

The `data-computed` attribute creates a read-only signal computed from an expression. Its value updates automatically when signals in the expression change. It can also define computed signals using key-value pairs with callables.

```html
<div data-computed:foo="$bar + $baz"></div>
```

```html
<div data-computed:foo="$bar + $baz"></div>
<div data-text="$foo"></div>
```

```html
<div data-computed="{foo: () => $bar + $baz}"></div>
```

--------------------------------

### Conditional Class Addition with data-class

Source: https://data-star.dev/reference/attributes

The `data-class` attribute adds or removes a class from an element based on an expression's evaluation. It can manage single classes or multiple classes simultaneously using an object of class names and their corresponding expressions.

```html
<div data-class:hidden="$foo"></div>
```

```html
<div data-class="{hidden: $foo, 'font-bold': $bar}"></div>
```

--------------------------------

### Modify Computed Signal Name Casing with data-computed Modifier

Source: https://data-star.dev/reference/attributes

The `__case` modifier for `data-computed` allows you to convert the casing of the computed signal name. Supported cases include camel (default), kebab, snake, and pascal.

```html
<div data-computed:my-signal__case.kebab="$bar + $baz"></div>
```

--------------------------------

### Add Custom Input Validation with data-custom-validity

Source: https://data-star.dev/reference/attributes

The `data-custom-validity` attribute allows for custom validation logic on form elements using a Datastar expression. The expression must evaluate to a string: an empty string means valid, a non-empty string is the validation error message. This Pro feature requires the Datastar Pro license.

```html
<form>
    <input data-bind:foo name="foo" />
    <input data-bind:bar name="bar"
           data-custom-validity="$foo === $bar ? '' : 'Values must be the same.'"
    />
    <button>Submit form</button>
</form>
```

--------------------------------

### Ignore Only Self with data-ignore__self Modifier

Source: https://data-star.dev/reference/attributes

The `__self` modifier for `data-ignore` specifies that only the element itself should be ignored, allowing its descendants to be processed by Datastar.

=== COMPLETE CONTENT === This response contains all available snippets from this library. No additional content exists. Do not make further requests.