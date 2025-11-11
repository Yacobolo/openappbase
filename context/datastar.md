### Datastar Signal Nesting Examples

Source: https://data-star.dev/docs

Demonstrates how to nest signals using dot-notation, object syntax, and two-way binding in HTML. This allows for more granular targeting of state on the backend, useful for managing repeated state like menu open/closed status.

```html
<div data-signals-foo.bar="1"></div>

```

```html
<div data-signals="{foo: {bar: 1}}"></div>

```

```html
<input data-bind-foo.bar />

```

```html
<div data-signals="{menu: {isOpen: {desktop: false, mobile: false}}}">
    <button data-on-click="@toggleAll({include: /^menu\.isOpen\./})">
        Open/close menu
    </button>
</div>

```

--------------------------------

### Combine data-signals, data-on, and data-text for Reactivity

Source: https://data-star.dev/guide/reactive_signals

This example demonstrates frontend reactivity by combining `data-signals` for initialization, `data-on-click` for user interaction, and `data-text` to display signal values. The button click updates the `$hal` signal, which is then reflected in the text content.

```html
<div data-signals-hal="'...'">
    <button data-on-click="$hal = 'Affirmative, Dave. I read you.'">
        HAL, do you read me?
    </button>
    <div data-text="$hal"></div>
</div>
```

--------------------------------

### Read Datastar Signals in PHP

Source: https://data-star.dev/docs

A simple PHP example showing how to read all signals from the current request using the `ServerSentEventGenerator::readSignals()` static method.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Reads all signals from the request.
$signals = ServerSentEventGenerator::readSignals();

```

--------------------------------

### Executing Backend Scripts via Button Click (HTML)

Source: https://data-star.dev/docs

This HTML demonstrates triggering a backend script execution when a button is clicked. The `@get('/endpoint')` syntax initiates a GET request to the specified endpoint. If the response is JavaScript, it will be executed on the frontend.

```html
<button data-on-click="@get('/endpoint')">
    What are you talking about, HAL?
</button>
```

--------------------------------

### Stream SSE Events in Rust

Source: https://data-star.dev/docs

This Rust example demonstrates streaming Server-Sent Events (SSE) using the `datastar` crate. It shows how to create an `Sse` stream that yields `PatchElements` and `PatchSignals` events.

```rust
use datastar::prelude::*;
use async_stream::stream;

Sse(stream! {
    // Patches elements into the DOM.
    yield PatchElements::new("<div id='question'>What do you put in a toaster?</div>").into();

    // Patches signals.
    yield PatchSignals::new("{response: '', answer: 'bread'}").into();
})
```

--------------------------------

### Stream SSE Events in C#

Source: https://data-star.dev/docs

This C# example illustrates how to configure a backend endpoint to stream Server-Sent Events (SSE) using DataStar. It shows how to add DataStar as a service and then use the `IDatastarService` to patch elements and signals.

```csharp
using StarFederation.Datastar.DependencyInjection;

// Adds Datastar as a service
builder.Services.AddDatastar();

app.MapGet("/", async (IDatastarService datastarService) =>
{
    // Patches elements into the DOM.
    await datastarService.PatchElementsAsync(@"<div id=\"question\">What do you put in a toaster?</div>");

    // Patches signals.
    await datastarService.PatchSignalsAsync(new { response = "", answer = "bread" });
});
```

--------------------------------

### C# DataStar ExecuteScriptAsync and PatchElementsAsync

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This C# example illustrates how to use `PatchElementsAsync` and `ExecuteScriptAsync` from the DataStar library for server-initiated redirects. It includes a `Task.Delay` to simulate a waiting period before the redirection.

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

### Install Datastar Locally

Source: https://data-star.dev/guide/getting_started

Integrate Datastar by hosting the script file yourself. Download the script or create your own bundle and reference it from your project's path.

```html
<script type="module" src="/path/to/datastar.js"></script>
```

--------------------------------

### Stream SSE Events in Python

Source: https://data-star.dev/docs

This Python example demonstrates streaming Server-Sent Events (SSE) using the `datastar_py` library, particularly with Litestar. It shows how to return a `DatastarResponse` containing multiple SSE events, including patching elements and signals.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.litestar import DatastarResponse

async def endpoint():
    return DatastarResponse([
        SSE.patch_elements('<div id="question">What do you put in a toaster?</div>'),
        SSE.patch_signals({"response": "", "answer": "bread"})
    ])
```

--------------------------------

### Listening for Data-Star Fetch Events

Source: https://data-star.dev/reference/actions

Provides an example of how to listen for `datastar-fetch` events dispatched during the lifecycle of a Data-Star fetch request. This allows for custom handling of events like 'error', 'started', 'finished', and more, enabling dynamic UI updates or logging based on request status.

```html
<div data-on-datastar-fetch="
    evt.detail.type === 'error' && console.log('Fetch error encountered')
"></div>
```

--------------------------------

### HTML Button for DataStar Redirect

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This snippet shows how to add a `data-on-click` attribute to an HTML button, which triggers a GET request to a specified backend endpoint when clicked. It also includes a placeholder `div` for visual feedback during the redirection process.

```html
<button data-on-click="@get('/endpoint')">
    Click to be redirected from the backend
</button>
<div id="indicator"></div>
```

--------------------------------

### Kotlin DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Kotlin example demonstrates using the `ServerSentEventGenerator` to patch elements and execute a script for redirection, including a 3-second delay.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements =
        """
        <div id="indicator">Redirecting in 3 seconds...</div>
        """.trimIndent()
)

Thread.sleep(3 * ONE_SECOND)

generator.executeScript(
    script = "window.location.href = '/success'",
)
```

--------------------------------

### Reading Signals in Ruby

Source: https://data-star.dev/guide/backend_requests

A Ruby example showing how to initialize DataStar with a request and response, and then access signals.

```ruby
# Setup with request
datastar = Datastar.new(request:, response:)

# Read signals
some_signal = datastar.signals[:some_signal]
```

--------------------------------

### Install Datastar via CDN

Source: https://data-star.dev/guide/getting_started

Include the Datastar framework in your project by adding this script tag to your HTML. This is the quickest method for integration and fetches the latest version from the CDN.

```html
<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.5/bundles/datastar.js"></script>
```

--------------------------------

### Configure Datastar Action with Options

Source: https://data-star.dev/docs

Example of configuring a Datastar action with multiple options including signal filtering, custom headers, and request cancellation. This demonstrates how to send specific signals, add authentication tokens, keep connections open when hidden, and disable automatic request cancellation.

```html
<button data-on-click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})"></button>
```

--------------------------------

### Stream SSE Events in PHP

Source: https://data-star.dev/docs

This PHP example shows how to stream Server-Sent Events (SSE) using the DataStar PHP library. It demonstrates creating a `ServerSentEventGenerator` instance and using its methods to patch DOM elements and signals.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
$sse = new ServerSentEventGenerator();

// Patches elements into the DOM.
$sse->patchElements(
    '<div id="question">What do you put in a toaster?</div>'
);

// Patches signals.
$sse->patchSignals(['response' => '', 'answer' => 'bread']);
```

--------------------------------

### Python Redirect with DataStar-py

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a redirect in Python using the datastar_py library. This asynchronous example yields SSE events for patching elements, sleeping, and executing a redirect script.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.execute_script('setTimeout(() => window.location = "/guide")')

```

--------------------------------

### Send GET Request to Patch Elements with Datastar

Source: https://data-star.dev/docs

This snippet demonstrates how to use the `@get()` action in Datastar to send a GET request to a specified URL. The response, if HTML, will be used to morph elements into the DOM. It requires the Datastar library.

```html
1<button data-on-click="@get('/endpoint')">
2    Open the pod bay doors, HAL.
3</button>
4<div id="hal"></div>
```

--------------------------------

### C# SDK Example for Dynamic List Loading

Source: https://data-star.dev/how_tos/load_more_list_items

This C# example uses ASP.NET Core and the DataStar SDK to implement dynamic list loading. It defines a signal record, reads the offset signal, and then uses the IDatastarService to patch elements into the list, update signals, or remove the load-more button.

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

### Server-Sent Events for DataStar Redirect (SSE)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This SSE example demonstrates how a backend can initiate a redirect. It first sends a `datastar-patch-elements` event to update an indicator, then waits for 3 seconds, and finally sends another SSE event to append a script that redirects the user.

```sse
event: datastar-patch-elements
data: elements <div id="indicator">Redirecting in 3 seconds...</div>

// Wait 3 seconds

event: datastar-patch-elements
selector body
mode append
data: elements <script>window.location.href = "/guide"</script>
```

--------------------------------

### Go Redirect with Server-Sent Events (SSE)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Handles redirects in Go using Server-Sent Events (SSE) by patching elements, sleeping for a duration, and then redirecting the client. This example uses the datastar-go library.

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

### Node.js DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Node.js example demonstrates using the `ServerSentEventGenerator` to stream SSE events for patching elements and executing a script to redirect the user after a 3-second delay using `setTimeout`.

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

### Clojure Redirect with DataStar API

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates a redirect in Clojure using the DataStar API. It sets up an SSE response that patches elements, pauses for 3 seconds, and then redirects the client.

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

### Rust DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Rust example utilizes DataStar's SSE capabilities to patch elements and execute a script for redirection, incorporating a Tokio-based delay.

```rust
use datastar::prelude::*;
use async_stream::stream;
use core::time::Duration;

Sse(stream! {
    yield PatchElements::new("<div id='indicator'>Redirecting in 3 seconds...</div>").into();
    tokio::time::sleep(core::time::Duration::from_secs(3)).await;
    yield ExecuteScript::new("window.location = '/guide'").into();
});
```

--------------------------------

### Read Datastar Signals in Go

Source: https://data-star.dev/docs

Provides an example of reading nested signals from an HTTP request in Go using the `datastar-go` library. It defines a `Signals` struct to match the expected JSON payload and uses `datastar.ReadSignals` to unmarshal the data.

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

### Multiple Element Patches via SSE with Datastar

Source: https://data-star.dev/docs

This example demonstrates sending multiple 'datastar-patch-elements' SSE events in a stream to update the DOM sequentially. This allows for dynamic content changes over time, such as showing a response and then updating it.

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
```

--------------------------------

### Submit Form Data using data-on-submit with ContentType 'form'

Source: https://data-star.dev/examples/form_data

This example shows how to initiate a GET request with contentType set to 'form' when a form is submitted. The @get() action is triggered by the 'data-on-submit' attribute on the form element, automatically sending the form's data.

```html
<form data-on-submit="@get('/endpoint', {contentType: 'form'})">
    foo: <input type="text" name="foo" required />
    <button>
        Submit form
    </button>
</form>
```

--------------------------------

### Clojure SDK Example for Dynamic List Loading

Source: https://data-star.dev/how_tos/load_more_list_items

This Clojure code demonstrates how to handle requests and send DataStar events using the SDK. It reads signals to get the current offset, patches new elements to the list in append mode, and either patches new signals or removes the load-more button based on the offset.

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
28         (d*/close-sse! sse)))}))

```

--------------------------------

### Patch Elements using Datastar Kotlin SDK

Source: https://data-star.dev/docs

This Kotlin code snippet shows how to use the ServerSentEventGenerator to patch HTML elements into the DOM. It includes an example of sending an initial patch and then updating it after a delay.

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

```

--------------------------------

### Stream SSE Events in Java

Source: https://data-star.dev/docs

This Java example shows how to stream Server-Sent Events (SSE) using the DataStar Java SDK. It involves creating a `ServerSentEventGenerator` and using its `send` method with `PatchElements` and `PatchSignals` builders to send data.

```java
import starfederation.datastar.utils.ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
AbstractResponseAdapter responseAdapter = new HttpServletResponseAdapter(response);
ServerSentEventGenerator generator = new ServerSentEventGenerator(responseAdapter);

// Patches elements into the DOM.
generator.send(PatchElements.builder()
    .data("<div id=\"question\">What do you put in a toaster?</div>")
    .build()
);

// Patches signals.
generator.send(PatchSignals.builder()
    .data("{\"response\": \"\", \"answer\": \"\"}")
    .build()
);
```

--------------------------------

### Alpine.js data-on-interval Basic Usage (JavaScript)

Source: https://data-star.dev/docs

A simple example of the 'data-on-interval' directive, which repeatedly executes an expression at a default interval (one second). This snippet increments a counter on each interval.

```html
<div data-on-interval="$count++"></div>

```

--------------------------------

### HTML Button for Click-to-Load

Source: https://data-star.dev/examples/click_to_load

This HTML snippet represents the button that initiates the 'click-to-load' functionality. It includes attributes for dynamic styling based on a fetching state and an event handler for the click action. The `data-on-click` attribute specifies the function to call, which in turn fetches more data.

```html
<button
    class="info wide"
    data-indicator-_fetching
    data-attr-aria-disabled="`${$_fetching}`"
    data-on-click="!$_fetching && @get('/examples/click_to_load/more')">
    Load More
</button>
```

--------------------------------

### Rust Redirect with DataStar Prelude

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Shows a redirect implementation in Rust using the datastar crate. It leverages async streams to send SSE events for patching elements, delaying, and executing a script for redirection.

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

### Python DataStar ExecuteScript and PatchElements (Sanic)

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Python example uses DataStar with the Sanic framework to perform server-initiated redirects. It yields SSE events for patching elements and executing a script, with an asynchronous delay.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response
import asyncio

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.execute_script('window.location = "/guide"')
```

--------------------------------

### PHP Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Shows how to perform a redirect in PHP using the ServerSentEventGenerator class. This includes patching elements, a short sleep, and then executing the redirect script.

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

### Go DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Go snippet shows the implementation of DataStar's `ExecuteScript` and `PatchElements` methods using Server-Sent Events. It includes a 3-second sleep before executing the script for redirection.

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

### Kotlin Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates how to implement redirects in Kotlin using the ServerSentEventGenerator. It involves patching HTML elements, pausing execution, and then initiating a redirect.

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

### Backend Redirect with Python Sanic Server-Sent Events

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Python example using the Sanic framework shows how to perform a backend redirect. It yields SSE events to patch an indicator and then initiate a redirect after a 3-second asynchronous sleep.

```python
from datastar_py import ServerSentEventGenerator as SSE
from datastar_py.sanic import datastar_response

@app.get("/redirect")
@datastar_response
async def redirect_from_backend():
    yield SSE.patch_elements('<div id="indicator">Redirecting in 3 seconds...</div>')
    await asyncio.sleep(3)
    yield SSE.redirect("/guide")
```

--------------------------------

### Java (Servlet): Patching Signals with SSE

Source: https://data-star.dev/docs

This Java example demonstrates how to patch signals using Server-Sent Events (SSE) within a Servlet environment using the DataStar SDK. It creates a `ServerSentEventGenerator` and uses its `send` method with `PatchSignals.builder()` to emit signal updates.

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
19
```

--------------------------------

### PHP DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This PHP snippet shows how to use the `ServerSentEventGenerator` to patch elements and execute a script for redirection, with a 3-second sleep in between.

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

### Alpine.js data-on-intersect with Modifiers (JavaScript)

Source: https://data-star.dev/docs

Demonstrates advanced usage of 'data-on-intersect' with modifiers like '__once' and '__full'. This example triggers an expression only once when the entire element is visible in the viewport.

```html
<div data-on-intersect__once__full="$fullyIntersected = true"></div>

```

--------------------------------

### Sequential SSE Signal Patching

Source: https://data-star.dev/guide/reactive_signals

This SSE example demonstrates sending multiple 'datastar-patch-signals' events in a stream, including a delay. This pattern is useful for orchestrating a sequence of UI updates, such as showing a response and then clearing it after a period.

```sse
1event: datastar-patch-signals
2data: signals {hal: 'Affirmative, Dave. I read you.'}
3
4// Wait 1 second
5
6event: datastar-patch-signals
7data: signals {hal: '...'}
```

--------------------------------

### Implement File Upload Progress with DataStar Pro

Source: https://data-star.dev/docs

This example demonstrates how to monitor file upload progress using DataStar Pro. It involves an HTML form configured for multipart/form-data submission. The `data-on-datastar-fetch` attribute is used to capture `upload-progress` events, updating signal values for `$uploading` and `$progress`, and displaying a progress bar.

```html
 1<form enctype="multipart/form-data"
 2    data-signals="{progress: 0, uploading: false}"
 3    data-on-submit__prevent="@post('https://example.com/upload', {contentType: 'form'})"
 4    data-on-datastar-fetch="
 5        if (evt.detail.type !== 'upload-progress') return;
 6
 7        const {progress, loaded, total} = evt.detail.argsRaw;
 8        $uploading = true;
 9        $progress = Number(progress);
10
11        if ($progress >= 100) {
12            $uploading = false;
13        }
14    "
15>
16    <input type="file" name="files" multiple />
17    <button type="submit">Upload</button>
18    <progress data-show="$uploading" data-attr-value="$progress" max="100"></progress>
19</form>
```

--------------------------------

### Executing JavaScript via Server-Sent Events (SSE) (HTML)

Source: https://data-star.dev/docs

This example demonstrates executing JavaScript using Server-Sent Events (SSE). A `script` tag is embedded within a `datastar-patch-elements` SSE event, allowing for dynamic script execution on the client-side.

```text
event: datastar-patch-elements
data: elements <div id="hal">
data: elements     <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
data: elements </div>
```

--------------------------------

### Go SDK Example for Dynamic List Loading

Source: https://data-star.dev/how_tos/load_more_list_items

This Go code demonstrates using the DataStar SDK to handle dynamic list loading. It reads offset signals, and then uses the SSE object to patch elements into the '#list' container with append mode, update signals, or remove the load-more button.

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

### Appending JavaScript to Body via SSE (HTML)

Source: https://data-star.dev/docs

This SSE example shows how to append a script tag directly to the `body` element of the DOM. This is useful for executing scripts that don't require specific element placement or patching.

```text
event: datastar-patch-elements
data: mode append
data: selector body
data: elements <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
```

--------------------------------

### Read Datastar Signals in Ruby

Source: https://data-star.dev/docs

Shows how to initialize Datastar with a request and response object in Ruby, and then access a specific signal (e.g., `some_signal`) using attribute-style access on the `datastar.signals` object.

```ruby
# Setup with request
datastar = Datastar.new(request:, response:)

# Read signals
some_signal = datastar.signals[:some_signal]

```

--------------------------------

### PHP: Patching Signals with SSE

Source: https://data-star.dev/docs

This PHP example shows how to patch signals using Server-Sent Events (SSE) with the DataStar SDK. It creates a `ServerSentEventGenerator` instance and uses its `patchSignals` method to send signal data, including a `sleep` call for a delay.

```php
 1use starfederation\datastar\ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4$sse = new ServerSentEventGenerator();
 5
 6// Patches signals.
 7$sse->patchSignals(['hal' => 'Affirmative, Dave. I read you.']);
 8
 9sleep(1)
10
11$sse->patchSignals(['hal' => '...']);
12
```

--------------------------------

### HTML Structure for Progressive Loading

Source: https://data-star.dev/examples/progressive_load

This HTML snippet defines the main structure of a page designed for progressive loading. It includes a button to trigger the loading and placeholders for content sections like header, article, and comments, which are intended to be populated dynamically via SSE.

```html
<div>
    <div class="actions">
        <button
            id="load-button"
            data-signals-load-disabled="false"
            data-on-click="$loadDisabled=true; @get('/examples/progressive_load/updates')"
            data-attr-disabled="$loadDisabled"
            data-indicator-progressive-Load
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

### PHP: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This PHP example demonstrates patching signals using Server-Sent Events (SSE) with the DataStar PHP library. It shows creating a `ServerSentEventGenerator` instance and using the `patchSignals` method to send signal updates.

```php
 1use starfederation\datastar\ServerSentEventGenerator;
 2
 3// Creates a new `ServerSentEventGenerator` instance.
 4$sse = new ServerSentEventGenerator();
 5
 6// Patches signals.
 7$sse->patchSignals(['hal' => 'Affirmative, Dave. I read you.']);
 8
 9sleep(1)
10
11$sse->patchSignals(['hal' => '...']);
```

--------------------------------

### C# ASP.NET Core: Patch Signals

Source: https://data-star.dev/guide/reactive_signals

This C# example shows how to use the DataStar SDK within an ASP.NET Core application to patch signals. It demonstrates adding DataStar as a service and then using `PatchSignalsAsync` to send updates, including a delay between them.

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

### Loading Indicator with data-indicator-fetching

Source: https://data-star.dev/guide/backend_requests

This HTML example utilizes the `data-indicator-fetching` attribute to manage a loading state. When a GET request is in flight (triggered by `data-on-click`), the element with `data-indicator-fetching` will have its associated signal set to true, allowing conditional styling like a loading indicator via `data-class-loading`.

```html
<div id="question"></div>
<button
    data-on-click="@get('/actions/quiz')"
    data-indicator-fetching
>
    Fetch a question
