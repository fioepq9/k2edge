import typer
from rich import print

def Alert():
    print("\n:broken_heart: [bold red]Alert! The script run failed.[/bold red]\n")

def Ok():
    print("\n:beers: [bold green]Finish! The script run success.[/bold green]\n")


app = typer.Typer()

@app.command()
def updateAPI(name: str = "worker-api"):
    # process
    import subprocess
    if name == "worker-api":
        r: subprocess.CompletedProcess = subprocess.run("goctl api go -api ./api/worker.api -dir ./worker -style goZero --home ./template", shell=True)
    elif name == "master-api":
        r: subprocess.CompletedProcess = subprocess.run("goctl api go -api ./api/master.api -dir ./master -style goZero --home ./template", shell=True)
    else:
        print("unsupported arguments:", name)

    # check result
    if r.returncode != 0:
        Alert()
    else:
        Ok()

@app.command()
def updateSwagger(name: str = "worker-api", port: int = 8080):
    # process
    import subprocess
    if name == "worker-api":
        r: subprocess.CompletedProcess = subprocess.run(f"goctl api plugin -plugin goctl-swagger='swagger -filename worker.json --host 127.0.0.1:{port}' -api ./api/worker.api -dir ./worker/swagger", shell=True)
    elif name == "master-api":
        r: subprocess.CompletedProcess = subprocess.run(f"goctl api plugin -plugin goctl-swagger='swagger -filename master.json --host 127.0.0.1:{port}' -api ./api/master.api -dir ./master/swagger", shell=True)
    else:
        print("unsupported arguments:", name)
    # check result
    if r.returncode != 0:
        Alert()
    else:
        Ok()

@app.command()
def runSwagger(name: str = "worker-api", port: int = 8083):
    # process
    import subprocess
    if name == "worker-api":
        r: subprocess.CompletedProcess = subprocess.run(f"docker run --rm -p {port}:8080 -e SWAGGER_JSON=/app/worker.json -v $PWD/worker/swagger:/app swaggerapi/swagger-ui", shell=True)
    elif name == "master-api":
        r: subprocess.CompletedProcess = subprocess.run(f"docker run --rm -p {port}:8080 -e SWAGGER_JSON=/app/master.json -v $PWD/master/swagger:/app swaggerapi/swagger-ui", shell=True)
    else:
        print("unsupported arguments:", name)
    # check result
    if r.returncode != 0:
        Alert()
    else:
        Ok()


@app.command()
def version(name: str):
    print("version:", "0.0.1")

if __name__ == "__main__":
    app()
