# Define a variable to store PIDs
PIDS_FILE=./pids.txt

run:
	@echo "Running applications..."
	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/Fno_Open_Pos && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/Fno_Ord_Dtls_Service && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/Comd_Open_Pos_Service && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/MTF_OPEN_POS && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
	

	@powershell -Command "Start-Process cmd -ArgumentList '/c cd Webserver && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
	@echo "Applications started."

stop:
	@echo "Stopping applications..."
	@powershell -Command "$pids = Get-Content -Path $(PIDS_FILE); foreach ($pid in $pids) { Stop-Process -Id [int]$pid -Force }"
	@del $(PIDS_FILE)
	@echo "Applications stopped."






# # Define a variable to store PIDs
# PIDS_FILE=pids.txt

# run:
# 	@echo "Running applications..."
# 	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/Open_Pos && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
# 	@powershell -Command "Start-Process cmd -ArgumentList '/c cd ApplicationServer/Ord_Dtls_Service && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
# 	@powershell -Command "Start-Process cmd -ArgumentList '/c cd Webserver && go run cmd/main.go' -NoNewWindow -PassThru | Select-Object -ExpandProperty Id | Out-File -Append $(PIDS_FILE)"
# 	@echo "Applications started."

# stop:
# 	@echo "Stopping applications..."
# 	@powershell -Command "$pids = Get-Content $(PIDS_FILE); foreach ($pid in $pids) { Stop-Process -Id [int]$pid -Force }"
# 	@del $(PIDS_FILE)
# 	@echo "Applications stopped."
