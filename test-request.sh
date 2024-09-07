#!/bin/bash

api_url="http://localhost:8080/foo"
response_json=$(curl -sS -X GET $api_url)
echo "Response from JSON data request:"
echo "$response_json"

api_url="http://localhost:8080/bar"
json_data='{"title": "Sample Post","body": "This is the body of the post", "userId": 1}'

response_json=$(curl -sS -X POST -H "Content-Type: application/json" -d "$json_data" $api_url)
echo "Response from JSON data request:"
echo "$response_json"

api_url="http://localhost:8080/bar"
json_data='{"title": "Sample Post","body": "This is the body of the post"}'
id_param="1"

response_json=$(curl -sS -X POST -H "Content-Type: application/json" -d "$json_data" $api_url)
echo "Response from JSON data request:"
echo "$response_json"
