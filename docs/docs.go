// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/create": {
            "post": {
                "description": "Creates a shortened version of the provided URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Shorten a URL",
                "parameters": [
                    {
                        "description": "URL and User ID",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.inputUrl"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.shortenedUrlResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON format or missing fields",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/delete": {
            "post": {
                "description": "Deletes a shortened URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Delete a shortened URL",
                "parameters": [
                    {
                        "description": "URL to delete",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.inputDelUrl"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "URL field is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "URL does not exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/list/{user_id}": {
            "get": {
                "description": "Returns a list of URLs for the specified user ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "List URLs for a specific user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of URLs for the user",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "User ID is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns a message indicating the server is online",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Check if the server is online",
                "responses": {
                    "200": {
                        "description": "Server is online",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{shortenedUrl}": {
            "get": {
                "description": "Redirects to the original URL from a shortened URL code",
                "tags": [
                    "urls"
                ],
                "summary": "Redirect to the original URL from a shortened URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL code",
                        "name": "shortenedUrl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirects to the original URL",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Shortened URL is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "URL not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.inputDelUrl": {
            "type": "object",
            "properties": {
                "url": {
                    "description": "URL to delete",
                    "type": "string"
                }
            }
        },
        "handlers.inputUrl": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.shortenedUrlResponse": {
            "type": "object",
            "properties": {
                "shortened_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "zexd.onrender.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ZexD API",
	Description:      "API for a URL Shortener Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
