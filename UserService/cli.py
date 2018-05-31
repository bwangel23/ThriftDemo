# -*- coding: utf-8 -*-

from model import Base, engine
from rpc.server import run_server
import click


@click.group()
def cli():
    pass


@click.command()
def createdb():
    Base.metadata.create_all(engine)
    click.echo("Create the database")


@click.command()
def recreatedb():
    Base.metadata.drop_all(engine)
    click.echo("Drop the databse")
    Base.metadata.create_all(engine)
    click.echo("Create the database")


@click.command()
def runserver():
    run_server()


if __name__ == '__main__':
    cli.add_command(createdb)
    cli.add_command(recreatedb)
    cli.add_command(runserver)
    cli()

