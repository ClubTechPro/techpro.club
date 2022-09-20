# File-folder structure

This document describes in detail the general folder/file structure in serial order.

# assets

This folder contains all the static files like logos, css and js. It is further divided into three folders.

### app

Static files for the main application after login

### home

Static files for project and contributor home pages

### logos

Application and third party logos.

# sources

This folder holds the main `Golang` code for the application.

### libraries

External third-party apis (including Authentication) for all the project repository providers will be handled here. Currently, we have integrated only Github.

### common

Common functions, environment variables and constants are defined here.

### mailers

This project will rely on third party mail providers to send emails. From time to time, we may change the service providers according to our needs. Hence, its separeted.

### templates

The code to handle each file in the frontend template will be written here. Further bifurcation has been made to separate contributors code from projects.

## users

Users management is the key to the overall application and also, it is common for all, which is why it has been kept in a separate folder

# templates

This folder is where the html files are stored. It has been divided into **home** and **app** where home holds the files for project and contributor home pages and finally, app holds the files for the pages for the rest of the application.

In future, other pages such as contact, about, etc will also go to home folder
