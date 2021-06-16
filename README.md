# virusprank.vbs
How many times did you try to create a Fake Virus or a prank virus to feel like a hacker or to prank someone? Sometimes you succeed sometimes not, there is often batch file creation by which you’ll create a simple virus that may also cause some issues on the computer.

Step 1. Writing Sample Code
1. Open Notepad or Notepad++.

Here is the code to generate a message box on a windows computer.

X=MsgBox(“Message to print in box”,0+16,”Title”)

Type the above code in Notepad and save this notepad on your desktop, Give it some convincing name like “Computer” and Save the file as .vbs  extension, and select all files instead of *txt.  Like Computer.vbs


![vs](https://user-images.githubusercontent.com/67673221/122180917-9552ef00-cea8-11eb-80f5-d6b0a0f0c40b.JPG)



![vs1](https://user-images.githubusercontent.com/67673221/122180961-9f74ed80-cea8-11eb-934f-621d41ecafda.JPG)




Step 2. Understanding the Code
In above Code  X=MsgBox(“Message Description”,0+16,”Title”)  here is the explanation of this code.

Message Description  – This is what you want to show as a Message.
Button – Type of button, like OK, Yes, No, Cancel, etc
 Icon – Type of icon you want to show like Info icon, critical icon, etc
Title – Title of message Box.

1. You can write any number from 1,2,3 or 4 instead of 0 (before the ‘+’ symbol)
Here is the meaning of these numbers:

0 = OK Button,
1 = OK / Cancel Button,
2 = Abort / Retry / Ignore Button,
3 = Yes / No / Cancel Button,
4 = Yes / No Button,
5 = Retry / Cancel Button

2. You can write 32 or 48 or 64 instead of 16.
Here is the meaning of each number:

16 = Critical Icon,
32 = Help Icon,
48 = Warning Icon,
64 = Information Icon,

Step 3. Changing the Icon Of Harmless Virus
You need to change the icon of this file so that victim is intended to open the file. You can change this icon to a Computer icon to make it looks original, but you can’t directly change the icon, follow this step to change the icon.

Copy and paste the file to another location like in C drive.
Create Shortcut – right-click on it and send it to desktop.
On the desktop you will get the shortcut, Right-click on it go to properties and click on the Change icon Now select the computer icon and hit OK.
