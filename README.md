docker build -t my-postgres .  
docker run -d --name my-postgres-container -p 5432:5432 my-postgres