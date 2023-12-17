# PlacementBuddy📊
A Simple HTTP(REST) API to make a dashboard like service for college placement. It helps keep a check of companies coming for recruitment.👩‍💻


## Functionalities👾
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



## Installation🛠️
#### Dependencies : GO, Gofr, Docker, mysql
1.  #### Install Go modules
   
    `go mod init folder-name/Placement-buddy`
    
    `go mod tidy`

3.  #### Install Gofr modules

     `go get gofr.dev/pkg/gofr`
   
3.  #### Create DB
    You can run the mysql server and create a database locally using the following docker command:
   
    `docker run --name placement-buddy -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=placement_buddy -p 3306:3306 -d mysql:8.0.30`

4.  #### Create Table
     Access placment_buddy database and create table companies with columns id and name:
   
    `docker exec -it placement-buddy mysql -uroot -proot123 placement_buddy -e "[SQL CODE FOR DB]"`

    can check in your mysql workbench if avlbl
5.  ####  RUN
     run your API using
    
     `go run main.go`




## Testing🧪
#### try out APIs using Postman

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/31381553-894c778a-6717-4224-a7e1-a24430b3be3f?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D31381553-894c778a-6717-4224-a7e1-a24430b3be3f%26entityType%3Dcollection%26workspaceId%3D17886b48-98fa-4cdf-972e-614275d26197)

#### Unit Test




## More about the API📑
#### Database Structure
<img width="381" alt="image" src="https://github.com/ananya-codes/PlacementBuddy/assets/77432683/8bed5140-1793-4711-b0ea-dda02c888b24">

#### Sequence Diagram
<img width="591" alt="image" src="https://github.com/ananya-codes/PlacementBuddy/assets/77432683/9b390126-1eb0-4834-9985-b0b53a213ca0">



## Refer🔗

[Go Official Documentaion](https://go.dev/)

[Gofr Official Documentaion](https://gofr.dev/)

[Gofr Github Repo](https://github.com/gofr-dev/gofr)




#### _Thankyou!_