</button>
<div data-class-loading="$fetching" class="indicator"></div>
```

--------------------------------

### Clojure DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Clojure snippet demonstrates using DataStar's `execute-script!` and `patch-elements!` functions within an SSE response. It includes a 3-second delay before executing the script for redirection.

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

### Alpine.js data-on-interval with Duration Modifier (JavaScript)

Source: https://data-star.dev/docs

Shows how to customize the interval duration using the '__duration' modifier with 'data-on-interval'. This example sets the interval to 500 milliseconds, causing the counter to increment twice as fast.

```html
<div data-on-interval__duration.500ms="$count++"></div>

```

--------------------------------

### Patch Elements via SSE with Datastar

Source: https://data-star.dev/docs

This example shows how to send Server-Sent Events (SSE) with the 'datastar-patch-elements' event type to update the DOM. It's useful for real-time updates and long-lived connections. The 'morph' strategy is used by default.

```html
1event: datastar-patch-elements
2data: elements <div id="hal">
3data: elements     I’m sorry, Dave. I’m afraid I can’t do that.
4data: elements </div>
```

--------------------------------

### Alpine.js data-on-load with Delay Modifier (JavaScript)

Source: https://data-star.dev/docs

Illustrates using the '__delay' modifier with 'data-on-load' to postpone the execution of the directive's expression. This example waits for 500 milliseconds after the element loads before initializing the count variable.

```html
<div data-on-load__delay.500ms="$count = 1"></div>

```

--------------------------------

### Initialize ServerSentEventGenerator

Source: https://data-star.dev/docs

This JavaScript snippet shows the initialization of the `ServerSentEventGenerator` which also sends the necessary headers for SSE. This is a prerequisite for streaming events.

```javascript
// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
ServerSentEventGenerator.stream(req, res, (stream) => {
```

--------------------------------

### Alpine.js data-on-click with Modifiers (JavaScript)

Source: https://data-star.dev/docs

Demonstrates using the 'data-on-click' attribute with modifiers like '__window', '__debounce', and '.leading' to control event triggering behavior. This example attaches a click listener to the window that only triggers once after a 500ms debounce with a leading edge execution.

```html
<button data-on-click__window__debounce.500ms.leading="$foo = ''"></button>

```

--------------------------------

### Reading Signals in PHP

Source: https://data-star.dev/guide/backend_requests

A PHP example demonstrating how to read all signals from an incoming request using the `ServerSentEventGenerator` class.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Reads all signals from the request.
$signals = ServerSentEventGenerator::readSignals();
```

--------------------------------

### HTML Input for Active Search with Debounce

Source: https://data-star.dev/examples/active_search

This HTML input element facilitates active searching. The `data-bind-search` attribute likely binds the input's value for searching, while `data-on-input__debounce.200ms` specifies that a GET request to '/examples/active_search/search' should be made 200ms after the user stops typing, passing the input value.

```html
<input
    type="text"
    placeholder="Search..."
    data-bind-search
    data-on-input__debounce.200ms="@get('/examples/active_search/search')"
/>
```

--------------------------------

### JavaScript Redirect with ServerSentEventGenerator

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a redirect in Node.js using the ServerSentEventGenerator. It streams SSE to patch the UI and then uses setTimeout to execute a script that redirects the client.

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

### C# Redirect with DataStar Service

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Implements a redirect by patching HTML elements, delaying, and then executing a JavaScript redirect script using the IDatastarService in C#.

```csharp
using StarFederation.Datastar.DependencyInjection;

app.MapGet("/redirect", async (IDatastarService datastarService) =>
{
    await datastarService.PatchElementsAsync("<div id=\"indicator\">Redirecting in 3 seconds...</div>");
    await Task.Delay(TimeSpan.FromSeconds(3));
    await datastarService.ExecuteScriptAsync("setTimeout(() => window.location = \"/guide\");");
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

### Ruby DataStar ExecuteScript and PatchElements

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Ruby snippet demonstrates using DataStar to patch elements and execute a script for redirection within a streaming context, including a sleep interval.

```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'
  sleep 3
  sse.execute_script 'window.location = "/guide"'
end
```

--------------------------------

### Java: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Java example demonstrates patching signals using Server-Sent Events (SSE) with the DataStar Java SDK. It involves creating a `ServerSentEventGenerator` and using its `send` method with `PatchSignals.builder()` to construct the event data.

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

### Reading Nested Signals in Go

Source: https://data-star.dev/guide/backend_requests

Provides a Go example for reading nested signals from an HTTP request using the `datastar-go` library.

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

### HTML with data-on-signal-patch Example

Source: https://data-star.dev/examples/on_signal_patch

This HTML snippet demonstrates the usage of the `data-on-signal-patch` plugin. It includes buttons to update a message and a counter, clear changes, and display current values. It also shows how to capture and display signal patches, both all changes and counter-specific changes, using `data-on-signal-patch-filter`.

```html
<div data-signals="{counter: 0, message: 'Hello World', allChanges: [], counterChanges: []}">
    <div class="actions">
        <button data-on-click="$message = `Updated: ${performance.now().toFixed(2)}`">
            Update Message
        </button>
        <button data-on-click="$counter++">
            Increment Counter
        </button>
        <button
            class="error"
            data-on-click="$allChanges.length = 0; $counterChanges.length = 0"
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

### C# Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

A C# example using `StarFederation.Datastar.DependencyInjection` to create an endpoint that patches elements with the current time. It utilizes `IDatastarService` for patching.

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

### Asynchronous External Scripts with Custom Events (JavaScript)

Source: https://data-star.dev/docs

This example shows how to handle asynchronous operations in external scripts by dispatching custom events. Datastar does not await asynchronous calls within expressions. The `myfunction` is asynchronous and dispatches a 'mycustomevent' with the result in `evt.detail.value`.

```html
<div data-signals-result>
    <input data-bind-foo 
           data-on-input="myfunction(el, $foo)"
           data-on-mycustomevent__window="$result = evt.detail.value"
    >
    <span data-text="$result"></span>
</div>
```

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

### Backend Redirect with Ruby Server-Sent Events

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This Ruby snippet illustrates performing a backend redirect using DataStar. It utilizes a stream to first patch an indicator element and then issue a redirect command after a 3-second delay.

```ruby
datastar = Datastar.new(request:, response:)

datastar.stream do |sse|
  sse.patch_elements '<div id="indicator">Redirecting in 3 seconds...</div>'

  sleep 3

  sse.redirect '/guide'
end
```

--------------------------------

### Ruby Redirect with Datastar Stream

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

Demonstrates redirect functionality in Ruby using the Datastar library. It streams Server-Sent Events (SSE) to patch elements, pauses, and then executes a JavaScript redirect.

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

### Backend Redirect with PHP Server-Sent Events

Source: https://data-star.dev/how_tos/redirect_the_page_from_the_backend

This PHP snippet demonstrates how to use ServerSentEventGenerator to create a redirect. It first patches an indicator element into the DOM and then issues a location redirect after a delay.

```php
$sse = new ServerSentEventGenerator();
$sse->patchElements(`
    <div id="indicator">Redirecting in 3 seconds...</div>
`);
sleep(3);
$sse->location('/guide');
```

--------------------------------

### SSE Event for Patching Signals

Source: https://data-star.dev/guide/reactive_signals

This example shows a Server-Sent Event (SSE) with the 'datastar-patch-signals' event type. This allows for more complex, real-time updates to frontend signals, enabling scenarios like sequential updates or delayed signal changes.

```sse
1event: datastar-patch-signals
2data: signals {hal: 'Affirmative, Dave. I read you.'}
```

--------------------------------

### HTML Lazy Tabs Structure with Datastar

Source: https://data-star.dev/examples/lazy_tabs

This HTML snippet defines the structure for a set of lazy tabs using Datastar. It includes buttons for each tab with 'data-on-click' attributes that trigger Datastar GET requests to fetch tab content. The 'aria-selected' attribute indicates the active tab. Dependencies include the Datastar framework. Inputs are tab button clicks, and outputs are fetched HTML content for the selected tab panel.

```html
<div id="demo">
    <div role="tablist">
        <button
            role="tab"
            aria-selected="true"
            data-on-click="@get('/examples/lazy_tabs/0')"
        >
            Tab 0
        </button>
        <button
            role="tab"
            aria-selected="false"
            data-on-click="@get('/examples/lazy_tabs/1')"
        >
            Tab 1
        </button>
        <button
            role="tab"
            aria-selected="false"
            data-on-click="@get('/examples/lazy_tabs/2')"
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

### Datastar Runtime Error Example

Source: https://data-star.dev/reference/attributes

Provides an example of a Datastar runtime error log. It shows the format of error messages, including a 'More info' link for detailed explanations and context-aware error pages.

```json
Uncaught datastar runtime error: textKeyNotAllowed
More info: https://data-star.dev/errors/runtime/text_key_not_allowed?metadata=%7B%22plugin%22%3A%7B%22name%22%3A%22text%22%2C%22type%22%3A%22attribute%22%7D%2C%22element%22%3A%7B%22id%22%3A%22%22%2C%22tag%22%3A%22DIV%22%7D%2C%22expression%22%3A%7B%22rawKey%22%3A%22textFoo%22%2C%22key%22%3A%22foo%22%2C%22value%22%3A%22%22%2C%22fnContent%22%3A%22%22%7D%7D
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

### Python (Sanic) Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

A Python example using the Sanic web framework and DataStar library (`datastar_py`) to create an endpoint that patches elements with the current time via SSE.

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

### Send GET Request with @get() Action (HTML)

Source: https://data-star.dev/guide/getting_started

This snippet demonstrates how to use the `@get()` Datastar action to send a GET request to a specified URL. The response, if it's HTML, will be used to patch elements in the DOM based on their IDs. Ensure the target element (e.g., 'hal') exists in the DOM for morphing to work.

```html
<button data-on-click="@get('/endpoint')">
    Open the pod bay doors, HAL.
</button>
<div id="hal"></div>
```

--------------------------------

### Copy Text to Clipboard with DataStar Pro

Source: https://data-star.dev/docs

This snippet illustrates the use of the `@clipboard()` DataStar Pro action to copy text to the user's clipboard. It shows examples for copying plain text and Base64 encoded text, which is decoded before being copied. This is useful for handling special characters or complex data safely.

```html
1<!-- Copy plain text -->
2<button data-on-click="@clipboard('Hello, world!')"></button>
3
4<!-- Copy base64 encoded text (will decode before copying) -->
5<button data-on-click="@clipboard('SGVsbG8sIHdvcmxkIQ==', true)"></button>
```

--------------------------------

### HTML Structure for Dynamic List Loading

Source: https://data-star.dev/how_tos/load_more_list_items

This HTML snippet sets up a list container with the ID 'list' and a 'load-more' button. The button uses 'data-signals-offset' to track the current offset and 'data-on-click' to trigger a GET request to the backend when clicked.

```html
1<div id="list">
2<div>Item 1</div>
3</div>
4<button id="load-more" 
5        data-signals-offset="1" 
6        data-on-click="@get('/how_tos/load_more/data')">
7Click to load another item
8</button>
```

--------------------------------

### HTML Progress Bar with SSE Updates

Source: https://data-star.dev/examples/progress_bar

This HTML structure defines a progress bar using SVG and includes a button that appears upon completion. It utilizes Datastar's data attributes for server-sent events (SSE) to dynamically update the progress and handle restart actions. The `data-on-load` attribute initiates the SSE connection, and `data-on-click` handles the restart functionality.

```html
<div
    id="progress-bar"
    data-on-load="@get('/examples/progress_bar/updates', {openWhenHidden: true})"
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
        data-indicator-_fetching
        data-attr-aria-disabled="`${$_fetching}`"
        data-on-click="
            !$_fetching && @get('/examples/progress_bar/updates', {openWhenHidden: true})
        "
    >
        <i class="material-symbols:check-circle"></i>
        Completed! Try again?
    </button>
</div>
```

--------------------------------

### HTML: Displaying a Read-Only Table Row

Source: https://data-star.dev/examples/edit_row

This code snippet displays a single row in a table with contact information (Name, Email) and an 'Edit' button. The 'Edit' button is configured to trigger a GET request to a specific URL when clicked, initiating the editing process.

```html
<tr>
    <td>Joe Smith</td>
    <td>joe@smith.org</td>
    <td>
        <button data-on-click="@get('/examples/edit_row/0')">
            Edit
        </button>
    </td>
</tr>
```

--------------------------------

### Stream SSE Events in Go

Source: https://data-star.dev/docs

This Go code snippet demonstrates how to stream Server-Sent Events (SSE) using the DataStar Go SDK. It shows the creation of a `ServerSentEventGenerator` and how to use its methods to patch elements and signals.

```go
import ("github.com/starfederation/datastar-go/datastar")

// Creates a new `ServerSentEventGenerator` instance.
sse := datastar.NewSSE(w,r)

// Patches elements into the DOM.
sse.PatchElements(
    `<div id="question">What do you put in a toaster?</div>`
)

// Patches signals.
sse.PatchSignals([]byte(`{response: '', answer: 'bread'}`))
```

--------------------------------

### GET Request API

Source: https://data-star.dev/reference/actions

Sends a GET request to the backend. By default, it sends signals as query parameters and includes a 'Datastar-Request: true' header. The SSE connection can be managed when the page is hidden.

```APIDOC
## GET /endpoint

### Description
Sends a `GET` request to the specified URI. This action is suitable for retrieving data and can handle Server-Sent Events (SSE).

### Method
GET

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint to send the GET request to.
- **options** (object) - Optional - Configuration for the request.
  - **contentType** (string) - Optional - Sets the content type for the request. Can be 'json' (default) or 'form'.
  - **filterSignals** (object) - Optional - Filters which signals are sent with the request. Contains `include` (RegExp) and `exclude` (RegExp) properties.
  - **selector** (string) - Optional - Specifies a form element to use when `contentType` is 'form'. Defaults to the closest form.
  - **headers** (object) - Optional - An object of headers to include with the request.
  - **openWhenHidden** (boolean) - Optional - If true, keeps the SSE connection open when the page is hidden. Defaults to `false`.
  - **retryInterval** (number) - Optional - The interval in milliseconds for retrying requests. Defaults to `1000`.
  - **retryScaler** (number) - Optional - The multiplier for scaling retry wait times. Defaults to `2`.
  - **retryMaxWaitMs** (number) - Optional - The maximum wait time in milliseconds between retries. Defaults to `30000`.
  - **retryMaxCount** (number) - Optional - The maximum number of retry attempts. Defaults to `10`.
  - **requestCancellation** (string | AbortController) - Optional - Controls request cancellation behavior. Defaults to `'auto'`.

### Request Example
```html
<button data-on-click="@get('/endpoint', { openWhenHidden: true })"></button>
<button data-on-click="@get('/endpoint', { contentType: 'form' })"></button>
<form enctype="multipart/form-data">
    <input type="file" name="file" />
    <button data-on-click="@get('/endpoint', { contentType: 'form' })"></button>
</form>
<button data-on-click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})></button>
```

### Response
#### Success Response (200)
- **Datastar SSE events** - Zero or more Datastar Server-Sent Events.

#### Response Example
```json
{
  "event": "update",
  "data": {
    "message": "Hello, world!"
  }
}
```
```

--------------------------------

### Linearly Interpolate Values with DataStar Pro's @fit()

Source: https://data-star.dev/docs

The `@fit()` DataStar Pro function linearly interpolates a value from a source range (oldMin, oldMax) to a target range (newMin, newMax). The optional `shouldClamp` parameter restricts the output to the target range, and `shouldRound` rounds the result to the nearest integer. Examples include converting slider values to RGB, Celsius to Fahrenheit, and mapping mouse position to opacity.

```html
 1<!-- Convert a 0-100 slider to 0-255 RGB value -->
 2<div>
 3    <input type="range" min="0" max="100" value="50" data-bind-slider-value>
 4    <div data-computed-rgb-value="@fit($sliderValue, 0, 100, 0, 255)">
 5        RGB Value: <span data-text="$rgbValue"></span>
 6    </div>
 7</div>
 8
 9<!-- Convert Celsius to Fahrenheit -->
10<div>
11    <input type="number" data-bind-celsius value="20" />
12    <div data-computed-fahrenheit="@fit($celsius, 0, 100, 32, 212)">
13        <span data-text="$celsius"></span>°C = <span data-text="$fahrenheit.toFixed(1)"></span>°F
14    </div>
15</div>
16
17<!-- Map mouse position to element opacity (clamped) -->
18<div
19    data-signals-mouse-x="0"
20    data-computed-opacity="@fit($mouseX, 0, window.innerWidth, 0, 1, true)"
21    data-on-mousemove__window="$mouseX = evt.clientX"
22    data-attr-style="'opacity: ' + $opacity"
23>
24    Move your mouse horizontally to change opacity
25</div>
```

--------------------------------

### Kotlin SDK Example for Dynamic List Loading

Source: https://data-star.dev/how_tos/load_more_list_items

This Kotlin code snippet shows how to read offset signals and initiate server-sent events for dynamic list loading using DataStar. It defines a data class for signals and uses ServerSentEventGenerator to patch elements, update signals, or remove elements based on the offset.

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

### Send GET Request with Datastar

Source: https://data-star.dev/reference/actions

The `@get()` action sends a GET request to a specified URI. It supports options for customizing request headers, filtering signals, controlling connection behavior when the page is hidden, and setting content types like 'form' or 'multipart/form-data'. Signals are sent as query parameters by default.

```html
<button data-on-click="@get('/endpoint')"></button>
<button data-on-click="@get('/endpoint', {openWhenHidden: true})"></button>
<button data-on-click="@get('/endpoint', {contentType: 'form'})></button>
<form enctype="multipart/form-data">
    <input type="file" name="file" />
    <button data-on-click="@get('/endpoint', {contentType: 'form'})></button>
</form>
<button data-on-click="@get('/endpoint', {
    filterSignals: {include: /^foo\./},
    headers: {
        'X-Csrf-Token': 'JImikTbsoCYQ9oGOcvugov0Awc5LbqFsZW6ObRCxuq',
    },
    openWhenHidden: true,
    requestCancellation: 'disabled',
})"></button>
```

--------------------------------

### Rust Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

A Rust example using `datastar` crate to generate SSEs for patching elements with the current time. It leverages `chrono` for time formatting and `async_stream` for stream handling.

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

### Clojure Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Provides a Clojure example using `starfederation.datastar.clojure.api` to patch elements with the current time at a set interval. It sets up a Server-Sent Events (SSE) response.

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
                  (LocalDateTime/.format (LocalDateTime/now) formatter)])))}))

        (d*/close-sse! sse)))
```

--------------------------------

### Patch Elements using Datastar C# SDK

Source: https://data-star.dev/docs

This C# code demonstrates how to add Datastar as a service using dependency injection and then use the IDatastarService to patch elements into the DOM. It shows updating content with a delay.

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

### Practical Use Case: Toggling Nested Menu Signals in HTML

Source: https://data-star.dev/guide/backend_requests

A practical example of using nested signals to manage the open/closed state of menus and a `toggleAll` action to control them, with HTML structure provided.

```html
<div data-signals="{menu: {isOpen: {desktop: false, mobile: false}}}">
    <button data-on-click="@toggleAll({include: /^menu\.isOpen\./})">
        Open/close menu
    </button>
</div>
```

--------------------------------

### Using External Scripts with Arguments and Return Values (JavaScript)

Source: https://data-star.dev/docs

This snippet demonstrates how to pass data into a function via arguments and return a result. It uses `data-*` attributes for binding and event handling, enabling reactivity. The function `myfunction` takes an argument and returns a formatted string.

```html
<div data-signals-result>
    <input data-bind-foo 
        data-on-input="$result = myfunction($foo)"
    >
    <span data-text="$result"></span>
</div>
```

```javascript
function myfunction(data) {
    return `You entered: ${data}`;
}
```

--------------------------------

### Kotlin Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

A Kotlin example demonstrating how to use `ServerSentEventGenerator` to patch elements with the current time at specified intervals. This snippet focuses on the backend generation of the SSE.

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

### Set JavaScript Response Headers with DataStar

Source: https://data-star.dev/docs

When serving JavaScript (`text/javascript`), you can set custom response headers like `datastar-script-attributes` to include JSON-encoded attributes for the script element. This allows for dynamic configuration of scripts served from the server. The example demonstrates setting the Content-Type and script attributes, along with a simple JavaScript response body.

```javascript
1response.headers.set('Content-Type', 'text/javascript')
2response.headers.set('datastar-script-attributes', JSON.stringify({ type: 'module' }))
3response.body = 'console.log("Hello from server!");'
```

--------------------------------

### Stream SSE Events in Ruby

Source: https://data-star.dev/docs

This Ruby snippet shows how to stream Server-Sent Events (SSE) using the `datastar` gem. It demonstrates instantiating `Datastar` and using the `stream` method with a block to send patch events for elements and signals.

```ruby
require 'datastar'

# Create a Datastar::Dispatcher instance

datastar = Datastar.new(request:, response:)

