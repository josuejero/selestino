import os
import subprocess

# Path to the file
file_path = "all_code.txt"

# Delete the content of the file
with open(file_path, "w") as file:
    pass

# Function to get multi-line input with a special end keyword
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

# Get user input for the problem with the code
problem_message = get_multiline_input("Enter the problem with the code:")

# Write the problem message and introductory text to the top of the file
with open(file_path, "a") as file:
    file.write(f"{problem_message}\n\n")
    file.write("Please analyze this project structure, description, and errors carefully in order to help me with the issue.\n\n")

# Append the content of README.md to the file
with open("README.md", "r") as readme_file:
    readme_content = readme_file.read()

with open(file_path, "a") as file:
    file.write("### README.md ###\n")
    file.write(readme_content)
    file.write("\n\n")

# Append the output of the tree command to the file
tree_output = subprocess.run(["tree"], capture_output=True, text=True)
with open(file_path, "a") as file:
    file.write(tree_output.stdout)
    file.write("\n")

# Define the Bash script content
bash_script_content = """#!/bin/bash

# Find and append relevant files to all_code.txt excluding .txt files and kubectl
find . -type f \\( -name "*.go" -o -name "*.yaml" -o -name "*.yml" -o -name "*.md" -o -name "*.sh" -o -name "*.sql" -o -name "Jenkinsfile" -o -name "LICENSE" -o -name "dockerfile" -o -name "go.mod" -o -name "go.sum" -o -name "kubeconfig" -o -name "*.py" -o -name "*.log" \\) \\
    -not -name "*.txt" \\
    -not -name "kubectl" \\
    -exec sh -c 'echo "### {} ###" >> all_code.txt && cat "{}" >> all_code.txt' \\;

echo "All code/text has been appended to all_code.txt."
"""

# Write the Bash script content to a temporary file
with open("append_code.sh", "w") as bash_script:
    bash_script.write(bash_script_content)

# Run the Bash script
subprocess.run(["bash", "append_code.sh"])

# Clean up the temporary Bash script
os.remove("append_code.sh")

# Copy the content of all_code.txt to the clipboard using pbcopy
with open(file_path, "r") as file:
    all_code_content = file.read()

subprocess.run("pbcopy", text=True, input=all_code_content)

print("All steps completed and all_code.txt has been updated and copied to clipboard.")
