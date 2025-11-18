### daisyUI 5 CDN Installation Example

Source: https://context7_llms

This HTML snippet demonstrates how to include daisyUI 5 and Tailwind CSS browser version directly from CDNs. It's a quick way to get started without a local installation, suitable for quick prototypes or simple projects.

```html
<link
  href="https://cdn.jsdelivr.net/npm/daisyui@5"
  rel="stylesheet"
  type="text/css"
/>
<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
```

---

### Basic daisyUI CSS Plugin Setup

Source: https://context7_llms

This CSS snippet shows the most basic way to include the daisyUI plugin without any custom configuration. It assumes daisyUI is installed as a CSS plugin.

```css
@plugin "daisyui";
```

---

### Install daisyUI with Bun

Source: https://daisyui.com/docs/install

Installs the latest version of daisyUI as a development dependency using Bun.

```bash
bun add -D daisyui@latest
```

---

### Install daisyUI with Deno

Source: https://daisyui.com/docs/install

Installs the latest version of daisyUI using Deno's package management.

```bash
deno i -D npm:daisyui@latest
```

---

### daisyUI Dock Component HTML Examples

Source: https://daisyui.com/components/dock

Demonstrates the basic structure and usage of the daisyUI Dock component. Includes examples for default, active states, and different size variations (extra small to extra large).

```html
<div class="dock">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

```html
<div class="dock dock-xs">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

```html
<div class="dock dock-sm">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

```html
<div class="dock dock-md">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

```html
<div class="dock dock-lg">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

```html
<div class="dock dock-xl">
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7"
      />
    </svg>
    Home
  </button>
  <button class="dock-tab dock-active">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z"
      />
    </svg>
    Inbox
  </button>
  <button class="dock-tab">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z"
      />
    </svg>
    Settings
  </button>
</div>
```

---

### Install daisyUI with Yarn

Source: https://daisyui.com/docs/install

Installs the latest version of daisyUI as a development dependency using Yarn.

```bash
yarn add -D daisyui@latest
```

---

### Cally Calendar Web Component Examples

Source: https://daisyui.com/components/calendar

Demonstrates how to integrate the Cally web component calendar with DaisyUI styles, including basic examples and a date picker with a dropdown. Cally works everywhere and can be imported from a CDN or installed via npm.

```html
<link
  rel="stylesheet"
  href="https://cdn.jsdelivr.net/npm/cally/dist/cally.css"
/>
<script type="module" src="https://cdn.jsdelivr.net/npm/cally"></script>

<cally></cally>
```

```html
<link
  rel="stylesheet"
  href="https://cdn.jsdelivr.net/npm/cally/dist/cally.css"
/>
<script type="module" src="https://cdn.jsdelivr.net/npm/cally"></script>

<div class="dropdown">
  <label tabindex="0" class="btn m-1">Pick a date</label>
  <div
    tabindex="0"
    class="dropdown-content card card-compact p-2 shadow bg-base-100"
  >
    <cally></cally>
  </div>
</div>
```

```html
// Import Cally web component from CDN // Or install as a dependency: npm i
cally // and import it in JS import "cally"; // Example usage in HTML: //
<cally></cally>
```

```html
// Import Cally web component from CDN // Or install as a dependency: npm i
cally // and import it in JS import "cally"; // Example usage in HTML for a
dropdown: //
<div class="dropdown">
  // <label tabindex="0" class="btn m-1">Pick a date</label> //
  <div
    tabindex="0"
    class="dropdown-content card card-compact p-2 shadow bg-base-100"
  >
    // <cally></cally> //
  </div>
  //
</div>
```

---

### DaisyUI Kbd Component Examples

Source: https://daisyui.com/components/kbd

This section showcases various ways to use the DaisyUI Kbd component. It includes examples for displaying single keys, different sizes of keys, keys within text, key combinations, function keys, a full keyboard layout, and arrow keys.

```html
<kbd>K</kbd>
```

```html
<kbd class="kbd-xs">Xsmall</kbd>
<kbd class="kbd-sm">Small</kbd>
<kbd class="kbd-md">Medium</kbd>
<kbd class="kbd-lg">Large</kbd>
<kbd class="kbd-xl">Xlarge</kbd>
```

```html
Press <kbd>F</kbd> to pay respects.
```

```html
<kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>del</kbd>
```

```html
<kbd>⌘</kbd>
<kbd>⌥</kbd>
<kbd>⇧</kbd>
<kbd>⌃</kbd>
```

```html
<kbd>q</kbd> <kbd>w</kbd> <kbd>e</kbd> <kbd>r</kbd> <kbd>t</kbd> <kbd>y</kbd>
<kbd>u</kbd> <kbd>i</kbd> <kbd>o</kbd> <kbd>p</kbd> <kbd>a</kbd> <kbd>s</kbd>
<kbd>d</kbd> <kbd>f</kbd> <kbd>g</kbd> <kbd>h</kbd> <kbd>j</kbd> <kbd>k</kbd>
<kbd>l</kbd> <kbd>z</kbd> <kbd>x</kbd> <kbd>c</kbd> <kbd>v</kbd> <kbd>b</kbd>
<kbd>n</kbd> <kbd>m</kbd> <kbd>/</kbd>
```

```html
<kbd>▲</kbd>

<kbd>◀︎</kbd> <kbd>▶︎</kbd>

<kbd>▼</kbd>
```

---

### DaisyUI Modal Component Examples (HTML)

Source: https://context7_llms

Shows different methods for implementing DaisyUI modals. Examples include using the HTML dialog element, a checkbox for toggling (legacy), and anchor links (legacy). Modals can be positioned using placement classes.

```html
<button onclick="my_modal.showModal()">Open modal</button>
<dialog id="my_modal" class="modal modal-bottom sm:modal-middle">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Press ESC key or click the button below to close</p>
  </div>
  <form method="dialog" class="modal-backdrop">
    <button>close</button>
  </form>
</dialog>
```

```html
<label for="my-modal" class="btn">Open modal</label>
<input type="checkbox" id="my-modal" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">You can wimdow content here!</p>
  </div>
  <label class="modal-backdrop" for="my-modal">Close</label>
</div>
```

```html
<a href="#my-modal" class="btn">Open modal</a>
<div class="modal" id="my-modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">You can wimdow content here!</p>
  </div>
</div>
```

---

### Pikaday Svelte Component Example

Source: https://daisyui.com/components/calendar

Provides an example of integrating Pikaday within a Svelte component. It includes installation instructions and shows how to manage the Pikaday instance lifecycle using Svelte's effect hook.

```bash
npm i pikaday
```

```svelte
<script lang="ts">
import { onMount, onDestroy } from 'svelte';
import Pikaday from 'pikaday';

let myDatepicker;
let pickerInstance;

onMount(() => {
  if (myDatepicker) {
    pickerInstance = new Pikaday({ field: myDatepicker });
  }
});

onDestroy(() => {
  if (pickerInstance) {
    pickerInstance.destroy();
  }
});
</script>

<input type="text" bind:this={myDatepicker} />
```

---

### Install DaisyUI using Package Managers

Source: https://daisyui.com/

Instructions for installing DaisyUI as a development dependency using popular Node.js package managers (npm, pnpm, yarn, bun) and Deno. Ensures the latest version is used.

```bash
npm i -D daisyui@latest
```

```bash
pnpm add -D daisyui@latest
```

```bash
yarn add -D daisyui@latest
```

```bash
bun add -D daisyui@latest
```

```bash
deno i -D npm:daisyui@latest
```

---

### Install daisyUI with PNPM

Source: https://daisyui.com/docs/install

Installs the latest version of daisyUI as a development dependency using PNPM.

```bash
pnpm add -D daisyui@latest
```

---

### Code Block Example - HTML

Source: https://daisyui.com/components/mockup-browser

Illustrates the basic structure for displaying code snippets using daisyUI's code component. This is essential for any documentation that includes code examples.

```html

```

// Your code here

```

```

---

### React Day Picker Component Example

Source: https://daisyui.com/components/calendar

Demonstrates the usage of the React Day Picker component for creating date pickers in React applications. It includes installation instructions and a basic example using `useState` to manage the selected date.

```bash
npm i react-day-picker
```

```typescript
import { useState } from "react";
import { DayPicker, DateRange } from "react-day-picker";
import "react-day-picker/dist/style.css";

export default function App() {
  const [date, setDate] = useState<DateRange | undefined>();

  return (
    <DayPicker
      mode="single"
      selected={date}
      onSelect={setDate}
    />
    {date ? date.toLocaleDateString() : "Pick a date"}
  );
}
```

---

### Install daisyUI with npm

Source: https://daisyui.com/

This command installs daisyUI as a development dependency in your project using npm. Ensure you have Node.js and npm installed.

```bash
npm i -D daisyui@latest
```

---

### DaisyUI Badge Component Examples (HTML)

Source: https://daisyui.com/components/badge

Demonstrates the basic structure and usage of DaisyUI badges, including different styles and sizes. These examples show how to apply classes to create various badge appearances.

```html
<!-- Default Badge -->
<div class="badge">Badge</div>

<!-- Outline Badge -->
<div class="badge badge-outline">Badge</div>

<!-- Dash Outline Badge -->
<div class="badge badge-dash">Badge</div>

<!-- Soft Badge -->
<div class="badge badge-soft">Badge</div>

<!-- Ghost Badge -->
<div class="badge badge-ghost">Badge</div>

<!-- Extra Small Badge -->
<div class="badge badge-xs">Badge</div>

<!-- Small Badge -->
<div class="badge badge-sm">Badge</div>

<!-- Medium Badge (default) -->
<div class="badge badge-md">Badge</div>

<!-- Large Badge -->
<div class="badge badge-lg">Badge</div>

<!-- Extra Large Badge -->
<div class="badge badge-xl">Badge</div>
```

---

### Pikaday Vue Component Example

Source: https://daisyui.com/components/calendar

Illustrates how to use Pikaday within a Vue.js application. It covers installation and demonstrates initializing the datepicker in the `mounted` lifecycle hook, referencing the input field via a ref.

```bash
npm i pikaday
```

```vue
<template>
  <input type="text" ref="myDatepicker" />
</template>

<script>
import Pikaday from "pikaday";

export default {
  mounted: function () {
    const picker = new Pikaday({ field: this.$refs.myDatepicker });
  },
  beforeDestroy: function () {
    // Ensure picker is destroyed to prevent memory leaks
    if (this.picker) {
      this.picker.destroy();
    }
  },
};
</script>
```

---

### Update Card component classes

Source: https://daisyui.com/docs/upgrade

This example demonstrates the renaming of the `card-bordered` class to `card-border` and the replacement of `card-compact` with `card-sm` for styling card elements.

```html
-
<div class="card card-bordered">
  -
  <div class="card-body">...</div>
  -
</div>
+
<div class="card card-border">
  +
  <div class="card-body">...</div>
  +
</div>
```

```html
-
<div class="card card-compact">
  -
  <div class="card-body">...</div>
  -
</div>
+
<div class="card card-sm">
  +
  <div class="card-body">...</div>
  +
</div>
```

---

### DaisyUI Badge as Indicator Examples (HTML)

Source: https://daisyui.com/components/indicator

Shows how to use badges as indicators on various elements. This includes examples for displaying notification counts on buttons, messages, and requests.

```html
<div class="indicator">
  <span class="indicator-item badge">12 inbox</span>
  <button class="btn">Inbox</button>
</div>
```

```html
<div class="indicator">
  <span class="indicator-item badge">8</span>
  <button class="btn btn-primary">Messages</button>
</div>
<div class="indicator">
  <span class="indicator-item badge">8</span>
  <button class="btn btn-accent">Notifications</button>
</div>
<div class="indicator">
  <span class="indicator-item badge">8</span>
  <button class="btn btn-secondary">Requests</button>
</div>
```

---

### Clock Countdown Example (HTML)

Source: https://daisyui.com/components/countdown

Illustrates how to implement a clock-style countdown using the DaisyUI Countdown component. This example shows the structure for displaying hours, minutes, and seconds.

```html
<div class="countdown">
  <span style="--value:10"></span>h <span style="--value:24"></span>m
  <span style="--value:59"></span>s
</div>
```

---

### Example daisyUI CSS Plugin Configuration

Source: https://context7_llms

This is an example of a comprehensive daisyUI CSS configuration. It enables all built-in themes, sets 'bumblebee' as the default and 'synthwave' as the prefers-dark theme. It also excludes 'rootscrollgutter' and 'checkbox', uses a 'daisy-' prefix for classes, and disables console logging.

```css
@plugin "daisyui" {
  themes: light, dark, cupcake, bumblebee --default, emerald, corporate, synthwave
      --prefersdark, retro, cyberpunk, valentine, halloween, garden, forest,
    aqua, lofi, pastel, fantasy, wireframe, black, luxury, dracula, cmyk,
    autumn, business, acid, lemonade, night, coffee, winter, dim, nord, sunset,
    caramellatte, abyss, silk;
  root: ":root";
  include: ;
  exclude: rootscrollgutter, checkbox;
  prefix: daisy-;
  logs: false;
}
```

---

### DaisyUI Indicator Placement: Middle-Start (HTML)

Source: https://daisyui.com/components/indicator

Demonstrates using `indicator-middle` and `indicator-start` to position an indicator vertically in the middle and horizontally at the start of a container.

```html
<div class="indicator indicator-middle indicator-start">
  <span class="indicator-item badge badge-secondary">New</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
```

---

### DaisyUI Menu Component Examples (HTML)

Source: https://context7_llms

Demonstrates the structure and variations of the DaisyUI menu component. It includes examples for vertical menus, horizontal menus, and dropdown menus. Modifiers for disabled, active, and focus states are also applicable.

```html
<ul class="menu">
  <li><button>Item</button></li>
</ul>
```

```html
<ul class="menu menu-horizontal">
  <li><button>Item</button></li>
</ul>
```

```html
<details class="menu-dropdown">
  <summary class="menu-dropdown-toggle">Item</summary>
  <ul class="menu-dropdown-content">
    <li><a>SubItem 1</a></li>
    <li><a>SubItem 2</a></li>
  </ul>
</details>
```

---

### Basic DaisyUI Alert Example

Source: https://daisyui.com/components/alert

Demonstrates a basic DaisyUI alert component. This is the foundational structure for displaying notifications.

```html
<div class="alert">12 unread messages. Tap to see.</div>
```

---

### Calendar Integration Examples (HTML)

Source: https://context7_llms

Provides HTML snippets for integrating DaisyUI styles with different calendar libraries like Cally, Pikaday, and React Day Picker.

```html
<!-- For Cally web component -->
<calendar-date class="cally">Content</calendar-date>

<!-- For Pikaday calendar input -->
<input type="text" class="input pika-single" />

<!-- For React Day Picker component -->
<DayPicker className="react-day-picker"></DayPicker>
```

---

### DaisyUI Image with Overlay Indicator Example (HTML)

Source: https://daisyui.com/components/indicator

Provides an example of placing an indicator, like a text label, in the center of an image. This is useful for displaying status or promotional messages.

```html
<div class="indicator">
  <span class="indicator-item">Only available for Pro users</span>
  <img
    src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
    class="w-full rounded-xl"
  />
</div>
```

---

### DaisyUI Avatar with Indicator Example (HTML)

Source: https://daisyui.com/components/indicator

Illustrates how to place an indicator on a user avatar. This is useful for showing online status or other badges.

```html
<div class="indicator">
  <span
    class="indicator-item badge indicator-bottom indicator-end badge-primary"
  ></span>
  <div class="avatar online">
    <div class="w-16 rounded-full">
      <img
        src="https://img.daisyui.com/images/profile/demo/batperson@192.webp"
      />
    </div>
  </div>
</div>
```

---

### DaisyUI Card with Button Indicator Example (HTML)

Source: https://daisyui.com/components/indicator

Shows how to use an indicator, specifically a button, to highlight an action on a card. This can be used for call-to-action elements.

```html
<div class="indicator">
  <span class="indicator-item"
    ><button class="btn btn-primary">Apply</button></span
  >
  <div class="card w-96 bg-base-100 shadow-xl">
    <div class="card-body">
      <h2 class="card-title">Job Title</h2>
      <p>Rerum reiciendis beatae tenetur excepturi</p>
    </div>
  </div>
</div>
```

---

### Chat Bubble Examples (HTML)

Source: https://context7_llms

Illustrates how to create chat bubbles for displaying conversations using DaisyUI. This includes different placements and colors for chat bubbles.

```html
<div class="chat chat-start">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img src="/images/stock/photo-1534528741761-a67e41c11709.jpg" />
    </div>
  </div>
  <div class="chat-header">
    Automated
    <time class="text-xs opacity-50">12:46</time>
  </div>
  <div class="chat-bubble">It's over there, let's go.</div>
</div>

<div class="chat chat-end">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img src="/images/stock/photo-1534528741761-a67e41c11709.jpg" />
    </div>
  </div>
  <div class="chat-header">
    You
    <time class="text-xs opacity-50">12:42</time>
  </div>
  <div class="chat-bubble chat-bubble-primary">That's good. Very good.</div>
</div>
```

---

### Update Artboard classes for responsive dimensions

Source: https://daisyui.com/docs/upgrade

This example shows the replacement of daisyUI's `artboard` and `phone-*` classes with Tailwind CSS's `w-*` and `h-*` classes for setting dimensions.

```html
-
<div class="artboard phone-1"></div>
+
<div class="w-[320px] h-[568px]"></div>
```

```html
-
<div class="artboard phone-2"></div>
+
<div class="w-[375px] h-[667px]"></div>
```

```html
-
<div class="artboard phone-3"></div>
+
<div class="w-[414px] h-[736px]"></div>
```

```html
-
<div class="artboard phone-4"></div>
+
<div class="w-[375px] h-[812px]"></div>
```

```html
-
<div class="artboard phone-5"></div>
+
<div class="w-[414px] h-[896px]"></div>
```

```html
-
<div class="artboard phone-6"></div>
+
<div class="w-[320px] h-[1024px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-1"></div>
+
<div class="w-[568px] h-[320px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-2"></div>
+
<div class="w-[667px] h-[375px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-3"></div>
+
<div class="w-[736px] h-[414px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-4"></div>
+
<div class="w-[812px] h-[375px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-5"></div>
+
<div class="w-[896px] h-[414px]"></div>
```

```html
-
<div class="artboard artboard-horizontal phone-6"></div>
+
<div class="w-[1024px] h-[320px]"></div>
```

---

### Button Component Examples (HTML)

Source: https://context7_llms

Illustrates various button styles, colors, sizes, and modifiers available in DaisyUI. Buttons are used to trigger actions by the user.

```html
<button class="btn">Button</button>
<button class="btn btn-primary">Primary Button</button>
<button class="btn btn-outline btn-secondary">Outline Secondary</button>
<button class="btn btn-error btn-lg">Large Error Button</button>
<button class="btn btn-disabled">Disabled Button</button>
<button class="btn btn-circle">?</button>
```

---

### DaisyUI Steps: Horizontal Layout Example

Source: https://daisyui.com/components/steps

Demonstrates the basic horizontal layout for the DaisyUI Steps component. This is suitable for standard sequential processes where space is not a major constraint.

```html
<ul class="steps">
  <li class="step step-neutral">Register</li>
  <li class="step step-neutral">Choose plan</li>
  <li class="step step-neutral">Purchase</li>
  <li class="step step-neutral">Receive Product</li>
</ul>
```

---

### Update Avatar component classes

Source: https://daisyui.com/docs/upgrade

This example shows the renaming of avatar status classes from generic names to more specific `avatar-*` prefixed classes.

```html
-
<div class="avatar online"></div>
+
<div class="avatar-online"></div>
```

```html
-
<div class="avatar offline"></div>
+
<div class="avatar-offline"></div>
```

```html
-
<div class="avatar placeholder"></div>
+
<div class="avatar-placeholder"></div>
```

---

### Pikaday CDN Integration Example

Source: https://daisyui.com/components/calendar

Shows how to use the Pikaday JavaScript datepicker library with a CDN link. This example initializes a Pikaday instance associated with an HTML element.

```html
<input type="text" id="myDatepicker" />
<script src="https://cdn.jsdelivr.net/npm/pikaday/dist/pikaday.js"></script>
<script>
  var picker = new Pikaday({ field: document.getElementById("myDatepicker") });
</script>
```

---

### DaisyUI Input with Indicator Example (HTML)

Source: https://daisyui.com/components/indicator

Demonstrates how to add a required indicator to an input field. This visually signals that the input is mandatory.

```html
<div class="indicator">
  <span class="indicator-item badge">Required</span>
  <input type="text" placeholder="Type here" class="input input-bordered" />
</div>
```

---

### Theme Controller: Swap Toggle Example

Source: https://daisyui.com/components/theme-controller

Showcases a theme controller using a 'swap' toggle, typically for switching between light and dark modes. This example includes hidden checkboxes and SVG icons for sun and moon, providing a visual indicator of the current theme state.

```html
<label class="swap swap-sunny swap-lg">
  <input type="checkbox" class="theme-controller" value="synthwave" />
  <svg
    class="swap-on fill-current w-10 h-10"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
  >
    <path d="M5.64,17l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,5.64,17Z" />
    <path d="M3,13H1a1,1,0,0,0,0,2H3a1,1,0,0,0,0-2Z" />
    <path
      d="M17.54,5.64l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71a1,1,0,0,0-1.41-1.41Z"
    />
    <path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M11.71,11.29a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,11.29Z"
    />
    <path d="M21,13H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M17.54,18.36a1,1,0,0,0-1.41,0l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,17.54,18.36Z"
    />
    <path d="M21,11H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M11.71,16.71a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,16.71Z"
    />
    <path d="M13,23H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M1,13H3a1,1,0,0,0,0-2H1a1,1,0,0,0,0,2Z" />
    <path
      d="M5.64,5.64a1,1,0,0,0-1.41,0l-.71-.71a1,1,0,1,0,1.41-1.41l.71.71A1,1,0,0,0,5.64,5.64Z"
    />
    <path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <circle cx="12" cy="12" r="3" />
  </svg>
  <svg
    class="swap-off fill-current w-10 h-10"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
  >
    <path
      d="M21.64,13a1,1,0,0,0-1.05-.14,1,1,0,0,0-.5.53,1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71A1,1,0,0,0,21.64,13Zm-9.42,10.26a1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.81-.33h0a1,1,0,0,0-.5.53,1,1,0,0,0,.14,1.05Z"
    />
    <path
      d="M12,2A10,10,0,1,0,22,12,10,10,0,0,0,12,2Zm0,18a8,8,0,1,1,8-8A8,8,0,0,1,12,20Z"
    />
    <path
      d="M19.4,12.25a1,1,0,0,0-.14.53,1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71A1,1,0,0,0,19.4,12.25Z"
    />
    <path
      d="M5.64,12.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path d="M12,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M12,23H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M23,13H21a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M1,13H3a1,1,0,0,0,0-2H1a1,1,0,0,0,0,2Z" />
    <path
      d="M19.4,7.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path
      d="M5.64,7.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path d="M12,11H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
  </svg>
</label>

<label class="swap swap-sunny swap-lg">
  <input type="checkbox" class="theme-controller" value="dark" />
  <svg
    class="swap-on fill-current w-10 h-10"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
  >
    <path d="M5.64,17l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,5.64,17Z" />
    <path d="M3,13H1a1,1,0,0,0,0,2H3a1,1,0,0,0,0-2Z" />
    <path
      d="M17.54,5.64l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71a1,1,0,0,0-1.41-1.41Z"
    />
    <path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M11.71,11.29a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,11.29Z"
    />
    <path d="M21,13H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M17.54,18.36a1,1,0,0,0-1.41,0l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,17.54,18.36Z"
    />
    <path d="M21,11H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path
      d="M11.71,16.71a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,16.71Z"
    />
    <path d="M13,23H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M1,13H3a1,1,0,0,0,0-2H1a1,1,0,0,0,0,2Z" />
    <path
      d="M5.64,5.64a1,1,0,0,0-1.41,0l-.71-.71a1,1,0,1,0,1.41-1.41l.71.71A1,1,0,0,0,5.64,5.64Z"
    />
    <path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <circle cx="12" cy="12" r="3" />
  </svg>
  <svg
    class="swap-off fill-current w-10 h-10"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
  >
    <path
      d="M21.64,13a1,1,0,0,0-1.05-.14,1,1,0,0,0-.5.53,1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71A1,1,0,0,0,21.64,13Zm-9.42,10.26a1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.81-.33h0a1,1,0,0,0-.5.53,1,1,0,0,0,.14,1.05Z"
    />
    <path
      d="M12,2A10,10,0,1,0,22,12,10,10,0,0,0,12,2Zm0,18a8,8,0,1,1,8-8A8,8,0,0,1,12,20Z"
    />
    <path
      d="M19.4,12.25a1,1,0,0,0-.14.53,1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71A1,1,0,0,0,19.4,12.25Z"
    />
    <path
      d="M5.64,12.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path d="M12,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M12,23H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M23,13H21a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
    <path d="M1,13H3a1,1,0,0,0,0-2H1a1,1,0,0,0,0,2Z" />
    <path
      d="M19.4,7.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path
      d="M5.64,7.25a1,1,0,0,0,.14.53,1,1,0,0,0,.5.33l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.64-.27,1,1,0,0,0-.14.53Z"
    />
    <path d="M12,11H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z" />
  </svg>
</label>
```

---

### Badge Component Examples (HTML)

Source: https://context7_llms

Demonstrates how to create badges with different styles, colors, and sizes using DaisyUI's badge component. Badges are used to inform users about the status of specific data.

```html
<span class="badge">Badge</span>
<span class="badge badge-outline">Badge</span>
<span class="badge badge-primary">Primary</span>
<span class="badge badge-secondary badge-lg">Large Secondary</span>
<span class="badge badge-error badge-xl">Extra Large Error</span>
```

---

### DaisyUI Steps: Responsive Layout Example

Source: https://daisyui.com/components/steps

Shows a responsive steps layout that defaults to vertical on small screens and transitions to horizontal on larger screens, optimizing the user experience across devices.

```html
<ul class="steps lg:steps-horizontal">
  <li class="step">Register</li>
  <li class="step">Choose plan</li>
  <li class="step">Purchase</li>
  <li class="step">Receive Product</li>
</ul>
```

---

### Pikaday React Component Example

Source: https://daisyui.com/components/calendar

Shows how to integrate Pikaday into a React application using functional components and hooks. It utilizes `useRef` to get a reference to the input element and `useEffect` to manage the Pikaday instance lifecycle.

```bash
npm i pikaday
```

```typescript
import { useEffect, useRef } from "react";
import Pikaday from "pikaday";

export default function App() {
  const myDatepicker = useRef<HTMLInputElement>(null);

  useEffect(() => {
    const picker = new Pikaday({
      field: myDatepicker.current as HTMLInputElement | null,
    });
    return () => {
      picker.destroy();
    };
  }, []);

  return <input type="text" ref={myDatepicker} />;
}
```

---

### Theme Controller: Dropdown Example

Source: https://daisyui.com/components/theme-controller

Presents a theme controller implemented using a dropdown menu. This approach conserves space and provides a clean way to present multiple theme options, allowing users to select their preferred theme from a list.

```html
<div class="dropdown">
  <div tabindex="0" class="btn m-1">Theme</div>
  <ul
    tabindex="0"
    class="dropdown-menu -mt-1 w-52 border border-base-content/10"
  >
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Retro"
        class="theme-controller "
        value="retro"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Cyberpunk"
        class="theme-controller "
        value="cyberpunk"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Valentine"
        class="theme-controller "
        value="valentine"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Aqua"
        class="theme-controller "
        value="aqua"
      />
    </li>
  </ul>
</div>
<div class="dropdown">
  <div tabindex="0" class="btn m-1">Theme</div>
  <ul
    tabindex="0"
    class="dropdown-menu -mt-1 w-52 border border-base-content/10"
  >
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Retro"
        class="theme-controller "
        value="retro"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Cyberpunk"
        class="theme-controller "
        value="cyberpunk"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Valentine"
        class="theme-controller "
        value="valentine"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Aqua"
        class="theme-controller "
        value="aqua"
      />
    </li>
  </ul>
</div>
<div class="dropdown">
  <div tabindex="0" class="btn m-1">Theme</div>
  <ul
    tabindex="0"
    class="dropdown-menu -mt-1 w-52 border border-base-content/10"
  >
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Retro"
        class="theme-controller "
        value="retro"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Cyberpunk"
        class="theme-controller "
        value="cyberpunk"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Valentine"
        class="theme-controller "
        value="valentine"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Aqua"
        class="theme-controller "
        value="aqua"
      />
    </li>
  </ul>
</div>
<div class="dropdown">
  <div tabindex="0" class="btn m-1">Theme</div>
  <ul
    tabindex="0"
    class="dropdown-menu -mt-1 w-52 border border-base-content/10"
  >
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Retro"
        class="theme-controller "
        value="retro"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Cyberpunk"
        class="theme-controller "
        value="cyberpunk"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Valentine"
        class="theme-controller "
        value="valentine"
      />
    </li>
    <li>
      <input
        type="radio"
        name="theme-controller"
        aria-label="Aqua"
        class="theme-controller "
        value="aqua"
      />
    </li>
  </ul>
</div>
```

---

### Theme Controller: Toggle with Text Example

Source: https://daisyui.com/components/theme-controller

Demonstrates a theme controller using a toggle that displays text, such as 'Default' and 'Current'. This provides clear visual feedback to the user about the selected theme.

```html
<label class="swap swap-sunny">
  <input type="checkbox" class="theme-controller" value="synthwave" />
  <span class="swap-on">Synthwave</span>
  <span class="swap-off">Default</span>
</label>
<label class="swap swap-sunny">
  <input type="checkbox" class="theme-controller" value="retro" />
  <span class="swap-on">Retro</span>
  <span class="swap-off">Default</span>
</label>
<label class="swap swap-sunny">
  <input type="checkbox" class="theme-controller" value="cyberpunk" />
  <span class="swap-on">Cyberpunk</span>
  <span class="swap-off">Default</span>
</label>
<label class="swap swap-sunny">
  <input type="checkbox" class="theme-controller" value="valentine" />
  <span class="swap-on">Valentine</span>
  <span class="swap-off">Default</span>
</label>
<label class="swap swap-sunny">
  <input type="checkbox" class="theme-controller" value="aqua" />
  <span class="swap-on">Aqua</span>
  <span class="swap-off">Default</span>
</label>
```

---

### DaisyUI Warning Color Alert Example

Source: https://daisyui.com/components/alert

Provides an example of the 'alert-warning' class for warnings or potential issues. Use this to alert users to non-critical problems.

```html
<div class="alert alert-warning">Warning: Invalid email address!</div>
```

---

### DaisyUI Info Color Alert Example

Source: https://daisyui.com/components/alert

Shows how to use the 'alert-info' class to create an informational alert. This style is typically used for non-critical updates or suggestions.

```html
<div class="alert alert-info">New software update available.</div>
```

---

### DaisyUI Steps: Vertical Layout Example

Source: https://daisyui.com/components/steps

Illustrates the vertical layout for the DaisyUI Steps component, ideal for narrower screens or when emphasizing a top-to-bottom progression.

```html
<ul class="steps steps-vertical">
  <li class="step">Register</li>
  <li class="step">Choose plan</li>
  <li class="step">Purchase</li>
  <li class="step">Receive Product</li>
</ul>
```

---

### DaisyUI Text Input: Email with Icon and Validator

Source: https://daisyui.com/components/input

Demonstrates an email input field styled with an icon and validation message using DaisyUI. This example guides users to enter a correctly formatted email address.

````html
```html
<label class="input input-bordered flex items-center gap-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 16 16"
    fill="currentColor"
    class="w-4 h-4 opacity-70"
  >
    <path
      d="M2.5 3A1.5 1.5 0 0 0 1 4.5v5.086a1.5 1.5 0 0 0 .437.971l3.147 3.147a1.5 1.5 0 0 0 2.072.012l3.147-3.147a1.5 1.5 0 0 0 .437-.971V4.5A1.5 1.5 0 0 0 11.5 3H2.5Z"
    />
    <path
      d="M1.75 5.75a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 .75.75v3.086a.75.75 0 0 1-.22.547l-3.147 3.147a.75.75 0 0 1-1.06.012l-3.147-3.147a.75.75 0 0 1-.22-.547V5.75Z"
    />
  </svg>
  <input type="email" class="grow" placeholder="Email" />
</label>
<label class="label">
  <span class="label-text-alt">Enter valid email address</span>
</label>
````

````

--------------------------------

### DaisyUI Indicator Placement: Top-Start (HTML)

Source: https://daisyui.com/components/indicator

Demonstrates the `indicator-top` and `indicator-start` classes to position an indicator at the top-left corner of a container element.

```html
<div class="indicator indicator-top indicator-start">
  <span class="indicator-item badge badge-secondary">New</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
````

---

### DaisyUI Drawer Sidebar Menu

Source: https://context7_llms

An example of a typical sidebar menu structure within a DaisyUI Drawer. This menu is a list (`<ul>`) with list items (`<li>`) and links (`<a>`). The `menu` class applies styling, and `p-4`, `w-80`, `min-h-full`, `bg-base-100`, and `text-base-content` are example utility classes for padding, width, height, background, and text color.

```html
<ul class="menu p-4 w-80 min-h-full bg-base-100 text-base-content">
  <li><a>Item 1</a></li>
  <li><a>Item 2</a></li>
</ul>
```

---

### DaisyUI Navbar: Icon at Start and End

Source: https://daisyui.com/components/navbar

Illustrates a DaisyUI Navbar configuration with icons placed at both the start and end of the navigation bar, offering more interactive elements.

```html
<div class="navbar bg-base-100">
  <div class="flex-none">
    <button class="btn btn-square btn-ghost">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h7"
        />
      </svg>
    </button>
  </div>
  <div class="flex-1">
    <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
  </div>
  <div class="flex-none">
    <button class="btn btn-square btn-ghost">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </button>
  </div>
</div>
```

---

### Theme Controller: Toggle with Icons Example

Source: https://daisyui.com/components/theme-controller

Illustrates a theme controller using a toggle with icons, typically sun and moon. This visual approach enhances user experience by providing an intuitive way to switch between themes.

```html
<label class="swap swap-sunny swap-lg">
  <input type="checkbox" class="theme-controller" value="synthwave" />
  <svg class="swap-on fill-current w-10 h-10" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M5.64,17l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,5.64,17Z"/><path d="M3,13H1a1,1,0,0,0,0,2H3a1,1,0,0,0,0-2Z"/><path d="M17.54,5.64l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71a1,1,0,0,0-1.41-1.41Z"/><path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z"/><path d="M11.71,11.29a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,11.29Z"/><path d="M21,13H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z"/><path d="M17.54,18.36a1,1,0,0,0-1.41,0l-.71.71a1,1,0,0,0,1.41,1.41l.71-.71A1,1,0,0,0,17.54,18.36Z"/><path d="M21,11H19a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z"/><path d="M11.71,16.71a1,1,0,0,0-1.41,0l-.71.71a1,1,0,1,0,1.41,1.41l.71-.71A1,1,0,0,0,11.71,16.71Z"/><path d="M13,23H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z"/><path d="M1,13H3a1,1,0,0,0,0-2H1a1,1,0,0,0,0,2Z"/><path d="M5.64,5.64a1,1,0,0,0-1.41,0l-.71-.71a1,1,0,1,0,1.41-1.41l.71.71A1,1,0,0,0,5.64,5.64Z"/><path d="M13,1H11a1,1,0,0,0,0,2h2a1,1,0,0,0,0-2Z"/><circle cx="12" cy="12" r="3"/></svg>
  <svg class="swap-off fill-current w-10 h-10" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M21.64,13a1,1,0,0,0-1.05-.14,1,1,0,0,0-.5.53,1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71A1,1,0,0,0,21.64,13Zm-9.42,10.26a1,1,0,0,0,.33.81l.71.71a1,1,0,1,0,1.41-1.41l-.71-.71a1,1,0,0,0-.81-.