# In a Rack handler, you can instantiate from the Rack env
# datastar = Datastar.from_rack_env(env)

# Start a streaming response
datastar.stream do |sse|
  # Patches elements into the DOM
  sse.patch_elements %(<div id="question">What do you put in a toaster?</div>)

  # Patches signals
  sse.patch_signals(response: '', answer: 'bread')
end
```

--------------------------------

### Attach Event Listeners with data-on Attribute

Source: https://data-star.dev/guide/reactive_signals

The `data-on` attribute attaches event listeners to elements, executing specified expressions when the event occurs. It supports standard and custom event names, enabling interactive UI elements. For example, `data-on-click` can reset a signal's value.

```html
<input data-bind-foo />
<button data-on-click="$foo = ''">
    Reset
</button>
```

--------------------------------

### Stream SSE Events in Clojure

Source: https://data-star.dev/docs

This snippet demonstrates how to set up a backend controller action in Clojure to stream Server-Sent Events (SSE) using the DataStar SDK. It shows how to create an SSE response and send patches for DOM elements and signals.

```clojure
(require '[starfederation.datastar.clojure.api :as d*])
(require '[starfederation.datastar.clojure.adapter.http-kit :refer [->sse-response on-open]])

;; in a ring handler
(defn handler [request]
  ;; Create an SSE response
  (->sse-response request
                  {on-open
                   (fn [sse]
                     ;; Patches elements into the DOM
                     (d*/patch-elements! sse
                                         "<div id=\"question\">What do you put in a toaster?</div>")

                     ;; Patches signals
                     (d*/patch-signals! sse "{response: '', answer: 'bread'}"))})
```

--------------------------------

### data-indicator Attribute Usage Examples

Source: https://data-star.dev/errors/runtime/indicator_key_or_value_required

Examples demonstrating the correct usage of the data-indicator attribute. The attribute requires either a key or a value to represent a signal's name. This signal is set to true when an SSE request is in progress and false otherwise.

```html
<div data-indicator-foo></div>
```

```html
<div data-indicator="foo"></div>
```

--------------------------------

### POST Request API

Source: https://data-star.dev/reference/actions

Sends a POST request to the backend. Similar to GET, but used for submitting data.

```APIDOC
## POST /endpoint

### Description
Sends a `POST` request to the specified URI. This action is used for creating or submitting data to the server.

### Method
POST

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint to send the POST request to.
- **options** (object) - Optional - Configuration for the request. See GET request options for details.

### Request Example
```html
<button data-on-click="@post('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response Body** - The server's response to the POST request.

#### Response Example
```json
{
  "status": "success",
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```
```

--------------------------------

### Patch Elements using Datastar Go SDK

Source: https://data-star.dev/docs

This Go code snippet illustrates how to create a ServerSentEventGenerator and use its PatchElements method to update the DOM. It includes sending an initial HTML fragment and then a subsequent update after a short delay.

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

### Generating SSE ExecuteScript Event with Go SDK (Go)

Source: https://data-star.dev/docs

This Go code snippet demonstrates using the DataStar SDK to generate a Server-Sent Event (SSE) that executes a JavaScript script on the frontend. The `ExecuteScript` helper function simplifies the process.

```go
sse := datastar.NewSSE(writer, request)
sse.ExecuteScript(`alert('This mission is too important for me to allow you to jeopardize it.')`)
```

--------------------------------

### HTML Structure for Web Component Interaction

Source: https://data-star.dev/examples/web_component

This HTML snippet sets up the user interface for the web component. It includes a label with an input field bound to the '_name' signal, a span to display the reversed text bound to the '_reversed' signal, and the custom 'reverse-component' itself, configured with event listeners and attribute bindings.

```html
<label>
    Reversed
    <input type="text" value="Your Name" data-bind-_name/>
</label>
<span data-signals-_reversed data-text="$_reversed"></span>
<reverse-component
    data-on-reverse="$_reversed = evt.detail.value"
    data-attr-name="$_name">
</reverse-component>
```

--------------------------------

### Patch Elements using Datastar Clojure SDK

Source: https://data-star.dev/docs

This Clojure code snippet shows how to use the Datastar SDK to create an SSE response and patch elements into the DOM. It includes sending an initial message and then updating it after a delay.

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
17                                         "<div id=\"hal\">Waiting for an order...</div>"))})
```

--------------------------------

### PHP Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Shows a PHP implementation using `ServerSentEventGenerator` to update an element with the current time. This example highlights the server-side generation of SSE for dynamic content.

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

### HTML Structure for DataStar Signals

Source: https://data-star.dev/guide/reactive_signals

This HTML snippet demonstrates how to set up elements to receive and display data-driven signals from the backend using DataStar attributes like 'data-signals-hal' and 'data-on-click'. It shows a button that triggers an API call and a div to display the 'hal' signal.

```html
1<div data-signals-hal="'...'" >
2    <button data-on-click="@get('/endpoint')">
3        HAL, do you read me?
4    </button>
5    <div data-text="$hal"></div>
6</div>
```

--------------------------------

### Example HTTP Response with Error Code

Source: https://data-star.dev/essays/im_a_teapot

This snippet demonstrates an HTTP response that includes an H1 tag displaying an error message. It's presented as an example within a discussion about HTTP status codes, highlighting a scenario where an error code might be embedded in an HTML response.

```http
HTTP/1.1 200 OK
Content-Type: text/html
<H1>Error 404</H1>
```

--------------------------------

### JavaScript Function: Synchronous Example

Source: https://data-star.dev/guide/datastar_expressions

Provides a simple synchronous JavaScript function `myfunction` that takes data as an argument and returns a formatted string. This function can be called from Datastar expressions.

```javascript
1function myfunction(data) {
2    return `You entered: ${data}`;
3}
```

--------------------------------

### Patch Elements using Datastar Service in C#

Source: https://data-star.dev/guide/getting_started

This C# example shows how to integrate Datastar into an ASP.NET Core application by adding it as a service. It then uses `IDatastarService.PatchElementsAsync` to send HTML content to patch elements in the DOM, including a delay between updates.

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

### Datastar: GET Request to Patch Elements with HTML Response

Source: https://data-star.dev/guide

Demonstrates how to use the '@get()' action in Datastar to send a GET request to a backend endpoint. If the response content type is 'text/html', the response's HTML elements are morphed into the DOM, replacing existing elements with matching IDs. This requires a corresponding element in the DOM with the specified ID for morphing to occur.

```html
<button data-on-click="@get('/endpoint')">
    Open the pod bay doors, HAL.
</button>
<div id="hal"></div>
```

```html
<div id="hal">
    I’m sorry, Dave. I’m afraid I can’t do that.
</div>
```

--------------------------------

### Submit Form Data via GET/POST with ContentType 'form'

Source: https://data-star.dev/examples/form_data

This snippet demonstrates how to submit form data using both GET and POST requests with the contentType option set to 'form'. This option automatically serializes form elements and sends them to the backend. A selector can be used to target a specific form.

```html
<form id="myform">
    foo:<input type="checkbox" name="checkboxes" value="foo" />
    bar:<input type="checkbox" name="checkboxes" value="bar" />
    baz:<input type="checkbox" name="checkboxes" value="baz" />
    <button data-on-click="@get('/endpoint', {contentType: 'form'})">
        Submit GET request
    </button>
    <button data-on-click="@post('/endpoint', {contentType: 'form'})">
        Submit POST request
    </button>
</form>

<button data-on-click="@get('/endpoint', {contentType: 'form', selector: '#myform'})">
    Submit GET request from outside the form
</button>
```

--------------------------------

### Go: Patching Signals with SSE

Source: https://data-star.dev/docs

This Go code snippet shows how to patch signals using Server-Sent Events (SSE) with the DataStar Go SDK. It initializes an `ServerSentEventGenerator` and uses its `PatchSignals` method to send signal updates, including a time delay between them.

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
14
```

--------------------------------

### Alpine.js data-on-load Basic Usage (JavaScript)

Source: https://data-star.dev/docs

This snippet demonstrates the 'data-on-load' directive, which executes an expression when the element is loaded into the DOM. Here, it initializes a count variable to 1.

```html
<div data-on-load="$count = 1"></div>

```

--------------------------------

### Implement Server-Sent Events with Rust

Source: https://data-star.dev/docs

This snippet demonstrates how to generate Server-Sent Events (SSE) using Rust and the `async_stream` crate. It involves yielding `PatchSignals` within a stream, with a delay between events.

```rust
use async_stream::stream;
use datastar::prelude::*;
use std::thread;
use std::time::Duration;

Sse(stream! {
    // Patches signals.
    yield PatchSignals::new("{hal: 'Affirmative, Dave. I read you.'}").into();

    thread::sleep(Duration::from_secs(1));
    
    yield PatchSignals::new("{hal: '...'}").into();
})
```

--------------------------------

### Read Datastar Signals in Python (FastAPI)

Source: https://data-star.dev/docs

Illustrates reading signals within a FastAPI endpoint in Python. It uses the `@datastar_response` decorator and the `read_signals` utility function to retrieve the frontend state as a dictionary.

```python
from datastar_py.fastapi import datastar_response, read_signals

@app.get("/updates")
@datastar_response
async def updates(request: Request):
    # Retrieve a dictionary with the current state of the signals from the frontend
    signals = await read_signals(request)

```

--------------------------------

### Datastar Expressions: Multiple Statements with Semicolon

Source: https://data-star.dev/docs

Shows how to execute multiple statements within a single Datastar expression by separating them with a semicolon. This allows for sequential operations.

```html
<div data-signals-foo="1">
    <button data-on-click="$landingGearRetracted = true; @post('/launch')">
        Force launch
    </button>
</div>
```

--------------------------------

### Send SSE Events with Datastar in C#

Source: https://data-star.dev/guide/backend_requests

This C# example shows how to configure Datastar as a service and handle SSE events within a MapGet endpoint. It demonstrates patching elements and signals asynchronously. Requires the Datastar .NET SDK.

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

### Patch Elements using Datastar Java SDK

Source: https://data-star.dev/docs

This Java code demonstrates using the ServerSentEventGenerator with an HttpServletResponseAdapter to patch elements into the DOM. It shows sending an initial HTML update and then a follow-up update after a one-second delay.

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

### Datastar Expressions: Multi-line Statements

Source: https://data-star.dev/docs

Demonstrates how Datastar expressions can span multiple lines. A semicolon is required to separate statements, as line breaks alone are not sufficient.

```html
<div data-signals-foo="1">
    <button data-on-click="
        $landingGearRetracted = true; 
        @post('/launch')
    ">
        Force launch
    </button>
</div>
```

--------------------------------

### Set Multiple HTML Attributes with Datastar data-attr

Source: https://data-star.dev/docs

Demonstrates using the `data-attr` attribute in Datastar to set multiple HTML attributes on an element simultaneously using a key-value pair object.

```html
<div data-attr="{title: $foo, disabled: $bar}"></div>
```

--------------------------------

### Datastar Expression: Accessing Signal Properties

Source: https://data-star.dev/guide/datastar_expressions

Illustrates how to access properties of a signal within a Datastar expression. This example displays the length of the string signal `$foo`.

```html
1<div data-text="$foo.length"></div>
```

--------------------------------

### Read Datastar Signals in C#

Source: https://data-star.dev/docs

Shows how to set up Datastar as a service and read nested signals from an incoming request using `IDatastarService` in a C# .NET environment. It defines a `Signals` record to represent the expected JSON structure.

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

### Send DELETE Request with Datastar

Source: https://data-star.dev/reference/actions

The `@delete()` action facilitates sending DELETE requests to a specified URI. It offers the same configuration options as the `@get()` action.

```html
<button data-on-click="@delete('/endpoint')"></button>
```

--------------------------------

### Send PATCH Request with Datastar

Source: https://data-star.dev/reference/actions

The `@patch()` action is designed for sending PATCH requests to a given URI. It shares the same set of options and behavior as the `@get()` action.

```html
<button data-on-click="@patch('/endpoint')"></button>
```

--------------------------------

### Interval Execution with `data-on-interval`

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Demonstrates how to use the `data-on-interval` attribute to run an expression, such as a GET request, at a specified interval. The `__duration` modifier sets the interval time.

```html
<div id="time"
     data-on-interval__duration.5s="@get('/endpoint')"
></div>
```

--------------------------------

### Send SSE Events with Datastar in Ruby

Source: https://data-star.dev/guide/backend_requests

This Ruby example demonstrates initiating a streaming SSE response using the Datastar gem. It shows how to patch elements and signals within a block passed to the `stream` method.

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

### Send PUT Request with Datastar

Source: https://data-star.dev/reference/actions

The `@put()` action allows for sending PUT requests to a specified URI. It mirrors the functionality and options available with the `@get()` action.

```html
<button data-on-click="@put('/endpoint')"></button>
```

--------------------------------

### Value Interpolation with @fit()

Source: https://data-star.dev/reference/actions

Demonstrates the `@fit()` directive for linear interpolation between two value ranges. Examples include converting a slider value to an RGB component, Celsius to Fahrenheit, and mapping mouse position to element opacity, with options for clamping and rounding.

```html
<!-- Convert a 0-100 slider to 0-255 RGB value -->
<div>
    <input type="range" min="0" max="100" value="50" data-bind-slider-value />
    <div data-computed-rgb-value="@fit($sliderValue, 0, 100, 0, 255)">
        RGB Value: <span data-text="$rgbValue"></span>
    </div>
</div>

<!-- Convert Celsius to Fahrenheit -->
<div>
    <input type="number" data-bind-celsius value="20" />
    <div data-computed-fahrenheit="@fit($celsius, 0, 100, 32, 212)">
        <span data-text="$celsius"></span>°C = <span data-text="$fahrenheit.toFixed(1)"></span>°F
    </div>
</div>

<!-- Map mouse position to element opacity (clamped) -->
<div
    data-signals-mouse-x="0"
    data-computed-opacity="@fit($mouseX, 0, window.innerWidth, 0, 1, true)"
    data-on-mousemove__window="$mouseX = evt.clientX"
    data-attr-style="'opacity: ' + $opacity"
>
    Move your mouse horizontally to change opacity
</div>
```

--------------------------------

### Clojure: Patching Signals with SSE

Source: https://data-star.dev/docs

This Clojure code snippet illustrates how to generate Server-Sent Events (SSE) to patch signals using the DataStar SDK. It sets up a Ring handler that creates an SSE response and uses `d*/patch-signals!` to send signal updates with a delay between them.

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

### Stream SSE Events in Kotlin

Source: https://data-star.dev/docs

This Kotlin snippet demonstrates how to stream Server-Sent Events (SSE) using the DataStar library. It shows the usage of `ServerSentEventGenerator` to patch elements and signals directly within a response context.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = "<div id=\"question\">What do you put in a toaster?</div>",
)

generator.patchSignals(
    signals = "{\"response\": \"\", \"answer\": \"bread\"}",
)
```

--------------------------------

### Global Keydown Listener with data-on

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

Listens for any keydown event globally on the window and executes a simple JavaScript expression. This serves as a basic example to show event capture.

```html
<div data-on-keydown__window="alert('Key pressed')"></div>
```

--------------------------------

### Datastar Expression: Access Signal Property Length

Source: https://data-star.dev/docs

Illustrates accessing a property of a signal and using JavaScript's `.length` method within a Datastar expression to display it.

```html
<div data-text="$foo.length"></div>
```

--------------------------------

### Color Throb Animation using HTML

Source: https://data-star.dev/examples/animations

This example demonstrates a simple color throb animation by maintaining a stable element ID across content swaps. Datastar swaps elements while preserving their IDs, enabling CSS transitions between old and new element states. The provided HTML sets initial colors and text for the element.

```html
<div
    id="color-throb"
    style="color: var(--blue-8); background-color: var(--orange-5);"
>
    blue on orange
</div>
```

--------------------------------

### HTML Table with Bulk Update Functionality

Source: https://data-star.dev/examples/bulk_update

This HTML snippet defines a table with checkboxes for row selection and buttons for bulk activation and deactivation. It uses custom data attributes for handling user interactions and binding data. The checkboxes in the header and rows manage selection states, and the buttons trigger PUT requests to specified endpoints. The table rows are dynamically updated by the server after the bulk operation.

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
                        data-bind-_all
                        data-on-change="$selections = Array(4).fill($_all)"
                        data-effect="$selections; $_all = $selections.every(Boolean)"
                        data-attr-disabled="$_fetching"
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
                        data-bind-selections
                        data-attr-disabled="$_fetching"
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
            data-on-click="@put('/examples/bulk_update/activate')"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
        >
            <i class="pixelarticons:user-plus"></i>
            Activate
        </button>
        <button
            class="error"
            data-on-click="@put('/examples/bulk_update/deactivate')"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
        >
            <i class="pixelarticons:user-x"></i>
            Deactivate
        </button>
    </div>
</div>
```

--------------------------------

### Attribute Order Example in HTML

Source: https://data-star.dev/reference/attributes

Demonstrates the order of attribute processing in Datastar. Attributes are evaluated in DOM depth-first order, and within an element, in the order they appear. This is crucial for dependencies, like ensuring an indicator is set up before a fetch request.

```html
<div data-indicator-fetching data-on-load="@get('/endpoint')"></div>
```

--------------------------------

### Send POST Request with Datastar

Source: https://data-star.dev/reference/actions

The `@post()` action is used to send a POST request to a specified URI. It functions similarly to `@get()` in terms of available options for request customization.

```html
<button data-on-click="@post('/endpoint')"></button>
```

--------------------------------

### Datastar Patch Elements Example

Source: https://data-star.dev/reference/sse_events

This code demonstrates how to use the `datastar-patch-elements` SSE event to update the DOM. It shows basic patching, removing elements, and using non-default morphing modes like 'inner' and specifying selectors. It also illustrates handling multi-line element data and using view transitions.

```text
event: datastar-patch-elements
data: elements <div id="foo">Hello world!</div>
```

```text
event: datastar-patch-elements
data: mode remove
data: selector #foo
```

```text
event: datastar-patch-elements
data: mode remove
data: elements <div id="foo"></div>
```

```text
event: datastar-patch-elements
data: mode inner
data: selector #foo
data: useViewTransition true
data: elements <div>
data:        Hello world!
data: </div>
```

--------------------------------

### Interactive Quiz with data-computed and data-show

Source: https://data-star.dev/guide/reactive_signals

This snippet creates an interactive quiz element. It uses `data-signals` to initialize response and answer signals, `data-computed` to determine correctness, `data-on-click` for user input via prompt, and `data-show` to conditionally display feedback based on the computed result.

```html
<div
    data-signals="{response: '', answer: 'bread'}"
    data-computed-correct="$response.toLowerCase() == $answer"
>
    <div id="question">What do you put in a toaster?</div>
    <button data-on-click="$response = prompt('Answer:') ?? ''">BUZZ</button>
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

### Datastar Expressions: Ternary and Logical Operators

Source: https://data-star.dev/docs

Demonstrates the use of JavaScript operators like the ternary operator (`?:`), logical OR (`||`), and logical AND (`&&`) within Datastar expressions for conditional rendering and actions.

```html
// Output one of two values, depending on the truthiness of a signal
<div data-text="$landingGearRetracted ? 'Ready' : 'Waiting'"></div>

// Show a countdown if the signal is truthy or the time remaining is less than 10 seconds
<div data-show="$landingGearRetracted || $timeRemaining < 10">
    Countdown
</div>

// Only send a request if the signal is truthy
<button data-on-click="$landingGearRetracted && @post('/launch')">
    Launch
</button>
```

--------------------------------

### Read Datastar Signals in Kotlin

Source: https://data-star.dev/docs

Demonstrates reading signals in Kotlin using `kotlinx.serialization`. It defines a `Signals` data class and a `JsonUnmarshaller` to deserialize the JSON request body, then uses a `readSignals` helper function.

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

### Web Components with Attributes and Custom Events (JavaScript)

Source: https://data-star.dev/docs

This snippet illustrates creating and using a reusable web component. Data is passed via attributes (`data-attr-src`), and custom events (`mycustomevent`) are used for communication. The `MyComponent` class extends `HTMLElement` and dispatches events when its 'src' attribute changes.

```html
<div data-signals-result="''">
    <input data-bind-foo />
    <my-component
        data-attr-src="$foo"
        data-on-mycustomevent="$result = evt.detail.value"
    ></my-component>
    <span data-text="$result"></span>
</div>
```

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

### JavaScript SSE Stream with Datastar

Source: https://data-star.dev/guide/reactive_signals

This JavaScript code snippet demonstrates creating a Server-Sent Event stream using Datastar. It initializes a `ServerSentEventGenerator` and uses `setTimeout` to send patch signals with a delay, simulating real-time updates. The input is a request and response object, and the output is a stream of patch signals.

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

### Datastar Button Click Event Handling

Source: https://data-star.dev/index

This HTML snippet demonstrates how to attach a click event listener to a button using Datastar. The `data-on-click` attribute specifies an HTTP GET request to '/endpoint' that will be triggered when the button is clicked. The result of the request is typically rendered into the `#hal` div.

