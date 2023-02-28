#!/usr/bin/env python3
import typer
from rich import print
import subprocess

def Alert():
    print("\n:broken_heart: [bold red]Alert! The script run failed.[/bold red]\n")

def Ok():
    print("\n:beers: [bold green]Finish! The script run success.[/bold green]\n")

def RunCommand(cmd: str) -> int:
    r: subprocess.CompletedProcess = subprocess.run(cmd, shell=True)
    return r.returncode


app = typer.Typer()

@app.command()
def updateAPI(w: bool = True, m: bool = True):
    code = RunCommand("goctl api format --dir ./api")
    if code != 0:
        Alert()
        return
    if w: RunCommand("goctl api go -api ./api/worker.api -dir ./worker -style goZero")
    if m: RunCommand("goctl api go -api ./api/master.api -dir ./master -style goZero --home ./template")


@app.command()
def updateSwagger(w: bool = True, m: bool = True, port: int = 8888):
    if w: RunCommand(f"goctl api plugin -plugin goctl-swagger='swagger -filename worker.json --host 127.0.0.1:{port}' -api ./api/worker.api -dir ./worker/swagger")
    if m: RunCommand(f"goctl api plugin -plugin goctl-swagger='swagger -filename master.json --host 127.0.0.1:{port}' -api ./script/swagger/tmp_master.api -dir ./master/swagger")

@app.command()
def runSwagger(name: str = "worker-api", port: int = 8083):
    if name == "worker-api":
        RunCommand(f"docker run --rm --privileged -d -p {port}:8080 -e SWAGGER_JSON=/app/worker.json -v $PWD/worker/swagger:/app swaggerapi/swagger-ui")
    elif name == "master-api":
        RunCommand(f"docker run --rm --privileged -d -p {port}:8080 -e SWAGGER_JSON=/app/master.json -v $PWD/master/swagger:/app swaggerapi/swagger-ui")
    else:
        print("unsupported arguments:", name)

@app.command()
def version(name: str):
    print("version:", "0.0.1")

if __name__ == "__main__":
    app()
