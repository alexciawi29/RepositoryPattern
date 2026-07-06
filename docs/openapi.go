package docs

import "github.com/gin-gonic/gin"

const OpenAPISpec = `{
    "openapi": "3.0.1",
    "info": {
        "title": "Tutorial Go API",
        "description": "REST API with Gin, GORM, and Swagger — v1 uses custom query params, v2 uses OData",
        "version": "v1"
    },
    "servers": [
        {
            "url": "http://localhost:8080",
            "description": "Local development server"
        }
    ],
    "security": [
        {"BearerAuth": []}
    ],
    "paths": {
        "/go/products": {
            "get": {
                "tags": ["Products"],
                "security": [],
                "summary": "Get all products",
                "description": "Returns a paginated list of products with optional sorting, filtering, and search",
                "parameters": [
                    {"name": "top", "in": "query", "description": "Number of records to return (default: 10)", "schema": {"type": "integer"}},
                    {"name": "skip", "in": "query", "description": "Number of records to skip (default: 0)", "schema": {"type": "integer"}},
                    {"name": "sortby", "in": "query", "description": "Field name to sort by (e.g. name, price)", "schema": {"type": "string"}},
                    {"name": "order", "in": "query", "description": "Sort direction: asc or desc (default: asc)", "schema": {"type": "string"}},
                    {"name": "search", "in": "query", "description": "Search term matched against product name", "schema": {"type": "string"}},
                    {"name": "filter", "in": "query", "description": "Filter by field value (e.g. filter[name]=Laptop)", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/PaginationResultProduct"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "post": {
                "tags": ["Products"],
                "security": [],
                "summary": "Create a product",
                "description": "Creates a new product with name and price",
                "requestBody": {
                    "required": true,
                    "content": {"application/json": {"schema": {"$ref": "#/components/schemas/CreateProductRequest"}}}
                },
                "responses": {
                    "201": {
                        "description": "Created",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Product"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/products/{id}": {
            "get": {
                "tags": ["Products"],
                "security": [],
                "summary": "Get product by ID",
                "description": "Returns a single product by ID",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "Product ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Product"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "404": {
                        "description": "Not Found",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "put": {
                "tags": ["Products"],
                "security": [],
                "summary": "Update a product",
                "description": "Updates a product by ID",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "Product ID", "schema": {"type": "integer"}}
                ],
                "requestBody": {
                    "required": true,
                    "content": {"application/json": {"schema": {"$ref": "#/components/schemas/UpdateProductRequest"}}}
                },
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Product"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "delete": {
                "tags": ["Products"],
                "security": [],
                "summary": "Delete a product",
                "description": "Deletes a product by ID (soft delete via GORM)",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "Product ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/MessageResponse"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/users": {
            "get": {
                "tags": ["Users"],
                "security": [],
                "summary": "Get all users",
                "description": "Returns a paginated list of users with optional sorting, filtering, and search",
                "parameters": [
                    {"name": "top", "in": "query", "description": "Number of records to return (default: 10)", "schema": {"type": "integer"}},
                    {"name": "skip", "in": "query", "description": "Number of records to skip (default: 0)", "schema": {"type": "integer"}},
                    {"name": "sortby", "in": "query", "description": "Field name to sort by (e.g. name, email)", "schema": {"type": "string"}},
                    {"name": "order", "in": "query", "description": "Sort direction: asc or desc (default: asc)", "schema": {"type": "string"}},
                    {"name": "search", "in": "query", "description": "Search term matched against name and email", "schema": {"type": "string"}},
                    {"name": "filter", "in": "query", "description": "Filter by field value (e.g. filter[name]=Herris)", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/PaginationResultUser"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "post": {
                "tags": ["Users"],
                "security": [],
                "summary": "Create a user",
                "description": "Creates a new user with name and email",
                "requestBody": {
                    "required": true,
                    "content": {"application/json": {"schema": {"$ref": "#/components/schemas/CreateUserRequest"}}}
                },
                "responses": {
                    "201": {
                        "description": "Created",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/User"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/users/{id}": {
            "get": {
                "tags": ["Users"],
                "security": [],
                "summary": "Get user by ID",
                "description": "Returns a single user by ID",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "User ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/User"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "404": {
                        "description": "Not Found",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "put": {
                "tags": ["Users"],
                "security": [],
                "summary": "Update a user",
                "description": "Updates a user by ID",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "User ID", "schema": {"type": "integer"}}
                ],
                "requestBody": {
                    "required": true,
                    "content": {"application/json": {"schema": {"$ref": "#/components/schemas/UpdateUserRequest"}}}
                },
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/User"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            },
            "delete": {
                "tags": ["Users"],
                "security": [],
                "summary": "Delete a user",
                "description": "Deletes a user by ID (soft delete via GORM)",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "User ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/MessageResponse"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/": {
            "get": {
                "tags": ["OData"],
                "summary": "OData v4 service document",
                "description": "Lists all available entity sets in the v2 API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ServiceDocumentResponse"}}}
                    }
                }
            }
        },
        "/go/odata/$metadata": {
            "get": {
                "tags": ["OData"],
                "summary": "OData v4 metadata document",
                "description": "Returns the EDMX metadata document describing all entity types and their properties",
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"text/xml": {"schema": {"type": "string"}}}
                    }
                }
            }
        },
        "/go/odata/products": {
            "get": {
                "tags": ["Products"],
                "summary": "Get all products (OData v4)",
                "description": "Returns a paginated product collection using OData v4 query options",
                "parameters": [
                    {"name": "$top", "in": "query", "description": "Number of records to return (default: 10)", "schema": {"type": "integer"}},
                    {"name": "$skip", "in": "query", "description": "Number of records to skip (default: 0)", "schema": {"type": "integer"}},
                    {"name": "$orderby", "in": "query", "description": "Sort field and direction (e.g. name asc, price desc)", "schema": {"type": "string"}},
                    {"name": "$filter", "in": "query", "description": "OData filter (e.g. name eq 'Laptop' and price gt 10000, contains(name,'lap'), startswith(name,'L'))", "schema": {"type": "string"}},
                    {"name": "$search", "in": "query", "description": "Search term matched against product name", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ODataPagedResponseProduct"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/products/$count": {
            "get": {
                "tags": ["Products"],
                "summary": "Get product count (OData v4)",
                "description": "Returns the total number of products as plain text",
                "parameters": [
                    {"name": "$filter", "in": "query", "description": "OData filter expression", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"text/plain": {"schema": {"type": "string"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/products/{id}": {
            "get": {
                "tags": ["Products"],
                "summary": "Get product by ID (OData v4)",
                "description": "Returns a single product with OData v4 context",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "Product ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"type": "object", "additionalProperties": true}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "404": {
                        "description": "Not Found",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/users": {
            "get": {
                "tags": ["Users"],
                "summary": "Get all users (OData v4)",
                "description": "Returns a paginated user collection using OData v4 query options",
                "parameters": [
                    {"name": "$top", "in": "query", "description": "Number of records to return (default: 10)", "schema": {"type": "integer"}},
                    {"name": "$skip", "in": "query", "description": "Number of records to skip (default: 0)", "schema": {"type": "integer"}},
                    {"name": "$orderby", "in": "query", "description": "Sort field and direction (e.g. name asc, email desc)", "schema": {"type": "string"}},
                    {"name": "$filter", "in": "query", "description": "OData filter (e.g. name eq 'Herris', contains(email,'@example.com'))", "schema": {"type": "string"}},
                    {"name": "$search", "in": "query", "description": "Search term matched against name and email", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ODataPagedResponseUser"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/users/$count": {
            "get": {
                "tags": ["Users"],
                "summary": "Get user count (OData v4)",
                "description": "Returns the total number of users as plain text",
                "parameters": [
                    {"name": "$filter", "in": "query", "description": "OData filter expression", "schema": {"type": "string"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"text/plain": {"schema": {"type": "string"}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        },
        "/go/odata/users/{id}": {
            "get": {
                "tags": ["Users"],
                "summary": "Get user by ID (OData v4)",
                "description": "Returns a single user with OData v4 context",
                "parameters": [
                    {"name": "id", "in": "path", "required": true, "description": "User ID", "schema": {"type": "integer"}}
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {"application/json": {"schema": {"type": "object", "additionalProperties": true}}}
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    },
                    "404": {
                        "description": "Not Found",
                        "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ErrorResponse"}}}
                    }
                }
            }
        }
    },
    "components": {
        "schemas": {
            "Product": {
                "type": "object",
                "properties": {
                    "id": {"type": "integer"},
                    "name": {"type": "string"},
                    "price": {"type": "number"},
                    "createdAt": {"type": "string"},
                    "updatedAt": {"type": "string"},
                    "deletedAt": {"$ref": "#/components/schemas/DeletedAt"}
                }
            },
            "User": {
                "type": "object",
                "properties": {
                    "id": {"type": "integer"},
                    "name": {"type": "string"},
                    "email": {"type": "string"},
                    "createdAt": {"type": "string"},
                    "updatedAt": {"type": "string"},
                    "deletedAt": {"$ref": "#/components/schemas/DeletedAt"}
                }
            },
            "DeletedAt": {
                "type": "object",
                "properties": {
                    "time": {"type": "string"},
                    "valid": {"type": "boolean", "description": "Valid is true if Time is not NULL"}
                }
            },
            "PaginationResultProduct": {
                "type": "object",
                "properties": {
                    "Data": {"type": "array", "items": {"$ref": "#/components/schemas/Product"}},
                    "Total": {"type": "integer"},
                    "Top": {"type": "integer"},
                    "Skip": {"type": "integer"},
                    "TotalPages": {"type": "integer"}
                }
            },
            "PaginationResultUser": {
                "type": "object",
                "properties": {
                    "Data": {"type": "array", "items": {"$ref": "#/components/schemas/User"}},
                    "Total": {"type": "integer"},
                    "Top": {"type": "integer"},
                    "Skip": {"type": "integer"},
                    "TotalPages": {"type": "integer"}
                }
            },
            "ErrorResponse": {
                "type": "object",
                "properties": {
                    "Error": {"type": "string", "example": "error message"}
                }
            },
            "MessageResponse": {
                "type": "object",
                "properties": {
                    "Message": {"type": "string", "example": "operation successful"}
                }
            },
            "CreateProductRequest": {
                "type": "object",
                "required": ["Name", "Price"],
                "properties": {
                    "Name": {"type": "string", "example": "Laptop"},
                    "Price": {"type": "number", "example": 15000000}
                }
            },
            "UpdateProductRequest": {
                "type": "object",
                "properties": {
                    "Name": {"type": "string", "example": "Laptop Pro"},
                    "Price": {"type": "number", "example": 18000000}
                }
            },
            "CreateUserRequest": {
                "type": "object",
                "required": ["Name", "Email"],
                "properties": {
                    "Name": {"type": "string", "example": "Herris"},
                    "Email": {"type": "string", "example": "herris@example.com"}
                }
            },
            "UpdateUserRequest": {
                "type": "object",
                "properties": {
                    "Name": {"type": "string", "example": "Herris Updated"},
                    "Email": {"type": "string", "example": "herris.new@example.com"}
                }
            },
            "ODataPagedResponseProduct": {
                "type": "object",
                "properties": {
                    "@odata.context": {"type": "string"},
                    "@odata.count": {"type": "integer"},
                    "@odata.nextLink": {"type": "string"},
                    "value": {"type": "array", "items": {"$ref": "#/components/schemas/Product"}}
                }
            },
            "ODataPagedResponseUser": {
                "type": "object",
                "properties": {
                    "@odata.context": {"type": "string"},
                    "@odata.count": {"type": "integer"},
                    "@odata.nextLink": {"type": "string"},
                    "value": {"type": "array", "items": {"$ref": "#/components/schemas/User"}}
                }
            },
            "ServiceDocumentEntry": {
                "type": "object",
                "properties": {
                    "name": {"type": "string"},
                    "kind": {"type": "string"},
                    "url": {"type": "string"}
                }
            },
            "ServiceDocumentResponse": {
                "type": "object",
                "properties": {
                    "@odata.context": {"type": "string"},
                    "value": {"type": "array", "items": {"$ref": "#/components/schemas/ServiceDocumentEntry"}}
                }
            }
        },
        "securitySchemes": {
            "BearerAuth": {
                "type": "http",
                "scheme": "bearer",
                "bearerFormat": "JWT"
            }
        }
    }
}`

func ServeOpenAPIJSON(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(200, OpenAPISpec)
}