```html
<button data-on-click="@get('/endpoint')">
    Open the pod bay doors, HAL.
</button>

<div id="hal">Waiting for an order...</div>
```

--------------------------------

### Kotlin: Patching Elements with Datastar SSE Generator

Source: https://data-star.dev/guide

Shows how to patch elements into the DOM in Kotlin using Datastar's ServerSentEventGenerator. This example initializes the generator and then calls the `patchElements` method to send updates to an HTML element, incorporating a one-second delay between the two update operations.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """<div id=\"hal\">I’m sorry, Dave. I’m afraid I can’t do that.</div>""",
)

Thread.sleep(ONE_SECOND)

generator.patchElements(
    elements = """<div id=\"hal\">Waiting for an order...</div>"""
)
```

--------------------------------

### Alpine.js data-on-my-event with Case Conversion (JavaScript)

Source: https://data-star.dev/docs

Illustrates the use of 'data-on-my-event' with the '__case.camel' modifier to convert event casing to camelCase. This is useful for custom events where consistent casing is important.

```html
<div data-on-my-event__case.camel="$foo = ''"></div>

```

--------------------------------

### Ruby Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Provides a Ruby example utilizing DataStar to generate Server-Sent Events (SSE) that update an element with the current time. It uses string interpolation for embedding the time.

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

### Send SSE Events with Datastar in PHP

Source: https://data-star.dev/guide/backend_requests

This PHP example utilizes the ServerSentEventGenerator to dispatch SSE events, including patching HTML elements and signals. It shows how to instantiate the generator and send data.

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

### Alpine.js data-on-intersect Basic Usage (JavaScript)

Source: https://data-star.dev/docs

Shows the basic implementation of the 'data-on-intersect' directive. When the target div becomes visible in the viewport, the expression '$intersected = true' is evaluated.

```html
<div data-on-intersect="$intersected = true"></div>

```

--------------------------------

### HTML Structure for DBMon Demo

Source: https://data-star.dev/examples/dbmon

This HTML snippet defines the user interface for the DBMon demo. It includes elements for displaying render times, mutation rates, FPS, and a table for database cluster statistics. It uses DataStar attributes for dynamic behavior and event handling.

```html
1 <div
2     id="demo"
3     data-on-load="@get('/examples/dbmon/updates')"
4     data-signals-_editing__ifmissing="false"
5 >
6     <p>
7         Average render time for entire page: { renderTime }
8     </p>
9     <div role="group">
10        <label>
11            Mutation Rate %
12            <input
13                type="number"
14                min="0"
15                max="100"
16                value="20"
17                data-on-focus="$_editing = true"
18                data-on-blur="@put('/examples/dbmon/inputs'); $_editing = false"
19                data-attr-data-bind-mutation-rate="$_editing"
20                data-attr-data-bind-_mutation-rate="!$_editing"
21            />
22        </label>
23        <label>
24            FPS
25            <input
26                type="number"
27                min="1"
28                max="144"
29                value="60"
30                data-on-focus="$_editing = true"
31                data-on-blur="@put('/examples/dbmon/inputs'); $_editing = false"
32                data-attr-data-bind-fps="$_editing"
33                data-attr-data-bind-_fps="!$_editing"
34            />
35        </label>
36    </div>
37    <table style="table-layout: fixed; width: 100%; word-break: break-all">
38        <tbody>
39            <!-- Dynamic rows generated by server -->
40            <tr>
41                <td>cluster1</td>
42                <td style="background-color: var(--_active-color)" class="success">
43                    8
44                </td>
45                <td aria-description="SELECT blah from something">
46                    12ms
47                </td>
48                <!-- More query cells... -->
49            </tr>
50            <!-- More database rows... -->
51        </tbody>
52    </table>
53</div>
```

--------------------------------

### Handle DataStar Fetch Events in HTML

Source: https://data-star.dev/docs

This snippet shows how to use the `data-on-datastar-fetch` attribute in HTML to listen for various stages of a `datastar-fetch` request lifecycle. It demonstrates logging a message to the console when a fetch request encounters an error, using event delegation within a `div` element.

```html
1<div data-on-datastar-fetch="
2    evt.detail.type === 'error' && console.log('Fetch error encountered')
3"></div>
```

--------------------------------

### Executing Inline JavaScript from Backend Response (JavaScript)

Source: https://data-star.dev/docs

This snippet shows a simple JavaScript alert executed directly from a backend response. If the response `content-type` is `text/javascript`, the browser will execute the provided script.

```javascript
alert('This mission is too important for me to allow you to jeopardize it.')
```

--------------------------------

### Datastar Patch Signals Example

Source: https://data-star.dev/reference/sse_events

This code illustrates the `datastar-patch-signals` SSE event for updating signals on the page. It covers basic signal patching, removing signals by setting values to null, and using the `onlyIfMissing` option to conditionally update signals.

```text
event: datastar-patch-signals
data: signals {foo: 1, bar: 2}
```

```text
event: datastar-patch-signals
data: signals {foo: null, bar: null}
```

```text
event: datastar-patch-signals
data: onlyIfMissing true
data: signals {foo: 1, bar: 2}
```

--------------------------------

### Send SSE Events with Datastar in Java

Source: https://data-star.dev/guide/backend_requests

This Java example uses the ServerSentEventGenerator to send SSE events, including patching elements and signals. It requires adapting the response to an AbstractResponseAdapter. Assumes HttpServletResponse is available.

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

### Go: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Go code snippet illustrates how to patch signals using Server-Sent Events (SSE) with the DataStar Go SDK. It shows creating an `ServerSentEventGenerator` and using its `PatchSignals` method to send updates.

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

### Rust SSE Stream with Datastar

Source: https://data-star.dev/guide/reactive_signals

This Rust code snippet demonstrates how to create a Server-Sent Event stream using Datastar's `async_stream` macro. It yields patch signals at intervals, simulating real-time updates. Dependencies include `async_stream`, `datastar`, and `std::thread` for sleeping.

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

### HTML Table Row Deletion with Confirmation

Source: https://data-star.dev/examples/delete_row

This snippet displays an HTML table structure with a 'Delete' button in each row. The button is configured with a `data-on-click` attribute that first prompts the user with a confirmation dialog and then executes a delete function if confirmed. It also includes attributes for visual feedback during the delete process.

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
                    data-on-click="confirm('Are you sure?') && @delete('/examples/delete_row/0')"
                    data-indicator-_fetching
                    data-attr-disabled="$_fetching"
                >
                    Delete
                </button>
            </td>
        </tr>
    </tbody>
</table>
```

--------------------------------

### HTML File Upload Form

Source: https://data-star.dev/examples/file_upload

This HTML snippet defines a file upload form. It includes a label with a file input and a submit button. The input uses 'data-bind-files' for automatic file binding, and the button's click event is conditionally enabled based on the presence of selected files. Files are automatically encoded as base64 when submitted.

```html
<label>
    <p>Pick anything less than 1MB</p>
    <input type="file" data-bind-files multiple/>
</label>
<button
    class="warning"
    data-on-click="$files.length && @post('/examples/file_upload')"
    data-attr-aria-disabled="`${!$files.length}`">
    Submit
</button>
```

--------------------------------

### Datastar Expression: Multiple Statements with Semicolon

Source: https://data-star.dev/guide/datastar_expressions

Explains how to use multiple statements within a single Datastar expression by separating them with a semicolon. This example updates a signal and then triggers a post request.

```html
1<div data-signals-foo="1">
2    <button data-on-click="$landingGearRetracted = true; @post('/launch')">
3        Force launch
4    </button>
5</div>
```

--------------------------------

### Datastar Expression: Access Element ID

Source: https://data-star.dev/docs

Shows how to use the `el` variable within a Datastar expression to access the current element's properties, specifically its `id` attribute.

```html
<div id="foo" data-text="el.id"></div>
```

--------------------------------

### JavaScript Integration: Synchronous Function Call

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates integrating external synchronous JavaScript functions into Datastar expressions. The example uses `data-on-input` to call `myfunction` with the signal value `$foo` and binds the result to `$result`.

```html
1<div data-signals-result>
2    <input data-bind-foo 
3        data-on-input="$result = myfunction($foo)"
4    >
5    <span data-text="$result"></span>
6</div>
```

--------------------------------

### Set HTML Attribute with Datastar data-attr

Source: https://data-star.dev/docs

Shows how to use the `data-attr` attribute in Datastar to set a single HTML attribute to the value of a Datastar expression, keeping it in sync with the expression's value.

```html
<div data-attr-title="$foo"></div>
```

--------------------------------

### Handle Click Events with data-on

Source: https://data-star.dev/guide/getting_started

Attach an event listener to an HTML element using the `data-on` attribute. This example shows how to trigger a JavaScript alert when a button is clicked. The attribute's value is a Datastar expression.

```html
<button data-on-click="alert('I’m sorry, Dave. I’m afraid I can’t do that.')">
    Open the pod bay doors, HAL.
</button>
```

--------------------------------

### Kotlin: Patching Signals with SSE

Source: https://data-star.dev/docs

This Kotlin code snippet demonstrates patching signals using Server-Sent Events (SSE) with the DataStar SDK. It initializes a `ServerSentEventGenerator` and uses the `patchSignals` function to send signal updates, including a pause between updates.

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
12
```

--------------------------------

### Generate SSE Events to Patch DOM Elements (Rust)

Source: https://data-star.dev/guide/getting_started

This Rust example uses the datastar crate to generate SSE events. It employs the async_stream macro to create a stream that yields PatchElements, including a one-second sleep between events.

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

### Direct JavaScript Execution via DataStar SSE

Source: https://data-star.dev/guide/datastar_expressions

Shows how to execute arbitrary JavaScript directly in the browser using Server-Sent Events (SSE) with DataStar. This example demonstrates embedding a script within a 'datastar-patch-elements' event to display an alert.

```html
event: datastar-patch-elements
data: elements <div id="hal">
data: elements     <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
data: elements </div>
```

--------------------------------

### JavaScript Function: Asynchronous Example with Custom Event Dispatch

Source: https://data-star.dev/guide/datastar_expressions

Presents an asynchronous JavaScript function `myfunction` that simulates an async operation using `setTimeout` and `Promise`. Upon completion, it dispatches a `CustomEvent` named `mycustomevent` to pass the result back to the Datastar expression.

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

### Python (Sanic): Patching Signals with SSE

Source: https://data-star.dev/docs

This Python snippet demonstrates patching signals using Server-Sent Events (SSE) within the Sanic web framework with the DataStar SDK. It uses `datastar_response` decorator and yields SSE objects to send signal updates asynchronously with a delay.

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
10
```

--------------------------------

### Generate SSE Events to Patch DOM Elements (Python)

Source: https://data-star.dev/guide/getting_started

This Python example utilizes the datastar_py library to generate SSE events for patching DOM elements. It defines an asynchronous Sanic endpoint that yields SSE patch elements after a short delay.

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

### Preserve Attributes During DOM Morphing with DataStar

Source: https://data-star.dev/docs

The `data-preserve-attr` attribute ensures that specified attribute values are maintained when DataStar morphs DOM elements. Multiple attributes can be preserved by separating their names with a space.

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

### Replicate Datastar Pro URL History and Scroll Effects with Free Version

Source: https://data-star.dev/essays

These HTML snippets demonstrate how to replicate two Datastar Pro features using the free version. The first replaces the current URL on load and when a variable changes, while the second scrolls an element into view. Standard HTML attributes and JavaScript expressions are utilized.

```html
1<!-- Replaces the current URL on load and whenever $page changes. -->
<div data-effect="window.history.replaceState({}, '', '/page/' + $page)"></div>

<!-- Scrolls the element into view. -->
<div data-on-load="el.scrollIntoView()"></div>
```

--------------------------------

### Ruby: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Ruby code demonstrates how to patch signals using Server-Sent Events (SSE) within a Rack handler using the DataStar gem. It shows creating a `Datastar::Dispatcher` and using its `stream` method with `sse.patch_signals`.

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

### Example: Using a Selector with contentType 'form'

Source: https://data-star.dev/errors/runtime/fetch_closest_form_not_found

This code snippet illustrates how to resolve the FetchClosestFormNotFound error by providing a 'selector' option when setting contentType to 'form'. This is useful when the button is not directly wrapped by its target form.

```html
<button data-on-click="@post('/endpoint', {contentType: 'form', selector: '#myform'})></button>
```

--------------------------------

### Datastar Expression: Display Signal Value

Source: https://data-star.dev/docs

Demonstrates a basic Datastar expression to display the value of a signal named 'foo'. The `data-text` attribute binds the element's text content to the signal's value.

```html
<div data-signals-foo="1">
    <div data-text="$foo"></div>
</div>
```

--------------------------------

### Replicate Datastar Pro Features with Free Version (HTML)

Source: https://data-star.dev/essays/greedy_developer

These HTML snippets demonstrate how to replicate two Datastar Pro features using the free version. The first snippet replaces the current URL on load and when a variable changes, while the second scrolls an element into view.

```HTML
1<!-- Replaces the current URL on load and whenever $page changes. -->
2<div data-effect="window.history.replaceState({}, '', '/page/' + $page)"></div>

3<!-- Scrolls the element into view. -->
4<div data-on-load="el.scrollIntoView()"></div>
```

--------------------------------

### Debounce Event Listener with data-on-resize and Modifiers

Source: https://data-star.dev/reference/attributes

The `data-on-resize` attribute can be combined with the `__debounce` modifier to delay the execution of an expression until after a certain period of inactivity. This example debounces the resize event listener to 10 milliseconds.

```html
<div data-on-resize__debounce.10ms="$count++"></div>
```

--------------------------------

### Clojure: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Clojure code snippet demonstrates how to generate Server-Sent Events (SSE) to patch signals using the DataStar SDK. It shows setting up a Ring handler that sends an 'affirmative' signal, waits, and then sends a reset signal.

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

### Create Signals with data-signals Attribute

Source: https://data-star.dev/guide/reactive_signals

The `data-signals` attribute is used to initialize signals directly on HTML elements. If a signal is accessed before creation, it defaults to an empty string. This attribute can also patch existing signals with new values or expressions.

```html
<div data-signals-foo="1"></div>
<div data-signals-form.foo="2"></div>
<div data-signals="{foo: 1, form: {foo: 2}}"></div>
```

--------------------------------

### Signal and Class Name Casing with __case Modifier

Source: https://data-star.dev/docs

The `__case` modifier can be applied to `data-bind`, `data-class`, and `data-computed` attributes to automatically convert signal or class names to different casing conventions like camelCase, kebab-case, snake_case, or PascalCase.

```html
<input data-bind-my-signal__case.kebab />
<div data-class-my-class__case.camel="$foo"></div>
<div data-computed-my-signal__case.kebab="$bar + $baz"></div>
```

--------------------------------

### Filter Signals with DataStar

Source: https://data-star.dev/docs

The `data-on-signal-patch-filter` attribute refines which signals trigger the `data-on-signal-patch` attribute. It accepts an object with `include` and/or `exclude` properties for regular expression filtering, allowing precise control over observed signal changes.

```html
<!-- Only react to counter signal changes -->
<div data-on-signal-patch-filter="{include: /^counter$/}"></div>

<!-- React to all changes except those ending with "changes" -->
<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>

<!-- Combine include and exclude filters -->
<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
```

--------------------------------

### Computed Signals with data-computed

Source: https://data-star.dev/docs

The `data-computed` attribute creates a read-only signal whose value is derived from an expression involving other signals. The computed signal automatically updates when its dependencies change. Expressions should not perform side effects.

```html
<div data-computed-foo="$bar + $baz"></div>
<div data-computed-foo="$bar + $baz"></div>
<div data-text="$foo"></div>
```

--------------------------------

### Ruby (Rack): Patching Signals with SSE

Source: https://data-star.dev/docs

This Ruby code snippet shows how to patch signals using Server-Sent Events (SSE) within a Rack application using the DataStar gem. It instantiates a `Datastar::Dispatcher` and uses its `stream` method to send signal updates with a `sleep` for delay.

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
19
```

--------------------------------

### Datastar Expression: Using the 'el' Variable

Source: https://data-star.dev/guide/datastar_expressions

Shows how to use the `el` variable, which represents the element the `data-*` attribute is attached to, within a Datastar expression. This example dynamically sets the text content to the element's ID.

```html
1<div id="foo" data-text="el.id"></div>
```

--------------------------------

### Run Expression on Signal Patch with DataStar

Source: https://data-star.dev/docs

The `data-on-signal-patch` attribute executes an expression when a signal is patched. The `patch` variable is available for signal patch details. It can be filtered using `data-on-signal-patch-filter` and modified with timing modifiers like `__delay`, `__debounce`, and `__throttle`.

```html
<div data-on-signal-patch="console.log('A signal changed!')"></div>
<div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
<div data-on-signal-patch__debounce.500ms="doSomething()"></div>
```

--------------------------------

### Example: Using a Wrapping Form with contentType 'form'

Source: https://data-star.dev/errors/runtime/fetch_closest_form_not_found

This code snippet demonstrates how to resolve the FetchClosestFormNotFound error by using a wrapping form element when setting the contentType option to 'form'. The button within the form will correctly associate with it.

```html
<form>
    <button data-on-click="@post('/endpoint', {contentType: 'form'} )></button>
</form>
```

--------------------------------

### Persist Signals in Session Storage with data-persist-__session

Source: https://data-star.dev/reference/attributes

The `__session` modifier for `data-persist` changes the storage mechanism from local storage to session storage. This example uses a custom key 'mykey' and stores data in session storage.

```html
<!-- Persists signals using a custom key `mykey` in session storage -->
<div data-persist-mykey__session></div>
```

--------------------------------

### JavaScript Web Component for String Reversal

Source: https://data-star.dev/examples/web_component

This JavaScript code defines a custom web component named 'ReverseComponent'. It observes the 'name' attribute, and upon changes, it reverses the string provided in the 'newValue' and dispatches a custom 'reverse' event with the reversed string in its detail.

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

### HTML: Displaying an Editable Table Row

Source: https://data-star.dev/examples/edit_row

This code snippet shows a table row in its editable state. It replaces the display text with input fields for 'Name' and 'Email', and updates the action buttons to 'Cancel' and 'Save'. The 'Cancel' button reverts the row to read-only, while the 'Save' button submits changes via a PATCH request.

```html
<tr>
    <td>
        <input type="text" data-bind-name>
    </td>
    <td>
        <input type="text" data-bind-email>
    </td>
    <td>
        <button data-on-click="@get('/examples/edit_row/cancel')">
            Cancel
        </button>
        <button data-on-click="@patch('/examples/edit_row/0')">
            Save
        </button>
    </td>
</tr>
```

--------------------------------

### Persist Signals in Local Storage with data-persist

Source: https://data-star.dev/docs

The `data-persist` attribute stores signals in local storage, preserving values across page loads. It supports filtering signals using `include` and `exclude` regular expressions, and custom storage keys can be defined using `data-persist-`.

```html
<div data-persist></div>
```

```html
<div data-persist="{include: /foo/, exclude: /bar/}"></div>
```

```html
<div data-persist-mykey></div>
```

--------------------------------

### Patch Elements using SSE Generator in Java (Spring)

Source: https://data-star.dev/guide/getting_started

This Java example shows how to use the Datastar SDK with a Spring `HttpServletResponseAdapter` to generate Server-Sent Events (SSE) for patching HTML elements. It utilizes `ServerSentEventGenerator` and `PatchElements.builder()` to construct and send the DOM update events, including a delay.

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

### Python (Sanic): Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Python snippet showcases how to patch signals using Server-Sent Events (SSE) within a Sanic web application using the DataStar Python library. It uses an async generator with `SSE.patch_signals` to send updates.

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

### Kotlin: Patch Signals via SSE

Source: https://data-star.dev/guide/reactive_signals

This Kotlin code snippet shows how to patch signals using Server-Sent Events (SSE) with the DataStar SDK. It utilizes the `ServerSentEventGenerator` and its `patchSignals` function to send signal data.

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

### Go Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Illustrates a Go implementation using `github.com/starfederation/datastar-go/datastar` to send Server-Sent Events (SSE) that update an element with the current time.

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

### Using Selector with contentType: 'form'

Source: https://data-star.dev/errors/runtime/sse_closest_form_not_found

This example shows how to fix the SseClosestFormNotFound error by providing a `selector` option to the `@post` action, targeting an existing form in the DOM when `contentType: 'form'` is used. This is useful when a direct wrapping form is not feasible.

```html
<button data-on-click="@post('/endpoint', {contentType: 'form', selector: '#myform'}) "></button>
```

--------------------------------

### Using Wrapping Form with contentType: 'form'

Source: https://data-star.dev/errors/runtime/sse_closest_form_not_found

This example demonstrates how to resolve the SseClosestFormNotFound error by placing the button within a wrapping HTML form when using `contentType: 'form'` with the `@post` action. This ensures a form element is available for the action to reference.

```html
<form>
    <button data-on-click="@post('/endpoint', {contentType: 'form'}) "></button>
