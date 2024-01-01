# Hackstore CLI Tool

This command-line interface tool enables management of programs, root domains, subdomains, and alive domains within Hackstore.

## Getting Started

### Prerequisites

- Go installed.
- MySQL database.

### Installation

1. Clone the repository:

   ```bash
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

For example:

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
Manage Subdomains
List subdomains associated with a root domain:

bash
Copy code
hackstore subdomains list -r example.com
Import subdomains from a file and associate with a root domain:

bash
Copy code
hackstore subdomains import -f subdomains.txt -r example.com
Add a subdomain to a root domain:

bash
Copy code
hackstore subdomains add-subdomain -r example.com -s sub.example.com
Delete a subdomain:

bash
Copy code
hackstore subdomains delete-subdomain -r example.com -s sub.example.com
Manage Alive Domains
List alive domains by root domain:

bash
Copy code
hackstore alivedomains list -r example.com
Import alive domains from a file and associate with a root domain:

bash
Copy code
hackstore alivedomains import -f alive_domains.txt -r example.com
Add an alive domain to a root domain:

bash
Copy code
hackstore alivedomains add-alive-domain -r example.com -a https://sub.example.com
Delete an alive domain:

bash
Copy code
hackstore alivedomains delete-alive-domain -r example.com -a https://sub.example.com

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
Subdomains
list: List subdomains associated with a root domain.
import: Import subdomains from a file.
add-subdomain: Add a subdomain to a root domain.
delete-subdomain: Delete a subdomain.
Alive Domains
list: List alive domains by root domain.
import: Import alive domains from a file.
add-alive-domain: Add an alive domain to a root domain.
delete-alive-domain: Delete an alive domain.
Additional Notes
Ensure the hackstore-config.yml file contains valid MySQL connection details.
For flags and additional options, refer to the help section of each command using --help.
Contributors
Your Name