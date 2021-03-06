package main

// taskPayload returns json schema for the payload part of the task definition
// please note we use a go string and do not load an external file, since we
// want this to be part of the compiled executable, and not rely on an external
// file
func taskPayloadSchema() string {
	return `{
  "id": "http://schemas.taskcluster.net/generic-worker/v1/payload.json#",
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Generic worker payload",
  "description": "This schema defines the structure of the ` + "`payload`" + ` property referred to in a Task Cluster Task definition.",
  "type": "object",
  "required": [
    "command",
    "maxRunTime"
  ],
  "properties": {
    "command": {
      "title": "Commands to run",
      "type": "array",
      "minItems": 1,
      "items": {
        "type": "string"
      },
	  "description": "One entry per command (consider each entry to be interpreted as a full line of a Windows™ .bat file). For example: ` + "`" + `[\"set\", \"echo hello world > hello_world.txt\", \"set GOPATH=C:\\\\Go\"]` + "`" + `."
    },
    "env": {
      "title": "Environment variable mappings.",
	  "description": "Example: ` + "```" + `{ \"PATH\": \"C:\\\\Windows\\\\system32;C:\\\\Windows\", \"GOOS\": \"darwin\" }` + "```" + `",
      "type": "object"
    },
    "maxRunTime": {
      "type": "number",
      "title": "Maximum run time in seconds",
      "description": "Maximum time the task container can run in seconds",
      "multipleOf": 1.0,
      "minimum": 1,
      "maximum": 86400
    },
    "artifacts": {
      "type": "array",
      "title": "Artifacts to be published",
	  "description": "Artifacts to be published. For example: ` + "`" + `{ \"type\": \"file\", \"path\": \"builds\\\\firefox.exe\", \"expires\": \"2015-08-19T17:30:00.000Z\" }` + "`" + `",
      "items": {
        "type": "object",
          "properties": {
          "type": {
            "title": "Artifact upload type.",
            "type": "string",
            "enum": ["file"],
			"description": "Currently only ` + "`file`" + ` is supported"
          },
          "path": {
            "title": "Artifact location",
            "type": "string",
            "description": "Location of artifact in container"
          },
          "expires": {
            "title": "Expiry date and time",
            "type": "string",
            "format": "date-time",
            "description": "Date when artifact should expire must be in the future"
          }
        },
        "required": ["type", "path", "expires"]
      }
    }
  }
}`
}
