# Slnix-task



# File Management System

This project is a file management API built in Golang, supporting file upload, retrieval, and user management with role-based access.

## Features

- User authentication (JWT-based)
- Organization-based folder and file management
- IP-based access control, one-time links, rate limiting
- Local file storage

cmd/main.go: The main application entry, setting up the router, middlewares, and starting the HTTP server.
config/config.go: Manages application configuration, loading values from environment variables.
controllers/: Controllers that define the route handlers for different resources (e.g., files, folders, orgs).
middlewares/: Various middlewares for authentication, role checks, IP restrictions, and rate limiting.
models/: Contains models for User, File, Folder, and Organization, each representing a database entity.
services/: Service layer, handling core business logic and abstracting operations on models.
storage/local_storage.go: Handles interactions with the file system for local storage, including reading/writing files.
utils/: Utility functions for token generation, logging, rate limiting, and UUID generation.
Dockerfile: Defines how to build a Docker image for the API.
nginx.conf: Configuration for NGINX reverse proxy, handling requests to subdomains and routing them to the server

