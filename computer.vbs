'0 = OK Button,' 
'1 = OK / Cancel Button, 
'2 = Abort / Retry / Ignore Button, 
'3 = Yes / No / Cancel Button, 
'4 = Yes / No Button, 
'5 = Retry / Cancel Button'


X=MsgBox("Error while opening Computer. Do you want to Fix this Error?",4+64,"Computer")
X=MsgBox("Unable to Fix this Error. Do you want to scan this Computer",3+48,"Computer")
X=MsgBox("Alert ! Virus Detected. Delete Virus? ",3+16,"Alert!")
X=MsgBox("Unable to delete this Virus",1+64,"Critical Error")
X=MsgBox("Virus Is activated",2+16,"Alert")
X=MsgBox("Deleling System Files.....",2+16,"File Deletion")
X=MsgBox("Virus is copying your password.....",2+48,"Virus Alert")
X=MsgBox("Please Wait, Uploading your data to Server. Do you want to Stop it",4+64,"File Transfer")
X=MsgBox("Could not stop. File Transfer Completed",1+16,"Completed")
X=MsgBox("Your Computer is Hacked",0+64,"Alert")
X=MsgBox("HAHA This was Prank",0+64,"Fooled You")
