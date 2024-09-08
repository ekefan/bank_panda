1.  Create a make file to hold crucial shell calls to setup the project easily
2.  Create a database schema ----- to me that would be after a system design and application requirement/user stories have been created
3.  Using migration files, create a database schema ---used go migrate
4.  Setup a docker postgres container, connect it and visualize using the psql console, or using table plus
5.  Run migrations to setup the database tables on the postgres server.
6.  Install SQLC, then write tests for the functions generated.... if you ask me, no need for that



Creating a Docker image
For single stage build: image size is big anyways
1. Define the base images for the application... using the FROM ... (since this is a golang application base image would be golang)
2. Declare the current working directory for the image using WORKDIR "<"/name_of_wd">
3. Move all needed files from root of the folder to the docker current working dir using COPY . . (first dot is root of folder second is the root of the image curr_working_dir)
4. Build the application. using RUN ... command to build the application go build...
5. EXPOSE: communicate the port of the application
6. Define the command to be run by default when a container is created using CMD ["<cmd>"]

For Multi-stage build to reduce size of the output image
use as to specify  as specific stage... eg FROM ... as builder

use From for each stage to get the images for that current stage, 

in run stage select a linux image, maybe alpine or something....
then copy with COPY --from=<"the build stage"> <"path to file from build stage"> <". or path to copy file to">