</form>
```

--------------------------------

### Throttle Event Listener with data-on-raf

Source: https://data-star.dev/reference/attributes

The `data-on-raf` attribute triggers an expression on the next animation frame. Modifiers like `__throttle` can be applied to limit the rate at which the event listener is invoked. This example throttles the listener to 10 milliseconds.

```html
<div data-on-raf__throttle.10ms="$count++"></div>
```

--------------------------------

### Python: Load More Data with SSE

Source: https://data-star.dev/how_tos/load_more_list_items

This Python snippet, using FastAPI and DataStar, implements a 'load more' functionality. It yields Server-Sent Events to patch elements, append new items, update the offset signal, or remove the 'load more' button when the maximum number of items is reached. It utilizes `ServerSentEventGenerator` and `ElementPatchMode` from the `datastar_py` library.

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

### JSON Payload for Patching Signals

Source: https://data-star.dev/guide/reactive_signals

This JSON object represents a typical payload used to patch frontend signals. When received with a 'content-type' of 'application/json', DataStar merges these values into the existing frontend signals, updating the UI dynamically.

```json
1{"hal": "Affirmative, Dave. I read you."}
```

--------------------------------

### Create Element References with DataStar Data-Ref

Source: https://data-star.dev/docs

The `data-ref` attribute creates a new signal that references the element it's attached to. The signal name can be defined in the attribute key or value. Modifiers like `__case` with options `.camel`, `.kebab`, `.snake`, and `.pascal` can alter the signal name's casing.

```html
<div data-ref-foo></div>
<div data-ref="foo"></div>

$foo is a reference to a <span data-text="$foo.tagName"></span> element

<div data-ref-my-signal__case.kebab></div>
```

--------------------------------

### Patch HTML Elements via SSE Event (HTML)

Source: https://data-star.dev/guide/getting_started

This example shows how to send a 'datastar-patch-elements' SSE event to update HTML elements in the DOM. The server sends HTML content prefixed with 'elements', and Datastar's morphing strategy updates the DOM based on element IDs.

```html
event: datastar-patch-elements
data: elements <div id="hal">
    I’m sorry, Dave. I’m afraid I can’t do that.
</div>
```

--------------------------------

### Reading Nested Signals in C#

Source: https://data-star.dev/guide/backend_requests

Demonstrates how to set up DataStar services and read nested signals from an incoming request using C# with ASP.NET Core.

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

### Content Security Policy for Datastar

Source: https://data-star.dev/docs

This snippet demonstrates how to configure a Content Security Policy (CSP) to allow Datastar's expression evaluation, which requires 'unsafe-eval'. This is crucial for preventing security vulnerabilities like Cross-Site Scripting (XSS) attacks when using Datastar expressions that execute JavaScript.

```html
<meta http-equiv="Content-Security-Policy"
    content="script-src 'self' 'unsafe-eval';"
>
```

--------------------------------

### Golang: Morph Multiple Random SVG Elements

Source: https://data-star.dev/examples/svg_morphing

This Go example shows how to update multiple SVG elements simultaneously using Datastar's SSE. It morphs three circles, each with a random radius and fill color, updating the SVG element with the ID 'multi-demo'.

```go
1svgMorphingRouter.Get("/multiple_elements", func(w http.ResponseWriter, r *http.Request) {
2    sse := datastar.NewSSE(w, r)
3    color1 := svgColors[rand.N(len(svgColors))]
4    color2 := svgColors[rand.N(len(svgColors))]
5    color3 := svgColors[rand.N(len(svgColors))]
6    r1 := 10 + rand.N(20) // radius 10-30
7    r2 := 10 + rand.N(20)
8    r3 := 10 + rand.N(20)
9    sse.PatchElements(fmt.Sprintf(`<svg id="multi-demo">
10        <circle cx="30" cy="30" r="%d" fill="%s" />
11        <circle cx="70" cy="30" r="%d" fill="%s" />
12        <circle cx="50" cy="70" r="%d" fill="%s" />
13    </svg>`, r1, color1, r2, color2, r3, color3))
14})
```

--------------------------------

### Patch Signals with data-signals Attribute

Source: https://data-star.dev/reference

The `data-signals` attribute allows for patching (adding, updating, or removing) signals. Values can be set directly, nested using dot-notation, or as multiple key-value pairs. Setting a signal to `null` removes it. Keys are converted to camel case. Signals starting with an underscore are not sent to the backend by default.

```html
<div data-signals-foo="1"></div>
```

```html
<div data-signals-foo.bar="1"></div>
```

```html
<div data-signals="{foo: {bar: 1, baz: 2}}"></div>
```

```html
<div data-signals="{foo: null}"></div>
```

--------------------------------

### Nesting Signals with Object Syntax in HTML

Source: https://data-star.dev/guide/backend_requests

Illustrates nesting signals using JavaScript object syntax within an HTML element's `data-signals` attribute.

```html
<div data-signals="{foo: {bar: 1}}"></div>
```

--------------------------------

### Patch Elements using SSE Generator in Go

Source: https://data-star.dev/guide/getting_started

This Go code demonstrates using the Datastar SDK to generate Server-Sent Events (SSE) for patching HTML elements into the DOM. It creates an `SSE` instance and uses `PatchElements` to send the HTML content, including a one-second delay between updates.

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

### HTML Structure for Templ Counter

Source: https://data-star.dev/examples/templ_counter

This HTML snippet defines the structure for the Templ Counter demo. It includes two buttons, one for a global counter and another for a user-specific counter. The elements are styled using flexbox and CSS variables. Data attributes are used to specify event handlers for loading and clicking, which trigger server-side updates.

```html
<div
    style="display: flex; gap: var(--size-6)"
    data-on-load="@get('/examples/templ_counter/updates')"
>
    <!-- Global Counter -->
    <button
        id="global"
        class="info"
        data-on-click="@patch('/examples/templ_counter/global')"
    >
        Global Clicks: 0
    </button>

    <!-- User Counter -->
    <button
        id="user"
        class="success"
        data-on-click="@patch('/examples/templ_counter/user')"
    >
        User Clicks: 0
    </button>
</div>
```

--------------------------------

### Datastar SSE PatchElements for DOM Updates

Source: https://data-star.dev/index

This Go snippet illustrates how to use Datastar's Server-Sent Events (SSE) to dynamically update the DOM. The `sse.PatchElements` function is used to replace the content of the `#hal` div with new HTML. A short delay is introduced before patching again.

```go
sse.PatchElements(`
    <div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>
`)
time.Sleep(1 * time.Second)
sse.PatchElements(`<div id="hal">Waiting for an order...</div>`)
```

--------------------------------

### Using data-custom-validity for Input Validation

Source: https://data-star.dev/errors/runtime/custom_validity_key_not_allowed

This example demonstrates how to use the `data-custom-validity` attribute for input validation. It ensures that the values of two input fields ('foo' and 'bar') are identical. If they are not, a custom error message is displayed. The attribute expects a string expression that evaluates to the validity message.

```html
<form>
    <input data-bind-foo name="foo" />
    <input data-bind-bar name="bar" 
        data-custom-validity="$foo === $bar ? '' : 'Field values must be the same.'" 
    />
    <button>
        Submit form
    </button>
</form>
```

--------------------------------

### Datastar: Patch Elements using Server-Sent Events (SSE)

Source: https://data-star.dev/guide

Illustrates how to send 'datastar-patch-elements' SSE events to update the DOM. This method allows for multiple element patches within a single stream and is suitable for long-lived connections. The example shows updating an element with ID 'hal' first with one message and then with another after a delay.

```sse
event: datastar-patch-elements
data: elements <div id="hal">
data:     I’m sorry, Dave. I’m afraid I can’t do that.
data: </div>
```

```sse
event: datastar-patch-elements
data: elements <div id="hal">
data:     Waiting for an order...
data: </div>
```

--------------------------------

### Attach Event Listeners with data-on

Source: https://data-star.dev/docs

The `data-on` attribute attaches an event listener to an element, executing a specified expression when the event is triggered. An `evt` variable, representing the event object, is available within the expression. This attribute supports standard browser events and custom events, with `data-on-submit` specifically preventing default form submission behavior.

```html
<button data-on-click="$foo = ''">Reset</button>
```

```html
<div data-on-myevent="$foo = evt.detail"></div>
```

--------------------------------

### HTML data-bind Attribute for Two-Way Data Binding

Source: https://data-star.dev/guide/reactive_signals

The `data-bind` attribute in Datastar enables two-way data binding for input elements. It automatically creates and manages a signal, synchronizing its value with the element's value. Changes in either the signal or the element's value are reflected in the other. This attribute supports direct signal assignment or assigning the signal name as a string value.

```html
<input data-bind-foo />
```

```html
<input data-bind="foo" />
```

--------------------------------

### HTML data-attr Attribute for Dynamic Attribute Binding

Source: https://data-star.dev/guide/reactive_signals

The `data-attr` attribute allows binding the value of any HTML attribute to a Datastar expression. This enables dynamic modification of attributes like `disabled` or `title` based on frontend state. Similar to `data-class`, it supports setting multiple attributes at once using an object literal, where keys are attribute names and values are their corresponding expressions.

```html
<input data-bind-foo />
<button data-attr-disabled="$foo == ''">
    Save
</button>
```

```html
<button data-attr="{disabled: $foo == '', title: $foo}">Save</button>
```

--------------------------------

### HTML data-class Attribute for Dynamic Class Management

Source: https://data-star.dev/guide/reactive_signals

The `data-class` attribute enables conditional addition or removal of CSS classes to an HTML element based on a Datastar expression. A single class can be toggled by providing the class name and a boolean expression. Multiple classes can be managed simultaneously using an object literal where keys are class names and values are their corresponding boolean expressions.

```html
<input data-bind-foo />
<button data-class-success="$foo != ''">
    Save
</button>
```

```html
<button data-class="{success: $foo != '', 'font-bold': $foo == 'bar'}">
    Save
</button>
```

--------------------------------

### HTML data-text Attribute for Displaying Signal Values

Source: https://data-star.dev/guide/reactive_signals

The `data-text` attribute dynamically sets the text content of an HTML element to the value of a Datastar signal. It requires the signal name to be prefixed with a `$`. This attribute also supports Datastar expressions, allowing JavaScript functions like `toUpperCase()` to be applied to the signal value before displaying it.

```html
<input data-bind-foo />
<div data-text="$foo"></div>
```

```html
<input data-bind-foo />
<div data-text="$foo.toUpperCase()"></div>
```

--------------------------------

### Node.js Backend for Real-time Time Updates

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Illustrates a Node.js implementation using the `node:http` module and `ServerSentEventGenerator` to stream updates containing the current time to the client.

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

### HTML data-show Attribute for Conditional Element Visibility

Source: https://data-star.dev/guide/reactive_signals

The `data-show` attribute controls the visibility of an HTML element based on the boolean evaluation of a Datastar expression. If the expression evaluates to `true`, the element is shown; otherwise, it is hidden. For preventing content flash, it's recommended to initially style the element with `display: none;`.

```html
<input data-bind-foo />
<button data-show="$foo != ''">
    Save
</button>
```

```html
<input data-bind-foo />
<button data-show="$foo != ''" style="display: none;">
    Save
</button>
```

--------------------------------

### Immediate Execution with `.leading` Modifier

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Shows how to combine the `data-on-interval` attribute with the `.leading` modifier to execute the expression immediately upon page load, in addition to the regular interval.

```html
<div id="time"
     data-on-interval__duration.5s.leading="@get('/endpoint')"
></div>
```

--------------------------------

### Dispatch and Listen to Custom Event with HTML and JavaScript

Source: https://data-star.dev/examples/custom_event

This snippet shows an HTML paragraph with data attributes and a JavaScript script. The script dispatches a custom event 'myevent' every second, and the HTML's data attributes listen to this event to update its text content with the event's details. It utilizes `CustomEvent` for dispatching and `data-on-myevent` for listening.

```html
<p
    id="foo"
    data-signals-_event-details
    data-on-myevent="$_eventDetails = evt.detail"
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

### HTML data-computed Attribute for Derived Signals

Source: https://data-star.dev/guide/reactive_signals

The `data-computed` attribute allows the creation of a new, read-only signal whose value is derived from a Datastar expression involving other signals. This computed signal automatically updates whenever any of its dependent signals change. It's useful for memoizing complex expressions and managing derived state efficiently.

```html
<input data-bind-foo />
<div data-computed-repeated="$foo.repeat(2)" data-text="$repeated"></div>
```

--------------------------------

### Datastar HTML with Embedded Bad Apple Video Benchmark

Source: https://data-star.dev/examples/bad_apple

This HTML snippet utilizes Datastar's data-signals attribute to manage the state of the Bad Apple video benchmark, including percentage and content. The data-on-load attribute triggers the frame streaming from a specified endpoint. The `pre` tag is dynamically updated with video frames, and a range input visually represents the playback progress.

```html
<label
    data-signals="{_percentage: 0, _contents: 'bad apple frames go here'}"
    data-on-load="@get('/examples/bad_apple/updates')"
>
    <span data-text="`Percentage: ${$_percentage.toFixed(2)}%`"></span>
    <input
        type="range"
        min="0"
        max="100"
        step="0.01"
        disabled
        style="cursor: default"
        data-attr-value="$_percentage"
    />
</label>
<pre style="line-height: 100%" data-text="$_contents"></pre>
```

--------------------------------

### Send SSE Events with Datastar in Clojure

Source: https://data-star.dev/guide/backend_requests

This snippet demonstrates how to set up a backend controller action in Clojure using Datastar's SDK to stream SSE events. It includes patching elements and signals into the DOM. Requires the `starfederation.datastar.clojure` SDK.

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
11                   :on-open
12                   (fn [sse] 
13                     ;; Patches elements into the DOM
14                     (d*/patch-elements! sse
15                                         "<div id=\"question\">What do you put in a toaster?</div>")
16
17                     ;; Patches signals
18                     (d*/patch-signals! sse "{response: '', answer: 'bread'}"))}))
```

--------------------------------

### Nesting Signals with Dot Notation in HTML

Source: https://data-star.dev/guide/backend_requests

Demonstrates how to nest signals using dot notation within an HTML element's `data-signals` attribute.

```html
<div data-signals-foo.bar="1"></div>
```

--------------------------------

### Datastar TodoMVC HTML Structure

Source: https://data-star.dev/examples/todomvc

This HTML code defines the user interface for the TodoMVC application. It includes input fields for new todos, checkboxes for marking completion, buttons for filtering and resetting the list, and a container for dynamically rendered todo items. Datastar-specific attributes like 'data-on-load', 'data-on-click', and 'data-signals-input' are used to bind UI elements to application logic and server-side operations.

```html
<section
    id="todomvc"
    data-on-load="@get('/examples/todomvc/updates')"
>
    <header id="todo-header">
        <input
            type="checkbox"
            data-on-click__prevent="@post('/examples/todomvc/-1/toggle')"
            data-on-load="el.checked = false"
        />
        <input
            id="new-todo"
            type="text"
            placeholder="What needs to be done?"
            data-signals-input
            data-bind-input
            data-on-keydown="
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
        <button class="small info" data-on-click="@put('/examples/todomvc/mode/0')">
            All
        </button>
        <button class="small" data-on-click="@put('/examples/todomvc/mode/1')">
            Pending
        </button>
        <button class="small" data-on-click="@put('/examples/todomvc/mode/2')">
            Completed
        </button>
        <button class="error small" aria-disabled="true">
            Delete
        </button>
        <button class="warning small" data-on-click="@put('/examples/todomvc/reset')">
            Reset
        </button>
    </div>
</section>
```

--------------------------------

### Send SSE Events with Datastar (JavaScript - Server-side)

Source: https://data-star.dev/guide/backend_requests

This JavaScript snippet demonstrates initiating an SSE stream from the server-side using Datastar. It sets up the necessary headers and provides a callback for sending events.

```javascript
1// Creates a new `ServerSentEventGenerator` instance (this also sends required headers)
2ServerSentEventGenerator.stream(req, res, (stream) => {
3});
```

--------------------------------

### Stream DOM Patching with SSE in Ruby

Source: https://data-star.dev/guide

Demonstrates streaming Server-Sent Events (SSE) in Ruby using the Datastar gem to patch elements into the DOM. It initializes a Datastar::Dispatcher and uses a stream block to send patch commands with a delay.

```ruby
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

### Rendering Current Time with Backend Templating

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Illustrates rendering dynamic content, like the current time, within an element that also uses `data-on-interval`. The backend is expected to provide the `now` variable.

```html
<div id="time"
     data-on-interval__duration.5s="@get('/endpoint')"
>
     {{ now }}
</div>
```

--------------------------------

### DataStar HTML with Web Component Integration

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates how to use a custom web component ('my-component') within HTML, passing data via attributes and listening for custom events. It utilizes 'data-bind' for input binding and 'data-attr-src'/'data-on-mycustomevent' for component interaction, driving reactivity with 'data-*' attributes.

```html
<div data-signals-result="''">
    <input data-bind-foo />
    <my-component
        data-attr-src="$foo"
        data-on-mycustomevent="$result = evt.detail.value"
    ></my-component>
    <span data-text="$result"></span>
</div>
```

--------------------------------

### JSON Output for Signal Changes

Source: https://data-star.dev/examples/on_signal_patch

These JSON snippets represent the initial state or filtered output for signal changes. The first shows an empty array for `counterChanges`, indicating no counter modifications yet. The second shows an empty array for `allChanges`, signifying no general signal modifications.

```json
{"counterChanges":[]}
```

```json
{"allChanges":[]}
```

--------------------------------

### PHP: Append Item and Update Offset Signal

Source: https://data-star.dev/how_tos/load_more_list_items

This PHP snippet shows how to append a new item to a list and update the offset signal using the ServerSentEventGenerator. It reads signals, defines limits, and conditionally patches elements or signals based on the current offset. The `patchElements` method is used for DOM updates, and `patchSignals` for signal synchronization.

```php
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

### Generate SSE Events to Patch DOM Elements (Ruby)

Source: https://data-star.dev/guide/getting_started

This Ruby snippet demonstrates using the Datastar gem to stream SSE events. It shows how to instantiate Datastar and use the stream method to send patch elements, including a one-second delay between updates.

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

--------------------------------

### Display Contact Details for Click to Edit

Source: https://data-star.dev/examples/click_to_edit

This HTML snippet displays contact information (First Name, Last Name, Email) and includes 'Edit' and 'Reset' buttons. The 'Edit' button fetches the editing UI from '/examples/click_to_edit/edit', and the 'Reset' button triggers a reset action at '/examples/click_to_edit/reset'. It utilizes data attributes for dynamic behavior and indicators.

```html
<div id="demo">
    <p>First Name: John</p>
    <p>Last Name: Doe</p>
    <p>Email: joe@blow.com</p>
    <div role="group">
        <button
            class="info"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
            data-on-click="@get('/examples/click_to_edit/edit')"
        >
            Edit
        </button>
        <button
            class="warning"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
            data-on-click="@patch('/examples/click_to_edit/reset')"
        >
            Reset
        </button>
    </div>
</div>
```

--------------------------------

### Reading Signals in Python (FastAPI)

Source: https://data-star.dev/guide/backend_requests

Illustrates reading signals within a FastAPI application in Python, using `read_signals` helper function.

```python
from datastar_py.fastapi import datastar_response, read_signals

@app.get("/updates")
@datastar_response
async def updates(request: Request):
    # Retrieve a dictionary with the current state of the signals from the frontend
    signals = await read_signals(request)
```

--------------------------------

### Editable Form for Contact Details

Source: https://data-star.dev/examples/click_to_edit

This HTML snippet renders an editable form for contact details, including input fields for First Name, Last Name, and Email. It features 'Save' and 'Cancel' buttons. The 'Save' button sends an update to '/examples/click_to_edit' using a PUT request, while 'Cancel' reverts changes by fetching from '/examples/click_to_edit/cancel'. Input fields are data-bound and disabled during fetching.

```html
<div id="demo">
    <label>
        First Name
        <input
            type="text"
            data-bind-first-name
            data-attr-disabled="$_fetching"
        >
    </label>
    <label>
        Last Name
        <input
            type="text"
            data-bind-last-name
            data-attr-disabled="$_fetching"
        >
    </label>
    <label>
        Email
        <input
            type="email"
            data-bind-email
            data-attr-disabled="$_fetching"
        >
    </label>
    <div role="group">
        <button
            class="success"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
            data-on-click="@put('/examples/click_to_edit')"
        >
            Save
        </button>
        <button
            class="error"
            data-indicator-_fetching
            data-attr-disabled="$_fetching"
            data-on-click="@get('/examples/click_to_edit/cancel')"
        >
            Cancel
        </button>
    </div>
</div>
```

--------------------------------

### Send SSE Events with Datastar in Go

