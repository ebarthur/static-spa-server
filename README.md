# Static SPA Server

A lightweight and efficient static server designed for serving single-page applications (SPAs). Built with Go, it features:

- **Embedded static files** for seamless deployment.
- **Graceful shutdown** for zero-downtime updates.
- **SPA-friendly routing** (automatically serves `index.html` for unknown routes).
- **Minimal configuration** and easy to use.

Perfect for deploying React, Vue, Angular, or any SPA framework. Ideal for small to medium-sized projects requiring a simple, reliable static server.

## Features

- Embedded file handling for hassle-free deployments.
- Graceful shutdown to ensure no requests are dropped during updates.
- SPA routing support for modern web applications.
- Lightweight and fast, built with Go.

## Use Cases

- Deploying static web applications.
- Serving SPAs with client-side routing.
- Simple projects needing a no-frills static server.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Makefile

Run build command with tests:

```bash
make all
```

Build the application:

```bash
make build
```

Run the application:

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up the binary from the last build:

```bash
make clean
```
