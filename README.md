# URL Shortener

## Overview

A URL shortener web service built for learning purposes. The initial prototype was developed in a single afternoon. Built with Go using the standard `net/http` package, this service allows users to interact via HTTP requests to create, retrieve, and delete shortened URLs.

## API Endpoints

| Method | Path    | Description                                         |
| ------ | ------- | --------------------------------------------------- |
| GET    | /:id    | Redirects the user to the original URL (if exists). |
| POST   | /create | Creates a new shortened URL, returns the id         |
| DELETE | /:id    | Deletes a previously set URL                        |

## Architecture

The application folows a layered architecture:

- **Handlers**: HTTP concerns (routing, status codes, JSON responses)
- **Service**: Business rules (ID Generation, collision handling, error treatment)
- **Repository**: Persistence abstraction with two implementations (SQL and InMemory)

Handlers depend only on the service interface, and the service depends only on repository interfaces.

## Features

- [x] Independent layers (handlers, service, and repository)
- [x] Redirect to the original address
- [x] Shorten URLs
- [x] Generate random IDs
- [x] SQL and in-memory repositories
- [x] Detect ID collisions
- [x] Allow deletion of URLs
- [x] Error handling
- [x] Documentation
- [ ] Docker
- [ ] Tests
- [ ] (Optional) Add expiration date for URLs
