# ICE 

[![Build Status](https://ci.sunho.kim/job/ICE/job/master/badge/icon)](https://ci.sunho.kim/job/ICE/job/master/display/redirect) 
Integrated Code Evaluator

## Run
You can run ice-judge by
```
docker-compose up
```
After this, site will be served on http://localhost:8080

For development, execute the script below
```
scripts/dev-compose-up.sh
```
This acts similarly but the site will use the host's ice/public/dist for assets and expose every service to  the outside. Basically this will reduce the need to rebuild the images.


