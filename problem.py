import os
import subprocess

def create_project_overview():
    overview_file = "project_overview.txt"
    with open(overview_file, "w") as f:
        # Project Description
        f.write("### Project Summary ###\n\n")
        f.write("Project: Recipe Website for Peruvian Dishes\n")
        f.write("Description: A web application where users can input ingredients and receive Peruvian dish recipes.\n\n")

        # Software and Tools
        f.write("### Software and Tools ###\n")
        f.write(f"Python: {subprocess.getoutput('python3 --version')}\n")
        f.write(f"Django: {subprocess.getoutput('django-admin --version')}\n")
        f.write(f"Java: {subprocess.getoutput('java -version 2>&1 | grep openjdk')}\n")
        f.write(f"Spring: {subprocess.getoutput('spring --version')}\n")
        f.write(f"PostgreSQL: {subprocess.getoutput('psql --version')}\n")
        f.write(f"Docker: {subprocess.getoutput('docker --version')}\n")
        kubectl_version_output = subprocess.getoutput("kubectl version --client --short")
        kubectl_version_line = next((line for line in kubectl_version_output.splitlines() if "Client Version" in line), "Kubectl version not found")
        f.write(f"Kubectl: {kubectl_version_line}\n")
        f.write(f"AWS CLI: {subprocess.getoutput('aws --version')}\n")
        f.write(f"Google Cloud SDK: {subprocess.getoutput('gcloud version | head -n 1')}\n")
        f.write(f"Jenkins: {subprocess.getoutput('jenkins --version 2>/dev/null || echo Jenkins not found')}\n")
        f.write(f"Git: {subprocess.getoutput('git --version')}\n\n")

        # Project Directory Structure
        f.write("### Project Directory Structure ###\n")
        f.write(subprocess.getoutput("tree -L 5") + "\n\n")

        # Git Repository Status
        f.write("### Git Repository Status ###\n")
        f.write(f"Current Branch: {subprocess.getoutput('git branch --show-current')}\n")
        f.write("Status:\n")
        f.write(subprocess.getoutput('git status') + "\n\n")

        # Disk Usage
        f.write("### Disk Usage ###\n")
        f.write(subprocess.getoutput('df -h') + "\n\n")

        # Network Configuration
        f.write("### Network Configuration ###\n")
        f.write(subprocess.getoutput('ifconfig') + "\n\n")

        # Get user input for the problem with the code
        problem_message = get_multiline_input("Enter the problem with the code:")
        f.write(f"### Problem Description ###\n{problem_message}\n\n")
        f.write("Please analyze this project structure, description, and errors carefully in order to help me with the issue.\n\n")

        # Append the content of README.md to the file if it exists
        readme_path = "README.md"
        if os.path.exists(readme_path):
            with open(readme_path, "r") as readme_file:
                readme_content = readme_file.read()

            f.write("### README.md ###\n")
            f.write(readme_content)
            f.write("\n\n")

        # Append relevant files to project_overview.txt
        bash_script_content = """
        #!/bin/bash

        find . -type d \\( -name 'logs' -o -name 'pids' -o -name 'lib-cov' -o -name 'coverage' -o -name '.nyc_output' -o -name '.grunt' -o -name 'bower_components' -o -name 'build' -o -name 'node_modules' -o -name 'jspm_packages' -o -name 'web_modules' -o -name '.npm' -o -name '.rpt2_cache' -o -name '.rts2_cache_cjs' -o -name '.rts2_cache_es' -o -name '.rts2_cache_umd' -o -name '.cache' -o -name '.parcel-cache' -o -name '.next' -o -name 'out' -o -name '.nuxt' -o -name 'dist' -o -name '.temp' -o -name '.serverless' -o -name '.fusebox' -o -name '.dynamodb' -o -name '.vscode-test' -o -name '.yarn' -o -name 'venv' \\) -prune -o \\
        -type f \\( -name "*.go" -o -name "*.yaml" -o -name "*.yml" -o -name "*.md" -o -name "*.sh" -o -name "*.sql" -o -name "Jenkinsfile" -o -name "LICENSE" -o -name "dockerfile" -o -name "go.mod" -o -name "go.sum" -o -name "kubeconfig" -o -name "*.py" \\) \\
        -not -name "*.log" \\
        -not -name "*.tsbuildinfo" \\
        -not -name "npm-debug.log*" \\
        -not -name "yarn-debug.log*" \\
        -not -name "yarn-error.log*" \\
        -not -name "lerna-debug.log*" \\
        -not -name ".pnpm-debug.log*" \\
        -not -name "report.[0-9]*.[0-9]*.[0-9]*.[0-9]*.json" \\
        -not -name ".lock-wscript" \\
        -not -name ".eslintcache" \\
        -not -name ".stylelintcache" \\
        -not -name "*.pid" \\
        -not -name "*.seed" \\
        -not -name "*.pid.lock" \\
        -not -name "*.tgz" \\
        -not -name ".yarn-integrity" \\
        -not -name ".env" \\
        -not -name ".env.*" \\
        -not -name ".node_repl_history" \\
        -not -name ".pnp.*" \\
        -not -name ".DS_Store" \\
        -not -name "jenkins_backup.tar" \\
        -not -name "jenkins/jenkins_backup.tar" \\
        -exec sh -c 'echo "### {} ###" >> project_overview.txt && cat "{}" >> project_overview.txt' \\;

        echo "All code/text has been appended to project_overview.txt."
        """

        with open("append_code.sh", "w") as bash_script:
            bash_script.write(bash_script_content)

    # Run the Bash script
    subprocess.run(["bash", "append_code.sh"])

    # Clean up the temporary Bash script
    os.remove("append_code.sh")

    # Copy the content of project_overview.txt to the clipboard using pbcopy
    with open(overview_file, "r") as file:
        all_code_content = file.read()

    subprocess.run("pbcopy", text=True, input=all_code_content)

    print("All steps completed and project_overview.txt has been updated and copied to clipboard.")

def get_multiline_input(prompt):
    print(prompt)
    print("Type 'END' on a new line to finish.")
    lines = []
    while True:
        line = input()
        if line.strip().upper() == 'END':
            break
        lines.append(line)
    return "\n".join(lines)

if __name__ == "__main__":
    create_project_overview()
