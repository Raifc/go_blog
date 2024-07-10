# Go Blog

A simple blogging application built with Golang, PostgreSQL, and GORM.

## Table of Contents

- Overview
- Features
- Prerequisites
- Installation
- Running the Application
- Running Tests
- API Endpoints
- Example Requests

## Overview

This project is a blogging platform where users can create and manage blog posts and comments. It uses Golang for the backend, PostgreSQL for the database, and GORM for ORM.

## Features

- Create, read, update, and delete blog posts
- Create, read, update, and delete comments on blog posts

## Prerequisites

- Docker
- Docker Compose

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-blog.git
```

2. Navigate to the project directory:
```bash
cd go-blog
```

3. Copy the example environment file and edit it with your configurations:
```bash
cp .env.example .env
```

## Running the Application

1. Build and start the application using Docker Compose:
```bash
docker-compose up --build
```
2. The application should now be running on `http://localhost:8080`.

## API Endpoints

### Posts

- `POST /posts` - Create a new post
- `GET /posts` - Get all posts
- `GET /posts/{id}` - Get a post by ID
- `PUT /posts/{id}` - Update a post by ID
- `DELETE /posts/{id}` - Delete a post by ID

### Comments

- `POST /posts/{postID}/comments` - Create a new comment for a post
- `GET /posts/{postID}/comments` - Get all comments for a post
- `GET /comments/{id}` - Get a comment by ID
- `PUT /comments/{id}` - Update a comment by ID
- `DELETE /comments/{id}` - Delete a comment by ID