```

---

### DaisyUI Drawer Collapsible Sidebar Example

Source: https://context7_llms

This example shows a DaisyUI Drawer with a sidebar that can collapse to show only icons and expand to show icons with text. It uses custom utility classes like `is-drawer-close:overflow-visible`, `is-drawer-close:w-14`, `is-drawer-open:w-64`, and `is-drawer-close:hidden` for conditional styling based on the drawer's state.

```html
<div class="drawer lg:drawer-open">
  <input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content">
    <!-- Page content here -->
  </div>
  <div class="drawer-side is-drawer-close:overflow-visible">
    <label
      for="my-drawer-4"
      aria-label="close sidebar"
      class="drawer-overlay"
    ></label>
    <div
      class="is-drawer-close:w-14 is-drawer-open:w-64 bg-base-200 flex flex-col items-start min-h-full"
    >
      <!-- Sidebar content here -->
      <ul class="menu w-full grow">
        <!-- list item -->
        <li>
          <button
            class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
            data-tip="Homepage"
          >
            🏠
            <span class="is-drawer-close:hidden">Homepage</span>
          </button>
        </li>
        <!-- list item -->
        <li>
          <button
            class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
            data-tip="Settings"
          >
            ⚙️
            <span class="is-drawer-close:hidden">Settings</span>
          </button>
        </li>
      </ul>
      <!-- button to open/close drawer -->
      <div
        class="m-2 is-drawer-close:tooltip is-drawer-close:tooltip-right"
        data-tip="Open"
      >
        <label
          for="my-drawer-4"
          class="btn btn-ghost btn-circle drawer-button is-drawer-open:rotate-y-180"
        >
          ↔️
        </label>
      </div>
    </div>
  </div>
</div>
```

---

### DaisyUI Drawer Always Visible Sidebar Example

Source: https://context7_llms

An example demonstrating a DaisyUI Drawer where the sidebar is always visible on large screens (`lg:drawer-open`) but can be toggled on smaller screens. It includes a page content section and a sidebar section. A label with `lg:hidden` is used to toggle the drawer on small screens.

```html
<div class="drawer lg:drawer-open">
  <input id="my-drawer-3" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content flex flex-col items-center justify-center">
    <!-- Page content here -->
    <label for="my-drawer-3" class="btn drawer-button lg:hidden">
      Open drawer
    </label>
  </div>
  <div class="drawer-side">
    <label
      for="my-drawer-3"
      aria-label="close sidebar"
      class="drawer-overlay"
    ></label>
    <ul class="menu bg-base-200 min-h-full w-80 p-4">
      <!-- Sidebar content here -->
      <li><button>Sidebar Item 1</button></li>
      <li><button>Sidebar Item 2</button></li>
    </ul>
  </div>
</div>
```

---

### DaisyUI Soft Style Alert Example

Source: https://daisyui.com/components/alert

Shows how to apply the 'alert-soft' style for a subtler alert appearance. This is useful for less intrusive notifications.

```html
<div class="alert alert-soft">12 unread messages. Tap to see.</div>
<div class="alert alert-info alert-soft">Your purchase has been confirmed!</div>
<div class="alert alert-warning alert-soft">
  Warning: Invalid email address!
</div>
<div class="alert alert-error alert-soft">Error! Task failed successfully.</div>
```

---

### Theme Controller: Checkbox Example

Source: https://daisyui.com/components/theme-controller

Demonstrates how to implement a theme controller using a checkbox input. This pattern is suitable for toggling between two themes, often light and dark. The provided HTML structure includes hidden inputs and labels for icons, allowing for a visually appealing toggle.

```html
<input type="checkbox" class="theme-controller" value="synthwave" />
<input type="checkbox" class="theme-controller" value="dark" />
<input type="checkbox" class="theme-controller" value="cupcake" />
<input type="checkbox" class="theme-controller" value="bumblebee" />
<input type="checkbox" class="theme-controller" value="emerald" />
<input type="checkbox" class="theme-controller" value="corporate" />
<input type="checkbox" class="theme-controller" value="synthwave" />
<input type="checkbox" class="theme-controller" value="retro" />
<input type="checkbox" class="theme-controller" value="cyberpunk" />
<input type="checkbox" class="theme-controller" value="valentine" />
<input type="checkbox" class="theme-controller" value="halloween" />
<input type="checkbox" class="theme-controller" value="garden" />
<input type="checkbox" class="theme-controller" value="forest" />
<input type="checkbox" class="theme-controller" value="aqua" />
<input type="checkbox" class="theme-controller" value="lofi" />
<input type="checkbox" class="theme-controller" value="pastel" />
<input type="checkbox" class="theme-controller" value="fantasy" />
<input type="checkbox" class="theme-controller" value="wireframe" />
<input type="checkbox" class="theme-controller" value="black" />
<input type="checkbox" class="theme-controller" value="luxury" />
<input type="checkbox" class="theme-controller" value="dracula" />
<input type="checkbox" class="theme-controller" value="cmyk" />
<input type="checkbox" class="theme-controller" value="autumn" />
<input type="checkbox" class="theme-controller" value="business" />
<input type="checkbox" class="theme-controller" value="acid" />
<input type="checkbox" class="theme-controller" value="lemonade" />
<input type="checkbox" class="theme-controller" value="night" />
<input type="checkbox" class="theme-controller" value="coffee" />
<input type="checkbox" class="theme-controller" value="winter" />
<input type="checkbox" class="theme-controller" value="dim" />
<input type="checkbox" class="theme-controller" value="nord" />
<input type="checkbox" class="theme-controller" value="sunset" />
<input type="checkbox" class="theme-controller" value="caramellatte" />
<input type="checkbox" class="theme-controller" value="abyss" />
<input type="checkbox" class="theme-controller" value="silk" />
```

---

### Import daisyUI in CSS file

Source: https://daisyui.com/docs/upgrade

This snippet demonstrates how to import daisyUI into the main CSS file, either as a general plugin or with specific themes enabled.

```css
@import "tailwindcss";
@plugin "daisyui";
```

```css
@import "tailwindcss";
@plugin "daisyui" {
  themes: light --default, dark --prefersdark, cupcake;
}
```

---

### File Input with Fieldset and Label - daisyUI

Source: https://daisyui.com/components/file-input

Shows how to group a file input with a fieldset and label, enhancing accessibility and user guidance. This example includes descriptive text like 'Pick a file' and size constraints.

```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">What's your profile picture?</span>
    <span class="label-text-alt">Max size 2MB</span>
  </label>
  <input type="file" class="file-input w-full max-w-xs" />
</div>
```

---

### Status Component with Bounce Animation (HTML)

Source: https://daisyui.com/components/status

Provides an example of applying a bounce animation to the Status component, often used for attention-grabbing notifications or new items. The example illustrates unread messages.

````html
```html
<div class="status status-info" style="--status-color: inherit;">
  Unread messages
</div>
<span
  class="status status-info animate-bounce"
  style="--status-color: inherit;"
></span>
<span class="status status-info" style="--status-color: inherit;"></span>

<div class="status status-warning" style="--status-color: inherit;">
  New notifications
</div>
<span
  class="status status-warning animate-bounce"
  style="--status-color: inherit;"
></span>
<span class="status status-warning" style="--status-color: inherit;"></span>
````

````

--------------------------------

### DaisyUI Browser Mockup Syntax (HTML)

Source: https://context7_llms

Shows how to create a browser mockup using DaisyUI. It includes the basic structure with a toolbar and content area. A URL can be set in the toolbar by using an input element.

```html
<div class="mockup-browser">
  <div class="mockup-browser-toolbar">
    <div class="input"><input type="text" value="daisyui.com" /></div>
  </div>
  <div class="mockup-browser-window">
    Content of the browser window
  </div>
</div>
````

---

### Theme Controller: Radio Input Example

Source: https://daisyui.com/components/theme-controller

Illustrates using radio inputs for the theme controller, allowing users to select one theme from a predefined list. This is useful for scenarios where a specific selection is required, such as choosing a primary theme for the application.

```html
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="retro"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="cyberpunk"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="valentine"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="aqua"
/>

<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="retro"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="cyberpunk"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="valentine"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="aqua"
/>

<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="retro"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="cyberpunk"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="valentine"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="aqua"
/>

<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="retro"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="cyberpunk"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="valentine"
/>
<input
  type="radio"
  name="theme-controller"
  class="theme-controller"
  value="aqua"
/>
```

---

### Stat with Download Data (HTML)

Source: https://daisyui.com/components/stat

This snippet illustrates a Stat component displaying download-related data, including volume and date range. It also shows examples with increase/decrease indicators for new users and registers.

```html
<div class="stats">
  <div class="stat">
    <div class="stat-title">Downloads</div>
    <div class="stat-value">31K</div>
    <div class="stat-desc">Jan 1st - Feb 1st</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Users</div>
    <div class="stat-value">4,200</div>
    <div class="stat-desc">↗︎ 400 (22%)</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Registers</div>
    <div class="stat-value">1,200</div>
    <div class="stat-desc">↘︎ 90 (14%)</div>
  </div>
</div>
```

---

### DaisyUI Code Mockup Syntax (HTML)

Source: https://context7_llms

Illustrates the DaisyUI code mockup component for displaying code snippets. It supports prefixes for lines and can be configured for line highlighting.

```html
<div class="mockup-code">
  <pre data-prefix="$"><code>npm i daisyui</code></pre>
  <pre data-prefix=">" data-line><code>console.log('Hello world!')</code></pre>
</div>
```

---

### DaisyUI Stack Component: Stacked Cards (Start Direction)

Source: https://daisyui.com/components/stack

Illustrates stacking card elements vertically using the 'stack-start' modifier, aligning them to the start (left) side of the stack. This is useful for left-aligned stacked content.

```html
<div class="stack stack-start">
  <div>A</div>
  <div>B</div>
  <div>C</div>
</div>
```

---

### Basic Code Mockup with daisyUI

Source: https://daisyui.com/components/mockup-code

Demonstrates the fundamental usage of the 'mockup-code' class to display a simple code block. This is useful for showing commands or basic code snippets.

```html
<div class="mockup-code">
  <pre data-prefix="$"><code>npm i daisyui</code></pre>
</div>
```

---

### Basic DaisyUI Timeline with Text and Icons (HTML)

Source: https://daisyui.com/components/timeline

Demonstrates a basic DaisyUI timeline structure with events displayed chronologically. This example utilizes the default vertical layout and includes event titles and descriptions.

```html
<ul class="timeline">
  <li>
    <div class="timeline-start">
      1984
      <span class="text-lg">First Macintosh computer</span>
    </div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>1984</time>
      First Macintosh computer
    </div>
  </li>
  <li>
    <div class="timeline-start">
      1998
      <span class="text-lg">iMac</span>
    </div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>1998</time>
      iMac
    </div>
  </li>
  <li>
    <div class="timeline-start">
      2001
      <span class="text-lg">iPod</span>
    </div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2001</time>
      iPod
    </div>
  </li>
  <li>
    <div class="timeline-start">
      2007
      <span class="text-lg">iPhone</span>
    </div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2007</time>
      iPhone
    </div>
  </li>
  <li>
    <div class="timeline-start">
      2015
      <span class="text-lg">Apple Watch</span>
    </div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2015</time>
      Apple Watch
    </div>
  </li>
</ul>
```

---

### DaisyUI Outline Style Alert Example

Source: https://daisyui.com/components/alert

Illustrates the 'alert-outline' class for alerts with a distinct outline. This style provides a clear visual separation.

```html
<div class="alert alert-outline">12 unread messages. Tap to see.</div>
<div class="alert alert-info alert-outline">
  Your purchase has been confirmed!
</div>
<div class="alert alert-warning alert-outline">
  Warning: Invalid email address!
</div>
<div class="alert alert-error alert-outline">
  Error! Task failed successfully.
</div>
```

---

### DaisyUI Toast Placement: Top-Start

Source: https://daisyui.com/components/toast

Shows how to position the DaisyUI Toast component at the top-start (top-left) corner of the page. This example uses the 'toast-top' and 'toast-start' classes for precise placement.

```html
<div class="toast toast-top toast-start">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Dash Style Alert Example

Source: https://daisyui.com/components/alert

Demonstrates the 'alert-dash' style for alerts with a dashed border. This offers a different visual emphasis compared to solid or outline styles.

```html
<div class="alert alert-dash">12 unread messages. Tap to see.</div>
<div class="alert alert-info alert-dash">Your purchase has been confirmed!</div>
<div class="alert alert-warning alert-dash">
  Warning: Invalid email address!
</div>
<div class="alert alert-error alert-dash">Error! Task failed successfully.</div>
```

---

### DaisyUI Indicator Placement: Top-Center (HTML)

Source: https://daisyui.com/components/indicator

Shows how to use `indicator-top` and `indicator-center` to place an indicator at the top-center of a container.

```html
<div class="indicator indicator-top indicator-center">
  <span class="indicator-item badge badge-secondary">New</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
```

---

### Replace Bottom Navigation component with Dock

Source: https://daisyui.com/docs/upgrade

This snippet illustrates the migration from the deprecated `bottom-nav` component and its associated classes to the new `dock` component and its corresponding classes.

```html
-
<div class="btm-nav">
  - <button>Legal</button> - <button class="btm-nav-active">Favorites</button> -
  <button>Settings</button> -
</div>
+
<div class="dock">
  + <button>🏠</button> + <button class="dock-active">🍿</button> +
  <button>⚙️</button> +
</div>
```

---

### Radial Progress with Background and Border (HTML)

Source: https://daisyui.com/components/radial-progress

Demonstrates how to add a background color and border to the radial-progress component. This example shows a 70% progress with a distinct background and border for better visibility.

```html
<div
  class="radial-progress bg-base-200 border border-base-300"
  style="--value:70"
  role="progressbar"
>
  70%
</div>
```

---

### DaisyUI Radio Button Examples

Source: https://daisyui.com/components/radio

Demonstrates the basic structure and usage of DaisyUI radio buttons. Radio buttons allow users to select a single option from a set. Ensure each radio button group has unique 'name' attributes.

```html
<!-- Basic Radio -->
<div class="form-control">
  <label class="label cursor-pointer">
    <span class="label-text">Remember me</span>
    <input type="radio" name="radio-10" class="radio checked:bg-blue-500" />
  </label>
</div>

<!-- Radio with different colors -->
<input type="radio" name="radio-10" class="radio radio-primary" checked />
<input type="radio" name="radio-10" class="radio radio-secondary" />
<input type="radio" name="radio-10" class="radio radio-accent" />
<input type="radio" name="radio-10" class="radio radio-ghost" />
```

---

### DaisyUI Window Mockup Syntax (HTML)

Source: https://context7_llms

Demonstrates the basic DaisyUI window mockup component. It consists of a container div where the window content can be placed.

```html
<div class="mockup-window">
  <div>
    <!-- Content inside the window -->
  </div>
</div>
```

---

### DaisyUI Indicator Placement: Top-End (Default) (HTML)

Source: https://daisyui.com/components/indicator

Illustrates the default top-end placement using `indicator-top` and `indicator-end` (where `indicator-end` is the default horizontal alignment).

```html
<div class="indicator indicator-top indicator-end">
  <span class="indicator-item badge badge-secondary">New</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
```

---

### DaisyUI Success Color Alert Example

Source: https://daisyui.com/components/alert

Illustrates the 'alert-success' class for displaying successful actions or confirmations. This is commonly used for positive feedback.

```html
<div class="alert alert-success">Your purchase has been confirmed!</div>
```

---

### Indicator Component Placement Example in HTML

Source: https://context7_llms

The indicator component places an element, such as a badge or icon, on the corner of another element. It uses 'indicator' and 'indicator-item' classes, with optional placement classes for positioning.

```html
<div class="indicator">
  <span class="indicator-item badge badge-secondary">!</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
```

---

### Radial Progress with Background and Border (Vue)

Source: https://daisyui.com/components/radial-progress

This Vue.js example shows how to style the radial-progress component with a background color and border, enhancing its visual presentation.

```vue
<template>
  <div
    class="radial-progress bg-base-200 border border-base-300"
    :style="{ '--value': 70 }"
    role="progressbar"
  >
    70%
  </div>
</template>

<script setup lang="ts">
// No script needed for basic example
</script>

<style>
/* DaisyUI styles are assumed to be globally imported */
</style>
```

---

### DaisyUI Error Color Alert Example

Source: https://daisyui.com/components/alert

Demonstrates the 'alert-error' class for critical errors or failures. This style should be used for significant problems that require user attention.

```html
<div class="alert alert-error">Error! Task failed successfully.</div>
```

---

### Accordion with Radio Inputs Example (HTML)

Source: https://daisyui.com/components/accordion

Demonstrates how to create an accordion component using HTML radio inputs. This method ensures that only one accordion item can be open at a time within a group sharing the same 'name' attribute. It's suitable for FAQs or sequential content presentation.

```html
<div>
  <div class="collapse collapse-arrow bg-base-200">
    <input type="radio" name="my-accordion-2" checked="checked" />
    <div class="collapse-title text-xl font-medium">
      Click to open this one and close others
    </div>
    <div class="collapse-content">
      <p>hello</p>
    </div>
  </div>
  <div class="collapse collapse-arrow bg-base-200">
    <input type="radio" name="my-accordion-2" />
    <div class="collapse-title text-xl font-medium">
      Click to open this one and close others
    </div>
    <div class="collapse-content">
      <p>hello</p>
    </div>
  </div>
  <div class="collapse collapse-arrow bg-base-200">
    <input type="radio" name="my-accordion-2" />
    <div class="collapse-title text-xl font-medium">
      Click to open this one and close others
    </div>
    <div class="collapse-content">
      <p>hello</p>
    </div>
  </div>
</div>
```

---

### DaisyUI Fieldset with Background and Border

Source: https://daisyui.com/components/fieldset

Shows how to apply background and border styles to a DaisyUI fieldset, enhancing its visual separation and grouping capabilities. This example builds upon the basic fieldset structure.

```html
<!-- Fieldset with background and border -->
<div class="form-control bg-base-200 p-4 rounded-lg border border-base-300">
  <label class="label">
    <span class="label-text">Page title</span>
  </label>
  <input type="text" placeholder="Type here" class="input input-bordered" />
</div>
```

---

### DaisyUI Toast Placement: Top-End

Source: https://daisyui.com/components/toast

Demonstrates how to align the DaisyUI Toast component to the top-end (top-right) of the viewport. This example employs the 'toast-top' and 'toast-end' classes.

```html
<div class="toast toast-top toast-end">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Join Component: Radio Inputs

Source: https://daisyui.com/components/join

Shows an example of using the DaisyUI 'join' component with radio inputs styled as buttons. Each radio input is a 'join-item' within the 'join' container, allowing for visually grouped radio selections.

```html
<div class="join">
  <input
    type="radio"
    name="options"
    class="join-item btn"
    checked
    aria-label="rdo1"
  />
  <input type="radio" name="options" class="join-item btn" aria-label="rdo2" />
</div>
<div class="join">
  <input
    type="radio"
    name="options"
    class="join-item btn"
    checked
    aria-label="rdo1"
  />
  <input type="radio" name="options" class="join-item btn" aria-label="rdo2" />
</div>
<div class="join">
  <input
    type="radio"
    name="options"
    class="join-item btn"
    checked
    aria-label="rdo1"
  />
  <input type="radio" name="options" class="join-item btn" aria-label="rdo2" />
</div>
<div class="join">
  <input
    type="radio"
    name="options"
    class="join-item btn"
    checked
    aria-label="rdo1"
  />
  <input type="radio" name="options" class="join-item btn" aria-label="rdo2" />
</div>
```

---

### DaisyUI Badge with Icon (HTML)

Source: https://daisyui.com/components/badge

Provides examples of how to include icons within DaisyUI badges. This enhances the visual communication of the badge's status or meaning.

```html
<!-- Example using a hypothetical SVG icon -->
<!-- Info Badge with Icon -->
<div class="badge badge-info gap-x-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    class="w-4 h-4 stroke-current"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M13 16h-1v2h-1V9h1m4 4h-1v6h-1M7 12a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm9 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"
    ></path>
  </svg>
  <span>Info</span>
</div>

<!-- Success Badge with Icon -->
<div class="badge badge-success gap-x-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    class="w-4 h-4 stroke-current"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M5 13l4 4L19 7"
    ></path>
  </svg>
  <span>Success</span>
</div>

<!-- Warning Badge with Icon -->
<div class="badge badge-warning gap-x-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    class="w-4 h-4 stroke-current"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M12 9v2m0 4h.01m-6.938 4h13.856c1.543 0 2.872-1.258 2.872-2.815V7.015c0-1.557-1.329-2.815-2.872-2.815H5.062C3.519 4.2 2.19 5.458 2.19 7.015v9.97c0 1.557 1.329 2.815 2.872 2.815z"
    ></path>
  </svg>
  <span>Warning</span>
</div>

<!-- Error Badge with Icon -->
<div class="badge badge-error gap-x-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    class="w-4 h-4 stroke-current"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M6 18L18 6M6 6l12 12"
    ></path>
  </svg>
  <span>Error</span>
</div>
```

---

### DaisyUI Navbar: Basic Title Only

Source: https://daisyui.com/components/navbar

This snippet shows the fundamental structure of a DaisyUI Navbar containing only a title. It serves as the most basic navigation bar setup.

```html
<div class="navbar bg-base-100">
  <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
</div>
```

---

### DaisyUI Box Tabs with HTML

Source: https://daisyui.com/components/tab

Demonstrates the 'tabs-box' style for DaisyUI tabs, which presents the tabs within a box container. This example applies the 'tabs-box' class to the tab list.

````html
### tabs-box Tab 1 Tab 2 Tab 3 ``` html
<div role="tablist" class="tabs tabs-box">
  <a role="tab" class="tab">Tab 1</a>
  <a role="tab" class="tab">Tab 2</a>
  <a role="tab" class="tab">Tab 3</a>
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 1
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 2
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 3
</div>
````

````

--------------------------------

### daisyUI Dock Custom Colors HTML Example

Source: https://daisyui.com/components/dock

Illustrates how to apply custom colors to the daisyUI Dock component using Tailwind CSS color classes. This allows for flexible theming and branding.

```html
<div class="dock bg-primary">
  <button class="dock-tab text-primary-content">
    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l9-9 9 9M5 10v7a2 2 0 002 2h10a2 2 0 002-2v-7" /></svg>
    Home
  </button>
  <button class="dock-tab dock-active text-primary-content bg-primary-focus">
    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 1a3 3 0 00-3 3v13a3 3 0 103 3h0v-13z" /></svg>
    Inbox
  </button>
  <button class="dock-tab text-primary-content">
    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-0.657 0-1.275-.344-1.658-.874-0.383-.53-.594-1.197-.594-1.879v0c0-1.38 1.12-2.5 2.5-2.5s2.5 1.12 2.5 2.5v0c0 0.682-.211 1.349-.594 1.879-.383.531-.999.874-1.658.874h0z" /></svg>
    Settings
  </button>
</div>
````

---

### Centered Stat Items (HTML)

Source: https://daisyui.com/components/stat

This example demonstrates how to center the items within the Stat component. It's useful for creating a more compact and visually balanced display of statistics.

```html
<div class="stats">
  <div class="stat place-items-center">
    <div class="stat-title">Downloads</div>
    <div class="stat-value">31K</div>
    <div class="stat-desc">From January 1st to February 1st</div>
  </div>
  <div class="stat place-items-center">
    <div class="stat-title">Users</div>
    <div class="stat-value">4,200</div>
    <div class="stat-desc">↗︎ 40 (2%)</div>
  </div>
  <div class="stat place-items-center">
    <div class="stat-title">New Registers</div>
    <div class="stat-value">1,200</div>
    <div class="stat-desc">↘︎ 90 (14%)</div>
  </div>
</div>
```

---

### Clock Countdown with Colons (HTML)

Source: https://daisyui.com/components/countdown

Provides an example of a clock countdown formatted with colons separating hours, minutes, and seconds. This layout is useful for time displays.

```html
<div class="countdown">
  <span style="--value:10"></span>: <span style="--value:24"></span>:
  <span style="--value:59"></span>
</div>
```

---

### Skeleton with Circle and Content - daisyUI

Source: https://daisyui.com/components/skeleton

This example shows how to implement a circular Skeleton component that can display content once loaded. It's useful for avatar placeholders or other circular elements. Ensure daisyUI is included in your project.

```html
<div class="skeleton circle"></div>
```

---

### DaisyUI Text Input: Basic and With Label

Source: https://daisyui.com/components/input

Demonstrates the basic DaisyUI text input field and how to incorporate a text label within the input container. This is a foundational example for creating user input elements.

````html
```html
<input
  type="text"
  placeholder="Text Input"
  class="input input-bordered w-full max-w-xs"
/>
````

```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">What is your name?</span>
  </label>
  <input
    type="text"
    placeholder="Type here"
    class="input input-bordered w-full max-w-xs"
  />
</div>
```

````

--------------------------------

### Checkbox Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Provides an example of validating a checkbox using DaisyUI. The 'validator' class is applied, and the 'validator-hint' indicates that the checkbox is required.

```html
<input type="checkbox" class="checkbox validator">
<p class="validator-hint">Required</p>
````

---

### DaisyUI Toast Placement: Bottom-End (Default)

Source: https://daisyui.com/components/toast

Demonstrates the default bottom-end (bottom-right) placement for the DaisyUI Toast component. This example relies on the default 'toast-bottom' and 'toast-end' behavior.

```html
<div class="toast toast-end toast-bottom">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### Radial Progress with Custom Size and Thickness (HTML)

Source: https://daisyui.com/components/radial-progress

Explains how to customize the size and thickness of the radial-progress component using the '--size' and '--thickness' CSS variables. The example shows a 70% progress with specified dimensions.

```html
<div
  class="radial-progress"
  style="--value:70; --size: 10rem; --thickness: 1rem;"
  role="progressbar"
>
  70%
</div>
```

---

### Basic Drawer Sidebar Implementation (HTML)

Source: https://daisyui.com/components/drawer

An example of a functional drawer sidebar. It includes a toggle button (represented by a label linked to a hidden checkbox) and content areas for the page and the sidebar. Clicking the toggle or sidebar items will manage the drawer's state.

```html
<div class="drawer">
  <input id="my-drawer" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content">
    {/* Page content here */}
    <label for="my-drawer" class="btn btn-primary drawer-button"
      >Open drawer</label
    >
  </div>
  <div class="drawer-side">
    <label
      for="my-drawer"
      aria-label="close sidebar"
      class="drawer-overlay"
    ></label>
    <ul class="menu p-4 w-80 bg-base-200 text-base-content">
      {/* Sidebar content here */}
      <li><a>Sidebar Item 1</a></li>
      <li><a>Sidebar Item 2</a></li>
    </ul>
  </div>
</div>
```

---

### Import daisyUI as a Plugin in Tailwind CSS (JS Config)

Source: https://daisyui.com/docs/v5

Shows the previous method of including daisyUI as a plugin within the `tailwind.config.js` file. This configuration is for older versions or specific setups.

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  plugins: [
    require('daisyui');
  ],
}
```

---

### Responsive Horizontal Divider

Source: https://daisyui.com/components/divider

Demonstrates how to apply a responsive horizontal divider using daisyUI's responsive prefixes. This example shows the divider behaving horizontally on larger screens.

```html
<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>

<div class="lg:divider-horizontal"></div>

<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>
```

```html
<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>

<div class="lg:divider-horizontal"></div>

<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>
```

```html
<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>

<div class="lg:divider-horizontal"></div>

<div>content</div>

<div class="lg:divider-horizontal">OR</div>

<div>content</div>
```

---

### Basic Skeleton Component - daisyUI

Source: https://daisyui.com/components/skeleton

This code snippet demonstrates the basic usage of the Skeleton component in daisyUI. It renders a simple placeholder div with a loading animation. No external dependencies are required beyond daisyUI.

```html
<div class="skeleton"></div>
```

---

### DaisyUI Button Default and Sizes

Source: https://daisyui.com/components/button

Demonstrates the default button appearance and how to apply different size modifiers (btn-xs, btn-sm, btn-md, btn-lg, btn-xl) to control button dimensions. These examples show basic button creation and size adjustments.

```html
<button class="btn">Default</button>
<button class="btn btn-xs">Xsmall</button>
<button class="btn btn-sm">Small</button>
<button class="btn btn-md">Medium</button>
<button class="btn btn-lg">Large</button>
<button class="btn btn-xl">Xlarge</button>
```

---

### Number Input Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Illustrates number input validation using DaisyUI. This example enforces that the number must be between 1 and 10.

```html
<input
  type="number"
  class="input input-bordered validator"
  placeholder="number"
/>
<p class="validator-hint">Must be between be 1 to 10</p>
```

---

### DaisyUI Rating with Warning Color

Source: https://daisyui.com/components/rating

This example shows a rating component using the 'mask-star-2' icon and applying a warning color. This is useful for highlighting negative feedback or alerts.

```html
<div class="rating">
  <input type="radio" name="rating-warning" class="mask mask-star-2" />
  <input type="radio" name="rating-warning" class="mask mask-star-2" checked />
  <input type="radio" name="rating-warning" class="mask mask-star-2" />
  <input type="radio" name="rating-warning" class="mask mask-star-2" />
  <input type="radio" name="rating-warning" class="mask mask-star-2" />
</div>
<div class="rating">
  <input
    type="radio"
    name="rating-warning-half"
    class="bg-warning mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-warning-half"
    class="bg-warning mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-warning-half"
    class="bg-warning mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-warning-half"
    class="bg-warning mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-warning-half"
    class="bg-warning mask mask-star-2"
    checked
  />
</div>
```

---

### Status Component with Ping Animation (HTML)

Source: https://daisyui.com/components/status

Demonstrates how to add a ping animation to the Status component, commonly used to indicate real-time updates or active states. The example shows a server status.

````html
```html
<div class="status status-error" style="--status-color: inherit;">
  Server is down
</div>
<span
  class="status status-error animate-ping"
  style="--status-color: inherit;"
></span>
<span class="status status-error" style="--status-color: inherit;"></span>

<div class="status status-success" style="--status-color: inherit;">
  Server is up
</div>
<span
  class="status status-success animate-ping"
  style="--status-color: inherit;"
></span>
<span class="status status-success" style="--status-color: inherit;"></span>
````

````

--------------------------------

### Accordion with Details Element Example (HTML)

Source: https://daisyui.com/components/accordion

Illustrates creating an accordion using the HTML details element. This approach provides a native browser mechanism for collapsible content and can be made searchable by browsers. It's ideal for content where semantic structure is important and browser searchability is desired.

```html
<div class="collapse collapse-arrow bg-base-200">
  <input type="checkbox" />
  <div class="collapse-title text-xl font-medium">
    Accordion Content
  </div>
  <div class="collapse-content">
    <p>Accordion content</p>
  </div>
</div>
````

---

### DaisyUI Navbar: Dropdown, Center Logo, and Icons

Source: https://daisyui.com/components/navbar

This example showcases a DaisyUI Navbar with a dropdown menu, a centered logo, and icons, suitable for a balanced layout with branding and navigation.

```html
<div class="navbar bg-base-100 rounded-box">
  <div class="navbar-start">
    <div class="dropdown">
      <label tabindex="0" class="btn btn-ghost lg:hidden">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h8m-8 6h8"
          />
        </svg>
      </label>
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li><a>Homepage</a></li>
        <li>
          <a>Parent</a>
          <ul class="p-2">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </li>
        <li><a>About</a></li>
      </ul>
    </div>
  </div>
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1">
      <li><a>Homepage</a></li>
      <li>
        <details>
          <summary>Parent</summary>
          <ul class="p-2 bg-base-100 rounded-t-none">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </details>
      </li>
      <li><a>About</a></li>
    </ul>
  </div>
  <div class="navbar-end">
    <a class="btn btn-ghost btn-circle">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </a>
  </div>
</div>
```

---

### DaisyUI Radio Tabs with Lift and Content (HTML)

Source: https://daisyui.com/components/tab

Illustrates using radio inputs for tab switching with the 'tabs-lift' style. This setup ensures that the correct tab content is displayed when a radio button is selected.

````html
### radio tabs-lift + tab content Tab content 1 Tab content 2 Tab content 3 ```
html
<div class="flex w-full flex-col gap-3">
  <div role="tablist" class="tabs tabs-lift">
    <input
      type="radio"
      name="my_tab_group_4"
      role="tab"
      class="tab"
      aria-label="Tab 1"
      checked
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 1
    </div>

    <input
      type="radio"
      name="my_tab_group_4"
      role="tab"
      class="tab"
      aria-label="Tab 2"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 2
    </div>

    <input
      type="radio"
      name="my_tab_group_4"
      role="tab"
      class="tab"
      aria-label="Tab 3"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 3
    </div>
  </div>
</div>
````

````

--------------------------------

### DaisyUI Footer with Form

Source: https://daisyui.com/components/footer

Presents a DaisyUI footer that includes a newsletter signup form. This example demonstrates how to embed input fields and a submit button within the footer structure for user engagement.

```html
<footer class="footer p-10 bg-base-200 text-base-content">
  <div>
    <span class="footer-title">Services</span>
    <a class="link link-hover">Branding</a>
    <a class="link link-hover">Design</a>
    <a class="link link-hover">Marketing</a>
    <a class="link link-hover">Advertisement</a>
  </div>
  <div>
    <span class="footer-title">Company</span>
    <a class="link link-hover">About us</a>
    <a class="link link-hover">Contact</a>
    <a class="link link-hover">Jobs</a>
    <a class="link link-hover">Press kit</a>
  </div>
  <div>
    <span class="footer-title">Legal</span>
    <a class="link link-hover">Terms of use</a>
    <a class="link link-hover">Privacy policy</a>
    <a class="link link-hover">Cookie policy</a>
  </div>
  <div>
    <span class="footer-title">Newsletter</span>
    <div class="form-control w-80">
      <label class="label">
        <span class="label-text">Enter your email address</span>
      </label>
      <div class="relative">
        <input type="text" placeholder="username@site.com" class="input input-bordered w-full pr-16" />
        <button class="btn btn-primary absolute top-0 right-0 rounded-l-none">Subscribe</button>
      </div>
    </div>
  </div>
</footer>
````

---

### Multi-line Code Mockup with daisyUI

Source: https://daisyui.com/components/mockup-code

Illustrates how to display multi-line code within the 'mockup-code' component. This is suitable for showing script content or complex commands with multiple lines.

```html
<div class="mockup-code">
  <pre data-prefix="$"><code>npm i daisyui
