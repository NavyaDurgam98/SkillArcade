# Sprint 1 Report

## Visual Demo Links
- [Frontend](https://drive.google.com/file/d/1wLm-qh37Ih8RsFMu-TMAoJNtcbJZStD_/view?usp=sharing)
- [Backend](https://uflorida-my.sharepoint.com/:v:/g/personal/spabbathi_ufl_edu/Ean2Ve1wj4xEtgGceMtBZ4sBMSX4zHehqoICCcBVmN8iBQ?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=0E9E5X)

## User Stories:
- Created 7 user stories spanning the features/user requirements of the Skill Arcade application. 
- Addition of more user stories depends on the scope of the project and resource availability in the future. 
- These user stories can be found under the issues tab with the label [userstories](https://github.com/ssaditya/Ecommute-SE_Project/issues?q=is%3Aopen+is%3Aissue+label%3A%22user+stories%22) in this repository.

## Development:
### Backend:  
- Set up Go using the Gin/Gonic framework and MongoDB as the database.  
- Established a database connection and implemented four API routes.  
- Created a test API with a POST operation to store data in the database.  
- Developed SignUp, SignIn, Forgot Password, and Reset Password APIs.  
- Implemented JWT authentication for session-based sign-in.  
- Integrated Twilio (a third-party mailing API) to handle Forgot Password requests.
### FrontEnd:
- Did the initial setup of the react application and installed required libraries/dependencies.
- Created the MainComponent and set up the routing paths.
- Created the Home Screen Page with the driver and rider modes.
- Developed the HeaderComponent.
- Created and partially developed the DriverComponent, RiderComponent with their respective navbars.
## Testing:
- Tested test-API using Postman and validated respective changes in DB.
- Tested CRUD operations in DB by calling APIs through Postman.
- Tested the User Interface by hosting on a localhost server.
