# PostgresUI Design System

## Overview

This design system uses Tailwind CSS v4 + DaisyUI with HSL-based theming inspired by Astro Starlight. The goal is to provide a consistent, accessible, and themeable UI across the entire application.

## Design Principles

1. **Lean on standards**: Use Tailwind's default spacing and sizing scales
2. **Semantic colors**: Always use DaisyUI semantic tokens instead of raw color values
3. **Minimal and clean**: Avoid custom values unless absolutely necessary
4. **Accessible**: WCAG AA compliance minimum, AAA preferred

## Color System

### Semantic Tokens (DaisyUI)

**Always use these instead of raw colors:**

| Token                | Usage                          | Example                           |
|----------------------|--------------------------------|-----------------------------------|
| `bg-base-100`        | Main background                | `<div class="bg-base-100">`       |
| `bg-base-200`        | Secondary background           | Cards, panels                     |
| `bg-base-300`        | Borders, dividers              | `<div class="border-base-300">`   |
| `text-base-content`  | Primary text                   | Body text                         |
| `text-base-content/70` | Muted text                   | Secondary text, descriptions      |
| `btn-primary`        | Primary actions                | Save, Submit buttons              |
| `btn-secondary`      | Secondary actions              | Cancel, Back                      |
| `btn-ghost`          | Tertiary actions               | Delete, Clear                     |
| `btn-accent`         | Accent actions                 | Special features                  |

### Additional Semantic Colors

| Token           | Usage                     | Example                           |
|-----------------|---------------------------|-----------------------------------|
| `alert-info`    | Informational messages    | Tips, notices                     |
| `alert-success` | Success messages          | Completed actions                 |
| `alert-warning` | Warning messages          | Cautions                          |
| `alert-error`   | Error messages            | Validation errors                 |

### Custom HSL Scale

Behind the scenes, themes generate these CSS variables:
- `--color-accent-{50-950}`: Primary brand colors
- `--color-gray-{50-950}`: Neutral grays

**You rarely need to use these directly.** Use DaisyUI tokens instead. However, if you need fine-grained control:

```css
/* Accent scale */
var(--color-accent-50)   /* Lightest */
var(--color-accent-500)  /* Base */
var(--color-accent-950)  /* Darkest */

/* Gray scale */
var(--color-gray-50)     /* Lightest */
var(--color-gray-500)    /* Base */
var(--color-gray-950)    /* Darkest */
```

## Spacing

Use Tailwind's default scale consistently:

| Class   | Size  | Usage                           |
|---------|-------|---------------------------------|
| `gap-1` | 4px   | Micro spacing (icon + text)     |
| `gap-2` | 8px   | Small spacing (button groups)   |
| `gap-4` | 16px  | Medium spacing (form fields)    |
| `gap-6` | 24px  | Large spacing (sections)        |
| `gap-8` | 32px  | XL spacing (major sections)     |

Same applies to padding (`p-*`), margin (`m-*`), width/height (`w-*`, `h-*`).

## Typography

### Font Families

- **Sans**: `Inter` (UI text, body, headings)
- **Mono**: `Fira Code` (code blocks, SQL queries)
- **Decorative**: `Gideon Roman` (optional for hero headings)

### Font Sizes

| Class      | Size | Usage                           |
|------------|------|---------------------------------|
| `text-xs`  | 12px | Micro labels, badges            |
| `text-sm`  | 14px | Small labels, captions          |
| `text-base`| 16px | Body text (default)             |
| `text-lg`  | 18px | Subheadings                     |
| `text-xl`  | 20px | Section headings                |
| `text-2xl` | 24px | Page titles                     |
| `text-3xl` | 30px | Hero headings                   |

### Font Weights

| Class           | Weight | Usage                     |
|-----------------|--------|---------------------------|
| `font-normal`   | 400    | Body text                 |
| `font-medium`   | 500    | Emphasized text           |
| `font-semibold` | 600    | Subheadings               |
| `font-bold`     | 700    | Headings, titles          |

## Component Patterns

### Buttons

