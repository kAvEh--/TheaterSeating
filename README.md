# TheaterSeating
### Theater Seating Algorithm/API implementation
- Create a data structure that defines a seating layout for a hall in a venue
- Given a list of “groups of users” per rank (basically sizes, e.g. (1, 3, 4, 4, 5, 1, 2, 4) in a specific order
- Design / Create a REST API to retrieve the layout of the allocations
- Improve the algorithm in such a way that no individual people sit alone
- Create a deployment of the application

#### Approach to solution:

This problem is a combination of CSP(Constraint Satisfaction Problem) and Bin Packaging problem.
Both of these problems are NP-hard problem and has no polynomial solutions.

CSPs are mathematical questions defined as a set of objects whose state must satisfy a number of constraints or limitations.
The bin Packaging is an optimization problem, in which items of different sizes must be packed into a finite number of bins or containers.

There are many approaches to Bin Packaging problem like next fit first fit, worst fit, and ...
In this repo three algorithm and compare between them and choose the best answer:
- direct seating algorithm. place the users continuously.
- First Fit algorithm. Find the nearest row that has enough space to that no individual people sit alone.
- Best Fit algorithm. Find the lowest capacity row that no individual people sit alone.

### Database
According to CAP theorem, because we need high consistency in this problem the SQL situation is chosen.

The PostgreSQl database will be run in Docker  

### API
the API implemented with these endpoints:
- PUT >> /hall : create new hall.
- GET >> /hall/:id : get hall data with id
- GET >> /user/:id : get user data and reservation list
- POST >> /reserve/:hall : send list of user group with rank

### Deployment

build application
```shell
make 
```
Create docker file-compose file:
```shell
docker-compose up -d --build
```
Implement the 12 factor application.

#### Dependecies
- Configuration: github.com/spf13/viper
- ORM: github.com/jinzhu/gorm
- Logging: github.com/sirupsen/logrus
- HTTP Server : github.com/gofiber/fiber