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
        r: subprocess.CompletedProcess = subprocess.run("goctl api go -api ./api/worker.api -dir ./worker -style goZero")
    elif name == "master-api":
        r: subprocess.CompletedProcess = subprocess.run("goctl api go -api ./api/master.api -dir ./master -style goZero")
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
