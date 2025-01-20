## Table of Contents

- [Overview](#overview)
- [Development Setup](#development-setup)
- [Development Mode](#development-mode)
- [Production Mode](#production-mode)
- [Access Information](#access-information)

## Overview

This is a prototype of an application for managing a company's human resources.

## Development Setup

### Create `.env` file for Backend

```sh
cp backend/app/.env.local.example backend/app/.env.local
```

### Create `.env` file for Frontend

```sh
cp frontend/.env.local.example frontend/.env.local
```

## Development Mode

This mode is used for local development.

### Build

```sh
docker compose build
```

### Start

#### Initialize and start with sample data

```sh
SEEDING=true docker compose up -d
```

#### Normal start

```sh
docker compose up -d
```

### Stop

```sh
docker compose down
```

## Production Mode

This mode is used for production environments. It can also be used to verify operations in an environment equivalent to production during development.

### Build

```sh
docker compose -f docker-compose.yml build
```

### Start

```sh
docker compose -f docker-compose.yml up -d
```

### Stop

```sh
docker compose down
```

## Access Information

In the local environment, you can access it at http://glissando.localhost.
