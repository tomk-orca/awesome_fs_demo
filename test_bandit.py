# test_bandit.py

import hashlib
import os
import subprocess

# 1. Hardcoded password (Bandit will flag this)
def connect_to_db():
    password = "my-secret-password"  # Hardcoded password (insecure)
    return f"Connecting to the database with password: {password}"

# 2. Use of eval() (Bandit will flag this as potentially dangerous)
def unsafe_eval(user_input):
    result = eval(user_input)  # Avoid using eval() with untrusted input
    return result

# 3. Use of weak MD5 hashing (Bandit will flag this)
def hash_password(password):
    return hashlib.md5(password.encode()).hexdigest()  # MD5 is insecure

# 4. Shell injection vulnerability (Bandit will flag this)
def run_system_command(command):
    subprocess.call(command, shell=True)  # Dangerous shell=True

# 5. Use of os.system (Bandit will flag this)
def run_os_command():
    os.system('ls -la')  # os.system should be avoided

if __name__ == "__main__":
    print(connect_to_db())
    print(unsafe_eval("5 + 10"))
    print(hash_password("password123"))
    run_system_command("echo Hello World")
    run_os_command()