installing...
Done!</code></pre>
</div>
```

---

### DaisyUI Toast Placement: Bottom-Center

Source: https://daisyui.com/components/toast

Illustrates the DaisyUI Toast component positioned at the bottom-center. This example uses 'toast-bottom' and 'toast-center', with 'toast-bottom' as the default vertical placement.

```html
<div class="toast toast-center toast-bottom">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Lifted Tabs with HTML

Source: https://daisyui.com/components/tab

Shows the 'tabs-lift' style for DaisyUI tabs, which gives the tabs a lifted appearance. This example includes the 'tabs-lift' class applied to the tab list container.

````html
### tabs-lift Tab 1 Tab 2 Tab 3 ``` html
<div role="tablist" class="tabs tabs-lift">
  <a role="tab" class="tab">Tab 1</a>
  <a role="tab" class="tab">Tab 2</a>
  <a role="tab" class="tab">Tab 3</a>
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 1
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 2
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 3
</div>
````

````

--------------------------------

### DaisyUI Stack Component: Stacked Notifications

Source: https://daisyui.com/components/stack

Provides an example of using the DaisyUI stack component to layer notification-like elements vertically. This showcases a practical application for presenting sequential messages or alerts.

```html
<div class="stack">
  <div>
    ## Notification 1
    You have 3 unread messages. Tap here to see.
  </div>
  <div>
    ## Notification 2
    You have 3 unread messages. Tap here to see.
  </div>
  <div>
    ## Notification 3
    You have 3 unread messages. Tap here to see.
  </div>
</div>
````

---

### Styling a button with Tailwind CSS and daisyUI

Source: https://daisyui.com/

This example demonstrates how to style a button using daisyUI's semantic class names in conjunction with Tailwind CSS utility classes. The 'Tailwind Button' comment indicates the expected output, showcasing a styled button.

````html
```html // Styling a simple button <button class="btn">Tailwind Button</button>
````

````

--------------------------------

### DaisyUI Icon-Only Menu Component

Source: https://daisyui.com/components/menu

Presents a menu where items are represented by icons only, suitable for toolbars or compact interfaces. Includes examples for both vertical and horizontal layouts.

```html
<ul class="menu menu-vertical w-12">
  <li>
    <a>
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M3 6h18M3 18h18" /></svg>
    </a>
  </li>
  <li>
    <a>
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M3 6h18M3 18h18" /></svg>
    </a>
  </li>
  <li>
    <a>
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M3 6h18M3 18h18" /></svg>
    </a>
  </li>
</ul>
````

```html
<ul class="menu menu-horizontal rounded-box">
  <li>
    <a>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
</ul>
```

---

### Window Mockup Component (HTML)

Source: https://daisyui.com/components/mockup-window

Provides HTML structure for creating a window mockup with optional border and background color using daisyUI classes. This can be used for UI demonstrations or documentation.

```html
<div class="mockup-window border border-base-300">
  <div class="flex justify-center px-4 py-6 bg-base-200">Hello!</div>
</div>

<div class="mockup-window border bg-primary">
  <div class="flex justify-center px-4 py-6 bg-base-200">Hello!</div>
</div>
```

---

### Hover Gallery Component Example in HTML

Source: https://context7_llms

The hover-gallery component displays multiple images that become visible on horizontal hover. It's suitable for e-commerce or portfolios and can contain up to 10 images. A max-width is required for proper display.

```html
<figure class="hover-gallery max-w-60">
  <img src="https://img.daisyui.com/images/stock/daisyui-hat-1.webp" />
  <img src="https://img.daisyui.com/images/stock/daisyui-hat-2.webp" />
  <img src="https://img.daisyui.com/images/stock/daisyui-hat-3.webp" />
  <img src="https://img.daisyui.com/images/stock/daisyui-hat-4.webp" />
</figure>
```

---

### Stat with Icons or Image (HTML)

Source: https://daisyui.com/components/stat

This example shows how to integrate icons or images within the daisyUI Stat component. It includes a 'stat-figure' element for visual elements alongside numerical data and descriptive text.

```html
<div class="stats">
  <div class="stat">
    <div class="stat-figure text-secondary">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        class="inline-block w-8 h-8 stroke-current"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2v-10m-6 4a2 2 0 11-4 0 2 2 0 014 0z"
        ></path>
      </svg>
    </div>
    <div class="stat-title">Total Likes</div>
    <div class="stat-value">25.6K</div>
    <div class="stat-desc">21% more than last month</div>
  </div>
  <div class="stat">
    <div class="stat-figure text-secondary">
      <img
        src="https://img.daisyui.com/images/profile/demo/anakeen@192.webp"
        alt="Tailwind CSS stat example component"
      />
    </div>
    <div class="stat-title">Page Views</div>
    <div class="stat-value">2.6M</div>
    <div class="stat-desc">21% more than last month</div>
  </div>
  <div class="stat">
    <div class="stat-figure text-secondary">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        class="inline-block w-8 h-8 stroke-current"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 10h18M7 10h10M9 14h6M11 18h0"
        ></path>
      </svg>
    </div>
    <div class="stat-title">Tasks done</div>
    <div class="stat-value">86%</div>
    <div class="stat-desc text-secondary">31 tasks remaining</div>
  </div>
</div>
```

---

### DaisyUI Pagination with Different Sizes

Source: https://daisyui.com/components/pagination

Shows how to apply different sizes to DaisyUI pagination components. This allows for responsive design adjustments based on the available space or user preference.

```html
<div class="join">
  <button class="join-item btn btn-xs">1</button>
  <button class="join-item btn btn-xs btn-active">2</button>
  <button class="join-item btn btn-xs">3</button>
  <button class="join-item btn btn-xs">4</button>
</div>
<div class="join">
  <button class="join-item btn btn-sm">1</button>
  <button class="join-item btn btn-sm btn-active">2</button>
  <button class="join-item btn btn-sm">3</button>
  <button class="join-item btn btn-sm">4</button>
</div>
<div class="join">
  <button class="join-item btn btn-md">1</button>
  <button class="join-item btn btn-md btn-active">2</button>
  <button class="join-item btn btn-md">3</button>
  <button class="join-item btn btn-md">4</button>
</div>
<div class="join">
  <button class="join-item btn btn-lg">1</button>
  <button class="join-item btn btn-lg btn-active">2</button>
  <button class="join-item btn btn-lg">3</button>
  <button class="join-item btn btn-lg">4</button>
</div>
```

---

### File Input Sizes - daisyUI

Source: https://daisyui.com/components/file-input

Provides examples of different size variants for the file input component, including extra-small (xs), small (sm), medium (md), large (lg), and extra-large (xl). This allows for flexible UI design based on content area and importance.

```html
<div>
  <input type="file" class="file-input file-input-xs" />
  <input type="file" class="file-input file-input-sm" />
  <input type="file" class="file-input file-input-md" />
  <input type="file" class="file-input file-input-lg" />
  <input type="file" class="file-input file-input-xl" />
</div>
```

---

### DaisyUI Stack Component: Stacked Cards

Source: https://daisyui.com/components/stack

Shows how to stack card elements vertically using the DaisyUI 'stack' component. This example demonstrates basic stacking without any specific alignment modifiers.

```html
<div class="stack">
  <div>A</div>
  <div>B</div>
  <div>C</div>
</div>
```

---

### Divider with No Text

Source: https://daisyui.com/components/divider

Shows how to implement a divider that appears as a simple line without any text, suitable for basic content separation. Provides examples for different visual styles.

```html
<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>
```

```html
<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>
```

```html
<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider"></div>

<div>content</div>
```

---

### DaisyUI Toast Placement: Middle-End

Source: https://daisyui.com/components/toast

Demonstrates positioning the DaisyUI Toast component to the middle-end (vertically centered, horizontally right). This example applies 'toast-middle' and 'toast-end' classes.

```html
<div class="toast toast-end toast-middle">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Text Input: Date, Time, and Datetime-Local

Source: https://daisyui.com/components/input

Provides examples of using DaisyUI styling for native HTML date, time, and datetime-local input types. These are essential for collecting specific temporal data.

````html
```html
<input type="date" class="input input-bordered w-full max-w-xs" />
<input type="time" class="input input-bordered w-full max-w-xs" />
<input type="datetime-local" class="input input-bordered w-full max-w-xs" />
````

````

--------------------------------

### DaisyUI Indicator Component: Basic Usage (HTML)

Source: https://daisyui.com/components/indicator

Demonstrates the fundamental structure for using DaisyUI indicators. The `indicator` class acts as a container, and `indicator-item` is the element placed on the corner. Placement classes like `indicator-start`, `indicator-end`, `indicator-top`, `indicator-middle`, and `indicator-bottom` control the exact position.

```html
<div class="indicator">
  <span class="indicator-item badge">+1</span>
  <div class="grid w-32 h-32 bg-base-300 place-items-center">Content</div>
</div>
````

---

### Loading Component Animations with DaisyUI

Source: https://context7_llms

Demonstrates the DaisyUI 'loading' component for displaying animations. Various styles like 'loading-dots', 'loading-spinner', and sizes from 'loading-xs' to 'loading-xl' are available.

```html
<span class="loading loading-spinner loading-lg"></span>
```

---

### URL Input Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Demonstrates how to validate URL inputs using DaisyUI's 'validator' and 'validator-hint' classes. This example ensures the entered text is a valid URL format.

```html
<input type="url" class="input input-bordered validator" placeholder="url" />
<p class="validator-hint">Must be valid URL</p>
```

---

### Divider Text Placement (Start, Default, End)

Source: https://daisyui.com/components/divider

Illustrates how to control the placement of text within a divider using classes like 'divider-start', 'divider-end', and the default centered alignment.

```html
<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>
```

```html
<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>
```

```html
<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>

<div class="divider"></div>

<div>Start</div>

<div class="divider divider-start">Default</div>

<div>End</div>
```

---

### DaisyUI Steps: With Scrollable Wrapper

Source: https://daisyui.com/components/steps

Demonstrates how to implement a scrollable wrapper for the steps component, which is useful when dealing with a large number of steps on a constrained screen.

```html
<div class="scroll-wrapper">
  <ul class="steps lg:steps-horizontal overflow-x-auto pb-4">
    <li class="step">start</li>
    <li class="step">2</li>
    <li class="step">3</li>
    <li class="step">4</li>
    <li class="step">5</li>
    <li class="step">6</li>
    <li class="step">7</li>
    <li class="step">8</li>
    <li class="step">9</li>
    <li class="step">10</li>
    <li class="step">11</li>
    <li class="step">12</li>
    <li class="step">13</li>
    <li class="step">14</li>
    <li class="step">15</li>
    <li class="step">16</li>
    <li class="step">17</li>
    <li class="step">18</li>
    <li class="step">19</li>
    <li class="step">20</li>
    <li class="step">21</li>
    <li class="step">22</li>
    <li class="step">23</li>
    <li class="step">end</li>
  </ul>
</div>
```

---

### DaisyUI Simple Hero Section (HTML)

Source: https://daisyui.com/components/hero

This code snippet presents a basic hero section with a heading and a call-to-action button. It's designed for highlighting key messages or guiding users to an important page. The structure is simple and easily customizable with DaisyUI's utility classes.

```html
### Hero with overlay image ### Hello there Provident cupiditate voluptatem et
in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti
eaque aut repudiandae et a id nisi. Get Started
```

---

### HTML structure comparison: Tailwind CSS vs. daisyUI

Source: https://daisyui.com/

This comparison illustrates the difference in HTML structure and class usage between standard Tailwind CSS and daisyUI. The daisyUI example shows significantly fewer class names for styling a similar element, emphasizing its efficiency.

````html
```html Click
<button class="btn">Tailwind only</button>

Tailwind + daisyUI
<button class="btn">Save</button>
````

````

--------------------------------

### DaisyUI Radio Tabs with Icons and Content (HTML)

Source: https://daisyui.com/components/tab

Shows an advanced usage of radio input tabs, incorporating icons within the tab labels. This example demonstrates how to visually enrich the tabs with icons alongside text.

```html
### radio tabs-lift with icons + tab content
Live Tab content 1
Laugh Tab content 2
Love Tab content 3
``` html
<div class="flex w-full flex-col gap-3">
  <div role="tablist" class="tabs tabs-lift">
    <input type="radio" name="my_tab_group_5" role="tab" class="tab" aria-label="Live" checked />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">Tab content 1</div>

    <input type="radio" name="my_tab_group_5" role="tab" class="tab" aria-label="Laugh" />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">Tab content 2</div>

    <input type="radio" name="my_tab_group_5" role="tab" class="tab" aria-label="Love" />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">Tab content 3</div>
  </div>
</div>
````

````

--------------------------------

### Phone Number Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Provides an example of validating phone numbers with DaisyUI. This implementation checks if the input consists of exactly 10 digits.

```html
<input type="tel" class="input input-bordered validator" placeholder="phone number">
<p class="validator-hint">Must be 10 digits</p>
````

---

### Dividers with Color Classes

Source: https://daisyui.com/components/divider

Presents examples of daisyUI dividers utilizing various color utility classes such as neutral, primary, secondary, accent, success, warning, info, and error.

```html
<div>Default</div>

<div class="divider"></div>

<div>Neutral</div>

<div class="divider divider-neutral"></div>

<div>Primary</div>

<div class="divider divider-primary"></div>

<div>Secondary</div>

<div class="divider divider-secondary"></div>

<div>Accent</div>

<div class="divider divider-accent"></div>

<div>Success</div>

<div class="divider divider-success"></div>

<div>Warning</div>

<div class="divider divider-warning"></div>

<div>Info</div>

<div class="divider divider-info"></div>

<div>Error</div>

<div class="divider divider-error"></div>
```

```html
<div>Default</div>

<div class="divider"></div>

<div>Neutral</div>

<div class="divider divider-neutral"></div>

<div>Primary</div>

<div class="divider divider-primary"></div>

<div>Secondary</div>

<div class="divider divider-secondary"></div>

<div>Accent</div>

<div class="divider divider-accent"></div>

<div>Success</div>

<div class="divider divider-success"></div>

<div>Warning</div>

<div class="divider divider-warning"></div>

<div>Info</div>

<div class="divider divider-info"></div>

<div>Error</div>

<div class="divider divider-error"></div>
```

```html
<div>Default</div>

<div class="divider"></div>

<div>Neutral</div>

<div class="divider divider-neutral"></div>

<div>Primary</div>

<div class="divider divider-primary"></div>

<div>Secondary</div>

<div class="divider divider-secondary"></div>

<div>Accent</div>

<div class="divider divider-accent"></div>

<div>Success</div>

<div class="divider divider-success"></div>

<div>Warning</div>

<div class="divider divider-warning"></div>

<div>Info</div>

<div class="divider divider-info"></div>

<div>Error</div>

<div class="divider divider-error"></div>
```

```html
<div>Default</div>

<div class="divider"></div>

<div>Neutral</div>

<div class="divider divider-neutral"></div>

<div>Primary</div>

<div class="divider divider-primary"></div>

<div>Secondary</div>

<div class="divider divider-secondary"></div>

<div>Accent</div>

<div class="divider divider-accent"></div>

<div>Success</div>

<div class="divider divider-success"></div>

<div>Warning</div>

<div class="divider divider-warning"></div>

<div>Info</div>

<div class="divider divider-info"></div>

<div>Error</div>

<div class="divider divider-error"></div>
```

---

### DaisyUI Phone Mockup Syntax (HTML)

Source: https://context7_llms

Provides the HTML structure for a phone mockup using DaisyUI. It includes separate divs for the camera and the display area, where content can be placed.

```html
<div class="mockup-phone">
  <div class="mockup-phone-camera"></div>
  <div class="mockup-phone-display">
    <!-- Content inside the phone display -->
  </div>
</div>
```

---

### DaisyUI Radio Tabs with Border and Content (HTML)

Source: https://daisyui.com/components/tab

Combines radio input-based tab functionality with the 'tabs-border' style. This example shows how to manage tab content visibility based on the selected radio input.

````html
### radio tabs-border + tab content Tab content 1 Tab content 2 Tab content 3
``` html
<div class="flex w-full flex-col gap-3">
  <div role="tablist" class="tabs tabs-border">
    <input
      type="radio"
      name="my_tab_group_3"
      role="tab"
      class="tab"
      aria-label="Tab 1"
      checked
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 1
    </div>

    <input
      type="radio"
      name="my_tab_group_3"
      role="tab"
      class="tab"
      aria-label="Tab 2"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 2
    </div>

    <input
      type="radio"
      name="my_tab_group_3"
      role="tab"
      class="tab"
      aria-label="Tab 3"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 3
    </div>
  </div>
</div>
````

````

--------------------------------

### Ghost Textarea Style - daisyUI

Source: https://daisyui.com/components/textarea

Illustrates the 'ghost' style for the Textarea component in daisyUI, which typically removes the background for a more minimalist appearance. Multiple examples showcase this style.

```html
```html
<textarea class="textarea textarea-ghost" placeholder="Bio"></textarea>
````

````

--------------------------------

### DaisyUI Outline Button Styles

Source: https://daisyui.com/components/button

Presents the 'outline' button style, characterized by a border and transparent background. This style is achieved with the 'btn-outline' class, and examples include different color variations.

```html
<button class="btn btn-outline">Default</button>
<button class="btn btn-outline btn-primary">Primary</button>
<button class="btn btn-outline btn-secondary">Secondary</button>
<button class="btn btn-outline btn-accent">Accent</button>
<button class="btn btn-outline btn-info">Info</button>
<button class="btn btn-outline btn-success">Success</button>
<button class="btn btn-outline btn-warning">Warning</button>
<button class="btn btn-outline btn-error">Error</button>
````

---

### Textarea with Form Control and Labels - daisyUI

Source: https://daisyui.com/components/textarea

Shows how to integrate the Textarea component within a form, including associated labels and optional indicators. This example demonstrates accessibility and usability best practices.

````html
```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">Your bio</span>
    <span class="label-text-alt">Optional</span>
  </label>
  <textarea
    class="textarea textarea-bordered w-full max-w-xs"
    placeholder="Bio"
  ></textarea>
  <label class="label">
    <span class="label-text-alt">Alt</span>
    <span class="label-text-alt">Optional</span>
  </label>
</div>
````

````

--------------------------------

### DaisyUI Swap Component: Activate Using Class Name (HTML)

Source: https://daisyui.com/components/swap

Explains and demonstrates how to activate the DaisyUI Swap component using a class name (`swap-active`) instead of a checkbox. This method allows for JavaScript-driven state changes, useful for dynamic UI interactions. Examples show toggling between different emojis.

```html
<div class="swap">
  <div class="swap-active">🥵</div>
  <div class="swap-inactive">🥶</div>
</div>

<div class="swap swap-active">
  <div class="swap-on">🥳</div>
  <div class="swap-off">😭</div>
</div>
````

---

### Username Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Illustrates the implementation of username validation rules using DaisyUI. This example enforces a length between 3 to 30 characters and allows only letters, numbers, or dashes.

```html
<input
  type="text"
  class="input input-bordered validator"
  placeholder="username"
/>
<p class="validator-hint">
  Must be 3 to 30 characters<br />
  containing only letters, numbers or dash
</p>
```

---

### Radial Progress with Custom Color (HTML)

Source: https://daisyui.com/components/radial-progress

Illustrates how to apply a custom color to the radial-progress component by setting the `--progress-color` CSS variable. The example shows a 70% progress with a custom color.

```html
<div
  class="radial-progress"
  style="--value:70; --progress-color:oklch(var(--p))"
  role="progressbar"
>
  70%
</div>
```

---

### Radial Progress with Background and Border (React TSX)

Source: https://daisyui.com/components/radial-progress

A React (TSX) example for styling the radial-progress component with a background color and border. This enhances the visual appearance of the progress indicator.

```tsx
function RadialProgressBackgroundBorderTsx() {
  return (
    <div
      className="radial-progress bg-base-200 border border-base-300"
      style={{ "--value": 70 } as React.CSSProperties}
      role="progressbar"
    >
      70%
    </div>
  );
}
```

---

### DaisyUI Join Component: With Extra Elements

Source: https://daisyui.com/components/join

Demonstrates that the 'join-item' styling in DaisyUI applies even when the item is not a direct child of the 'join' container. This example shows grouping inputs and buttons with text.

```html
<div class="join">
  <input class="join-item input input-bordered" placeholder="Filter" />
  <button class="join-item btn">Search</button>
</div>
<div class="join">
  <input class="join-item input input-bordered" placeholder="Filter" />
  <button class="join-item btn">Search</button>
</div>
<div class="join">
  <input class="join-item input input-bordered" placeholder="Filter" />
  <button class="join-item btn">Search</button>
</div>
<div class="join">
  <input class="join-item input input-bordered" placeholder="Filter" />
  <button class="join-item btn">Search</button>
</div>
```

---

### DaisyUI Pagination Syntax (HTML)

Source: https://context7_llms

Shows the DaisyUI pagination component structure. Each item within the pagination should have the `join-item` class. This component is built using the `join` class.

```html
<div class="join">
  <button class="join-item btn btn-active">1</button>
  <button class="join-item btn">2</button>
  <button class="join-item btn">3</button>
  <button class="join-item btn">4</button>
</div>
```

---

### Code Mockup Without Prefix in daisyUI

Source: https://daisyui.com/components/mockup-code

Illustrates the 'mockup-code' component without a line prefix. This is useful for displaying code snippets where a command prompt or path is not relevant.

```html
<div class="mockup-code">
  <code>without prefix</code>
</div>
```

---

### Horizontal Divider with Text

Source: https://daisyui.com/components/divider

Illustrates how to use the horizontal divider with text content, showing different ways to position the text within the divider line. Includes examples with 'OR' separators.

```html
<div>content</div>

<div class="divider">OR</div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider">OR</div>

<div>content</div>
```

```html
<div>content</div>

<div class="divider">OR</div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider">OR</div>

<div>content</div>
```

```html
<div>content</div>

<div class="divider">OR</div>

<div>content</div>

<div class="divider"></div>

<div>content</div>

<div class="divider">OR</div>

<div>content</div>
```

---

### Date Input Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Shows an example of date input validation with DaisyUI, specifically checking if the input matches the year 2025. This can be adapted for other date-related validation rules.

```html
<input type="text" class="input input-bordered validator" placeholder="date" />
<p class="validator-hint">Must be 2025</p>
```

---

### Browser Mockup with Background Color - HTML

Source: https://daisyui.com/components/mockup-browser

Shows how to implement a browser mockup with a specified background color using daisyUI. This allows for more design flexibility when presenting web interfaces.

```html
<div class="mockup-browser bg-primary">
  <div class="mockup-browser-toolbar">
    <input
      class="input input-bordered input-primary"
      value="https://daisyui.com"
    />
  </div>
  <div class="flex justify-center px-4 py-16 bg-base-200">Hello!</div>
</div>
```

---

### DaisyUI Card with Centered Content

Source: https://daisyui.com/components/card

This example illustrates a DaisyUI card with its content centered and includes padding. It features an image at the top, a title, descriptive text, and a call-to-action button. This card is suitable for promotional content or featured items.

```html
<div class="card w-96 image-full shadow-xl">
  <figure>
    <img
      src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
      alt="Shoes"
    />
  </figure>
  <div class="card-body">
    <h2 class="card-title">Card Title</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
    <div class="card-actions justify-center">
      <button class="btn btn-primary">Buy Now</button>
    </div>
  </div>
</div>
```

---

### Basic DaisyUI Table Structure

Source: https://daisyui.com/components/table

Demonstrates the fundamental structure for creating a table using DaisyUI classes. This includes the main table element and header/row definitions. No external dependencies are required beyond DaisyUI.

```html
<div class="overflow-x-auto">
  <table class="table">
    <!-- head -->
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
        <th>Favorite Color</th>
      </tr>
    </thead>
    <tbody>
      <!-- row 1 -->
      <tr>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
      </tr>
      <!-- row 2 -->
      <tr>
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
      </tr>
      <!-- row 3 -->
      <tr>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
      </tr>
    </tbody>
  </table>
</div>
```

---

### DaisyUI Text Input: Username with Icon and Validator

Source: https://daisyui.com/components/input

Illustrates a username input field with an accompanying icon and validation message using DaisyUI classes. This example enforces specific character and length constraints for usernames.

````html
```html
<label class="input input-bordered flex items-center gap-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 16 16"
    fill="currentColor"
    class="w-4 h-4 opacity-70"
  >
    <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z" />
    <path
      fill-rule="evenodd"
      d="M13.682 2.318A9.887 9.887 0 0 0 8 1c-2.87 0-5.432 1.772-6.735 4.464a9.887 9.887 0 0 0 2.034 1.964C3.787 7.712 5.474 9 8 9c2.526 0 4.213-1.288 5.307-2.722a9.887 9.887 0 0 0 2.034-1.964A9.875 9.875 0 0 0 13.682 2.318Z"
      clip-rule="evenodd"
    />
  </svg>
  <input type="text" class="grow" placeholder="Username" />
</label>
<label class="label">
  <span class="label-text-alt"
    >Must be 3 to 30 characters containing only letters, numbers or dash</span
  >
</label>
````

````

--------------------------------

### DaisyUI Sticky Rows and Columns Table

Source: https://daisyui.com/components/table

Demonstrates how to create tables with sticky rows and columns using `table-pin-rows` and `table-pin-cols` modifiers. This is useful for large tables where headers or specific columns need to remain visible during scrolling. Requires basic HTML table setup.

```html
<div class="overflow-x-auto">
  <table class="table table-pin-rows table-pin-cols">
    <!-- head -->
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
        <th>Favorite Color</th>
      </tr>
    </thead>
    <tbody>
      <!-- row 1 -->
      <tr>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
      </tr>
      <!-- row 2 -->
      <tr>
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
      </tr>
      <!-- row 3 -->
      <tr>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
      </tr>
    </tbody>
  </table>
</div>
````

---

### Colored Code Mockup with daisyUI

Source: https://daisyui.com/components/mockup-code

Shows how to apply custom colors to the 'mockup-code' component. This allows for visual customization to match theme requirements or highlight specific code blocks.

```html
<div class="mockup-code" data-theme="emerald">
  <code>can be any color!</code>
</div>
```

---

### DaisyUI Rating with Hidden Input

Source: https://daisyui.com/components/rating

This example shows how to use the 'rating-hidden' modifier. A hidden radio input at the beginning allows the user to clear their selection, providing a 'no rating' option.

```html
<div class="rating rating-hidden ">
  <input type="radio" name="rating-hidden" class="mask mask-star-2" />
  <input type="radio" name="rating-hidden" class="mask mask-star-2" checked />
  <input type="radio" name="rating-hidden" class="mask mask-star-2" />
  <input type="radio" name="rating-hidden" class="mask mask-star-2" />
  <input type="radio" name="rating-hidden" class="mask mask-star-2" />
</div>
```

---

### Radial Progress with Different Values (Vue)

Source: https://daisyui.com/components/radial-progress

This Vue.js example displays the radial-progress component at different percentage values, showing how to dynamically set the '--value' CSS variable for various progress states.

```vue
<template>
  <div class="radial-progress" :style="{ '--value': 0 }" role="progressbar">
    0%
  </div>
  <div class="radial-progress" :style="{ '--value': 20 }" role="progressbar">
    20%
  </div>
  <div class="radial-progress" :style="{ '--value': 60 }" role="progressbar">
    60%
  </div>
  <div class="radial-progress" :style="{ '--value': 80 }" role="progressbar">
    80%
  </div>
  <div class="radial-progress" :style="{ '--value': 100 }" role="progressbar">
    100%
  </div>
</template>

<script setup lang="ts">
// No script needed for basic example
</script>

<style>
/* DaisyUI styles are assumed to be globally imported */
</style>
```

---

### Password Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Shows how to implement password validation requirements using DaisyUI's 'validator' and 'validator-hint' classes. The example includes checks for length, numbers, lowercase, and uppercase letters.

```html
<input
  type="password"
  class="input input-bordered validator"
  placeholder="password"
/>
<p class="validator-hint">
  Must be more than 8 characters, including<br />
  At least one number<br />
  At least one lowercase letter<br />
  At least one uppercase letter
</p>
```

---

### Basic Radial Progress (HTML)

Source: https://daisyui.com/components/radial-progress

Demonstrates the basic usage of the radial-progress component using HTML and CSS variables. It requires the '--value' CSS variable to function and can be styled further with '--size' and '--thickness'.

```html
<div class="radial-progress" style="--value:70" role="progressbar">70%</div>
```

---

### Select Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Demonstrates validating a select dropdown using DaisyUI. The example includes instructions to click the button before picking an option to see the error color, indicating a required field.

```html
<select class="select select-bordered validator">
  <option disabled selected>Choose:</option>
  <option>Tabs</option>
  <option>Spaces</option>
</select>
<p class="validator-hint">Required</p>
```

---

### DaisyUI Soft Button Styles

Source: https://daisyui.com/components/button

Demonstrates the 'soft' button style, which provides a lighter, less prominent appearance compared to default buttons. This style is applied using the 'btn-soft' class, with examples shown for various color options.

```html
<button class="btn btn-soft">Default</button>
<button class="btn btn-soft btn-primary">Primary</button>
<button class="btn btn-soft btn-secondary">Secondary</button>
<button class="btn btn-soft btn-accent">Accent</button>
<button class="btn btn-soft btn-info">Info</button>
<button class="btn btn-soft btn-success">Success</button>
<button class="btn btn-soft btn-warning">Warning</button>
<button class="btn btn-soft btn-error">Error</button>
```

---

### DaisyUI Swap Component: Hamburger Button Toggle (HTML)

Source: https://daisyui.com/components/swap

Provides an example of creating a hamburger button that transforms into a close icon upon clicking, using the DaisyUI Swap component. The `swap` class is used, with a hidden checkbox controlling the state change between the hamburger and close icons.

```html
<label class="swap">
  <input type="checkbox" />
  <svg
    class="swap-on fill-current"
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path d="M3 6h18M3 12h18M3 18h18"></path>
  </svg>
  <svg
    class="swap-off fill-current"
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path d="M18 6L6 18M6 6l12 12"></path>
  </svg>
</label>
```

---

### Dock Component Syntax (HTML)

Source: https://context7_llms

Illustrates the HTML structure for a dock component, which functions as a bottom navigation bar. Content is typically a list of buttons with labels.

```html
<div class="dock {MODIFIER}">{CONTENT}</div>
```

```html
<button>
  <svg>{icon}</svg>
  <span class="dock-label">Text</span>
</button>
```

---

### Browser Mockup with Border - HTML

Source: https://daisyui.com/components/mockup-browser

Demonstrates how to create a browser mockup with a border using daisyUI's HTML structure. This is useful for showcasing web content within a realistic browser frame.

```html
<div class="mockup-browser border border-base-300">
  <div class="mockup-browser-toolbar">
    <input class="input input-bordered" value="https://daisyui.com" />
  </div>
  <div class="flex justify-center px-4 py-16 bg-base-200">Hello!</div>
</div>
```

---

### DaisyUI Join Component: Custom Border Radius

Source: https://daisyui.com/components/join

Illustrates how to apply custom border radius to elements within the DaisyUI 'join' component. This example shows applying a rounded class to the 'join' container itself.

```html
<div class="join rounded-lg">
  <button class="join-item btn">Subscribe</button>
</div>
<div class="join rounded-lg">
  <button class="join-item btn">Subscribe</button>
</div>
<div class="join rounded-lg">
  <button class="join-item btn">Subscribe</button>
</div>
<div class="join rounded-lg">
  <button class="join-item btn">Subscribe</button>
</div>
```

---

### DaisyUI Rating with Green Star Icon

Source: https://daisyui.com/components/rating

This example showcases a rating component using the 'mask-star-2' icon and applying a specific green color ('green-500'). This allows for custom branding or thematic consistency in the UI.

```html
<div class="rating">
  <input type="radio" name="rating-green-star" class="mask mask-star-2" />
  <input
    type="radio"
    name="rating-green-star"
    class="mask mask-star-2"
    checked
  />
  <input type="radio" name="rating-green-star" class="mask mask-star-2" />
  <input type="radio" name="rating-green-star" class="mask mask-star-2" />
  <input type="radio" name="rating-green-star" class="mask mask-star-2" />
</div>
<div class="rating">
  <input
    type="radio"
    name="rating-green-half"
    class="bg-green-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-green-half"
    class="bg-green-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-green-half"
    class="bg-green-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-green-half"
    class="bg-green-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-green-half"
    class="bg-green-500 mask mask-star-2"
    checked
  />
</div>
```

---

### Basic DaisyUI Card Component

Source: https://daisyui.com/components/card

A fundamental card component structure in DaisyUI. It typically includes a figure for an image or icon, a body section for text content, and distinct areas for titles and actions. This example demonstrates the basic markup.

```html
<div class="card w-96 bg-base-100 shadow-xl">
  <figure>
    <img
      src="https://img.daisyui.com/images/stock/photo-1606107557116-37ef5577107a.webp"
      alt="Shoes"
    />
  </figure>
  <div class="card-body">
    <h2 class="card-title">Card title!</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Buy Now</button>
    </div>
  </div>
</div>
```

---

### DaisyUI Card with Bottom Image Example

Source: https://daisyui.com/components/card

This code snippet demonstrates a basic DaisyUI card component featuring an image positioned at the bottom. It includes a title and descriptive text within the card body. This component is useful for displaying product information or articles.

```html
<div class="card w-96 bg-base-100 shadow-xl">
  <div class="card-body">
    <h2 class="card-title">Card Title</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
  </div>
  <figure>
    <img
      src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
      alt="Shoes"
    />
  </figure>
</div>
```

---

### Basic Textarea Component - daisyUI

Source: https://daisyui.com/components/textarea

Demonstrates the fundamental usage of the Textarea component in daisyUI. It shows a simple textarea element with a placeholder, ready for user input.

```html
<textarea class="textarea" placeholder="Bio"></textarea>
```

---

### DaisyUI Card with Custom Color

Source: https://daisyui.com/components/card

This example showcases a DaisyUI card with a custom background color, indicated by 'bg-primary'. It includes a title, descriptive text, and a 'Buy Now' button. Customizing colors allows for better theme integration and visual hierarchy.

```html
<div class="card w-96 bg-primary text-primary-content shadow-xl">
  <div class="card-body">
    <h2 class="card-title">Card title!</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
    <div class="card-actions justify-end">
      <button class="btn">Buy Now</button>
    </div>
  </div>
</div>
```

---

### Basic Checkbox Usage (HTML)

Source: https://daisyui.com/components/checkbox

