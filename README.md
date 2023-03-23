# Pass-secure


This project is a simple golang microservice. It goes along with the project
https://self-service-password.readthedocs.io/en/latest/

The microservice allows the main program LDAP Tool Box Self Service Password to perform an additional security test on password strength. 
At the replacement stage, the password is sent to the microservice with a post request, 
the weakpass microservice looks for the password in its database using the API. 
The database is updated separately every six months. 
After verification, the microservice returns the answer whether the password is safe or not.
