# Hackstore CLI Tool
This command-line interface tool enables the management of programs, root domains, subdomains, and alive domains within Hackstore using mysql database.

## Getting Started Prerequisites:
* Go installed.
* MySQL database.


## Creating a Database from the SQL Dump File

Follow these steps to create a database using the provided SQL dump file: `db_structure.sql`

### Prerequisites

- Ensure you have MySQL installed on your system.

### Steps

1. **Download the SQL Dump File:** Obtain the `db_structure.sql` file containing the database structure.

2. **Open Terminal (Command Prompt):** Launch your terminal or command prompt.

3. **Login to MySQL:** Log in to your MySQL server using your MySQL credentials.

   ```bash
   mysql -u your_username -p

## Create a New Database

### Create a New Empty Database

Execute the following SQL command to create a new empty database where you want to import the structure:

```sql
> CREATE DATABASE hackstore;
```
### Now Import the structure from SQL dump file

```sql
$ mysql -u your_username -p hackstore < db_structure.sql
```

## Installation
Clone the repository:

`git clone https://github.com/abhi-recon/hackstore.git`

`cd hackstore`

Install dependencies:

`go mod tidy`

Build tool

`go build -o hackstore`

## Configuration
Create a configuration file `hackstore-config.yml` in the path `~/.config/hackstore/`. The file should contain the following database connection details:

```yaml
username: your_username
password: your_password
host: your_host
port: your_port
database: your_database_name
```

## Usage
To start using the Hackstore CLI tool, execute the following command:

`hackstore [command]`

Replace [command] with one of the available commands:

```bash
programs: Manage programs.
root-domains: Manage root domains.
subdomains: Manage subdomains.
alivedomains: Manage alive domains. 
```
# Example Usages
### Manage Programs

List all programs:


`hackstore programs list`

Import programs from a file:

`hackstore programs import -f programs.txt`

Add a new program:

`hackstore programs add-program -p test_program`

Delete a program and associated data:

`hackstore programs delete-program -p test_program`

### Manage Root Domains
List root domains by program:

`hackstore root-domains list -p test_program`

Import root domains from a file and associate with a program:

`hackstore root-domains import -f root_domains.txt -p test_program`

Add root domains by program:

`hackstore root-domains add-root-domains -p test_program -r example.com -r anotherexample.com`

Delete a root domain and associated data:

`hackstore root-domains delete-root-domain -r example.com -p test_program`

[... Continue this pattern for managing subdomains and alive domains ...]


### Additional Notes

Ensure the `hackstore-config.yml` file contains valid MySQL connection details.
For flags and additional options, refer to the help section of each command using --help.

# Contributors
Abhishek Karle
## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://linktr.ee/abhishekkarle)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/abhishek-karle-47a84717a/)

[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/AbhishekKarle3)