Demonstrates the basic HTML structure for a simple checkbox using daisyUI classes. No external dependencies are required beyond daisyUI.

````html
```html <input type="checkbox" class="checkbox" />
````

````

--------------------------------

### Floating Label with Different Sizes (HTML/Tailwind CSS)

Source: https://daisyui.com/components/label

Presents examples of floating labels with various input sizes supported by daisyUI, including Extra Small, Small, Medium, Large, and Extra Large. This allows for flexible form design according to the application's needs.

```html
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-xs input-bordered peer" placeholder=" " />
  <label class="label absolute top-0 left-0 ml-2 px-1 text-xs peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300">
    <span class="label-text">Extra Small</span>
  </label>
</div>
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-sm input-bordered peer" placeholder=" " />
  <label class="label absolute top-0 left-0 ml-2 px-1 text-xs peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300">
    <span class="label-text">Small</span>
  </label>
</div>
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-md input-bordered peer" placeholder=" " />
  <label class="label absolute top-0 left-0 ml-2 px-1 text-xs peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300">
    <span class="label-text">Medium</span>
  </label>
</div>
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-lg input-bordered peer" placeholder=" " />
  <label class="label absolute top-0 left-0 ml-2 px-1 text-xs peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300">
    <span class="label-text">Large</span>
  </label>
</div>
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-xl input-bordered peer" placeholder=" " />
  <label class="label absolute top-0 left-0 ml-2 px-1 text-xs peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300">
    <span class="label-text">Extra Large</span>
  </label>
</div>
````

---

### Radial Progress with Different Values (React TSX)

Source: https://daisyui.com/components/radial-progress

A React (TSX) example showcasing the radial-progress component with various percentage values. This demonstrates how to dynamically set the '--value' CSS variable for different progress states.

```tsx
function RadialProgressDifferentValuesTsx() {
  return (
    <>
      <div
        className="radial-progress"
        style={{ "--value": 0 } as React.CSSProperties}
        role="progressbar"
      >
        0%
      </div>
      <div
        className="radial-progress"
        style={{ "--value": 20 } as React.CSSProperties}
        role="progressbar"
      >
        20%
      </div>
      <div
        className="radial-progress"
        style={{ "--value": 60 } as React.CSSProperties}
        role="progressbar"
      >
        60%
      </div>
      <div
        className="radial-progress"
        style={{ "--value": 80 } as React.CSSProperties}
        role="progressbar"
      >
        80%
      </div>
      <div
        className="radial-progress"
        style={{ "--value": 100 } as React.CSSProperties}
        role="progressbar"
      >
        100%
      </div>
    </>
  );
}
```

---

### DaisyUI Navbar Syntax (HTML)

Source: https://context7_llms

Demonstrates the basic structure of a DaisyUI navbar. Content can be positioned within the navbar using `navbar-start`, `navbar-center`, and `navbar-end` classes.

```html
<div class="navbar bg-base-100">
  <div class="navbar-start">
    <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
  </div>
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1">
      <li><a>Link 1</a></li>
      <li><a>Link 2</a></li>
    </ul>
  </div>
  <div class="navbar-end">
    <button class="btn btn-ghost btn-circle">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </button>
  </div>
</div>
```

---

### Update tailwind.config.js for Tailwind CSS 4 and daisyUI 5

Source: https://daisyui.com/docs/upgrade

This snippet shows how to modify the tailwind.config.js file to remove daisyUI themes and plugins before running the Tailwind CSS upgrade tool, and how to re-add daisyUI with its themes after the upgrade.

```javascript
module.exports = {
  content: ["./your-files/**/*.{html,js}"],
  // other stuff...
  // - daisyui: {
  // - themes: ['light', 'dark', 'cupcake'],
  // - },
  // - plugins: [require("daisyui")],
};
```

```javascript
module.exports = {
  content: ["./your-files/**/*.{html,js}"],
  // other stuff...
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/colors").light,
          primary: "#007bff",
        },
      },
      "dark",
      "cupcake",
    ],
  },
};
```

---

### Filter without HTML Form - daisyUI

Source: https://daisyui.com/components/filter

This example shows how to create a filter interface without using a full HTML form. This approach is useful when form submission is not required or when integrating the filter into a different UI structure. It still utilizes radio buttons for selection.

```html
<div class="filter">
  <input
    type="radio"
    name="filter-group"
    id="optionA"
    class="filter-radio"
    checked
  />
  <label for="optionA">Option A</label>

  <input type="radio" name="filter-group" id="optionB" class="filter-radio" />
  <label for="optionB">Option B</label>

  <input type="radio" name="filter-group" id="optionC" class="filter-radio" />
  <label for="optionC">Option C</label>

  <button class="filter-reset">Reset</button>
</div>
```

---

### DaisyUI Dash Button Styles

Source: https://daisyui.com/components/button

Features the 'dash' button style, which uses a dashed border for a distinct look. This style is applied using the 'btn-dash' class, with examples showing its integration with various color options.

```html
<button class="btn btn-dash">Default</button>
<button class="btn btn-dash btn-primary">Primary</button>
<button class="btn btn-dash btn-secondary">Secondary</button>
<button class="btn btn-dash btn-accent">Accent</button>
<button class="btn btn-dash btn-info">Info</button>
<button class="btn btn-dash btn-success">Success</button>
<button class="btn btn-dash btn-warning">Warning</button>
<button class="btn btn-dash btn-error">Error</button>
```

---

### Label for Select Element (HTML/Tailwind CSS)

Source: https://daisyui.com/components/label

Illustrates using the daisyUI Label component with a select dropdown. This allows for clear labeling of select options and enhances user experience. Examples include different types of select inputs.

```html
<label class="label">
  <span class="label-text">Choose your favorite</span>
</label>
<select class="select select-bordered w-full max-w-xs">
  <option disabled selected>Pick one</option>
  <option>Star Wars</option>
  <option>Harry Potter</option>
  <option>Lord of the Rings</option>
  <option>Planet of the Apes</option>
  <option>The Matrix</option>
</select>
```

```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">Personal or Business</span>
  </label>
  <select class="select select-bordered">
    <option>Personal</option>
    <option>Business</option>
  </select>
</div>
```

---

### Basic DaisyUI Menu Component

Source: https://daisyui.com/components/menu

Demonstrates the basic structure of a DaisyUI menu with simple list items. This serves as a foundation for more complex menu configurations.

```html
<ul class="menu">
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>
```

---

### Diff Component Syntax (HTML)

Source: https://context7_llms

Demonstrates the structure for a diff component, used for comparing two items side-by-side. Aspect ratio classes can be added to maintain proportions.

```html
<figure class="diff">
  <div class="diff-item-1">{item1}</div>
  <div class="diff-item-2">{item2}</div>
  <div class="diff-resizer"></div>
</figure>
```

---

### DaisyUI Table with Active Row

Source: https://daisyui.com/components/table

Shows how to highlight an active row within a DaisyUI table. This is typically achieved by applying a specific class to the desired row element, visually indicating selection or focus. The example uses standard HTML table rows.

```html
<div class="overflow-x-auto">
  <table class="table">
    <!-- head -->
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
        <th>Favorite Color</th>
      </tr>
    </thead>
    <tbody>
      <!-- row 1 -->
      <tr>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
      </tr>
      <!-- row 2 - Active -->
      <tr class="active">
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
      </tr>
      <!-- row 3 -->
      <tr>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
      </tr>
    </tbody>
  </table>
</div>
```

---

### Floating Label Input (HTML/Tailwind CSS)

Source: https://daisyui.com/components/label

Demonstrates the 'floating-label' pattern in daisyUI, where the label floats above the input field when it's focused or has content. This provides a clean and modern UI for forms. Examples include different input sizes.

```html
<div class="form-control relative">
  <input type="text" class="input input-bordered peer" placeholder=" " />
  <label
    class="label absolute top-0 left-0 ml-3 px-2 peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 peer-hover:text-xs peer-hover:-top-2 transition-all duration-300"
  >
    <span class="label-text">Your Email</span>
  </label>
</div>
```

```html
<div class="form-control relative w-full max-w-xs">
  <input type="text" class="input input-bordered peer" placeholder=" " />
  <label
    class="label absolute top-0 left-0 ml-3 px-2 peer-focus:text-xs peer-focus:ml-1 peer-focus:-top-2 transition-all duration-300"
  >
    <span class="label-text">Username</span>
  </label>
</div>
```

---

### DaisyUI Responsive Button Sizes

Source: https://daisyui.com/components/button

Illustrates how to create responsive buttons that adapt their size based on the viewport. This uses DaisyUI's responsive prefixes to define different button sizes for various screen widths, enabling adaptive UI design.

```html
<button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-lg xl:btn-xl">
  Responsive
</button>
```

---

### Textarea Component Color Variations - daisyUI

Source: https://daisyui.com/components/textarea

Presents various color options for the Textarea component in daisyUI, including neutral, primary, secondary, accent, info, success, warning, and error states. Each example applies a specific color class.

````html
```html
<textarea class="textarea textarea-neutral" placeholder="Bio"></textarea>
<textarea class="textarea textarea-primary" placeholder="Bio"></textarea>
<textarea class="textarea textarea-secondary" placeholder="Bio"></textarea>
<textarea class="textarea textarea-accent" placeholder="Bio"></textarea>
<textarea class="textarea textarea-info" placeholder="Bio"></textarea>
<textarea class="textarea textarea-success" placeholder="Bio"></textarea>
<textarea class="textarea textarea-warning" placeholder="Bio"></textarea>
<textarea class="textarea textarea-error" placeholder="Bio"></textarea>
````

````

--------------------------------

### Filter Component with HTML Form in HTML

Source: https://context7_llms

The filter component groups radio buttons, allowing only one selection at a time and displaying a reset button. This example shows its implementation using an HTML form, with classes for the component and the reset button.

```html
<form class="filter">
  <input class="btn btn-square" type="reset" value="×"/>
  <input class="btn" type="radio" name="my-radio" aria-label="Tab 1 title"/>
  <input class="btn" type="radio" name="my-radio" aria-label="Tab 2 title"/>
</form>
````

---

### daisyUI 5: Importing Library Parts via Plugin

Source: https://daisyui.com/docs/v5

Shows the syntax for importing specific parts of the daisyUI library using a plugin, allowing for modular inclusion and exclusion of components. This is a new feature in daisyUI 5 for enhanced customization.

```css
@plugin "daisyui";
```

---

### DaisyUI Dropdown Syntax (Details/Summary)

Source: https://context7_llms

Implementation of a DaisyUI Dropdown using the HTML `<details>` and `<summary>` elements. The `<details>` element has the `dropdown` class, and the dropdown content is placed within a `<ul>` with the `dropdown-content` class. This method leverages native HTML semantics for dropdown functionality.

```html
<details class="dropdown">
  <summary>Button</summary>
  <ul class="dropdown-content">
    {CONTENT}
  </ul>
</details>
```

---

### Basic File Input - daisyUI

Source: https://daisyui.com/components/file-input

Demonstrates the fundamental usage of the file input component in daisyUI. This code snippet shows a simple file input field without any specific styling or variations.

```html
<input type="file" class="file-input" />
```

---

### FAB and Speed Dial (Vertical) - With SVG Icons

Source: https://daisyui.com/components/fab

Illustrates how to integrate SVG icons into the FAB and Speed Dial buttons. This example shows the structure for both the main FAB button and the individual speed dial buttons, incorporating SVG elements for visual representation.

```html
<div class="fab">
  <button class="fab-main-action">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-6 w-6"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M4 6h16M4 12h16M4 18h16"
      />
    </svg>
  </button>
  <div class="fab-content">
    <button class="btn btn-circle btn-xs">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-4 w-4"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </button>
    <button class="btn btn-circle btn-xs">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-4 w-4"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </button>
    <button class="btn btn-circle btn-xs">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-4 w-4"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </button>
  </div>
</div>
```

---

### DaisyUI Modal with Hidden Checkbox

Source: https://daisyui.com/components/modal

This example shows how to implement a modal using a hidden checkbox input. Clicking a label associated with the checkbox toggles the modal's visibility. This method provides a pure HTML/CSS approach to modal control.

```html
<label for="my-modal" class="btn">open modal</label>

<input type="checkbox" id="my-modal" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with a hidden checkbox!</p>
    <div class="modal-action">
      <label for="my-modal" class="btn">Close</label>
    </div>
  </div>
</div>
```

---

### DaisyUI Join Component: Responsive Layout

Source: https://daisyui.com/components/join

Shows how to achieve a responsive layout with the DaisyUI 'join' component. The items are displayed vertically on small screens and switch to a horizontal layout on larger screens, controlled by utility classes.

```html
<div class="join lg:join-vertical ">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join lg:join-vertical ">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join lg:join-vertical ">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join lg:join-vertical ">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join lg:join-vertical ">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
```

---

### Label for Input Field (HTML/Tailwind CSS)

Source: https://daisyui.com/components/label

Demonstrates how to use the daisyUI Label component to associate text with an input field. This setup ensures proper accessibility and styling for form elements. It covers basic input fields and input fields with labels placed at the end.

```html
<label class="label">
  <span class="label-text">Your email</span>
</label>
<input
  type="text"
  placeholder="Type here"
  class="input input-bordered w-full max-w-xs"
/>
```

```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">What's in your pocket?</span>
  </label>
  <input
    type="text"
    placeholder="Type here"
    class="input input-bordered w-full max-w-xs"
  />
  <label class="label">
    <span class="label-text-alt">Alternative label</span>
  </label>
</div>
```

---

### daisyUI 5: Micro CSS File for Toggle Component

Source: https://daisyui.com/docs/v5

Demonstrates how to include only the CSS for the toggle component from daisyUI's CDN for projects without a build step, significantly reducing the amount of CSS loaded. This showcases the 'Micro CSS files' feature for no-build projects.

```html
<link
  href="https://cdn.jsdelivr.net/npm/daisyui@5/components/toggle.css"
  rel="stylesheet"
  type="text/css"
/>
```

---

### DaisyUI Select Component Basic Usage

Source: https://daisyui.com/components/select

Demonstrates the basic HTML structure for a DaisyUI select component, allowing users to pick a color from a predefined list. It uses standard HTML select elements with DaisyUI classes for styling.

```html
<label class="form-control w-full max-w-xs">
  <div class="label">
    <span class="label-text">Pick a color</span>
  </div>
  <select class="select select-bordered w-full max-w-xs">
    <option disabled selected>Crimson</option>
    <option>Amber</option>
    <option>Velvet</option>
  </select>
</label>
```

---

### FAB and Speed Dial (Vertical) - With Rectangle Buttons

Source: https://daisyui.com/components/fab

Demonstrates using standard rectangle buttons instead of circular buttons for the speed dial. This example shows the FAB structure with the main action button and the content area containing rectangular speed dial buttons.

```html
<div class="fab">
  <button class="fab-main-action">F</button>
  <div class="fab-content">
    <button class="btn btn-xs">Button A</button>
    <button class="btn btn-xs">Button B</button>
    <button class="btn btn-xs">Button C</button>
  </div>
</div>
```

---

### Breadcrumbs with Max-Width and Scrolling - daisyUI

Source: https://daisyui.com/components/breadcrumbs

This example illustrates how to implement breadcrumbs that handle long text or a large number of items by applying a maximum width and enabling horizontal scrolling. This prevents layout issues on smaller screens or with extensive navigation paths. The container must have a defined width for the scrolling to function.

```html
<div class="max-w-xs ">
  <div role="Breadcrumb" class="breadcrumbs">
    <ul>
      <li>
        <a href="#">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="w-4 h-4 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6a2 2 0 00-2-2H5a2 2 0 00-2 2z"
            ></path>
          </svg>
          Long text 1
        </a>
      </li>
      <li>
        <a href="#">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="w-4 h-4 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 13h6m-3-3v6m5 5H9a2 2 0 01-2-2V5a2 2 0 012-2h5.5a2 2 0 011.54.535M9 22V7h5.5a2 2 0 011.548.535l3.452 3.452a2 2 0 01.547 1.548L21 17a2 2 0 01-2 2H9z"
            ></path>
          </svg>
          Long text 2
        </a>
      </li>
      <li>
        <a href="#">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="w-4 h-4 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 13h6m-3-3v6m5 5H9a2 2 0 01-2-2V5a2 2 0 012-2h5.5a2 2 0 011.54.535M9 22V7h5.5a2 2 0 011.548.535l3.452 3.452a2 2 0 01.547 1.548L21 17a2 2 0 01-2 2H9z"
            ></path>
          </svg>
          Long text 3
        </a>
      </li>
      <li>
        <a href="#">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="w-4 h-4 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 13h6m-3-3v6m5 5H9a2 2 0 01-2-2V5a2 2 0 012-2h5.5a2 2 0 011.54.535M9 22V7h5.5a2 2 0 011.548.535l3.452 3.452a2 2 0 01.547 1.548L21 17a2 2 0 01-2 2H9z"
            ></path>
          </svg>
          Long text 4
        </a>
      </li>
      <li>
        <a href="#">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="w-4 h-4 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 13h6m-3-3v6m5 5H9a2 2 0 01-2-2V5a2 2 0 012-2h5.5a2 2 0 011.54.535M9 22V7h5.5a2 2 0 011.548.535l3.452 3.452a2 2 0 01.547 1.548L21 17a2 2 0 01-2 2H9z"
            ></path>
          </svg>
          Long text 5
        </a>
      </li>
    </ul>
  </div>
</div>
```

---

### DaisyUI Pagination using Radio Inputs

Source: https://daisyui.com/components/pagination

Demonstrates using radio inputs within the DaisyUI pagination component. This approach is useful for scenarios where only one page can be selected at a time, offering a different interaction model.

```html
<div class="join">
  <input
    type="radio"
    name="options"
    aria-label="1"
    class="join-item btn btn-square"
    checked
  />
  <input
    type="radio"
    name="options"
    aria-label="2"
    class="join-item btn btn-square"
  />
  <input
    type="radio"
    name="options"
    aria-label="3"
    class="join-item btn btn-square"
  />
  <input
    type="radio"
    name="options"
    aria-label="4"
    class="join-item btn btn-square"
  />
</div>
```

---

### Basic DaisyUI Chat Bubble (HTML)

Source: https://daisyui.com/components/chat

Demonstrates the basic structure of a chat bubble using DaisyUI classes. It requires 'chat-start' or 'chat-end' for horizontal alignment and utilizes various color classes for the bubble itself.

```html
<div class="chat chat-start">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img
        src="https://img.daisyui.com/images/stock/photo-1534528736660-8977d6180181.webp"
      />
    </div>
  </div>
  <div class="chat-header">
    Obi-Wan Kenobi
    <time class="text-xs opacity-50">12:45</time>
  </div>
  <div class="chat-bubble chat-bubble-neutral">It is true. All of it.</div>
  <div class="chat-footer opacity-50">Delivered</div>
</div>

<div class="chat chat-end">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img
        src="https://img.daisyui.com/images/stock/photo-1534528736660-8977d6180181.webp"
      />
    </div>
  </div>
  <div class="chat-header">
    Anakin Skywalker
    <time class="text-xs opacity-50">12:46</time>
  </div>
  <div class="chat-bubble chat-bubble-primary">Not all of it.</div>
  <div class="chat-footer opacity-50">Seen</div>
</div>
```

---

### DaisyUI Toggle Sizes (HTML)

Source: https://daisyui.com/components/toggle

Illustrates how to apply different size variants (xs, sm, md, lg, xl) to the DaisyUI toggle component.

```html
<input type="checkbox" class="toggle toggle-xs" />
<input type="checkbox" class="toggle toggle-sm" />
<input type="checkbox" class="toggle" />
<input type="checkbox" class="toggle toggle-lg" />
<input type="checkbox" class="toggle toggle-xl" />
```

---

### DaisyUI Basic Tabs with HTML

Source: https://daisyui.com/components/tab

Demonstrates the basic implementation of DaisyUI tabs using HTML. This includes the main 'tabs' container and individual 'tab' elements. The 'tab-content' class is used for the content associated with each tab.

````html
### tabs Tab 1 Tab 2 Tab 3 ``` html
<div role="tablist" class="tabs tabs-lift">
  <a role="tab" class="tab">Tab 1</a>
  <a role="tab" class="tab">Tab 2</a>
  <a role="tab" class="tab">Tab 3</a>
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 1
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 2
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 3
</div>
````

````

--------------------------------

### Highlighted Line Code Mockup with daisyUI

Source: https://daisyui.com/components/mockup-code

Shows how to highlight a specific line within the 'mockup-code' component. This feature is useful for drawing attention to particular commands or output lines, like errors.

```html
<div class="mockup-code">
  <pre data-prefix="$"><code>npm i daisyui
installing...
Error!</code></pre>
</div>
````

---

### Dropdown using details and summary (HTML)

Source: https://daisyui.com/components/dropdown

Implements a dropdown using native HTML 'details' and 'summary' elements. The content is toggled by clicking the summary, and can also be controlled programmatically via the 'open' attribute.

```html
<details>
  <summary>Click me</summary>
  <ul class="p-2">
    <li>Item 1</li>
    <li>Item 2</li>
  </ul>
</details>
```

---

### DaisyUI Join Component: Basic Usage

Source: https://daisyui.com/components/join

Demonstrates the fundamental usage of the DaisyUI 'join' component for grouping buttons. It shows how 'join' and 'join-item' classes are applied to create a horizontally grouped set of buttons.

```html
<div class="join">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
```

---

### Ghost Styled File Input - daisyUI

Source: https://daisyui.com/components/file-input

Illustrates how to apply a 'ghost' style to the file input component for a more minimalist appearance. This style is often used when the input needs to blend seamlessly with the background or other elements.

```html
<input type="file" class="file-input file-input-ghost" />
```

---

### DaisyUI Pagination with Outline Buttons

Source: https://daisyui.com/components/pagination

Shows how to create DaisyUI pagination with outline buttons for 'Previous' and 'Next' pages. This style provides a less prominent visual distinction for navigation.

```html
<div class="join">
  <button class="join-item btn btn-outline">Previous page</button>
  <button class="join-item btn btn-outline">Next</button>
</div>
```

---

### Helper Dropdown with Text in daisyUI

Source: https://daisyui.com/components/dropdown

Illustrates creating a 'helper' dropdown that appears alongside regular text. This pattern is useful for providing additional context or details without cluttering the main content. It uses daisyUI's 'dropdown' component.

```html
<div class="dropdown">
  <label tabindex="0" class="m-1">A normal text and a helper dropdown</label>
  <div class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
    <p>You needed more info? --------------------- Here is a description!</p>
  </div>
</div>
```

---

### HTML Skeleton Loader

Source: https://context7_llms

A placeholder component to indicate a loading state. Uses the 'skeleton' class and requires height and width utility classes (e.g., 'h-_', 'w-_').

```html
<div class="skeleton h-20 w-full"></div>
<div class="skeleton h-32 w-32 rounded-full"></div>
```

---

### Kbd Component for Keyboard Shortcuts with DaisyUI

Source: https://context7_llms

Shows how to use the 'kbd' component in DaisyUI to display keyboard shortcuts. The 'kbd' class is applied to the element, and optional size modifiers like 'kbd-lg' can be used.

```html
<kbd class="kbd kbd-lg">Ctrl + S</kbd>
```

---

### Chat Component Syntax (HTML)

Source: https://context7_llms

Defines the basic structure for a chat message bubble, including options for placement, avatar, and color. Requires placement classes like 'chat-start' or 'chat-end'.

```html
<div class="chat {PLACEMENT}">
  <div class="chat-image"></div>
  <div class="chat-header"></div>
  <div class="chat-bubble {COLOR}">Message text</div>
  <div class="chat-footer"></div>
</div>
```

---

### Drawer with Navbar and Mobile Sidebar (HTML)

Source: https://daisyui.com/components/drawer

Demonstrates a responsive layout where a navbar is visible on desktop, and a drawer sidebar appears on mobile. The screen size change triggers the visibility of the respective navigation elements.

```html
<div class="drawer lg:drawer-open">
  <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content flex flex-col items-center justify-center">
    {/* Page content here */}
    <label for="my-drawer-2" class="btn drawer-button lg:hidden"
      >Open drawer</label
    >
  </div>
  <div class="drawer-side">
    <label
      for="my-drawer-2"
      aria-label="close sidebar"
      class="drawer-overlay"
    ></label>
    <ul class="menu p-4 w-80 bg-base-200 text-base-content">
      {/* Sidebar content here */}
      <li><a>Sidebar Item 1</a></li>
      <li><a>Sidebar Item 2</a></li>
    </ul>
  </div>
</div>
```

---

### Drawer Component Structure (HTML)

Source: https://context7_llms

Outlines the basic HTML structure for a drawer component, which creates a sidebar layout. It includes toggles, content, sidebars, and overlays.

```html
<div class="drawer {MODIFIER}">
  <div class="drawer-content"></div>
  <div class="drawer-side">
    <label for="my-drawer" class="drawer-overlay"></label>
    <ul class="menu p-4 w-80 bg-base-100 text-base-content"></ul>
  </div>
</div>
```

---

### daisyUI CSS Plugin with Light Theme Only

Source: https://context7_llms

This CSS configuration snippet enables the daisyUI plugin but restricts the available themes to only 'light', setting it as the default theme.

```css
@plugin "daisyui" {
  themes: light --default;
}
```

---

### Basic DaisyUI Toggle (HTML)

Source: https://daisyui.com/components/toggle

Demonstrates the fundamental structure of a DaisyUI toggle switch using an HTML checkbox and associated label.

```html
<label class="toggle">
  <input type="checkbox" />
</label>
```

---

### Centered Hero Component - HTML

Source: https://daisyui.com/components/hero

This HTML code snippet demonstrates how to implement a centered hero component using DaisyUI classes. It includes a title, description, and a call-to-action button. No external JavaScript dependencies are required for basic rendering.

```html
<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center">
    <div class="max-w-md">
      <h1 class="text-5xl font-bold">Hello there</h1>
      <p class="py-6">
        Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
        excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a
        id nisi.
      </p>
      <button class="btn btn-primary">Get Started</button>
    </div>
  </div>
</div>
```

---

### DaisyUI Toast with Alert - Basic Usage

Source: https://daisyui.com/components/toast

Demonstrates the basic usage of the DaisyUI Toast component, wrapping an alert message. This snippet shows how to implement a simple notification within the toast container.

```html
<div class="toast">
  <div class="alert alert-info">
    <div>
      <span>New message arrived.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Steps: With Data Content

Source: https://daisyui.com/components/steps

Utilizes the `data-content` attribute to display custom text or icons within each step, providing more detailed information at a glance.

```html
<ul class="steps">
  <li data-content="?" class="step step-neutral">Step 1</li>
  <li data-content="!" class="step step-neutral">Step 2</li>
  <li data-content="✓" class="step step-neutral">Step 3</li>
  <li data-content="✕" class="step step-neutral">Step 4</li>
  <li data-content="→" class="step step-neutral">Step 5</li>
  <li data-content="←" class="step step-neutral">Step 6</li>
  <li data-content="★" class="step step-neutral">Step 7</li>
</ul>
```

---

### DaisyUI Alert Component Syntax (HTML)

Source: https://context7_llms

Shows the basic HTML structure for a DaisyUI alert component. Alerts are used to display important messages to users. Various styles, colors, and directions can be applied using modifier classes.

```html
<div role="alert" class="alert {MODIFIER}">{CONTENT}</div>
```

---

### Card Component Structure (HTML)

Source: https://context7_llms

Demonstrates the basic HTML structure for creating cards using DaisyUI, including image placement, title, body, and actions. Cards are used to group and display content.

```html
<div class="card w-96 bg-base-100 shadow-xl">
  <figure>
    <img src="/img/stock/photo-1606107557116-3565ded65004.jpg" alt="Shoes" />
  </figure>
  <div class="card-body">
    <h2 class="card-title">Card title!</h2>
    <p>
      If a dog chews on his bag of bones, would a puppy have an appetite for the
      same?
    </p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Buy Now</button>
    </div>
  </div>
</div>
```

---

### Basic Status Component Usage (HTML)

Source: https://daisyui.com/components/status

Demonstrates the fundamental usage of the Status component using its default class name. This serves as the base for all other status variations.

````html
```html <span class="status"></span>
````

````

--------------------------------

### DaisyUI Dropdown Syntax (Popover API)

Source: https://context7_llms

This method utilizes the Popover API to create a DaisyUI Dropdown. A button is anchored to the dropdown content using `anchor-name` and `--{anchor}` CSS variables. The dropdown content itself is a `<ul>` with the `dropdown-content` class, marked as a `popover` and positioned using `position-anchor`.

```html
<button popovertarget="{id}" style="anchor-name:--{anchor}">{button}</button>
<ul class="dropdown-content" popover id="{id}" style="position-anchor:--{anchor}">{CONTENT}</ul>
````

---

### Divider Component Syntax (HTML)

Source: https://context7_llms

Shows the basic HTML for a divider, used to separate content. It can include text and various modifiers for color, direction, and placement.

```html
<div class="divider {MODIFIER}">{text}</div>
```

---

### DaisyUI Basic Footer Component

Source: https://daisyui.com/components/footer

Demonstrates the fundamental structure of a DaisyUI footer, organized into columns for services, company information, and legal links. It utilizes vertical stacking by default, suitable for smaller screens.

```html
<footer class="footer p-10 bg-base-200 text-base-content">
  <div>
    <span class="footer-title">Services</span>
    <a class="link link-hover">Branding</a>
    <a class="link link-hover">Design</a>
    <a class="link link-hover">Marketing</a>
    <a class="link link-hover">Advertisement</a>
  </div>
  <div>
    <span class="footer-title">Company</span>
    <a class="link link-hover">About us</a>
    <a class="link link-hover">Contact</a>
    <a class="link link-hover">Jobs</a>
    <a class="link link-hover">Press kit</a>
  </div>
  <div>
    <span class="footer-title">Legal</span>
    <a class="link link-hover">Terms of use</a>
    <a class="link link-hover">Privacy policy</a>
    <a class="link link-hover">Cookie policy</a>
  </div>
</footer>
```

---

### Basic FAB - HTML

Source: https://context7_llms

A minimal Floating Action Button (FAB) structure consisting of a container and a single button. This serves as the base for more complex FAB implementations.

```html
<div class="fab">
  <button class="btn btn-lg btn-circle">{IconOriginal}</button>
</div>
```

---

### Primary Color File Input - daisyUI

Source: https://daisyui.com/components/file-input

Demonstrates how to apply the primary color theme to the file input component. This snippet shows variations using primary, secondary, accent, info, success, warning, and error colors.

```html
<div>
  <input type="file" class="file-input file-input-primary" />
  <input type="file" class="file-input file-input-secondary" />
  <input type="file" class="file-input file-input-accent" />
  <input type="file" class="file-input file-input-info" />
  <input type="file" class="file-input file-input-success" />
  <input type="file" class="file-input file-input-warning" />
  <input type="file" class="file-input file-input-error" />
</div>
```

---

### Theme Controller Input for DaisyUI Themes

Source: https://context7_llms

Allows users to switch themes by selecting a checkbox or radio input with the 'theme-controller' class. The input's value should correspond to a valid DaisyUI theme name.

```html
<input type="checkbox" value="{theme-name}" class="theme-controller" />
```

---

### DaisyUI Text Input: Data List Suggestions

Source: https://daisyui.com/components/input

Demonstrates implementing a text input with data list suggestions using HTML's datalist element and DaisyUI styling. This provides users with predefined options to choose from.

````html
```html
<input class="input input-bordered w-full max-w-xs" list="datalistOptions" />
<datalist id="datalistOptions">
  <option value="San Francisco" />
  <option value="New York" />
  <option value="London" />
  <option value="Paris" />
</datalist>
````

````

--------------------------------

### daisyUI CSS Plugin with Default Themes

Source: https://context7_llms

This CSS configuration sets up the daisyUI plugin to include both 'light' (as default) and 'dark' (as prefers-dark) themes. It also specifies the root element and leaves include/exclude/prefix options empty.

```css
@plugin "daisyui" {
  themes: light --default, dark --prefersdark;
  root: ":root";
  include: ;
  exclude: ;
  prefix: ;
  logs: true;
}
````

---

### DaisyUI Progress Bar Syntax (HTML)

Source: https://context7_llms

Provides the basic HTML structure for a DaisyUI progress bar. This component is used to visually indicate the progress of a task or the passage of time.

```html
<progress class="progress" value="70" max="100"></progress>
```

---

### DaisyUI Drawer Component Syntax

Source: https://context7_llms

The basic HTML structure for a DaisyUI Drawer. It includes a checkbox input to control visibility, the main content area, and a sidebar. A label associated with the checkbox input is used to toggle the drawer's state. Modifiers can be applied for placement and responsiveness.

```html
<div class="drawer {MODIFIER}">
  <input id="my-drawer" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content">{CONTENT}</div>
  <div class="drawer-side">{SIDEBAR}</div>
</div>
```

---

### DaisyUI Steps Component Syntax

Source: https://context7_llms

The DaisyUI steps component visualizes a process flow. It supports vertical and horizontal orientations, colors, and allows for icons and custom data content within each step.

```html
<ul class="steps {MODIFIER}">
  <li class="step">{step content}</li>
</ul>
```

---

### DaisyUI Toggle Colors (HTML)

Source: https://daisyui.com/components/toggle

Demonstrates applying various color classes (primary, secondary, accent, neutral, success, warning, info, error) to DaisyUI toggles.

```html
<input type="checkbox" class="toggle toggle-primary" />
<input type="checkbox" class="toggle toggle-secondary" />
<input type="checkbox" class="toggle toggle-accent" />
<input type="checkbox" class="toggle toggle-neutral" />
<input type="checkbox" class="toggle toggle-success" />
<input type="checkbox" class="toggle toggle-warning" />
<input type="checkbox" class="toggle toggle-info" />
<input type="checkbox" class="toggle toggle-error" />
```

---

### Basic Radial Progress (Vue)

Source: https://daisyui.com/components/radial-progress

A Vue.js implementation of the radial-progress component, demonstrating its basic usage with CSS variables for value and accessibility attributes.

```vue
<template>
  <div class="radial-progress" :style="{ '--value': 70 }" role="progressbar">
    70%
  </div>
</template>

<script setup lang="ts">
// No script needed for basic example
</script>

<style>
/* DaisyUI styles are assumed to be globally imported */
</style>
```

---

### Enable DaisyUI 'toggle' Component

Source: https://daisyui.com/docs/v5

Shows the basic syntax for including the 'toggle' component from DaisyUI. This is a simple directive to enable a specific component's functionality.

```CSS
{ include: toggle; }
```

---

### DaisyUI Alert with Buttons (HTML)

Source: https://daisyui.com/components/alert

This snippet demonstrates a DaisyUI alert component with 'Deny' and 'Accept' buttons. It's useful for user confirmations or notifications where user interaction is required. No external JavaScript dependencies are explicitly shown.

