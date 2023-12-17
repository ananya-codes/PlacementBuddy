# PlacementBuddy📊
A Simple HTTP(REST) API to make a dashboard like service for college placement. It helps keep a check of companies coming for recruitment.👩‍💻

## Functionalities
#### VIEW
  - `/dashboard`    view Welcome Dashboard 
  
  - `/company`      view All Companies and thier Info

  you can view _Company's name, Package, Selected students count, Ongoing or not_
#### ADD

  - `/company/{name}`      Add New Company
  
  - `/company/{name}/{package}`       Add New Company and it's Annual Package
   
  - `/selected/{id}/{selected}`       Add no of students selected by the company
  
#### UPDATE
  - `/stop/{id}`        Change the status of hiring to : no more hiring!

  - `/start/{id}`      Change the status of hiring to : now hiring!
   
  - `/selected/{id}/{selected}`     Update no of students selected by the company

#### DELETE
  - `/company/{id}`    Delete Company Info from the system using its ID

## Installation

#### Dependencies : GO, Gofr, Docker, mysql
1.  #### Create DB
    You can run the mysql server and create a database locally using the following docker command:
   
    `docker run --name placement-buddy -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=placement_buddy -p 3306:3306 -d mysql:8.0.30`

 2.  #### Create Table
     Access placment_buddy database and create table companies with columns id and name:
   
    `docker exec -it placement-buddy mysql -uroot -proot123 placement_buddy -e "[SQL CODE FOR DB]"`
