# HTMX Guide - Implementation Notes

## What Was Built

A standalone guide showcasing HTMX + Alpine.js + SSE interactive components:
- **Modals** - 4 variants (basic, confirm with loading, Alpine-enhanced, form)
- **Drawers** - 3 variants (right, left nav, bottom sheet)
- **Toasts** - SSE-powered with queue, auto-dismiss, icons
- **SSE** - Dead simple single-channel implementation

**Philosophy**: Full frontend interactivity, bare minimum server.

---

## Final Structure

```
guide/
├── main.go              # ~95 lines - routes + handlers
├── sse.go               # ~35 lines - SSE endpoint
├── go.mod               # Go 1.25 + templ
├── README.md            # Full documentation
├── CONTEXT.md           # This file
└── templates/
    ├── layout.templ     # Base HTML with CDN deps
    ├── index.templ      # Demo page
    ├── modal.templ      # 4 modal variants
    ├── drawer.templ     # 3 drawer variants
    └── toast.templ      # Toast component
```

---

## Versions Used

| Package | Version |
|---------|---------|
| Go | 1.25 |
| Templ | 0.3.960 |
| HTMX | 2.0.8 |
| HTMX SSE Extension | 2.2.2 |
| Alpine.js | 3.15.3 |
| Alpine Focus | 3.15.3 |
| DaisyUI | 5 |
| Tailwind CSS | 4 (browser build) |

---

## DaisyUI 5 Breaking Changes Applied

| Old (v4) | New (v5) |
|----------|----------|
| `active` | `menu-active` |
| `form-control` + `label-text` | `fieldset` + `legend` |
| `tailwindcss.com` CDN | `@tailwindcss/browser@4` |
| `daisyui@4.x/full.css` | `daisyui@5` |

---

## Key Gotchas Discovered

### 1. SSE Toast Requires Hidden Swap Element
```html
// WITHOUT this, SSE events don't fire!
<div sse-swap="sse-toast" style="display:none"></div>
```

### 2. HTMX SSE Event Name Uses Hyphens
```html
// Correct (hyphens, not camelCase)
x-on:htmx:sse-before-message.window="handleSSE($event)"
```

### 3. Alpine x-trap Needs Focus Plugin
```html
// Load BEFORE Alpine.js
<script defer src="https://unpkg.com/@alpinejs/focus@3.15.3/dist/cdn.min.js"></script>
<script defer src="https://unpkg.com/alpinejs@3.15.3/dist/cdn.min.js"></script>
```

### 4. DaisyUI 5 CDN Setup
```html
// Must use Tailwind 4 browser build with DaisyUI 5
<link href="https://cdn.jsdelivr.net/npm/daisyui@5" rel="stylesheet" type="text/css"/>
<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
```

---

## Running

```bash
templ generate
go run .
# http://localhost:8080
```

---

## Source Components (Original Project)

| Component | Original Location |
|-----------|-------------------|
| Toast | `app/service-admin/web/components/toast/toast.templ` |
| Modal | `app/service-admin/web/components/modal/modal.templ` |
| Drawer | `app/service-admin/web/components/drawer/drawer.templ` |
| SSE Server | `app/service-admin/sse/server.go` |
| Broker | `app/service-admin/pubsub/broker.go` (not used - simplified) |