```html
<div>
  <div class="alert">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      class="stroke-info shrink-0 w-6 h-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M13 16h.01M11 20h.01M12 18h.01M12 6h.01M12 10h.01M12 14h.01M7 22h10a2 2 0 002-2V6a2 2 0 00-2-2H7a2 2 0 00-2 2v12a2 2 0 002 2z"
      ></path>
    </svg>
    <span>we use cookies for no reason.</span>
    <div>
      <button class="btn btn-sm btn-ghost">Deny</button>
      <button class="btn btn-sm btn-primary">Accept</button>
    </div>
  </div>
</div>
```

---

### Checkbox Sizes (HTML)

Source: https://daisyui.com/components/checkbox

Illustrates how to apply different size variations to checkboxes using daisyUI's size utility classes. These classes control the visual dimensions of the checkbox component.

````html
```html
<input type="checkbox" class="checkbox checkbox-xs" />
<input type="checkbox" class="checkbox checkbox-sm" />
<input type="checkbox" class="checkbox checkbox-md" />
<input type="checkbox" class="checkbox checkbox-lg" />
<input type="checkbox" class="checkbox checkbox-xl" />
````

````

--------------------------------

### DaisyUI Pagination with Active Button

Source: https://daisyui.com/components/pagination

Demonstrates how to create a DaisyUI pagination component with an active button. This involves using the 'join' and 'join-item' classes to group navigation links or buttons.

```html
<div class="join">
  <button class="join-item btn">1</button>
  <button class="join-item btn btn-active">2</button>
  <button class="join-item btn">3</button>
  <button class="join-item btn">4</button>
</div>
````

---

### Multi-Column Footer Layout

Source: https://daisyui.com/components/footer

This snippet presents a footer with a three-column layout. The columns are labeled 'Services', 'Company', and 'Legal', each containing a list of related links. A final section includes company information and a copyright notice. This structure is ideal for organizing extensive footer content.

```html
### Two footer ###### Services Branding Design Marketing Advertisement ######
Company About us Contact Jobs Press kit ###### Legal Terms of use Privacy policy
Cookie policy ACME Industries Ltd. Providing reliable tech since 1992
```

```html
###### Services Branding Design Marketing Advertisement ###### Company About us
Contact Jobs Press kit ###### Legal Terms of use Privacy policy Cookie policy
ACME Industries Ltd. Providing reliable tech since 1992
```

```html
###### Services Branding Design Marketing Advertisement ###### Company About us
Contact Jobs Press kit ###### Legal Terms of use Privacy policy Cookie policy
ACME Industries Ltd. Providing reliable tech since 1992
```

```html
###### Services Branding Design Marketing Advertisement ###### Company About us
Contact Jobs Press kit ###### Legal Terms of use Privacy policy Cookie policy
ACME Industries Ltd. Providing reliable tech since 1992
```

---

### Link Styling with Color Variants in daisyUI

Source: https://daisyui.com/components/link

Shows how to use various color utility classes with the 'link' class to style links according to the daisyUI theme. These classes include link-neutral, link-primary, link-secondary, link-accent, link-success, link-info, link-warning, and link-error.

```html
<a href="#" class="link link-neutral">Click me</a>
<a href="#" class="link link-primary">Click me</a>
<a href="#" class="link link-secondary">Click me</a>
<a href="#" class="link link-accent">Click me</a>
<a href="#" class="link link-success">Click me</a>
<a href="#" class="link link-info">Click me</a>
<a href="#" class="link link-warning">Click me</a>
<a href="#" class="link link-error">Click me</a>
```

---

### Basic Link Styling with daisyUI

Source: https://daisyui.com/components/link

Demonstrates how to apply the 'link' class to an anchor tag to restore default link styling, overriding Tailwind CSS's default reset. This is the fundamental step for using daisyUI link styles.

```html
<a href="#" class="link">Click me</a>
```

---

### File Input Component in HTML

Source: https://context7_llms

The file-input component allows users to upload files. It supports various modifiers for styling, including ghost, color, and size variations. The basic syntax is an HTML input tag of type file with the 'file-input' class.

```html
<input type="file" class="file-input file-input-bordered w-full max-w-xs" />
```

---

### Breadcrumbs Navigation (HTML)

Source: https://context7_llms

Shows the basic structure for creating breadcrumbs navigation using DaisyUI. Breadcrumbs help users navigate through the site hierarchy.

```html
<div class="breadcrumbs">
  <ul>
    <li><a>Link 1</a></li>
    <li><a>Link 2</a></li>
    <li><a>Link 3</a></li>
  </ul>
</div>
```

---

### DaisyUI Avatar Component Syntax (HTML)

Source: https://context7_llms

Illustrates the HTML markup for DaisyUI avatar components, used for displaying user thumbnails or icons. This includes single avatars and avatar groups, with modifiers for online/offline status and placeholders.

```html
<div class="avatar {MODIFIER}">
  <div>
    <img src="{image-url}" />
  </div>
</div>
```

---

### DaisyUI Rating with Half Stars

Source: https://daisyui.com/components/rating

Demonstrates how to implement half-star ratings using the 'rating-half' modifier. This allows for more granular user input, such as 3.5-star ratings.

```html
<div class="rating rating-half">
  <input
    type="radio"
    name="rating-half-1"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-1"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-2"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-2"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-3"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-3"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-4"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-4"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-5"
    class="bg-yellow-500 mask mask-star-2"
  />
  <input
    type="radio"
    name="rating-half-5"
    class="bg-yellow-500 mask mask-star-2"
  />
</div>
```

---

### Checkbox Component Syntax (HTML)

Source: https://context7_llms

Illustrates the HTML structure for a checkbox with DaisyUI styling. Supports various color and size modifiers.

```html
<input type="checkbox" class="checkbox {MODIFIER}" />
```

---

### Timeline Component Syntax for Event Display

Source: https://context7_llms

A component for displaying events chronologically. It supports vertical and horizontal layouts, along with modifiers for snapping icons and compact views. The 'timeline-start', 'timeline-middle', and 'timeline-end' classes define the layout of each event item.

```html
<ul class="timeline {MODIFIER}">
  <li>
    <div class="timeline-start">{start}</div>
    <div class="timeline-middle">{icon}</div>
    <div class="timeline-end">{end}</div>
  </li>
</ul>
```

---

### Collapse Component Syntax (HTML)

Source: https://context7_llms

Shows the HTML for a collapsible content section. It includes a title and content area, with support for different modifiers like 'collapse-arrow' and 'collapse-plus'.

```html
<div tabindex="0" class="collapse {MODIFIER}">
  <div class="collapse-title">{title}</div>
  <div class="collapse-content">{CONTENT}</div>
</div>
```

---

### Basic Divider Usage in HTML

Source: https://daisyui.com/components/divider

Demonstrates the fundamental use of the daisyUI divider component to create a simple horizontal line separator in HTML. It shows how to insert content before and after the divider.

```html
<div>content</div>

<div class="divider"></div>

<div>content</div>
```

```html
<div>content</div>

<div class="divider">OR</div>

<div>content</div>
```

---

### DaisyUI Accordion Component Syntax (HTML)

Source: https://context7_llms

Demonstrates the structure for a DaisyUI accordion, which uses radio inputs to manage collapsible content. Each accordion item requires a unique name for its radio group to ensure correct functionality. Modifiers like `collapse-arrow` or `collapse-plus` can be applied.

```html
<div class="collapse {MODIFIER}">
  <input type="radio" name="{name}" checked="{checked}" />
  <div class="collapse-title">{title}</div>
  <div class="collapse-content">{CONTENT}</div>
</div>
```

---

### Input Component Styling with DaisyUI

Source: https://context7_llms

Demonstrates how to style input fields using DaisyUI's 'input' component class. Modifiers for style, color, and size can be applied. The 'input' class should be used on the parent element when multiple elements are contained within the input structure.

```html
<input
  type="text"
  placeholder="Type here"
  class="input input-bordered input-primary input-lg"
/>
```

---

### Link Component Styling with DaisyUI

Source: https://context7_llms

Explains how to style anchor tags as links using DaisyUI's 'link' component. Includes options for 'link-hover' and various color modifiers.

```html
<a class="link link-primary link-hover" href="#">Click me</a>
```

---

### Basic Stat Display (HTML)

Source: https://daisyui.com/components/stat

This code snippet demonstrates the basic structure for displaying statistics using the daisyUI Stat component. It includes a container for multiple stats and individual stat blocks with titles, values, and descriptions.

```html
<div class="stats">
  <div class="stat">
    <div class="stat-title">Total Page Views</div>
    <div class="stat-value">89,400</div>
    <div class="stat-desc">21% more than last month</div>
  </div>
</div>
```

---

### DaisyUI Fieldset for Login Form

Source: https://daisyui.com/components/fieldset

Presents a practical application of the DaisyUI fieldset component within a login form. It includes fields for email and password, adhering to accessibility best practices with proper `id` and `for` attributes.

```html
<!-- Login form with fieldset -->
<div class="space-y-4">
  <label class="label">
    <span class="label-text">Login</span>
  </label>
  <div class="form-control">
    <label class="label">
      <span class="label-text">Email</span>
    </label>
    <input
      type="email"
      placeholder="email@example.com"
      class="input input-bordered"
    />
  </div>
  <div class="form-control">
    <label class="label">
      <span class="label-text">Password</span>
    </label>
    <input
      type="password"
      placeholder="password"
      class="input input-bordered"
    />
  </div>
  <button class="btn btn-primary">Login</button>
</div>
```

---

### FAB and Speed Dial (Vertical) - Basic Structure

Source: https://daisyui.com/components/fab

Demonstrates the basic HTML structure for a vertical FAB and Speed Dial. It emphasizes the use of a focusable div with `tabIndex` and `role="button"` for accessibility, and shows how to structure the main FAB and the revealed speed dial buttons.

```html
<div class="fab">
  <button class="fab-main-action">F</button>
  <div class="fab-content">
    <button class="btn btn-xs">A</button>
    <button class="btn btn-xs">B</button>
    <button class="btn btn-xs">C</button>
  </div>
</div>
```

---

### Card as Dropdown in daisyUI

Source: https://daisyui.com/components/dropdown

Demonstrates how to use a card element as a dropdown trigger in daisyUI. This allows for richer content within dropdown menus by leveraging the card component's styling and structure. No external dependencies are required beyond daisyUI.

```html
<div class="dropdown">
  <label tabindex="0" class="btn m-1">Click</label>
  <ul
    tabindex="0"
    class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52"
  >
    <li><a>This is a card. You can use any element as a dropdown.</a></li>
  </ul>
</div>
```

---

### DaisyUI Range Slider with Steps and Measures

Source: https://daisyui.com/components/range

Shows how to implement a range slider with defined steps and visual measures. This is useful for scenarios where specific increments are required, enhancing user guidance.

```html
<input type="range" min="0" max="100" value="50" class="range" />
<div class="w-full flex justify-between text-xs px-2">
  <span>|</span><span>|</span><span>|</span><span>|</span><span>|</span>
</div>

<input type="range" min="0" max="100" value="30" class="range range-primary" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input
  type="range"
  min="0"
  max="100"
  value="70"
  class="range range-secondary"
/>
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input type="range" min="0" max="100" value="90" class="range range-accent" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input type="range" min="0" max="100" value="10" class="range range-success" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input type="range" min="0" max="100" value="40" class="range range-warning" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input type="range" min="0" max="100" value="60" class="range range-info" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>

<input type="range" min="0" max="100" value="80" class="range range-error" />
<div class="w-full flex justify-between text-xs px-2">
  <span>1</span><span>2</span><span>3</span><span>4</span><span>5</span>
</div>
```

---

### Basic DaisyUI Fieldset Usage

Source: https://daisyui.com/components/fieldset

Demonstrates the fundamental structure of a DaisyUI fieldset, including the `fieldset-legend` for the title and `label` for descriptions. This serves as a basic container for related form elements.

```html
<!-- Basic Fieldset -->
<div class="form-control">
  <label class="label">
    <span class="label-text">Page title</span>
  </label>
  <input type="text" placeholder="Type here" class="input input-bordered" />
</div>
```

---

### DaisyUI Navbar: Responsive with Dropdown and Center Menu

Source: https://daisyui.com/components/navbar

This snippet demonstrates a responsive DaisyUI Navbar that features a dropdown menu on small screens and a centered menu on larger screens, adapting to different viewports.

```html
<div
  class="navbar bg-base-100 rounded-box flex-col lg:flex-row justify-between"
>
  <div class="navbar-start w-full">
    <div class="dropdown lg:hidden w-full">
      <label tabindex="0" class="btn btn-ghost btn-circle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h8m-8 6h8"
          />
        </svg>
      </label>
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-full"
      >
        <li><a>Item 1</a></li>
        <li>
          <details>
            <summary>Parent</summary>
            <ul class="p-2 bg-base-100 rounded-t-none">
              <li><a>Submenu 1</a></li>
              <li><a>Submenu 2</a></li>
            </ul>
          </details>
        </li>
        <li><a>Item 3</a></li>
      </ul>
    </div>
    <a class="btn btn-ghost normal-case text-xl lg:hidden">daisyUI</a>
  </div>
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1">
      <li><a>Item 1</a></li>
      <li>
        <details>
          <summary>Parent</summary>
          <ul class="p-2 bg-base-100 rounded-t-none">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </details>
      </li>
      <li><a>Item 3</a></li>
    </ul>
  </div>
  <div class="navbar-end w-full lg:w-auto">
    <button class="btn btn-primary">Button</button>
  </div>
</div>
```

---

### Radial Progress with Different Values (HTML)

Source: https://daisyui.com/components/radial-progress

Shows how to display radial progress at various percentages (0%, 20%, 60%, 80%, 100%) using the '--value' CSS variable. This illustrates the component's flexibility in representing different progress states.

```html
<div class="radial-progress" style="--value:0" role="progressbar">0%</div>
<div class="radial-progress" style="--value:20" role="progressbar">20%</div>
<div class="radial-progress" style="--value:60" role="progressbar">60%</div>
<div class="radial-progress" style="--value:80" role="progressbar">80%</div>
<div class="radial-progress" style="--value:100" role="progressbar">100%</div>
```

---

### Integrate DaisyUI into CSS

Source: https://daisyui.com/

Adds DaisyUI to your project's main CSS file by importing Tailwind CSS followed by the DaisyUI plugin. This step is crucial for enabling DaisyUI's styles and components.

```css
@import "tailwindcss";
@plugin "daisyui";
```

---

### FAB with Vertical Buttons and Labels - HTML

Source: https://context7_llms

This FAB configuration includes labels for each of the expanding buttons, placed before the button element. This provides more context for the action associated with each button.

```html
<div class="fab">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <div>{Label1}<button class="btn btn-lg btn-circle">{Icon1}</button></div>
  <div>{Label2}<button class="btn btn-lg btn-circle">{Icon2}</button></div>
  <div>{Label3}<button class="btn btn-lg btn-circle">{Icon3}</button></div>
</div>
```

---

### DaisyUI Steps: Custom Content in Step Icon

Source: https://daisyui.com/components/steps

Demonstrates how to use custom content, such as emojis, within the step icons for enhanced visual feedback or specific step representation.

```html
<ul class="steps">
  <li class="step step-neutral">😕Step 1</li>
  <li class="step step-neutral">😃Step 2</li>
  <li class="step step-neutral">😍Step 3</li>
</ul>
```

---

### DaisyUI Select Component Ghost Style

Source: https://daisyui.com/components/select

Illustrates how to apply a 'ghost' style to the DaisyUI select component, removing the background for a more minimalist appearance. This is useful for forms where visual separation is less critical.

```html
<label class="form-control w-full max-w-xs">
  <div class="label">
    <span class="label-text">Pick a font</span>
  </div>
  <select class="select select-ghost w-full max-w-xs">
    <option disabled selected>Inter</option>
    <option>Poppins</option>
    <option>Raleway</option>
  </select>
</label>
```

---

### DaisyUI Swap Component: Volume Icon Toggle (HTML)

Source: https://daisyui.com/components/swap

Illustrates how to use the DaisyUI Swap component to toggle between two icons, specifically for volume (on and off states). A hidden checkbox controls the swap logic, with SVG elements representing the icons.

```html
<label class="swap">
  <input type="checkbox" />
  <svg
    class="swap-on fill-current"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path d="M1 12v10l3-3 3 3v-10L6 9l-3-3z"></path>
    <path d="M9 9h3v12h-3zm7-3v12h-3zm4-6v12h-3z"></path>
  </svg>
  <svg
    class="swap-off fill-current"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path d="M1 12v10l3-3 3 3v-10L6 9l-3-3z"></path>
    <path d="M9 9h3v12h-3zm7-3v12h-3zm4-6v12h-3z"></path>
  </svg>
</label>
```

---

### Status Component Sizes (HTML)

Source: https://daisyui.com/components/status

Illustrates how to apply different size variations to the Status component using size-specific classes. Available sizes range from extra small (xs) to extra large (xl).

````html
```html
<span class="status status-xs"></span>
<span class="status status-sm"></span>
<span class="status status-md"></span>
<span class="status status-lg"></span>
<span class="status status-xl"></span>
````

````

--------------------------------

### DaisyUI Range Slider Basic Usage

Source: https://daisyui.com/components/range

Demonstrates the basic implementation of a range slider using DaisyUI classes. This component allows users to select a value by sliding a handle and can be styled with different color classes.

```html
<input type="range" class="range">
<input type="range" class="range range-primary">
<input type="range" class="range range-secondary">
<input type="range" class="range range-accent">
<input type="range" class="range range-success">
<input type="range" class="range range-warning">
<input type="range" class="range range-info">
<input type="range" class="range range-error">
````

---

### Long Line Scrolling in Code Mockup with daisyUI

Source: https://daisyui.com/components/mockup-code

Demonstrates how the 'mockup-code' component handles long lines of text by enabling horizontal scrolling. This ensures readability for code or text that exceeds the container width.

```html
<div class="mockup-code">
  <pre><code>Magnam dolore beatae necessitatibus nemopsum itaque sit. Et porro quae qui et et dolore ratione.</code></pre>
</div>
```

---

### DaisyUI Dropdown Menu Component

Source: https://daisyui.com/components/menu

Illustrates how to create a collapsible dropdown menu using DaisyUI's menu component. Requires JS for toggle functionality.

```html
<ul class="menu">
  <li>
    <details>
      <summary>Menu Item</summary>
      <ul>
        <li>Sub Item 1</li>
        <li>Sub Item 2</li>
      </ul>
    </details>
  </li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>
```

---

### Exclude Specific DaisyUI Components with Plugins

Source: https://daisyui.com/docs/v5

Demonstrates how to exclude specific components from DaisyUI when using its plugin system. This is useful for optimizing the library by only including necessary parts. It shows the syntax for both including and excluding components.

```CSS
@plugin "daisyui" {
  exclude: scrollbar;
}
```

---

### DaisyUI Hero Section with Form (HTML)

Source: https://daisyui.com/components/hero

This snippet demonstrates a hero section with an integrated form for user login. It includes fields for email and password, along with a 'Forgot password?' link and a submit button. This pattern is useful for creating prominent call-to-action areas on a webpage.

```html
### Hero with form ### Login now! Provident cupiditate voluptatem et in. Quaerat
fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut
repudiandae et a id nisi. Email Password Forgot password? Login
```

---

### DaisyUI Card with Image Overlay

Source: https://daisyui.com/components/card

This code demonstrates a DaisyUI card where the image acts as an overlay for the content. It includes a title, descriptive text, and a 'Buy Now' button positioned within the image area. This design is effective for visually engaging content.

```html
<div class="card image-full w-96 shadow-xl">
  <figure>
    <img
      src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
      alt="Shoes"
    />
  </figure>
  <div class="card-body">
    <h2 class="card-title">Card Title</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Buy Now</button>
    </div>
  </div>
</div>
```

---

### DaisyUI Text Input: Colors

Source: https://daisyui.com/components/input

Demonstrates how to apply different color variations (primary, secondary, accent, info, success, warning, error) to DaisyUI text inputs. This allows for visual feedback and categorization.

````html
```html
<input
  type="text"
  placeholder="Primary"
  class="input input-primary input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Secondary"
  class="input input-secondary input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Accent"
  class="input input-accent input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Info"
  class="input input-info input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Success"
  class="input input-success input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Warning"
  class="input input-warning input-bordered w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Error"
  class="input input-error input-bordered w-full max-w-xs"
/>
````

````

--------------------------------

### DaisyUI Navbar: Search Input and Dropdown

Source: https://daisyui.com/components/navbar

This snippet demonstrates a DaisyUI Navbar incorporating a search input field and a user profile dropdown menu. It combines utility with user interaction elements.

```html
<div class="navbar bg-base-100 shadow-lg">
  <div class="flex-1">
    <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
  </div>
  <div class="flex-none">
    <div class="form-control">
      <input type="text" placeholder="Search" class="input input-bordered" />
    </div>
    <div class="dropdown dropdown-end">
      <label tabindex="0" class="btn btn-ghost btn-circle avatar">
        <div class="w-10 rounded-full">
          <img src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
        </div>
      </label>
      <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
        <li>
          <a class="justify-between">
            Profile
            <span class="badge">New</span>
          </a>
        </li>
        <li><a>Settings</a></li>
        <li><a>Logout</a></li>
      </ul>
    </div>
  </div>
</div>
````

---

### DaisyUI Toast Placement: Middle-Start

Source: https://daisyui.com/components/toast

Shows the implementation of the DaisyUI Toast component positioned in the middle-start (vertically centered, horizontally left) of the page. This uses 'toast-middle' and 'toast-start'.

```html
<div class="toast toast-start toast-middle">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### Basic Breadcrumbs - daisyUI

Source: https://daisyui.com/components/breadcrumbs

This snippet shows the fundamental implementation of breadcrumbs using daisyUI classes. It's designed for simple navigation where users can easily track their location within a site's hierarchy. No external dependencies are required beyond daisyUI and Tailwind CSS.

```html
<div role="Breadcrumb" class="breadcrumbs">
  <ul>
    <li>
      <a href="#">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          class="w-4 h-4 stroke-current"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6a2 2 0 00-2-2H5a2 2 0 00-2 2z"
          ></path>
        </svg>
        Home
      </a>
    </li>
    <li>
      <a href="#">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          class="w-4 h-4 stroke-current"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 13h6m-3-3v6m5 5H9a2 2 0 01-2-2V5a2 2 0 012-2h5.5a2 2 0 011.54.535M9 22V7h5.5a2 2 0 011.548.535l3.452 3.452a2 2 0 01.547 1.548L21 17a2 2 0 01-2 2H9z"
          ></path>
        </svg>
        Documents
      </a>
    </li>
    <li>Add Document</li>
  </ul>
</div>
```

---

### Skeleton with Rectangle and Content - daisyUI

Source: https://daisyui.com/components/skeleton

This code snippet illustrates the use of a rectangular Skeleton component, suitable for placeholder content like text blocks or image areas. It provides a visual cue for loading data. This requires the daisyUI framework.

```html
<div class="skeleton rectangle"></div>
```

---

### DaisyUI Stack Component: Basic Usage

Source: https://daisyui.com/components/stack

Demonstrates the fundamental usage of the DaisyUI 'stack' component to vertically align three divs. This component utilizes Tailwind CSS classes to achieve the stacking effect.

```html
<div class="stack">
  <div>1</div>
  <div>2</div>
  <div>3</div>
</div>
```

---

### Radial Progress with Custom Size and Thickness (Vue)

Source: https://daisyui.com/components/radial-progress

Illustrates customizing the size and thickness of the radial-progress component in Vue.js using '--size' and '--thickness' CSS variables for a tailored look.

```vue
<template>
  <div
    class="radial-progress"
    :style="{
      '--value': 70,
      '--size': '10rem',
      '--thickness': '1rem',
    }"
    role="progressbar"
  >
    70%
  </div>
</template>

<script setup lang="ts">
// No script needed for basic example
</script>

<style>
/* DaisyUI styles are assumed to be globally imported */
</style>
```

---

### DaisyUI Badge with Outline Style and Colors (HTML)

Source: https://daisyui.com/components/badge

Shows how to use the 'outline' style for badges combined with different color classes. This results in a bordered appearance without a filled background.

```html
<!-- Outline Primary Badge -->
<div class="badge badge-outline badge-primary">Primary</div>

<!-- Outline Secondary Badge -->
<div class="badge badge-outline badge-secondary">Secondary</div>

<!-- Outline Accent Badge -->
<div class="badge badge-outline badge-accent">Accent</div>

<!-- Outline Info Badge -->
<div class="badge badge-outline badge-info">Info</div>

<!-- Outline Success Badge -->
<div class="badge badge-outline badge-success">Success</div>

<!-- Outline Warning Badge -->
<div class="badge badge-outline badge-warning">Warning</div>

<!-- Outline Error Badge -->
<div class="badge badge-outline badge-error">Error</div>
```

---

### Hero Component Structure in HTML

Source: https://context7_llms

The hero component provides a large container for key content, often used for landing pages. It utilizes the 'hero' class and can include 'hero-content' for text and 'hero-overlay' for background image effects.

```html
<div
  class="hero min-h-screen"
  style="background-image: url(https://img.daisyui.com/images/stock/photo-1508762717628-14311d6a4f5c.jpg);"
>
  <div class="hero-overlay bg-opacity-60"></div>
  <div class="hero-content text-center text-neutral-content">
    <div class="max-w-md">
      <h1 class="mb-5 text-5xl font-bold">Hello there</h1>
      <p class="mb-5">
        Provident cupidatat officia deserunt mollit anim id est laborum. Seitan
        Shirley est cillum dolore eu fugiat nulla pariatur. Excepteur sint
        occaecat cupidatat non dolore magna aliqua.
      </p>
      <button class="btn btn-primary">Get Started</button>
    </div>
  </div>
</div>
```

---

### HTML Stack Layout

Source: https://context7_llms

Arranges elements visually on top of each other using the 'stack' class. Optional modifiers like 'stack-top' and 'stack-bottom' can control placement. Width and height can be set with utility classes.

```html
<div class="stack">
  <div>1</div>
  <div>2</div>
  <div>3</div>
</div>
<div class="stack stack-bottom h-32 w-32">
  <div class="bg-primary">1</div>
  <div class="bg-secondary">2</div>
  <div class="bg-accent">3</div>
</div>
```

---

### Countdown Component Syntax (HTML)

Source: https://context7_llms

Provides the HTML for a countdown timer with a transition effect. Requires a CSS variable '--value' to set the number and JavaScript to update it.

```html
<span class="countdown">
  <span style="--value:{number};">number</span>
</span>
```

---

### DaisyUI Tab Component Syntax (Radio Inputs)

Source: https://context7_llms

This DaisyUI tab component utilizes radio inputs for managing tab states, often paired with 'tabs-box' styling. Radio inputs are essential for the tab content to function correctly upon clicking.

```html
<div role="tablist" class="tabs tabs-box">
  <input type="radio" name="my_tabs" class="tab" aria-label="Tab" />
</div>
```

---

### DaisyUI Swap Component Syntax (Checkbox)

Source: https://context7_llms

The DaisyUI swap component allows toggling between two elements using a checkbox. It provides 'swap-on' and 'swap-off' classes for the content to be displayed in active and inactive states, respectively.

```html
<label class="swap {MODIFIER}">
  <input type="checkbox" />
  <div class="swap-on">{content when active}</div>
  <div class="swap-off">{content when inactive}</div>
</label>
```

---

### DaisyUI Navbar: Icon, Indicator, and Dropdown

Source: https://daisyui.com/components/navbar

Presents a DaisyUI Navbar with an icon, a notification indicator, and a dropdown menu. This is suitable for e-commerce or user-centric applications.

```html
<div class="navbar bg-base-100 shadow-lg">
  <div class="flex-1">
    <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
  </div>
  <div class="flex-none">
    <div class="dropdown dropdown-end">
      <label tabindex="0" class="btn btn-ghost btn-circle">
        <div class="indicator">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 17h5l1-1.769-1-1.769h-6c0-1.066-.864-1.927-1.93-2.001C11.066 10.864 10 11.725 10 13v4h5zM6 17v4a2 2 0 002 2h10a2 2 0 002-2v-4"
            />
          </svg>
          <span class="badge badge-sm indicator-item">8</span>
        </div>
      </label>
      <div
        tabindex="0"
        class="mt-3 card card-compact dropdown-content w-52 bg-base-100 shadow"
      >
        <div class="card-body">
          <span class="font-bold text-lg">8 Items</span>
          <span class="text-info">Subtotal: $999</span>
          <div class="card-actions">
            <button class="btn btn-primary btn-block btn-sm">View cart</button>
          </div>
        </div>
      </div>
    </div>
    <div class="dropdown dropdown-end">
      <label tabindex="0" class="btn btn-ghost btn-circle avatar">
        <div class="w-10 rounded-full">
          <img
            src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
          />
        </div>
      </label>
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li>
          <a class="justify-between">
            Profile
            <span class="badge">New</span>
          </a>
        </li>
        <li><a>Settings</a></li>
        <li><a>Logout</a></li>
      </ul>
    </div>
  </div>
</div>
```

---

### Label Component for Input Fields with DaisyUI

Source: https://context7_llms

Demonstrates two ways to use the 'label' component in DaisyUI: regular labels associated with inputs and 'floating-label' for a dynamic label effect. The 'input' class styles the parent container for regular labels.

```html
<label class="input">
  <span class="label">Email</span>
  <input type="email" placeholder="Type here" />
</label>
```

```html
<label class="floating-label">
  <input type="text" placeholder="Type here" class="input" />
  <span>Username</span>
</label>
```

---

### DaisyUI Badge with Soft Style and Colors (HTML)

Source: https://daisyui.com/components/badge

Demonstrates the 'soft' style applied to badges in conjunction with various color classes. This creates a less prominent, filled appearance suitable for different contexts.

```html
<!-- Soft Primary Badge -->
<div class="badge badge-soft badge-primary">Primary</div>

<!-- Soft Secondary Badge -->
<div class="badge badge-soft badge-secondary">Secondary</div>

<!-- Soft Accent Badge -->
<div class="badge badge-soft badge-accent">Accent</div>

<!-- Soft Info Badge -->
<div class="badge badge-soft badge-info">Info</div>

<!-- Soft Success Badge -->
<div class="badge badge-soft badge-success">Success</div>

<!-- Soft Warning Badge -->
<div class="badge badge-soft badge-warning">Warning</div>

<!-- Soft Error Badge -->
<div class="badge badge-soft badge-error">Error</div>
```

---

### Centered Footer with Logo and Social Icons

Source: https://daisyui.com/components/footer

This snippet demonstrates a centered footer layout that includes a company name, a tagline, and a copyright notice. It's designed to be a common footer pattern for websites and applications. No specific external dependencies are required beyond a basic HTML structure and potentially CSS for styling.

```html
### Centered footer with logo and social icons ACME Industries Ltd. Providing
reliable tech since 1992 Copyright © 2025 - All right reserved
```

```html
ACME Industries Ltd. Providing reliable tech since 1992 Copyright © {new
Date().getFullYear()} - All right reserved
```

```html
ACME Industries Ltd. Providing reliable tech since 1992 Copyright © {new
Date().getFullYear()} - All right reserved
```

```html
ACME Industries Ltd. Providing reliable tech since 1992 Copyright © {new
Date().getFullYear()} - All right reserved
```

```html
ACME Industries Ltd. Providing reliable tech since 1992 Copyright © {new
Date().getFullYear()} - All right reserved
```

---

### FAB with Main Action Button - HTML

Source: https://context7_llms

A FAB that features a distinct 'main action' button displayed when the FAB is expanded. This button is visually different and intended for a primary action.

```html
<div class="fab">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <div class="fab-main-action">
    {LabelMainAction}<button class="btn btn-circle btn-secondary btn-lg">
      {IconMainAction}
    </button>
  </div>
  <div>{Label1}<button class="btn btn-lg btn-circle">{Icon1}</button></div>
  <div>{Label2}<button class="btn btn-lg btn-circle">{Icon2}</button></div>
  <div>{Label3}<button class="btn btn-lg btn-circle">{Icon3}</button></div>
</div>
```

---

### DaisyUI Swap Component: Basic Text Toggle (HTML)

Source: https://daisyui.com/components/swap

Demonstrates the fundamental usage of the DaisyUI Swap component for toggling between two text elements (ON/OFF). It utilizes a hidden checkbox to control the visibility of the 'swap-on' and 'swap-off' child elements.

```html
<label class="swap">
  <input type="checkbox" />
  <div class="swap-on">ON</div>
  <div class="swap-off">OFF</div>
</label>
```

---

### Join Component for Grouping Elements with DaisyUI

Source: https://context7_llms

Illustrates the use of the 'join' and 'join-item' classes in DaisyUI to group elements like buttons or inputs. Supports 'join-vertical' and 'join-horizontal' for orientation and responsive adjustments with 'lg:join-horizontal'.

```html
<div class="join">
  <button class="btn join-item">Button 1</button>
  <button class="btn join-item">Button 2</button>
  <button class="btn join-item">Button 3</button>
</div>
```

---

### Responsive Drawer: Always Visible on Large Screens (HTML)

Source: https://daisyui.com/components/drawer

Illustrates a responsive drawer pattern where the sidebar is persistently visible on large screens (using `lg:drawer-open`) and can be toggled on smaller screens via a button.

```html
<div class="drawer lg:drawer-open">
  <input id="my-drawer-3" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content flex flex-col items-center justify-center">
    {/* Page content here */}
    <label for="my-drawer-3" class="btn drawer-button lg:hidden"
      >Open drawer</label
    >
  </div>
  <div class="drawer-side">
    <label
      for="my-drawer-3"
      aria-label="close sidebar"
      class="drawer-overlay"
    ></label>
    <ul class="menu p-4 w-80 bg-base-200 text-base-content">
      {/* Sidebar content here */}
      <li><a>Sidebar Item 1</a></li>
      <li><a>Sidebar Item 2</a></li>
    </ul>
  </div>
</div>
```

---

### Status Component Colors (HTML)

Source: https://daisyui.com/components/status

Shows how to style the Status component with various color options provided by daisyUI, including neutral, primary, secondary, accent, info, success, warning, and error.

````html
```html
<span class="status status-neutral"></span>
<span class="status status-primary"></span>
<span class="status status-secondary"></span>
<span class="status status-accent"></span>
<span class="status status-info"></span>
<span class="status status-success"></span>
<span class="status status-warning"></span>
<span class="status status-error"></span>
````

````

--------------------------------

### Drawer Component Structure (HTML)

Source: https://daisyui.com/components/drawer

Defines the basic HTML structure for a DaisyUI Drawer component, including the main container, a hidden toggle checkbox, content area, sidebar wrapper, overlay, and sidebar content.

