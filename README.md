Hackstore CLI Tool
This command-line interface tool enables the management of programs, root domains, subdomains, and alive domains within Hackstore.

Getting Started
Prerequisites
Go installed.
MySQL database.
Installation

Clone the repository:

bash
Copy code
git clone <repository_url>

Install dependencies:

bash
Copy code
go mod tidy
Configuration
Create a configuration file hackstore-config.yml in the path ~/.config/hackstore/. The file should contain the following database connection details:

yaml
Copy code
username: your_username
password: your_password
host: your_host
port: your_port
database: your_database_name
Usage
To start using the Hackstore CLI tool, execute the following command:

bash
Copy code
hackstore [command]
Replace [command] with one of the available commands:

programs: Manage programs.
root-domains: Manage root domains.
subdomains: Manage subdomains.
alivedomains: Manage alive domains.
Example Usages
Manage Programs
List all programs:

bash
Copy code
hackstore programs list
Import programs from a file:

bash
Copy code
hackstore programs import -f programs.txt
Add a new program:

bash
Copy code
hackstore programs add-program -p test_program
Delete a program and associated data:

bash
Copy code
hackstore programs delete-program -p test_program
Manage Root Domains
List root domains by program:

bash
Copy code
hackstore root-domains list -p test_program
Import root domains from a file and associate with a program:

bash
Copy code
hackstore root-domains import -f root_domains.txt -p test_program
Add root domains by program:

bash
Copy code
hackstore root-domains add-root-domains -p test_program -r example.com -r anotherexample.com
Delete a root domain and associated data:

bash
Copy code
hackstore root-domains delete-root-domain -r example.com -p test_program
[... Continue this pattern for managing subdomains and alive domains ...]

Commands
Programs
list: List all programs.
import: Import programs from a file.
add-program: Add a new program.
delete-program: Delete a program and associated data.
Root Domains
list: List root domains by program.
import: Import root domains from a file.
add-root-domains: Add root domains by program.
delete-root-domain: Delete a root domain and associated data.
[... Continue listing commands for subdomains and alive domains ...]

Additional Notes
Ensure the hackstore-config.yml file contains valid MySQL connection details.
For flags and additional options, refer to the help section of each command using --help.
Contributors
Your Name