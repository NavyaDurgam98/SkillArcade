# Sprint 1 Report

## Visual Demo Links
- [Frontend](https://drive.google.com/file/d/1wLm-qh37Ih8RsFMu-TMAoJNtcbJZStD_/view?usp=sharing)
- [Backend](https://uflorida-my.sharepoint.com/:v:/g/personal/spabbathi_ufl_edu/Ean2Ve1wj4xEtgGceMtBZ4sBMSX4zHehqoICCcBVmN8iBQ?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=0E9E5X)

## User Stories:
- Created 7 user stories spanning the features/user requirements of the Skill Arcade application. 
- Additional user stories were introduced in the Later sprint to address emerging requirements, refine existing features, and ensure better alignment with project goals and user needs.
- These user stories can be found under the issues tab with the label [userstories](https://github.com/NavyaDurgam98/SkillArcade/issues) in this repository.


## Issues Planned to Address  
 
### User Authentication Implementation  
- Develop APIs for **SignUp, SignIn, Forgot Password, and Reset Password**.  
- Implement **JWT authentication** for session-based login.  
- Validate **user input** and handle errors during authentication.  

### Frontend Development  
- Set up **Angular** and configure the project structure.  
- Implement **routing** between login and signup pages.  
- Develop **user input validation** and display appropriate error messages.  
- Integrate the frontend with **backend authentication APIs**.  

### Database Integration  
- Connect the **Go backend to MongoDB** for data storage.  
- Ensure **efficient querying and data handling**.  

### Bug Fixes and UI Enhancements  
- Resolve **Angular routing issues** preventing navigation between components.  
- Improve the **UI for a better user experience**.  

### Testing and Debugging  
- Use **Postman** to test APIs and validate database interactions.  
- Debug issues related to **form validation, authentication, and API responses**.  



## Development:
### Backend:  
- Set up Go using the Gin/Gonic framework and MongoDB as the database.  
- Established a database connection and implemented four API routes.  
- Created a test API with a POST operation to store data in the database.  
- Developed SignUp, SignIn, Forgot Password, and Reset Password APIs.  
- Implemented JWT authentication for session-based sign-in.  
- Integrated Twilio (a third-party mailing API) to handle Forgot Password requests.
#### Testing:
- Tested test-API using Postman and validated respective changes in DB.
- Tested CRUD operations in DB by calling APIs through Postman.
- Tested the User Interface by hosting on a localhost server.
### FrontEnd:
- Set up Angular by installing Node.js and Angular CLI, creating a project with ng new, and running ng serve.  
- Created three components: App, Login, and Signup, where the App component is the main one, and routing happens here.   
- Validation for all pages is done, and the user is displayed with error messages when they miss the required fields.

## Issues Not Completed and Reasons  

While significant progress was made, the following issues were not completed due to time constraints and learning curve challenges:  

### Frontend Out of Scope Tasks  
- **Dashboard development is still in progress** *(out of Sprint 1 scope)*.  
- **CORS issues** are causing API request failures *(out of Sprint 1 scope)*.  

### Backend Pending Tasks  
- **CORS implementation** is pending. *(out of Sprint 1 scope)*
- **Password hashing/encryption** is not yet implemented.  
- **Full API integration with the frontend** is still ongoing.  

## Reasons for Incompletion  

- Since team was working with **Go** and **Angular** for the first time, a significant portion of the sprint was spent learning these technologies, setting up the project structure, and understanding best practices.  

- Initial time was spent getting familiar with **Git workflows** and setting up the directory structure properly.  

- Integrating **MongoDB with Go** and structuring **API requests for the frontend** required additional time and debugging.  

- Issues like **routing errors in Angular** and **API CORS restrictions** delayed progress in API integration.  



## Closing Note  

Despite these challenges, **Sprint 1 was successful** in establishing a solid foundation for the project.  

Moving forward, our focus will be on:  
- Completing the **Backlog tasks**.  
- Resolving **CORS issues**.  
- Further improving **frontend-backend integration**.  