```html
<!-- Primary action -->
<button class="btn btn-primary">Save</button>

<!-- Secondary action -->
<button class="btn btn-secondary">Cancel</button>

<!-- Tertiary/Ghost -->
<button class="btn btn-ghost">Delete</button>

<!-- Sizes -->
<button class="btn btn-xs">Extra Small</button>
<button class="btn btn-sm">Small</button>
<button class="btn">Default</button>
<button class="btn btn-lg">Large</button>

<!-- Loading state -->
<button class="btn btn-primary loading">Processing...</button>

<!-- Disabled -->
<button class="btn btn-primary" disabled>Disabled</button>
```

### Cards

```html
<div class="card bg-base-200 shadow-sm">
    <div class="card-body p-4">
        <h3 class="card-title text-lg">Card Title</h3>
        <p class="text-base-content/70">Card content goes here.</p>
        <div class="card-actions justify-end mt-4">
            <button class="btn btn-primary btn-sm">Action</button>
        </div>
    </div>
</div>

<!-- Compact card -->
<div class="card bg-base-200 border border-base-300">
    <div class="card-body p-3">
        <h4 class="text-base font-medium">Compact Card</h4>
    </div>
</div>
```

### Forms

```html
<div class="form-control">
    <label class="label">
        <span class="label-text">Field Label</span>
        <span class="label-text-alt">Optional</span>
    </label>
    <input 
        type="text" 
        class="input input-bordered" 
        placeholder="Enter text..." />
    <label class="label">
        <span class="label-text-alt text-error">Error message</span>
    </label>
</div>

<!-- Select -->
<div class="form-control">
    <label class="label">
        <span class="label-text">Dropdown</span>
    </label>
    <select class="select select-bordered">
        <option disabled selected>Choose one</option>
        <option>Option 1</option>
        <option>Option 2</option>
    </select>
</div>

<!-- Textarea -->
<div class="form-control">
    <label class="label">
        <span class="label-text">Description</span>
    </label>
    <textarea 
        class="textarea textarea-bordered" 
        placeholder="Enter description..."
        rows="4"></textarea>
</div>

<!-- Checkbox -->
<div class="form-control">
    <label class="label cursor-pointer">
        <span class="label-text">Enable feature</span>
        <input type="checkbox" class="checkbox" />
    </label>
</div>

<!-- Toggle -->
<div class="form-control">
    <label class="label cursor-pointer">
        <span class="label-text">Active</span>
        <input type="checkbox" class="toggle toggle-primary" />
    </label>
</div>
```

### Tables

```html
<table class="table table-zebra">
    <thead class="bg-base-200">
        <tr>
            <th class="text-sm font-semibold">Column 1</th>
            <th class="text-sm font-semibold">Column 2</th>
            <th class="text-sm font-semibold">Actions</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Data 1</td>
            <td>Data 2</td>
            <td>
                <button class="btn btn-ghost btn-xs">Edit</button>
            </td>
        </tr>
        <tr>
            <td>Data 3</td>
            <td>Data 4</td>
            <td>
                <button class="btn btn-ghost btn-xs">Edit</button>
            </td>
        </tr>
    </tbody>
</table>

<!-- Compact table -->
<table class="table table-sm">
    <!-- ... -->
</table>
```

### Alerts

```html
<!-- Info -->
<div class="alert alert-info">
    <iconify-icon icon="mdi:information" class="text-lg"></iconify-icon>
    <span>This is an informational message.</span>
</div>

<!-- Success -->
<div class="alert alert-success">
    <iconify-icon icon="mdi:check-circle" class="text-lg"></iconify-icon>
    <span>Action completed successfully!</span>
</div>

<!-- Warning -->
<div class="alert alert-warning">
    <iconify-icon icon="mdi:alert" class="text-lg"></iconify-icon>
    <span>Please review this warning.</span>
</div>

<!-- Error -->
<div class="alert alert-error">
    <iconify-icon icon="mdi:alert-circle" class="text-lg"></iconify-icon>
    <span>An error occurred.</span>
</div>
```

### Badges

```html
<span class="badge">Default</span>
<span class="badge badge-primary">Primary</span>
<span class="badge badge-secondary">Secondary</span>
<span class="badge badge-accent">Accent</span>
<span class="badge badge-ghost">Ghost</span>

<!-- Sizes -->
<span class="badge badge-xs">XS</span>
<span class="badge badge-sm">SM</span>
<span class="badge badge-md">MD</span>
<span class="badge badge-lg">LG</span>
```

