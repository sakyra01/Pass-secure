# Pass-secure


This project is a simple golang microservice. It goes along with the project
https://self-service-password.readthedocs.io/en/latest/

The microservice allows the main program LDAP Tool Box Self Service Password to perform an additional security test on password strength. 
At the replacement stage, the password is sent to the microservice with a post request, 
the weakpass microservice looks for the password in its database using the API. 
The database is updated separately every six months. 
After verification, the microservice returns the answer whether the password is safe or not.

# How to use:

  1. At the moment, you need to add changes to the code of the main program to enable the microservice to work, the test php code is in test.php.
  2. You need configurate your .env file for postgresql docker connection.
  3. Project in progress
  4. There is binary file which you coud just run by ./weakpass, binary file was build for linux arch amd64
