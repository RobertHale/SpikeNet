REM // NPM Install first
echo "npm install"
cd C:\Users\rthsp\GolandProjects\SpikeNet\web\SpikeNet
call npm config set strict-ssl=false
call npm install

REM // Angular build
start "NG Build" cmd /k "echo on & cd C:\Users\rthsp\GolandProjects\SpikeNet\web\SpikeNet & ng build --watch --output-path C:\Users\rthsp\GolandProjects\SpikeNet\build\dist\local\spikenet"

REM // webapp
start "SpikeNet" cmd /k "echo on & cd C:\Users\rthsp\GolandProjects\SpikeNet\webapp & go run main.go"