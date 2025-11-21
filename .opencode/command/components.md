---
description: Update component showcase with latest CSS changes
agent: build
---

# Update Component Showcase

The component showcase is located at `docs/components.html`. Review the latest CSS changes and update the showcase to reflect any new components or styling changes.

## Current CSS Files
The following CSS files define the components:
- `web/resources/styles/components.css` - Component-specific contexts and layouts
- `web/resources/styles/variables.css` - Design tokens and CSS custom properties
- `web/resources/styles/base.css` - Global tag defaults

## Task
Review recent changes to the CSS files and:
1. Identify any new components or component modifications
2. Update `docs/components.html` to include examples for any new components
3. Ensure all component examples showcase the current design accurately
4. Test that the showcase renders correctly by opening it in a browser

## Component Categories in Showcase
- Typography
- Buttons  
- Forms
- Cards
- Modals
- Tables
- Badges
- Breadcrumb
- Loading
- Pagination
- Tabs
- Toolbar
- Connections
- Empty State
- Error Messages

## Notes
- Keep examples simple and focused on demonstrating the component
- Use semantic HTML following the "95% clean HTML" philosophy
- Showcase all states: default, hover, active, disabled, etc.
- The showcase uses relative path to import CSS: `../web/resources/static/index.css`
