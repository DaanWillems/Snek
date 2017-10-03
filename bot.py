import discord
from discord.ext import commands

description = '''Mainframe'''
bot = commands.Bot(command_prefix='!', description=description)

@bot.event
async def on_ready():
    print('Logged in as')
    print(bot.user.name)
    print(bot.user.id)
    print('------')

@bot.command()
async def woop():
    """Woop woop"""
    await bot.say("Woop we have a bot now!")


bot.run('MzY0MzczODExODA3OTExOTQ2.DLPUqg.tWvujqihsOdUGUsr_owzvo7Rhuo')