Source: https://data-star.dev/guide/backend_requests

This Go snippet illustrates creating a ServerSentEventGenerator instance to send SSE events. It shows how to patch HTML elements and JSON signals using the datastar-go library.

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

### Implement Dynamic Polling Interval in Go

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Go snippet shows how to set dynamic polling intervals for Server-Sent Events using the Datastar Go library. It calculates the duration based on the current second, switching to 5-second intervals normally and 1-second intervals for the last 10 seconds of a minute. It uses `datastar.NewSSE` to create an SSE instance.

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

### Copy Text to Clipboard using @clipboard()

Source: https://data-star.dev/reference/actions

This snippet shows how to use the `@clipboard()` directive to copy text to the user's clipboard. It supports copying plain text directly and also handles Base64 encoded text by decoding it before copying, which is useful for special characters or complex data.

```html
<!-- Copy plain text -->
<button data-on-click="@clipboard('Hello, world!')"></button>

<!-- Copy base64 encoded text (will decode before copying) -->
<button data-on-click="@clipboard('SGVsbG8sIHdvcmxkIQ==', true)"></button>
```

--------------------------------

### Backend Event Handling for DataStar List Updates

Source: https://data-star.dev/how_tos/load_more_list_items

These snippets illustrate backend events used by DataStar to update the UI. 'datastar-patch-elements' is used to append new items to the '#list' container, and 'datastar-patch-signals' updates the offset. The 'remove' mode is used to hide the button when all items are loaded.

```text
1event: datastar-patch-elements
2data: selector #list
3data: mode append
4data: elements <div>Item 2</div>
```

```text
1event: datastar-patch-signals
2data: signals {offset: 2}
```

```text
1event: datastar-patch-elements
2data: selector #load-more
3data: mode remove
```

--------------------------------

### Datastar Starfield Web Component Attributes

Source: https://data-star.dev/index

This snippet shows how to configure a Datastar web component (`ds-starfield`) using reactive signals. The attributes `data-attr-center-x`, `data-attr-center-y`, and `data-attr-speed` are bound to JavaScript variables ($x, $y, $speed) for dynamic updates.

```html
<ds-starfield
    data-attr-center-x="$x"
    data-attr-center-y="$y"
    data-attr-speed="$speed"
></ds-starfield>
```

--------------------------------

### DataStar Backend Action for Script Execution

Source: https://data-star.dev/guide/datastar_expressions

Illustrates how a button click can trigger a backend action that returns JavaScript to be executed on the frontend. The 'data-on-click' attribute is used to specify the endpoint, and if the response has a 'text/javascript' content type, the script is run in the browser.

```html
<button data-on-click="@get('/endpoint')">
    What are you talking about, HAL?
</button>
```

--------------------------------

### Datastar Expression: Basic Signal Usage

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates the basic usage of Datastar expressions by referencing a signal `$foo` within a `data-text` attribute. The output reflects the current value of the signal.

```html
1<div data-signals-foo="1">
2    <div data-text="$foo"></div>
3</div>
```

--------------------------------

### Send Multiple HTML Patches via SSE Events (HTML)

Source: https://data-star.dev/guide/getting_started

This snippet illustrates sending a sequence of 'datastar-patch-elements' SSE events to update the DOM over time. It first sets a response and then, after a delay, updates the same element with new content, demonstrating dynamic content changes.

```html
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

### Configure Content Security Policy for Datastar

Source: https://data-star.dev/reference/security

This snippet demonstrates how to configure a Content Security Policy (CSP) to allow Datastar expressions to be evaluated. The `script-src` directive must include `'unsafe-eval'` to permit the use of the `Function()` constructor, which Datastar relies on for expression evaluation. Ensure this is used cautiously and only when necessary.

```html
<meta http-equiv="Content-Security-Policy" 
    content="script-src 'self' 'unsafe-eval';">

```

--------------------------------

### Backend Response with `datastar-patch-elements` Event

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

Details the structure of a `datastar-patch-elements` event response from the backend, which contains updated HTML for the specified element, including dynamic content.

```http
event: datastar-patch-elements
data: elements <div id="time" data-on-interval__duration.5s="@get('/endpoint')">
data: elements     {{ now }}
data: elements </div>
```

--------------------------------

### Clojure: Generating SSE for Patching Elements with Datastar

Source: https://data-star.dev/guide

Shows how to generate Server-Sent Events (SSE) in Clojure using the Datastar SDK to patch elements into the DOM. It demonstrates creating an SSE response and asynchronously sending patch commands to update an HTML element with a specific ID. A delay is included between patches.

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

### Go SDK for Executing Scripts via SSE

Source: https://data-star.dev/guide/datastar_expressions

Provides a Go code snippet using the DataStar SDK to generate a Server-Sent Event (SSE) that executes a JavaScript alert. This function simplifies the process of sending script execution commands from the backend to the frontend.

```go
sse := datastar.NewSSE(writer, request)
sse.ExecuteScript(`alert('This mission is too important for me to allow you to jeopardize it.')`)
```

--------------------------------

### Generate SSE Events to Patch DOM Elements (PHP)

Source: https://data-star.dev/guide/getting_started

This snippet shows how to create a ServerSentEventGenerator instance in PHP and use it to patch elements into the DOM. It demonstrates sending an initial message, pausing for a second, and then sending an updated message.

```php
use starfederation\datastar\ServerSentEventGenerator;

// Creates a new `ServerSentEventGenerator` instance.
$sse = new ServerSentEventGenerator();

// Patches elements into the DOM.
$sse->patchElements(
    '<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>'
);

sleep(1)

$sse->patchElements(
    '<div id="hal">Waiting for an order...</div>'
);
```

--------------------------------

### Send Multiple SSE Events using Python (DatastarResponse)

Source: https://data-star.dev/guide/backend_requests

This Python snippet shows how to construct a DatastarResponse that includes multiple Server-Sent Events (SSE). It uses the `SSE` helper class to create `patch_elements` and `patch_signals` events, returning them within a list. This is a common pattern for returning SSE responses in Python.

```python
return DatastarResponse([
    SSE.patch_elements('<div id="question">...</div>'),
    SSE.patch_elements('<div id="instructions">...</div>'),
    SSE.patch_signals({"answer": "...", "prize": "..."})
])
```

--------------------------------

### HTML with SortableJS and DataStar Event Binding

Source: https://data-star.dev/examples/sortable

This HTML snippet sets up a sortable list using SortableJS and binds a custom event handler using DataStar's `data-on-*` attribute. The `data-signals-order-info` attribute is used to display signal data, and the `data-on-reordered` attribute listens for the 'reordered' event to update the `$orderInfo` signal.

```html
<div data-signals-order-info="'Initial order'" data-text="$orderInfo"></div>
<div id="sortContainer" data-on-reordered="$orderInfo = event.detail.orderInfo">
    <button>Item 1</button>
    <button>Item 2</button>
    <button>Item 3</button>
    <button>Item 4</button>
    <button>Item 5</button>
</div>
```

--------------------------------

### Send POST Request using Datastar Backend Action

Source: https://data-star.dev/guide/backend_requests

This snippet shows how to trigger a POST request to '/actions/quiz' when a button is clicked, utilizing Datastar's `@post()` backend action. It requires a button element with the `data-on-click` attribute.

```html
<button data-on-click="@post('/actions/quiz')">
    Submit answer
</button>
```

--------------------------------

### HTML Form for File Upload with Progress

Source: https://data-star.dev/reference/actions

This HTML form demonstrates how to enable file upload progress monitoring with Datastar Pro. It utilizes `enctype='multipart/form-data'` and `data-on-submit__prevent` to trigger a POST request. The `data-on-datastar-fetch` attribute listens for 'upload-progress' events to update the UI with the upload status.

```html
<form enctype="multipart/form-data"
    data-signals="{progress: 0, uploading: false}"
    data-on-submit__prevent="@post('https://example.com/upload', {contentType: 'form'})"
    data-on-datastar-fetch="
        if (evt.detail.type !== 'upload-progress') return;

        const {progress, loaded, total} = evt.detail.argsRaw;
        $uploading = true;
        $progress = Number(progress);

        if ($progress >= 100) {
            $uploading = false;
        }
    "
>
    <input type="file" name="files" multiple />
    <button type="submit">Upload</button>
    <progress data-show="$uploading" data-attr-value="$progress" max="100"></progress>
</form>
```

--------------------------------

### Reading Nested Signals in Kotlin

Source: https://data-star.dev/guide/backend_requests

Shows how to define a data class for signals and read them from a request using Kotlin, leveraging JSON decoding.

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

### Server-Sent HTML Response Handling in Data-Star

Source: https://data-star.dev/reference/actions

Illustrates how to configure a server response to send HTML content back to the client. It includes setting the 'Content-Type' to 'text/html' and specifies Data-Star's 'datastar-selector' and 'datastar-mode' headers for targeted DOM patching.

```javascript
response.headers.set('Content-Type', 'text/html')
response.headers.set('datastar-selector', '#my-element')
response.headers.set('datastar-mode', 'inner')
response.body = '<p>New content</p>'
```

--------------------------------

### Kotlin: Append Item and Update Offset Signal

Source: https://data-star.dev/how_tos/load_more_list_items

This Kotlin snippet demonstrates appending a new item to a list and updating the offset signal. It checks if the new offset is less than the maximum limit to decide whether to continue signaling or remove a 'load more' button. It uses `generator.patchElements` for DOM manipulation and `generator.patchSignals` for signal updates.

```kotlin
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
        signals = "{\"offset\": $newOffset}",
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
```

--------------------------------

### Implement Dynamic Polling Interval in Python

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Python snippet demonstrates dynamic polling intervals for Server-Sent Events using StarFederation Datastar. It calculates the duration based on the current second, applying a 1-second interval during the last 10 seconds of each minute. It uses `DatastarResponse` and `SSE.patch_elements` for sending SSE updates.

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

### Two-Way Binding for Nested Signals in HTML

Source: https://data-star.dev/guide/backend_requests

Shows how to achieve two-way binding for nested signals using the `data-bind` attribute in HTML.

```html
<input data-bind-foo.bar />
```

--------------------------------

### Send Multiple SSE Events using Ruby (Yield)

Source: https://data-star.dev/guide/backend_requests

This Ruby snippet shows an alternative way to send Server-Sent Events (SSE) using `yield` with `PatchElements` and `PatchSignals` classes. This pattern is useful in specific DSL contexts or when building event streams programmatically. It requires explicit creation of event objects.

```ruby
yield PatchElements::new("<div id='question'>...</div>").into()
yield PatchElements::new("<div id='instructions'>...</div>").into()
yield PatchSignals::new("{answer: '...', prize: '...'}").into()
```

--------------------------------

### HTML Structure with Data Attributes

Source: https://data-star.dev/guide/backend_requests

This HTML snippet demonstrates the use of data attributes for frontend interactivity. `data-signals` initializes signals, `data-computed-correct` defines a computed property, `data-on-click` triggers actions on click events, and `data-show` conditionally displays elements. It also shows how backend events can update the DOM and signals.

```html
<div
    data-signals="{response: '', answer: ''}"
    data-computed-correct="$response.toLowerCase() == $answer"
>
    <div id="question"></div>
    <button data-on-click="@get('/actions/quiz')">Fetch a question</button>
    <button
        data-show="$answer != ''"
        data-on-click="$response = prompt('Answer:') ?? ''"
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

### Implement Dynamic Polling Interval in Rust

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Rust snippet demonstrates dynamic polling intervals for Server-Sent Events using the Datastar Rust crate. It calculates the interval duration based on the current second, applying a 1-second interval during the last 10 seconds of each minute. It uses `async_stream::stream` and `PatchElements::new` to construct the SSE data.

```rust
1use datastar::prelude::*;
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

### Generate SSE Events to Patch DOM Elements (JavaScript)

Source: https://data-star.dev/guide/getting_started

This JavaScript snippet illustrates generating SSE events for patching DOM elements using the ServerSentEventGenerator. It sets up a stream that sends an initial message, waits for 1000 milliseconds, and then sends an updated message.

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

### Send SSE Events with Datastar in Rust

Source: https://data-star.dev/guide/backend_requests

This Rust snippet shows how to create an SSE stream using the datastar crate and `async_stream`. It yields `PatchElements` and `PatchSignals` to the stream.

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

### Send Multiple SSE Events using Java

Source: https://data-star.dev/guide/backend_requests

This Java code demonstrates sending multiple Server-Sent Events (SSE) using Datastar's builder pattern. It involves creating `PatchElements` and `PatchSignals` objects and sending them via a generator. This approach provides a structured way to build and dispatch SSE events.

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

### Generate SSE Patch Elements in Clojure

Source: https://data-star.dev/guide/getting_started

This Clojure code demonstrates how to generate Server-Sent Events (SSE) to patch elements into the DOM using the Datastar SDK. It utilizes `->sse-response` and `d*/patch-elements!` to send HTML content that updates specific DOM elements, with a delay between updates.

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

### JavaScript for SortableJS Initialization and Event Dispatching

Source: https://data-star.dev/examples/sortable

This JavaScript code initializes SortableJS on the 'sortContainer' element. It configures the `animation` and `ghostClass` options and defines an `onEnd` callback. Within `onEnd`, a custom 'reordered' event is dispatched with the old and new index details, which is then caught by DataStar's `data-on-reordered` attribute.

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

### HTML for Infinite Scroll Trigger

Source: https://data-star.dev/examples/infinite_scroll

This HTML snippet demonstrates the core of the infinite scroll pattern. A div with the `data-on-intersect` attribute is used as a trigger. When this element scrolls into the viewport, the specified URL (`/examples/infinite_scroll/more`) is requested, and the response is appended to the page. This attribute is a custom directive for triggering actions based on element visibility.

```html
<div data-on-intersect="@get('/examples/infinite_scroll/more')">
    Loading...
</div>
```

--------------------------------

### Datastar Expressions: Ternary and Logical Operators

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates the use of JavaScript operators within Datastar expressions, including the ternary operator for conditional output, logical OR for showing content based on multiple conditions, and logical AND for conditional actions like sending a request.

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
10<button data-on-click="$landingGearRetracted && @post('/launch')">
11    Launch
12</button>
```

--------------------------------

### Patch Elements using SSE Generator in Kotlin

Source: https://data-star.dev/guide/getting_started

This Kotlin snippet demonstrates using the Datastar SDK's `ServerSentEventGenerator` to patch HTML elements into the DOM. It sends an initial HTML patch, waits for a second, and then sends another patch to update the content dynamically.

```kotlin
val generator = ServerSentEventGenerator(response)

generator.patchElements(
    elements = """<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>""",
)

Thread.sleep(ONE_SECOND)

generator.patchElements(
    elements = """<div id="hal">Waiting for an order...</div>"""
)
```

--------------------------------

### Implement Dynamic Polling Interval in Ruby

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Ruby snippet shows how to dynamically set polling intervals for Server-Sent Events. It calculates the duration based on the current second, switching to 1-second polling for the last 10 seconds of each minute. It uses the Datastar library to patch elements with the updated interval duration.

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

### Custom Storage Key for Persistence with data-persist-

Source: https://data-star.dev/reference/attributes

A custom storage key for `data-persist` can be defined by appending it after `data-persist-`. This allows for distinct storage of different sets of signals.

```html
<div data-persist-mykey></div>
```

--------------------------------

### HTML Form for Inline Email Validation

Source: https://data-star.dev/examples/inline_validation

This HTML snippet defines a form with input fields for email, first name, and last name. The email input is configured for inline validation, triggering a POST request on keydown. It also includes a paragraph for displaying validation messages.

```html
<div id="demo">
    <label>
        Email Address
        <input
            type="email"
            required
            aria-live="polite"
            aria-describedby="email-info"
            data-bind-email
            data-on-keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <p id="email-info" class="info">The only valid email address is "test@test.com".</p>
    <label>
        First Name
        <input
            type="text"
            required
            aria-live="polite"
            data-bind-first-name
            data-on-keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <label>
        Last Name
        <input
            type="text"
            required
            aria-live="polite"
            data-bind-last-name
            data-on-keydown__debounce.500ms="@post('/examples/inline_validation/validate')"
        />
    </label>
    <button
        class="success"
        data-on-click="@post('/examples/inline_validation')"
    >
        <i class="material-symbols:person-add"></i>
        Sign Up
    </button>
</div>
```

--------------------------------

### Sync Query String Params with data-query-string

Source: https://data-star.dev/reference/attributes

The `data-query-string` attribute synchronizes query string parameters with signal values on page load and updates the query string when signals change. This enables bookmarkable states.

```html
<div data-query-string></div>
```

--------------------------------

### Send Multiple SSE Events using Ruby

Source: https://data-star.dev/guide/backend_requests

This Ruby code demonstrates sending multiple Server-Sent Events (SSE) within a `datastar.stream` block. It uses the `sse` object to call `patch_elements` and `patch_signals`. This provides a clear and idiomatic way to handle SSE streams in Ruby.

```ruby
datastar.stream do |sse|
  sse.patch_elements('<div id="question">...</div>')
  sse.patch_elements('<div id="instructions">...</div>')
  sse.patch_signals(answer: '...', prize: '...')
end
```

--------------------------------

### Controlling Request Cancellation with Data-Star Options

Source: https://data-star.dev/reference/actions

Shows how to manage Data-Star's request cancellation behavior. 'disabled' prevents automatic cancellation, allowing concurrent requests. A custom AbortController provides fine-grained manual control over request cancellation.

```html
<!-- Allow concurrent requests (no automatic cancellation) -->
<button data-on-click="@get('/endpoint', {requestCancellation: 'disabled'})">Allow Multiple</button>

<!-- Custom abort controller for fine-grained control -->
<div data-signals-controller="new AbortController()">
    <button data-on-click="@get('/endpoint', {requestCancellation: $controller})">Start Request</button>
    <button data-on-click="$controller.abort()">Cancel Request</button>
</div>
```

--------------------------------

### Send Multiple SSE Events using C# (Stream)

Source: https://data-star.dev/guide/backend_requests

This C# code snippet shows how to send multiple Server-Sent Events (SSE) using the `sse` object, likely within a streaming context. It includes `PatchElements` and `PatchSignals` methods for updating UI elements and sending data. Note the byte array conversion for signals.

```csharp
sse.PatchElements(`<div id="question">...</div>`)
sse.PatchElements(`<div id="instructions">...</div>`)
sse.PatchSignals([]byte(`{answer: '...', prize: '...'}`))
```

--------------------------------

### Server-Sent JSON Response Handling in Data-Star

Source: https://data-star.dev/reference/actions

Demonstrates sending JSON data from the server with Data-Star. It sets the 'Content-Type' to 'application/json' and utilizes the 'datastar-only-if-missing' header to conditionally patch signals that do not already exist in the client's state.

```javascript
response.headers.set('Content-Type', 'application/json')
response.headers.set('datastar-only-if-missing', 'true')
response.body = JSON.stringify({ foo: 'bar' })
```

--------------------------------

### Send Multiple SSE Events using JavaScript

Source: https://data-star.dev/guide/backend_requests

This JavaScript code snippet demonstrates sending multiple Server-Sent Events (SSE) using a `stream` object. It calls `patchElements` and `patchSignals` methods, passing HTML strings and an object for signals. This is typical for client-side or Node.js backend implementations using Datastar.

```javascript
stream.patchElements('<div id="question">...</div>');
stream.patchElements('<div id="instructions">...</div>');
stream.patchSignals({'answer': '...', 'prize': '...'});
```

--------------------------------

### Wrap Expression in View Transition API with data-on-load modifier

Source: https://data-star.dev/reference

Wraps the expression in `document.startViewTransition()` when the View Transition API is available. This modifier, `__viewtransition`, allows for smooth visual transitions between DOM states triggered by the event listener.

```html
<div data-on-load__viewtransition="doSomethingAnimated()"></div>
```

--------------------------------

### Create Loading Indicators with DataStar

Source: https://data-star.dev/reference/attributes

The `data-indicator` attribute creates a signal that is true while a fetch request is in flight and false otherwise, useful for loading indicators. The signal name can be specified in the key or value. It can control attributes like `disabled` or show/hide elements.

```html
<button data-on-click="@get('/endpoint')"
        data-indicator-fetching
></button>
```

```html
<button data-on-click="@get('/endpoint')"
        data-indicator-fetching
        data-attr-disabled="$fetching"
></button>
<div data-show="$fetching">Loading...</div>
```

```html
<button data-indicator="fetching"></button>
```

```html
<div data-indicator-fetching data-on-load="@get('/endpoint')"></div>
```

--------------------------------

### Implement Dynamic Polling Interval in Kotlin

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Kotlin snippet demonstrates dynamic polling intervals for Server-Sent Events. It calculates the interval duration based on the current second, applying a 1-second interval during the last 10 seconds of each minute. It utilizes `ServerSentEventGenerator` for creating SSE responses.

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

### Datastar Attribute Casing for Signals and Events

Source: https://data-star.dev/reference

