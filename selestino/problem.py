import os

def list_files(startpath):
    with open('project_structure_and_code.txt', 'w') as f:
        for root, dirs, files in os.walk(startpath):
            # Exclude the venv directory
            dirs[:] = [d for d in dirs if d != 'venv']
            level = root.replace(startpath, '').count(os.sep)
            indent = ' ' * 4 * (level)
            f.write('{}{}/\n'.format(indent, os.path.basename(root)))
            subindent = ' ' * 4 * (level + 1)
            for file in files:
                f.write('{}{}\n'.format(subindent, file))
                if file.endswith('.py'):
                    f.write('\n{}# Content of {}:\n'.format(subindent, file))
                    with open(os.path.join(root, file), 'r') as py_file:
                        for line in py_file:
                            f.write('{}{}'.format(subindent, line))
                    f.write('\n')

if __name__ == "__main__":
    project_root = os.path.dirname(os.path.abspath(__file__))
    list_files(project_root)
    print("Project structure and code have been written to 'project_structure_and_code.txt'")
