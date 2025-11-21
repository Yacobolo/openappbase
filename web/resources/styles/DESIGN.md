# PostgresUI Design System

## Philosophy: Context over Classes

- **95% Clean HTML**: Location dictates style, not classes
- **Global Defaults**: Style naked tags first (`button`, `table`, `input`)
- **Context**: DOM structure drives design (`aside button`, `.card table`)
- **5% Exceptions**: Classes only when breaking parent rules

## File Structure

```text
web/resources/styles/
├── variables.css    # Custom properties
├── reset.css        # Minimal reset
├── base.css         # Global tag defaults
├── layout.css       # header, aside, main contexts
├── components.css   # Specific contexts
└── utilities.css    # Exception classes only
```

## Variables

```css
--color-{bg,text,primary,danger}
--space-{sm,md,lg}
--header-height, --sidebar-width
```

## HTML Style

Clean semantic markup:
```html
<button>Save</button>              <!-- no classes -->
<button class="danger">Delete</button>  <!-- exception -->
```

## Layout

Grid-based: `body { display: grid }` with header spanning full width, sidebar + main columns

## Exception Classes

- `.danger` - Destructive
- `.primary` - Primary action
- `.active` - Active state
- `.count` - Badge

## States

Use pseudo-classes: `:hover`, `:focus-visible`, `:active`, `:disabled`, `details[open]`

## Icons

`<iconify-icon icon="mdi:name">` - context handles spacing

## Datastar

Forms: `data-bind:signal.field`, actions: `data-on:click="@post(url)"`
Templ: `datastar.PostSSE("/path")`

## Build

`task build:styles` - Compiles styles.css → index.css
