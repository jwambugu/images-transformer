# Images Transformer

This project is part of the lessons from [Gophercises](https://gophercises.com/) where we learn how to build how to
build an application for transforming images to different shapes
using [Primitive](https://github.com/fogleman/primitive).

## Run Locally

Clone the project

```bash
  git@github.com:jwambugu/images-transformer.git
```

Go to the project directory

```bash
  cd images-transformer
```

Copy the configuration file and update your configuration

```bash
  cp .keys.sample.json .keys.json
```

Start the server

```go
   go run cmd/api/*
```

## API Reference

#### Get all available modes

```http
  GET /v1/modes
```

#### Get the number of shapes

```http
  GET /v1/modes/no-of-shapes
```

#### Upload image to transform

```http
  POST /api/images
```

| Parameter | Type       | Description                                                |
| :-------- | :-------   | :----------------------------------------------------------|
| `photos`  | `file`     | **Required**. The image to upload.                         |
| `mode`    | `int`      | **Required**. The primitive mode to use. Expects **0-8**   |
| `shapes`  | `int`      | **Required**. The number of shapes to render.              |