```html
.drawer // The root container
  .drawer-toggle // A hidden checkbox to toggle the visibility of the sidebar
  .drawer-content // All your page content goes here
    // navbar, content, footer
    .drawer-side // Sidebar wrapper
      .drawer-overlay // A dark overlay that covers the whole page when the drawer is open
      // Sidebar content (menu or anything)
````

---

### DaisyUI Toggle with Custom Colors (HTML)

Source: https://daisyui.com/components/toggle

Illustrates how to apply custom colors to a DaisyUI toggle switch, allowing for more design flexibility.

```html
<input type="checkbox" data-toggle-bg="#673ab7" class="toggle" />
<input type="checkbox" data-toggle-bg="#e91e63" class="toggle" />
```

---

### Basic Input Field in HTML

Source: https://context7_llms

A standard text input field. This component is a basic building block for forms and can be styled using various DaisyUI classes for appearance and size.

```html
<input
  type="text"
  placeholder="Type here..."
  class="input input-bordered w-full max-w-xs"
/>
```

---

### DaisyUI Text Input: Ghost Style

Source: https://daisyui.com/components/input

Shows how to apply the 'ghost' style to DaisyUI text inputs, creating a minimalist appearance. This style is useful for elements that should blend seamlessly with the background.

````html
```html
<input
  type="text"
  placeholder="Ghost"
  class="input input-bordered input-ghost w-full max-w-xs"
/>
````

````

--------------------------------

### List Component for Structured Data with DaisyUI

Source: https://context7_llms

Shows the DaisyUI 'list' component for displaying information in rows. Uses 'list' for the container and 'list-row' for each item. Modifiers like 'list-col-wrap' and 'list-col-grow' can alter layout.

```html
<ul class="list">
  <li class="list-row">
    <span>Item 1</span>
    <span>Details 1</span>
  </li>
  <li class="list-row">
    <span>Item 2</span>
    <span>Details 2</span>
  </li>
</ul>
````

---

### daisyUI 5: Improved Color Variables

Source: https://daisyui.com/docs/v5

Illustrates the new, more readable, and customizable CSS variables for colors in daisyUI 5, utilizing the 'oklch' format for better control and compatibility with Tailwind CSS 4's color-mix() function. This replaces the older, harder-to-read format.

```css
/* Before
* hard to read variable names, hard to customize values
*/
 {
  --b1: 100% 0 0;
  --b2: 96.1151% 0 0;
  --b3: 92.4169% 0.00108 197.137559;
  --bc: 27.8078% 0.029596 256.847952;
  --p: 49.12% 0.3096 275.75;
  --pc: 89.824% 0.06192 275.75;
  --s: 69.71% 0.329 342.55;
  --sc: 98.71% 0.0106 342.55;
  --a: 76.76% 0.184 183.61;
  --ac: 15.352% 0.0368 183.61;
  --n: 32.1785% 0.02476 255.701624;
  --nc: 89.4994% 0.011585 252.096176;
  --in: 72.06% 0.191 231.6;
  --inc: 0% 0 0;
  --su: 64.8% 0.15 160;
  --suc: 0% 0 0;
  --wa: 84.71% 0.199 83.87;
  --wac: 0% 0 0;
  --er: 71.76% 0.221 22.18;
  --erc: 0% 0 0;
}

/* After
* standard color format, easy to customize in browser
*/
 {
  --color-base-100: oklch(100% 0 0);
  --color-base-200: oklch(96.115% 0 0);
  --color-base-300: oklch(92.416% 0.001 197.137);
  --color-base-content: oklch(27.807% 0.029 256.847);
  --color-primary: oklch(49.12% 0.309 275.75);
  --color-primary-content: oklch(89.824% 0.061 275.75);
  --color-secondary: oklch(69.71% 0.329 342.55);
  --color-secondary-content: oklch(98.71% 0.01 342.55);
  --color-accent: oklch(76.76% 0.184 183.61);
  --color-accent-content: oklch(15.352% 0.036 183.61);
  --color-neutral: oklch(20% 0.024 255.701);
  --color-neutral-content: oklch(89.499% 0.011 252.096);
  --color-info: oklch(72.06% 0.191 231.6);
  --color-info-content: oklch(0% 0 0);
  --color-success: oklch(64.8% 0.15 160);
  --color-success-content: oklch(0% 0 0);
  --color-warning: oklch(84.71% 0.199 83.87);
  --color-warning-content: oklch(0% 0 0);
  --color-error: oklch(71.76% 0.221 22.18);
  --color-error-content: oklch(0% 0 0);
}
```

---

### DaisyUI Tab Component Syntax (Buttons)

Source: https://context7_llms

The DaisyUI tab component displays links in a tabbed format using button elements. It supports various styles like 'tabs-box' and 'tabs-border', and placement options like 'tabs-top'.

```html
<div role="tablist" class="tabs {MODIFIER}">
  <button role="tab" class="tab">Tab</button>
</div>
```

---

### Basic Radial Progress (React TSX)

Source: https://daisyui.com/components/radial-progress

Provides a React (TSX) implementation of the radial-progress component. It uses CSS variables for value, size, and thickness, and includes accessibility attributes.

```tsx
function RadialProgressTsx() {
  return (
    <div
      className="radial-progress"
      style={{ "--value": 70 } as React.CSSProperties}
      role="progressbar"
    >
      70%
    </div>
  );
}
```

---

### DaisyUI Bordered Tabs with HTML

Source: https://daisyui.com/components/tab

Illustrates how to apply a border style to the DaisyUI tabs component. This variation uses the 'tabs-border' class for a distinct visual appearance. The structure remains similar to basic tabs.

````html
### tabs-border Tab 1 Tab 2 Tab 3 ``` html
<div role="tablist" class="tabs tabs-border">
  <a role="tab" class="tab">Tab 1</a>
  <a role="tab" class="tab">Tab 2</a>
  <a role="tab" class="tab">Tab 3</a>
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 1
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 2
</div>
<div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
  Tab content 3
</div>
````

````

--------------------------------

### DaisyUI Menu Sizes

Source: https://daisyui.com/components/menu

Illustrates how to apply different size modifiers to DaisyUI menu items, ranging from extra small to extra large. This allows for flexible visual design.

```html
<ul class="menu menu-xs">
  <li>Xsmall 1</li>
  <li>Xsmall 2</li>
</ul>
<ul class="menu menu-sm">
  <li>Small 1</li>
  <li>Small 2</li>
</ul>
<ul class="menu menu-md">
  <li>Medium 1</li>
  <li>Medium 2</li>
</ul>
<ul class="menu menu-lg">
  <li>Large 1</li>
  <li>Large 2</li>
</ul>
<ul class="menu menu-xl">
  <li>Xlarge 1</li>
  <li>Xlarge 2</li>
</ul>
````

---

### DaisyUI Toggle with Fieldset and Label (HTML)

Source: https://daisyui.com/components/toggle

Shows how to integrate a DaisyUI toggle within a fieldset and label for better form organization and accessibility.

```html
<div class="form-control">
  <label class="label cursor-pointer">
    <span class="label-text">Remember me</span>
    <input type="checkbox" class="toggle" checked />
  </label>
</div>
```

---

### DaisyUI Steps: Custom Colors

Source: https://daisyui.com/components/steps

Applies custom color classes to the steps to visually differentiate them or match a specific theme, using DaisyUI's color palette.

```html
<ul class="steps">
  <li class="step step-primary">Fly to moon</li>
  <li class="step step-secondary">Shrink the moon</li>
  <li class="step step-accent">Grab the moon</li>
  <li class="step step-info">Sit on toilet</li>
</ul>
```

---

### DaisyUI Select Component with Fieldset and Labels

Source: https://daisyui.com/components/select

Shows how to integrate the DaisyUI select component within a fieldset, using legend for the main title and labels for individual options. This structure improves accessibility and organization for grouped form elements.

```html
<fieldset class="space-y-6">
  <legend class="sr-only">Browsers</legend>
  <div class="form-control w-full max-w-xs">
    <label class="label">
      <span class="label-text">Pick a browser</span>
    </label>
    <select class="select select-bordered w-full max-w-xs">
      <option disabled selected>Chrome</option>
      <option>FireFox</option>
      <option>Safari</option>
    </select>
    <label class="label">
      <span class="label-text-alt">Optional</span>
    </label>
  </div>
</fieldset>
```

---

### DaisyUI Stack Component: Stacked Images

Source: https://daisyui.com/components/stack

Illustrates how to use the DaisyUI 'stack' component to layer images vertically. This is useful for creating visual effects where images overlap or are presented in a stacked format.

```html
<div class="stack">
  <img
    src="https://img.daisyui.com/images/stock/photo-1572635148818-ef6fd45eb394.webp"
    alt="Image 1"
  />
  <img
    src="https://img.daisyui.com/images/stock/photo-1565098772267-60af42b81ef2.webp"
    alt="Image 2"
  />
  <img
    src="https://img.daisyui.com/images/stock/photo-1559703248-dcaaec9fab78.webp"
    alt="Image 3"
  />
</div>
```

---

### DaisyUI Tab Sizes with HTML

Source: https://daisyui.com/components/tab

Demonstrates the different size variations available for DaisyUI tabs, including extra-small (xs), small (sm), medium (md), large (lg), and extra-large (xl). Each size is applied using corresponding class names.

````html
### Sizes Xsmall Xsmall Xsmall Small Small Small Medium Medium Medium Large
Large Large Xlarge Xlarge Xlarge ``` html
<div class="flex flex-col gap-3">
  <div>
    <div role="tablist" class="tabs tabs-lift">
      <a role="tab" class="tab tab-xs">Xsmall</a>
      <a role="tab" class="tab tab-xs">Xsmall</a>
      <a role="tab" class="tab tab-xs">Xsmall</a>
    </div>
    <div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
      Tab content
    </div>
  </div>
  <div>
    <div role="tablist" class="tabs tabs-lift">
      <a role="tab" class="tab tab-sm">Small</a>
      <a role="tab" class="tab tab-sm">Small</a>
      <a role="tab" class="tab tab-sm">Small</a>
    </div>
    <div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
      Tab content
    </div>
  </div>
  <div>
    <div role="tablist" class="tabs tabs-lift">
      <a role="tab" class="tab tab-md">Medium</a>
      <a role="tab" class="tab tab-md">Medium</a>
      <a role="tab" class="tab tab-md">Medium</a>
    </div>
    <div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
      Tab content
    </div>
  </div>
  <div>
    <div role="tablist" class="tabs tabs-lift">
      <a role="tab" class="tab tab-lg">Large</a>
      <a role="tab" class="tab tab-lg">Large</a>
      <a role="tab" class="tab tab-lg">Large</a>
    </div>
    <div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
      Tab content
    </div>
  </div>
  <div>
    <div role="tablist" class="tabs tabs-lift">
      <a role="tab" class="tab tab-xl">Xlarge</a>
      <a role="tab" class="tab tab-xl">Xlarge</a>
      <a role="tab" class="tab tab-xl">Xlarge</a>
    </div>
    <div class="p-4 border border-t-0 border-base-content/20 rounded-b-box">
      Tab content
    </div>
  </div>
</div>
````

````

--------------------------------

### DaisyUI Join Component: Vertical Grouping

Source: https://daisyui.com/components/join

Illustrates how to use the 'join-vertical' class with the DaisyUI 'join' component to stack grouped items vertically. This is useful for creating vertical lists of buttons or other form elements.

```html
<div class="join join-vertical">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join join-vertical">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join join-vertical">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join join-vertical">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
<div class="join join-vertical">
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
  <button class="join-item btn">Button</button>
</div>
````

---

### DaisyUI Chat Bubble with Image (HTML)

Source: https://daisyui.com/components/chat

Shows how to integrate an image within a DaisyUI chat bubble, typically used for author avatars. The 'chat-image' class is essential for this layout.

```html
<div class="chat chat-start">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img src="https://img.daisyui.com/images/profile/demo/kenobee@192.webp" />
    </div>
  </div>
  <div class="chat-bubble">
    It was said that you would, destroy the Sith, not join them.
  </div>
</div>
<div class="chat chat-start">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img src="https://img.daisyui.com/images/profile/demo/kenobee@192.webp" />
    </div>
  </div>
  <div class="chat-bubble">It was you who would bring balance to the Force</div>
</div>
<div class="chat chat-start">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img src="https://img.daisyui.com/images/profile/demo/kenobee@192.webp" />
    </div>
  </div>
  <div class="chat-bubble">Not leave it in Darkness</div>
</div>

<div class="chat chat-end">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img
        src="https://img.daisyui.com/images/profile/demo/chuckles@192.webp"
      />
    </div>
  </div>
  <div class="chat-bubble">I wish that were so!</div>
</div>
<div class="chat chat-end">
  <div class="chat-image avatar">
    <div class="w-10 rounded-full">
      <img
        src="https://img.daisyui.com/images/profile/demo/chuckles@192.webp"
      />
    </div>
  </div>
  <div class="chat-bubble">I did, and I will!</div>
</div>
```

---

### DaisyUI Tabs with Radio Inputs and HTML

Source: https://daisyui.com/components/tab

Explains how to implement DaisyUI tabs using radio inputs for interactive tab switching. This approach requires unique 'name' attributes for each tab group and uses hidden radio inputs to control the active state.

````html
### tabs-box using radio inputs ``` html
<div class="flex w-full flex-col gap-3">
  <div role="tablist" class="tabs tabs-boxed">
    <input
      type="radio"
      name="my_tab_group_1"
      role="tab"
      class="tab"
      aria-label="Tab 1"
      checked
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 1
    </div>

    <input
      type="radio"
      name="my_tab_group_1"
      role="tab"
      class="tab"
      aria-label="Tab 2"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 2
    </div>

    <input
      type="radio"
      name="my_tab_group_1"
      role="tab"
      class="tab"
      aria-label="Tab 3"
    />
    <div role="tabpanel" class="tab-content border rounded-t-none p-4">
      Tab content 3
    </div>
  </div>
</div>
````

````

--------------------------------

### DaisyUI Swap Component: Icons with Flip Effect (HTML)

Source: https://daisyui.com/components/swap

Demonstrates the flip effect for icon toggling using the DaisyUI Swap component. The `swap-flip` class is applied to the swap container, allowing for a 3D flip transition between two child elements (e.g., devil and angel emojis).

```html
<label class="swap swap-flip">
  <input type="checkbox" />
  <div class="swap-on">😈</div>
  <div class="swap-off">😇</div>
</label>
````

---

### DaisyUI Custom Theme with Tailwind CSS

Source: https://context7_llms

Defines a custom DaisyUI theme by importing Tailwind CSS and daisyUI plugins. It specifies various CSS variables for colors, border-radius, sizes, and component depth/noise. This allows for a fully branded user interface.

```css
@import "tailwindcss";
@plugin "daisyui";
@plugin "daisyui/theme" {
  name: "mytheme";
  default: true; /* set as default */
  prefersdark: false; /* set as default dark mode (prefers-color-scheme:dark) */
  color-scheme: light; /* color of browser-provided UI */

  --color-base-100: oklch(98% 0.02 240);
  --color-base-200: oklch(95% 0.03 240);
  --color-base-300: oklch(92% 0.04 240);
  --color-base-content: oklch(20% 0.05 240);
  --color-primary: oklch(55% 0.3 240);
  --color-primary-content: oklch(98% 0.01 240);
  --color-secondary: oklch(70% 0.25 200);
  --color-secondary-content: oklch(98% 0.01 200);
  --color-accent: oklch(65% 0.25 160);
  --color-accent-content: oklch(98% 0.01 160);
  --color-neutral: oklch(50% 0.05 240);
  --color-neutral-content: oklch(98% 0.01 240);
  --color-info: oklch(70% 0.2 220);
  --color-info-content: oklch(98% 0.01 220);
  --color-success: oklch(65% 0.25 140);
  --color-success-content: oklch(98% 0.01 140);
  --color-warning: oklch(80% 0.25 80);
  --color-warning-content: oklch(20% 0.05 80);
  --color-error: oklch(65% 0.3 30);
  --color-error-content: oklch(98% 0.01 30);

  --radius-selector: 1rem; /* border radius of selectors (checkbox, toggle, badge) */
  --radius-field: 0.25rem; /* border radius of fields (button, input, select, tab) */
  --radius-box: 0.5rem; /* border radius of boxes (card, modal, alert) */
  /* preferred values for --radius-* : 0rem, 0.25rem, 0.5rem, 1rem, 2rem */

  --size-selector: 0.25rem; /* base size of selectors (checkbox, toggle, badge). Value must be 0.25rem unless we intentionally want bigger selectors. In so it can be 0.28125 or 0.3125. If we intentionally want smaller selectors, it can be 0.21875 or 0.1875 */
  --size-field: 0.25rem; /* base size of fields (button, input, select, tab). Value must be 0.25rem unless we intentionally want bigger fields. In so it can be 0.28125 or 0.3125. If we intentionally want smaller fields, it can be 0.21875 or 0.1875 */

  --border: 1px; /* border size. Value must be 1px unless we intentionally want thicker borders. In so it can be 1.5px or 2px. If we intentionally want thinner borders, it can be 0.5px */

  --depth: 1; /* only 0 or 1 – Adds a shadow and subtle 3D depth effect to components */
  --noise: 0; /* only 0 or 1 - Adds a subtle noise (grain) effect to components */
}
```

---

### HTML Dialog Modal with JS Open and Click Outside Close (DaisyUI)

Source: https://daisyui.com/components/modal

This snippet shows a DaisyUI modal implemented with an HTML dialog element. It is opened via JavaScript's `showModal()` method. The modal includes a 'modal-backdrop' class, which, when clicked, closes the modal, providing a common user interaction pattern for dismissing dialogs.

```html
<button class="btn" onclick="my_modal_2.showModal()">open modal</button>
<dialog id="my_modal_2" class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Press ESC key or click outside to close</p>
    <div class="modal-action">
      <form method="dialog">
        <button class="btn">close</button>
      </form>
    </div>
  </div>
  <form method="dialog" class="modal-backdrop">
    <button>close</button>
  </form>
</dialog>
```

```javascript
// Open the modal using document.getElementById('ID').showModal() method
document.getElementById("my_modal_2").showModal();
```

---

### FAB Flower with Tooltips - HTML

Source: https://context7_llms

A FAB in the flower arrangement that uses tooltips to provide labels for the expanding buttons. This is suitable when there isn't enough space for visible text labels.

```html
<div class="fab fab-flower">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <button class="fab-main-action btn btn-circle btn-lg">
    {IconMainAction}
  </button>
  <div class="tooltip tooltip-left" data-tip="{Label1}">
    <button class="btn btn-lg btn-circle">{Icon1}</button>
  </div>
  <div class="tooltip tooltip-left" data-tip="{Label2}">
    <button class="btn btn-lg btn-circle">{Icon2}</button>
  </div>
  <div class="tooltip tooltip-left" data-tip="{Label3}">
    <button class="btn btn-lg btn-circle">{Icon3}</button>
  </div>
</div>
```

---

### Countdown in Boxes (HTML)

Source: https://daisyui.com/components/countdown

Presents the DaisyUI Countdown component styled within boxes, offering a visually distinct way to display time units. Each unit is contained within its own background box.

```html
<div class="grid grid-flow-col gap-5 text-center auto-rows-max">
  <div
    class="flex flex-col p-2 bg-neutral-800 rounded-box text-neutral-content"
  >
    <span class="countdown font-mono text-5xl">
      <span style="--value:15"></span>
    </span>
    days
  </div>
  <div
    class="flex flex-col p-2 bg-neutral-800 rounded-box text-neutral-content"
  >
    <span class="countdown font-mono text-5xl">
      <span style="--value:10"></span>
    </span>
    hours
  </div>
  <div
    class="flex flex-col p-2 bg-neutral-800 rounded-box text-neutral-content"
  >
    <span class="countdown font-mono text-5xl">
      <span style="--value:24"></span>
    </span>
    min
  </div>
  <div
    class="flex flex-col p-2 bg-neutral-800 rounded-box text-neutral-content"
  >
    <span class="countdown font-mono text-5xl">
      <span style="--value:59"></span>
    </span>
    sec
  </div>
</div>
```

---

### Customizing Component Border Size with CSS Variables

Source: https://daisyui.com/docs/v5

Demonstrates how to customize the border size for DaisyUI components like buttons and inputs using the `--border` CSS variable. This allows for consistent border styling across different themes or globally.

```CSS
/* Example for customizing border size */
:root {
  --border: 2px; /* Adjust border width */
}
```

---

### DaisyUI Validator for Form Inputs (HTML)

Source: https://daisyui.com/components/validator

Demonstrates the basic usage of the 'validator' class for form inputs like text, select, and textarea. It shows how an invalid input triggers an error color, and a valid input triggers a success color. The 'validator-hint' class is used for displaying validation messages.

```html
<input
  type="email"
  class="input input-bordered validator"
  placeholder="email"
/>
<p class="validator-hint">Enter valid email address</p>

<select class="select select-bordered validator">
  <option disabled selected>Choose option</option>
  <option>option 1</option>
  <option>option 2</option>
</select>
<p class="validator-hint">Choose option</p>

<textarea
  class="textarea textarea-bordered validator"
  placeholder="Bio"
></textarea>
<p class="validator-hint">Enter valid email address</p>
```

---

### HTML: daisyUI List with Default Second Column Grow

Source: https://daisyui.com/components/list

This snippet demonstrates the default behavior of the daisyUI List component where the second column within a 'list-row' automatically grows to fill remaining space. It is useful for displaying data in a structured, responsive manner.

```html
<div class="list">
  <div class="list-row">
    <div>Most played songs this week</div>
    <div>
      <img
        src="https://img.daisyui.com/images/profile/demo/1@94.webp"
        alt="User Avatar"
      />
      Dio Lupa Remaining Reason
    </div>
  </div>
  <div class="list-row">
    <div></div>
    <div>
      <img
        src="https://img.daisyui.com/images/profile/demo/4@94.webp"
        alt="User Avatar"
      />
      Ellie Beilish Bears of a fever
    </div>
  </div>
  <div class="list-row">
    <div></div>
    <div>
      <img
        src="https://img.daisyui.com/images/profile/demo/3@94.webp"
        alt="User Avatar"
      />
      Sabrino Gardener Cappuccino
    </div>
  </div>
</div>
```

---

### Import daisyUI as a Plugin in Tailwind CSS 4 (CSS)

Source: https://daisyui.com/docs/v5

Demonstrates how to import daisyUI as a plugin directly within your main CSS file when using Tailwind CSS 4. This is the new recommended approach for integration.

```css
/* app.css */
@import "tailwindcss";
@plugin "daisyui";
```

---

### Dropdown using Popover API and anchor positioning (HTML)

Source: https://daisyui.com/components/dropdown

Utilizes the HTML Popover API for dropdowns, which renders content in a top layer, simplifying z-index and overflow issues. CSS Anchor Positioning is used for relative placement, though browser support for anchor positioning varies.

```html
<button popovertarget="my-popover" anchor-name="--my-anchor">Button</button>
<div id="my-popover" popover anchor-name="--my-anchor">
  <ul class="p-2">
    <li>Item 1</li>
    <li>Item 2</li>
  </ul>
</div>
```

---

### DaisyUI Radio Button Colors

Source: https://daisyui.com/components/radio

Shows how to apply different color themes to DaisyUI radio buttons. This includes neutral, primary, secondary, accent, success, warning, info, and error color variants.

```html
<!-- Radio Neutral -->
<input type="radio" name="radio-color" class="radio radio-neutral" checked />

<!-- Radio Primary -->
<input type="radio" name="radio-color" class="radio radio-primary" checked />

<!-- Radio Secondary -->
<input type="radio" name="radio-color" class="radio radio-secondary" checked />

<!-- Radio Accent -->
<input type="radio" name="radio-color" class="radio radio-accent" checked />

<!-- Radio Success -->
<input type="radio" name="radio-color" class="radio radio-success" checked />

<!-- Radio Warning -->
<input type="radio" name="radio-color" class="radio radio-warning" checked />

<!-- Radio Info -->
<input type="radio" name="radio-color" class="radio radio-info" checked />

<!-- Radio Error -->
<input type="radio" name="radio-color" class="radio radio-error" checked />
```

---

### DaisyUI Alert with Title and Description (HTML)

Source: https://daisyui.com/components/alert

This code shows a DaisyUI alert component featuring a title ('New message!') and a description ('You have 1 unread message') followed by a call-to-action link ('See'). It's suitable for displaying important notifications or updates.

```html
<div>
  <div class="alert alert-info">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      class="stroke-current shrink-0 w-6 h-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M13 16h.01M11 20h.01M12 18h.01M12 6h.01M12 10h.01M12 14h.01M7 22h10a2 2 0 002-2V6a2 2 0 00-2-2H7a2 2 0 00-2 2v12a2 2 0 002 2z"
      ></path>
    </svg>
    <div>
      <h3 class="font-bold">New message!</h3>
      <div class="text-xs">You have 1 unread message</div>
    </div>
    <button class="btn btn-sm">See</button>
  </div>
</div>
```

---

### Customizing DaisyUI Themes with Custom Fonts (CSS)

Source: https://daisyui.com/docs/v5

This snippet demonstrates how to customize daisyUI themes by applying custom font families. It shows how to use the `@plugin` directive to modify existing themes like 'cyberpunk' and 'wireframe' with specific font stacks. This is useful for enforcing consistent typography across your project.

```css
@plugin "daisyui";
@plugin "daisyui/theme" {
  name: cyberpunk;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;
}
@plugin "daisyui/theme" {
  name: wireframe;
  font-family: Chalkboard, comic sans ms, "sans-serif";
}
```

---

### DaisyUI Navbar: Title and Icon

Source: https://daisyui.com/components/navbar

Demonstrates how to create a DaisyUI Navbar that includes both a title and an icon. This enhances the visual identity of the navigation bar.

```html
<div class="navbar bg-base-100">
  <a class="btn btn-ghost normal-case text-xl">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 12l2-2m0 0l7-7 7 7M5 12v6a2 2 0 002 2h10a2 2 0 002-2v-6"
      />
    </svg>
    daisyUI
  </a>
</div>
```

---

### DaisyUI Dropdown Syntax (CSS Focus)

Source: https://context7_llms

A DaisyUI Dropdown implemented using CSS focus and `tabindex`. The trigger element has `tabindex="0"` and `role="button"`. The dropdown content, a `<ul>` with the `dropdown-content` class, is made focusable with `tabindex="-1"`. This approach relies on keyboard interaction for opening/closing.

```html
<div class="dropdown">
  <div tabindex="0" role="button">Button</div>
  <ul tabindex="-1" class="dropdown-content">
    {CONTENT}
  </ul>
</div>
```

---

### DaisyUI Toggle with Icons Inside (HTML)

Source: https://daisyui.com/components/toggle

Demonstrates how to embed icons within a DaisyUI toggle switch, typically used for visual indication of the on/off state.

```html
<label class="toggle">
  <input type="checkbox" />
  <span class="label-text">ON</span>
  <span class="label-text">OFF</span>
</label>
```

---

### DaisyUI Pagination with Disabled Button

Source: https://daisyui.com/components/pagination

Illustrates how to create a DaisyUI pagination component with a disabled button, typically used for the 'previous' or 'next' links when they are not available. This uses the 'join' and 'join-item' classes.

```html
<div class="join">
  <button class="join-item btn">1</button>
  <button class="join-item btn">2</button>
  <button class="join-item btn btn-disabled">...</button>
  <button class="join-item btn">99</button>
  <button class="join-item btn btn-active">100</button>
</div>
```

---

### DaisyUI Fieldset with Multiple Inputs

Source: https://daisyui.com/components/fieldset

Illustrates creating a DaisyUI fieldset that contains multiple input fields. It emphasizes the importance of unique `id` and `for` attributes for accessibility, linking labels to their respective inputs.

```html
<!-- Fieldset with multiple inputs -->
<div class="space-y-4">
  <label class="label">
    <span class="label-text">Page details</span>
  </label>
  <div class="form-control">
    <label class="label">
      <span class="label-text">Title</span>
    </label>
    <input type="text" placeholder="Title" class="input input-bordered" />
  </div>
  <div class="form-control">
    <label class="label">
      <span class="label-text">Slug</span>
    </label>
    <input type="text" placeholder="Slug" class="input input-bordered" />
  </div>
  <div class="form-control">
    <label class="label">
      <span class="label-text">Author</span>
    </label>
    <input type="text" placeholder="Author" class="input input-bordered" />
  </div>
</div>
```

---

### DaisyUI Swap Component: Icons with Rotate Effect (HTML)

Source: https://daisyui.com/components/swap

Shows how to implement the rotate effect on icon toggling using the DaisyUI Swap component. The `swap-rotate` class is applied to the swap container, and a hidden checkbox controls the state change between two icons (e.g., sun and moon).

```html
<label class="swap swap-rotate">
  <input type="checkbox" />
  <svg
    class="swap-on fill-current"
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <circle cx="12" cy="12" r="4"></circle>
    <path
      d="M16 12h4M20 12h-4M12 16h.01M12 20h.01M12 8h.01M12 4h.01M8 12h.01M4 12h.01M20 4l-2 2M18 20l2-2M6 4l2 2M4 20l2-2"
    ></path>
  </svg>
  <svg
    class="swap-off fill-current"
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
  </svg>
</label>
```

---

### DaisyUI Badge with Dash Style and Colors (HTML)

Source: https://daisyui.com/components/badge

Illustrates the 'dash' style for badges, applied with various color classes. This creates a dashed border effect, offering a distinct visual alternative to solid outlines.

```html
<!-- Dash Primary Badge -->
<div class="badge badge-dash badge-primary">Primary</div>

<!-- Dash Secondary Badge -->
<div class="badge badge-dash badge-secondary">Secondary</div>

<!-- Dash Accent Badge -->
<div class="badge badge-dash badge-accent">Accent</div>

<!-- Dash Info Badge -->
<div class="badge badge-dash badge-info">Info</div>

<!-- Dash Success Badge -->
<div class="badge badge-dash badge-success">Success</div>

<!-- Dash Warning Badge -->
<div class="badge badge-dash badge-warning">Warning</div>

<!-- Dash Error Badge -->
<div class="badge badge-dash badge-error">Error</div>
```

---

### Hero with Figure Component - HTML

Source: https://daisyui.com/components/hero

This HTML code snippet shows a DaisyUI hero component featuring an image on one side and content on the other. It's suitable for introducing a product or service with visual context. The layout can be reversed using specific classes.

```html
<div class="hero min-h-screen bg-base-200">
  <div class="hero-content flex-col lg:flex-row">
    <img
      src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
      class="max-w-sm rounded-lg shadow-2xl"
    />
    <div>
      <h1 class="text-5xl font-bold">Box Office News!</h1>
      <p class="py-6">
        Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
        excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a
        id nisi.
      </p>
      <button class="btn btn-primary">Get Started</button>
    </div>
  </div>
</div>
```

---

### Large Text Countdown with Labels (HTML)

Source: https://daisyui.com/components/countdown

Demonstrates how to use the DaisyUI Countdown component for displaying larger text with descriptive labels for units like days, hours, minutes, and seconds.

```html
<div class="grid grid-flow-col gap-5 text-neutral-content">
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-5xl">
      <span style="--value:15"></span>
    </span>
    days
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-5xl">
      <span style="--value:10"></span>
    </span>
    hours
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-5xl">
      <span style="--value:24"></span>
    </span>
    min
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-5xl">
      <span style="--value:59"></span>
    </span>
    sec
  </div>
</div>
```

---

### DaisyUI Textarea Component Syntax

Source: https://context7_llms

The DaisyUI textarea component allows multi-line text input from users. It can be styled with 'textarea-ghost' and supports various color and size modifiers for customization.

```html
<textarea class="textarea {MODIFIER}" placeholder="Bio"></textarea>
```

---

### DaisyUI Rating Component Basic Structure

Source: https://daisyui.com/components/rating

This snippet shows the fundamental HTML structure for a DaisyUI rating component. It uses radio inputs within a div with the class 'rating'. Each radio button represents a rating option.

```html
<div class="rating">
  \n <input type="radio" name="rating-1" class="mask mask-star" />\n
  <input type="radio" name="rating-2" class="mask mask-star" checked />\n
  <input type="radio" name="rating-3" class="mask mask-star" />\n
  <input type="radio" name="rating-4" class="mask mask-star" />\n
  <input type="radio" name="rating-5" class="mask mask-star" />\n
</div>
```

---

### DaisyUI Badge within Text (HTML)

Source: https://daisyui.com/components/badge

Demonstrates how DaisyUI badges can be integrated within standard text elements like headings and paragraphs. This allows for context-specific labeling.

```html
<h1>Heading 1 <span class="badge">Badge</span></h1>
<h2>Heading 2 <span class="badge">Badge</span></h2>
<h3>Heading 3 <span class="badge">Badge</span></h3>
<h4>Heading 4 <span class="badge">Badge</span></h4>
<h5>Heading 5 <span class="badge">Badge</span></h5>
<p>Paragraph text with a <span class="badge">Badge</span> inside.</p>
```

---

### DaisyUI Empty Badge (HTML)

Source: https://daisyui.com/components/badge

Shows the implementation of an empty badge in DaisyUI. This can be used as a placeholder or for visual elements that do not contain text content.

```html
<div class="badge"></div>
<div class="badge badge-outline"></div>
<div class="badge badge-soft"></div>
<div class="badge badge-ghost"></div>
<div class="badge badge-dash"></div>
```

---

### DaisyUI Responsive Menu (Vertical/Horizontal)

Source: https://daisyui.com/components/menu

Shows a menu that displays vertically on small screens and horizontally on larger screens. Achieved using DaisyUI's responsive modifiers.

```html
<ul class="menu menu-vertical lg:menu-horizontal">
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>
```

---

### Dropdown in Navbar with daisyUI

Source: https://daisyui.com/components/dropdown

Shows how to implement a dropdown menu within a daisyUI navigation bar. This is useful for organizing navigation links or actions. The snippet utilizes the 'dropdown' and 'navbar' components from daisyUI.

```html
<div class="navbar bg-base-100">
  <div class="dropdown">
    <label tabindex="0" class="btn btn-ghost lg:btn-circle">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16m-7 6h7"
        />
      </svg>
    </label>
    <ul
      tabindex="0"
      class="menu menu-sm dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
    >
      <li>
        <a>Dropdown</a>
        <ul class="p-2">
          <li><a>Item 1</a></li>
          <li><a>Item 2</a></li>
        </ul>
      </li>
    </ul>
  </div>
  <div class="flex-1 px-2 mx-2">daisyUI</div>
  <div class="flex-none">
    <button class="btn btn-square btn-ghost">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </button>
  </div>
</div>
```

---

### DaisyUI Navbar: Menu and Submenu

Source: https://daisyui.com/components/navbar

Shows how to implement a DaisyUI Navbar with a dropdown menu, including parent items and nested submenus. This is useful for complex navigation structures.

```html
<div class="navbar bg-base-100 rounded-box">
  <div class="flex-1 px-2 mx-2">
    <span class="text-lg font-bold">daisyUI</span>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal px-1">
      <li><a>Link</a></li>
      <li>
        <details>
          <summary>Parent</summary>
          <ul class="p-2 bg-base-100 rounded-t-none">
            <li><a>Link 1</a></li>
            <li><a>Link 2</a></li>
          </ul>
        </details>
      </li>
    </ul>
  </div>
