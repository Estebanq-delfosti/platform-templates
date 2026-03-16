# Platform Templates

Catálogo de templates de Backstage para scaffolding de nuevos servicios y librerías.

## Templates disponibles

| Template | Descripción |
|---|---|
| [`go-backend`](./go-backend/) | Servicio HTTP en Go — chi, arquitectura hexagonal, Docker, K8s |
| [`go-lambda`](./go-lambda/) | Función AWS Lambda en Go — API GW HTTP v2, Terraform, arquitectura hexagonal |
| [`nodejs-backend`](./nodejs-backend/) | Servicio HTTP en Node.js |
| [`nodejs-library`](./nodejs-library/) | Librería npm en Node.js |
| [`react-frontend`](./react-frontend/) | Aplicación frontend en React |
| [`angular-frontend`](./angular-frontend/) | Aplicación frontend en Angular |
| [`rust-backend`](./rust-backend/) | Servicio HTTP en Rust |
| [`rust-library`](./rust-library/) | Librería en Rust |

## Estructura

Cada template contiene:

```
<template>/
├── template.yaml   # Definición del scaffolder de Backstage
└── skeleton/       # Archivos que se copian al nuevo repositorio
```

El índice central está en [`location.yaml`](./location.yaml).