### Modals

```html
<!-- Trigger button -->
<button class="btn" onclick="my_modal.showModal()">Open Modal</button>

<!-- Modal -->
<dialog id="my_modal" class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg mb-4">Modal Title</h3>
        <p class="text-base-content/70">Modal content goes here.</p>
        <div class="modal-action">
            <form method="dialog">
                <button class="btn btn-ghost">Close</button>
                <button class="btn btn-primary">Confirm</button>
            </form>
        </div>
    </div>
    <form method="dialog" class="modal-backdrop">
        <button>close</button>
    </form>
</dialog>
```

### Loading States

```html
<!-- Spinner -->
<span class="loading loading-spinner loading-xs"></span>
<span class="loading loading-spinner loading-sm"></span>
<span class="loading loading-spinner loading-md"></span>
<span class="loading loading-spinner loading-lg"></span>

<!-- Dots -->
<span class="loading loading-dots loading-sm"></span>

<!-- In button -->
<button class="btn btn-primary">
    <span class="loading loading-spinner loading-sm"></span>
    Processing...
</button>
```

### Icons

Use Iconify with the `iconify-icon` web component:

```html
<!-- Basic icon -->
<iconify-icon icon="mdi:home" class="text-xl"></iconify-icon>

<!-- With text -->
<button class="btn btn-primary gap-2">
    <iconify-icon icon="mdi:plus"></iconify-icon>
    Add New
</button>

<!-- Common icons -->
mdi:pencil          <!-- Edit -->
mdi:delete          <!-- Delete -->
mdi:refresh         <!-- Refresh -->
mdi:check           <!-- Success -->
mdi:close           <!-- Close -->
mdi:information     <!-- Info -->
mdi:alert           <!-- Warning -->
mdi:alert-circle    <!-- Error -->
```

## Layout Patterns

### Container

```html
<!-- Full width container -->
<div class="container mx-auto px-4">
    <!-- Content -->
</div>

<!-- Max width container -->
<div class="max-w-4xl mx-auto px-4">
    <!-- Content -->
</div>
```

### Grid Layouts

```html
<!-- Responsive grid -->
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <div>Column 1</div>
    <div>Column 2</div>
    <div>Column 3</div>
</div>

<!-- Sidebar + Main -->
<div class="grid grid-cols-1 lg:grid-cols-[250px_1fr] gap-4">
    <aside>Sidebar</aside>
    <main>Main content</main>
</div>
```

### Flex Layouts

```html
<!-- Horizontal layout -->
<div class="flex items-center gap-4">
    <div>Item 1</div>
    <div>Item 2</div>
</div>

<!-- Space between -->
<div class="flex items-center justify-between">
    <div>Left</div>
    <div>Right</div>
</div>

<!-- Vertical stack -->
<div class="flex flex-col gap-4">
    <div>Item 1</div>
    <div>Item 2</div>
</div>
```

## Theme Configuration

Admins can configure themes at `/admin/theme`:

### Theme Parameters

- **Accent Hue**: 0-360 degrees (primary brand color)
  - Examples: 250 (Indigo), 150 (Emerald), 350 (Rose)
- **Gray Hue**: 0-360 degrees (neutral tones)
  - Examples: 240 (Cool), 0 (Warm), 180 (Cyan-tinted)
- **Contrast Level**: Low, Normal, High
  - **Low**: 50% saturation (softer, less vibrant)
  - **Normal**: 70% saturation (balanced)
  - **High**: 90% saturation (vibrant, WCAG AAA)

### Configuration Priority

The theme system uses a cascading configuration approach:

1. **Database** (Highest Priority)
   - UI-configured via `/admin/theme`
   - Stored in SQLite `themes` table
   - Allows runtime theme changes
   
2. **Config File** (Medium Priority)
   - `config/theme.yml`
   - For IaC/GitOps deployments
   - Version-controlled theme settings
   
3. **Defaults** (Fallback)
   - Hardcoded in application
   - Indigo theme (hue: 250, gray: 240, contrast: normal)

### Theme Presets

Built-in presets available in the theme editor:

