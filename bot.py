import discord
import time
from discord.ext import commands

description = '''Mainframe'''
bot = commands.Bot(command_prefix='!', description=description)

@bot.event
async def on_ready():
    print('Logged in as')
    print(bot.user.name)
    print(bot.user.id)
    print('------')

@bot.command(pass_context=True)
async def woop(ctx):
    """Woop woop"""
    await bot.say(ctx.message.author.display_name+ " is excited!")

@bot.command(pass_context=True)
async def kill(ctx):
    """Make the bot commit suicide"""
    await bot.say("Go fuck yourself "+ctx.message.author.display_name+". You monster...")

@bot.command()
async def repeat():
    """Woop woop"""
    time.sleep(2)
    await bot.say("!repeat")


bot.run('MzY0MzczODExODA3OTExOTQ2.DLPUqg.tWvujqihsOdUGUsr_owzvo7Rhuo')