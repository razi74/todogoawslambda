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


rem aws lambda create-function --region ap-southeast-1 --function-name AddTaskFunction --description AddTaskFunction --zip-file fileb://addtask.zip --runtime go1.x --tracing-config Mode=Active --role arn:aws:iam::011540970954:role/devrole --handler addtask

rem aws lambda create-function --region ap-southeast-1 --function-name GetTaskFunction --description GetTaskFunction --zip-file fileb://gettask.zip --runtime go1.x --tracing-config Mode=Active --role arn:aws:iam::011540970954:role/devrole --handler gettask

rem aws lambda create-function --region ap-southeast-1 --function-name DeleteTaskFunction --description DeleteTaskFunction --zip-file fileb://deletetask.zip --runtime go1.x --tracing-config Mode=Active --role arn:aws:iam::011540970954:role/devrole --handler deletetask