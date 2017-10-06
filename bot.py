import discord
import time
import random
from discord.ext import commands

description = '''Mainframe'''
bot = commands.Bot(command_prefix='!', description=description)

responses = ["hiss", "i bite you", "i am darkness", "heck off", "am scary cober", "i do a flat", "ill do you a chomp", "this snoot will NOT boop", "i am nice snek", "SNEK", "STAHP", "doing a shed"]

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

@bot.command()
async def woopwoop():
    """CALLING THE COPS!!!"""
    await bot.say("That's the sound of the police!")

@bot.command(pass_context=True)
async def kill(ctx):
    """Make the bot commit suicide"""
    await bot.say("Go fuck yourself "+ctx.message.author.display_name+". You monster...")

@bot.command()
async def fiteme():
    """Fight the bot"""
    await bot.say("Wot the fok did ye just say 2 me m8? i dropped out of newcastle primary skool im the sickest bloke ull ever meet & ive nicked ova 300 chocolate globbernaughts frum tha corner shop. im trained in street fitin' & im the strongest foker in tha entire newcastle gym. yer nothin to me but a cheeky lil bellend w/ a fit mum & fakebling. ill waste u and smash a fokin bottle oer yer head bruv, i swer 2 christ. ya think u can fokin run ya gabber at me whilst sittin on yer arse behind a lil screen? think again wanka. im callin me homeboys rite now preparin for a proper scrap. A roomble thatll make ur nan sore jus hearin about it. yer a waste bruv. me crew be all over tha place & ill beat ya to a proper fokin pulp with me fists wanka. if i aint satisfied w/ that ill borrow me m8s cricket paddle & see if that gets u the fok out o' newcastle ya daft kunt. if ye had seen this bloody fokin mess commin ye might a' kept ya gabber from runnin. but it seems yea stupid lil twat, innit? ima ****e fury & ull drown in it m8. ur ina proper mess knob.")

@bot.command()
async def repeat():
    """Woop woop"""
    time.sleep(2)
    await bot.say("!repeat")

@bot.command()
async def suicide():
    """Give up on life"""
    await bot.say("Do a flip!")

@bot.command()
async def hiss():
    """Talk to snek!"""
    await bot.say(random.choice(responses))






bot.run('MzY0Nzg4NjQwNzE2NTU0MjQw.DLU3aA.JxU5cqP_m2wHp4t-lqddKF2FXCA')
