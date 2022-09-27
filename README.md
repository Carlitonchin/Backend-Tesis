# Instrucciones para correr el proyecto

### Backend:
`sudo docker compose up`

### Frontend:
`cd frontend/`\
`yarn install`\
`yarn dev`

El backend utiliza varios puertos para ejecutar la servidor en gin, la base de datos postgres, y la base de datos redis, por lo que puede saltar un error si tiene algún otro proceso en su pc utilizando los puertos 8080, 5432, 6379