</div>
```

---

### DaisyUI Neutral Badge with Outline or Dash Style (HTML)

Source: https://daisyui.com/components/badge

Demonstrates the neutral badge style combined with outline and dash styles. These are intended for use on light backgrounds and feature dark text for better contrast.

```html
<!-- Neutral Outline Badge -->
<div class="badge badge-neutral badge-outline">Outline</div>

<!-- Neutral Dash Badge -->
<div class="badge badge-neutral badge-dash">Dash</div>
```

---

### DaisyUI Swap Component Syntax (Class Name)

Source: https://context7_llms

This variant of the DaisyUI swap component toggles visibility between two elements using a class name modifier, typically controlled via JavaScript. It also uses 'swap-on' and 'swap-off' for content visibility.

```html
<div class="swap {MODIFIER}">
  <div class="swap-on">{content when active}</div>
  <div class="swap-off">{content when inactive}</div>
</div>
```

---

### Basic FAB with Vertical Buttons - HTML

Source: https://context7_llms

A standard FAB with a main action button that expands to reveal three additional circular buttons arranged vertically. This is the default behavior for FABs.

```html
<div class="fab">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <button class="btn btn-lg btn-circle">{Icon1}</button>
  <button class="btn btn-lg btn-circle">{Icon2}</button>
  <button class="btn btn-lg btn-circle">{Icon3}</button>
</div>
```

---

### Utilizing Responsive Modifiers in DaisyUI

Source: https://daisyui.com/docs/v5

Highlights that all DaisyUI component modifiers are now responsive by default. This allows developers to use standard Tailwind CSS responsive prefixes like `md:` and `lg:` with any DaisyUI modifier for adaptive designs.

```HTML
<!-- Example: Button with different styles on different screen sizes -->
<button class="btn btn-primary md:btn-secondary lg:btn-accent">Responsive Button</button>
```

---

### Centered Footer with Social Icons and Copyright

Source: https://daisyui.com/components/footer

This snippet shows a centered footer containing links for 'About us', 'Contact', 'Jobs', and 'Press kit', followed by a copyright notice attributing the work to ACME Industries Ltd. It's suitable for a minimalist footer design. The dynamic year calculation is handled by a JavaScript expression.

```html
### Centered footer with social icons About us Contact Jobs Press kit Copyright
© 2025 - All right reserved by ACME Industries Ltd
```

```html
About us Contact Jobs Press kit Copyright © {new Date().getFullYear()} - All
right reserved by ACME Industries Ltd
```

```html
About us Contact Jobs Press kit Copyright © {new Date().getFullYear()} - All
right reserved by ACME Industries Ltd
```

```html
About us Contact Jobs Press kit Copyright © {new Date().getFullYear()} - All
right reserved by ACME Industries Ltd
```

```html
About us Contact Jobs Press kit Copyright © {new Date().getFullYear()} - All
right reserved by ACME Industries Ltd
```

---

### Checkbox Colors (HTML)

Source: https://daisyui.com/components/checkbox

Demonstrates how to apply various color themes to checkboxes using daisyUI's color utility classes. These classes change the primary color of the checkbox.

````html
```html
<input type="checkbox" class="checkbox checkbox-primary" />
<input type="checkbox" class="checkbox checkbox-secondary" />
<input type="checkbox" class="checkbox checkbox-accent" />
<input type="checkbox" class="checkbox checkbox-neutral" />
<input type="checkbox" class="checkbox checkbox-success" />
<input type="checkbox" class="checkbox checkbox-warning" />
<input type="checkbox" class="checkbox checkbox-info" />
<input type="checkbox" class="checkbox checkbox-error" />
````

````

--------------------------------

### HTML: daisyUI List with Third Column Grow

Source: https://daisyui.com/components/list

This snippet illustrates how to configure the daisyUI List component so that the third column within a 'list-row' takes up the remaining space. This is achieved by applying appropriate modifiers and is useful for layouts requiring specific column sizing.

```html
<div class="list">
  <div class="list-row">
    <div>Most played songs this week</div>
    <div>01</div>
    <div>
      <img src="https://img.daisyui.com/images/profile/demo/1@94.webp" alt="User Avatar">
      Dio Lupa
      Remaining Reason
    </div>
  </div>
  <div class="list-row">
    <div></div>
    <div>02</div>
    <div>
      <img src="https://img.daisyui.com/images/profile/demo/4@94.webp" alt="User Avatar">
      Ellie Beilish
      Bears of a fever
    </div>
  </div>
  <div class="list-row">
    <div></div>
    <div>03</div>
    <div>
      <img src="https://img.daisyui.com/images/profile/demo/3@94.webp" alt="User Avatar">
      Sabrino Gardener
      Cappuccino
    </div>
  </div>
</div>
````

---

### Responsive Dialog Modal

Source: https://daisyui.com/components/modal

This modal is designed to be responsive, positioning at the bottom on small screens (SM) and in the middle on medium screens (MD). It can be opened using JavaScript.

```html
<!-- Responsive Modal goes bottom on SM screen size, goes middle on MD screen size -->
<dialog id="my_modal_5" class="modal modal-bottom sm:modal-middle">
  <div class="modal-box">
    <form method="dialog">
      <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
        ✕
      </button>
    </form>
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Press ESC key or click the button below to close</p>
    <div class="modal-action">
      <form method="dialog"><button class="btn">Close</button></form>
    </div>
  </div>
</dialog>
<button class="btn" onclick="my_modal_5.showModal()">open modal</button>
```

---

### Link Styling with Hover Underline in daisyUI

Source: https://daisyui.com/components/link

Illustrates how to make a link's underline appear only when the user hovers over it. This is achieved by combining the 'link' class with the 'link-hover' class for a more interactive user experience.

```html
<a href="#" class="link link-hover">Click me</a>
```

---

### DaisyUI Icon-Only Menu with Tooltips

Source: https://daisyui.com/components/menu

Demonstrates an icon-only menu where each icon is paired with a tooltip for better usability. Applicable to both vertical and horizontal menu orientations.

```html
<ul class="menu menu-vertical w-16 p-2">
  <li>
    <a data-tooltip="Tooltip 1">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a data-tooltip="Tooltip 2">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a data-tooltip="Tooltip 3">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
</ul>
```

```html
<ul class="menu menu-horizontal rounded-box">
  <li>
    <a data-tooltip="Tooltip 1">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a data-tooltip="Tooltip 2">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
  <li>
    <a data-tooltip="Tooltip 3">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12h18M3 6h18M3 18h18"
        />
      </svg>
    </a>
  </li>
</ul>
```

---

### Label for Date Input (HTML/Tailwind CSS)

Source: https://daisyui.com/components/label

Shows how to apply daisyUI labels to date input fields. This ensures that date input elements are clearly identified and accessible to users. Multiple styling options are presented.

```html
<label class="label">
  <span class="label-text">Publish date</span>
</label>
<input type="date" class="input input-bordered w-full max-w-xs" />
```

```html
<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">Event date</span>
  </label>
  <input type="date" class="input input-bordered" />
</div>
```

---

### DaisyUI Text Input: Sizes

Source: https://daisyui.com/components/input

Presents various size options (xs, sm, md, lg, xl) for DaisyUI text inputs. This allows developers to control the visual footprint of input fields to match design requirements.

````html
```html
<input
  type="text"
  placeholder="Extra small"
  class="input input-bordered input-xs w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Small"
  class="input input-bordered input-sm w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Medium"
  class="input input-bordered input-md w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Large"
  class="input input-bordered input-lg w-full max-w-xs"
/>
<input
  type="text"
  placeholder="Extra large"
  class="input input-bordered input-xl w-full max-w-xs"
/>
````

````

--------------------------------

### DaisyUI Centered Footer Layout

Source: https://daisyui.com/components/footer

Shows how to center-align the content within a DaisyUI footer using the 'footer-center' class. This is useful for creating a more compact and visually balanced footer layout.

```html
<footer class="footer footer-center p-10 bg-base-200 text-base-content rounded">
  <div>
    <svg width="50" height="50" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill-rule="evenodd" clip-rule="evenodd" class="inline-block fill-current"><path d="M22.672 17.011c0 3.903-3.167 7.011-7.011 7.011s-7.011-3.107-7.011-7.011c0-3.207 2.167-5.955 5.153-6.761V4.217c0-1.066-.868-1.931-1.931-1.931H9.716c-1.066 0-1.931.865-1.931 1.931v4.023c-1.493.697-2.868 1.734-3.838 3.013-.97 1.279-1.484 2.829-1.484 4.471 0 3.903 3.167 7.011 7.011 7.011s7.011-3.107 7.011-7.011c0-1.741-.514-3.389-1.484-4.728-.97-1.339-2.345-2.376-3.838-3.013v-4.023c0-1.066-.868-1.931-1.931-1.931H7.011c-1.066 0-1.931.865-1.931 1.931v10.207c0 1.652 1.101 3.107 2.688 3.563.968.291 1.987.437 3.011.437s2.043-.146 3.011-.437c1.587-.456 2.688-1.911 2.688-3.563V10.218c0-1.066-.868-1.931-1.931-1.931H9.716c-1.066 0-1.931.865-1.931 1.931v5.726c1.503.717 3.181 1.185 4.963 1.185 1.782 0 3.46-.468 4.963-1.185.774-.373 1.438-.929 1.974-1.616.536-.687.973-1.443 1.29-2.247.317-.804.475-1.657.475-2.539z"></path></svg>
    <p class="font-bold">
      ACME Industries Ltd.
    </p>
    <p>
      Providing reliable tech since 1992
    </p>
  </div>
  <div>
    <span class="footer-title">Services</span>
    <a class="link link-hover">Branding</a>
    <a class="link link-hover">Design</a>
    <a class="link link-hover">Marketing</a>
    <a class="link link-hover">Advertisement</a>
  </div>
  <div>
    <span class="footer-title">Company</span>
    <a class="link link-hover">About us</a>
    <a class="link link-hover">Contact</a>
    <a class="link link-hover">Jobs</a>
    <a class="link link-hover">Press kit</a>
  </div>
  <div>
    <span class="footer-title">Legal</span>
    <a class="link link-hover">Terms of use</a>
    <a class="link link-hover">Privacy policy</a>
    <a class="link link-hover">Cookie policy</a>
  </div>
</footer>
````

---

### Radial Progress with Custom Size and Thickness (React TSX)

Source: https://daisyui.com/components/radial-progress

This React (TSX) code illustrates customizing the size and thickness of the radial-progress component using '--size' and '--thickness' CSS variables for a tailored visual representation.

```tsx
function RadialProgressCustomSizeThicknessTsx() {
  return (
    <div
      className="radial-progress"
      style={
        {
          "--value": 70,
          "--size": "10rem",
          "--thickness": "1rem",
        } as React.CSSProperties
      }
      role="progressbar"
    >
      70%
    </div>
  );
}
```

---

### Countdown Component with 2 or 3 Digits (CSS)

Source: https://daisyui.com/components/countdown

Shows how to configure the DaisyUI Countdown component to display a minimum of 2 or 3 digits using CSS variables. This ensures consistent formatting regardless of the number's magnitude.

```css
--digits: 2;
/* or */
--digits: 3;
```

---

### Toast Component for Corner Notifications

Source: https://context7_llms

A wrapper component used to stack elements and position them in the corner of the page. Modifiers can be used to control the placement of the toast, such as 'toast-start', 'toast-center', or 'toast-end'.

```html
<div class="toast {MODIFIER}">{CONTENT}</div>
```

---

### Dialog Modal with Custom Width

Source: https://daisyui.com/components/modal

Demonstrates how to create a dialog modal with a customizable width using DaisyUI's 'w-' and 'max-w-' utility classes. The modal can be opened programmatically.

```html
<!-- Dialog modal with custom width -->
<dialog id="my_modal_4" class="modal">
  <div class="modal-box w-11/12 max-w-5xl">
    <form method="dialog">
      <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
        ✕
      </button>
    </form>
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Click the button below to close</p>
    <div class="modal-action">
      <form method="dialog"><button class="btn">Close</button></form>
    </div>
  </div>
</dialog>
<button class="btn" onclick="my_modal_4.showModal()">open modal</button>
```

---

### DaisyUI Badge Color Variations (HTML)

Source: https://daisyui.com/components/badge

Illustrates how to apply different color schemes to DaisyUI badges using color-related classes. This allows for semantic meaning and visual distinction based on status or type.

```html
<!-- Neutral Badge -->
<div class="badge badge-neutral">Neutral</div>

<!-- Primary Badge -->
<div class="badge badge-primary">Primary</div>

<!-- Secondary Badge -->
<div class="badge badge-secondary">Secondary</div>

<!-- Accent Badge -->
<div class="badge badge-accent">Accent</div>

<!-- Info Badge -->
<div class="badge badge-info">Info</div>

<!-- Success Badge -->
<div class="badge badge-success">Success</div>

<!-- Warning Badge -->
<div class="badge badge-warning">Warning</div>

<!-- Error Badge -->
<div class="badge badge-error">Error</div>
```

---

### DaisyUI Modal with Anchor Links (Alternative)

Source: https://daisyui.com/components/modal

This is another variation of a DaisyUI modal controlled by anchor links, specifically demonstrating the 'Close' functionality within the modal box. It's a common pattern for simple modal interactions.

```html
<label for="my_modal" class="btn">open modal</label>

<input type="checkbox" id="my_modal" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with anchor links</p>
    <div class="modal-action">
      <label for="my_modal" class="btn">Close</label>
    </div>
  </div>
</div>
```

---

### DaisyUI Indeterminate Toggle (JavaScript)

Source: https://daisyui.com/components/toggle

Explains how to programmatically set a DaisyUI toggle to an indeterminate state using JavaScript. This is useful for representing a state that is neither fully on nor fully off.

```javascript
document.getElementById("my-toggle").indeterminate = true;
```

---

### DaisyUI Text Input: With Fieldset and Legend

Source: https://daisyui.com/components/input

Illustrates using a fieldset and legend with DaisyUI text inputs to group related form elements semantically. This enhances accessibility and organization.

````html
```html
<fieldset class="w-full">
  <legend class="text-lg font-semibold">What is your name?</legend>
  <label class="input input-bordered flex items-center gap-2">
    <input type="text" class="grow" placeholder="Optional" />
  </label>
</fieldset>
````

````

--------------------------------

### Mask Component for Shaping Content with DaisyUI

Source: https://context7_llms

Illustrates how to use the DaisyUI 'mask' component to apply shapes to elements, particularly images. Requires a style modifier (e.g., 'mask-circle', 'mask-heart') and can be combined with size classes.

```html
<img class="mask mask-squircle w-32 h-32" src="{image-url}" alt="Shaped Image" />
````

---

### FAB Flower Arrangement - HTML

Source: https://context7_llms

This FAB uses the 'fab-flower' class to arrange the expanding buttons in a quarter-circle shape instead of a vertical line. It also includes a main action button.

```html
<div class="fab fab-flower">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <button class="fab-main-action btn btn-circle btn-lg">
    {IconMainAction}
  </button>
  <button class="btn btn-lg btn-circle">{Icon1}</button>
  <button class="btn btn-lg btn-circle">{Icon2}</button>
  <button class="btn btn-lg btn-circle">{Icon3}</button>
</div>
```

---

### Footer Component Structure in HTML

Source: https://context7_llms

The footer component is designed to hold logos, copyright information, and navigation links. It uses the 'footer' class and can include 'footer-title' for section headings. Modifiers for placement and direction are optional.

```html
<footer class="footer p-10 bg-neutral text-neutral-content">
  <div>
    <svg
      width="50"
      height="50"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
      fill="currentColor"
    >
      <path d="M12 2L0 22h24L12 2zM3.535 20 12 3.878 20.465 20H3.535z"></path>
    </svg>
    <p>ACME Industries Ltd.<br />Providing reliable tech since 1992</p>
  </div>
  <div>
    <span class="footer-title">Services</span>
    <a class="link link-hover">Branding</a>
    <a class="link link-hover">Design</a>
    <a class="link link-hover">Marketing</a>
    <a class="link link-hover">Advertisement</a>
  </div>
  <div>
    <span class="footer-title">Company</span>
    <a class="link link-hover">About us</a>
    <a class="link link-hover">Contact</a>
    <a class="link link-hover">Jobs</a>
    <a class="link link-hover">Press kit</a>
  </div>
  <div>
    <span class="footer-title">Legal</span>
    <a class="link link-hover">Terms of use</a>
    <a class="link link-hover">Privacy policy</a>
    <a class="link link-hover">Cookie policy</a>
  </div>
</footer>
```

---

### DaisyUI Modal with Anchor Link

Source: https://daisyui.com/components/modal

This snippet demonstrates how to create a modal that is triggered by clicking an anchor link. It utilizes HTML structure and specific DaisyUI classes for functionality. No external JavaScript is explicitly required for this basic implementation.

```html
<label for="my_modal_8" class="btn">open modal</label>

<input type="checkbox" id="my_modal_8" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with anchor links</p>
    <div class="modal-action">
      <label for="my_modal_8" class="btn">Yay!</label>
    </div>
  </div>
</div>
```

---

### HTML Dialog Modal with JS Open/Close (DaisyUI)

Source: https://daisyui.com/components/modal

Demonstrates how to use an HTML dialog element as a modal with DaisyUI. The modal is opened using JavaScript's `showModal()` method and can be closed programmatically with `close()`. It includes a button to trigger the opening and a close button within the modal. Accessibility is enhanced by allowing closure via the ESC key.

```html
<button class="btn" onclick="my_modal_1.showModal()">open modal</button>
<dialog id="my_modal_1" class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Press ESC key or click the button below to close</p>
    <div class="modal-action">
      <form method="dialog">
        <button class="btn">Close</button>
      </form>
    </div>
  </div>
</dialog>
```

```javascript
// Open the modal using document.getElementById('ID').showModal() method
document.getElementById("my_modal_1").showModal();
```

---

### DaisyUI Button Color Variants

Source: https://daisyui.com/components/button

Showcases the range of predefined color options available for DaisyUI buttons, including neutral, primary, secondary, accent, info, success, warning, and error. Each color variant is applied using corresponding class names.

```html
<button class="btn btn-neutral">Neutral</button>
<button class="btn btn-primary">Primary</button>
<button class="btn btn-secondary">Secondary</button>
<button class="btn btn-accent">Accent</button>
<button class="btn btn-info">Info</button>
<button class="btn btn-success">Success</button>
<button class="btn btn-warning">Warning</button>
<button class="btn btn-error">Error</button>
```

---

### Fieldset Component - HTML

Source: https://context7_llms

Defines a fieldset component for grouping related form elements. It includes a legend for the title and a paragraph for the description, with specific DaisyUI classes.

```html
<fieldset class="fieldset">
  <legend class="fieldset-legend">{title}</legend>
  {CONTENT}
  <p class="label">{description}</p>
</fieldset>
```

---

### Applying 'soft' and 'dash' Styles to DaisyUI Components

Source: https://daisyui.com/docs/v5

Shows how to apply new 'soft' and 'dash' style modifiers to DaisyUI components such as buttons and badges. The 'soft' style provides a gentler appearance, while 'dash' adds a dashed border.

```HTML
<!-- Example for a soft button -->
<button class="btn btn-soft">Soft Button</button>

<!-- Example for a dashed alert -->
<div class="alert alert-dash">Dashed Alert</div>
```

---

### Basic Countdown Component Usage (HTML)

Source: https://daisyui.com/components/countdown

Demonstrates the basic HTML structure for the DaisyUI Countdown component. It requires JavaScript to update the span text and the '--value' CSS variable. The value must be between 0 and 999.

```html
<div class="countdown">
  <span></span>
</div>
```

---

### Checkbox with Custom Colors (HTML)

Source: https://daisyui.com/components/checkbox

Demonstrates how to apply custom colors to checkboxes by defining and using CSS variables. This allows for more granular control over the checkbox appearance.

````html
```html
<input type="checkbox" class="checkbox" style="--chkbg: #f00; --chkfg: #fff;" />
<input
  type="checkbox"
  class="checkbox checkbox-primary"
  style="--chkbg: #0f0; --chkfg: #fff;"
/>
<input
  type="checkbox"
  class="checkbox checkbox-accent"
  style="--chkbg: #00f; --chkfg: #fff;"
/>
````

````

--------------------------------

### DaisyUI Toast Placement: Bottom-Start

Source: https://daisyui.com/components/toast

Shows how to set the DaisyUI Toast component to the bottom-start (bottom-left) position. This uses the 'toast-bottom' and 'toast-start' classes, with 'toast-bottom' being the default vertical alignment.

```html
<div class="toast toast-start toast-bottom">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
````

---

### DaisyUI Disabled Toggle (HTML)

Source: https://daisyui.com/components/toggle

Shows how to create a disabled DaisyUI toggle switch by adding the 'disabled' attribute to the input element.

```html
<input type="checkbox" class="toggle" checked disabled />
<input type="checkbox" class="toggle" disabled />
```

---

### DaisyUI Text Input: Search with Icon

Source: https://daisyui.com/components/input

Shows a search input field styled with an icon using DaisyUI. This provides a clear visual cue for the search functionality within a form.

````html
```html
<label class="input input-bordered flex items-center gap-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 16 16"
    fill="currentColor"
    class="w-4 h-4 opacity-70"
  >
    <path
      fill-rule="evenodd"
      d="M9.965 11.026a2 2 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z"
      clip-rule="evenodd"
    />
  </svg>
  <input type="text" class="grow" placeholder="Search" />
</label>
````

````

--------------------------------

### Filter with HTML Form, Checkboxes, and Reset Button - daisyUI

Source: https://daisyui.com/components/filter

This snippet demonstrates a multi-select filter using checkboxes within an HTML form. Unlike radio buttons, checkboxes allow users to select multiple options simultaneously. The `filter` class is not needed here, only the checkboxes and the reset button.

```html
<form>
  <div class="checkbox-group">
    <input type="checkbox" id="multiselect1" class="filter-checkbox" checked>
    <label for="multiselect1">Multi-Select 1</label>

    <input type="checkbox" id="multiselect2" class="filter-checkbox">
    <label for="multiselect2">Multi-Select 2</label>

    <input type="checkbox" id="multiselect3" class="filter-checkbox">
    <label for="multiselect3">Multi-Select 3</label>
  </div>
  <button type="reset" class="filter-reset">Reset Filters</button>
</form>
````

---

### Dropdown using CSS focus (HTML)

Source: https://daisyui.com/components/dropdown

A dropdown implemented using CSS focus, where content appears when a focusable element (like a div with tabindex) is focused. This method is often used with a workaround for Safari's focus bug on buttons.

```html
<div tabindex="0" role="button" class="focus:bg-base-300">
  Click
  <ul class="absolute hidden group-focus:block bg-base-100 p-2">
    <li>Item 1</li>
    <li>Item 2</li>
  </ul>
</div>
```

---

### Checkbox with Fieldset and Label (HTML)

Source: https://daisyui.com/components/checkbox

Shows how to group checkboxes within a fieldset and associate them with labels for better accessibility and user experience. This structure is common for form elements.

````html
```html
<div class="form-control">
  <label class="label cursor-pointer">
    <span class="label-text">Remember me</span>
    <input type="checkbox" class="checkbox" checked />
  </label>
</div>
````

````

--------------------------------

### FAB and Speed Dial (Vertical) - With Labels and fab-close

Source: https://daisyui.com/components/fab

Illustrates a FAB configuration where the main action button transforms into a 'Close' button when the speed dial is open, and includes text labels for the speed dial items. This provides a user-friendly way to dismiss the speed dial.

```html
<div class="fab">
  <button class="fab-close">✕</button>
  <div class="fab-content">
    <div class="flex items-center gap-2">
      <label>Label A</label>
      <button class="btn btn-xs">A</button>
    </div>
    <div class="flex items-center gap-2">
      <label>Label B</label>
      <button class="btn btn-xs">B</button>
    </div>
    <div class="flex items-center gap-2">
      <label>Label C</label>
      <button class="btn btn-xs">C</button>
    </div>
  </div>
</div>
````

---

### DaisyUI Horizontal Footer Layout

Source: https://daisyui.com/components/footer

Illustrates how to arrange footer columns horizontally using the 'footer-horizontal' class, making better use of horizontal space on larger screens. This is often used in conjunction with responsive design breakpoints.

```html
<footer class="footer p-10 bg-base-200 text-base-content">
  <div class="grid grid-flow-col gap-4">
    <a>
      <svg
        width="50"
        height="50"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
        fill-rule="evenodd"
        clip-rule="evenodd"
        class="inline-block fill-current"
      >
        <path
          d="M22.672 17.011c0 3.903-3.167 7.011-7.011 7.011s-7.011-3.107-7.011-7.011c0-3.207 2.167-5.955 5.153-6.761V4.217c0-1.066-.868-1.931-1.931-1.931H9.716c-1.066 0-1.931.865-1.931 1.931v4.023c-1.493.697-2.868 1.734-3.838 3.013-.97 1.279-1.484 2.829-1.484 4.471 0 3.903 3.167 7.011 7.011 7.011s7.011-3.107 7.011-7.011c0-1.741-.514-3.389-1.484-4.728-.97-1.339-2.345-2.376-3.838-3.013v-4.023c0-1.066-.868-1.931-1.931-1.931H7.011c-1.066 0-1.931.865-1.931 1.931v10.207c0 1.652 1.101 3.107 2.688 3.563.968.291 1.987.437 3.011.437s2.043-.146 3.011-.437c1.587-.456 2.688-1.911 2.688-3.563V10.218c0-1.066-.868-1.931-1.931-1.931H9.716c-1.066 0-1.931.865-1.931 1.931v5.726c1.503.717 3.181 1.185 4.963 1.185 1.782 0 3.46-.468 4.963-1.185.774-.373 1.438-.929 1.974-1.616.536-.687.973-1.443 1.29-2.247.317-.804.475-1.657.475-2.539z"
        ></path>
      </svg>
    </a>
    <div>
      <span class="footer-title">Services</span>
      <a class="link link-hover">Branding</a>
      <a class="link link-hover">Design</a>
      <a class="link link-hover">Marketing</a>
      <a class="link link-hover">Advertisement</a>
    </div>
    <div>
      <span class="footer-title">Company</span>
      <a class="link link-hover">About us</a>
      <a class="link link-hover">Contact</a>
      <a class="link link-hover">Jobs</a>
      <a class="link link-hover">Press kit</a>
    </div>
    <div>
      <span class="footer-title">Legal</span>
      <a class="link link-hover">Terms of use</a>
      <a class="link link-hover">Privacy policy</a>
      <a class="link link-hover">Cookie policy</a>
    </div>
  </div>
</footer>
```

---

### DaisyUI Stats Component Syntax

Source: https://context7_llms

The DaisyUI stats component displays statistical information. It can be modified with direction classes and contains individual stat elements with titles, values, and descriptions.

```html
<div class="stats {MODIFIER}">
  <div class="stat">{CONTENT}</div>
</div>
```

---

### DaisyUI Status Component Syntax

Source: https://context7_llms

The DaisyUI status component is a small icon used to visually represent an element's current status (e.g., online, offline, error). It supports various color and size modifiers.

```html
<span class="status {MODIFIER}"></span>
```

---

### DaisyUI Timeline with Bottom Side Only (HTML)

Source: https://daisyui.com/components/timeline

Illustrates a DaisyUI timeline where all content is positioned on the bottom side. This layout is suitable for simpler timelines where a dual-sided display is not required.

```html
<ul class="timeline">
  <li>
    <div class="timeline-start">First Macintosh computer</div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>1984</time>
    </div>
  </li>
  <li>
    <div class="timeline-start">iMac</div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>1998</time>
    </div>
  </li>
  <li>
    <div class="timeline-start">iPod</div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2001</time>
    </div>
  </li>
  <li>
    <div class="timeline-start">iPhone</div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2007</time>
    </div>
  </li>
  <li>
    <div class="timeline-start">Apple Watch</div>
    <div class="timeline-middle">
      <div class="w-40"></div>
    </div>
    <div class="timeline-end">
      <time>2015</time>
    </div>
  </li>
</ul>
```

---

### DaisyUI Toast Placement: Top-Center

Source: https://daisyui.com/components/toast

Illustrates positioning the DaisyUI Toast component at the top-center of the page. This configuration utilizes the 'toast-top' and 'toast-center' classes.

```html
<div class="toast toast-top toast-center">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Table Component Syntax

Source: https://context7_llms

The DaisyUI table component structures data in a tabular format. It supports modifiers for zebra striping, row/column pinning, and various sizes. An 'overflow-x-auto' class is recommended for horizontal scrolling on smaller screens.

```html
<div class="overflow-x-auto">
  <table class="table {MODIFIER}">
    <thead>
      <tr>
        <th></th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <th></th>
      </tr>
    </tbody>
  </table>
</div>
```

---

### HTML Progress Bar with Color Modifiers

Source: https://context7_llms

Displays a progress bar with customizable colors using DaisyUI's progress component. Requires 'value' and 'max' attributes. The optional {MODIFIER} can be one of the predefined color classes.

```html
<progress class="progress progress-primary" value="50" max="100"></progress>
<progress class="progress progress-secondary" value="75" max="100"></progress>
<progress class="progress progress-accent" value="25" max="100"></progress>
<progress class="progress progress-info" value="90" max="100"></progress>
<progress class="progress progress-success" value="10" max="100"></progress>
<progress class="progress progress-warning" value="60" max="100"></progress>
<progress class="progress progress-error" value="40" max="100"></progress>
<progress class="progress progress-neutral" value="100" max="100"></progress>
```

---

### Customizing Component Size Scale with CSS Variables

Source: https://daisyui.com/docs/v5

Illustrates how to customize the size scale of DaisyUI components using CSS variables. The `--size-field` variable controls the base size for inputs and buttons, while `--size-selector` affects checkboxes and toggles. This allows for global or theme-specific adjustments.

```CSS
/* Example for customizing input/button size */
:root {
  --size-field: 10px; /* Adjust base size for fields */
}

/* Example for customizing selector size */
:root {
  --size-selector: 8px; /* Adjust base size for selectors */
}
```

---

### Modal using Anchor Links (Legacy)

Source: https://daisyui.com/components/modal

A legacy method that controls modal visibility by adding a parameter to the URL. The modal only appears when the URL contains the specific parameter. Closing the modal scrolls the page to the top due to the anchor link. May not work well with SPAs.

```html
<!-- Modal using anchor link -->
<div id="my_modal_8" class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with anchor links</p>
    <div class="modal-action">
      <a href="#" class="btn">Yay!</a>
    </div>
  </div>
</div>
<a href="#my_modal_8" class="btn">open modal</a>
```

---

### DaisyUI Zebra-Striped Table

Source: https://daisyui.com/components/table

Illustrates how to apply zebra striping to table rows for better readability using the `table-zebra` modifier. This enhancement aids in visually distinguishing rows. It relies on standard HTML table structure with the added DaisyUI class.

```html
<div class="overflow-x-auto">
  <table class="table table-zebra">
    <!-- head -->
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
        <th>Favorite Color</th>
      </tr>
    </thead>
    <tbody>
      <!-- row 1 -->
      <tr>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
      </tr>
      <!-- row 2 -->
      <tr>
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
      </tr>
      <!-- row 3 -->
      <tr>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
      </tr>
    </tbody>
  </table>
</div>
```

---

### DaisyUI Card with Action on Top

Source: https://daisyui.com/components/card

A DaisyUI card variation where the primary action is placed at the top of the card body, above the main content. This can be used for quick access to a primary function associated with the card's content.

```html
<div class="card w-96 bg-base-100 shadow-xl">
  <div class="card-body">
    <h2 class="card-title">Card Title</h2>
    <p>We are using cookies for no reason.</p>
    <div class="card-actions">
      <button class="btn btn-primary">Primary Action</button>
    </div>
  </div>
</div>
```

---

### Toggle Requirement Validator (HTML)

Source: https://daisyui.com/components/validator

Shows how to validate a toggle switch using DaisyUI. Similar to checkboxes, the 'validator' class is used, and the 'validator-hint' indicates it is a required field.

```html
<input type="checkbox" class="toggle validator" />
<p class="validator-hint">Required</p>
```

---

### DaisyUI Table Sizing (Extra Small to Extra Large)

Source: https://daisyui.com/components/table

Shows how to adjust the size of DaisyUI tables using size modifiers like `table-xs`, `table-sm`, `table-md` (default), `table-lg`, and `table-xl`. This allows for responsive table design based on content and layout needs. Standard HTML table structure applies.

```html
<div class="overflow-x-auto">
  <table class="table table-xs">
    ...
  </table>
  <table class="table table-sm">
    ...
  </table>
  <table class="table">
    ...
  </table>
  <!-- Default table-md -->
  <table class="table table-lg">
    ...
  </table>
  <table class="table table-xl">
    ...
  </table>
</div>
```

---

### DaisyUI Drawer Toggle Button

Source: https://context7_llms

How to create a button that toggles a DaisyUI Drawer. A label element is used with its `for` attribute pointing to the `id` of the `drawer-toggle` input. Clicking this label will change the checked state of the input, thereby opening or closing the drawer.

```html
<label for="my-drawer" class="btn drawer-button">Open/close drawer</label>
```

---

### Large Text Countdown with Labels Underneath (HTML)

Source: https://daisyui.com/components/countdown

An alternative layout for displaying a countdown with labels, where the labels appear directly below the numerical values. This uses flexbox for arrangement.

```html
<div class="grid grid-flow-row gap-3 text-center auto-rows-max">
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-4xl">
      <span style="--value:15"></span>
    </span>
    days
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-4xl">
      <span style="--value:10"></span>
    </span>
    hours
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-4xl">
      <span style="--value:24"></span>
    </span>
    min
  </div>
  <div class="flex flex-col p-2 bg-base-100 rounded-box">
    <span class="countdown font-mono text-4xl">
      <span style="--value:59"></span>
    </span>
    sec
  </div>
</div>
```

---

### Toggle Switch for DaisyUI Forms

Source: https://context7_llms

A styled checkbox input that mimics a switch button. It supports various color and size modifiers to match the application's design. This component is useful for binary choices.

```html
<input type="checkbox" class="toggle {MODIFIER}" />
```

---

### DaisyUI Card with No Image

Source: https://daisyui.com/components/card

This snippet shows a DaisyUI card component that does not include an image. It focuses solely on text content, featuring a title, descriptive text, and a call-to-action button. This is suitable for content-heavy cards where visuals are not primary.

```html
<div class="card w-96 bg-base-100 shadow-xl">
  <div class="card-body">
    <h2 class="card-title">Card title!</h2>
    <p>
      A card component has a figure, a body part, and inside body there are
      title and actions parts
    </p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Buy Now</button>
    </div>
  </div>
</div>
```

---

### Textarea Component Size Variations - daisyUI

Source: https://daisyui.com/components/textarea

Demonstrates the different size variants available for the Textarea component in daisyUI, ranging from extra-small (xs) to extra-large (xl), including the default medium (md) size.

