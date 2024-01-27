#!/bin/bash
curl http://localhost:8081 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'

