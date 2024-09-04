import os

def list_files(startpath):
    with open('project_structure_and_code.txt', 'w') as f:
        for root, dirs, files in os.walk(startpath):
            # Exclude the 'env' and 'django_extensions' directories
            dirs[:] = [d for d in dirs if d not in ['venv', 'django_extensions', "admin"]]
            level = root.replace(startpath, '').count(os.sep)
            indent = ' ' * 4 * (level)
            f.write('{}{}/\n'.format(indent, os.path.basename(root)))
            subindent = ' ' * 4 * (level + 1)
            for file in files:
                f.write('{}{}\n'.format(subindent, file))
                
                # List of file extensions to include their content
                if file.endswith(('.py', '.pem', '.ini', '.html', '.css', '.js', '.sh', 'Dockerfile', '.yml', 'Jenkinsfile')):
                    f.write('\n{}# Content of {}:\n'.format(subindent, file))
                    try:
                        with open(os.path.join(root, file), 'r') as specific_file:
                            for line in specific_file:
                                f.write('{}{}'.format(subindent, line))
                    except Exception as e:
                        f.write('{}# Could not read file: {}\n'.format(subindent, e))
                    f.write('\n')

if __name__ == "__main__":
    project_root = os.path.dirname(os.path.abspath(__file__))
    list_files(project_root)
    print("Project structure and code have been written to 'project_structure_and_code.txt'")
