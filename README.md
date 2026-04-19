# HH Service Property

A simple CRUD microservice for managing map nodes and their type styles using Go and MongoDB.

## Setup

1. Install Go 1.19+ and MongoDB.

2. Clone the repo.

3. Copy .env.example to .env and update MONGO_URI and DB_NAME.

4. Run `go mod tidy`

5. Run `go build ./cmd/app`

6. Run `./app` or `go run ./cmd/app`

## API Endpoints

### Node Type Styles

- POST /api/node-type-styles : Create

- GET /api/node-type-styles : Get all styles

- GET /api/node-type-styles?type=<type> : Get style by type (query parameter, treated as primary key)

- GET /api/node-type-styles/:id : Get by ID

- PUT /api/node-type-styles/:id : Update

- DELETE /api/node-type-styles/:id : Delete

### Map Nodes

- POST /api/map-nodes : Create

- GET /api/map-nodes : Get all nodes

- GET /api/map-nodes?type=<type> : Get nodes by type (returns list of nodes matching the type)

- GET /api/map-nodes/:id : Get by ID

- PUT /api/map-nodes/:id : Update

- DELETE /api/map-nodes/:id : Delete

## Database Configuration

**Database Name:** `househunt`

**Collections:**
- `node_type_styles` - Stores styling information for map node types
- `map_nodes` - Stores map node instances

## Schema

### NodeTypeStyle Collection (`node_type_styles`)

```json
{
  "_id": "ObjectId",
  "type": "string",
  "color": "string",
  "maxRadius": "float64"
}
```

**Fields:**
- `_id` (ObjectId): MongoDB auto-generated unique identifier
- `type` (string): Node type identifier (e.g., 'home', 'house', 'snack_shop', 'gym', 'transit', 'salon', 'other')
- `color` (string): Hex color code for map visualization (e.g., '#FF5733')
- `maxRadius` (float64): Maximum influence radius in meters

### MapNode Collection (`map_nodes`)

```json
{
  "_id": "ObjectId",
  "type": "string",
  "label": "string",
  "latitude": "float64",
  "longitude": "float64",
  "description": "string (optional)"
}
```

**Fields:**
- `_id` (ObjectId): MongoDB auto-generated unique identifier
- `type` (string): Node type from NodeTypeStyle (e.g., 'home', 'house', etc.)
- `label` (string): Display label for the node
- `latitude` (float64): Geographic latitude coordinate
- `longitude` (float64): Geographic longitude coordinate
- `description` (string, optional): Additional information about the node