Explains how Datastar handles casing for data attributes. Keys for signal-defining attributes (e.g., `data-signals-*`) are converted to camelCase, while other attribute keys default to kebab-case. The `__case` modifier allows explicit conversion between casing styles.

```html
<div data-signals-my-signal></div>
<div data-class-text-blue-700></div>
<div data-on-widget-loaded__case.camel></div>
```

--------------------------------

### Enable History Support with data-query-string__history

Source: https://data-star.dev/reference/attributes

The `__history` modifier for `data-query-string` enables browser history support, adding a new history entry each time a matching signal changes. This allows users to navigate back and forth through application states using browser navigation.

```html
<div data-query-string__filter__history></div>
```

--------------------------------

### HTML with DataStar Event Bubbling

Source: https://data-star.dev/examples/event_bubbling

This HTML structure sets up a demo for event bubbling using DataStar. A parent div with `data-on-click` listens for click events, and child buttons have `data-id` attributes. The clicked button's ID updates a span element. The `pointer-events: none;` style on the button container is crucial.

```html
1<div id="demo">
 2    Key pressed: <span data-text="$key"></span>
 3    <div id="button-container" data-on-click="$key = evt.target.dataset.id">
 4        <button data-id="KEY ELSE" class="gray">KEY<br/>ELSE</button>
 5        <button data-id="CM">CM</button>
 6        <button data-id="OM">OM</button>
 7        <button data-id="FETCH">FETCH</button>
 8        <button data-id="SET">SET</button>
 9        <button data-id="EXEC">EXEC</button>
10        <button data-id="TEST ALARM" class="gray">TEST<br/>ALARM</button>
11        <button data-id="3">3</button>
12        <button data-id="2">2</button>
13        <button data-id="1">1</button>
14        <button data-id="ENTER">ENTER</button>
15        <button data-id="CLEAR">CLEAR</button>
16    </div>
17</div>
18
19<style>
20    #button-container {
21        pointer-events: none;
22    }
23</style>
```

--------------------------------

### Implement Dynamic Polling Interval in Clojure

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Clojure snippet demonstrates how to dynamically set the polling interval for Server-Sent Events. It calculates the duration based on the current second of the minute, defaulting to 5 seconds but switching to 1 second for the last 10 seconds. It requires the starfederation.datastar.clojure.api and http-kit libraries.

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
26                   current-time]))))
27
28         (d*/close-sse! sse))}))
```

--------------------------------

### Send Multiple SSE Events (Patch Elements and Signals)

Source: https://data-star.dev/guide/backend_requests

Demonstrates sending multiple Server-Sent Events (SSE) in a single response. This includes patching HTML elements and sending signal data, showcased in a concise DSL format. It's useful for updating multiple parts of the UI or sending complex state changes.

```dsl
(d*/patch-elements! sse "<div id=\"question\">...</div>")
(d*/patch-elements! sse "<div id=\"instructions\">...</div>")
(d*/patch-signals! sse "{answer: '...', prize: '...'}")
```

--------------------------------

### data-on Keydown Listener for Enter or Ctrl+L

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

Combines conditions to create a global keydown listener that triggers an alert for either the 'Enter' key press or the 'Ctrl+L' key combination. This demonstrates logical OR operations within event expressions.

```html
<div data-on-keydown__window="(evt.key === 'Enter' || (evt.ctrlKey && evt.key === 'l')) && alert('Key pressed')"></div>
```

--------------------------------

### Default Request Cancellation in Data-Star

Source: https://data-star.dev/reference/actions

Demonstrates Data-Star's default behavior where initiating a new fetch request on an element automatically cancels any existing request on the same element. This is useful for preventing race conditions when an element triggers multiple rapid requests.

```html
<button data-on-click="@get('/slow-endpoint')">Load Data</button>
```

--------------------------------

### Send Multiple SSE Events using C#

Source: https://data-star.dev/guide/backend_requests

This C# code snippet illustrates how to send multiple Server-Sent Events (SSE) using the Datastar service. It includes calls to `PatchElementsAsync` for updating HTML elements and `PatchSignalsAsync` for sending structured data. This is part of Datastar's .NET integration.

```csharp
datastarService.PatchElementsAsync(@"<div id=\"question\">...</div>");
datastarService.PatchElementsAsync(@"<div id=\"instructions\">...</div>");
datastarService.PatchSignalsAsync(new { answer = "...", prize = "..." } );
```

--------------------------------

### Send Multiple SSE Events using PHP

Source: https://data-star.dev/guide/backend_requests

This PHP code snippet demonstrates sending multiple Server-Sent Events (SSE) using the `$sse` object. It includes calls to `patchElements` and `patchSignals` methods. The `patchSignals` method accepts an associative array, which is then serialized.

```php
$sse->patchElements('<div id="question">...</div>');
$sse->patchElements('<div id="instructions">...</div>');
$sse->patchSignals(['answer' => '...', 'prize' => '...']);
```

--------------------------------

### Send Multiple SSE Events using Python

Source: https://data-star.dev/guide/backend_requests

This Python code snippet shows how to send multiple Server-Sent Events (SSE) using Datastar's `generator`. It utilizes `patchElements` and `patchSignals` methods with multiline strings for HTML and JSON data. This is a clean way to manage SSE updates in Python applications.

```python
generator.patchElements(
    elements = """<div id=\"question\">...</div>""",
)
generator.patchElements(
    elements = """<div id=\"instructions\">...</div>""",
)
generator.patchSignals(
    signals = """{"answer": "...", "prize": "..."}""",
)
```

--------------------------------

### Send SSE Events with Datastar in Kotlin

Source: https://data-star.dev/guide/backend_requests

This Kotlin snippet demonstrates using the ServerSentEventGenerator to patch elements and signals. It assumes a `response` object is available and utilizes string interpolation for data.

```kotlin
1val generator = ServerSentEventGenerator(response)
2
generator.patchElements(
3    elements = """<div id=\"question\">What do you put in a toaster?</div>""",
4)
5
generator.patchSignals(
6    signals = "{\"response\": \"\", \"answer\": \"bread\"}",
7)
```

--------------------------------

### Datastar Aliased Script Inclusion

Source: https://data-star.dev/reference/attributes

Shows how to include the Datastar aliased version using a script tag. This is useful for avoiding conflicts with legacy libraries when custom aliases are needed.

```html
<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.5/bundles/datastar-aliased.js"></script>
```

--------------------------------

### Set All Signals with @setAll()

Source: https://data-star.dev/reference/actions

The @setAll() action sets the value of multiple signals to a specified value. It can optionally filter signals using a regular expression for inclusion and exclusion. This is useful for bulk updates or resetting signal states.

```html
<!-- Sets the `foo` signal only -->
<div data-signals-foo="false">
    <button data-on-click="@setAll(true, {include: /^foo$/})"></button>
</div>

<!-- Sets all signals starting with `user.` -->
<div data-signals="{user: {name: '', nickname: ''}}">
    <button data-on-click="@setAll('johnny', {include: /^user\./})"></button>
</div>

<!-- Sets all signals except those ending with `_temp` -->
<div data-signals="{data: '', data_temp: '', info: '', info_temp: ''}">
    <button data-on-click="@setAll('reset', {include: /.*/, exclude: /_temp$/})"></button>
</div>
```

--------------------------------

### Implement Dynamic Polling Interval in C#

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This C# code snippet implements dynamic polling intervals for Server-Sent Events using StarFederation Datastar. It determines the interval duration based on the current second, switching to 1-second polling in the last 10 seconds of a minute. It utilizes `IDatastarService` for patching elements.

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

### Custom Web Component Implementation in JavaScript

Source: https://data-star.dev/guide/datastar_expressions

Defines a custom HTML element 'my-component' that extends HTMLElement. It observes the 'src' attribute and dispatches a 'mycustomevent' with the processed value whenever the attribute changes. This showcases the basic structure for creating reusable web components.

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

### Server-Sent JavaScript Execution in Data-Star

Source: https://data-star.dev/reference/actions

Shows how to send JavaScript code from the server to be executed in the browser using Data-Star. The 'Content-Type' is set to 'text/javascript', and the 'datastar-script-attributes' header allows for setting arbitrary attributes on the dynamically created script tag, such as 'type="module"'.

```javascript
response.headers.set('Content-Type', 'text/javascript')
response.headers.set('datastar-script-attributes', JSON.stringify({ type: 'module' }))
response.body = 'console.log("Hello from server!");'
```

--------------------------------

### Send SSE Events with Datastar in Python (Litestar)

Source: https://data-star.dev/guide/backend_requests

This Python snippet shows how to generate SSE events using Datastar within a Litestar application. It returns a DatastarResponse containing both patched elements and signals.

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

### PATCH Request API

Source: https://data-star.dev/reference/actions

Sends a PATCH request to the backend. Used for applying partial modifications to a resource.

```APIDOC
## PATCH /endpoint

### Description
Sends a `PATCH` request to the specified URI. This action is used for applying partial modifications to a resource on the server.

### Method
PATCH

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint to send the PATCH request to.
- **options** (object) - Optional - Configuration for the request. See GET request options for details.

