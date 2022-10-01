# Go Installation 

---

- ### linux
**Download Go Installer for Linux** . Click [here](https://go.dev/dl/go1.19.1.linux-amd64.tar.gz)

1. Remove any previous Go installation by deleting the */usr/local/go* folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
`$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz `
(You may need to run the command as root or through sudo).

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

2. Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

`export PATH=$PATH:/usr/local/go/bin`
Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

3. Verify that you've installed Go by opening a command prompt and typing the following command:
`$ go version`
4. Confirm that the command prints the installed version of Go.

---

- ### Windows

**Download Go Installer for Windows** . Click [here](https://go.dev/dl/go1.19.1.windows-amd64.msi)

1. Open the MSI file you downloaded and follow the prompts to install Go.
By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

2. Verify that you've installed Go.
   - In Windows, click the Start menu.
   - In the menu's search box, type cmd, then press the Enter key.
   - In the Command Prompt window that appears, type the following command:
   `$ go version`
   - Confirm that the command prints the installed version of Go.

---

### Mac

**Download Go Installer for Mac** . Click [here](https://go.dev/dl/go1.19.1.darwin-amd64.pkg)
1. Open the package file you downloaded and follow the prompts to install Go.
The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

2. Verify that you've installed Go by opening a command prompt and typing the following command:
`$ go version`
3. Confirm that the command prints the installed version of Go.

---

**For more details about installation visit [Go official installation page](https://go.dev/doc/install).**