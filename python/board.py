import base64

import requests

from conf import loadConf
from issue import Issue

CONF = loadConf()
TOKEN = base64.b64encode('{user}:{password}'.format(user=CONF['auth']['user'],password=CONF['auth']['password']).encode('UTF-8'))
AUTH = {'Authorization': 'Basic ' + TOKEN.decode('UTF-8')}

BASE = 'https://homesmediasolutions.atlassian.net/rest'
ISSUE = BASE + '/api/2/issue/'
BOARD = BASE + '/agile/1.0/board/{board}'
BACKLOG = BASE + '/agile/1.0/board/{board}/backlog'

class Board():
    def __init__(self, id):
        self.id = id

        self.update()

    def update(self):
        self._raw = requests.get(BOARD.format(board=self.id), headers=AUTH).json()

class Backlog():
    def __init__(self, board, fields):
        self.board = board
        self.fields = fields
        self._issues = {}

        self.fetch()

    def fetch(self):
        params = {}
        raw = requests.get(BACKLOG.format(board=self.board.id), headers=AUTH, params=self.fields).json()
        self.update(raw)
        pages = list(range(0, raw['total']+1, raw['maxResults']))
        for page in pages:
            params.update({'startAt': page})
            params.update(self.fields)
            raw = requests.get(BACKLOG.format(board=self.board.id), headers=AUTH, params=params).json()
            self.update(raw)

    def update(self, raw):
        for issue in raw['issues']:
            i = Issue(issue)
            self._issues[i.key] = i
        # self._issues = [Issue(i) for i in self._raw['issues']]

    def __len__(self):
        return len(self._issues)

    def __getitem__(self, idx):
        return self._issues[idx]

    def __contains__(self, key):
        return any(key == i.key for i in self._issues)

    def __iter__(self):
        for i in self._issues.values():
            yield i