````html
```html
<textarea class="textarea textarea-xs" placeholder="Bio"></textarea>
<textarea class="textarea textarea-sm" placeholder="Bio"></textarea>
<textarea class="textarea textarea-md" placeholder="Bio"></textarea>
<textarea class="textarea textarea-lg" placeholder="Bio"></textarea>
<textarea class="textarea textarea-xl" placeholder="Bio"></textarea>
````

````

--------------------------------

### DaisyUI Rating with Heart Icon and Multiple Colors

Source: https://daisyui.com/components/rating

This snippet demonstrates using a heart icon ('mask-heart') for the rating component and applies multiple colors to different rating options, allowing for a visually distinct rating scale.

```html
<div class="rating">
  <input type="radio" name="rating-heart" class="mask mask-heart" />
  <input type="radio" name="rating-heart" class="mask mask-heart" checked />
  <input type="radio" name="rating-heart" class="mask mask-heart" />
  <input type="radio" name="rating-heart" class="mask mask-heart" />
  <input type="radio" name="rating-heart" class="mask mask-heart" />
</div>
<div class="rating">
  <input type="radio" name="rating-heart-colors" class="bg-red-400 mask mask-heart" />
  <input type="radio" name="rating-heart-colors" class="bg-orange-400 mask mask-heart" />
  <input type="radio" name="rating-heart-colors" class="bg-yellow-400 mask mask-heart" />
  <input type="radio" name="rating-heart-colors" class="bg-lime-400 mask mask-heart" checked />
  <input type="radio" name="rating-heart-colors" class="bg-green-400 mask mask-heart" />
</div>
````

---

### DaisyUI Rating Sizes (XS to XL)

Source: https://daisyui.com/components/rating

This snippet illustrates the different size variants for the DaisyUI Rating component, ranging from extra small ('rating-xs') to extra large ('rating-xl').

```html
<!-- Extra small -->
<div class="rating rating-xs">
  <input type="radio" name="rating-xs" class="mask mask-star-2" />
  <input type="radio" name="rating-xs" class="mask mask-star-2" checked />
  <input type="radio" name="rating-xs" class="mask mask-star-2" />
  <input type="radio" name="rating-xs" class="mask mask-star-2" />
  <input type="radio" name="rating-xs" class="mask mask-star-2" />
</div>

<!-- Small -->
<div class="rating rating-sm">
  <input type="radio" name="rating-sm" class="mask mask-star-2" />
  <input type="radio" name="rating-sm" class="mask mask-star-2" checked />
  <input type="radio" name="rating-sm" class="mask mask-star-2" />
  <input type="radio" name="rating-sm" class="mask mask-star-2" />
  <input type="radio" name="rating-sm" class="mask mask-star-2" />
</div>

<!-- Medium (Default) -->
<div class="rating rating-md">
  <input type="radio" name="rating-md" class="mask mask-star-2" />
  <input type="radio" name="rating-md" class="mask mask-star-2" checked />
  <input type="radio" name="rating-md" class="mask mask-star-2" />
  <input type="radio" name="rating-md" class="mask mask-star-2" />
  <input type="radio" name="rating-md" class="mask mask-star-2" />
</div>

<!-- Large -->
<div class="rating rating-lg">
  <input type="radio" name="rating-lg" class="mask mask-star-2" />
  <input type="radio" name="rating-lg" class="mask mask-star-2" checked />
  <input type="radio" name="rating-lg" class="mask mask-star-2" />
  <input type="radio" name="rating-lg" class="mask mask-star-2" />
  <input type="radio" name="rating-lg" class="mask mask-star-2" />
</div>

<!-- Extra large -->
<div class="rating rating-xl">
  <input type="radio" name="rating-xl" class="mask mask-star-2" />
  <input type="radio" name="rating-xl" class="mask mask-star-2" checked />
  <input type="radio" name="rating-xl" class="mask mask-star-2" />
  <input type="radio" name="rating-xl" class="mask mask-star-2" />
  <input type="radio" name="rating-xl" class="mask mask-star-2" />
</div>
```

---

### Carousel Structure (HTML)

Source: https://context7_llms

Shows the HTML structure for a basic carousel component in DaisyUI. Carousels are used to display images or content in a scrollable area.

```html
<div class="carousel w-full">
  <div id="slide1" class="carousel-item relative w-full">
    <img
      src="/images/stock/photo-1609621038510-57099756704b.jpg"
      class="w-full"
    />
    <div
      class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2"
    >
      <a href="#slide4" class="btn btn-circle">❮</a>
      <a href="#slide2" class="btn btn-circle">❯</a>
    </div>
  </div>
  <div id="slide2" class="carousel-item relative w-full">
    <img
      src="/images/stock/photo-1609621038510-57099756704b.jpg"
      class="w-full"
    />
    <div
      class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2"
    >
      <a href="#slide1" class="btn btn-circle">❮</a>
      <a href="#slide3" class="btn btn-circle">❯</a>
    </div>
  </div>
</div>
```

---

### DaisyUI Radio Button Sizes

Source: https://daisyui.com/components/radio

Illustrates how to control the size of DaisyUI radio buttons using size utility classes. Available sizes include extra small (xs), small (sm), medium (md), large (lg), and extra large (xl).

```html
<!-- Radio Extra Small -->
<input type="radio" name="radio-size" class="radio radio-xs" checked />

<!-- Radio Small -->
<input type="radio" name="radio-size" class="radio radio-sm" checked />

<!-- Radio Medium -->
<input type="radio" name="radio-size" class="radio radio-md" checked />

<!-- Radio Large -->
<input type="radio" name="radio-size" class="radio radio-lg" checked />

<!-- Radio Extra Large -->
<input type="radio" name="radio-size" class="radio radio-xl" checked />
```

---

### HTML Radial Progress Indicator

Source: https://context7_llms

Implements a circular progress indicator using a div element and CSS variables. The `--value` CSS variable controls the progress (0-100). Accessible via `aria-valuenow`. Size and thickness can be adjusted with `--size` and `--thickness`.

```html
<div
  class="radial-progress"
  style="--value:70;"
  aria-valuenow="70"
  role="progressbar"
>
  70%
</div>
<div
  class="radial-progress"
  style="--value:30; --size: 8rem; --thickness: 4px;"
  aria-valuenow="30"
  role="progressbar"
>
  30%
</div>
```

---

### FAB with Rectangle Buttons - HTML

Source: https://context7_llms

A FAB variant that uses rectangular buttons instead of circular ones. This allows for more content within the button, such as text labels, making them more descriptive.

```html
<div class="fab">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <button class="btn btn-lg">{Label1}</button>
  <button class="btn btn-lg">{Label2}</button>
  <button class="btn btn-lg">{Label3}</button>
</div>
```

---

### HTML Star Rating Input

Source: https://context7_llms

A star rating component built with radio buttons and 'mask-star' class. Each set needs a unique 'name' attribute. 'rating-half' and 'rating-hidden' modifiers are available for half-star selection and clearing.

```html
<div class="rating">
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" checked />
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" />
</div>
<div class="rating rating-half">
  <input type="radio" name="rating-2" class="bg-warning mask mask-star-2" />
  <input
    type="radio"
    name="rating-2"
    class="bg-warning mask mask-star-2"
    checked
  />
</div>
<div class="rating rating-hidden">
  <input type="radio" name="rating-3" class="mask mask-star" />
  <input type="radio" name="rating-3" class="mask mask-star" />
  <input type="radio" name="rating-3" class="mask mask-star" />
  <input type="radio" name="rating-3" class="mask mask-star" />
  <input type="radio" name="rating-3" class="mask mask-star" />
</div>
```

---

### HTML Select Dropdown

Source: https://context7_llms

A customizable select dropdown component. The {MODIFIER} can include style, color, and size classes.

```html
<select class="select select-bordered">
  <option disabled selected>Select option</option>
  <option>Option 1</option>
  <option>Option 2</option>
  <option>Option 3</option>
</select>
<select class="select select-primary select-lg">
  <option>Option A</option>
  <option>Option B</option>
</select>
```

---

### Vertical Stat Layout (HTML)

Source: https://daisyui.com/components/stat

This code snippet shows how to arrange Stat components vertically. This layout is typically used when space is limited or for specific design requirements.

```html
<div class="stats stats-vertical lg:stats-horizontal">
  <div class="stat">
    <div class="stat-title">Downloads</div>
    <div class="stat-value">31K</div>
    <div class="stat-desc">Jan 1st - Feb 1st</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Users</div>
    <div class="stat-value">4,200</div>
    <div class="stat-desc">↗︎ 400 (22%)</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Registers</div>
    <div class="stat-value">1,200</div>
    <div class="stat-desc">↘︎ 90 (14%)</div>
  </div>
</div>
```

---

### Responsive DaisyUI Card (Vertical/Horizontal)

Source: https://daisyui.com/components/card

A responsive DaisyUI card that displays vertically on small screens and horizontally on larger screens. It includes an image, a title, descriptive text, and a call-to-action button, adapting its layout for optimal viewing across different devices.

```html
<div
  class="card lg:card-side bg-base-100 shadow-xl max-w-xl flex-col lg:flex-row"
>
  <figure class="w-full lg:w-1/3">
    <img
      src="https://img.daisyui.com/images/stock/photo-1494232410401-ad00d5433cfa.webp"
      alt="Album"
      class="w-full h-full object-cover"
    />
  </figure>
  <div class="card-body lg:w-2/3 p-4">
    <h2 class="card-title">New album is released!</h2>
    <p>Click the button to listen on Spotiwhy app.</p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Listen</button>
    </div>
  </div>
</div>
```

---

### DaisyUI Radio Button with Custom Colors

Source: https://daisyui.com/components/radio

Explains how to apply custom colors to DaisyUI radio buttons using inline styles or CSS variables. This allows for greater design flexibility beyond the predefined color themes.

```html
<!-- Radio with custom color -->
<input
  type="radio"
  name="radio-custom"
  class="radio border-2 border-purple-600 text-purple-600 checked:bg-purple-600"
  checked
/>

<!-- Radio with custom color and label -->
<label class="cursor-pointer label">
  <span class="label-text">Custom Color Radio</span>
  <input
    type="radio"
    name="radio-custom-label"
    class="radio border-2 border-green-500 text-green-500 checked:bg-green-500"
  />
</label>
```

---

### Modal using Checkbox (Legacy)

Source: https://daisyui.com/components/modal

A legacy method for modal control using a hidden checkbox. Labels can toggle the checkbox to open/close the modal. This method might have compatibility issues with modern practices.

```html
<!-- Modal using checkbox -->
<input type="checkbox" id="my-modal-checkbox" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with a hidden checkbox!</p>
    <div class="modal-action">
      <label for="my-modal-checkbox" class="btn">Close!</label>
    </div>
  </div>
</div>
<label for="my-modal-checkbox" class="btn">open modal</label>
```

---

### Validator Class for Form Input Feedback

Source: https://context7_llms

Applies error or success styling to form elements based on validation rules. It includes the 'validator-hint' class for displaying feedback messages. Use this with input, select, and textarea elements.

```html
<input type="{type}" class="input validator" required />
<p class="validator-hint">Error message</p>
```

---

### DaisyUI Text Input: Email with Icon, Validator, and Button

Source: https://daisyui.com/components/input

Combines an email input with an icon, validation message, and an adjacent button for actions like joining a newsletter. This is a common pattern for subscription forms.

````html
```html
<label class="input input-bordered flex items-center gap-2">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 16 16"
    fill="currentColor"
    class="w-4 h-4 opacity-70"
  >
    <path
      d="M2.5 3A1.5 1.5 0 0 0 1 4.5v5.086a1.5 1.5 0 0 0 .437.971l3.147 3.147a1.5 1.5 0 0 0 2.072.012l3.147-3.147a1.5 1.5 0 0 0 .437-.971V4.5A1.5 1.5 0 0 0 11.5 3H2.5Z"
    />
    <path
      d="M1.75 5.75a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 .75.75v3.086a.75.75 0 0 1-.22.547l-3.147 3.147a.75.75 0 0 1-1.06.012l-3.147-3.147a.75.75 0 0 1-.22-.547V5.75Z"
    />
  </svg>
  <input type="email" class="grow" placeholder="Email" />
</label>
<button class="btn btn-primary">Join</button>
<label class="label">
  <span class="label-text-alt">Enter valid email address</span>
</label>
````

````

--------------------------------

### Modal with Backdrop Click to Close (Legacy)

Source: https://daisyui.com/components/modal

This legacy modal uses a hidden checkbox and a 'modal-backdrop' class to allow closing the modal by clicking outside of it. Labels are used to toggle the checkbox state.

```html
<!-- Modal that closes when clicked outside -->
<input type="checkbox" id="my-modal-checkbox-backdrop" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">This modal works with a hidden checkbox!</p>
    <div class="modal-action">
      <label for="my-modal-checkbox-backdrop" class="btn">Close</label>
    </div>
  </div>
</div>
<label for="my-modal-checkbox-backdrop" class="btn">open modal</label>
<label class="modal-backdrop" for="my-modal-checkbox-backdrop"></label>
````

---

### Disabled File Input - daisyUI

Source: https://daisyui.com/components/file-input

Shows how to disable the file input component, making it non-interactive and visually indicating that it cannot be used. This is useful for forms where file uploads are conditional or not yet allowed.

```html
<div>
  <input
    type="file"
    placeholder="disabled placeholder"
    class="file-input file-input-bordered w-full max-w-xs"
    disabled
  />
  <input
    type="file"
    class="file-input file-input-ghost w-full max-w-xs"
    disabled
  />
</div>
```

---

### DaisyUI Range Slider Sizes

Source: https://daisyui.com/components/range

Illustrates the different size variations available for the DaisyUI range slider component, including extra small (xs), small (sm), medium (md), large (lg), and extra large (xl).

```html
<input type="range" class="range range-xs" />
<input type="range" class="range range-sm" />
<input type="range" class="range range-md" />
<input type="range" class="range range-lg" />
<input type="range" class="range range-xl" />
```

---

### FAB and Speed Dial (Vertical) - With Text Labels

Source: https://daisyui.com/components/fab

Shows how to add text labels to the speed dial buttons. This structure includes the main FAB and the content for the speed dial buttons, with each button having associated text labels.

```html
<div class="fab">
  <button class="fab-main-action">F</button>
  <div class="fab-content">
    <div class="flex items-center gap-2">
      <label>Label A</label>
      <button class="btn btn-xs">A</button>
    </div>
    <div class="flex items-center gap-2">
      <label>Label B</label>
      <button class="btn btn-xs">B</button>
    </div>
    <div class="flex items-center gap-2">
      <label>Label C</label>
      <button class="btn btn-xs">C</button>
    </div>
  </div>
</div>
```

---

### Filter with HTML Form, Radio Buttons, and Reset Button - daisyUI

Source: https://daisyui.com/components/filter

This snippet demonstrates how to implement a filter using an HTML form with radio buttons. When a radio button is selected, other options are hidden, and a reset button appears next to the selected one. This is ideal for single-choice filtering.

```html
<form>
  <div class="filter">
    <input
      type="radio"
      name="filter-group"
      id="option1"
      class="filter-radio"
      checked
    />
    <label for="option1">Option 1</label>

    <input type="radio" name="filter-group" id="option2" class="filter-radio" />
    <label for="option2">Option 2</label>

    <input type="radio" name="filter-group" id="option3" class="filter-radio" />
    <label for="option3">Option 3</label>

    <button type="reset" class="filter-reset">Reset</button>
  </div>
</form>
```

---

### HTML Range Slider Input

Source: https://context7_llms

A customizable range slider input using the 'range' class. Requires 'min', 'max', and 'value' attributes. Color and size modifiers can be applied.

```html
<input type="range" min="0" max="100" value="40" class="range range-primary" />
<input
  type="range"
  min="0"
  max="200"
  value="150"
  class="range range-lg range-success"
/>
```

---

### HTML Radio Button Group

Source: https://context7_llms

Creates a group of radio buttons using the 'radio' class. Each radio input should have a unique 'name' attribute for proper grouping. Optional color and size modifiers can be applied.

```html
<input type="radio" name="radio-group-1" class="radio radio-primary" checked />
<input type="radio" name="radio-group-1" class="radio radio-secondary" />
<input type="radio" name="radio-group-1" class="radio radio-lg radio-accent" />
<input type="radio" name="radio-group-2" class="radio radio-sm radio-warning" />
```

---

### Filter Component without HTML Form in HTML

Source: https://context7_llms

This demonstrates the filter component used without an HTML form, employing a div container instead. It includes the 'filter-reset' class for the reset radio button and standard radio inputs for options.

```html
<div class="filter">
  <input class="btn filter-reset" type="radio" name="my-radio" aria-label="×" />
  <input class="btn" type="radio" name="my-radio" aria-label="Tab 1 title" />
  <input class="btn" type="radio" name="my-radio" aria-label="Tab 2 title" />
</div>
```

---

### HTML Stat Display

Source: https://context7_llms

Component for displaying statistics or numerical data in a structured block. Uses 'stats' for the container and parts like 'stat', 'stat-title', 'stat-value', etc. Supports horizontal and vertical layouts.

```html
<div class="stats stats-vertical shadow">
  <div class="stat">
    <div class="stat-title">Total Page Views</div>
    <div class="stat-value">89,400</div>
    <div class="stat-desc">21% more than last month</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Users</div>
    <div class="stat-value">4,200</div>
    <div class="stat-desc">40% more than last month</div>
  </div>
</div>
<div class="stats stats-horizontal shadow">
  <div class="stat place-items-center">
    <div class="stat-title">Downloads</div>
    <div class="stat-value">31K</div>
    <div class="stat-desc">Jan 1st - Feb 1st</div>
  </div>
</div>
```

---

### Radial Progress with Custom Color (Vue)

Source: https://daisyui.com/components/radial-progress

Demonstrates how to apply a custom color to the radial-progress component in Vue.js by setting the '--progress-color' CSS variable for a 70% progress indicator.

```vue
<template>
  <div
    class="radial-progress"
    :style="{ '--value': 70, '--progress-color': 'oklch(var(--p))' }"
    role="progressbar"
  >
    70%
  </div>
</template>

<script setup lang="ts">
// No script needed for basic example
</script>

<style>
/* DaisyUI styles are assumed to be globally imported */
</style>
```

---

### DaisyUI Neutral Button with Outline/Dash

Source: https://daisyui.com/components/button

Shows how to combine the neutral button style with outline or dash modifiers for specific design needs. These buttons are intended for light backgrounds and offer a subtle visual distinction.

```html
<button class="btn btn-neutral btn-outline">Outline</button>
<button class="btn btn-neutral btn-dash">Dash</button>
```

---

### Apply Bold Star Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a bold star mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a thicker star shape. No external JavaScript dependencies are required.

```css
.mask-star-2 {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,5 L 61.2,37.8 L 96.3,37.8 L 67.5,58.3 L 78.7,91.2 L 50,70.5 L 21.3,91.2 L 32.5,58.3 L 3.7,37.8 L 38.8,37.8 Z'/%3E%3C/svg%3E");
}
```

---

### DaisyUI Stack Component: Stacked Cards (Top Direction)

Source: https://daisyui.com/components/stack

Demonstrates stacking card elements vertically with the 'stack-top' modifier, aligning them to the top of the stack. This affects how the elements are positioned relative to each other within the stack.

```html
<div class="stack stack-top">
  <div>A</div>
  <div>B</div>
  <div>C</div>
</div>
```

---

### Apply Triangle Mask (Pointing Top) with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a triangle mask pointing upwards to an element using DaisyUI CSS classes. This utility crops the content of the element into a top-pointing triangle shape. No external JavaScript dependencies are required.

```css
.mask-triangle {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,10 L 90,90 L 10,90 Z'/%3E%3C/svg%3E");
}
```

---

### DaisyUI Image Card with Side Image

Source: https://daisyui.com/components/card

This DaisyUI card features an image positioned on the side, with the text content adjacent to it. It's suitable for displaying articles, products, or media where the image is a key visual element alongside descriptive text and a call to action.

```html
<div class="card card-side bg-base-100 shadow-xl max-w-md">
  <figure>
    <img
      src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
      alt="Movie"
    />
  </figure>
  <div class="card-body">
    <h2 class="card-title">New movie is released!</h2>
    <p>Click the button to watch on Jetflix app.</p>
    <div class="card-actions justify-end">
      <button class="btn btn-primary">Watch</button>
    </div>
  </div>
</div>
```

---

### Apply Triangle Mask (Pointing Left) with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a triangle mask pointing leftwards to an element using DaisyUI CSS classes. This utility crops the content of the element into a left-pointing triangle shape. No external JavaScript dependencies are required.

```css
.mask-triangle-3 {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 90,50 L 10,10 L 10,90 Z'/%3E%3C/svg%3E");
}
```

---

### FAB with Close Button - HTML

Source: https://context7_llms

This FAB includes a specific 'close' button that appears when the FAB is open, replacing the original icon. It is styled with a red error button and a '✕' symbol.

```html
<div class="fab">
  <div tabindex="0" role="button" class="btn btn-lg btn-circle btn-primary">
    {IconOriginal}
  </div>
  <div class="fab-close">
    Close <span class="btn btn-circle btn-lg btn-error">✕</span>
  </div>
  <div>{Label1}<button class="btn btn-lg btn-circle">{Icon1}</button></div>
  <div>{Label2}<button class="btn btn-lg btn-circle">{Icon2}</button></div>
  <div>{Label3}<button class="btn btn-lg btn-circle">{Icon3}</button></div>
</div>
```

---

### Dialog Modal with Close Button

Source: https://daisyui.com/components/modal

A basic dialog modal with a close button in the corner. It can be opened using JavaScript's showModal() method. Closing is handled by the '✕' button or by pressing the ESC key.

```html
<!-- Dialog modal with a close button at corner -->
<dialog id="my_modal_3" class="modal">
  <div class="modal-box">
    <form method="dialog">
      <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
        ✕
      </button>
    </form>
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Press ESC key or click on ✕ button to close</p>
  </div>
</dialog>
<button class="btn" onclick="my_modal_3.showModal()">open modal</button>
```

---

### Apply Pentagon Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a pentagon mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a pentagon (5-sided polygon) shape. No external JavaScript dependencies are required.

```css
.mask-pentagon {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,10 C 55.5228,10 60,14.4772 60,20 L 60,40 C 60,45.5228 55.5228,50 50,50 C 44.4772,50 40,45.5228 40,40 L 40,20 C 40,14.4772 44.4772,10 50,10 Z M 70.7107,40 C 74.6401,40 77.7107,43.0706 77.7107,47 L 77.7107,53 C 77.7107,56.9294 74.6401,60 70.7107,60 L 50,80 C 44.4772,80 40,75.5228 40,70 L 40,47 C 40,43.0706 43.0706,40 47,40 L 70.7107,40 Z'/%3E%3C/svg%3E");
}
```

---

### FAB and Speed Dial (Flower Shape) - Basic Structure

Source: https://daisyui.com/components/fab

Shows the HTML structure for a FAB that opens its speed dial buttons in a quarter-circle ('flower') arrangement. This uses the `fab-flower` modifier to alter the display.

```html
<div class="fab fab-flower">
  <button class="fab-main-action">F</button>
  <div class="fab-content">
    <button class="btn btn-xs">A</button>
    <button class="btn btn-xs">B</button>
    <button class="btn btn-xs">C</button>
  </div>
</div>
```

---

### Centered DaisyUI Card with Neutral Color

Source: https://daisyui.com/components/card

This DaisyUI card is centered and uses a neutral color scheme. It includes a title and two action buttons, 'Accept' and 'Deny'. This is useful for confirmation dialogs or cookie consent banners.

```html
<div class="card w-96 bg-neutral text-neutral-content shadow-xl text-center">
  <div class="card-body">
    <h2 class="card-title">Cookies!</h2>
    <p>We are using cookies for no reason.</p>
    <div class="card-actions justify-center">
      <button class="btn btn-primary">Accept</button>
      <button class="btn">Deny</button>
    </div>
  </div>
</div>
```

---

### Apply Star Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a star mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a star shape. No external JavaScript dependencies are required.

```css
.mask-star {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,0 L 61.8,38.2 L 100,38.2 L 68.2,58.3 L 75,100 L 50,78.2 L 25,100 L 31.8,58.3 L 0,38.2 L 38.2,38.2 Z'/%3E%3C/svg%3E");
}
```

---

### Apply Circle Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a circle mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a perfect circle. No external JavaScript dependencies are required.

```css
.mask-circle {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Ccircle cx='50' cy='50' r='50'/%3E%3C/svg%3E");
}
```

---

### Apply Triangle Mask (Pointing Right) with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a triangle mask pointing rightwards to an element using DaisyUI CSS classes. This utility crops the content of the element into a right-pointing triangle shape. No external JavaScript dependencies are required.

```css
.mask-triangle-4 {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 10,50 L 90,10 L 90,90 Z'/%3E%3C/svg%3E");
}
```

---

### Apply Triangle Mask (Pointing Down) with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a triangle mask pointing downwards to an element using DaisyUI CSS classes. This utility crops the content of the element into a bottom-pointing triangle shape. No external JavaScript dependencies are required.

```css
.mask-triangle-2 {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,90 L 10,10 L 90,10 Z'/%3E%3C/svg%3E");
}
```

---

### Apply Squircle Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a squircle mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a squircle shape. No external JavaScript dependencies are required.

```css
.mask-squircle {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 85.4477,21.6518 C 89.7135,21.6518 93.1266,25.0649 93.1266,29.3307 L 93.1266,70.6693 C 93.1266,74.9351 89.7135,78.3482 85.4477,78.3482 L 14.5523,78.3482 C 10.2865,78.3482 6.8734,74.9351 6.8734,70.6693 L 6.8734,29.3307 C 6.8734,25.0649 10.2865,21.6518 14.5523,21.6518 Z M 75.8168,34.1832 C 74.1769,32.5433 71.7675,32.5433 70.1276,34.1832 L 50.0000,54.3105 L 29.8724,34.1832 C 28.2325,32.5433 25.8231,32.5433 24.1832,34.1832 C 22.5433,35.8231 22.5433,38.2325 24.1832,39.8724 L 45.7017,61.4000 L 45.7017,61.4000 C 47.3416,63.0400 49.7510,63.0400 51.3909,61.4000 L 51.3909,61.4000 L 75.8168,39.8724 C 77.4567,38.2325 77.4567,35.8231 75.8168,34.1832 Z'/%3E%3C/svg%3E');
}
```

---

### DaisyUI Toast Placement: Middle-Center

Source: https://daisyui.com/components/toast

Illustrates the DaisyUI Toast component placed in the middle-center (vertically and horizontally centered) of the screen. This configuration uses 'toast-middle' and 'toast-center'.

```html
<div class="toast toast-center toast-middle">
  <div class="alert alert-info">
    <div>
      <span>New mail arrived.</span>
      <span>Message sent successfully.</span>
    </div>
  </div>
</div>
```

---

### DaisyUI Stack Component: Stacked Cards (End Direction)

Source: https://daisyui.com/components/stack

Shows how to stack card elements vertically using the 'stack-end' modifier, aligning them to the end (right) side of the stack. This is useful for right-aligned stacked content.

```html
<div class="stack stack-end">
  <div>A</div>
  <div>B</div>
  <div>C</div>
</div>
```

---

### Apply Square Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a square mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a square shape. No external JavaScript dependencies are required.

```css
.mask-square {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 10,10 H 90 V 90 H 10 Z'/%3E%3C/svg%3E");
}
```

---

### Apply Diamond Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a diamond mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a diamond shape. No external JavaScript dependencies are required.

```css
.mask-diamond {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,10 L 90,50 L 50,90 L 10,50 Z'/%3E%3C/svg%3E");
}
```

---

### Apply First Half Mask Modifier with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a modifier to mask the first half of an element's content using DaisyUI CSS. This is used in conjunction with other mask classes to create partial masks. No external JavaScript dependencies are required.

```css
.mask-half-1 {
  mask-composite: intersect;
  -webkit-mask-composite: source-in;
  mask-image: linear-gradient(to right, rgba(0, 0, 0, 1) 50%, transparent 50%);
}
```

---

### Hero with Reversed Figure - HTML

Source: https://daisyui.com/components/hero

This HTML code snippet illustrates a DaisyUI hero component where the image is positioned to the left of the content. This variation is useful for A/B testing layouts or when a specific visual flow is desired. It uses standard DaisyUI and Tailwind CSS classes.

```html
<div class="hero min-h-screen bg-base-200">
  <div class="hero-content flex-col lg:flex-row-reverse">
    <img
      src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
      class="max-w-sm rounded-lg shadow-2xl"
    />
    <div>
      <h1 class="text-5xl font-bold">Box Office News!</h1>
      <p class="py-6">
        Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
        excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a
        id nisi.
      </p>
      <button class="btn btn-primary">Get Started</button>
    </div>
  </div>
</div>
```

---

### Radial Progress with Custom Color (React TSX)

Source: https://daisyui.com/components/radial-progress

This React (TSX) snippet demonstrates applying a custom color to the radial-progress component using the '--progress-color' CSS variable. It shows a 70% progress indicator with a specified color.

```tsx
function RadialProgressCustomColorTsx() {
  return (
    <div
      className="radial-progress"
      style={
        {
          "--value": 70,
          "--progress-color": "oklch(var(--p))",
        } as React.CSSProperties
      }
      role="progressbar"
    >
      70%
    </div>
  );
}
```

---

### Disabled Checkbox States (HTML)

Source: https://daisyui.com/components/checkbox

Shows how to render checkboxes in a disabled state, preventing user interaction. This is achieved by adding the `disabled` attribute to the input element.

````html
```html
<input type="checkbox" class="checkbox" disabled />
<input type="checkbox" class="checkbox" disabled checked />
````

````

--------------------------------

### DaisyUI Range Slider with Custom Color and No Fill

Source: https://daisyui.com/components/range

Demonstrates advanced customization of the range slider, allowing for a custom color and disabling the fill effect. This provides greater control over the component's appearance.

```html
<input type="range" class="range range-xs" style="--value: 60;" value="60">
<input type="range" class="range range-sm" style="--value: 60;" value="60">
<input type="range" class="range range-md" style="--value: 60;" value="60">
<input type="range" class="range range-lg" style="--value: 60;" value="60">
<input type="range" class="range range-xl" style="--value: 60;" value="60">

<input type="range" class="range range-primary" style="--value: 60;" value="60">
<input type="range" class="range range-secondary" style="--value: 60;" value="60">
<input type="range" class="range range-accent" style="--value: 60;" value="60">
<input type="range" class="range range-success" style="--value: 60;" value="60">
<input type="range" class="range range-warning" style="--value: 60;" value="60">
<input type="range" class="range range-info" style="--value: 60;" value="60">
<input type="range" class="range range-error" style="--value: 60;" value="60">
````

---

### Apply Second Half Mask Modifier with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a modifier to mask the second half of an element's content using DaisyUI CSS. This is used in conjunction with other mask classes to create partial masks. No external JavaScript dependencies are required.

```css
.mask-half-2 {
  mask-composite: intersect;
  -webkit-mask-composite: source-in;
  mask-image: linear-gradient(to left, rgba(0, 0, 0, 1) 50%, transparent 50%);
}
```

---

### Indeterminate Checkbox State (JavaScript)

Source: https://daisyui.com/components/checkbox

Provides JavaScript code to set a checkbox to an indeterminate state, which is visually distinct from checked or unchecked. This is useful for hierarchical selections.

````javascript
```javascript
document.getElementById("my-checkbox").indeterminate = true;
````

````

--------------------------------

### Apply Decagon Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a decagon mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a decagon (10-sided polygon) shape. No external JavaScript dependencies are required.

```css
.mask-decagon {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 79.3893,20.6107 C 84.035,22.7609 86.8374,27.4881 85.1203,32.1340 L 72.0557,72.9443 C 70.3386,77.5901 65.6114,80.3925 60.9657,78.6754 L 39.0343,71.3246 C 34.3886,69.6075 31.5862,64.8803 33.3033,60.2346 L 40.6533,40.0000 L 33.3033,19.7654 C 31.5862,15.1197 34.3886,10.3925 39.0343,8.6754 L 60.9657,1.3246 C 65.6114,-0.3925 70.3386,2.4099 72.0557,7.0557 L 79.3893,20.6107 Z'/%3E%3C/svg%3E");
}
````

---

### Apply Hexagon Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a hexagon mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a hexagon shape. No external JavaScript dependencies are required.

```css
.mask-hexagon {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 75,25 L 25,25 L 0,50 L 25,75 L 75,75 L 100,50 Z'/%3E%3C/svg%3E");
}
```

---

### Apply Heart Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a heart mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a heart shape. No external JavaScript dependencies are required.

```css
.mask-heart {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,25 C 38.9543,25 25,38.9543 25,50 C 25,61.0457 30.7724,71.3803 41.3168,77.7119 C 41.3168,77.7119 43.0786,78.7341 45.0000,77.0000 C 46.9214,78.7341 48.6832,77.7119 48.6832,77.7119 C 59.2276,71.3803 65,61.0457 65,50 C 65,38.9543 51.0457,25 50,25 Z'/%3E%3C/svg%3E");
}
```

---

### Disabled Textarea Component - daisyUI

Source: https://daisyui.com/components/textarea

Illustrates how to disable the Textarea component in daisyUI, making it non-interactive and visually indicating its inactive state. This is useful for form states where input is not permitted.

````html
```html <textarea class="textarea" placeholder="Bio" disabled></textarea>
````

````

--------------------------------

### DaisyUI Disabled Radio Button

Source: https://daisyui.com/components/radio

Demonstrates how to disable DaisyUI radio buttons, making them unselectable by the user. Disabled radio buttons typically appear visually distinct to indicate their state.

```html
<!-- Disabled Radio -->
<input type="radio" name="radio-disabled" class="radio radio-disabled" checked disabled />

<!-- Disabled Radio with Label -->
<label class="cursor-not-allowed label">
  <span class="label-text">Disabled Radio</span>
  <input type="radio" name="radio-disabled-label" class="radio radio-disabled" disabled />
</label>
````

---

### DaisyUI Text Input: Disabled State

Source: https://daisyui.com/components/input

Shows how to disable a DaisyUI text input field, making it non-interactive and visually distinct. Disabled inputs are useful for preventing user input in certain states.

````html
```html
<input
  type="text"
  placeholder="Disabled placeholder"
  class="input input-bordered input-disabled w-full max-w-xs"
  disabled
/>
````

````

--------------------------------

### Apply Horizontal Hexagon Mask with DaisyUI CSS

Source: https://daisyui.com/components/mask

Applies a horizontal hexagon mask to an element using DaisyUI CSS classes. This utility crops the content of the element into a horizontally oriented hexagon shape. No external JavaScript dependencies are required.

```css
.mask-hexagon-2 {
  mask-type: alpha;
  -webkit-mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Cpath d='M 50,0 L 100,25 L 100,75 L 50,100 L 0,75 L 0,25 Z'/%3E%3C/svg%3E");
}
````

=== COMPLETE CONTENT === This response contains all available snippets from this library. No additional content exists. Do not make further requests.
