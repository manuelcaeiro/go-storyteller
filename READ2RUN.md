As already mentioned on the README, the app looks and behaves like a desktop program, but actually runs on a local server and opens the computer default browser to use it as its Graphical User Interface (GUI).

In practice this means that in order to run it needs to start a local server on a preassigned port - and that should be shown on a command line window (on Windows) or on a terminal window (on Linux) that opens before the browser window when you click or double click on the executable file (storyteller/storyteller.exe).

![cmd](https://github.com/manuelcaeiro/go-storyteller/blob/master/screenshots/cmd_info.JPG)

![terminal](https://github.com/manuelcaeiro/go-storyteller/blob/master/screenshots/terminal_info.png)

If your Linux desktop presents you a pop up window with the option "execute in terminal" or "run in terminal" choose that option.

![execute in terminal](https://github.com/manuelcaeiro/go-storyteller/blob/master/screenshots/pop.png)

However, this cmd/terminal may not open on some Desktop Environments (DE), and the program just opens the browser. This is not an error and it only depends on your DE settings.

In this case, the local server process that supports the execution of the app cannot be stopped simply by closing the cmd/terminal, and will keep running unless you stop it. (1)

So, the right way to start the application next time would be from the command line or terminal. And that's actually very simple.

Open the folder where you have 2 files of the app that you downloaded previously and:

- On Windows: Put the mouse cursor anywhere on the empty space, click the right button with the Shift key pressed and choose from the shown menu the option "open command line here"; on the cmd window type (where the cursor is blinking) storyteller.exe and wait until the browser opens; (Be aware that closing the cmd window will stop the execution of both the program and the server.)
- On Linux: Right click anywhere on the empty space and the shown menu will have immediately the option "open terminal"; open the terminal, type ./storyteller and wait for the browser to open.

(1) On Windows: Right click on the Taskbar -> Open the Task Manager -> Click on the tab Processes -> Select the line starting with storyteller.exe and click the stop button on the right-down.<br>
On Linux: Open a terminal window from the start menu and type (or copy/paste) kill $(lsof -t -i :5016)
