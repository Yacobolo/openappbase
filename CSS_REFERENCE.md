# Context-Based CSS Reference Guide

Quick reference for working with the new context-based CSS architecture.

## Philosophy: Context over Classes

**95% Clean HTML** - Element location dictates style  
**5% Exceptions** - Classes only for rule-breaking elements

## CSS File Structure

```
web/resources/styles/
├── variables.css      # Design tokens (colors, spacing, typography)
├── reset.css         # Minimal modern reset
├── base.css          # Global tag styles (button, input, table, etc)
├── layout.css        # Layout contexts (header, aside, main)
├── components.css    # Component contexts (cards, modals, tabs)
├── utilities.css     # Exception classes only (.active, .danger, etc)
└── styles.css        # Entry point (imports all above)
```

## Design Tokens

### Colors

```css
/* Semantic Colors */
--color-bg              /* Main background (white) */
--color-bg-subtle       /* Subtle background (gray-50) */
--color-bg-muted        /* Muted background (gray-100) */
--color-border          /* Main border (gray-300) */
--color-border-subtle   /* Subtle border (gray-200) */
--color-text            /* Main text (gray-900) */
--color-text-muted      /* Muted text (gray-500) */

/* Functional Colors */
--color-primary         /* Primary actions (blue-700) */
--color-primary-hover   /* Primary hover (blue-800) */
--color-primary-light   /* Primary background (blue-50) */
--color-danger          /* Destructive actions (red-500) */
--color-success         /* Success states (teal-400) */
```

### Spacing

```css
--space-xs    /* 4px */
--space-sm    /* 8px */
--space-md    /* 16px */
--space-lg    /* 24px */
--space-xl    /* 32px */
--space-2xl   /* 48px */
```

### Typography

```css
--text-xs     /* 0.75rem */
--text-sm     /* 0.875rem */
--text-base   /* 1rem */
--text-lg     /* 1.125rem */
--text-xl     /* 1.25rem */
--text-2xl    /* 1.5rem */
--text-3xl    /* 1.875rem */

--font-weight-normal    /* 400 */
--font-weight-medium    /* 500 */
--font-weight-semibold  /* 600 */
--font-weight-bold      /* 700 */
```

## Common Patterns

### Buttons

```html
<!-- Default button (styled globally) -->
<button>Action</button>

<!-- Primary action (exception class) -->
<button class="primary">Save</button>

<!-- Destructive action (exception class) -->
<button class="danger">Delete</button>

<!-- Ghost button (exception class) -->
<button class="ghost">Cancel</button>
```

### Forms

```html
<!-- Forms are styled globally -->
<form>
  <label>Email</label>
  <input type="email" placeholder="your@email.com"/>
  
  <label>Message</label>
  <textarea placeholder="Your message..."></textarea>
  
  <button type="submit" class="primary">Submit</button>
</form>
```

### Tables

```html
<!-- Tables are styled globally -->
<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Email</th>
      <th>Status</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>John Doe</td>
      <td>john@example.com</td>
      <td><span class="badge badge-success">Active</span></td>
    </tr>
  </tbody>
</table>
```

### Navigation

```html
<!-- Sidebar navigation (styled by context) -->
<aside>
  <header>
    <h1>App Name</h1>
  </header>
  
  <nav>
    <details open>
      <summary>Section Name</summary>
      <ul>
        <li><a href="/" class="active">Home</a></li>
        <li><a href="/about">About</a></li>
      </ul>
    </details>
  </nav>
  
  <footer>
    <span class="badge badge-success">Connected</span>
  </footer>
</aside>
```

### Collapsible Sections

```html
<!-- Use native <details> instead of DaisyUI collapse -->
<details open>
  <summary>
    Section Title
    <span class="count">5</span>
  </summary>
  <div>
    Content goes here
  </div>
</details>
```

### Badges

```html
<!-- Status badges -->
<span class="badge">Default</span>
<span class="badge badge-primary">Primary</span>
<span class="badge badge-success">Active</span>
<span class="badge badge-danger">Error</span>
<span class="badge badge-outline">Outline</span>

<!-- Small badge -->
<span class="badge badge-sm">Small</span>

<!-- Count badge -->
<span class="count">42</span>
```

### Cards

```html
<div class="card">
  <header>
    <h2>Card Title</h2>
  </header>
  
  <p>Card content goes here.</p>
  
  <footer>
    <button>Action</button>
  </footer>
</div>
```

### Tabs

```html
<div class="tabs">
  <button class="active">Tab 1</button>
  <button>Tab 2</button>
  <button>Tab 3</button>
</div>
```

### Breadcrumbs

```html
<div class="breadcrumb">
  <a href="/">Home</a>
  <a href="/section">Section</a>
  <span>Current Page</span>
</div>
```

### Toast Notifications

```html
<!-- Add to toast container -->
<div id="toast-container" class="toast">
  <div>
    <span>Your message here</span>
    <button>×</button>
  </div>
</div>
```

### Empty States

```html
<div class="empty-state">
  <svg>...</svg>
  <h3>No Data Found</h3>
  <p>Get started by adding your first item.</p>
  <button class="primary">Add Item</button>
</div>
```

### Loading Indicators

```html
<!-- Inline loading spinner -->
<div class="loading"></div>

<!-- Full overlay -->
<div class="loading-overlay">
  <div class="loading"></div>
</div>
```

## Exception Classes (Use Sparingly!)

### State Classes

