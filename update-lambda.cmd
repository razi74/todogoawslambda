set GOOS=linux
set GOARCH=amd64
echo %GOPATH%
set GOPATH=Z:\work\Gostuff\GoPathFolder;Z:\work\Gostuff\Goworkspace\todogoawslambda
go install addtask
go install gettask
go install deletetask
Z:\work\Gostuff\GoPathFolder\bin\build-lambda-zip.exe -o addtask.zip Z:\work\Gostuff\Goworkspace\todogoawslambda\bin\linux_amd64\addtask
Z:\work\Gostuff\GoPathFolder\bin\build-lambda-zip.exe -o gettask.zip Z:\work\Gostuff\Goworkspace\todogoawslambda\bin\linux_amd64\gettask
Z:\work\Gostuff\GoPathFolder\bin\build-lambda-zip.exe -o deletetask.zip Z:\work\Gostuff\Goworkspace\todogoawslambda\bin\linux_amd64\deletetask

aws lambda update-function-code --region ap-southeast-1 --function-name arn:aws:lambda:ap-southeast-1:011540970954:function:AddTaskFunction --zip-file fileb://addtask.zip

aws lambda update-function-code --region ap-southeast-1 --function-name arn:aws:lambda:ap-southeast-1:011540970954:function:GetTaskFunction --zip-file fileb://gettask.zip

aws lambda update-function-code --region ap-southeast-1 --function-name arn:aws:lambda:ap-southeast-1:011540970954:function:DeleteTaskFunction --zip-file fileb://deletetask.zip