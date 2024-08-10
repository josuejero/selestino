import os
import subprocess

def run_command(command, description):
    print(f"Running: {description}...\nCommand: {command}\n")
    result = subprocess.run(command, shell=True, capture_output=True, text=True)
    if result.returncode != 0:
        output = f"Error: {result.stderr.strip()}\n"
    else:
        output = result.stdout.strip()
    print(output + "\n")
    return output

def main():
    # Prompting user for custom names and paths
    project_name = input("Enter your Django project name (e.g., 'myproject'): ")
    db_name = input("Enter your PostgreSQL database name (e.g., 'myprojectdb'): ")
    db_user = input("Enter your PostgreSQL username (e.g., 'myprojectuser'): ")
    db_password = input("Enter your PostgreSQL password: ")
    github_username = input("Enter your GitHub username: ")
    
    verification_output = []

    # Step 1: Install Python and Django
    verification_output.append("Step 1: Install Python and Django")

    # 1.1 Check Python Installation
    verification_output.append("Check Python Installation:")
    python_version = run_command("python3 --version", "Checking Python version")
    verification_output.append(python_version)

    # 1.2 Check pip and virtualenv installation
    verification_output.append("Check pip and virtualenv Installation:")
    pip_version = run_command("python3 -m pip --version", "Checking pip version")
    verification_output.append(pip_version)
    virtualenv_version = run_command("virtualenv --version", "Checking virtualenv version")
    verification_output.append(virtualenv_version)

    # 1.3 Check Virtual Environment Activation
    verification_output.append("Check Virtual Environment Activation:")
    env_name = os.getenv('VIRTUAL_ENV')
    if env_name:
        verification_output.append(f"Virtual environment is activated: {env_name}")
    else:
        verification_output.append("Virtual environment is NOT activated")

    # 1.4 Check Django Installation
    verification_output.append("Check Django Installation:")
    django_version = run_command("django-admin --version", "Checking Django version")
    verification_output.append(django_version)

    # 1.5 Check Django Project Structure
    verification_output.append("Check Django Project Structure:")
    project_structure = run_command(f"ls {project_name}", f"Listing Django project files in {project_name}")
    verification_output.append(project_structure)

    # Step 2: Database Setup with PostgreSQL
    verification_output.append("\nStep 2: Database Setup with PostgreSQL")

    # 2.1 Check PostgreSQL Installation
    verification_output.append("Check PostgreSQL Installation:")
    psql_path = run_command("which psql", "Checking PostgreSQL installation")
    verification_output.append(psql_path)

    # 2.2 Verify Database Creation
    verification_output.append("Verify PostgreSQL Database Creation:")
    db_list = run_command("psql -c '\\l'", "Listing PostgreSQL databases")
    verification_output.append(db_list)

    # 2.3 Test Django-PostgreSQL Connection
    verification_output.append("Test Django-PostgreSQL Connection:")
    migrate_result = run_command(f"python manage.py migrate", "Testing database migration")
    verification_output.append(migrate_result)

    # Step 3: Version Control with Git
    verification_output.append("\nStep 3: Version Control with Git")

    # 3.1 Check Git Status
    verification_output.append("Check Git Status:")
    git_status = run_command("git status", "Checking git status")
    verification_output.append(git_status)

    # 3.2 Check Git Log
    verification_output.append("Check Git Log:")
    git_log = run_command("git log --oneline", "Checking git log")
    verification_output.append(git_log)

    # Write the outputs to a file
    with open("verification_output.txt", "w") as f:
        for line in verification_output:
            f.write(line + "\n")

    print("All steps completed. Output saved to 'verification_output.txt'.")

if __name__ == "__main__":
    main()