```css
.active     /* Active navigation links, tabs */
```

### Color Variants

```css
.primary    /* Primary actions */
.danger     /* Destructive actions */
.success    /* Success actions */
.ghost      /* Transparent buttons */
```

### Badge Variants

```css
.badge                /* Base badge */
.badge-outline        /* Outline style */
.badge-primary        /* Primary color */
.badge-success        /* Success color */
.badge-danger         /* Danger color */
.badge-sm             /* Small size */
```

### Size Variants

```css
.count      /* Small count badges */
.text-sm    /* Small text */
.text-lg    /* Large text */
```

### Layout Helpers

```css
.flex           /* Display flex */
.flex-col       /* Flex column */
.items-center   /* Align items center */
.justify-center /* Justify content center */
.gap-sm         /* Small gap */
.w-full         /* Full width */
```

### Display Utilities

```css
.hidden          /* Hide element */
.sr-only         /* Screen reader only */
.mobile-only     /* Show on mobile only */
.desktop-only    /* Hide on mobile */
.text-center     /* Center text */
.truncate        /* Truncate with ellipsis */
```

## Layout Structure

### Page Layout

```html
<body>
  <!-- Header spans full width -->
  <header>
    <div><!-- Left section --></div>
    <div><!-- Center section --></div>
    <div><!-- Right section --></div>
  </header>
  
  <!-- Mobile drawer toggle -->
  <input id="my-drawer-2" type="checkbox"/>
  <div class="drawer-overlay"></div>
  
  <!-- Sidebar (auto-positioned) -->
  <aside>...</aside>
  
  <!-- Main content (auto-positioned) -->
  <main>...</main>
</body>
```

### Responsive Breakpoints

```css
/* Mobile First */
@media (min-width: 768px)  { /* Tablet */ }
@media (min-width: 1024px) { /* Desktop */ }
```

## Migration Guide

### DaisyUI → Context-Based

| DaisyUI | Context-Based |
|---------|---------------|
| `<div class="collapse">` | `<details>` |
| `<div class="badge">` | `<span class="badge">` |
| `<button class="btn btn-primary">` | `<button class="primary">` |
| `<button class="btn btn-ghost">` | `<button class="ghost">` |
| `<input class="input input-bordered">` | `<input type="text">` |
| `<div class="card bg-base-100">` | `<div class="card">` |
| `<div class="tabs tabs-bordered">` | `<div class="tabs">` |
| `<ul class="menu">` | `<ul>` (styled by context) |

### Utility Classes → Context

| Utility | Solution |
|---------|----------|
| `class="flex items-center"` | Context styling (header, aside) |
| `class="p-4 mb-2"` | Context styling or semantic HTML |
| `class="text-sm text-gray-500"` | Global tag styles |
| `class="border border-gray-200"` | Global tag styles |
| `class="bg-base-100"` | Global defaults |

## Best Practices

### DO ✅

- Use semantic HTML elements (`<button>`, `<nav>`, `<aside>`, `<details>`)
- Let context drive styling (location matters)
- Use exception classes only when truly needed
- Reference design tokens in custom CSS
- Keep HTML clean and readable

### DON'T ❌

- Add utility classes for spacing/sizing
- Create new exception classes without discussion
- Override context styles with inline styles
- Use generic `<div>` when semantic element exists
- Add classes to every element

## Examples

### Before (Tailwind + DaisyUI)

```html
<div class="w-80 bg-base-200 p-4 flex flex-col shrink-0">
  <header class="mb-4 text-center border-b border-base-100 pb-4">
    <h1 class="text-2xl font-bold text-primary">Explorer</h1>
  </header>
  <div class="mb-4">
    <input type="text" class="input input-bordered w-full"/>
  </div>
  <nav class="flex-grow overflow-y-auto">
    <div class="collapse collapse-arrow bg-base-100 rounded-lg">
      <input type="checkbox" checked/>
      <div class="collapse-title font-medium">Schema</div>
    </div>
  </nav>
</div>
```

### After (Context-Based)

```html
<aside>
  <header>
    <h1>Explorer</h1>
  </header>
  <div>
    <input type="text"/>
  </div>
  <nav>
    <details open>
      <summary>Schema</summary>
    </details>
  </nav>
</aside>
```

## Development Workflow

### Build Commands

```bash
# Build CSS
task build:styles

# Build templ files
task build:templ

# Build everything
task build

# Dev server with hot reload
task live
```

### File Watching

The `task live` command watches for changes in:
- `*.templ` files
- `web/resources/styles/*.css` files

Changes automatically trigger rebuilds.

## Troubleshooting

### CSS Not Updating

```bash
# Rebuild CSS manually
task build:styles

# Check for syntax errors in CSS files
```

### Styles Not Applied

1. Check if element is in correct context (header, aside, main)
2. Verify no inline styles overriding
3. Check browser DevTools for applied styles
4. Ensure CSS build succeeded

### Layout Issues

1. Verify body grid structure exists
2. Check header/aside/main placement
3. Test responsive breakpoints
4. Verify drawer checkbox for mobile

## Further Reading

- [IMPLEMENTATION_SUMMARY.md](./IMPLEMENTATION_SUMMARY.md) - Full migration details
- [CONTEXT_CSS_TRANSITION_PLAN.md](./CONTEXT_CSS_TRANSITION_PLAN.md) - Original plan
- [AGENTS.md](./AGENTS.md) - Development guidelines