| Preset   | Accent Hue | Gray Hue | Description           |
|----------|------------|----------|-----------------------|
| Indigo   | 250        | 240      | Default, professional |
| Emerald  | 150        | 150      | Fresh, organic        |
| Rose     | 350        | 340      | Warm, friendly        |
| Custom   | Variable   | Variable | User-defined          |

### Exporting Themes

Themes can be exported to YAML for:
- Version control
- Deployment via CI/CD
- Sharing across environments

Example `config/theme.yml`:

```yaml
theme:
  name: "indigo"
  accent_hue: 250
  gray_hue: 240
  contrast_level: normal
```

## Migration Guide

### Migrating from Raw Colors

#### ❌ Before (Inconsistent)
```html
<div class="bg-gray-100 text-red-800 p-3">
    <button class="bg-blue-600 text-white px-4 py-2 rounded">
        Click me
    </button>
</div>
```

#### ✅ After (Semantic)
```html
<div class="bg-base-200 text-base-content p-4">
    <button class="btn btn-primary">
        Click me
    </button>
</div>
```

### Common Replacements

| Old (Raw)                  | New (Semantic)              |
|----------------------------|-----------------------------|
| `bg-gray-50`               | `bg-base-100`               |
| `bg-gray-100`              | `bg-base-200`               |
| `bg-gray-200`              | `bg-base-300`               |
| `text-gray-900`            | `text-base-content`         |
| `text-gray-600`            | `text-base-content/70`      |
| `bg-blue-600`              | `btn-primary` or `bg-primary` |
| `border-gray-300`          | `border-base-300`           |
| `text-red-600`             | `text-error`                |
| `text-green-600`           | `text-success`              |

### Migration Checklist

- [ ] Replace all `bg-gray-*` with `bg-base-*`
- [ ] Replace all `text-gray-*` with `text-base-content` variants
- [ ] Replace all `border-gray-*` with `border-base-*`
- [ ] Replace raw color buttons with `btn-*` classes
- [ ] Use DaisyUI form components (`input-bordered`, `select-bordered`)
- [ ] Use semantic alert classes (`alert-info`, `alert-success`, etc.)
- [ ] Test theme switching to ensure colors adapt correctly

## Accessibility Guidelines

### Contrast Requirements

- **WCAG AA**: 4.5:1 for normal text, 3:1 for large text
- **WCAG AAA**: 7:1 for normal text, 4.5:1 for large text

Use the "High" contrast level in theme settings for AAA compliance.

### Focus States

Always ensure keyboard focus is visible:

```html
<!-- Good: Default focus ring -->
<button class="btn btn-primary">Click me</button>

<!-- Custom focus ring -->
<a href="#" class="focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-2">
    Link
</a>
```

### ARIA Labels

Provide descriptive labels for interactive elements:

```html
<button class="btn btn-ghost" aria-label="Delete connection">
    <iconify-icon icon="mdi:delete"></iconify-icon>
</button>

<input 
    type="text" 
    class="input input-bordered"
    aria-label="Database name"
    aria-describedby="db-name-hint" />
<span id="db-name-hint" class="text-sm text-base-content/70">
    Enter a unique name for this connection
</span>
```

### Semantic HTML

Use proper HTML elements:

```html
<!-- Good -->
<nav>
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/admin">Admin</a></li>
    </ul>
</nav>

<!-- Avoid -->
<div class="nav">
    <div class="link" onclick="...">Home</div>
</div>
```

## Best Practices

1. **Consistency over customization**: Use existing patterns before creating new ones
2. **Mobile-first**: Design for mobile, enhance for desktop
3. **Progressive enhancement**: Start with semantic HTML, enhance with CSS/JS
4. **Performance**: Minimize custom CSS, leverage Tailwind's utility classes
5. **Accessibility**: Test with keyboard navigation and screen readers
6. **Dark mode ready**: Use semantic tokens so dark mode works automatically (future)

## Resources

- [Tailwind CSS Documentation](https://tailwindcss.com/docs)
- [DaisyUI Components](https://daisyui.com/components/)
- [Astro Starlight Theme System](https://starlight.astro.build/guides/css-and-tailwind/)
- [WCAG Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)
- [Iconify Icon Sets](https://icon-sets.iconify.design/)

## Questions?

For theme configuration issues, see `/admin/theme` in the application.

For design system questions, review this document or consult the team.