### Request Example
```html
<button data-on-click="@patch('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response Body** - The server's response to the PATCH request.

#### Response Example
```json
{
  "status": "partially updated",
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```
```

--------------------------------

### Patch DOM Elements with SSE in Rust

Source: https://data-star.dev/guide

Implements Server-Sent Events (SSE) for patching DOM elements in Rust using the `datastar` crate. It utilizes the `stream!` macro and `PatchElements` to send updates with a one-second sleep.

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

### Implement Dynamic Polling Interval in PHP

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This PHP snippet shows how to dynamically set polling intervals for Server-Sent Events using the StarFederation Datastar library. It determines the interval duration based on the current second, switching to 1-second polling for the last 10 seconds of each minute. It uses `ServerSentEventGenerator` to patch elements.

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

### PUT Request API

Source: https://data-star.dev/reference/actions

Sends a PUT request to the backend. Typically used for updating existing resources.

```APIDOC
## PUT /endpoint

### Description
Sends a `PUT` request to the specified URI. This action is used for updating an existing resource on the server.

### Method
PUT

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint to send the PUT request to.
- **options** (object) - Optional - Configuration for the request. See GET request options for details.

### Request Example
```html
<button data-on-click="@put('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response Body** - The server's response to the PUT request.

#### Response Example
```json
{
  "status": "updated",
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```
```

--------------------------------

### Delay Event Listener with data-on-load

Source: https://data-star.dev/reference/attributes

The `data-on-load__delay` attribute delays the execution of an event listener. It accepts time values in milliseconds (e.g., `.500ms`) or seconds (e.g., `.1s`). This is useful for staggering UI updates or animations.

```html
<div data-on-load__delay.500ms="$count = 1"></div>
```

--------------------------------

### Sync Query String Params with data-query-string

Source: https://data-star.dev/reference

The `data-query-string` attribute synchronizes query string parameters with signal values. On page load, it updates signals from the URL, and on signal change, it updates the URL. Filtering can be applied using `include` and `exclude` regular expressions.

```html
<div data-query-string></div>
```

```html
<div data-query-string="{include: /foo/, exclude: /bar/}"></div>
```

--------------------------------

### data-on-interval: Execute Code at Intervals

Source: https://data-star.dev/reference/attributes

The `data-on-interval` directive executes an expression at a regular interval, with a default duration of one second. The `__duration` modifier can be used to set custom intervals, and other modifiers like `__leading` can control initial execution.

```html
<div data-on-interval="$count++"></div>
<div data-on-interval__duration.500ms="$count++"></div>
```

--------------------------------

### Datastar Expression with Signal and Element Variable

Source: https://data-star.dev/reference/attributes

Illustrates a Datastar expression used in a data attribute. It shows how to reference signals (prefixed with '$') and the element itself (available as the 'el' variable).

```html
<div id="bar" data-text="$foo + el.id"></div>
```

--------------------------------

### Datastar Expression: Multi-line Statements

Source: https://data-star.dev/guide/datastar_expressions

Shows how Datastar expressions can span multiple lines, emphasizing the requirement of using a semicolon to separate statements, unlike standard JavaScript where line breaks suffice.

```html
1<div data-signals-foo="1">
2    <button data-on-click="
3        $landingGearRetracted = true; 
4        @post('/launch')
5    ">
6        Force launch
7    </button>
8</div>
```

--------------------------------

### data-on Keydown Listener for Enter Key

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

Attaches a global keydown event listener to the window that triggers an alert only when the 'Enter' key is pressed. It utilizes the `evt.key` property for condition checking.

```html
<div data-on-keydown__window="evt.key === 'Enter' && alert('Key pressed')"></div>
```

--------------------------------

### Implement Dynamic Polling Interval in Node.js

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This Node.js snippet shows how to dynamically set polling intervals for Server-Sent Events. It calculates the duration based on the current second, switching to 1-second polling for the last 10 seconds of each minute. It utilizes `ServerSentEventGenerator` to stream SSE data.

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

### Control Signal Casing with data-ref Modifiers

Source: https://data-star.dev/reference/attributes

Modifiers for `data-ref` allow control over the casing of the generated signal name. Options include `.camel` (default), `.kebab`, `.snake`, and `.pascal`. This helps in maintaining consistent naming conventions.

```html
<div data-ref-my-signal__case.kebab></div>
```

--------------------------------

### Persist Signals with Custom Key and Session Storage

Source: https://data-star.dev/reference

The `data-persist` attribute supports custom storage keys by appending them after `data-persist-`. The `__session` modifier changes the storage target from local storage to session storage.

```html
<div data-persist-mykey></div>
```

```html
<!-- Persists signals using a custom key `mykey` in session storage -->
<div data-persist-mykey__session></div>
```

--------------------------------

### Execute Expressions on Animation Frames with data-on-raf

Source: https://data-star.dev/reference/attributes

The `data-on-raf` Pro attribute executes a given expression on every `requestAnimationFrame` event. This is useful for animations or continuous updates that need to synchronize with the browser's rendering cycle.

```html
<div data-on-raf="$count++"></div>
```

--------------------------------

### Patch Signals with data-signals Attribute

Source: https://data-star.dev/reference/attributes

The `data-signals` attribute allows for patching (adding, updating, or removing) signals. It supports direct signal assignment, nested signals using dot notation, and patching multiple signals using JavaScript object notation or JSON. Setting a signal to null removes it. Signal names are camel-cased in attributes and can include modifiers like `__case` for casing conversion and `__ifmissing` to set defaults.

```html
<div data-signals-foo="1"></div>
<div data-signals-foo.bar="1"></div>
<div data-signals="{foo: {bar: 1, baz: 2}}"></div>
<div data-signals="{foo: null}"></div>
<div data-signals-my-signal__case.kebab="1"
     data-signals-foo__ifmissing="1"
></div>
```

--------------------------------

### Create Element Reference Signal with data-ref

Source: https://data-star.dev/reference/attributes

The `data-ref` attribute creates a new signal that references the element it is placed on. The signal name can be specified in the attribute key or value. The signal can then be used to access element properties like `tagName`.

```html
<div data-ref-foo></div>
```

```html
<div data-ref="foo"></div>
```

```html
$foo is a reference to a <span data-text="$foo.tagName"></span> element
```

--------------------------------

### data-on: Attach Event Listeners

Source: https://data-star.dev/reference/attributes

The `data-on` directive attaches an event listener to an element. The specified expression executes when the event is triggered. An `evt` variable representing the event object is available. It supports custom events and modifiers for behavior customization.

```html
<button data-on-click="$foo = ''">Reset</button>
<div data-on-myevent="$foo = evt.detail"></div>
<button data-on-click__window__debounce.500ms.leading="$foo = ''"></button>
<div data-on-my-event__case.camel="$foo = ''"></div>
```

--------------------------------

### Filter Signals with data-on-signal-patch-filter

Source: https://data-star.dev/reference/attributes

The `data-on-signal-patch-filter` attribute filters which signals trigger `data-on-signal-patch`. It accepts an object with `include` and/or `exclude` properties that are regular expressions. This allows for precise control over which signal changes are reacted to.

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

### data-on-load: Execute Code on Element Load

Source: https://data-star.dev/reference/attributes

The `data-on-load` directive executes an expression when the element is loaded into the DOM. This can occur during page load, when an element is patched, or when the attribute itself is modified. It's useful for initializing component state or performing actions upon element readiness.

```html
<div data-on-load="$count = 1"></div>
```

--------------------------------

### Golang: Animated SVG Morph Sequence

Source: https://data-star.dev/examples/svg_morphing

This Go code implements an automatic sequence of SVG morphs to create an animation effect. It uses Datastar's SSE to update a single SVG circle's properties (radius and color) over time with short delays between each change, culminating in a reset.

```go
1svgMorphingRouter.Get("/animated_morph", func(w http.ResponseWriter, r *http.Request) {
2    sse := datastar.NewSSE(w, r)
3    
4    // First morph
5    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="30" fill="red" /></svg>`)
6    time.Sleep(500 * time.Millisecond)
7    
8    // Second morph
9    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="45" fill="orange" /></svg>`)
10    time.Sleep(500 * time.Millisecond)
11    
12    // Third morph
13    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="60" fill="yellow" /></svg>`)
14    time.Sleep(500 * time.Millisecond)
15    
16    // Reset
17    sse.PatchElements(`<svg id="animated-demo"><circle cx="50" cy="50" r="20" fill="green" /></svg>`)
18})
```

--------------------------------

### data-on Keydown Listener for Ctrl+L Combination

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

Implements a global keydown event listener to detect when the 'Ctrl' key is held down along with the 'L' key, triggering an alert. It checks both `evt.ctrlKey` and `evt.key` properties.

```html
<div data-on-keydown__window="evt.ctrlKey && evt.key === 'l' && alert('Key pressed')"></div>
```

--------------------------------

### Patch DOM Elements and Signals in JavaScript

Source: https://data-star.dev/guide/backend_requests

This JavaScript code uses the `stream.patchElements` and `stream.patchSignals` functions to update the DOM and frontend signals. `patchElements` replaces an existing DOM element, while `patchSignals` updates signal values. Ensure the target element exists in the DOM before patching.

```javascript
stream.patchElements(`<div id="question">What do you put in a toaster?</div>`);
stream.patchSignals({'response':  '', 'answer': 'bread'});
```

--------------------------------

### Debounce Event Listener with data-on-signal-patch

Source: https://data-star.dev/reference/attributes

The `data-on-signal-patch__debounce` attribute debounces an event listener. It accepts time values like `.500ms` or `.1s`, and can be configured with `.leading` or `.notrail` options to control execution timing. This prevents rapid, repeated function calls.

```html
<div data-on-signal-patch__debounce.500ms="doSomething()"></div>
```

--------------------------------

### JavaScript Integration: Asynchronous Function with Custom Event

Source: https://data-star.dev/guide/datastar_expressions

Illustrates how to handle asynchronous JavaScript functions within Datastar expressions. Since Datastar does not await asynchronous calls, the function dispatches a custom event (`mycustomevent`) with the result in `evt.detail.value` to update the `$result` signal.

```html
1<div data-signals-result>
2    <input data-bind-foo 
3           data-on-input="myfunction(el, $foo)"
4           data-on-mycustomevent__window="$result = evt.detail.value"
5    >
6    <span data-text="$result"></span>
7</div>
```

--------------------------------

### Set View Transition Name with data-view-transition

Source: https://data-star.dev/reference/attributes

The `data-view-transition` attribute explicitly sets the `view-transition-name` CSS style for an element, enabling it to participate in smooth page transitions. Page-level transitions are handled automatically if the View Transition API is supported.

```html
<div data-view-transition="$foo"></div>
```

--------------------------------

### Throttle Event Listener with data-on-signal-patch

Source: https://data-star.dev/reference/attributes

The `data-on-signal-patch__throttle` attribute throttles an event listener. Similar to debounce, it accepts time values and can be configured with `.noleading` or `.trail` options. Throttling ensures a function is called at most once within a specified time interval.

```html
<div data-on-signal-patch__throttle.500ms="doSomething()"></div>
```

--------------------------------

### Access Signals Without Subscription with @peek()

Source: https://data-star.dev/reference/actions

The @peek() action allows accessing signal values within DataStar expressions without subscribing to their changes. This prevents re-evaluation of the expression when the peeked signal changes, optimizing performance. It takes a callable function as an argument that returns the signal value.

```html
<div data-text="$foo + @peek(() => $bar)"></div>
```

--------------------------------

### DELETE Request API

Source: https://data-star.dev/reference/actions

Sends a DELETE request to the backend. Used for removing a resource.

```APIDOC
## DELETE /endpoint

### Description
Sends a `DELETE` request to the specified URI. This action is used for removing a resource from the server.

### Method
DELETE

### Endpoint
`/endpoint`

### Parameters
#### Query Parameters
- **uri** (string) - Required - The endpoint to send the DELETE request to.
- **options** (object) - Optional - Configuration for the request. See GET request options for details.

### Request Example
```html
<button data-on-click="@delete('/endpoint')"></button>
```

### Response
#### Success Response (200)
- **Response Body** - The server's response to the DELETE request.

#### Response Example
```json
{
  "status": "deleted",
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```
```

--------------------------------

### DataStar SSE for Script Execution with Append Mode

Source: https://data-star.dev/guide/datastar_expressions

Demonstrates a DataStar Server-Sent Event (SSE) that appends a script tag to the body of the HTML document. This method is useful for executing scripts that need to be injected into a specific location, such as the document body.

```html
event: datastar-patch-elements
data: mode append
data: selector body
data: elements <script>alert('This mission is too important for me to allow you to jeopardize it.')</script>
```

--------------------------------

### Filter Query String Sync with data-query-string

Source: https://data-star.dev/reference/attributes

The `data-query-string` attribute can filter signals to be synced using `include` and `exclude` regular expressions. This allows control over which parameters are reflected in the URL.

```html
<div data-query-string="{include: /foo/, exclude: /bar/}"></div>
```

--------------------------------

### React to Signal Patches with data-on-signal-patch

Source: https://data-star.dev/reference/attributes

The `data-on-signal-patch` attribute triggers an expression when one or more signals are patched. The `patch` variable, containing signal patch details, is available within the expression. This is useful for tracking data changes and updating computed values.

```html
<div data-on-signal-patch="console.log('A signal changed!')"></div>
```

```html
<div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
```

--------------------------------

### Frontend Interval Polling with Data Attributes

Source: https://data-star.dev/how_tos/poll_the_backend_at_regular_intervals

This snippet demonstrates using HTML data attributes to configure client-side polling. The `data-on-interval__duration` attribute specifies the polling interval in seconds and the endpoint to fetch data from. The content within the div is updated with the current time, likely to indicate the last poll time or trigger updates.

```html
<div id="time"
     data-on-interval__duration.${duration}s="@get('/endpoint)')"
>
  ${currentTime.toISOString()}
</div>
```

--------------------------------

### Golang: Random SVG Shape Transformation

Source: https://data-star.dev/examples/svg_morphing

This Go code utilizes Datastar's SSE to morph an SVG element into a different shape. Each time the endpoint is hit, it selects a random shape from a predefined list (`svgShapes`) and updates the SVG element with the ID 'shape-demo'.

```go
1svgMorphingRouter.Get("/shape_transform", func(w http.ResponseWriter, r *http.Request) {
2    sse := datastar.NewSSE(w, r)
3    shape := svgShapes[rand.N(len(svgShapes))]
4    sse.PatchElements(fmt.Sprintf(`<svg id="shape-demo">%s</svg>`, shape))
5})
```

--------------------------------

### Two-Way Data Binding with data-bind

Source: https://data-star.dev/reference

The `data-bind` attribute creates or uses an existing signal for two-way data binding with an element's value. It supports various input elements and preserves signal types, including arrays for multiple selections. Modifiers like `__case` can alter signal name casing.

```html
<input data-bind-foo />
```

```html
<input data-bind="foo" />
```

```html
<input data-bind-foo value="bar" />
```

```html
<div data-signals-foo="baz">
    <input data-bind-foo value="bar" />
</div>
```

```html
<div data-signals-foo="0">
    <select data-bind-foo>
        <option value="10">10</option>
    </select>
</div>
```

```html
<div data-signals-foo="[]">
    <input data-bind-foo type="checkbox" value="bar" />
    <input data-bind-foo type="checkbox" value="baz" />
</div>
```

```html
<input data-bind-my-signal__case.kebab />
```

--------------------------------

### Animated Scrolling with data-scroll-into-view__smooth

Source: https://data-star.dev/reference/attributes

The `__smooth` modifier for `data-scroll-into-view` enables animated scrolling to the element, providing a smoother user experience compared to an instant jump.

```html
<div data-scroll-into-view__smooth></div>
```

--------------------------------

### Trigger Expression on Element Resize with data-on-resize

Source: https://data-star.dev/reference/attributes

The `data-on-resize` attribute executes an expression whenever an element's dimensions change. This is useful for responsive designs or when an element's size is dynamically altered.

```html
<div data-on-resize="$count++"></div>
```

--------------------------------

### Execute Effects with DataStar

Source: https://data-star.dev/reference/attributes

The `data-effect` attribute executes an expression on page load and whenever any signals within that expression change. This is ideal for performing side effects like updating other signals, making backend requests, or manipulating the DOM.

```html
<div data-effect="$foo = $bar + $baz"></div>
```

--------------------------------

### data-on Keydown Listener with PreventDefault

Source: https://data-star.dev/how_tos/bind_keydown_events_to_specific_keys

Shows how to prevent the default behavior of a keydown event, such as form submission on 'Enter' press, while still executing custom JavaScript. It uses `evt.preventDefault()` within the expression.

```html
<div data-on-keydown__window="evt.key === 'Enter' && (evt.preventDefault(), alert('Key pressed'))"></div>
```

--------------------------------

### Persist Signals in Local Storage with data-persist

Source: https://data-star.dev/reference

The `data-persist` attribute saves signal values to local storage, preserving them across page loads. It can be configured with `include` and `exclude` regular expressions to filter which signals are persisted.

```html
<div data-persist></div>
```

```html
<div data-persist="{include: /foo/, exclude: /bar/}"></div>
```

--------------------------------

### Persist Signals in Local Storage with data-persist

Source: https://data-star.dev/reference/attributes

The `data-persist` attribute saves signal values to local storage, preserving them across page loads. By default, signals are stored under the key 'datastar'.

```html
<div data-persist></div>
```

--------------------------------

### Set Inline CSS Styles with data-style Attribute

Source: https://data-star.dev/reference

The `data-style` attribute dynamically sets inline CSS styles based on expressions. It supports single style properties or multiple properties defined as key-value pairs. Style properties can be in camelCase or kebab-case. Falsy values (empty string, null, undefined, false) restore original inline styles or remove the property.

```html
<div data-style-background-color="$usingRed ? 'red' : 'blue'"></div>
<div data-style-display="$hiding && 'none'"></div>
```

```html
<div data-style="{
   display: $hiding ? 'none' : 'flex',
   flexDirection: 'column',
   color: $usingRed ? 'red' : 'green'
}"></div>
```

```html
<!-- When $x is false, color remains red from inline style -->
<div style="color: red;" data-style-color="$x && 'green'"></div>

<!-- When $hiding is true, display becomes none; when false, reverts to flex from inline style -->
<div style="display: flex;" data-style-display="$hiding && 'none'"></div>
```

--------------------------------

### Apply Modifiers with data-signals Attribute

Source: https://data-star.dev/reference

Modifiers can be used with the `data-signals` attribute to alter signal patching behavior. `__case` converts signal name casing, and `__ifmissing` only patches if the signal key doesn't exist, useful for setting defaults.

```html
<div data-signals-my-signal__case.kebab="1"
     data-signals-foo__ifmissing="1"
></div>
```

--------------------------------

### Toggle Boolean Signals with @toggleAll()

Source: https://data-star.dev/reference/actions

The @toggleAll() action toggles the boolean value of matching signals. Similar to @setAll(), it supports filtering signals by regular expressions for inclusion and exclusion. This is ideal for toggling states like 'active', 'enabled', or 'visible'.

```html
<!-- Toggles the `foo` signal only -->
<div data-signals-foo="false">
    <button data-on-click="@toggleAll({include: /^foo$/})"></button>
</div>

<!-- Toggles all signals starting with `is` -->
<div data-signals="{isOpen: false, isActive: true, isEnabled: false}">
    <button data-on-click="@toggleAll({include: /^is/})"></button>
</div>

<!-- Toggles signals starting with `settings.` -->
<div data-signals="{settings: {darkMode: false, autoSave: true}}">
    <button data-on-click="@toggleAll({include: /^settings\./})"></button>
</div>
```

--------------------------------

### Animate Attributes with data-animate

Source: https://data-star.dev/reference/attributes

The `data-animate` Pro attribute enables animation of element attributes over time. These animations are reactive, meaning they update automatically whenever the signals used within the animation expression change.

--------------------------------

### Set HTML Attribute with DataStar data-attr

Source: https://data-star.dev/reference/attributes

The `data-attr` attribute in DataStar allows you to dynamically set and maintain the value of any HTML attribute. It can target a single attribute or multiple attributes using an object literal for key-value pairs.

```html
<div data-attr-title="$foo"></div>
```

```html
<div data-attr="{title: $foo, disabled: $bar}"></div>
```

--------------------------------

### Modify Computed Signal Naming with DataStar Modifiers

Source: https://data-star.dev/reference

Modifiers like `__case` can be applied to `data-computed` to change the casing of the computed signal name. Supported cases include `.camel` (default), `.kebab`, `.snake`, and `.pascal`. This allows for flexible naming conventions based on project requirements.

```html
<div data-computed-my-signal__case.kebab="$bar + $baz"></div>
```

--------------------------------

### HTML Structure for Namespaced SVG

Source: https://data-star.dev/examples/svg_morphing

This HTML structure demonstrates how to properly wrap an inner SVG element within an outer `<svg>` tag. This is crucial for ensuring the inner SVG elements are created under the correct namespace, which is necessary for SVG morphing in Datastar.

```html
1<svg>
2    <svg id="target">
3        <circle cx="50" cy="100" r="50" fill="red" />
4    </svg>
5    <circle cx="150" cy="100" r="50" fill="red" />
6</svg>
```

--------------------------------

### Compact JSON Output with DataStar Data-Json-Signals Modifiers

Source: https://data-star.dev/reference

The `__terse` modifier for `data-json-signals` outputs JSON in a compact format without extra whitespace. This is beneficial for displaying filtered data inline or when space is limited. It enhances readability for concise data representations.

```html
<!-- Display filtered signals in a compact format -->
<pre data-json-signals__terse="{include: /counter/}"></pre>
```

--------------------------------

### Display Reactive JSON Signals with DataStar

Source: https://data-star.dev/reference/attributes

The `data-json-signals` attribute sets an element's text content to a reactive, JSON stringified version of signals, useful for debugging. Optional filters can include or exclude signals using regular expressions. The `__terse` modifier outputs compact JSON.

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

```html
<!-- Display filtered signals in a compact format -->
<pre data-json-signals__terse="{include: /counter/}"></pre>
```

--------------------------------

### Set Inline CSS Styles with data-style Attribute

Source: https://data-star.dev/reference/attributes

The `data-style` attribute dynamically sets inline CSS styles on an element based on expressions. It supports setting individual style properties (e.g., `background-color`) or multiple properties using a JavaScript object. Style property names can be camelCase or kebab-case. Falsy values (empty string, null, undefined, false) can be used to remove or revert styles, allowing for conditional styling with the logical AND operator.

```html
<div data-style-background-color="$usingRed ? 'red' : 'blue'"></div>
<div data-style-display="$hiding && 'none'"></div>
<div data-style="{
   display: $hiding ? 'none' : 'flex',
   flexDirection: 'column',
   color: $usingRed ? 'red' : 'green'
}"></div>
<!-- When $x is false, color remains red from inline style -->
<div style="color: red;" data-style-color="$x && 'green'"></div>
<!-- When $hiding is true, display becomes none; when false, reverts to flex from inline style -->
<div style="display: flex;" data-style-display="$hiding && 'none'"></div>
```

--------------------------------

### Two-Way Data Binding with DataStar data-bind

Source: https://data-star.dev/reference/attributes

The `data-bind` attribute facilitates two-way data binding between an HTML element's value and a DataStar signal. It supports various input elements and automatically updates the signal when the element changes, and vice-versa. Signal type preservation and array binding for checkboxes are supported.

```html
<input data-bind-foo />
```

```html
<input data-bind="foo" />
```

```html
<input data-bind-foo value="bar" />
```

```html
<div data-signals-foo="baz">
    <input data-bind-foo value="bar" />
</div>
```

```html
<div data-signals-foo="0">
    <select data-bind-foo>
        <option value="10">10</option>
    </select>
</div>
```

```html
<div data-signals-foo="[]">
    <input data-bind-foo type="checkbox" value="bar" />
    <input data-bind-foo type="checkbox" value="baz" />
</div>
```

--------------------------------

### Golang: Change SVG Circle Color

Source: https://data-star.dev/examples/svg_morphing

This Go code snippet shows how to change the fill color of an SVG circle using Datastar's SSE (Server-Sent Events) for real-time updates. It selects a random color from a predefined list and applies it to the SVG element with the ID 'circle-demo'.

```go
1svgMorphingRouter.Get("/circle_color", func(w http.ResponseWriter, r *http.Request) {
2    sse := datastar.NewSSE(w, r)
3    color := svgColors[rand.N(len(svgColors))]
4    sse.PatchElements(fmt.Sprintf(`<svg id="circle-demo"><circle cx="50" cy="50" r="40" fill="%s" /></svg>`, color))
5})
```

--------------------------------

### Display Reactive JSON Signals with DataStar Data-Json-Signals

Source: https://data-star.dev/reference

The `data-json-signals` attribute displays a reactive, JSON stringified version of signals within an element's text content. It's useful for debugging and troubleshooting. Optional filter objects can be provided to include or exclude specific signals using regular expressions.

```html
<!-- Display all signals -->
<pre data-json-signals></pre>

<!-- Only show signals that include "user" in their path -->
<pre data-json-signals="{include: /user/}"></pre>

<!-- Show all signals except those ending with "temp" -->
<pre data-json-signals="{exclude: /temp$/}"></pre>

<!-- Combine include and exclude filters -->
<pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>
```

--------------------------------

### DataStar data-bind Case Modifier

Source: https://data-star.dev/reference/attributes

The `__case` modifier for `data-bind` allows you to control the casing of the signal name during data binding. Supported formats include camel, kebab, snake, and pascal case.

```html
<input data-bind-my-signal__case.kebab />
```

--------------------------------

### Conditionally Show/Hide Element with data-show

Source: https://data-star.dev/reference/attributes

The `data-show` attribute conditionally shows or hides an element based on a boolean expression. For elements that should be hidden initially before Datastar processes them, a `display: none` style can be applied to prevent flickering.

```html
<div data-show="$foo"></div>
```

```html
<div data-show="$foo" style="display: none"></div>
```

--------------------------------

### Replace URL Without Reloading with data-replace-url

Source: https://data-star.dev/reference/attributes

The `data-replace-url` attribute updates the browser's URL without causing a page reload. The attribute's value, which can be an evaluated expression, specifies the new URL.

```html
<div data-replace-url="`/page${page}`"></div>
```

--------------------------------

### Filter Persisted Signals with data-persist

Source: https://data-star.dev/reference/attributes

The `data-persist` attribute can filter which signals are persisted using `include` and `exclude` regular expressions within its value. This allows for selective storage of signal data.

```html
<div data-persist="{include: /foo/, exclude: /bar/}"></div>
```

--------------------------------

### data-on-intersect: Trigger on Viewport Intersection

Source: https://data-star.dev/reference/attributes

The `data-on-intersect` directive runs an expression when an element intersects with the viewport. It supports various modifiers to control the intersection threshold, timing, and behavior, such as triggering only once or when a certain percentage is visible.

```html
<div data-on-intersect="$intersected = true"></div>
<div data-on-intersect__once__full="$fullyIntersected = true"></div>
```

--------------------------------

### DataStar data-class Case Modifier

Source: https://data-star.dev/reference/attributes

The `__case` modifier for `data-class` enables you to specify the casing for the CSS class name. Options include camel, kebab, snake, and pascal case.

```html
<div data-class-my-class__case.camel="$foo"></div>
```

--------------------------------

### Golang: Change SVG Circle Radius

Source: https://data-star.dev/examples/svg_morphing

This Go code demonstrates morphing the radius of an SVG circle using Datastar. When triggered, it generates a random radius between 15 and 60 and updates the SVG element with the ID 'size-demo', effectively changing the circle's size.

```go
1svgMorphingRouter.Get("/circle_size", func(w http.ResponseWriter, r *http.Request) {
2    sse := datastar.NewSSE(w, r)
3    radius := 15 + rand.N(45) // Random radius between 15-60
4    sse.PatchElements(fmt.Sprintf(`<svg id="size-demo"><circle cx="50" cy="50" r="%d" fill="green" /></svg>`, radius))
5})
```

--------------------------------

### Bind Element Text Content with data-text Attribute

Source: https://data-star.dev/reference/attributes

The `data-text` attribute binds the text content of an element to the evaluation of a given expression. This is useful for displaying dynamic text values within your HTML.

```html
<div data-text="$foo"></div>
```

--------------------------------

### Create Computed Signals with DataStar

Source: https://data-star.dev/reference/attributes

The `data-computed` attribute creates a read-only signal whose value is automatically updated when its dependencies change. It's useful for memoizing expressions and can be used in other expressions. Computed signals should not perform actions.

```html
<div data-computed-foo="$bar + $baz"></div>
<div data-text="$foo"></div>
```

```html
<div data-computed-my-signal__case.kebab="$bar + $baz"></div>
```

--------------------------------

### Conditional Class Toggling with DataStar data-class

Source: https://data-star.dev/reference/attributes

The `data-class` attribute dynamically adds or removes CSS classes from an element based on the evaluation of an expression. It supports toggling single classes or multiple classes simultaneously using an object literal.

```html
<div data-class-hidden="$foo"></div>
```

```html
<div data-class="{hidden: $foo, 'font-bold': $bar}"></div>
```

--------------------------------

### Scroll Element Into View with data-scroll-into-view

Source: https://data-star.dev/reference/attributes

The `data-scroll-into-view` attribute automatically scrolls the element into the viewport. This is particularly useful when dynamically adding content to the DOM and needing to focus the user's attention on it.

```html
<div data-scroll-into-view></div>
```

--------------------------------

### Custom Input Validation with data-custom-validity

Source: https://data-star.dev/reference/attributes

The `data-custom-validity` Pro attribute allows for custom validation logic on form elements. An expression is evaluated, and if it returns a non-empty string, the input is considered invalid with that string as the validation message. An empty string signifies a valid input.

```html
<form>
    <input data-bind-foo name="foo" />
    <input data-bind-bar name="bar"
           data-custom-validity="$foo === $bar ? '' : 'Values must be the same.'"
    />
    <button>Submit form</button>
</form>
```

--------------------------------

### Ignore Morphing with DataStar

Source: https://data-star.dev/reference/attributes

The `data-ignore-morph` attribute tells the `PatchElements` watcher to skip processing an element and its children during element morphing. To remove this attribute, patch the element with the attribute absent.

```html
<div data-ignore-morph>
    This element will not be morphed.
</div>
```

--------------------------------

### Conditional Class Toggling with data-class

Source: https://data-star.dev/reference

The `data-class` attribute adds or removes CSS classes from an element based on an expression's truthiness. It supports toggling single classes or multiple classes using an object of class names and their corresponding expressions. The `__case` modifier can adjust class name casing.

```html
<div data-class-hidden="$foo"></div>
```

```html
<div data-class="{hidden: $foo, 'font-bold': $bar}"></div>
```

```html
<div data-class-my-class__case.camel="$foo"></div>
```

--------------------------------

### Create Computed Signals with DataStar

Source: https://data-star.dev/reference

The `data-computed` attribute creates a read-only signal computed from an expression involving other signals. Its value updates automatically when dependent signals change. Useful for memoizing expressions and their values can be used in other expressions. Avoid using computed signals for performing actions.

```html
<div data-computed-foo="$bar + $baz"></div>
<div data-text="$foo"></div>
```

--------------------------------

### HTML Form with data-custom-validity for Validation

Source: https://data-star.dev/errors/runtime/custom_validity_invalid_expression

This HTML snippet demonstrates the correct usage of the `data-custom-validity` attribute. It binds two input fields and uses an expression to compare their values. If the values differ, a custom error message is displayed; otherwise, the input is considered valid.

```html
<form>
    <input data-bind-foo name="foo" />
    <input data-bind-bar name="bar" 
        data-custom-validity="$foo === $bar ? '' : 'Field values must be the same.'"
    />
    <button>
        Submit form
    </button>
</form>
```

--------------------------------

### Preserve Attribute Value with data-preserve-attr

Source: https://data-star.dev/reference/attributes

The `data-preserve-attr` attribute preserves the value of specified attributes when morphing DOM elements. Multiple attributes can be preserved by separating them with a space. This is useful for maintaining state on elements like `open` for `<details>` tags.

```html
<details open data-preserve-attr="open">
    <summary>Title</summary>
    Content
</details>
```

```html
<details open class="foo" data-preserve-attr="open class">
    <summary>Title</summary>
    Content
</details>
```

--------------------------------

### Ignore Elements with DataStar

Source: https://data-star.dev/reference/attributes

The `data-ignore` attribute prevents DataStar from processing an element and its descendants. This is useful for avoiding conflicts with third-party libraries or unescaped user input. The `__self` modifier can be used to ignore only the element itself.

```html
<div data-ignore data-show-thirdpartylib="">
    <div>
        Datastar will not process this element.
    </div>
</div>
```