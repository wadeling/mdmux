# mdmux
multi-dimension event multiplexer

# usage
1. vgo build
2. ./mdmux

# test
1. events trigger:
curl -v -X POST http://localhost:8080/trigger/events -d '{ "src": "github","uuid":"abc","ip":"1.1.1.1" }'  -H "Content-Type: application/json"
2. get current event list:
curl -v -X GET http://localhost:8080/trigger/events 
3. delete event-array's top events:
curl -v -X DELETE http://localhost:8080/trigger/events 
