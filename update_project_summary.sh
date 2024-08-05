#!/bin/bash

# Create a new summary file
summary_file="project_summary.txt"
touch "$summary_file"

# Project Description
echo "### Project Summary ###" > "$summary_file"
echo "" >> "$summary_file"
echo "Project: Recipe Website for Peruvian Dishes" >> "$summary_file"
echo "Description: A web application where users can input ingredients and receive Peruvian dish recipes." >> "$summary_file"
echo "" >> "$summary_file"

# Software and Tools
echo "### Software and Tools ###" >> "$summary_file"
{
    echo "Python: $(python3 --version)"
    echo "Django: $(django-admin --version)"
    echo "Java: $(java -version 2>&1 | grep 'openjdk')"
    echo "Spring: $(spring --version)"
    echo "PostgreSQL: $(psql --version)"
    echo "Docker: $(docker --version)"
    echo "Kubectl: $(kubectl version --client -o yaml | awk '/gitVersion/{print $2;}')"
    echo "AWS CLI: $(aws --version)"
    echo "Google Cloud SDK: $(gcloud version | head -n 1)"
    echo "Jenkins: $(jenkins --version 2>/dev/null || echo 'Jenkins not found')"
    echo "Git: $(git --version)"
} >> "$summary_file"
echo "" >> "$summary_file"

# Project Directory Structure
echo "### Project Directory Structure ###" >> "$summary_file"
tree -L 3 >> "$summary_file"
echo "" >> "$summary_file"

# Git Repository Status
echo "### Git Repository Status ###" >> "$summary_file"
{
    echo "Current Branch: $(git branch --show-current)"
    echo "Status:"
    git status
} >> "$summary_file"
echo "" >> "$summary_file"

# Disk Usage
echo "### Disk Usage ###" >> "$summary_file"
df -h >> "$summary_file"
echo "" >> "$summary_file"

# Network Configuration
echo "### Network Configuration ###" >> "$summary_file"
ifconfig >> "$summary_file"
echo "" >> "$summary_file"

echo "Project summary updated successfully."
