# Make sure you have the full Python tools
sudo apt update
sudo apt install python3-full python3-venv python3-pip -y

# Create a virtual environment
python3 -m venv ~/nds-env

# Activate it
source ~/nds-env/bin/activate

# Install ndspy inside
pip install ndspy

