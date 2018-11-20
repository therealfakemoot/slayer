import configparser
import argparse

def loadConf():
    conf = configparser.ConfigParser()
    conf.read_file(open('conf.ini'))
    return conf

def setupFlags():
    p = argparse.ArgumentParser()
    p.add_argument("-b", "--board", dest="board", default="67", help="Board ID.")
    p.add_argument("-f", "--filter", dest="filter", default="RM", help="JIRA Keys to include in the weighted output.")
    p.add_argument("--debug", dest="debug", type=bool, help="Debug mode.")

    return p
