# Mocking API
Backend service to create mocking endpoint  
For CRUD mock api, this project use url :
```  
GET /api/v1/mock  
POST /api/v1/mock 
PUT /api/v1/mock 
DELETE /api/v1/mock  
```  

## RUNNING DOCKER COMPOSE
```  
docker-compose up -d 
docker ps 
```  
  
## RUNNING IMAGES TO CONTAINER  
```  
docker images docker run -p 8080:8080 --name my-golang-app my-image-name
```  
  
example :  
```  
docker run -p 8080:8080 --name mocking_api_docker_file mocking_api_docker_file
```