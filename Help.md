



# How to Kill Port
1) open CMD/Powertshell as admistator

2) Run This Commands and port accordingly
    netstat -ano | findstr :50051
    netstat -ano | findstr :50052
    netstat -ano | findstr :8080

3) If Task is Running With PID id 
4) then RUn this Command
    taskkill /F /PID <PID>
    Example :- 
        taskkill /PID 19148 /F
        taskkill /PID 12492 /F
        taskkill /PID 18080 /F

5) Then Task Will be Killed Successfully