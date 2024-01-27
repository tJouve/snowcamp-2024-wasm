#!/bin/bash
hey -n 100 -c 1 -m POST \
-d 'John Doe' \
"http://localhost:8081" 


#hey -n 100 -c 10 -m POST \
#-d 'John Doe' \
#"http://localhost:8081" 
